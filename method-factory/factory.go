package methodfactory

import "github.com/punkestu/method-factory-task/method-factory/product"

// interface untuk hasil dari method factory
type IMakanan interface {
	Buat()
	Makan()
}

type Restoran struct {
}

const (
	AYAM_GORENG = iota
	NASI_PECEL
)

// method yang digunakan untuk menghasilkan objek makanan
func (r *Restoran) Pesan(tipe int) IMakanan {
	switch tipe {
	case AYAM_GORENG:
		return &product.AyamGoreng{}
	case NASI_PECEL:
		return &product.NasiPecel{}
	default:
		return nil
	}
}
