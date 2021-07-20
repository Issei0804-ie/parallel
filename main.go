package main

import (
	"fmt"
	"log"
	"math"
)

type Item struct {
	Value  int
	Weight int
}

func main() {
	var n, weights int
	_, err := fmt.Scanf("%v %v", &n, &weights)
	if err != nil {
		log.Fatalf(err.Error())
	}

	var items []Item

	for i := 0; i < n; i++ {
		var value, weight int
		_, err = fmt.Scanf("%v %v", &value, &weight)
		if err != nil {
			log.Fatalf(err.Error())
		}
		items = append(items, Item{
			Value:  value,
			Weight: weight,
		})
	}

	w, v, b := solve(items, weights)
	fmt.Printf("w = %v v = %v b = %v", w, v, b)
}

func solve(items []Item, weight int) (w, v, bits int) {
	fmt.Println(items)
	p := math.Pow(2, float64(len(items)))
	w = 0
	v = 0
	bits = 0
	for j := 0; j < int(p); j++ {
		fmt.Println(j)
		w2, v2 := calcWeight(items, j)
		if w2 <= weight && w < w2 {
			w = w2
			v = v2
			bits = j
		}
	}
	return
}

func calcWeight(items []Item, bits int) (weight, value int) {
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
