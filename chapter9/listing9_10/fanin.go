package listing9_10

import "sync"

func FanIn[K any](quit <-chan int, allChannels ...<-chan K) chan K {
	wg := sync.WaitGroup{} // Creates a waitgroup with size equal to the number of channels
	wg.Add(len(allChannels))
	output := make(chan K) // Creates the output channel
	for _, c := range allChannels {
		go func(channel <-chan K) { // Starts goroutine for each input channel
			defer wg.Done() // Mark the waitgroup as Done when the goroutine terminates
			for i := range channel {
				select {
				case output <- i: // Forward each received message to the shared output channel
				case <-quit: // Terminate the procedure if the quit channel is closed
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
