package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetShelfLifeDays(name, category string) (int, error) {
	apiKey := os.Getenv("OPEN_ROUTER_API_KEY")
	if apiKey == "" {
		return 7, fmt.Errorf("OPEN_ROUTER_API_KEY not set")
	}

	url := "https://openrouter.ai/api/v1/chat/completions"

	prompt := fmt.Sprintf(
		`Kamu adalah ahli keamanan pangan yang sangat ketat.

Tugas:
Tentukan masa simpan aman (dalam hari) untuk makanan atau minuman yang disimpan DI DALAM KULKAS dengan suhu konstan 4°C (bukan suhu ruang, bukan freezer).

Input:
- Nama makanan/minuman: "%s"
- Kategori: "%s"

Aturan Utama:
- Kembalikan HANYA angka (jumlah hari). Tanpa teks tambahan.
- Semua asumsi HARUS berdasarkan penyimpanan di kulkas (4°C).
- Gunakan standar keamanan pangan konservatif (pilih durasi paling aman/pendek).
- Jika ragu, pilih nilai yang lebih kecil (lebih aman).

Asumsi Penyimpanan:
- Disimpan di kulkas (bukan freezer, bukan suhu ruang).
- Dalam wadah tertutup.
- Kondisi higienis normal rumah tangga.

Aturan Kondisi:
- Makanan mentah → pendek (1–3 hari jika berisiko tinggi).
- Makanan matang/olahan → 2–5 hari.
- Makanan potong/kupas → lebih cepat rusak.
- Produk kemasan dibuka → lebih cepat rusak.
- Produk kemasan belum dibuka → lebih lama tapi tetap konservatif.

Kasus Khusus (WAJIB):
- Non-perishable (madu, gula, garam) → 3650
- Air dan es (termasuk es batu) → 3650
- Alkohol:
  - Belum dibuka → 3650
  - Sudah dibuka → 30–180 (pilih konservatif)
- Minuman susu → 3–5 hari setelah dibuka
- Minuman segar → 1–3 hari
- Minuman kemasan dibuka → 3–7 hari
- Makanan fermentasi → 7–30 hari
- Makanan beku yang dicairkan → perlakukan sebagai makanan segar

Panduan Kategori:
- sayuran → 5–14 hari
- buah → 3–10 hari
- makanan_matang → 2–5 hari
- daging_mentah → 1–3 hari
- seafood → 1–2 hari
- telur → 21–35 hari
- produk_susu → 5–14 hari
- roti_kue → 3–7 hari
- saus_bumbu → 30–180 hari
- makanan_fermentasi → 7–30 hari
- makanan_beku_cair → 1–3 hari
- minuman → 1–7 hari
- bahan_kering → 30–180 hari

Fallback:
- Jika tidak diketahui, gunakan kategori terdekat dan pilih nilai paling aman.

Output:
HANYA angka (integer).`,
		name, category,
	)

	reqBody := map[string]interface{}{
		"model": "openai/gpt-4o-mini",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)

	if err != nil {
		return 7, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 7, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 7, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return 7, fmt.Errorf("API error: %s", string(bodyBytes))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return 7, err
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	content := message["content"].(string)

	clean := strings.TrimSpace(content)
	days, err := strconv.Atoi(clean)
	if err != nil {
		return 7, fmt.Errorf("invalid response: %s", clean)
	}

	if days <= 0 {
		return 7, nil
	}

	return days, nil
}
