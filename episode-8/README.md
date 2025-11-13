# Belajar Golang Dasar #8: Function

Fungsi (function) adalah bagian penting dalam pemrograman Golang. Fungsi membantu kita memisahkan kode ke dalam blok-blok kecil yang bisa digunakan kembali, membuat program jadi lebih rapi, mudah dibaca, dan mudah dikelola.

Di artikel ini, kita akan belajar dasar penggunaan fungsi di Golang, mulai dari cara mendeklarasikan, memanggil, hingga parameter dan nilai kembaliannya.

## Cara Membuat Fungsi di Golang

Fungsi dideklarasikan dengan kata kunci `func`, diikuti dengan nama fungsi, daftar parameter, dan tipe nilai yang dikembalikan (jika ada).

Contoh fungsi sederhana tanpa parameter dan tanpa return.

```go
func sapa() {
    fmt.Println("Halo dari fungsi!")
}
```

Untuk menjalankannya:

```go
sapa()
```

## Fungsi dengan Paramter

Kita bisa mengirimkan nilai ke dalam fungsi menggunakan parameter:

```go
func halo(nama string) {
    fmt.Println("Halo", nama)
}
```

Pemanggilan:

```go
halo("Golang")
```

Fungsi juga bisa punya beberapa parameter:

```go
func tambah(a int, b int) {
    fmt.Println("Hasil:", a + b)
}
```

Atau disingkat

```go
func tambah(a, b int) {
    fmt.Println("Hasil:", a + b)
}
```

## Fungsi dengan Return Value

Fungsi bisa mengembalikan nilai menggunakan `return`.

```go
func kali (a, b int) int {
    return a * b
}
```

Contoh penggunaan:

```go
hasil := kali(3, 4)
fmt.Println(hasil) // Output: 12
```

## Function dengan Multiple Return Value

Golang mendukung fungsi yang mengembalikan lebih dari satu nilai.

```go
func hitung(a, b int) (int, int) {
    return a + b, a * b
}
```

Pemanggilan:

```go
jumlah, perkalian := hitung(2, 5)
fmt.Println("Jumlah:", jumlah)
fmt.Println("Perkalian:", perkalian)
```

## Fungsi dengan Return Bernama

Kita juga bisa memberi nama nilai yang dikembalikan, lalu cukup gunakan `return` tanpa argument

```go
func bagi(a, b int) (hasil int) {
    hasil = a / b
    return
}
```

## Fungsi di dalam Variable (Function as Value)

Fungsi bisa disimpan dalam variable:

```go
f := func(nama string) {
    fmt.Println("Hai", nama)
}

f("Golang")
```

## Kesimpulan

Fungsi (function) di Golang membantu menulis kode yang lebih terstruktur, efisien, dan mudah dikelola. Kita bisa membuat fungsi sederhana hingga fungsi dengan parameter dan banyak return.

Pada artikel selanjutnya kita akan bahas tentang Pointer yaitu cara bekerja dengan referensi memori di Golang.
