package gocoroutine

import (
	"errors"
	"fmt"
	"time"
)


func doSomeWork() (int, error){
	fmt.Println("Do Some API calls!");
	time.Sleep(3 * time.Second);
	return 10, nil;
}

// Now issue with is func, Even TimeOut happen and retun the error, 
// Goroutine will not exits, It will still there, This leads to a Goroutine leak.
// Fixed version is "TimeOutCodeV2"
func TimeOutCode() (int, error){
	var result int;
	var err error;
	//
	done := make(chan struct{});
	go func ()  {
		result, err = doSomeWork();
		if err != nil {
			fmt.Println("Opps! Error Happen!");
		}
		fmt.Println("The Result is: ", result);
		close(done);
	}();

	select {
	case <- done:
		return result, err;
	case <-time.After(2 * time.Second):
		return 0, errors.New("opps timeout");
	}
}

func TimeOutCodeV2() (int, error){
	var result int;
	var err error;
	cancel := make(chan struct{});
	//
	done := make(chan struct{});
	go func ()  {
		select {
		case <-cancel:
			return;
		default:
			result, err = doSomeWork();
			if err != nil {
				fmt.Println("Opps! Error Happen!");
			}
			fmt.Println("The Result is: ", result);			
			close(done);
		}
	}();
	//
	select {
	case <- done:
		return result, err;
	case <-time.After(2 * time.Second):
		close(cancel);
		return 0, errors.New("opps timeout");
	}
}