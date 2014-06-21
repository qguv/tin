package main

import "testing"
import (
	"fmt"
	"math/rand"
	"time"
)

func TestHawaiianNameString(t *testing.T) {
	randomSource := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(randomSource)

	var n name

	n = name{
		origin: hawaiian,
		min:    2,
		max:    5,
	}

	fmt.Println("Hawaiian:")
	for i := 0; i < 20; i++ {
		fmt.Println("\t", n.String(r))
	}
}

func TestJapaneseNameString(t *testing.T) {
	randomSource := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(randomSource)

	var n name

	n = name{
		origin: japanese,
		min:    2,
		max:    5,
	}

	fmt.Println("Japanese:")
	for i := 0; i < 20; i++ {
		fmt.Println("\t", n.String(r))
	}
}
