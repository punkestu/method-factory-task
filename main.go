package main

import methodfactory "github.com/punkestu/method-factory-task/method-factory"

func main() {
	// membuat objek factory
	var restoran = methodfactory.Restoran{}

	// membuat objek makanan dengan parameter AYAM_GORENG
	var ayamGoreng = restoran.Pesan(methodfactory.AYAM_GORENG)
	// panggil method sesuai dengan interface
	ayamGoreng.Buat()
	ayamGoreng.Makan()

	// membuat objek makanan dengan parameter NASI_PECEL
	var nasiPecel = restoran.Pesan(methodfactory.NASI_PECEL)
	// panggil method sesuai dengan interface
	nasiPecel.Buat()
	nasiPecel.Makan()
}
