package main
import "fmt"


const NMAX = 20

type Klub struct {
	Nama           string
	Main           int
	Menang         int
	Seri           int
	Kalah          int
	GolMasuk       int
	GolKemasukan   int
	SelisihGol     int
	Poin           int
}

type Liga [NMAX]Klub

type Jadwal struct {
	Klub1, Klub2 string
	Pekan        int
}

type DaftarJadwal [1000]Jadwal


func cariKlub(liga Liga, nama string) int {
	var i int
	for i = 0; i < NMAX; i++ {
		if liga[i].Nama == nama {
			return i
		}
	}
	return -1
}


func hitungKlub(liga Liga) int {
	var sum int = 0
	var i int
	for i = 0; i < NMAX; i++ {
		if liga[i].Nama != "" {
			sum++
		}
	}
	return sum
}




func tambahKlub(liga *Liga) {
	var nama string
	var n int
	fmt.Print("Masukkan jumlah grup yang ingin diinput: ")
	fmt.Scan(&n)

	i := 0
	for i < n {
		fmt.Print("Masukkan nama klub (3 huruf): ")
		fmt.Scan(&nama)

		if cariKlub(*liga, nama) != -1 {
			fmt.Println("Klub sudah terdaftar.")
			i++
		} else {
			j := 0
			for j < NMAX && liga[j].Nama != "" {
				j++
			}

			if j < NMAX {
				liga[j].Nama = nama
				fmt.Println("Klub berhasil ditambahkan.")
				i++
			} else {
				fmt.Println("Kuota klub penuh.")
				i = n
			}
		}
	}
}


func ubahKlub(liga *Liga) {
	var namaLama, namaBaru string
	var idx int
	
	fmt.Print("Masukkan nama klub yang ingin diubah: ")
	fmt.Scan(&namaLama)
	fmt.Print("Masukkan nama baru (3 huruf): ")
	fmt.Scan(&namaBaru)

	idx = cariKlub(*liga, namaLama)
	if idx != -1 {
		liga[idx].Nama = namaBaru
		fmt.Println("Nama klub berhasil diubah.")
	} else {
		fmt.Println("Klub tidak ditemukan.")
	}
}


func hapusKlub(liga *Liga) {
	var nama string
	var idx int
	
	fmt.Print("Masukkan nama klub yang ingin dihapus: ")
	fmt.Scan(&nama)

	idx = cariKlub(*liga, nama)
	if idx != -1 {
		liga[idx].Nama = ""
		fmt.Println("Klub berhasil dihapus.")
	} else {
		fmt.Println("Klub tidak ditemukan.")
	}
}


func buatJadwal(liga Liga, jadwal *DaftarJadwal, jumlahKlub int) int {
	var klubAktif [NMAX]string
	var k, i, n, idx, totalPekan,pekan int
	k = 0
	for i = 0; i < NMAX; i++ {
		if liga[i].Nama != "" {
			klubAktif[k] = liga[i].Nama
			k++
		}
	}

	n = jumlahKlub
	if n%2 != 0 {
		klubAktif[n] = "Input Salah"
		n++
	}

	idx = 0
	totalPekan = (n - 1) * 2

	for pekan = 0; pekan < totalPekan/2; pekan++ {
		for i = 0; i < n/2; i++ {
			if klubAktif[i] != "Input Salah" && klubAktif[n-1-i] != "Input Salah" {
				jadwal[idx] = Jadwal{Klub1: klubAktif[i], Klub2: klubAktif[n-1-i], Pekan: pekan + 1}
				idx++
			}
		}
		last := klubAktif[n-1]
		for i = n - 1; i > 1; i-- {
			klubAktif[i] = klubAktif[i-1]
		}
		klubAktif[1] = last
	}

	for pekan = 0; pekan < totalPekan/2; pekan++ {
		for i = 0; i < n/2; i++ {
			if klubAktif[i] != "Input Salah" && klubAktif[n-1-i] != "Input Salah" {
				jadwal[idx] = Jadwal{Klub1: klubAktif[n-1-i], Klub2: klubAktif[i], Pekan: pekan + 1 + totalPekan/2}
				idx++
			}
		}
		last := klubAktif[n-1]
		for i = n - 1; i > 1; i-- {
			klubAktif[i] = klubAktif[i-1]
		}
		klubAktif[1] = last
	}
	return idx
}

func tampilkanJadwal(jadwal DaftarJadwal, jumlah int) {
	var i int
	for i = 0; i < jumlah; i++ {
		fmt.Printf("Pekan %d: %s vs %s\n", jadwal[i].Pekan, jadwal[i].Klub1, jadwal[i].Klub2)
	}
}

func inputHasil(liga *Liga) bool {
	var klub1, klub2 string
	var skor1, skor2 int
	var i, j int
	
	fmt.Print("Masukkan nama klub 1: ")
	fmt.Scan(&klub1)
	fmt.Print("Masukkan nama klub 2: ")
	fmt.Scan(&klub2)
	fmt.Print("Masukkan skor klub 1: ")
	fmt.Scan(&skor1)
	fmt.Print("Masukkan skor klub 2: ")
	fmt.Scan(&skor2)

	i = cariKlub(*liga, klub1)
	j = cariKlub(*liga, klub2)

	if i == -1 || j == -1 {
		fmt.Println("Salah satu klub tidak ditemukan.")
		return false
	}

	liga[i].Main++
	liga[j].Main++
	liga[i].GolMasuk += skor1
	liga[i].GolKemasukan += skor2
	liga[j].GolMasuk += skor2
	liga[j].GolKemasukan += skor1

	if skor1 > skor2 {
		liga[i].Menang++
		liga[j].Kalah++
		liga[i].Poin += 3
	} else if skor2 > skor1 {
		liga[j].Menang++
		liga[i].Kalah++
		liga[j].Poin += 3
	} else {
		liga[i].Seri++
		liga[j].Seri++
		liga[i].Poin++
		liga[j].Poin++
	}

	liga[i].SelisihGol = liga[i].GolMasuk - liga[i].GolKemasukan
	liga[j].SelisihGol = liga[j].GolMasuk - liga[j].GolKemasukan
	return true
}

func tampilkanPeringkat(liga Liga) {
	dikunjungi := [NMAX]bool{}
	total := hitungKlub(liga)

	for urutan := 1; urutan <= total; urutan++ {
		maxIdx := -1
		for i := 0; i < NMAX; i++ {
			if liga[i].Nama != "" && !dikunjungi[i] {
				if maxIdx == -1 || liga[i].Poin > liga[maxIdx].Poin ||
					(liga[i].Poin == liga[maxIdx].Poin && liga[i].SelisihGol > liga[maxIdx].SelisihGol) {
					maxIdx = i
				}
			}
		}

		if maxIdx != -1 {
			fmt.Printf("%2d. %-5s | Poin: %3d | SG: %3d\n", urutan, liga[maxIdx].Nama, liga[maxIdx].Poin, liga[maxIdx].SelisihGol)
			dikunjungi[maxIdx] = true
		}
	}
}

func sortLigaByPoin(liga *Liga, total int) {
	for i := 0; i < total-1; i++ {
		for j := i + 1; j < total; j++ {
			if liga[i].Poin < liga[j].Poin ||
				(liga[i].Poin == liga[j].Poin && liga[i].SelisihGol < liga[j].SelisihGol) {
				liga[i] = liga[j]
				liga[j] = liga[i]
			}
		}
	}
}

func binarySearchByPoin(liga Liga, total, carPoin int) int {
	low := 0
	high := 0

	for low <= high {
		mid := (low + high) / 2
		if liga[mid].Poin == carPoin {
			return mid
		} else if liga[mid].Poin < carPoin {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func cariByPoin(liga Liga) {
	var carPoin int
    total := hitungKlub(liga)
    
    if total > 0 {
        fmt.Print("Masukkan jumlah poin yang ingin dicari: ")
        fmt.Scan(&carPoin)

        for i := 0; i < total-1; i++ {
            for j := i + 1; j < total; j++ {
                if liga[i].Poin < liga[j].Poin || 
                   (liga[i].Poin == liga[j].Poin && liga[i].SelisihGol < liga[j].SelisihGol) {
                    temp := liga[i]
                    liga[i] = liga[j]
                    liga[j] = temp
                }
            }
        }
        found := false
        for i := 0; i < total; i++ {
            if liga[i].Poin == carPoin {
                if !found {
                    fmt.Printf("\nKlub dengan %d poin:\n", carPoin)
                    found = true
                }
                fmt.Printf("- %s | Main: %d | Menang: %d | Seri: %d | Kalah: %d | GM: %d | GK: %d | SG: %d\n",
                    liga[i].Nama, liga[i].Main, liga[i].Menang, liga[i].Seri, liga[i].Kalah,
                    liga[i].GolMasuk, liga[i].GolKemasukan, liga[i].SelisihGol)
            }
        }
        if !found {
            fmt.Println("Tidak ada klub dengan poin tersebut.")
        }
    } else {
        fmt.Println("Belum ada klub yang terdaftar.")
    }
}



func main() {
	var liga Liga
	var jadwal DaftarJadwal
	var jumlahJadwal int

	for {
		fmt.Println("\n--- MENU EPL MANAGER ---")
		fmt.Println("1. Tambah Klub")
		fmt.Println("2. Ubah Klub")
		fmt.Println("3. Hapus Klub")
		fmt.Println("4. Buat Jadwal Pertandingan")
		fmt.Println("5. Lihat Jadwal")
		fmt.Println("6. Input Hasil Pertandingan")
		fmt.Println("7. Lihat Peringkat")
		fmt.Println("8. Cari Berdasarkan Poin")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var menu int
		fmt.Scan(&menu)

		switch menu {
		case 1:
			tambahKlub(&liga)
		case 2:
			ubahKlub(&liga)
		case 3:
			hapusKlub(&liga)
		case 4:
			jumlahKlub := hitungKlub(liga)
			jumlahJadwal = buatJadwal(liga, &jadwal, jumlahKlub)
			fmt.Println("Jadwal berhasil dibuat.")
		case 5:
			tampilkanJadwal(jadwal, jumlahJadwal)
		case 6:
			inputHasil(&liga)
		case 7:
			tampilkanPeringkat(liga)
		case 8:
			cariByPoin(liga)
		case 0:
			fmt.Println()
			return
		default:
			fmt.Println("Menu tidak valid")
		}
	}
}
