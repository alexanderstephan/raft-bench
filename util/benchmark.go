package util

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	numKeys   = 1
	mil       = 1000000
	runs      = 10
	wait      = 3000
	firstWait = 10000
)

func Bench(test bool, logFile string, read func(string) bool, write func(string, string) bool) {
	defer WaitForCtrlC()
	if !test {
		return
	}

	f, err := os.Create(logFile)
	if err != nil {
		log.Fatal("unable to create csv log")
	}
	defer f.Close()

	time.Sleep(firstWait)
	log.Printf("Starting benchmark...\n")
	for i := 0; i < runs; i++ {
		time.Sleep(wait)

		start := time.Now()
		k := 0
		success := 0
		for k < numKeys*mil {
			v := rand.Int()
			go func() {
				ok := write(string(k), string(v))
				if ok {
					success++
				}
			}()
			k += 1
		}
		_, _ = f.WriteString(fmt.Sprintf("write,%v,%v,%v,%v", i+1, success, numKeys*mil, time.Since(start)))

		time.Sleep(wait)
		start = time.Now()
		k = 0
		success = 0
		for k < numKeys*mil {
			go func() {
				ok := read(string(k))
				if ok {
					success++
				}
			}()
			k += 1
		}
		_, _ = f.WriteString(fmt.Sprintf("read,%v,%v,%v,%v", i+1, success, numKeys*mil, time.Since(start)))
	}
	log.Printf("BENCHMARK COMPLETE\n")
}
