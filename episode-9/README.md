# Belajar Golang Dasar #9: Pointer

Dalam pemrograman, terutama bahasa seperti Golang, pointer adalah konsep penting yang perlu kita pahami. Pointer memungkinkan kita untuk bekerja langsung dengan alamat memori suatu variable. Ini berguna untuk efisiensi memori, manipulasi data, dan juga dasar dari konsep pass by reference.

## Apa itu Pointer?

Pointer adalah variable yang menyimpan alamat memori dari variable lain. Kalau biasanya kita menyimpan dan mengakses nilai, pointer menyimpan dan mengakses alamat dari nilai tersebut.

## Deklarasi dan Penggunaan Pointer

### 1. Mengambil Alamat Variable (`&`)

```go
var angka int = 10
var ptr *int = &angka

fmt.Println("Nilai angka:", angka)
fmt.Println("Alamat angka:", &angka)
fmt.Println("Pointer ptr:", ptr)
```

Penjelasan

-   `&angka` - mengambil alamat dari variable `angka`
-   `*int` - menyatakan bahwa `ptr` adalah pointer ke `int`

### 2. Mengakses Nilai dari Pointer (`*`)

```go
fmt.Println("Nilai melalui pointer:", *ptr)
```

`*ptr` artinya kita mengambil nilai dari alamat yang disimpan di pointer `ptr`.

## Pointer dan Fungsi

Pointer berguna banget saat kita ingin mengubah nilai variable dari dalam fungsi.

### 1. Contoh tanpa pointer (nilai tidak berubah):

```go
func ubah(x int) {
    x = 100
}

func main() {
    a := 10
    ubah(a)
    fmt.Println(a) // Tetap 10
}
```

### 2. Contoh dengan pointer (nilai berubah):

```go
func ubah(x *int) {
    *x = 100
}

func main() {
    a := 10
    ubah(&a)
    fmt.Println(a) // Jadi 100
}
```

## Pointer Default (`new`)

Kita juga bisa membuat pointer baru dengan `new`:

```go
angka := new(int)
*angka = 50
fmt.Println(*angka) // Output: 50
```

`new(int)` akan mengalokasikan memori untuk `int`. dan mengembalikan pointer ke alamat tersebut.

## Kesimpulan

Pointer di Golang memungkinkan kita bekerja langsung dengan referensi memori. Dengan memahami cara mengambil alamat (`&`) dan mengakses nilai dari pointer (`*`), kita bisa menulis kode yang lebih efisien dan fleksibel, terutama saat bekerja dengan fungsi.

Pointer juga jadi dasar penting saat kita mulai masuk ke struktur data yang lebih kompleks seperti struct, slice, dan interface.

Pada artikel selanjutnya, kita semua akan belajar tentang Array, Slice, dan Map di dalam Golang.
