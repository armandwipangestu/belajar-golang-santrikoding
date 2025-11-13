# Belajar Golang Dasar #6: Operator

Dalam Golang, operator digunakan untuk melakukan operasi terhadap nilai atau variable, seperti penjumlahan, perbandingan, atau logika.

Operator ini jadi bagian penting dari logika program, karena akan sering kita gunakan saat membuat kondisi, perhitungan, dan kontrol alur. Dan berikut ini jenis-jenis operator yang ada di Golang.

## 1. Operator Aritmatika

Operator ini digunakan untuk operasi matematika dasar:

| Operator | Keterangan          |
| -------- | ------------------- |
| `+`      | Penjumlahan         |
| `-`      | Pengurangan         |
| `*`      | Perkalian           |
| `/`      | Pembagian           |
| `%`      | Modulus (sisa bagi) |

Contoh:

```go
a := 10
b := 3
fmt.Println(a + b) // Output: 13
fmt.Println(a % b) // Output: 1
```

## 2. Operator Perbandingan

Operator ini menghasilkan nilai boolean (`true` atau `false`) berdasarkan hasil perbandingan dua nilai.

| Operator | Keterangan        |
| -------- | ----------------- |
| `==`     | Sama dengan       |
| `!=`     | Tidak sama dengan |
| `>`      | Lebih besar       |
| `<`      | Lebih kecil       |
| `>=`     | Lebih besar/sama  |
| `<=`     | Lebih kecil/sama  |

Contoh:

```go
x := 5
y := 7
fmt.Println(x < y) // true
fmt.Println(x == y) // false
```

## 3. Operator Logika

Digunakan untuk menggabungkan kondisi boolean.

| Operator | Keterangan |
| -------- | ---------- |
| `&&`     | AND        |
| `!!`     | OR         |
| `!`      | NOT        |

Contoh:

```go
a := true
b := false
fmt.Println(a && b) // false
fmt.Println(a || b) // true
fmt.Println(!a) // False
```

## Kesimpulan

Operator di Golang sangat penting untuk membangun ekspresi dan logika dalam program. Kita bisa melakukan perhitungan, perbandingan, dan pengolahan kondisi dengan mudah menggunakan operator aritmatika, perbandingan, dan logika.

Di artikel selanjutnya, kita akan mempelajari bagaimana menggunakan Control Flow seperti `if`, `switch`, dan perulangan `for` di bahasa golang.
