package model

import "sync"

type LabubuStore struct {
	sync.Mutex
	StockName string
	Count int
}