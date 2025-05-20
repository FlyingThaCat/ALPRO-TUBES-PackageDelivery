package paket

import (
	"PackageDelivery/datas"
	"PackageDelivery/types"
	"PackageDelivery/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


func UbahPaket(paket types.Paket) {
	utils.ClearScreen()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Ubah Paket ===")
	
	inputWithDefault := func(prompt string, current string) string {
		fmt.Printf("%s [%s]: ", prompt, current)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			return current
		}
		return text
	}

	// Tipe
	paket.Tipe = inputWithDefault("Tipe", string(paket.Tipe))

	// Berat (float)
	beratStr := inputWithDefault("Berat", fmt.Sprintf("%.2f", paket.Berat))
	fmt.Sscanf(beratStr, "%f", &paket.Berat)

	// Harga (float)
	hargaStr := inputWithDefault("Harga", fmt.Sprintf("%.0f", paket.Harga))
	fmt.Sscanf(hargaStr, "%f", &paket.Harga)

	// Kota Pengirim
	paket.SenderCity = inputWithDefault("Kota Pengirim", paket.SenderCity)

	// Kota Tujuan
	paket.ReceiverCity = inputWithDefault("Kota Tujuan", paket.ReceiverCity)

	// Status (untuk contoh ini kita tampilkan sebagai satu string, tidak array)
	statusStr := inputWithDefault("Status (pisah dengan koma jika banyak)", strings.Join(paket.Status, ","))
	paket.Status = strings.Split(statusStr, ",")

	// Update waktu diperbarui
	paket.UpdatedAt = time.Now()

	// Simpan perubahan ke database (simulasi)
	for i, p := range datas.PaketDB {
		if p.NoResi == paket.NoResi {
			datas.PaketDB[i] = paket
			break
		}
	}

	utils.ClearScreen()
	fmt.Println("\nPaket berhasil diubah:")
	fmt.Printf("No Resi: %s\n", paket.NoResi)
	fmt.Printf("Tipe: %s\n", paket.Tipe)
	fmt.Printf("Berat: %.2f\n", paket.Berat)
	fmt.Printf("Harga: %.0f\n", paket.Harga)
	fmt.Printf("Kota Pengirim: %s\n", paket.SenderCity)
	fmt.Printf("Kota Tujuan: %s\n", paket.ReceiverCity)
	fmt.Printf("Status: %v\n", paket.Status)
	fmt.Printf("Dibuat pada: %s\n", paket.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("Diperbarui pada: %s\n", paket.UpdatedAt.Format("2006-01-02 15:04:05"))
	utils.EnterToContinue()
}
