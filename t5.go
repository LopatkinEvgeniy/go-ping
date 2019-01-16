package main

import (
	"fmt"
	"sync"
	"sort"
)


func main() {
	contents := []string{"172.20.13.136 - HOST ACTIVE", "172.20.13.130 - HOST ACTIVE", "172.20.13.137 - HOST ACTIVE", "172.20.13.135 - HOST DISABLED", "172.20.13.131 - HOST ACTIVE", "172.20.13.132 - HOST ACTIVE"}
	sampleChan := make(chan string, 1)
	var sampleList []string
	var wg sync.WaitGroup
	for _, line := range contents {
		wg.Add(1)

		go func(line string, ch chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- line

		}(line, sampleChan, &wg)
	}
	go func (channel chan string, sampleList *[]string) {
		for s := range channel {
			*sampleList = append(*sampleList, s)
		}

	}(sampleChan, &sampleList)
	wg.Wait()
	sort.Strings(sampleList)
	fmt.Println(sampleList)
}
