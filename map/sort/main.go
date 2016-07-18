package main

import (
	"fmt"
	"sort"
)

type kv struct {
	key   string
	value float64
}

func main() {
	myMap := map[string]float64{"c": 1.78, "b": 4.12, "a": 0.55}
	for _, v := range sortMapFloat(myMap) {
		fmt.Printf("%s %f\n", v.key, v.value)
	}
}

// sortMapFloat sorts map[string]float64 w/ value
func sortMapFloat(m map[string]float64) []kv {
	n := map[float64][]string{}
	var (
		a []float64
		r []kv
	)
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			r = append(r, kv{s, k})
		}
	}
	return r
}
