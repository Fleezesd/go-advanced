package Concurrency

import (
	"fmt"

	"golang.org/x/sync/singleflight"
)

// reference: https://encore.dev/blog/advanced-go-concurrency

type Info struct {
	TempC, TempF int    // temperature in Celsius and Farenheit
	Conditions   string // "sunny", "snowing", etc
}

var group singleflight.Group

// City 同一时间内并发请求相同city结果相同, 减少DB损耗
func City(city string) (*Info, error) {
	// group.Do 减少高并发请求下相同key的重复工作
	results, err, _ := group.Do(city, func() (interface{}, error) {
		info, err := fetchWeatherFromDB(city) // slow operation
		return info, err
	})
	if err != nil {
		return nil, fmt.Errorf("weather.City %s: %w", city, err)
	}
	return results.(*Info), nil
}

func fetchWeatherFromDB(city string) (*Info, error) {
	info := &Info{
		TempC:      20,
		TempF:      68,
		Conditions: "sunny",
	}
	return info, nil
}
