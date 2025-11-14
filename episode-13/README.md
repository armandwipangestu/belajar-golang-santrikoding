# Belajar Golang Dasar #13: Interface

Kalau sebelumnya kita belajar cara membuat struct dan memberi method ke struct tersebut, sekarang kita masuk ke fitur yang bikin Golang fleksibel dan elegan: Interface.

Interface di Golang adalah alat untuk mendesain kode yang lebih modular dan mudah diperluas, tanpa harus terikat pada hierarki class seperti di bahasa OOP klasik. Konsepnya sederhana, tapi penggunaannya bisa sangat kuat kalau kita udah paham.

## Apa itu Interface?

Interface adalah kumpulan definisi method tanpa implementasi. Struct yang punya method sesuai dengan definisi di interface secara otomatis mengimplementasikan interface tersebut, tapna perlu deklarasi eksplisit seperti `implemenets` di Java atau `: InterfaceName` di bahasa lain.

## Contoh Interface

Misalnya kita punya interface `BisaBerjalan`:

```go
type BisaBerjalan interface {
    Jalan()
}
```

Lalu kita punya dua struct:

```go
type Manusia struct {
    Name string
}

type Robot struct {
    Kode string
}
```

Masing-masing punya method `Jalan()`:

```go
func (m Manusia) Jalan() {
    fmt.Println(m.Nama, "Sedang berjalan")
}

func (r Robot) Jalan() {
    fmt.Println("Robot", r.Kode, "berjalan otomatis")
}
```

Karena kedua struct tersebut punya method `Jalan()`, otomatis mereka mengimplementasikan interface `BisaBerjalan`.

## Menggunakan Interface

Kita bisa bikin fungsi yang menerima parameter bertipe interface:

```go
func Gerakkan(b BisaBerjalan) {
    b.Jalan()
}
```

Lalu bisa dipakai dengan berbagai tipe:

```go
func main() {
    m := Manusia{Nama: "Andi"}
    r := Robot{Kode: "RX-01"}

    Gerakkan(m)
    Gerakkan(r)
}
```

Output

```go
Andi sedang berjalan
Robot RX-01 berjalan otomatis
```

## Interface Kosong (`interface{}`)

Di Golang, ada juga tipe `interface{}` yang disebut interface kosong. Ini bisa menerima sembarang tipe data. Mirip seperti `any` di TypeScript atau `Object` di Java.

```go
func CetakData(data interface{}) {
    fmt.Println(data)
}
```

Tapi, kalau kita mau "mengakses" nilai aslinya, kita perlu type assertion atau switch.

## Kesimpulan

Interface adalah cara Golang mengimplementasikan prinsip OOP dengan gaya yang simpel dan ringan. Dengan interface, kita bisa membuat program yang modular, mudah diuji, dan fleksibel menghadapi perubahan.

Yang unik dari Golang adalah, implementasi interface terjadi secara otomatis, selama method-nya cocok, tanpa perlu deklarasi tambahan.
