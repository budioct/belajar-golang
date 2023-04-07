package main

import "fmt"

func main() {

	/**
	Operator
	operator yang bisa digunakan di .Go Secara umum terdapat 3 kategori operator:
	- aritmatika,
	- perbandingan,
	- logika

	Operator Aritmatika // result number
	Tanda	Penjelasan
	+		penjumlahan
	-		pengurangan
	*		perkalian
	/		pembagian
	%		modulus / sisa hasil pembagian

	Operator Perbandingan // result boolean
	Tanda	Penjelasan
	==		apakah nilai kiri sama dengan nilai kanan
	!=		apakah nilai kiri tidak sama dengan nilai kanan
	<		apakah nilai kiri lebih kecil daripada nilai kanan
	<=		apakah nilai kiri lebih kecil atau sama dengan nilai kanan
	>		apakah nilai kiri lebih besar dari nilai kanan
	>=		apakah nilai kiri lebih besar atau sama dengan nilai kanan

	// Operator Logika
	Tanda	Penjelasan
	&&		kiri dan kanan
	||		kiri atau kanan
	!		negasi / nilai kebalikan

	*/

	// Operator Aritmatika
	// secara default pembuatan (var nameVariable value) number, demikian tipe data nya adalah int8 ~ int64
	var penjumlahan = 10 + 5
	var pengurangan = 10 - 1000
	var perkalian = 10 * 5
	var pembagian = 10 / 5
	var moduluse = 10 % 5
	fmt.Println("hasil penjumlahan: ", penjumlahan)
	fmt.Println("hasil pengurangan: ", pengurangan)
	fmt.Println("hasil perkalian: ", perkalian)
	fmt.Println("hasil pembagian: ", pembagian)
	fmt.Println("hasil moduluse: ", moduluse)

	// Operator Perbandingan
	var value1 = 10
	var value2 = 5

	fmt.Printf("operator > : %d > %x hasil= %t \n", value1, value2, value1 > value2)
	fmt.Printf("operator >= : %d >= %x hasil= %t \n", value1, value2, value1 >= value2)
	fmt.Printf("operator < : %d < %x hasil= %t \n", value1, value2, value1 < value2)
	fmt.Printf("operator <= : %d >= %x hasil= %t \n", value1, value2, value1 <= value2)
	fmt.Printf("operator == : %d == %x hasil= %t \n", value1, value2, value1 == value2)
	fmt.Printf("operator != : %d != %x hasil= %t \n", value1, value2, value1 != value2)

	// Operator Logika
	var satu = true
	var nol = false

	fmt.Printf("operator AND && : %t > %t hasil= %t \n", satu, nol, satu && nol)
	fmt.Printf("operator OR || : %t > %t hasil= %t \n", satu, nol, satu || nol)
	fmt.Printf("operator NEGASI ! : %t hasil= %t \n", satu, !satu)
	fmt.Printf("operator NEGASI ! : %t hasil= %t \n", nol, !nol)

}
