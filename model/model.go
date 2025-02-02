package model

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
	Rows Rows
}