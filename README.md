Tentang Proyek
Proyek ini merupakan aplikasi Go yang dibuat dengan menggunakan pola desain bersih (clean architecture) dengan tiga lapisan utama: controller, service, dan repo. Aplikasi ini memiliki API dengan semua metode terpakai (GET, POST, PUT, PATCH, DELETE) dan menyimpan data dari API ke database tanpa menggunakan ORM, namun tetap mengimplementasikan pola Repository.

Relational Database
Proyek ini memiliki setidaknya satu hubungan antara dua tabel di database. Penggunaan transaksi digunakan ketika ada penyimpanan dua data secara bersamaan untuk meningkatkan integritas data.

Cara Menjalankan
Deskripsi singkat tentang cara menjalankan proyek Anda. Pastikan untuk menyertakan prasyarat seperti versi Go, dependensi, dan langkah-langkah yang diperlukan.

bash
Copy code
# Contoh langkah menjalankan proyek
go run main.go
Design Pattern
Pola desain bersih (clean architecture) digunakan dalam proyek ini untuk memisahkan tanggung jawab di antara tiga lapisan utama: controller, service, dan repo. Dengan ini, proyek dapat dengan mudah diperluas dan diuji.

Dokumentasi Postman
Dokumentasi Postman dapat ditemukan [di sini untuk](https://documenter.getpostman.com/view/25921875/2s9YkoeN5N) memudahkan penggunaan API.

Presentasi
Link ke video presentasi menggunakan [Loom] (https://www.loom.com/share/de51348e20084e5eaad7d80d54a21a61?sid=72ffa04f-aa4c-4b36-b9e4-a0ed9c07d37f)
