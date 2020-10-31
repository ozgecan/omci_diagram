package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanFile(logfile string) {
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		value := sc.Text()
		fmt.Println("line", value)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
		return
	}
}
