package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trueNumber(number int, pcNumber int) (sonuc int) {
	if number < pcNumber {
		sonuc = 3
	} else if number > pcNumber {
		sonuc = 2
	} else {
		sonuc = 1
	}
	return
}

func main() {
	var sayi int
	min := 0
	max := 100
	rand.Seed(time.Now().UnixNano())
	pcNumber := rand.Intn(max-min+1) + min

	fmt.Println("SAYI TAHMİN OYUNUNA HOŞ GELDİNİZ")
	fmt.Printf("%d ve %d arasında bir sayıyı bulmanız gerekiyor\n", min, max)
	fmt.Println("Lütfen sayınızı giriniz")

	for {
		fmt.Scan(&sayi)
		sonuc := trueNumber(sayi, pcNumber)

		if sonuc == 1 {
			fmt.Println("Tebrikler! Doğru sayıyı buldunuz.")
			break
		} else if sonuc == 2 {
			fmt.Println("Daha küçük bir sayı olmalı")
		} else if sonuc == 3 {
			fmt.Println("Daha büyük bir sayı olmalı")
		}

	}
}
