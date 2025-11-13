# Belajar Golang Dasar #4: Variable

## Apa itu Variable?

Variable adalah tempat untuk menyimpan sebuah nilai. Nilai ini bisa berupa angka, teks, boolean, atau tipe data lainnya.

Nama variable sebaiknya deskriptif agar mudah dipahami, dan mengikuti aturan penulisan yang berlaku (tidak diawali angka, tidak mengandung spasi, dan tidak menggunakan simbol khusus).

## Cara Mendeklarasikan Variable di Golang

Ada beberapa cara untuk mendeklarasikan variable di Golang:

### 1. Dengan Keyword `var`

```go
var nama string = "Arman"
var umur int = 22
```

Kita bisa mendeklarasikan variable tanpa langsung mengisi value atau nilainya

```go
var status bool
```

Variable `status` di atas akan otomatis memiliki nilai awal (zero value), yaitu `false` karena bertipe boolean.

### 2. Dengan Type Inference

Golang bisa menebak tipe data dari nilai yang diberikan, jadi kita bisa menulis seperti ini:

```go
var kota = "Jakarta" // otomatis bertipe string
```

### 3. Dengan Short Declaration

Deklarasi singkat menggunakan `:=` hanya bisa digunakan di dalam fungsi, dan menjadi cara yang paling umum dipakai:

```go
judul := "Belajar Golang"
nilai := 90
```

## Multiple Declaration

Kita juga bisa mendeklarasikan beberapa variable sekaligus:

```go
var a, b, c int = 1, 2, 3
```

Atau tanpa menyebut tipe data secara eksplisit:

```go
x, y, z := "A", "B", "C"
```

## Mengganti Nilai Variable

Nilai variable bisa diubah selama program berjalan:

```go
var jumlah int = 10
jumlah = 15
```

Jika kita menggunakan `:=` untuk variable yang sudah pernah dideklarasikan sebelumnya, akan muncul error. Maka gunakan `=` untuk mengubah nilai, bukan `:=`.

## Variable di Luar Fungsi

Kita juga bisa mendekalarasikan variable di luar fungsi `main`, tapi hanya dengan `var`, bukan `:=`:

```go
var versi = "1.0.0"

func main() {
    fmt.Println("Versi: ", versi)
}
```

## Kesimpulan

Variable di Golang dapat dideklarasikan dengan beberapa cara, seperti menggunakan `var`, deklarasi singkat `:=`, maupun deklarasi sekaligus beberapa varible.

Memahami penggunaan variable sangat penting karena hampir seluruh program akan bergantung pada penyimpanan dan pengolahan nilai melalui variable.

Pada artikel selanjutnya, kita semua akan belajar bersama-sama tentang Konstanta di dalam bahasa Golang.
