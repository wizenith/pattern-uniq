package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

var (
	pattern   string
	src_path  string
	dest_path string
)

func init() {
	flag.StringVar(&src_path, "o", "", "the filepath of target")
	flag.StringVar(&dest_path, "d", "", "the filepath of destination")
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
	src_path = "/home/wizenith/Desktop/git_alias"
	dest_path = "/home/wizenith/Desktop/git_alias_filtered"
	flag.Parse()
	file, err := os.Open(src_path)
	if err != nil {
		panic("read file" + err.Error())
	}
	defer file.Close()

	w, err := os.Create(dest_path)
	if err != nil {
		panic(err)
	}
	defer w.Close()
	// let files of input and output be ready before processing any logic
	// var result_arr []string
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

	for _, data := range map_ele {
		_, err := w.WriteString(data + "\n")
		if err != nil {
			panic(err)
		}
	}
	w.Sync()
	fmt.Println("done")

}
