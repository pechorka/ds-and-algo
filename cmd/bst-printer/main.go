package main

import (
	"ds-and-algo/pkg/bst"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if len(os.Args) != 2 {
		return nil
	}

	treeValsInput := os.Args[1]
	bst := bst.New()
	for _, s := range strings.Split(treeValsInput, ",") {
		s = strings.TrimSpace(s)
		v, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		bst.Insert(v)
	}

	fmt.Println(bst.DrawTree())
	return nil
}
