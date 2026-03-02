package model

import "sync"

type LabubuStore struct {
	StockName string
	Count     int
	Mu        sync.Mutex
}