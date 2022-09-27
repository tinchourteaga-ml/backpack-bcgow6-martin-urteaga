package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(valores []int, ch chan float64) {
	inicio := time.Now()
	for i := 0; i < len(valores)-1; i++ {
		for j := 0; j < len(valores)-i-1; j++ {
			if valores[j] > valores[j+1] {
				valores[j], valores[j+1] = valores[j+1], valores[j]
			}
		}
	}
	ch <- float64(time.Since(inicio))
}

func selectionSort(valores []int, ch chan float64) {
	inicio := time.Now()
	for i := 0; i < len(valores); i++ {
		var minIdx = i
		for j := i; j < len(valores); j++ {
			if valores[j] < valores[minIdx] {
				minIdx = j
			}
		}
		valores[i], valores[minIdx] = valores[minIdx], valores[i]
	}
	ch <- float64(time.Since(inicio))
}

func insertionSort(valores []int, ch chan float64) {
	inicio := time.Now()
	for i := 1; i < len(valores); i++ {
		j := i
		for j > 0 {
			if valores[j-1] > valores[j] {
				valores[j-1], valores[j] = valores[j], valores[j-1]
			}
			j = j - 1
		}
	}
	ch <- float64(time.Since(inicio))
}

func main() {
	ch1 := make(chan float64)
	ch2 := make(chan float64)
	ch3 := make(chan float64)
	sliceOfSlices := [][]int{}
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	sliceOfSlices = append(sliceOfSlices, variable1)
	sliceOfSlices = append(sliceOfSlices, variable2)
	sliceOfSlices = append(sliceOfSlices, variable3)

	for i, slc := range sliceOfSlices {
		go bubbleSort(slc, ch1)
		fmt.Printf("Bubble sort array %d: %.2fms\n", i, <-ch1)
	}

	for i, slc := range sliceOfSlices {
		go selectionSort(slc, ch2)
		fmt.Printf("Selection sort array %d: %.2fms\n", i, <-ch2)
	}

	for i, slc := range sliceOfSlices {
		go insertionSort(slc, ch3)
		fmt.Printf("Insertion sort array %d: %.2fms\n", i, <-ch3)
	}
}
