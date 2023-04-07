package main

import "fmt"

// method main() method eksekusi golang
func main() {

	/**
	Go memiliki aturan unik yang jarang dimiliki bahasa lain, yaitu tidak boleh ada satupun variabel yang menganggur.
	Artinya, semua variabel yang dideklarasikan harus digunakan. Jika ada variabel yang tidak digunakan tapi dideklarasikan,
	error pada Compailer dan program tidak akan bisa di-run.

	Skema penggunaan keyword var:
	var <nama-variabel> <tipe-data>
	var <nama-variabel> <tipe-data> = <nilai>

	command perintah, go run namafile.go
	go run variable.go // menjalankan file.go
	*/

	// deklarasi variable dengan keyword var
	// var <nama-variabel> <tipe-data> = <nilai>
	var firstName string = "budhi"
	var lastName string = "octaviansyah"

	fmt.Println("halo dunia!") // Println(a ...any) (n int, err error) // Println() // cetak void
	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName) // Printf(format string, a ...any) (n int, err error) // Printf() // cetak dengan alias
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// deklarasi variable dengan keyword var dan operator :=
	// key (var) dan operator (:=)
	var namaDepan = "slamet" // menggunakan key var
	namaBelakang := "wiguna" // tidak menggunakan key var
	fmt.Printf("halo %s %s!\n", namaDepan, namaBelakang)

	// operator (:=) hanya digunakan sekali di awal pada saat deklarasi. Untuk assignment nilai selanjutnya harus menggunakan tanda =, contoh:
	// jika menggunakan variable yang sama akan Error: No new variables on the left side of ':='
	nama := "intan"
	nama = "juleha"
	nama = "marijem"
	fmt.Printf("halo %s!\n", nama) // Printf()

	// deklarasi multi variable
	var pertama, kedua, ketiga string
	pertama, kedua, ketiga = "pertama", "kedua", "ketiga"
	fmt.Printf("halo %s %s %s!\n", pertama, kedua, ketiga)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Printf("halo %s %s %s!\n", fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Printf("halo %s %s %s!\n", seventh, eight, ninth)

	/**
	Format alias Printf()
	%s -> String
	%d -> nilai Bulat Global
	%x -> nilai Bulat
	%f -> nilai Pecahan
	%t -> Bolean
	*/

	nilaibulat, kondisi, nilaipecahan, setering := 1, true, 2.0, "hello"
	fmt.Printf("halo %x %t %f %s!\n", nilaibulat, kondisi, nilaipecahan, setering)

	// deklarasi variable dengan keyword _
	// Variabel Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
	_ = "Belajar Golang"
	_ = "Makan Bakso"
	//ismail, _ := "ismail", "wick" // Variabel underscore adalah predefined. tidak perlu :=.. cukup  dengan  = saja sudah bisa

	// deklarasi variable dengan keyword new
	ismail := new(string)
	fmt.Println(ismail)  // 0xc000088370 // menampilkan alamat memory string
	fmt.Println(*ismail) // "" // * tanda asterik // tidak menampilkan apapun

	/*
		// deklarasi variable dengan keyword make
		- channel
		- slice
		- map
	*/

}
