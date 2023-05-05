package gocoroutine

import (
	"sync"
)


func WaitGropWithChan(in <-chan int, processor func(int) int, num int) []int {
	out := make(chan int, num); // define buffered channel
	var wg sync.WaitGroup;
	wg.Add(num); // Tell how many time need to wait.
	//
	for i := 0; i < num; i++ {
		go func ()  {
			defer wg.Done();
			for v := range in {
				out <- processor(v);
			}
		}();
	}
	// After all wait groups finished, This will execute and close the out channel.
	go func ()  {
		wg.Wait();
		close(out);
	}();

	var result []int;
	for v := range out { // After closed the out channel this range loop will terminated.
		result = append(result, v);
	}
	return result;
}



