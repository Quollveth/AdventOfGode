package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileFull(name string) []byte {
	c, e := os.ReadFile(name)
	if e != nil {
		log.Fatal(e)
	}

	return c
}

func ReadFileLines(name string) []string {
	f, e := os.Open(name)
	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	var c []string
	for s.Scan() {
		c = append(c, s.Text())
	}

	return c
}

func ReadFileWords(name string) []string {
	f, e := os.Open(name)
	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	var c []string
	for s.Scan() {
		c = append(c, s.Text())
	}

	return c
}
