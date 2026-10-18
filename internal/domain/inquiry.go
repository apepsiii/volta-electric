package domain

import (
	"fmt"
	"net/url"
)

// InquiryRequest holds calculation and quote request parameters
type InquiryRequest struct {
	ClientName  string `json:"client_name" form:"client_name"`
	TankType    string `json:"tank_type" form:"tank_type"`
	CapacityKVA string `json:"capacity_kva" form:"capacity_kva"`
	Winding     string `json:"winding" form:"winding"`
	Voltage     string `json:"voltage" form:"voltage"`
	Notes       string `json:"notes" form:"notes"`
}

// GenerateWhatsAppURL generates a pre-filled direct WhatsApp link
func (i *InquiryRequest) GenerateWhatsAppURL(waNumber string) string {
	client := i.ClientName
	if client == "" {
		client = "Pelanggan voltra-electric.id"
	}

	text := fmt.Sprintf(
		"Halo PT Mega Karya Perkasa,\n\nSaya (%s) ingin menanyakan penawaran harga & spesifikasi Trafo VOLTRA:\n- Tipe Tangki: %s\n- Kapasitas: %s\n- Bahan Kumparan: %s\n- Rasio Tegangan: %s\n\nCatatan Tambahan: %s\n\nMohon informasi ketersediaan teknis dan surat penawarannya. Terima kasih.",
		client,
		i.TankType,
		i.CapacityKVA,
		i.Winding,
		i.Voltage,
		i.Notes,
	)

	return fmt.Sprintf("https://wa.me/%s?text=%s", waNumber, url.QueryEscape(text))
}

