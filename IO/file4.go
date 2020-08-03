package main

import (
	"fmt"
	"os"
)

type Student struct {
	name string
	roll int
}

func main() {
	f1, err := os.Create("tmp1.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f1.Close()
	student := []Student{{"rrm", 1}, {"ajp", 2}, {"rjt", 3}, {"psg", 4}}
	for _, x := range student {
		f1.WriteString(x.name + "\t")
		f1.WriteString(fmt.Sprintf("%d\n", x.roll))
	}
}
