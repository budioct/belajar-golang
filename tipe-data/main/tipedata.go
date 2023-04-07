package main

import (
	"fmt"
)

func main() {

	/**
	Tipe Data
	- numerik,
	- (desimal & non-desimal),
	- string,
	- boolean.

	Tipe Data Numerik Non-Desimal
	 - uint, tipe data untuk bilangan cacah (bilangan positif).
	 - int, tipe data untuk bilangan bulat (bilangan negatif dan positif).

	Tipe data	Cakupan bilangan
	uint8		0 ↔ 255
	uint16		0 ↔ 65535
	uint32		0 ↔ 4294967295
	uint64		0 ↔ 18446744073709551615
	uint		sama dengan uint32 atau uint64 (tergantung nilai)
	byte		sama dengan uint8
	int8		-128 ↔ 127
	int16		-32768 ↔ 32767
	int32		-2147483648 ↔ 2147483647
	int64		-9223372036854775808 ↔ 9223372036854775807
	int	s		ama dengan int32 atau int64 (tergantung nilai)
	rune		sama dengan int32
	*/

	// Penggunaan Fungsi fmt.Print()
	fmt.Println("sadam huesin")
	fmt.Println("sadam", "huesin")

	fmt.Print("sadam husein\n")
	fmt.Print("sadam", "husein\n")
	fmt.Print("sadam", " ", "husein\n")

	// Konstanta = Final (tidak bisa diubah sejak pertama kali dibuat) // cara buatnya dengan keyword const
	const data = "ismail"
	fmt.Print("hello ", data, " broo!\n") // Println(a ...any) (n int, err error)

	// teknik type inference
	const data2 = "badrul"
	fmt.Println("nice to meet you", data2, "broo!") // Println(a ...any) (n int, err error)

	// Deklarasi Multi Konstanta = Final (tidak bisa diubah sejak pertama kali dibuat) // cara buatnya dengan keyword const
	const (
		square   string  = "kotak"
		hariIni  bool    = true
		numberic uint8   = 1
		floatNum float32 = 2.2
	)

	// Deklarasi Konstanta dengan tipe data dan nilai yang sama
	// tipe data dan nilai tidak dituliskan dalam deklarasi konstanta, maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
	const (
		a = "konstanta"
		b
	)

	// Deklarasi Konstanta gabungan
	const (
		today string = "jum'at"
		now
		isToday = true
	)

	// Deklarasi Konstanta dalam satu baris
	const satu, dua = 1, 2
	const jumat, sabtu string = "jum'at", "sabtu"

	fmt.Printf("Konstanta:  %x  %x \n", satu, dua)
	fmt.Printf("Konstanta:  %s  %s \n", jumat, sabtu)

}
