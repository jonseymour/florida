package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var Florida = "Florida"

// child knows whether it is a boy or girl and has a name (if it is a girl)
type child struct {
	girl   bool
	name   string
	family *family
}

// names a girl with a specified probability of naming it Florida
func (c *child) nameGirl(probability float64) {
	if !c.girl {
		return
	}
	if rand.Float64() <= probability {
		c.name = Florida
	}
}

func (c *child) isAGirlNamedFlorida() bool {
	return c.girl && c.name == Florida
}

// a family has exactly two children
type family struct {
	first  *child
	second *child
}

func (f *family) atLeastOneGirl() bool {
	return f.first.girl || f.second.girl
}

func (f *family) hasAGirlNamedFlorida() bool {
	return f.first.isAGirlNamedFlorida() || f.second.isAGirlNamedFlorida()
}

func (f *family) isGG() bool {
	return f.first.girl && f.second.girl
}

func (f *family) presentRandomOrder() (*child, *child) {
	if coinFlip() {
		return f.first, f.second
	} else {
		return f.second, f.first
	}
}

// returns true roughly half the time
func coinFlip() bool {
	for {
		f := rand.Float64()
		if f < 0.5 {
			return true
		} else if f >= 0.5 {
			return false
		}
	}
}

// generates a stream of families
func generateFamilies(n int) func(out chan<- *family) {
	return func(out chan<- *family) {
		for i := 0; i < n; i++ {
			f := &family{
				first:  &child{girl: coinFlip()},
				second: &child{girl: coinFlip()},
			}
			f.first.family = f
			f.second.family = f
			out <- f
		}
		close(out)
	}
}

// discards BB families since we are not interested in these
func atLeastOneGirl() func(in <-chan *family, out chan<- *family) {
	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			if !f.atLeastOneGirl() {
				continue
			}
			out <- f
		}
		close(out)
	}
}

// names each girl in a family
func namingProcess(probability float64) func(in <-chan *family, out chan<- *family) {
	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			f.first.nameGirl(probability)
			f.second.nameGirl(probability)
			out <- f
		}
		close(out)
	}
}

// names each girl in a family
func family2child() func(in <-chan *family, out chan<- *child) {
	return func(in <-chan *family, out chan<- *child) {
		for f := range in {
			out <- f.first
			out <- f.second
		}
		close(out)
	}
}

// The child is girl named florida.
func girls() func(in <-chan *child, out chan<- *child) {
	return func(in <-chan *child, out chan<- *child) {
		for c := range in {
			if !c.girl {
				continue
			}
			out <- c
		}
		close(out)
	}
}

// The child is girl named florida.
func florida() func(in <-chan *child, out chan<- *child) {
	return func(in <-chan *child, out chan<- *child) {
		for c := range in {
			if !c.isAGirlNamedFlorida() {
				continue
			}
			out <- c
		}
		close(out)
	}
}

// names each girl in a family
func child2family() func(in <-chan *child, out chan<- *family) {
	return func(in <-chan *child, out chan<- *family) {
		for c := range in {
			out <- c.family
		}
		close(out)
	}
}

// accounting stage`
func accounting(verbose bool) func(in <-chan *family, done chan<- struct{}) {
	return func(in <-chan *family, done chan<- struct{}) {
		countTotal := 0
		countGirlGirl := 0
		countFlorida := 0
		countGirlGirlFlorida := 0

		report := func() {
			fmt.Fprintf(
				os.Stdin,
				"n=%d, gg=%d (%f), florida=%d (%f), florida && gg=%d (%f)\n",
				countTotal,
				countGirlGirl,
				float64(countGirlGirl)/float64(countTotal),
				countFlorida,
				float64(countFlorida)/float64(countTotal),
				countGirlGirlFlorida,
				float64(countGirlGirlFlorida)/float64(countFlorida))
		}

		for f := range in {
			countTotal++
			gg := false
			if f.isGG() {
				countGirlGirl++
				gg = true
			}
			if f.hasAGirlNamedFlorida() {
				countFlorida++
				if gg {
					countGirlGirlFlorida++
				}
				if verbose {
					report()
				}
			}
		}
		report()
		done <- struct{}{}
	}
}

func main() {
	var probability float64
	var n int
	var verbose bool
	var girlsFlag bool
	var floridaFlag bool

	flag.Float64Var(&probability, "probability", 0.001, "Probability of naming a girl florida.")
	flag.IntVar(&n, "iterations", 1000000, "Number of families.")
	flag.BoolVar(&girlsFlag, "girls", false, "Select girls, not families.")
	flag.BoolVar(&floridaFlag, "florida", false, "Select girls named Florida.")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output.")
	flag.Parse()

	if floridaFlag && !girlsFlag {
		log.Fatalf("--florida is only valid if --girls is also specified.")
	}

	p1 := make(chan *family)
	p2 := make(chan *family)
	p3 := make(chan *family)
	p4 := make(chan *child)
	p5 := make(chan *family)
	p6 := make(chan *child)
	p7 := make(chan *child)
	done := make(chan struct{})

	go generateFamilies(n)(p1)
	go namingProcess(probability)(p1, p2)
	go atLeastOneGirl()(p2, p3)
	if girlsFlag {
		go family2child()(p3, p7)
		go girls()(p7, p4)
		if floridaFlag {
			go florida()(p4, p6)
			p4 = p6
		}
		go child2family()(p4, p5)
		p3 = p5
	}
	go accounting(verbose)(p3, done)

	<-done
}
