# Tentang Proyek
Proyek ini merupakan aplikasi Go yang dibuat dengan menggunakan pola desain bersih (clean architecture) dengan tiga lapisan utama: controller, service, dan repo. Aplikasi ini memiliki API dengan semua metode terpakai (GET, POST, PUT, PATCH, DELETE) dan menyimpan data dari API ke database tanpa menggunakan ORM, namun tetap mengimplementasikan pola Repository.

# Relational Database
Proyek ini memiliki setidaknya satu hubungan antara dua tabel di database. di tabel tasks dan users yang dimana user memiliki koneksi one to many ke tabel tasks

# cara  menjalankan proyek
import kode dari github ke local 
install dependensi yang dibutuhkan projek 
dan jalankan ``go run main.go``

# Design Pattern
Pola desain bersih (clean architecture) digunakan dalam proyek ini untuk memisahkan tanggung jawab di antara tiga lapisan utama: controller, service, dan repo. Dengan ini, proyek dapat dengan mudah diperluas dan diuji.

# Dokumentasi Postman
Dokumentasi Postman dapat ditemukan [di sini untuk](https://documenter.getpostman.com/view/25921875/2s9YkoeN5N) memudahkan penggunaan API.

# Presentasi
Link ke video presentasi menggunakan [Loom] (https://www.loom.com/share/de51348e20084e5eaad7d80d54a21a61?sid=72ffa04f-aa4c-4b36-b9e4-a0ed9c07d37f)
