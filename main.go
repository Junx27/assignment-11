package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

// Membuat fungsi mengambil data yang dimasukan pengguna dan menghitung hasil kuadratnya
func processNumber(bilangan int, wg *sync.WaitGroup, hasilBilangan chan<- int) {
	defer wg.Done()
	kuadrat := bilangan * bilangan
	hasilBilangan <- kuadrat
}

func main() {
	var wg sync.WaitGroup
	var input string
	numbers := []int{}
	hasilBilangan := make(chan int, len(numbers))
	resultKuadrat := []int{}
	fmt.Println("Masukkan minimal 10 bilangan bulat (tekan 'x' untuk mengakhiri):")

	for {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&input)
		//Validasi jumlah angka yang dimasukan
		if strings.ToLower(input) == "x" {
			if len(numbers) < 10 {
				fmt.Println("Anda harus memasukkan minimal 10 bilangan.")
				continue
			}
			break
		}
		//Validasi input pengguna
		bilangan, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan bilangan bulat.")
			continue
		}
		//Memasukan data ke numbers
		numbers = append(numbers, bilangan)

		// Hitung kuadrat secara langsung
		wg.Add(1)
		go processNumber(bilangan, &wg, hasilBilangan)
	}

	// Menunggu semua goroutine selesai
	go func() {
		wg.Wait()
		close(hasilBilangan)
	}()

	// Mengambil hasil dari saluran hasilBilangan
	for hasil := range hasilBilangan {
		resultKuadrat = append(resultKuadrat, hasil)
	}
	//Menampilkan hasil masukan pengguna
	fmt.Println("Hasil masukan pengguna:", numbers)
	fmt.Println("Hasil kuadrat:", resultKuadrat)
}
