package main

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
	"time"
)

func BenchmarkFunction1(b *testing.B) {
	// Panggil fungsi yang akan diuji
	for n := 0; n < b.N; n++ {
		Function1()
	}
}

func BenchmarkFunction2(b *testing.B) {
	// Panggil fungsi yang akan diuji
	for n := 0; n < b.N; n++ {
		Function2()
	}
}

func Function1() {
	// Fungsi yang akan diuji
	time.Sleep(10 * time.Millisecond)
}

func Function2() {
	// Fungsi yang akan diuji
	time.Sleep(20 * time.Millisecond)
}

func main() {
	// Buat tabwriter untuk menampilkan tabel
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	// Jalankan benchmark
	fmt.Fprintf(w, "Function\tTime\n")
	fmt.Fprintf(w, "Function1\t%s\n", testing.Benchmark(BenchmarkFunction1).String())
	fmt.Fprintf(w, "Function2\t%s\n", testing.Benchmark(BenchmarkFunction2).String())

	// Tampilkan tabel
	w.Flush()
}
