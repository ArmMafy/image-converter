package main

import (
	"education/images/pkg/converter"
	"flag"
	"fmt"
	"sync"
)

var (
	maxWorkers           int = 2
	sourceDirectory      string
	destinationDirectory string
	wg                   sync.WaitGroup
)

func spawnWorkers(n int, wg *sync.WaitGroup, ch chan *converter.Converter, destinationDirectory string) {
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go converter.Worker(wg, ch, destinationDirectory)
		fmt.Println("Workers:")
		fmt.Println(i)
	}
	fmt.Println("Workers spawned")
}

func main() {
	flag.StringVar(&sourceDirectory, "s", sourceDirectory, "Source Directory")
	flag.StringVar(&destinationDirectory, "d", destinationDirectory, "Destination Directory")
	flag.IntVar(&maxWorkers, "n", maxWorkers, "Quantity of Workers")
	flag.Parse()
	fmt.Println(maxWorkers)
	fmt.Println("Starting")
	// Channel for sharing readed files among workers
	imagePool := make(chan *converter.Converter, maxWorkers)

	// Creates workers that takes files from imagePool channel, converts their colors to grey, encodes new converted versions of images to destination directory
	spawnWorkers(maxWorkers, &wg, imagePool, destinationDirectory)
	// Read files from directory, ignores subdirectories and any other file that doesn't have jpeg or png extension, sends images to channel
	converter.GetImages(sourceDirectory, imagePool)
	wg.Wait()
	fmt.Println("Done")
}
