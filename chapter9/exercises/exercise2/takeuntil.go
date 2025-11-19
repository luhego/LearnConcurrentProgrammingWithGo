package exercise2

func TakeUntil[K any](f func(K) bool, quit chan int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		predicate := true
		moreData := true
		var msg K
		for predicate && moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					predicate = f(msg)
					if predicate {
						output <- msg
					}
				}
			case <-quit:
				return
			}

		}
		if !predicate {
			close(quit)
		}
	}()
	return output
}
