package main

import "fmt"

const idx_max = 500

type barang_mentah struct {
	nama         string
	berat        float64
	harga        int
	barang_hasil string
}
type barang_jadi struct {
	nama        string
	berat       float64
	harga       int
	asal_barang string
}

func main() {
	var array_mentah [idx_max]barang_mentah
	var array_jadi [idx_max]barang_jadi
	var input_user int
	array_mentah[1].nama = "bawang"
	array_jadi[1].nama = "nasi goreng"
	fmt.Scan(&input_user)
	for input_user != 0000 {
		menu()
	}
}

func menu() {
	var input_user int
	fmt.Println("1. Tampilkan Barang")
	fmt.Println("2. Cari Barang")
	fmt.Println("3. Edit Barang")
	fmt.Println("4. Hapus Barang")
	if input_user == 1 {
		tampil_barang(&barang_jadi)
	}
}

func tampil_barang(tab *barang_mentah) {
	var i int
	var input_user int
	for i <= idx_max {
		fmt.Println(i, ". ", *tab[i])
		if input_user == 1 {
			sorting_mentah(&barang_mentah)
		} else if input_user == 2 {
			sorting_jadi(&barang_jadi)
		}
	}
}

func sorting_mentah(tab *barang_mentah) {
	for i := 0; i < len(*tab); i++ {
		min_index := 1
		for j := i + 1; j < len(*tab); j++ {
			if (*tab)[min_index] > (*tab)[j] {
				min_index = j
			}
		}
		tmp := (*tab)[i]
		(*tab)[i] = (*tab)[min_index]
		(*tab)[min_index] = tmp
	}
}

func sorting_jadi(tab *array_mentah) {
	var min_index int
	var j int
	for i := 0; i < len(*tab); i++ {
		min_index = 1
		for j := i + 1; j < len(*tab); j++ {
			if (*tab)[min_index] > array_mentah[j] {
				min_index = j
			}
		}
		tmp := (*tab)[j]
		*tab = *tab[min_index]
		(*tab)[min_index] = tmp
	}
}
