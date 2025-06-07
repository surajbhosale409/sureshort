package pkg

import (
	"fmt"
	"sort"
	"sync"
)

type Stats struct {
	valueFrequency map[string]int
	mux            *sync.RWMutex
}

func NewStats() *Stats {
	return &Stats{
		valueFrequency: make(map[string]int),
		mux:            &sync.RWMutex{},
	}
}

func (s *Stats) Observe(value string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.valueFrequency[value]++
}

func (s *Stats) Top(n int) []string {
	s.mux.RLock()
	defer s.mux.RUnlock()

	type kv struct {
		Value     string
		Frequency int
	}

	var list []kv
	for value, frequency := range s.valueFrequency {
		list = append(list, kv{value, frequency})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Frequency > list[j].Frequency
	})

	top := []string{}
	for i := 0; i < len(list) && i < n; i++ {
		top = append(top, fmt.Sprintf("%s: %d", list[i].Value, list[i].Frequency))
	}
	return top
}
