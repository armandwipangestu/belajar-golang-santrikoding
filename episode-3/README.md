# Belajar Golang Dasar #3: Tipe Data Dasar

## 1. Tipe Data Angka (Number)

### Integer

Tipe data integer digunakan untuk menyimpan bilangan bulat, baik positif maupun negatif.

```go
var usia int = 22
```

Golang menyediakan beberapa variasi integer: `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.

> [!NOTE]
> **Catatan**: `int` biasanya cukup untuk kebutuhan umum, karena ukurannya menyesuaikan arsitektur sistem (32-bit atau 64-bit).
>
> Perbedaan antara `int` dan `uint` (Unsigned Int) adalah:
>
> -   jika `int` bisa menyimpan bilangan negatif dan positif, contohnya `-10`, `0`, `42`, dimana range nya tergantung arsitektur nya (misalnya di 64-bit: `-9223372036854775808` sampai `9223372036854775808`).
>
> -   jika `uint` hanya bisa menyimpan bilangan positif (termasuk nol), contohnya `0`, `42`, `99999`, range nya dimulai dari `0` sampai `18446744073709551615` pada 64-bit.

### Float

Untuk bilangan desimal, kita bisa menggunakan tipe `float32` atau `float64`.

```go
var tinggi float32 = 175.5
```

## 2. Tipe Data String

Tipe `string` digunakan untuk menyimpan teks.

```go
var nama string = "Golang"
```

Kita juga bisa menggunakan tanda backtick (`) untuk string multi-baris:

```go
pesan := `Halo,
Ini adalah teks
dengan banyak baris.`
```

## 3. Tipe Data Boolean

Tipe boolean hanya memiliki dua nilai yaitu `true` atau `false`.

```go
var aktif bool = true
```

Boolean sangat sering digunakan dalam kondisi, seperti pada penyataan `if`.

## 4. Tipe Data Default (Deklarasi tanpa Tipe)

Golang mendukung deklarasi variabel tanpa menyebutkan tipe secara eksplisit. Tipe data akan ditentukan berdasarkan nilai yang diberikan.

```go
nama := "Belajar Golang"
umur := 22
lulus := false
```

Penulisan ini dikenal dengan nama `short declaration`.

## 5. Zero Value

Jika sebuah variable dideklarasikan tanpa diberi nilai awal, maka Golang akan memberikan nilai default (zero value) tergantung pada tipenya:

| Tipe Data | Zero Value  |
| --------- | ----------- |
| int       | 0           |
| float     | 0.0         |
| bool      | false       |
| string    | "" (kosong) |

Contoh:

```go
var angka int
fmt.Println(angka) // Output: 0
```

## 6. Konversi Tipe Data

Kita bisa mengubah (konversi) tipe data secara eksplisit menggunakan cara seperti ini:

```go
var x int = 10
var y float64 = float64(x)
```

Namun, Golang tidak mendukung konversi otomatis seperti bahasa lain. Kita harus selalu menyebutkan secara eksplisit tipe yang ingin digunakan.

## Kesimpulan

Tipe data dasar di Golang meliputi angka (integer dan float), string, boolean, serta aturan terkait zero value dan konversi tipe.
