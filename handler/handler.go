package handler

import (
	"net/http"
	"slices"

	"github.com/RexKizzy22/goth-demo/components"
	"github.com/RexKizzy22/goth-demo/model"
	"github.com/RexKizzy22/goth-demo/toast"
	"github.com/RexKizzy22/goth-demo/util"
)

type Handler struct {
	state *model.State
}

func New(state *model.State) *Handler {
	return &Handler{
		state: state,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	l := len(h.state.Rows)
	h.state.Pagination = util.Paginate(r, util.DEFAULT_PAGE_SIZE, l)
	s := h.state
	p := h.state.Pagination
	components.Page(s.Rows[p.PageOffset:(p.PageOffset+util.DEFAULT_PAGE_SIZE)], p).Render(r.Context(), w)
}

func (h *Handler) Table(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("search")
	if s == "" {
		p := util.Paginate(r, util.DEFAULT_PAGE_SIZE, len(h.state.Rows))

		rr := len(h.state.Rows) - p.PageOffset
		if rr > util.DEFAULT_PAGE_SIZE {
			components.Table(h.state.Rows[p.PageOffset:(p.PageOffset+util.DEFAULT_PAGE_SIZE)]).Render(r.Context(), w)
		} else {
			components.Table(h.state.Rows[p.PageOffset:(p.PageOffset+rr)]).Render(r.Context(), w)
		}
		return
	}

	p := util.Paginate(r, util.DEFAULT_PAGE_SIZE, len(h.state.SearchResults))

	rr := len(h.state.SearchResults) - p.PageOffset
	if rr > util.DEFAULT_PAGE_SIZE {
		components.Table(h.state.SearchResults[p.PageOffset:(p.PageOffset+util.DEFAULT_PAGE_SIZE)]).Render(r.Context(), w)
	} else {
		components.Table(h.state.SearchResults[p.PageOffset:(p.PageOffset+rr)]).Render(r.Context(), w)
	}
}

func (h *Handler) Main(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("search")
	if s == "" {
		http.Redirect(w, r, "/refresh", http.StatusSeeOther)
		return
	}

	rows := util.FilterStocks(h.state.Rows, s)
	h.state.SearchResults = rows
	l := len(rows)
	p := util.Paginate(r, util.DEFAULT_PAGE_SIZE, l)

	if l == 0 {
		components.Main(rows, p, s).Render(r.Context(), w)
	} else {
		rr := l - p.PageOffset
		if rr > util.DEFAULT_PAGE_SIZE {
			components.Main(rows[p.PageOffset:(p.PageOffset+util.DEFAULT_PAGE_SIZE)], p, s).Render(r.Context(), w)
		} else {
			components.Main(rows[p.PageOffset:(p.PageOffset+rr)], p, s).Render(r.Context(), w)
		}
	}
}

func (h *Handler) RefreshMain(w http.ResponseWriter, r *http.Request) {
	l := len(h.state.Rows)
	h.state.Pagination = util.Paginate(r, util.DEFAULT_PAGE_SIZE, l)
	s := h.state
	p := h.state.Pagination
	components.Main(s.Rows[p.PageOffset:(p.PageOffset+util.DEFAULT_PAGE_SIZE)], p, "").Render(r.Context(), w)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	nr := util.ParseForm(r)
	slices.Reverse(h.state.Rows)
	h.state.Rows = append(h.state.Rows, nr)
	slices.Reverse(h.state.Rows)

	json := toast.Jsonify(toast.SUCCESS, "Successful!", "Index was successfully updated!")
	w.Header().Set("HX-Trigger", json)
	components.TableRow(nr).Render(r.Context(), w)
}

func (h *Handler) NextSlide(w http.ResponseWriter, r *http.Request) {
	p := h.state.Pagination
	if p.HasNextSlide() {
		p.CurrentSlide = p.CurrentSlide + 1
		p.SlideOffset = p.SlideOffset + util.PAGES_PER_SLIDE
	}

	components.Pagination(p).Render(r.Context(), w)
}

func (h *Handler) PrevSlide(w http.ResponseWriter, r *http.Request) {
	p := h.state.Pagination
	if p.HasPrevSlide() {
		p.CurrentSlide = p.CurrentSlide - 1
		p.SlideOffset = p.SlideOffset - util.PAGES_PER_SLIDE
	}

	components.Pagination(p).Render(r.Context(), w)
}
