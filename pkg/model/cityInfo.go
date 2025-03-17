package model

import "sync"

type City struct {
	MeanSum float32
	Min     float32
	Max     float32
	Count   float32
	Name    string
	Lock    *sync.Mutex
}
