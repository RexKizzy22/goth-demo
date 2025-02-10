package handler

import (
	"net/http"
	"slices"

	"github.com/RexKizzy22/goth-demo/components"
	"github.com/RexKizzy22/goth-demo/model"
	"github.com/RexKizzy22/goth-demo/toast"
	"github.com/RexKizzy22/goth-demo/util"
)

const DEFAULT_PAGE_SIZE = 15

type Handler struct {
	state *model.State
}

func New(state *model.State) *Handler {
	return &Handler{
		state: state,
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.state.Pagination = util.Paginate(r, DEFAULT_PAGE_SIZE, h.state.Rows)
	s := h.state
	p := h.state.Pagination
	components.Page(s.Rows[p.PageOffset:(p.PageOffset+DEFAULT_PAGE_SIZE)], p).Render(r.Context(), w)
}

func (h *Handler) Table(w http.ResponseWriter, r *http.Request) {
	p := util.Paginate(r, DEFAULT_PAGE_SIZE, h.state.Rows)

	rr := len(h.state.Rows) - p.PageOffset
	if rr > DEFAULT_PAGE_SIZE {
		components.Table(h.state.Rows[p.PageOffset:(p.PageOffset+DEFAULT_PAGE_SIZE)]).Render(r.Context(), w)
	} else {
		components.Table(h.state.Rows[p.PageOffset:(p.PageOffset+rr)]).Render(r.Context(), w)
	}
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("search")
	if s == "" {
		http.Redirect(w, r, "/table", http.StatusSeeOther)
		return
	}

	rows := h.state.Rows
	p := util.Paginate(r, DEFAULT_PAGE_SIZE, rows)
	ns := util.FilterStocks(rows, s)

	if len(ns) <= 0 {
		components.NotFound(s).Render(r.Context(), w)
	} else {
		if len(ns) > DEFAULT_PAGE_SIZE {
			components.Table(ns[p.PageOffset:(p.PageOffset+DEFAULT_PAGE_SIZE)]).Render(r.Context(), w)
		} else {
			components.Table(ns).Render(r.Context(), w)
		}
	}
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
