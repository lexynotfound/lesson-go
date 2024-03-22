package main

import (
	"fmt"
	"os"
)

// Struct untuk menyimpan data teman

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Data teman-teman di kelas

var dataTeman = map[int]Teman{
	1: {"Ramadhan", "Jl. Bogor Ciputat", "Developer", "Ingin mempelajarin bahasa Go untuk meningkatkan skill pemerograman"},
	2: {"Syarul", "Jl Meruya Ilir", "Developer", "Mendapat rekomendasi dari rekan kerja untuk mengikuti kelasi ini"},
	3: {"Wafiq", "Jl. Bandung Merdeka", "Developer Flutter", "Ingin Memepelajarin bahasa go untuk di jadikan sebuah back-end pada bahasa go"},
	4: {"Galih", "Jl Ahmad Yani", "Data Scientis", "Ingin beralih menjadi seorang back-end developer menggunakan bahasa go"},
	// 957356840-778: {"Kurnia Raihan Ardian", "Jl. Senayan GBK", "Developer Mobile", "Ingin beralih ke backend developer menggunakan bahasa go"},
	778: {"Kurnia Raihan Ardian", "Jl. Senayan GBK", "Developer Mobile", "Ingin beralih ke backend developer menggunakan bahasa go"},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run biodata.go <1957356840-778>")
		return
	}

	nomorAbsen := args[1]
	temen, err := getTemanByNomorAbsen(nomorAbsen)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Menampilkan data teman
	fmt.Println("Nama:", temen.Nama)
	fmt.Println("Alamat:", temen.Alamat)
	fmt.Println("Pekerjaan:", temen.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", temen.Alasan)
}

func getTemanByNomorAbsen(nomorAbsen string) (Teman, error) {
	//Konversi nomor absen ke dalam bentuk integer

	var absen int
	_, err := fmt.Sscanf(nomorAbsen, "%d", &absen)
	if err != nil {
		return Teman{}, fmt.Errorf("nomor absen harus berupa angka")
	}

	temen, ok := dataTeman[absen]
	if !ok {
		return Teman{}, fmt.Errorf("tidak ada teman dengan nomor absen %d", absen)
	}
	return temen, nil
}
