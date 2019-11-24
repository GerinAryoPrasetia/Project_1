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
type array_mentah [idx_max]barang_mentah
type array_jadi [idx_max]barang_jadi

func main() {
	var input_user byte
	array_mentah[1].nama = "bawang"
	array_jadi[1].nama = "nasi goreng"
	fmt.Scan(&input_user)
	for input_user != 0000 {
		fmt.Println("1. Tampilkan Barang")
		fmt.Println("2. Cari Barang")
		fmt.Println("3. Edit Barang")
		fmt.Println("4. Hapus Barang")
		if input_user == 1 {
			tampil_barang()
		} else if input_user == 2 {
			cari_barang()
		} else if input_user == 3 {
			edit_barang()
		} else if input_user == 4 {
			hapus_barang()
		}
	}
}

func tampil_barang() {
	var i int
	for i <= idx_max {
		fmt.Println(i, ". ", array_mentah[i])
		if input_user == 9 {
			for i {

			}
		}
	}
}

func sorting() {
	var N int = len(array_jadi)
	for i = 0; i < N; i++ {

	}
}
