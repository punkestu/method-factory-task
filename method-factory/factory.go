package methodfactory

import "github.com/punkestu/method-factory-task/method-factory/product"

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
