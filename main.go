package main // File yang memiliki package main akan di eksekusi pertama kali ketika program dijalankan

import "fmt"  // package fmt memiliki fungsi untuk keperluan I/O

func main(){ // Fungsi main harus ada didalam program yang memiliki file package main
	fmt.Println("Hello World") // Digunakan untuk memunculkan text ke layar
}