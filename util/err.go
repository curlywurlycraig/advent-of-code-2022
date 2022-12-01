package util

import "fmt"

func NoErr(err error) {
	if err != nil {
		fmt.Println("got an error. the elves didn't sign up for this:", err)
		panic(err)
	}
}
