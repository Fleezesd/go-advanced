package Concurrency

import (
	"sync"

	"golang.org/x/sync/errgroup"
)

func Cities(cities ...string) ([]*Info, error) {
	var g errgroup.Group
	var mu sync.Mutex
	res := make([]*Info, len(cities)) // res[i] corresponds to cities[i]

	for i, city := range cities {
		i, city := i, city // create locals for closure below
		g.Go(func() error {
			info, err := City(city)
			mu.Lock()
			res[i] = info
			mu.Unlock()
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}
