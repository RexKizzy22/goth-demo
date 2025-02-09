package model

type Pagination struct {
	NoOfPages    int
	PageOffset   int
	NoOfSlides   int
	SlideOffset  int
	CurrentSlide int
}

func (p *Pagination) HasNextSlide() bool {
	return p.CurrentSlide+1 <= p.NoOfSlides
}

func (p *Pagination) HasPrevSlide() bool {
	return p.CurrentSlide > 1
}

type Row struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

type Rows = []Row

type State struct {
	Rows       Rows
	Pagination *Pagination
}
