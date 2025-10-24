package example

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Concurrency() {
	path := flag.String("file", "example/url.txt", "path to URL file")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	urlSlice := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)
	for _, url := range urlSlice {
		go ping(url, respCh, errCh)
	}
	for range urlSlice {
		select {
		case err := <-errCh:
			fmt.Println(err)
		case res := <-respCh:
			fmt.Println(res)
		}
	}
}

func ping(url string, respCh chan int, errCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	respCh <- resp.StatusCode
}

// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
// numGoroutines := 3
// arrPartLen := len(arr) / numGoroutines
// arrChunks := slices.Chunk(arr, arrPartLen)
// sumCh := make(chan int, numGoroutines)
// for chunk := range arrChunks {
// 	go sumPart(chunk, sumCh)
// }
// var finalSum int
// for range numGoroutines {
// 	finalSum += <-sumCh
// }
// fmt.Printf("Final sum: %d\n", finalSum)

// code := make(chan int)
// var wg sync.WaitGroup
// for range 10 {
// 	wg.Go(func() {
// 		getHttpCode(code)
// 	})
// }
// go func() {
// 	wg.Wait()
// 	close(code)
// }()
// for res := range code {
// 	fmt.Printf("Status code: %d\n", res)
// }
