package Concurrency

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

// BoundedCities 有限并发
func BoundedCities(cities ...string) ([]*Info, error) {
	var g errgroup.Group
	var mu sync.Mutex
	res := make([]*Info, len(cities)) // res[i] corresponds to cities[i]
	sem := make(chan struct{}, 10)    // 通过channel来做有限并发
	for i, city := range cities {
		i, city := i, city // create locals for closure below
		sem <- struct{}{}
		g.Go(func() error {
			info, err := City(city)
			mu.Lock()
			res[i] = info
			mu.Unlock()
			<-sem
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}
