# Belajar Golang Dasar #10: Array, Slice, dan Map

Ketika kita membuat program, sering kali kita nggak cuma bekerja dengan satu atau dua data saja. Misalnya, kita ingin menyimpan daftar nama mahasiswa, kumpulan nilai, atau data produk.

Nah, menyimpan data satu per satu dalam variable terpisah jelas merepotkan dan nggak efisien. Di sinilah struktur data seperti array, slice, dan map jadi sangat berguna. Di Golang, kita punya tiga jenis utama untuk menyimpan kumpulan data:

-   Array, untuk menyimpan sejumlah data dengan ukuran tetap
-   Slice, versi fleksibelnya array
-   Map, untuk menyimpan pasangan key-value

Ketiganya punya peran masing-masing, dan pemahaman dasar tentang cara kerja mereka akan bantu banget saat membangun program yang lebih kompleks nantinya.

## 1. Array

Array adalah kumpulan elemen dengan jumlah tetap dan tipe data yang sama.

```go
var angka [3]int

angka[0] = 10
angka[1] = 20
angka[3] = 30

// Atau

angka := [3]int{10, 20, 30}
```

Kita bisa akses elemennya pakai indeks:

```go
fmt.Println(angka[1]) // Output: 20
```

Catatan

-   Ukuran array tetap dan tidak bisa diubah
-   Indeks dimulai dari 0

## 2. Slice

Slice mirip array, tapi lebih fleksibel karena ukurannya bisa berubah. Slice sebenarnya adalah wrapper di atas array.

```go
angka := []int{10, 20, 30}
fmt.Println(angka[0]) // Output: 10
```

Kita juga bisa membuat slice dari array:

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Ambil elemen ke-1 sampai sebelum ke-4
fmt.Println(slice) // [2 3 4]
```

## 3. Map

Map adalah struktur data pasangan key-value. Mirip seperti dictionary di Python atau object di JavaScript.

```go
nilai := map[string]int{
    "Andi": 90,
    "Budi": 85,
}

fmt.Println(nilai["Andi"]) // Output: 90
```

Menambah data:

```go
nilai["Citra"] = 88
```

Menghapus data:

```go
delete(nilai, "Budi")
```

Cek apakah key ada:

```go
val, ok := nilai["Dina"]

if ok {
    fmt.Println("Ada:", val)
} else {
    fmt.Println("Tidak ditemukan")
}
```

## Perbandingan Singkat

| Struktur | Ukuran  | Bisa Diubah? | Cocok Untuk                  |
| -------- | ------- | ------------ | ---------------------------- |
| Array    | Tetap   | Tidak        | Data tetap, kecil            |
| Slice    | Dinamis | Ya           | Koleksi data yang fleksibel  |
| Map      | Dinamis | Ya           | Data berpasangan (key-value) |

## Kesimpulan

Array, Slice, dan Map adalah tiga cara utama untuk menyimpan data dalam Golang. Array cocok untuk jumlah data tetap, Slice untuk koleksi data yang fleksibel, dan Map untuk data pasangan key-value.

Pada artikel selanjutnya, kita semua akan belajar tentang Struct di dalam bahasa pemrograman golang.
