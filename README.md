# VOLTRA Electric (voltra-electric.id)

Aplikasi web landing page performa tinggi untuk **[voltra-electric.id](https://voltra-electric.id)** yang dibangun menggunakan **Golang** dengan framework **Echo v4**, arsitektur modular yang rapi dan terorganisir (*Clean Architecture*), serta data resmi dari profil perusahaan **PT. MEGA KARYA PERKASA** (Merek Dagang **VOLTRA**).

---

## 🏗️ Struktur Folder Proyek

```
volta-electric/
├── bin/                          # Binary hasil kompilasi
│   └── voltra-server.exe
├── cmd/
│   └── web/
│       └── main.go               # Entry point aplikasi & graceful shutdown
├── config/
│   └── config.go                 # Konfigurasi aplikasi & environment
├── internal/
│   ├── delivery/
│   │   └── http/
│   │       ├── handler.go        # HTTP Handlers (HTML & REST API)
│   │       ├── renderer.go       # Custom Echo html/template renderer
│   │       └── router.go         # Registrasi rute & middleware Echo
│   ├── domain/
│   │   ├── company.go            # Model data entitas perusahaan
│   │   ├── product.go            # Model data produk trafo & pengujian FAT
│   │   └── inquiry.go            # Model kalkulator penawaran & generator WhatsApp
│   └── repository/
│       └── memory.go             # Data resmi dari PDF PT Mega Karya Perkasa
├── web/
│   ├── static/                   # Aset publik statis
│   │   ├── css/
│   │   │   └── style.css         # Custom styling
│   │   ├── js/
│   │   │   └── app.js           # Filter dinamis, modal API, & handler WhatsApp
│   │   └── img/                  # Foto pabrik & trafo resolusi tinggi dari PDF
│   └── templates/                # Template HTML terfragmentasi rapi
│       ├── layouts/
│       │   └── base.html         # Base layout master
│       ├── pages/
│       │   └── index.html        # Halaman utama komposisi
│       └── partials/             # Komponen antarmuka modular
│           ├── navbar.html       # Navigasi sticky & top notice
│           ├── hero.html         # Hero showcase banner
│           ├── profile.html      # Profil resmi PT Mega Karya Perkasa
│           ├── products.html     # Kartu katalog produk trafo VOLTRA
│           ├── components.html   # Inti CRGO, kumparan, tangki & isolasi
│           ├── testing.html      # 9 Tahap pengujian mutu FAT bersertifikat
│           ├── calculator.html   # Kalkulator spesifikasi & konsultasi
│           ├── modal.html        # Pop-up rincian spesifikasi teknis
│           └── footer.html       # Footer korporat tema gelap
├── .env.example                  # Template variabel lingkungan
├── go.mod                        # Modul Go & dependencies
├── go.sum                        # Checksum hash dependensi
└── README.md
```

---

## ⚡ Cara Menjalankan Aplikasi

### 1. Prasyarat
- Go versi 1.21+ (teruji pada Go 1.25.6)

### 2. Menjalankan Langsung (Development)
```bash
go run ./cmd/web
```
Server akan aktif di: `http://localhost:8080`

### 3. Kompilasi Binary (Production)
```bash
# Windows
go build -o bin/voltra-server.exe ./cmd/web

# Linux / Server Production
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/voltra-server ./cmd/web
```

---

## 📡 Rute & REST API Tersedia

| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/` | Merender landing page utama |
| `GET` | `/health` | Healthcheck status server |
| `GET` | `/api/products` | Mengambil seluruh katalog produk trafo (JSON) |
| `GET` | `/api/products/:id` | Mengambil detail teknis trafo untuk modal popup (JSON) |
| `POST` | `/api/inquiry` | Memproses permintaan spesifikasi & membuat link WhatsApp |
| `GET` | `/static/*` | Menyajikan berkas statis (gambar, CSS, JS) |

---

## 📋 Data Profil PT Mega Karya Perkasa yang Terintegrasi
- **Nama Legal**: PT. MEGA KARYA PERKASA
- **Merek Dagang**: VOLTRA
- **Tahun Pendirian**: 2008
- **Bidang Usaha**: Trafo Distribusi berpendingin minyak dan kelengkapannya
- **Standar**: IEC (International Electrotechnical Commission) & SPLN (Standar PLN)
- **Kekuatan Dielektrik Minyak**: Minimum 50 kV
- **Proteksi Ketebalan Cat**: Minimum 80 mikron
- **9 Rangkaian Uji Pabrik (FAT)**:
  1. Visual, pembuatan dan dimensi
  2. Tahanan belitan (HV & LV)
  3. Rasio belitan & pergeseran fasa (Vector Group)
  4. Uji dielektrik minyak isolasi ($\ge$ 50 kV)
  5. Rugi beban nol & arus beban nol
  6. Rugi beban penuh & persentase impedansi
  7. Tegangan terapan (Applied voltage test)
  8. Tegangan induksi lebih (Induced overvoltage test)
  9. Tes kebocoran tangki bertekanan

