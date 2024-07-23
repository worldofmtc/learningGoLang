package main

import "fmt"

func main() {

	var bakiye, para float32
	var islem int
	var ad = "Muhammed Talha"
	var soyad = "Çorbacı"
	var inputSifre string
	var mSifre = "1234"
	var banka = "GoBank"

	bakiye = 1000.0

	fmt.Printf("Merhaba %s Müşterimiz \nLütfen Şifrenizi Giriniz: ", banka)
	fmt.Scanln(&inputSifre)

	if inputSifre == mSifre {
		fmt.Printf("HoşGeldiniz Sayın %s %s\n", ad, soyad)

		fmt.Printf("Bakiyeniz : %f Hangi işlemi yapmak istersiniz \n Eğer Para yatırmak isterseniz %d \n Eğer para çekmek istiyorsanız %d  e basınız \n", bakiye, 1, 2)
		fmt.Scanln(&islem)

		load := "yükleniyor" // bunu ilerde paramtere olarak saniye verip heryerde kullanabilceğemiz bir load fonksiyonu yapabiliriz
		for i := 0; i < 10; i++ {
			fmt.Printf("%s ...... \n", load)
		}

		switch islem {
		case 1:
			fmt.Println("Kaç tl yatırmak istersiniz ? ")
			fmt.Scan(&para)
			bakiye = para + bakiye
			fmt.Printf("İşleminiz Gerçekleşti \n  Yeni bakyeniz %f tl ", bakiye)

		case 2:

			fmt.Println("Kaç tl çekmek istersiniz ? ")
			fmt.Scan(&para)
			if para > bakiye {
				fmt.Println("çekmek istediğiniz miktar bakiyenizden büyük ")
				return
			} else {
				bakiye = bakiye - para
				fmt.Printf("İşleminiz Gerçekleşti \n Yeni bakyeniz %f  tl", bakiye)
			}

		default:

			fmt.Println("Geçersiz bir işlem seçtiniz")

		}

	} else {
		fmt.Println("Maalesef Şifreniz Hatalı")
		return
	}

}
