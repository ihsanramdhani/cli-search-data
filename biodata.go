package main

import (
	"fmt"
	"os"
	"strconv"
)

// membuat struct berisi nama, alamat, pekerjaan, dan alasan
type Biodata struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	//membuat data dummy
	var data = []Biodata{
		{Nama: "dummy1", Alamat: "alamat1", Pekerjaan: "Engineer", Alasan: "alasan1"},
		{Nama: "dummy2", Alamat: "alamat2", Pekerjaan: "Engineer", Alasan: "alasan2"},
		{Nama: "dummy3", Alamat: "alamat3", Pekerjaan: "Engineer", Alasan: "alasan3"},
		{Nama: "dummy4", Alamat: "alamat4", Pekerjaan: "Engineer", Alasan: "alasan4"},
		{Nama: "dummy5", Alamat: "alamat5", Pekerjaan: "Finance", Alasan: "alasan5"},
		{Nama: "dummy6", Alamat: "alamat6", Pekerjaan: "Finance", Alasan: "alasan6"},
		{Nama: "dummy7", Alamat: "alamat7", Pekerjaan: "UI", Alasan: "alasan7"},
		{Nama: "dummy8", Alamat: "alamat8", Pekerjaan: "UI", Alasan: "alasan8"},
		{Nama: "dummy9", Alamat: "alamat9", Pekerjaan: "UX", Alasan: "alasan9"},
		{Nama: "dummy10", Alamat: "alamat10", Pekerjaan: "UX", Alasan: "alasan10"},
	}

	//looping setiap data pada variabel data
	flag := "false"

	for index, item := range data {
		//mengambil argumen pertama
		arg := os.Args[1]
		//menambah nilai index agar index dimulai dari 1
		index = index + 1
		//merubah tipe data index dari int menjadi string
		index := strconv.Itoa(index)
		//kondisi yg dilakukan ketika nilai argumen yg dimasukkan dan index sama
		if arg == index {
			fmt.Printf("Nama: %s\n", item.Nama)
			fmt.Printf("Alamat: %s\n", item.Alamat)
			fmt.Printf("Pekerjaan: %s\n", item.Pekerjaan)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", item.Alasan)
			flag = "true"
		}
	}
	if flag == "false" {
		fmt.Println("Data tidak ditemukan")
	}
}
