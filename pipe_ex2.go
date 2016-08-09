package main

import "fmt"


func sumFiles1() {
	c := make(chan result)
	errc := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				select {
				case <- done: return
				case c<-hanle():
				}
			}()

			select {
			case <-done: return
			default: 
			}
		}

		go func() {
			wg.Wait()
			close(c)
		}()

		errc <- err
	}()

	return c, errc
}

func walkFiles() {
	paths := make(chan string)
	errc := make(chan error, 1)

	// do sth
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func() {
			select {
			case <-done: return
			case paths <- path:
			}
		}
		)
	}()

	return paths, errc
}

func digest(done <-chan struct{}, paths, result) {
	for path := range paths {
		// do sth
		select {
		case <-done: return
		case result<-path: // do sth
		}
	}
}

func MD5All(root string) () {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)
	
	cnt := 20
	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			digest(done)
			wg.Done()
		} ()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c, errc
}

func main() {
	m, err := MD5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// var path []string
	// for path := range m {
	// }

	// print
}
