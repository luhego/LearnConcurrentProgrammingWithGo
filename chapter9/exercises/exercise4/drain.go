package exercise4

func Drain[T any](quit <-chan int, input <-chan T) {
	go func() {
		moreData := true
		for moreData {
			select {
			case _, moreData = <-input:
				// drain message
			case <-quit:
				return
			}
		}
	}()
}
