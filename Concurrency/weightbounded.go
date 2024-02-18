package Concurrency

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

func WeightBoundedCities(cities ...string) ([]*Info, error) {
	ctx := context.TODO() // replace with a real context
	var g errgroup.Group
	var mu sync.Mutex
	res := make([]*Info, len(cities))
	sem := semaphore.NewWeighted(100) // 100 chars processed concurrently
	for i, city := range cities {
		i, city := i, city
		cost := int64(len(city))
		if err := sem.Acquire(ctx, cost); err != nil {
			break
		}
		g.Go(func() error {
			info, err := City(city)
			mu.Lock()
			res[i] = info
			mu.Unlock()
			sem.Release(cost)
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	} else if err := ctx.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

/*
	errgroup.Group：用于并发执行任务，同时捕获并返回第一个遇到的错误, 进行记录，至于对于任务是否停止，是通过context来禁止
	sync.Mutex：确保对共享资源（这里是结果切片res）的访问是线程安全的。
	semaphore.Weighted：通过加权信号量控制并发的总成本，这里以城市名称的长度作为成本度量，限制了同时处理的城市名称总长度不超过100。这是一个灵活的方式来限制基于资源预算的并发，比如防止过多的数据库连接。
	上下文（context.Context）：提供了一种机制来控制和取消长时间运行的操作。在实际应用中，应该使用具有超时或取消能力的上
*/
