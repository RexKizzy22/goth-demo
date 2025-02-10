package util

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RexKizzy22/goth-demo/model"
)

const DEFAULT_PAGE_SIZE = 15
const PAGES_PER_SLIDE = 10

func LoadAppState() model.State {
	stocks, err := readCsv("data/table.csv")
	if err != nil {
		log.Fatal(err)
	}

	state := model.State{
		Rows: parseStocks(&stocks),
	}

	return state
}

func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return rows[1:], nil
}

func parseStocks(stocks *[][]string) (rows model.Rows) {
	for _, stock := range *stocks {
		// date, _ := time.Parse(time.DateOnly, stock[0])
		open, _ := strconv.ParseFloat(stock[1], 64)
		high, _ := strconv.ParseFloat(stock[2], 64)
		low, _ := strconv.ParseFloat(stock[3], 64)
		close, _ := strconv.ParseFloat(stock[4], 64)
		adj, _ := strconv.ParseFloat(stock[6], 64)

		rows = append(rows, model.Row{
			// Date:     fmt.Sprint(date.Date()),
			Date:     stock[0],
			Open:     fmt.Sprintf("%.2f", open),
			High:     fmt.Sprintf("%.2f", high),
			Low:      fmt.Sprintf("%.2f", low),
			Close:    fmt.Sprintf("%.2f", close),
			Volume:   stock[5],
			AdjClose: fmt.Sprintf("%.2f", adj),
		})
	}

	return rows
}

func ParseForm(r *http.Request) model.Row {
	date := time.Now().Format(time.DateOnly)
	open, _ := strconv.ParseFloat(r.FormValue("open"), 64)
	high, _ := strconv.ParseFloat(r.FormValue("high"), 64)
	low, _ := strconv.ParseFloat(r.FormValue("low"), 64)
	close, _ := strconv.ParseFloat(r.FormValue("close"), 64)
	volume := r.FormValue("volume")
	adjClose, _ := strconv.ParseFloat(r.FormValue("adjClose"), 64)

	return model.Row{
		Date:     date,
		Open:     fmt.Sprintf("%.2f", open),
		High:     fmt.Sprintf("%.2f", high),
		Low:      fmt.Sprintf("%.2f", low),
		Close:    fmt.Sprintf("%.2f", close),
		Volume:   volume,
		AdjClose: fmt.Sprintf("%.2f", adjClose),
	}
}

func FilterStocks(stocks model.Rows, search string) (filteredStocks model.Rows) {
	for _, s := range stocks {
		if strings.Contains(s.Date, search) {
			filteredStocks = append(filteredStocks, s)
		}
	}

	return
}

func Paginate(r *http.Request, size int, rl int) *model.Pagination {
	pg := &model.Pagination{}
	p, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		p = 1
	}

	po := generatePageOffset(p, size)
	nop := generateNumberOfPages(rl, size)

	if rl < size {
		return &model.Pagination{
			CurrentSlide: 1,
			PageOffset:   po,
			NoOfSlides:   1,
			SlideOffset:  1,
			NoOfPages:    nop,
		}
	}

	var s int
	if nop > PAGES_PER_SLIDE {
		s = PAGES_PER_SLIDE
	} else {
		s = nop
	}
	nos := calculateTotalPagesOrSlides(nop, s)
	pg = generateSlideStats(pg)
	pg.NoOfSlides = nos
	pg.NoOfPages = nop
	pg.PageOffset = po
	return pg
}

func generateNumberOfPages(rl int, size int) (nop int) {
	if rl > DEFAULT_PAGE_SIZE {
		nop = calculateTotalPagesOrSlides(rl, size)
	} else {
		nop = 1
	}
	return
}

func generatePageOffset(page int, size int) int {
	return (page - 1) * size
}

func calculateTotalPagesOrSlides(total, size int) (p int) {
	var r int

	if total >= size {
		r = total % size
		p = total / size
	} else {
		r = 0
		p = total
	}

	if r > 0 {
		return p + 1
	} else {
		return p
	}
}

func generateSlideStats(p *model.Pagination) *model.Pagination {
	if p.CurrentSlide == 0 {
		p.CurrentSlide = 1
	}

	return setSlideOffset(p)
}

func setSlideOffset(p *model.Pagination) *model.Pagination {
	if p.CurrentSlide == 1 {
		p.SlideOffset = p.CurrentSlide
	} else {
		p.SlideOffset = (PAGES_PER_SLIDE * (p.CurrentSlide - 1)) + 1
	}

	return p
}
