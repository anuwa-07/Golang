package gocoroutine
import (
	"fmt"
)

// The Done Channel Pattern
func searchData(rangeCount int, done *chan struct{}) <-chan int {
	ch := make(chan int);
	go func ()  {
		for i := 0; i < rangeCount; i++ {
			select {
			case <-*done:
				{
				return;					
				}
			default:
				ch <- i; // Writing the value to channel
			}
		}
		// After the writting is done on the channel, We need to close the channel.
		// If not that will reason for a DeadLock.
		fmt.Println("All looping is Finished in Goroutine!");
		defer close(ch);
	}();
	// retun the channel, To read values from another goroutine.
	return ch;
}

func ExecuteSearchData() {
	done := make(chan struct{});
	//
	for val := range searchData(10, &done) {
		fmt.Println("values written to the Channel: ", val);
		if val >= 5 {
			fmt.Println("Values are grater than 5, Need to Exits");
			break;
		}
	}
	close(done);
	//
	fmt.Println("All Proccess Finished!");
}



// 
func CountTo(max int) (<- chan int, func()) {
	ch := make(chan int);
	done := make(chan struct{});

	// Define the callback
	cancel := func ()  {
		close(done);
	}
	//
	go func ()  {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				{
					fmt.Println("Fuck!");
					return;					
				}
			default:
				ch <- i;
			}
		}
		close(ch);
	}();
	return ch, cancel
}

//
func CountToMain(){
	ch, cancel := CountTo(10);
	for i := range ch {
		if i > 5 {
			break;
		}
		fmt.Println("Value is: ", i);
	}
	cancel() // trigger the callback function.
}

