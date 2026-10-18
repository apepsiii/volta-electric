package repository

import (
	"voltra-electric/internal/domain"
)

// Repository defines data access interface
type Repository interface {
	GetCompanyProfile() domain.CompanyProfile
	GetMetrics() []domain.MetricStat
	GetProducts() []domain.Product
	GetProductByID(id string) (*domain.Product, bool)
	GetComponents() []domain.ComponentDetail
	GetTestingProcedures() []domain.TestingProcedure
}

type memoryRepo struct {
	company    domain.CompanyProfile
	metrics    []domain.MetricStat
	products   []domain.Product
	components []domain.ComponentDetail
	testing    []domain.TestingProcedure
}

// NewMemoryRepository initializes and populates data from PT Mega Karya Perkasa PDF profile
func NewMemoryRepository() Repository {
	company := domain.CompanyProfile{
		LegalName:       "PT. MEGA KARYA PERKASA",
		BrandName:       "VOLTRA",
		EstablishedYear: 2008,
		CoreBusiness:    "Industri Trafo Distribusi berpendingin minyak dan kelengkapannya",
		Standards:       []string{"Standar IEC (International Electrotechnical Commission)", "Standar Industri Indonesia (SPLN)"},
		Mission:         "Menghasilkan produk berkualitas tinggi sesuai standar Nasional dan Internasional demi kepuasan pelanggan melalui keseimbangan antara kualitas dan kinerja.",
		QualitySystem:   "Sistem pemeriksaan on-line telah diterapkan pada setiap tahap produksi.",
	}

	metrics := []domain.MetricStat{
		{Value: "2008", Label: "Tahun Berdiri", Sublabel: "Pengalaman Manufaktur Trafo", Highlight: "PT Mega Karya Perkasa"},
		{Value: "≥ 50 kV", Label: "Dielektrik Minyak", Sublabel: "Minimum Ketahanan Tembus", Highlight: "Garansi Awet & Aman"},
		{Value: "≥ 80 µm", Label: "Lapisan Proteksi Cat", Sublabel: "Ketebalan Standar Anti-Karat", Highlight: "Tahan Iklim Tropis"},
		{Value: "SPLN & IEC", Label: "Standarisasi Resmi", Sublabel: "Sertifikasi Uji Laboratorium", Highlight: "100% On-Line QC"},
	}

	products := []domain.Product{
		{
			ID:          "hermetik",
			Name:        "Trafo Hermetik Fin Wall",
			Series:      "VOLTRA H-Series",
			Tag:         "Corrugated Radiator",
			Category:    "hermetik",
			Subtitle:    "Bebas Perawatan Rutin • Desain Kompak & Kedap Udara",
			Description: "Trafo distribusi berpendingin minyak tipe hermetik tanpa tangki konservator. Sirip gelombang (fin wall) fleksibel mengembang dan mengempis mengikuti ekspansi minyak saat temperatur naik/turun, mencegah kontak minyak dengan oksigen luar secara total.",
			Capacity:    "25 s/d 2500 kVA",
			ImageURL:    "/static/img/unit_hermetik.jpg",
			Badge:       "Zero Maintenance",
			Features: []string{
				"Radiator sirip gelombang (fin wall) kompak menghemat ruang gardu",
				"Kedap udara 100% sehingga oli tidak teroksidasi dan bebas penggantian silica gel",
				"Lapisan cat anti-karat minimum 80 mikron untuk lingkungan luar ruangan",
				"Memenuhi standar SPLN D3.002:2014 & IEC 60076",
			},
			Specs: map[string]string{
				"Tipe Pendingin": "ONAN (Oil Natural Air Natural)",
				"Tegangan Primer": "20.000 V (20 kV)",
				"Tegangan Sekunder": "400 V / 231 V",
				"Vektor Group": "Dyn5 / Yyn0",
				"Kekuatan Dielektrik": "≥ 50 kV",
			},
		},
		{
			ID:          "conservator",
			Name:        "Trafo Conservator Panel",
			Series:      "VOLTRA C-Series",
			Tag:         "Panel Radiator",
			Category:    "conservator",
			Subtitle:    "Dilengkapi Tangki Konservator • Andal untuk Beban Berat",
			Description: "Trafo distribusi yang dilengkapi tangki konservator di bagian atas dan silica gel dehydrating breather. Sangat cocok untuk kapasitas daya besar, gardu outdoor, dan kebutuhan operasional industri dengan fluktuasi beban tinggi.",
			Capacity:    "100 s/d 3150 kVA",
			ImageURL:    "/static/img/unit_conservator.jpg",
			Badge:       "Heavy Duty",
			Features: []string{
				"Indikator level oli visual jelas untuk pemantauan volume minyak berkala",
				"Dilengkapi silica gel breather untuk menyaring uap air udara masuk",
				"Panel radiator modular kokoh dengan laju pendinginan maksimal",
				"Dapat dilengkapi proteksi Buchholz Relay atau DGPT2",
			},
			Specs: map[string]string{
				"Tipe Pendingin": "ONAN / ONAF",
				"Tegangan Primer": "20 kV / Multi-tappings",
				"Tegangan Sekunder": "400 V / 380 V",
				"Proteksi Korosi": "Washing + Coating + Cat ≥ 80 µm",
				"Standar Uji": "IEC & SPLN Certified",
			},
		},
		{
			ID:          "nitrogen",
			Name:        "Trafo Hermetik Gas Nitrogen",
			Series:      "VOLTRA N-Series",
			Tag:         "N2 Gas Cushion",
			Category:    "nitrogen",
			Subtitle:    "Bantalan Gas Nitrogen • Proteksi Oksidasi Maksimal",
			Description: "Trafo hermetik yang menggunakan bantalan gas Nitrogen kering pada ruang kosong di atas permukaan minyak. Nitrogen sebagai gas inert mencegah degradasi oksidasi minyak trafo dan penyerapan kelembaban dari lingkungan luar 100%.",
			Capacity:    "Custom Sesuai Pesanan",
			ImageURL:    "/static/img/unit_nitrogen.jpg",
			Badge:       "Ultra Protection",
			Features: []string{
				"Bantalan gas nitrogen mencegah penuaan dini minyak isolasi trafo",
				"Sangat tahan terhadap iklim tropis dengan kelembaban tinggi dan area pesisir",
				"Dilengkapi pressure relief device dan katup pengaman tekanan gas",
				"Uji kebocoran tangki bertekanan ganda sebelum pengiriman",
			},
			Specs: map[string]string{
				"Gas Pelindung": "Nitrogen Kering (Dry Nitrogen Gas)",
				"Kekuatan Dielektrik": "≥ 50 kV",
				"Material Tangki": "Plat Baja Hitam Cold/Hot Rolled",
				"Uji Tekanan": "Pressure Test 100% Lulus",
				"Sertifikasi": "Standar IEC & SPLN",
			},
		},
		{
			ID:          "custom",
			Name:        "Trafo Custom Industri & PLN",
			Series:      "VOLTRA Spec-Series",
			Tag:         "SPLN Standard",
			Category:    "custom",
			Subtitle:    "Desain Presisi Sesuai Spesifikasi Teknis Tender & Gardu",
			Description: "Layanan manufaktur trafo khusus oleh tim R&D PT Mega Karya Perkasa. Kami merancang trafo sesuai spesifikasi tender PLN, gardu pabrik, pertambangan, dan pembangkit dengan rugi daya rendah (low loss) dan keandalan tinggi.",
			Capacity:    "25 s/d 2500+ kVA",
			ImageURL:    "/static/img/unit_custom.jpg",
			Badge:       "Custom Order",
			Features: []string{
				"Pilihan belitan kawat tembaga murni (Cu) atau alumunium murni (Al)",
				"Inti besi Cold Rolled Grain Oriented Silicon Sheet kelas tertinggi",
				"Oil duct presisi untuk disipasi panas optimal dan ketahanan hubung singkat",
				"Didukung sertifikat Factory Acceptance Test (FAT) 9 parameter uji lengkap",
			},
			Specs: map[string]string{
				"Tegangan": "Sesuai Pesanan (Multi-ratio)",
				"Material Inti": "CRGO Silicon Sheet Low Loss",
				"Bahan Belitan": "Tembaga (Cu) / Alumunium (Al)",
				"Jaminan Mutu": "Pengawasan On-Line & Garansi Resmi",
				"Uji Laboratorium": "FAT 9 Tahap Lengkap",
			},
		},
	}

	components := []domain.ComponentDetail{
		{
			ID:          "tank",
			Title:       "Tanki Baja Presisi & Radiator Fin Wall",
			Subtitle:    "Struktur Mekanikal & Sistem Pendingin",
			Category:    "Mekanikal",
			Description: "Terbuat dari plat baja hitam (Cold Rolled atau Hot Rolled) dengan fin wall (Corrugated Radiator) atau Panel Radiator yang disesuaikan dengan kapasitas trafo. Setiap tangki melewati 4 tahap proteksi ketat sebelum perakitan.",
			KeyPoints: []string{
				"Pemeriksaan kebocoran bertekanan untuk memastikan tangki 100% kedap",
				"Pencucian kimiawi (washing) dan coating threadment pelindung permukaan",
				"Pengecatan tahan karat dengan ketebalan minimum 80 mikron",
				"Pilihan tangki conservator, hermetik, atau hermetik dengan gas nitrogen",
			},
			ImageURLs: []string{
				"/static/img/process_clamp_measuring.jpg",
				"/static/img/process_fin_grinding.jpg",
				"/static/img/process_tank_welding.jpg",
				"/static/img/process_spray_painting.jpg",
			},
		},
		{
			ID:          "core-coil",
			Title:       "Inti Besi (Core) & Kumparan (Coil)",
			Subtitle:    "Sistem Elektrikal & Efisiensi Induksi",
			Category:    "Elektrikal",
			Description: "Didesain dengan keseimbangan antara kualitas dan kinerja menggunakan seleksi bahan berkualitas tinggi untuk menghasilkan efisiensi induksi maksimal dan rugi besi yang sangat rendah.",
			KeyPoints: []string{
				"Inti besi terbuat dari Cold Rolled Grain Oriented Silicon Sheet dengan rugi histeresis rendah",
				"Kumparan menggunakan bahan tembaga murni (Cu) dan alumunium (Al) murni",
				"Kawat bulat berenamel untuk tegangan tinggi & lembaran berisolasi kertas untuk tegangan rendah",
				"Oil Duct (saluran sirkulasi minyak) kompak dirancang cermat untuk transfer panas dan penyekat hubung singkat",
			},
			ImageURLs: []string{
				"/static/img/core_silicon.jpg",
				"/static/img/coil_winding.jpg",
			},
		},
		{
			ID:          "insulation",
			Title:       "Bahan Isolasi & Minyak Berkualitas Tinggi",
			Subtitle:    "Proteksi Dielektrik & Ketahanan Arus Hubung Singkat",
			Category:    "Isolasi",
			Description: "Material isolasi dipilih khusus untuk memperkuat daya tahan kumparan saat terjadi lonjakan arus hubung singkat dinamis dan menjaga kestabilan dielektrik jangka panjang.",
			KeyPoints: []string{
				"Kertas isolasi (Insulation Paper) dan Diamond Dotted Paper sebagai media isolasi utama belitan",
				"Press Board berkualitas tinggi sebagai penyangga kokoh antara kumparan dan core clamp",
				"Isolasi minyak jenis mineral, alami, atau silikon dengan titik api tinggi",
				"Kekuatan dielektrik minyak minimum 50 kV untuk menjamin garansi trafo tahan lama",
			},
			ImageURLs: []string{
				"/static/img/insulation_paper.jpg",
			},
		},
		{
			ID:          "rd",
			Title:       "Penelitian & Pengembangan (R&D)",
			Subtitle:    "Didukung Tenaga Ahli Berpengalaman",
			Category:    "Engineering",
			Description: "Setiap bahan utama trafo—seperti inti besi, kumparan, bahan isolasi, penyangga inti, dan tangki—didesain secara tepat dan diawasi terus-menerus oleh tim ahli yang berpengalaman di bidangnya.",
			KeyPoints: []string{
				"Simulasi dan kalkulasi elektromagnetik presisi sebelum proses manufaktur",
				"Seleksi ketat bahan baku bersertifikasi uji lab",
				"Pengawasan on-line di setiap tahapan produksi dari nol hingga unit jadi",
				"Penyempurnaan desain kontinu sesuai perkembangan standar IEC & SPLN",
			},
			ImageURLs: []string{
				"/static/img/rd_team.jpg",
			},
		},
	}

	testing := []domain.TestingProcedure{
		{
			StepNumber:  1,
			Code:        "FAT-01",
			Name:        "Visual, Pembuatan dan Dimensi",
			Description: "Pemeriksaan menyeluruh terhadap fisik eksterior tangki, kelurusan sirip radiator fin wall, pengecatan, kelengkapan nameplate, dan toleransi dimensi gardu.",
			Standard:    "SPLN & IEC 60076",
			IsCrucial:   false,
		},
		{
			StepNumber:  2,
			Code:        "FAT-02",
			Name:        "Tahanan Belitan (Winding Resistance)",
			Description: "Pengukuran resistansi DC pada setiap fasa kumparan sisi tegangan tinggi (HV) dan tegangan rendah (LV) untuk memastikan keseragaman belitan.",
			Standard:    "IEC 60076-1",
			IsCrucial:   false,
		},
		{
			StepNumber:  3,
			Code:        "FAT-03",
			Name:        "Perbandingan Rasio Belitan (Turns Ratio)",
			Description: "Pengujian perbandingan transformasi tegangan pada setiap posisi tap changer dan verifikasi sudut pergeseran fasa (Vector Group Dyn5/Yyn0).",
			Standard:    "IEC 60076-1",
			IsCrucial:   false,
		},
		{
			StepNumber:  4,
			Code:        "FAT-04",
			Name:        "Kekuatan Dielektrik Oli (≥ 50 kV)",
			Description: "Uji tegangan tembus minyak isolasi (dielectric breakdown test) dengan batas minimum 50 kV untuk menjamin keamanan isolasi listrik internal.",
			Standard:    "IEC 60156 / SPLN",
			IsCrucial:   true,
		},
		{
			StepNumber:  5,
			Code:        "FAT-05",
			Name:        "Rugi Beban Nol & Arus Beban Nol",
			Description: "Pengukuran losses inti besi CRGO (no-load loss) dan arus magnetisasi saat kumparan diberi tegangan nominal tanpa beban.",
			Standard:    "SPLN D3.002:2014",
			IsCrucial:   false,
		},
		{
			StepNumber:  6,
			Code:        "FAT-06",
			Name:        "Rugi Beban Penuh & Impedansi",
			Description: "Pengujian rugi tembaga (load loss) dan persentase tegangan impedansi (%Z) pada arus pengenal untuk efisiensi distribusi daya.",
			Standard:    "IEC 60076-1",
			IsCrucial:   false,
		},
		{
			StepNumber:  7,
			Code:        "FAT-07",
			Name:        "Tegangan Terapan (Applied Voltage Test)",
			Description: "Uji ketahanan dielektrik isolasi kumparan terhadap bumi dan antar kumparan menggunakan sumber tegangan tinggi frekuensi daya eksternal.",
			Standard:    "IEC 60076-3",
			IsCrucial:   true,
		},
		{
			StepNumber:  8,
			Code:        "FAT-08",
			Name:        "Tegangan Induksi Lebih (Induced Overvoltage)",
			Description: "Pengujian ketahanan isolasi antar lilitan kumparan (inter-turn withstand test) dengan frekuensi dua kali lipat tegangan nominal.",
			Standard:    "IEC 60076-3",
			IsCrucial:   false,
		},
		{
			StepNumber:  9,
			Code:        "FAT-09",
			Name:        "Tes Kebocoran Tangki Bertekanan",
			Description: "Pengujian tekanan tangki (pressure test) dengan minyak dan udara bertekanan untuk menjamin 0% risiko rembesan atau kebocoran minyak saat operasi.",
			Standard:    "SPLN & Factory Standard",
			IsCrucial:   true,
		},
	}

	return &memoryRepo{
		company:    company,
		metrics:    metrics,
		products:   products,
		components: components,
		testing:    testing,
	}
}

func (r *memoryRepo) GetCompanyProfile() domain.CompanyProfile {
	return r.company
}

func (r *memoryRepo) GetMetrics() []domain.MetricStat {
	return r.metrics
}

func (r *memoryRepo) GetProducts() []domain.Product {
	return r.products
}

func (r *memoryRepo) GetProductByID(id string) (*domain.Product, bool) {
	for _, p := range r.products {
		if p.ID == id {
			return &p, true
		}
	}
	return nil, false
}

func (r *memoryRepo) GetComponents() []domain.ComponentDetail {
	return r.components
}

func (r *memoryRepo) GetTestingProcedures() []domain.TestingProcedure {
	return r.testing
}

