package main

import methodfactory "github.com/punkestu/method-factory-task/method-factory"

func main() {
	var restoran = methodfactory.Restoran{}
	var ayamGoreng = restoran.Pesan(methodfactory.AYAM_GORENG)
	ayamGoreng.Buat()
	ayamGoreng.Makan()

	var nasiPecel = restoran.Pesan(methodfactory.NASI_PECEL)
	nasiPecel.Buat()
	nasiPecel.Makan()
}
