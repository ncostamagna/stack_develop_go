package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

const numberOfThreads int = 60

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup = sync.WaitGroup{}
)

func findArea(inputChannel chan string, i int) {
	for pointsStr := range inputChannel {
		var points []Point2D
		for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		value := math.Abs(area) / 2.0

		if value < 10000 {
			// Solo se frena un thread de los 8
			fmt.Printf("%f %d STOPPING\n", value, i)
			time.Sleep(1 * time.Second)
			fmt.Printf("%f %d CONTINUE\n", value, i)
		} else {
			fmt.Printf("%f %d\n", value, i)
		}

	}
	waitGroup.Done()
}

func main() {
	absPath, _ := filepath.Abs("./threadpool/")
	dat, _ := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(dat)

	inputChannel := make(chan string, 1000)

	// definimos que vamos a crear 8 threads
	for i := 0; i < numberOfThreads; i++ {
		go findArea(inputChannel, i)
	}
	waitGroup.Add(numberOfThreads)
	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s \n", elapsed)
}
