package main

import "fmt"

const idx_max = 500	
type barang_mentah struct{
	nama string
	berat float64
	harga int
	barang_hasil string
	}
type barang_jadi struct{
	nama string
	berat float64
	harga int
	asal_barang string
}
type array_mentah[idx_max]barang_mentah
type array_jadi[idx_max]barang_jadi

array_mentah[1].nama = "bawang"
array_jadi[1].nama = "nasi goreng"

func main() {
		
}

func tampil_barang(){
	var i int
	for i <= idx_max{
		fmt.Println(i, ". ", array_mentah[i])
		if input_user == 9{
			for i 
		}
	}
}
