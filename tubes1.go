package main

import "fmt"

const NMAX int = 100

type IdeStartup struct {
	namaIde   string
	kategori  string
	deskripsi string
	skor      int
}

// Variabel global HANYA untuk array utama (Sesuai Aturan)
var arrIde [NMAX]IdeStartup

func main() {
	// nIde dipindah ke lokal main (Sesuai Aturan Variabel Global)
	var nIde int = 0
	var pilihan int
	var aplikasiJalan bool = true

	// Penggunaan variabel boolean pengganti break (Sesuai Aturan)
	for aplikasiJalan == true {
		fmt.Println("\n=== APLIKASI IDE STARTUP ===")
		fmt.Println("1. Tambah Ide Baru")
		fmt.Println("2. Tampilkan Semua Ide")
		fmt.Println("3. Edit Ide")
		fmt.Println("4. Hapus Ide")
		fmt.Println("5. Urutkan Data Ide")
		fmt.Println("6. Cari Ide by Skor (Binary Search - Ascending)")
		fmt.Println("7. Cari Ide by Kategori (Sequential Search)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahIde(&arrIde, &nIde)
		} else if pilihan == 2 {
			tampilIde(&arrIde, nIde)
		} else if pilihan == 3 {
			editIde(&arrIde, nIde)
		} else if pilihan == 4 {
			hapusIde(&arrIde, &nIde)
		} else if pilihan == 5 {
			var jenisSort, arah int
			fmt.Println("\nPilih Algoritma Pengurutan:")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&jenisSort)

			fmt.Println("\nPilih Arah Pengurutan:")
			fmt.Println("1. Ascending (Kecil ke Besar)")
			fmt.Println("2. Descending (Besar ke Kecil)")
			fmt.Print("Pilih (1/2): ")
			fmt.Scan(&arah)

			if jenisSort == 1 && arah == 1 {
				selectionSortSkorAsc(&arrIde, nIde)
				fmt.Println("Data berhasil diurutkan (Selection Sort - Ascending).")
			} else if jenisSort == 1 && arah == 2 {
				selectionSortSkorDesc(&arrIde, nIde)
				fmt.Println("Data berhasil diurutkan (Selection Sort - Descending).")
			} else if jenisSort == 2 && arah == 1 {
				insertionSortSkorAsc(&arrIde, nIde)
				fmt.Println("Data berhasil diurutkan (Insertion Sort - Ascending).")
			} else if jenisSort == 2 && arah == 2 {
				insertionSortSkorDesc(&arrIde, nIde)
				fmt.Println("Data berhasil diurutkan (Insertion Sort - Descending).")
			} else {
				fmt.Println("Pilihan algoritma atau arah tidak valid.")
			}
		} else if pilihan == 6 {
			// Membuat salinan (copy) array agar data asli tidak teracak
			var arrCopy [NMAX]IdeStartup = arrIde
			var nCopy int = nIde

			// Urutkan hanya array salinannya saja
			insertionSortSkorAsc(&arrCopy, nCopy)

			var cariSkor int
			fmt.Print("Masukkan Skor yang dicari: ")
			fmt.Scan(&cariSkor)

			// Lakukan Binary Search pada array salinan yang sudah terurut
			var hasil int = binSearchSkorAsc(&arrCopy, nCopy, cariSkor)
			if hasil != -1 {
				// Tampilkan hasil dari array salinan
				fmt.Printf("Ide ditemukan! Nama: %s | Kategori: %s | Deskripsi: %s\n", arrCopy[hasil].namaIde, arrCopy[hasil].kategori, arrCopy[hasil].deskripsi)
			} else {
				fmt.Println("Ide dengan skor tersebut tidak ditemukan.")
			}
		} else if pilihan == 7 {
			cariIdeKategori(&arrIde, nIde)
		} else if pilihan == 0 {
			aplikasiJalan = false
			fmt.Println("Program selesai dijalankan.")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Spesifikasi: Menambahkan data ide baru ke dalam array
// I.S.: Array T terdefinisi, n berisi jumlah data saat ini
// F.S.: Data ide baru ditambahkan ke array T, n bertambah 1 jika kapasitas belum penuh
func tambahIde(T *[NMAX]IdeStartup, n *int) {
	if *n < NMAX {
		fmt.Print("Masukkan Nama Ide (1 kata): ")
		fmt.Scan(&T[*n].namaIde)
		fmt.Print("Masukkan Kategori (1 kata): ")
		fmt.Scan(&T[*n].kategori)
		fmt.Print("Masukkan Deskripsi Ide (1 kata): ")
		fmt.Scan(&T[*n].deskripsi)
		fmt.Print("Masukkan Skor (1-10): ")
		fmt.Scan(&T[*n].skor)

		*n = *n + 1
		fmt.Println("Ide berhasil ditambahkan!")
	} else {
		fmt.Println("Penyimpanan penuh!")
	}
}

// Spesifikasi: Menampilkan seluruh data ide startup yang tersimpan
// I.S.: Array T terdefinisi, n berisi jumlah data saat ini
// F.S.: Seluruh elemen array T dari indeks 0 hingga n-1 ditampilkan ke layar
func tampilIde(T *[NMAX]IdeStartup, n int) {
	var i int = 0
	if n == 0 {
		fmt.Println("Belum ada data ide startup.")
	} else {
		fmt.Println("----------------------------------------------------------------------------------------------------")
		for i < n {
			fmt.Printf("%d. Nama: %s | Kategori: %s | Deskripsi: %s | Skor: %d\n", i+1, T[i].namaIde, T[i].kategori, T[i].deskripsi, T[i].skor)
			i = i + 1
		}
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}

// Spesifikasi: Mencari indeks data berdasarkan Nama Ide menggunakan Sequential Search
// I.S.: Array T terdefinisi berisi n data, target string yang dicari diketahui
// F.S.: Mengembalikan indeks array jika ditemukan, atau -1 jika tidak ditemukan
func seqSearchNama(T *[NMAX]IdeStartup, n int, target string) int {
	var ketemu int = -1
	var k int = 0
	for ketemu == -1 && k < n {
		if T[k].namaIde == target {
			ketemu = k
		}
		k = k + 1
	}
	return ketemu
}

// Spesifikasi: Mengubah data kategori, deskripsi, dan skor pada ide tertentu
// I.S.: Array T terdefinisi berisi n data
// F.S.: Jika ide ditemukan, datanya diubah sesuai inputan baru
func editIde(T *[NMAX]IdeStartup, n int) {
	var target string
	var idx int
	fmt.Print("Nama Ide yang ingin diedit: ")
	fmt.Scan(&target)

	idx = seqSearchNama(T, n, target)
	if idx != -1 {
		fmt.Print("Masukkan Kategori Baru: ")
		fmt.Scan(&T[idx].kategori)
		fmt.Print("Masukkan Deskripsi Baru: ")
		fmt.Scan(&T[idx].deskripsi)
		fmt.Print("Masukkan Skor Baru: ")
		fmt.Scan(&T[idx].skor)
		fmt.Println("Data berhasil diubah!")
	} else {
		fmt.Println("Ide tidak ditemukan.")
	}
}

// Spesifikasi: Menghapus data ide tertentu dengan menggeser elemen array
// I.S.: Array T terdefinisi berisi n data
// F.S.: Jika ide ditemukan, data dihapus, elemen sebelah kanannya digeser ke kiri, n berkurang 1
func hapusIde(T *[NMAX]IdeStartup, n *int) {
	var target string
	var idx, i int
	fmt.Print("Nama Ide yang ingin dihapus: ")
	fmt.Scan(&target)

	idx = seqSearchNama(T, *n, target)
	if idx != -1 {
		i = idx
		for i <= *n-2 {
			T[i] = T[i+1]
			i = i + 1
		}
		*n = *n - 1
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("Ide tidak ditemukan.")
	}
}

// Spesifikasi: Mengurutkan array berdasarkan skor dari besar ke kecil (Descending)
// I.S.: Array A terdefinisi berisi N data acak
// F.S.: Array A terurut secara Descending menggunakan algoritma Selection Sort
func selectionSortSkorDesc(A *[NMAX]IdeStartup, N int) {
	var pass, idx, i int
	var temp IdeStartup

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if A[idx].skor < A[i].skor {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

// Spesifikasi: Mengurutkan array berdasarkan skor dari kecil ke besar (Ascending)
// I.S.: Array A terdefinisi berisi N data acak
// F.S.: Array A terurut secara Ascending menggunakan algoritma Selection Sort
func selectionSortSkorAsc(A *[NMAX]IdeStartup, N int) {
	var pass, idx, i int
	var temp IdeStartup

	pass = 1
	for pass <= N-1 {
		idx = pass - 1
		i = pass
		for i < N {
			if A[idx].skor > A[i].skor {
				idx = i
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

// Spesifikasi: Mengurutkan array berdasarkan skor dari kecil ke besar (Ascending)
// I.S.: Array A terdefinisi berisi N data acak
// F.S.: Array A terurut secara Ascending menggunakan algoritma Insertion Sort
func insertionSortSkorAsc(A *[NMAX]IdeStartup, N int) {
	var pass, i int
	var temp IdeStartup

	pass = 1
	for pass < N {
		temp = A[pass]
		i = pass - 1
		// Tidak menggunakan break, diganti dengan kondisi and di loop
		for i >= 0 && A[i].skor > temp.skor {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = temp
		pass = pass + 1
	}
}

// Spesifikasi: Mengurutkan array berdasarkan skor dari besar ke kecil (Descending)
// I.S.: Array A terdefinisi berisi N data acak
// F.S.: Array A terurut secara Descending menggunakan algoritma Insertion Sort
func insertionSortSkorDesc(A *[NMAX]IdeStartup, N int) {
	var pass, i int
	var temp IdeStartup

	pass = 1
	for pass < N {
		temp = A[pass]
		i = pass - 1
		for i >= 0 && A[i].skor < temp.skor {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = temp
		pass = pass + 1
	}
}

// Spesifikasi: Mencari data berdasarkan skor secara spesifik menggunakan Binary Search
// I.S.: Array T terdefinisi dan WAJIB terurut secara Ascending
// F.S.: Mengembalikan indeks array jika ditemukan, atau -1 jika tidak ditemukan
func binSearchSkorAsc(T *[NMAX]IdeStartup, N int, targetSkor int) int {
	var left int = 0
	var right int = N - 1
	var found int = -1
	var mid int

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if targetSkor < T[mid].skor {
			right = mid - 1
		} else if targetSkor > T[mid].skor {
			left = mid + 1
		} else {
			found = mid
		}
	}
	return found
}

// Spesifikasi: Mencari dan menampilkan seluruh ide berdasarkan kategori tertentu
// I.S.: Array T terdefinisi berisi n data, target string kategori diketahui
// F.S.: Menampilkan semua elemen yang memiliki kategori sama dengan target (Sequential Search)
func cariIdeKategori(T *[NMAX]IdeStartup, n int) {
	var target string
	var i int = 0
	var ada bool = false

	fmt.Print("Masukkan Kategori yang ingin dicari (1 kata): ")
	fmt.Scan(&target)

	fmt.Println("----------------------------------------------------------------------------------------------------")
	for i < n {
		if T[i].kategori == target {
			fmt.Printf("Nama: %s | Deskripsi: %s | Skor: %d\n", T[i].namaIde, T[i].deskripsi, T[i].skor)
			ada = true
		}
		i = i + 1
	}

	if ada == false {
		fmt.Println("Ide dengan kategori tersebut tidak ditemukan.")
	}
	fmt.Println("----------------------------------------------------------------------------------------------------")
}