package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	data := []float64{3, 4, 1, 2, 5, 6, 3, 3, 3, 3, 3, 6, 6, 6, 6, 6}
	fmt.Println("Data:", data)

	rata2 := hitungMean(data)
	median := hitungMedian(data)
	mode := hitungMode(data)

	fmt.Println("Rata-rata:", rata2)
	fmt.Println("Median:", median)
	fmt.Println("Modus:", strings.Trim(fmt.Sprint(mode), "[]"))
}

func hitungMean(data []float64) (rata2 float64) {
	var jumlah float64 = 0

	// Menghitung nilai rata-rata
	for i := 0; i < len(data); i++ {
		jumlah = jumlah + data[i]
	}
	rata2 = jumlah / float64(len(data))

	return rata2
}

func hitungMedian(data []float64) (median float64) {
	// Mengurutkan slice data
	sort.Float64s(data)

	// Menghitung median, tergantung dr apakah jumlah datanya ganjil atau genap
	if len(data)%2 == 0 {
		var index1 float64 = data[(len(data)/2)-1]
		var index2 float64 = data[((len(data)/2)+1)-1]
		median = (index1 + index2) / 2
	} else {
		median = data[((len(data)+1)/2)-1]
	}

	return median
}

func hitungMode(data []float64) (mode []float64) {
	// Menggunakan map untuk menghitung banyaknya setiap nilai dalam slice data
	m := make(map[float64]int)
	for _, v := range data {
		if _, found := m[v]; found {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	// Mencari nilai yang terbanyak muncul (modus) dari map yg telah dibuat
	sliceVals := make([]int, len(m))

	for _, v := range m {
		sliceVals = append(sliceVals, v)
	}
	sort.Ints(sliceVals)

	for i, v := range m {
		if v == sliceVals[len(sliceVals)-1] {
			mode = append(mode, i)
		} else {
			continue
		}
	}

	return mode
}
