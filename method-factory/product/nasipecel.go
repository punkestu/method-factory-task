package product

import "fmt"

// salah satu implementasi dari interface makanan
type NasiPecel struct {
}

func (n *NasiPecel) Buat() {
	fmt.Println(">> Membuat nasi pecel")
	fmt.Println("+  Siapkan nasi")
	fmt.Println("+  Siapkan pecel")
	fmt.Println("+  Siram dengan bumbu kacang")
}

func (n *NasiPecel) Makan() {
	fmt.Println(">> Makan nasi pecel")
}
