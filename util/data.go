package util

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RexKizzy22/goth-demo/model"
)

func ReadCsv() ([][]string, error) {
	f, err := os.Open("data/table.csv")
	if err != nil {
		return nil, err
	}

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return rows[1:], nil
}

func ParseStocks(stocks *[][]string) (rows model.Rows) {
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

func Paginate(r *http.Request, size int) int {
	p, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		p = 1
	}
	o := generateOffset(p, size)

	return o
}

func generateOffset(page int, size int) int {
	return (page - 1) * size
}
