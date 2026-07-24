# 🚀 Startup Idea Manager (CLI)

Aplikasi command-line interface (CLI) sederhana berbasis Go untuk mengelola data ide startup — mendukung operasi CRUD penuh (Create, Read, Update, Delete) serta dilengkapi implementasi algoritma sorting dan searching sebagai mesin pengurutan & pencarian data.

> Tugas Besar — Mata Kuliah Algoritma Pemrograman

## ✨ Fitur

| Menu | Fungsi |
|---|---|
| 1. Tambah Ide Baru | Menambahkan data ide startup baru ke dalam array |
| 2. Tampilkan Semua Ide | Menampilkan seluruh data ide yang tersimpan |
| 3. Edit Ide | Mengubah data ide berdasarkan nama (Sequential Search) |
| 4. Hapus Ide | Menghapus data ide dan menggeser elemen array |
| 5. Urutkan Data Ide | Sorting berdasarkan skor (Selection Sort / Insertion Sort, Ascending/Descending) |
| 6. Cari Ide by Skor | Binary Search pada array yang sudah terurut |
| 7. Cari Ide by Kategori | Sequential Search berdasarkan kategori |
| 0. Keluar | Keluar dari aplikasi |

## 🧠 Algoritma yang Diimplementasikan

- **Selection Sort** (Ascending & Descending) — mengurutkan data berdasarkan skor
- **Insertion Sort** (Ascending & Descending) — alternatif algoritma pengurutan
- **Binary Search** — pencarian skor pada data yang sudah terurut (kompleksitas O(log n))
- **Sequential Search** — pencarian berdasarkan nama ide (untuk edit/hapus) dan kategori

## 🗂️ Struktur Data

```go
type IdeStartup struct {
    namaIde   string
    kategori  string
    deskripsi string
    skor      int
}
```

Data disimpan dalam array statis berukuran tetap (`NMAX = 100`), dengan variabel penghitung jumlah data (`n`) dikelola secara lokal (bukan variabel global) sesuai kaidah pemrograman terstruktur.

## ▶️ Cara Menjalankan

Pastikan [Go](https://go.dev/dl/) sudah terinstal, lalu jalankan:

```bash
git clone https://github.com/akmal234dcn/startup-idea-manager-cli-go.git
cd startup-idea-manager-cli-go
go run main.go
```

Atau build menjadi executable:

```bash
go build -o startup-idea-manager
./startup-idea-manager
```

## 🛠️ Tech Stack

- **Bahasa:** Go (Golang)
- **Paradigma:** Pemrograman terstruktur (array, struct, pointer, procedural)

## 📚 Catatan Akademik

Project ini dikerjakan sebagai bagian dari tugas besar mata kuliah **Algoritma Pemrograman**, dengan fokus penerapan struktur data array, struct, serta algoritma sorting & searching dasar secara manual (tanpa memakai fungsi bawaan/library sorting Go).
