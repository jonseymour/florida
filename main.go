package main

import (
	"flag"
	"fmt"
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
	duplicate bool
	first     *child
	second    *child
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
func generate(n int) func(out chan<- *family) {
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
func name(probability float64, noDuplicates bool) func(in <-chan *family, out chan<- *family) {
	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			f.first.nameGirl(probability)
			probability2 := probability
			if noDuplicates {
				if f.first.isAGirlNamedFlorida() {
					probability2 = -1.0
				}
			}
			f.second.nameGirl(probability2)
			out <- f
		}
		close(out)
	}
}

// names each girl in a family
func family2child() func(in <-chan *family, out chan<- *child) {
	return func(in <-chan *family, out chan<- *child) {
		for f := range in {
			f.duplicate = false
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

func floridaFamily() func(in <-chan *family, out chan<- *family) {
	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			if !f.hasAGirlNamedFlorida() {
				continue
			}
			out <- f
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

// select families that have not already been seen
func uniqueFamilies() func(in <-chan *family, out chan<- *family) {
	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			if f.duplicate {
				continue
			}
			f.duplicate = true
			out <- f
		}
		close(out)
	}
}

// Reproduce Part D of the Mary, Jane K problem.
//
// Randomly sample families and daughters. If we pick the same family and
// the same daughter, we output the family otherwise we drop it.
func sampleAndMatch(n int, daughtersMatchFlag bool) func(in <-chan *family, out chan<- *family) {
	families := []*family{}
	daughters := []*child{}

	return func(in <-chan *family, out chan<- *family) {
		for f := range in {
			families = append(families, f)
			if f.first.girl {
				daughters = append(daughters, f.first)
			}
			if f.second.girl {
				daughters = append(daughters, f.second)
			}
		}
		i := 0
		for i < n {
			x := rand.Int31n(int32(len(families)))
			y := rand.Int31n(int32(len(daughters)))
			if families[x] == daughters[y].family {
				if daughtersMatchFlag {
					d := families[x].first
					if !d.girl || families[x].second.girl && coinFlip() {
						d = families[x].second
					}
					if d != daughters[y] {
						continue
					}
				}
				out <- families[x]
				i++
			}
		}
		close(out)
	}
}

// accounting stage
func accounting(n int, girls bool) func(in <-chan *family, done chan<- struct{}) {
	return func(in <-chan *family, done chan<- struct{}) {
		c := 0
		gg := 0
		for f := range in {
			c++
			if f.isGG() {
				gg++
			}
		}
		tn := "families"
		if girls {
			tn = "girls"
		}
		fmt.Fprintf(
			os.Stdin,
			"[%s] n=%d, c=%d, gg=%d, gg/c=%0.1f%%\n",
			tn,
			n,
			c,
			gg,
			float64(gg)/float64(c)*100.0)
		done <- struct{}{}
	}
}

func main() {
	var probability float64
	var n int
	var girlsFlag bool
	var floridaFlag bool
	var uniqueFlag bool
	var atLeastOneGirlFlag bool
	var sampleAndMatchFlag bool
	var daughtersMatchFlag bool
	var noDuplicatesFlag bool

	flag.Float64Var(&probability, "probability", 0.001, "Probability of naming a girl florida.")
	flag.IntVar(&n, "families", 1000000, "Number of families.")
	flag.BoolVar(&girlsFlag, "girls", false, "Select girls, not families.")
	flag.BoolVar(&floridaFlag, "florida", false, "Select girls named Florida.")
	flag.BoolVar(&atLeastOneGirlFlag, "at-least-one-girl", false, "Select families with at least one girl.")
	flag.BoolVar(&sampleAndMatchFlag, "sample-and-match", false, "Sample daughters and families and output if matched.")
	flag.BoolVar(&daughtersMatchFlag, "daughters-match", true, "Insist daughters match during sampling.")
	flag.BoolVar(&uniqueFlag, "unique-families", false, "Count unique families of girls, not girls.")
	flag.BoolVar(&noDuplicatesFlag, "no-duplicate-names", false, "Do not allow duplicate names.")
	flag.Parse()

	p1 := make(chan *family)
	p2 := make(chan *family)
	p3 := make(chan *family)
	done := make(chan struct{})

	go generate(n)(p1)
	go name(probability, noDuplicatesFlag)(p1, p2)
	if atLeastOneGirlFlag {
		p4 := make(chan *family)
		go atLeastOneGirl()(p2, p4)
		p3 = p4
	} else {
		p3 = p2
	}
	if girlsFlag {
		p4 := make(chan *child)
		p5 := make(chan *family)
		p7 := make(chan *child)

		go family2child()(p3, p7)
		if floridaFlag {
			go florida()(p7, p4)
		} else {
			go girls()(p7, p4)
		}
		go child2family()(p4, p5)
		p3 = p5
	} else if floridaFlag {
		p4 := make(chan *family)
		go floridaFamily()(p3, p4)
		p3 = p4
	}
	if sampleAndMatchFlag {
		p4 := make(chan *family)
		go sampleAndMatch(n, daughtersMatchFlag)(p3, p4)
		p3 = p4
	}
	if uniqueFlag {
		p4 := make(chan *family)
		go uniqueFamilies()(p3, p4)
		p3 = p4
	}
	go accounting(n, girlsFlag && !uniqueFlag)(p3, done)

	<-done
}
