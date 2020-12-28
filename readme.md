# Talent Factory Competition (Backend) - Task 3

> Buat algoritma bebas dengan menggunakan Go

## Muhammad Yusril Yuriis

### Algoritma Statistik Deskriptif (mean, median, mode)

Algoritma ini digunakan untuk menghitung rata-rata, median, dan modus dari suatu data tunggal.

- Mean (rata-rata):

  1. Jumlahkan seluruh elemen data.
  2. Bagi hasil penjumlahan tersebut dengan banyaknya data.

- Median (nilai tengah):

  1. Urutkan data dari nilai terkecil ke nilai terbesar.
  2. Untuk banyaknya data bernilai ganjil, median merupakan nilai tengahnya. Untuk banyaknya data bernilai genap, median merupakan hasil bagi penjumlahan nilai tengah-1 dan nilai tengah+1.

- Mode (nilai yg paling banyak muncul):
  1. Gunakan map (hash table) untuk menyimpan masing-masing elemen data dengan banyaknya elemen tersebut.
  2. Gunakan slice untuk menyimpan banyaknya masing-masing elemen data, lalu diurutkan.
  3. Ambil nilai terbesar dari slice no. 2, lalu pindahkan ke slice mode (karena modus bisa lebih dari 1 elemen).
