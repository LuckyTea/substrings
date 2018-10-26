package main

import (
	"bufio"
	"flag"
	"os"
	"sort"
	"strings"
)

func main() {
	var line, path string

	flag.StringVar(&line, "l", "ATGGAGAAAATAGTRCTTCTTCTTGCAATAGTCAGTCTTGTTAAAAGTGATCAGATTTGCATTGGTTACCATGCAAACAATTCAACAGAGCAGGTTGACACAATCATGGAAAAGAACGTTACTGTTACACATGCCCAAGACATACTGGAAAAGACACACAACGGGAAGCTCTGCGATCTAGATGGAGTGAAGCCTCTAATTTTAAGAGATTGTAGTGTAGCTGGATGGCTCCTCGGGAACCCAATGTGTGACGAATTCATCAATGTACCGGAATGGTCTTACATAGTGGAGAAGGCCAATCCAACCAATGACCTCTGTTACCCAGGGAGTTTCAACGACTATGAAGAACTGAAACATCTATTGAGCAGAATAAACCATTTTGAGAAAATTCAAATCATCCCCAAAAGTTCTTGGTCCGATCATGAAGCCTCATCAGGAGTGAGCTCAGCATGTCCATACCTGGGAAGTCCCTCCTTTTTTAGAAATGTGGTATGGCTTATCAAAAAGAACAGTACATACCCAACAATAAAGAAAAGCTACAATAATACCAACCAAGAAGATCTTTTGGTACTGTGGGGAATTCACCATCCTAATGATGCGGCAGAGCAGACAAGGCTATATCAAAACCCAACCACCTATATTTCCATTGGGACATCAACACTAAACCAGAGATTGGTACCAAAAATAGCTACTAGATCCAAAGTAAACGGGCAAAGTGGAAGGATGGAGTTCTTCTGGGCAATTTTAAAACCTAATGATGCAATCAACTTCGAGAGTAATGGAAATTTCATTGCTCCAGAATATGCATACAAAATTGTCAAGAAAGGGGACTCAGCAATTATGAAAAGTGAATTGGAATATGGTAACTGCAACACCAAGTGTCAAACTCCAATGGGGGCGATAAACTCTAGTATGCCATTCCACAACATACACCCTCTCACCATCGGGGAATGCCCCAAATATGTGAAATCAAACAGATTAGTCCTTGCAACAGGGCTCAGAAATAGCCCTCAAAGAGAGAGCAGAAGAAAAAAGAGAGGACTATTTGGAGCTATAGCAGGTTTTATAGAGGGAGGATGGCAGGGAATGGTAGATGGCTGGTATGGGTACCACCATAGCAATGAGCAGGGGAGTGGGTACGCTGCAGACAAAGAATCCACTCAAAAGGCAATAGATGGAGTCACCAATAAGGTCAACTCAATTATTGACAAAATGAACACTCAGTTTGAGGCTGTTGGAAGGGAATTTAATAACTTAGAAAGGAGAATAGAGAATTTAAACAAGAAGATGGAAGACGGGTTTCTAGATGTTTGGACTTATAATGCCGAACTTCTGGTTCTCATGGAAAATGAGAGAACTCTAGACTTTCATGACTCAAATGTTAAGAACCTCTACGACAAGGTCCGACTACAGCTTAGGGATAATGCAAAAGAGCTGGGTAACGGTTGTTTCGAGTTCTATCACAAATGTGATAATGAATGTATGGAAAGTATAAGAAACGGAACGTACAACTATCCGCAGTATTCAGAAGAAGCAAGATTAAAAAGAGAGGAAATAAGTGGGGTAAAATTGGAATCAATAGGAACTTACCAAATACTGTCAATTTATTCAACAGTAGCGAGTTCCCTAGCACTGGCAATCATGATAGCTGGTCTATCTTTATGGATGTGCTCCAATGGATCGTTACAATGCAGAATTTGCATTTAA", "string to parse")
	flag.StringVar(&path, "p", "samples.txt", "path to file with substrings")
	flag.Parse()

	arr, err := readSubstrings(path, len(line))
	if err != nil {
		panic(err)
	}

	sortedArr := make([]string, len(arr))
	copy(sortedArr, arr)
	sort.Sort(byLen(sortedArr))

	answer := split(line, sortedArr)

	for k, v := range answer {
		if k != 0 {
			print(",")
		}
		print(indexOf(v, arr))
	}
}

// readSubstrings - read file and collect substrings //////////////////////////
// assume that Id always start at 0, and the numbers are consecutive //////////
func readSubstrings(path string, maxSize int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var substrings []string
	var substring string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		substring = strings.Split(scanner.Text(), ",")[1]
		// substring must not be longer than line /////////////////////////////
		if len(substring) <= maxSize {
			substrings = append(substrings, substring)
		}
	}

	return substrings, scanner.Err()
}

// split line by sublines /////////////////////////////////////////////////////
func split(line string, array []string) []string {
	var answer []string

	for len(line) > 0 {
		for _, v := range array {
			if strings.HasPrefix(line, v) {
				answer = append(answer, v)
				line = line[len(v):]

				break
			}
		}
	}

	return answer
}

type byLen []string

func (a byLen) Len() int           { return len(a) }
func (a byLen) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
func (a byLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// indexOf - returns the index of the element in the original slice ///////////
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}

	return -1
}
