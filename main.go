package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	pattern  string
	filePath string
	destPath string
)

func init() {
	flag.StringVar(&filePath, "o", "", "the filepath of target")
	flag.StringVar(&destPath, "d", "", "the filepath of destination")
	flag.StringVar(&pattern, "p", "='?(.*)", "the pattern you want to match in regular expression")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: pattern-uniq [options] [root]\n")
	fmt.Fprintf(os.Stderr, " It was used for getting rid of similar elements and make it to be the unique one:\n")
	fmt.Fprintf(os.Stderr, " for example: fisrt line is \"alias abc='echo abc'\", second line is \"alias echoabc='echo abc'\"  \n")
	fmt.Fprintf(os.Stderr, " you can keep either one of lines you need by using regex pattern such `='?(.*)` \n")
	flag.PrintDefaults()
}

func main() {

	flag.Parse()
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		panic("read file" + err.Error())
	}

	var result_arr []string
	map_ele := map[string]string{}
	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(pattern)
	for scanner.Scan() {
		cur_line := scanner.Text()
		match := regex.FindStringSubmatch(cur_line)
		if match == nil || match[1] == "" {
			continue
		}

		map_ele[match[1]] = cur_line
	}

	for _, e := range map_ele {
		result_arr = append(result_arr, e)
	}

	w, err := os.Create(destPath)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	for _, data := range result_arr {
		_, err := w.WriteString(data + "\n")
		if err != nil {
			panic(err)
		}
	}
	w.Sync()
	fmt.Println("done")

}
