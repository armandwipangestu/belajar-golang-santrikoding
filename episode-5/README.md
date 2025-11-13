# Belajar Golang Dasar #5: Konstanta

## Apa itu Konstanta?

Konstanta (constant) adalah variable khusus yang nilainya tidak dapat diubah setelah dideklarasikan. Biasanya digunakan untuk nilai-nilai yang sudah pasti, seperti konfigurasi, batasan, atau data tetap lainnya. Contoh penggunaan:

1. Nilai pi (`3.14`)
2. Jumlah hari dalam seminggu (`7`)
3. Pesan default, versi aplikasi, dll

## Cara Mendeklarasikan Konstanta

Konstanta dideklarasikan dengan kata kunci `const`.

```go
const pi = 3.14
const aplikasi = "Belajar Golang"
```

Tidak seperti variable, kita tidak bisa menggunakan `:=` untuk mendeklarasikan konstanta.

## Konstanta Bertipe dan Tanpa Tipe

Golang memperbolehkan dua jenis konstanta:

### 1. Konstanta Tanpa Tipe

```go
const angka = 10 // Tipe akan disesuaikan saat digunakan
```

### 2. Konstanta Dengan Tipe Tertentu

```go
const umur int = 25
```

Konstanta tanpa tipe lebih fleksibel karena tipe-nya akan disesuaikan dengan konteks saat dipakai.

## Grup Konstanta

Kita bisa mendeklarasikan beberapa konstanta sekaligus dalam satu blok:

```go
const (
    satu = 1
    dua = 2
    tiga = 3
)
```

## Keyword `iota` (Untuk Nilai Berurutan)

Golang menyediakan keyword `iota` untuk membuat urutan konstanta secara otomatis. Biasanya digunakan dalam enum atau deklarasi bertingkat.

```go
const (
    A = iota    // 0
    B           // 1
    C           // 2
)
```

Setiap kali `iota` muncul, nilainya akan meningkat 1 secara otomatis.

## Perbedaan Variable dan Konstanta

| Aspek        | Variable        | Konstanta   |
| ------------ | --------------- | ----------- |
| Dapat diubah | Ya              | Tidak       |
| Keyword      | `var` atau `:=` | `const`     |
| Digunakan    | Data dinamis    | Nilai tetap |

## Kesimpulan

Konstanta di Golang digunakan untuk menyimpan nilai tetap yang tidak berubah selama program berjalan. Kita menggunakan `const` untuk mendeklarasikannya, dan bisa memanfaatkan `iota` untuk membuat nilai berurutan secara otomatis.
