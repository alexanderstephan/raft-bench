package util

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numKeys = 1
	mil     = 1000000
	runs    = 3
	wait    = 3000
)

func Bench(test bool, read func(string), write func(string, string)) {
	defer WaitForCtrlC()
	if !test {
		return
	}

	for i := 0; i < runs; i++ {
		time.Sleep(wait)

		start := time.Now()
		k := 0
		for k < numKeys*mil {
			v := rand.Int()
			go write(string(k), string(v))
			k += 1
		}
		fmt.Printf("Write test, %v, %v, %v\n", i+1, numKeys*mil, time.Since(start))

		time.Sleep(wait)
		start = time.Now()
		k = 0
		for k < numKeys*mil {
			go read(string(k))
			k += 1
		}
		fmt.Printf("Read test, %v, %v, %v\n", i+1, numKeys*mil, time.Since(start))
	}
}