package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(8)
	rand.Seed(time.Now().UnixNano())
}

const largePrime = 15485863

var wordList = []string{
	"harbor", "cylinder", "glide", "casualty", "sources", "regrets", "gates", "selectors", "instruction",
	"overload", "agreements", "alignment", "flakes", "tickets", "ladders", "member", "insignias",
	"freedom", "bypass", "bridge", "control", "grease", "tackles", "condensers", "stencil", "security",
	"duration", "skirts", "catalogs", "thickness", "formations", "minutes", "chain", "segments", "rakes",
	"waves", "leapers", "boilers", "rowboats", "watchstanding", "recombination", "selector", "acceptors",
	"bed", "quiet", "laughs", "pens", "chills", "tachometers", "rubbish", "developments", "brackets",
	"churn", "chins", "driver", "additions", "twirls", "parameters", "spares", "beings", "macro", "congress",
	"platter", "profit", "wrist", "facts", "policy", "lake", "increment", "gyro", "interchanges", "claims",
	"compositions", "interviewer", "degrees", "tear", "wheels", "taps", "arrests", "substitute", "points",
	"entries", "wax", "societies", "knowledge", "disassemblies", "splashes", "hub", "thermometers", "possession",
	"magazines", "honors", "subprograms", "codes", "bullet", "strikes", "man", "schoolrooms", "father", "worksheets",
}

// START OMIT

func DoWork(n int, c chan string) {
	for m := range c {
		time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
		fmt.Println("Worker", n, "processed word:", m)
	}
}

func main() {
	c := make(chan string)

	for i := 0; i < 8; i++ {
		go DoWork(i, c)
	}
	for _, w := range wordList {
		c <- w
	}

	close(c)
}

// END OMIT
