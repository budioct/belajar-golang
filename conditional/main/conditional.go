package main

import "fmt"

func main() {

	/**
	Seleksi Kondisi
	if (asal == usul){
		return true
		} else{
		return false
		}

	note: Go tidak mendukung seleksi kondisi menggunakan ternary.
	Statement  var data = (isExist ? true : false)
	*/

	// Kondisi Menggunakan Keyword if, else if, & else
	var nilai = 10
	if nilai == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if nilai > 5 {
		fmt.Println("lulus")
	} else if nilai == 4 {
		fmt.Println("tidak lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", nilai)
	}

	// Variabel Temporary Pada if - else
	var point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Kondisi Menggunakan Keyword switch - case
	var angka = 7
	switch angka {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Pemanfaatan case Untuk Banyak Kondisi
	var angka1 = 8
	switch angka1 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Kurung Kurawal Pada Keyword case & default
	var angka2 = 100
	switch angka2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	// Switch Dengan Gaya if - else
	var angka3 = 6
	switch {
	case angka3 == 8:
		fmt.Println("perfect")
	case (angka3 < 8) && (angka3 > 3):
		fmt.Println("ayam goreng")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Penggunaan Keyword fallthrough Dalam switch
	// Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya,
	// jadi satu case di pengecekan selanjutnya tersebut selalu dianggap benar (meskipun aslinya adalah salah).
	// Dalam sebuah switch lebih dari satu fallthrough bisa di tempatkan untuk memaksa melanjutkan proses pengecekan ke satu case setelahnya.
	var angka4 = 6
	switch {
	case angka4 == 8:
		fmt.Println("perfect")
	case (angka4 < 8) && (angka4 > 3):
		fmt.Println("awesome")
		fallthrough
	case angka4 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// Seleksi Kondisi Bersarang
	var angka5 = 3
	if angka5 <= 10 {
		if angka5 > 7 {
			switch angka5 {
			case 10:
				fmt.Println("perfect!")
			default:
				fmt.Println("nice!")
			}
		} else {
			if angka5 == 5 {
				fmt.Println("not bad")
			} else if (angka5 < 5) && (angka5 > 2) {
				fmt.Println("keep trying")
			} else {
				fmt.Println("you can do it")
				if angka5 == 0 {
					fmt.Println("try harder!")
				}
			}
		}
	} else {
		fmt.Println("Tidak Boleh Lebih dari angka 10")
	}
}
