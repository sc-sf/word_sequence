package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Words string
	Count int
}

type Pairs []Pair

// Implementing the sort interface for customSort
func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool { return p[i].Count < p[j].Count }

func getStdin() {
	file, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if file.Size() > 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		inputFiles[file.Name()] = string(bytes)
	}
}

func getArgs() {
	files := os.Args[1:]
	if len(files) > 0 {
		for _, filename := range files {
			parseFile(filename)
		}
	}
}

func parseFile(filePath string) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to open %s", file)
	}
	inputFiles[filePath] = string(file)
}

type sortedList []map[string]Pairs

var inputFiles = make(map[string]string)
var sortedlist sortedList

func customSort(filename string, trios map[string]int) {
	sortInterface := make(Pairs, len(trios))
	sorted := make(map[string]Pairs)
	i := 0
	for k, v := range trios {
		sortInterface[i] = Pair{k, v}
		i++
	}
	sort.Sort(sortInterface) // This sorts the sortInterface in place
	sorted[filename] = sortInterface
	sortedlist = append(sortedlist, sorted)
}

// wordTokenizer splits file content into stream of lowered case word tokens,
// punctuations trimed and store three-words string into a map for sortting.

func wordTokenizer(fileName string, fileContent string) {
	sr := bufio.NewScanner(strings.NewReader(fileContent))
	sr.Split(bufio.ScanWords)
	var words []string
	for sr.Scan() {
		words = append(words, strings.Trim(strings.ToLower(sr.Text()), ",.?!"))
	}
	trios := make(map[string]int)
	for i := 0; i < len(words); i++ {
		if i+2 < len(words) {
			trio := fmt.Sprintf("%s %s %s", words[i], words[i+1], words[i+2])
			// count the occurance of each three words string in a map
			if _, ok := trios[trio]; ok {
				trios[trio] += 1
			} else {
				trios[trio] = 1
			}
		}
	}
	customSort(fileName, trios)
}

func printMostCommon(top int) {
	if len(sortedlist) > 0 {
		for _, m := range sortedlist {
			for filename, pairs := range m {
				fmt.Println()
				fmt.Println(filename, ":\n")
				count := 0
				for i := len(pairs) - 1; i > 0; i-- {
					fmt.Println(pairs[i].Words, "-", pairs[i].Count)
					count += 1
					if count > top {
						break
					}
				}
			}
		}
	}

}
func main() {
	// getStdin() and getArgs() methods collect files provided at command line
	getStdin()
	getArgs()

	if len(inputFiles) > 0 {
		for filename, content := range inputFiles {
			wordTokenizer(filename, content)
		}
		printMostCommon(100)
	} else {
		fmt.Println("Please provide files to be parsed")
	}

}
