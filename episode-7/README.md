# Belajar Golang Dasar #7: Control Flow (if, switch, loop)

Control flow adalah struktur yang mengatur urutan eksekusi kode dalam program. Dengan control flow, kita bisa membuat program mengambil keputusan (menggunakan kondisi) dan melakukan perulangan (looping). Di Golang, struktur control flow yang umum digunakan adalah: `if`, `swtich`, dan `for`.

## 1. Struktur Kondisional: `if`, `else if`, `else`

Struktur `if` digunakan untuk menjalankan blok kode hanya jika suatu kondisi bernilai `true`.

```go
nilai := 85

if nilai >= 90 {
    fmt.Println("Nilai A")
} else if nilai >= 80 {
    fmt.Println("Nilai B")
} else {
    fmt.Println("Nilai C atau di bawahnya")
}
```

> [!NOTE] > **Catatan**: Tidak perlu tanda kurung di sekitar kondisi (`if nilai >= 90` langsung saja), tapi wajib pakai `{}` untuk menulis blok kode.

## 2. Struktur Pemilihan: `switch`

`switch` digunakan untuk mengecek satu nilai terhadap beberapa kemungkinan secara lebih rapi dibanding banyak `if-else`.

```go
hari := "Senin"

switch hari {
case "Senin":
    fmt.Println("Weekday")
case "Sabtu", "Minggu":
    fmt.Println("Weekend")
default:
    fmt.Println("Hari biasa")
}
```

Beberapa catatan:

-   Tidak perlu `break`, karena Golang otomatis keluar setelah satu `case` terpenuhi.
-   Bisa menuliskan beberapa nilai dalam satu `case` (seperti `"Sabtu", "Minggu"`).

## 3. Struktur Perulangan: `for`

Golang hanya punya satu kata kunci untuk perulangan, yaitu `for`. Tapi fleksibel banget, bisa dipakai sebagai `while`, `do-while`, atau `for-each`.

### 1. Standar (mirip `for` di C/C++)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 2. Tanpa Kondisi (infinite loop)

```go
for {
    fmt.Println("Terus jalan...")
    break // keluar dari loop
}
```

### 3. Dengan `range`

Digunakan untuk melakukan iterasi pada array, slice, map, atau string.

```go
angka := []int{1, 2, 3, 4}

for index, value := range angka {
    fmt.Println(index, value)
}
```

Kalau kita cuma butuh nilainya saja:

```go
for _, value := range angka {
    fmt.Println(value)
}
```

## Kesimpulan

Control flow di golang membantu kita mengatur logika program dengan baik. Dengan `if`, `switch`, dan `for`, kita bisa membuat program yang cerdas dalam mengambil keputusan dan melakukan perulangan.

Memahami struktur-struktur ini sangat penting sebelum masuk ke konsep yang lebih kompleks seperti function, struct, atau goroutine.

Selanjutnya kita akan masuk ke pembahasan tentang fungsi (function), blok kode yang bisa digunakan berulang kali di Golang.
