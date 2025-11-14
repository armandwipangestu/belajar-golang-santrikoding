# Belajar Golang Dasar #12: Method

Di artikel sebelumnya, kita udah belajar tentang `struct`, cara bikin tipe data sendiri dengan field-field yang bsia kita sesuaikan. Nah, sekarang bayangin kalau kita bisa ngasih fungsi ke struct itu semacam "kemampuan" tambahan yang nempel ke data. Di Golang, itu bisa dilakukan pakai method.

## Apa itu Method?

Method adalah fungsi yang menempel ke suatu tipe data (biasanya struct). Dengan method, kita bisa bikin fungsi yang bisa langsung dipanggil oleh objek dari struct tersebut, mirip kayak OOP di bahasa lain, tapi dengan gaya khas Golang.

## Cara Mendefinisikan Method

Bentuk dasarnya seperti ini

```go
func (r ReceiverType) namaMethod() {
    // isi fungsi
}
```

-   `ReceiverType` adalah tipe data yang mau kita kasih method (biasanya struct).
-   Receiver-nya bisa disingkat menjadi satu atau dua huruf (bebas)
-   Method bisa pakai value receiver atau pointer receiver (nanti kita bahas bedanya).

## Contoh Method Sederhana

```go
type Mahasiswa struct {
    Nama string
    Umur int
}

func (m Mahasiswa) Perkenalan() {
    fmt.Printf("Halo, nama saya %s dan umur saya %d\n", m.Nama, m.Umur)
}
```

Penggunaan:

```go
func main() {
    mhs := Mahasiswa{Nama: "Adit", Umur: 21}
    mhs.Perkenalan()
}
```

## Method dengan Return Value

Method bisa juga mengembalikan nilai:

```go
func (m Mahasiswa) TahunLahir(tahunSekarang int) int {
    return tahunSekarang - m.Umur
}
```

## Pointer Receiver vs Value Receiver

Saat kita bikin method di Golang, kita bisa memilih apakah method tersebut menerima value dari struct-nya atau menerima pointer ke struct-nya. Kedua pendekatan ini punya efek yang berbeda, terutama soal apakah data aslinya akan berubah atau tidak.

Kalau pakai value receiver, struct yang dikirim ke method hanyalah salinan. Jadi perubahan yang dilakukan di dalam method nggak akan mempengaruhi data aslinya.

Sedangkan kalu pakai pointer receiver, method menerima alamat dari struct tersebut. Artinya, perubahan yang kita lakukan di method akan langsung mengubah isi struct yang asli.

Supaya lebih jelas, yuk kita lihat contoh dan perbandingannya.

### 1. Value Receiver (copy)

```go
func (m Mahasiswa) UbahNamaBaru(nama string) {
    m.Nama = nama
}
```

Efeknya tidak permanen, karena `m` hanya salinan dari struct aslinya.

### 2. Value Pointer (langsung ubah)

```go
func (m *Mahasiswa) UbahNama(nama string) {
    m.Nama = nama
}
```

Kalau pakai `*Mahasiswa`, kita beneran mengubah data aslinya. Contoh:

```go
func main() {
    mhs := Mahasiswa{Nama: "Rudi", Umur: 22}
    mhs.UbahNama("Santi")
    fmt.Println(mhs.Nama) // Output: Santi
}
```

## Kenapa Method Penting?

Method bikin struct kita lebih rapi dan terstruktur. Daripada bikin fungsi terpisah yang terima struct sebagai parameter, kita bisa langsung nempelin fungsinya ke struct itu sendiri. Ini bikin kode lebih jelas dan modular.

## Kesimpulan

Method di Golang adalah fungsi yang menempel ke tipe data (biasanya struct). Dengan method, kita bisa ngasih "kemampuan" tambahan ke data yang kita buat.

Pada artikel selanjutnya, kita semua akan belajar dan memahami apa tentang interface di dalam pemrograman golang.
