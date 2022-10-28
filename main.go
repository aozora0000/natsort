package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

var reverse = flag.Bool("reverse", true, "sort reverse")

func main() {
	flag.Parse()
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	list := strings.Split(string(bytes), "\n")
	if *reverse {
		sort.Sort(sort.Reverse(stringSlice(list)))
	} else {
		sort.Sort(stringSlice(list))
	}
	fmt.Println(strings.Join(list, "\n"))
}
