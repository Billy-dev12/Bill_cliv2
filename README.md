# 🛠️ Bill CLI v2 - Auto Push GitHub

Sebuah *tool* CLI (Command Line Interface) interaktif berbasis **Go** yang dirancang khusus untuk mempermudah *workflow* ngoding di Linux (dioptimalkan untuk pengguna Arch Linux/Hyprland). 

Aplikasi ini dibuat untuk menyelesaikan masalah klasik Linux: Malas mengetik rentetan perintah Git dan bosan dimintain *username/password* terus-menerus saat menggunakan URL HTTPS!

---

## ✨ Fitur Utama

* **Smart Git Config:** Otomatis mendeteksi jika Git global kamu belum disetting (Email & Nama), dan meminta konfigurasinya sekali saja.
* **Auto Push Modal Link:** Cukup masukkan link repository GitHub kosong dan pesan commit, aplikasi langsung mengeksekusi `git init` sampai `git push`.
* **Zero Password Prompt (Bypass Token):** Menyimpan Token GitHub kamu secara aman di folder konfigurasi lokal dan menyisipkannya langsung ke URL Git. Kamu bebas dari mengetik password selamanya!
* **Auto Conflict Solver (Clean Origin):** Otomatis menghapus *remote origin* lama jika folder yang kamu gunakan sudah pernah diotak-atik Git-nya, mencegah terjadinya bentrok jalur *push*.

---

## 📂 Struktur Penyimpanan Token (Khusus Linux)

Agar kamu tidak perlu menginput token baru setiap kali berpindah folder project, aplikasi ini akan membuat file rahasia di folder Home kamu:
`~/.config/bill/config.json`

> ⚠️ **Catatan Keamanan:** File ini diatur dengan izin hak akses `0600`, yang berarti hanya akun pengguna Linux kamu saja yang bisa membaca dan menulis file ini!

---

## 🚀 Cara Menggunakan

1. Jalankan aplikasi Bill CLI di dalam folder project yang ingin kamu upload.
2. Saat pertama kali digunakan, aplikasi akan meminta **Personal Access Token (PAT)** GitHub kamu.
3. Masukkan link repository GitHub yang sudah kamu buat di web.
4. Tulis pesan commit (atau tekan Enter untuk menggunakan pesan *default* `"Commit by bill cli"`).
5. *Boom!* Kodingan kamu mendarat di GitHub tanpa ribet!

---

## 🔑 Cara Mengambil Token di GitHub

Karena GitHub sudah tidak mengizinkan login menggunakan password biasa di terminal, kamu wajib membuat Token sebagai penggantinya:

1. Buka **GitHub** -> Klik Foto Profil -> **Settings**.
2. Scroll ke menu paling bawah di sebelah kiri -> Klik **Developer Settings**.
3. Pilih **Personal Access Tokens** -> Klik **Tokens (classic)**.
4. Klik **Generate new token (classic)**.
5. Beri nama token tersebut (misal: `Laptop Arch Bill`).
6. **PENTING:** Centang bagian **`repo`** (Full control of private repositories).
7. Klik **Generate token**, lalu copy kodenya dan masukkan saat aplikasi Bill CLI memintanya!

---
*Dibuat dengan penuh dedikasi, kopi hangat, dan olahraga kelingking di atas keyboard Mechanical Blue Switch selama 8 jam.* ☕💪🚀
