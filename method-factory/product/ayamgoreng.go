package product

import "fmt"

// salah satu implementasi dari interface makanan
type AyamGoreng struct {
}

func (a *AyamGoreng) Buat() {
	fmt.Println(">> Membuat ayam gorang")
	fmt.Println("+  Siapkan ayam")
	fmt.Println("+  Siapkan bumbu")
	fmt.Println("+  Goreng ayam")
}

func (a *AyamGoreng) Makan() {
	fmt.Println(">> Makan ayam goreng")
}
