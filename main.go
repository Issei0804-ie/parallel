package main

import (
	"fmt"
	"log"
	"math"
)

type Item struct {
	Value  uint64
	Weight uint64
}

/*
6 65
120 10
130 12
80 7
100 9
250 21
185 16
*/

func main() {
	var n, weights, i uint64
	_, err := fmt.Scanf("%v %v", &n, &weights)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var items []Item

	for i = 0; i < n; i++ {
		var value, weight uint64
		_, err = fmt.Scanf("%v %v", &value, &weight)
		if err != nil {
			log.Fatalf(err.Error())
		}
		items = append(items, Item{
			Value:  value,
			Weight: weight,
		})
	}

	//w, v, b := solve(items, weights)
	w, v, b := parallelSolve(items, weights)
	fmt.Printf("w = %v v = %v b = %v", w, v, b)
}

func solve(items []Item, weight uint64) (w, v, bits uint64) {
	var j uint64
	p := math.Pow(2, float64(len(items)))
	w = 0
	v = 0
	bits = 0
	for j = 0; j < uint64(p); j++ {
		w2, v2 := calcWeight(items, j)
		if w2 <= weight && w < w2 {
			w = w2
			v = v2
			bits = j
		}
	}
	return
}

func calcWeight(items []Item, bits uint64) (weight, value uint64) {
	weight = 0
	value = 0
	for i := 0; i < len(items); i++ {
		if bits&1 == 1 {
			weight += items[i].Weight
			value += items[i].Value
		}
		bits = bits >> 1
	}
	return
}

type Result struct {
	w, v, b uint64
	isOver  bool
}

func parallelSolve(items []Item, weight uint64) (w, v, bits uint64) {
	var j uint64
	p := math.Pow(2, float64(len(items)))
	w = 0
	v = 0
	bits = 0

	r := make(chan Result)

	for j = 0; j < uint64(p); j++ {
		go parallelCalcWeight(items, j, r)
	}

	for j = 0; j < uint64(p); j++ {
		result := <-r
		w2 := result.w
		v2 := result.v
		b2 := result.b
		if w2 <= weight && w < w2 {
			w = w2
			v = v2
			bits = b2
		}
	}
	return
}

func parallelCalcWeight(items []Item, bits uint64, r chan Result) {
	var weight, value uint64
	weight = 0
	value = 0
	shiftBits := bits
	result := Result{}
	for i := 0; i < len(items); i++ {
		if shiftBits&1 == 1 {
			weight += items[i].Weight
			value += items[i].Value
		}
		shiftBits = shiftBits >> 1
	}
	result.w = weight
	result.v = value
	result.b = bits

	r <- result
}
