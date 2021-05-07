package main

import (
	"education/images/pkg/converter"
	"flag"
	"fmt"
	"sync"
)

var (
	maxWorkers           int = 5
	sourceDirectory      string
	destinationDirectory string
	wg                   sync.WaitGroup
)

//third parametr chan Convertible
func spawnWorkers(n int, wg *sync.WaitGroup, ch chan converter.Converter, destinationDirectory string) {
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go converter.Worker(wg, ch, destinationDirectory)
	}
}

func main() {
	flag.StringVar(&sourceDirectory, "Source Directory", sourceDirectory, "s")
	flag.StringVar(&destinationDirectory, "Target Directory", destinationDirectory, "d")
	flag.IntVar(&maxWorkers, "Number of workers spawned", maxWorkers, "n")
	flag.Parse()

	imagePool := make(chan converter.Converter, maxWorkers)
	//We will take path of the  source directory and read all files with walkpath. We'll ignore directories
	spawnWorkers(maxWorkers, &wg, imagePool, destinationDirectory)
	converter.GetImages(sourceDirectory, imagePool)
	//WAitGroup
	wg.Wait()
	fmt.Println("Done")
}
