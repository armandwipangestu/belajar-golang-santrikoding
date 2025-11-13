# Belajar Golang Dasar #11: Struct

Setelah kita belajar cara menyimpan kumpulan data dengan array, slice, dan map, sekarang saatnya kenalan dengan fitur penting lainnya di Golang yaitu Struct.

Struct (singkatan dari structure) memungkinkan kita membuat tipe data sendiri yang bisa menampung berbagai informasi dalam satu wadah. Ini sangat berguna saat kita mau merepresentasikan sesuatu yang kompleks, seperti data pengguna, produk, buku, dan sebagainya.

## Apa itu Struct?

Strucrt adalah kumpulan dari beberapa field (variable) yang dikelompokkan dalam satu tipe data. Field-field ini bisa berbeda tipe, dan kita bisa memberi nama sesuai kebutuhan.

Ibaratnya struct itu kayak cetakan atau blueprint untuk bikin objek/data yang strukturnya sama.

## Membuat Struct

```go
type Mahasiswa struct {
    Nama string
    Umur int
    Aktif bool
}
```

Kita baru saja membuat tipe data baru bernama `Mahasiswa` dengan tiga field: `Nama`, `Umur`, dan `Aktif`.

## Cara Menggunakan Struct

### 1. Membuat Variable dari Struct

```go
var mhs1 Mahasiswa
mhs1.Nama = "Budi"
mhs1.Umur = 20
mhs1.Aktif = true

fmt.Println(mhs1)
```

### 2. Struct Literal (langsung isi)

```go
mhs2 := Mahasiswa{
    Nama: "Ani",
    Umur: 22,
    Aktif: false,
}
```

### 3. Akses Field

```go
fmt.Println("Nama:", mhs2.Nama)
fmt.Println("Umur:", mhs2.Umur)
```

## Struct di Dalam Struct

Struct juga bisa menampung struct lain.

```go
type Alamat struct {
    Kota string
    Jalan string
}

type Mahasiswa struct {
    Nama string
    Umur int
    Alamat Alamat
}
```

Penggunaan:

```go
mhs := Mahasiswa{
    Nama: "Citra",
    Umur: 21,
    Alamat: Alamat{
        Kota: "Bandung",
        Jalan: "Jl. Merdeka",
    },
}

fmt.Println(mhs.Alamat.Kota) // Output: Bandung
```

## Pointer ke Struct

Struct juga bisa digunakan bersama pointer.

```go
func ubahUmur(m *Mahasiswa) {
    m.Umur = 25
}

func main() {
    m := Mahasiswa{Nama: "Doni", Umur: 20}
    ubahUmur(&m)
    fmt.Println(m.Umur) // Output: 25
}
```

Dengan pointer, perubahan di dalam fungsi akan berdampang langsung ke data aslinya.

## Kesimpulan

Struct adalah fondasi untuk membuat data yang lebih kompleks di Golang. Dengan struct, kita bisa mengelompokkan data yang saling berhubungan ke dalam satu kesatuan yang rapi dan mudah digunakan.

Struct juga jadi dasar penting untuk OOP (Object-Oriented Programming) versi Golang. termasuk saat kita belajar tentang method dan interface nanti.

Pada artikel selanjutnya, kita semua akan belajar tentang Method di dalam pemrograman Golang.
