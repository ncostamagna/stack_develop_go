package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metarClose    = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

func parseToArray(textChannel chan string, metarChannel chan []string) {
	fmt.Println("Entra parseToArray")
	// recorremos todos los valores que se van agregando al canal
	// cuando termina (se ejecuta el close), sale del for
	for text := range textChannel {
		fmt.Println("parseToArray 		---->		textChannel ")
		lines := strings.Split(text, "\n")
		metarSlice := make([]string, 0, len(lines))
		metarStr := ""
		for _, line := range lines {
			if tafValidation.MatchString(line) {
				break
			}
			if !comment.MatchString(line) {
				metarStr += strings.Trim(line, " ")
			}
			if metarClose.MatchString(line) {
				metarSlice = append(metarSlice, metarStr)
				metarStr = ""
			}
		}

		// ponemos valor en el canal para que lo tome el otro thread
		metarChannel <- metarSlice
	}
	fmt.Println("Close channel")
	close(metarChannel)
}

func extractWindDirection(metarChannel chan []string, windsChannel chan []string) {
	fmt.Println("Entra extractWindDirection")
	for metars := range metarChannel {
		fmt.Println("extractWindDirection 		---->		metarChannel")
		winds := make([]string, 0, len(metars))
		for _, metar := range metars {
			if windRegex.MatchString(metar) {
				winds = append(winds, windRegex.FindAllStringSubmatch(metar, -1)[0][1])
			}
		}
		// ponemos valor en el canal para que lo tome el otro thread
		windsChannel <- winds
	}
	fmt.Println("Close channel")
	close(windsChannel)
}

func mineWindDistribution(windsChannel chan []string, distChannel chan [8]int) {
	fmt.Println("Entra mineWindDistribution")
	for winds := range windsChannel {
		fmt.Println("mineWindDistribution 		---->		windsChannel")
		for _, wind := range winds {
			if variableWind.MatchString(wind) {
				for i := 0; i < 8; i++ {
					windDist[i]++
				}
			} else if validWind.MatchString(wind) {
				windStr := windDirOnly.FindAllStringSubmatch(wind, -1)[0][1]
				if d, err := strconv.ParseFloat(windStr, 64); err == nil {
					dirIndex := int(math.Round(d/45.0)) % 8
					windDist[dirIndex]++
				}
			}
		}
	}
	fmt.Println("Close channel")
	distChannel <- windDist
	close(distChannel)
}

func main() {
	textChannel := make(chan string)
	metarChannel := make(chan []string)
	windsChannel := make(chan []string)
	resultsChannel := make(chan [8]int)

	// Los corro en hilos y les voy pasando los valores en el for
	//1. Change to array, each metar report is a separate item in the array
	// extrae info y sigue al siguiente paso
	go parseToArray(textChannel, metarChannel)

	//2. Extract wind direction, EGLL 312350Z 07004KT CAVOK 12/09 Q1016 NOSIG= -> 070
	// extrae info y sigue al siguiente paso
	go extractWindDirection(metarChannel, windsChannel)

	//3. Assign to N, NE, E, SE, S, SW, W, NW, 070 -> E + 1
	// extrae info y sigue al siguiente paso
	go mineWindDistribution(windsChannel, resultsChannel)

	absPath, _ := filepath.Abs("./metarfiles/")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()
	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(dat)

		// agregamos el valor al canal para que se vaya procesando
		textChannel <- text
		//time.Sleep(time.Second)
	}
	close(textChannel) // hay que cerrar el canal para indicar que finalizo y se cierren los threads
	results := <-resultsChannel
	elapsed := time.Since(start)
	fmt.Printf("%v\n", results)
	fmt.Printf("Processing took %s\n", elapsed)
}
