# 🚀 Tutorial Golang - Dari Nol Sampai Backend

Repository ini berisi tutorial Golang lengkap dengan contoh kode yang bisa langsung dijalankan. Setiap file adalah satu konsep yang harus dipelajari secara berurutan.

> ⚠️ **PENTING:** Jangan lompat-lompat! Ikuti urutan dari atas ke bawah.

---

## 📁 Struktur Repository

```
golang-tutorial/
├── main.go                    # 1. Hello World
├── variabel.go                # 2. Variabel & Tipe Data
├── loop.go                    # 3. Kondisi & Loop
├── function.go                # 4. Function & Multiple Return
├── array.go                   # 5. Slice (Array Dinamis)
├── error-handling.go          # 6. Error Handling
├── struct.go                  # 7. Struct (OOP Lite)
├── interfce.go                # 8. Interface
├── Goroutine-channel.go       # 9. Concurrency
├── backend.go                 # 10. HTTP Server Dasar
├── API.go                     # 11. REST API dengan Gin
└── README.md
```

---

## 🎯 Urutan Belajar (WAJIB DIIKUTI)

### 1️⃣ Hello World → `main.go`

**Konsep:** Package, Import, Function main()

```bash
go run main.go
```

**Output:**
```
Hello World!
```

**Yang Dipelajari:**
- Struktur program Go
- `package main` sebagai entry point
- `func main()` sebagai starting point
- `fmt.Println()` untuk output

---

### 2️⃣ Variabel & Tipe Data → `variabel.go`

**Konsep:** Deklarasi variabel, short declaration (`:=`)

```bash
go run variabel.go
```

**Output:**
```
10 20 Iqbal true
```

**Yang Dipelajari:**
- `var x int = 10` (deklarasi lengkap)
- `y := 20` (short declaration - CIRI KHAS GO)
- Tipe data otomatis
- String, int, bool

---

### 3️⃣ Kondisi & Loop → `loop.go`

**Konsep:** if-else, for loop

```bash
go run loop.go
```

**Output:**
```
Besar
0
1
2
3
4
Count: 0
Count: 1
Count: 2
```

**Yang Dipelajari:**
- `if-else` statement
- `for` loop (Go tidak punya `while`!)
- `for` sebagai while
- Infinite loop dengan `break`

---

### 4️⃣ Function → `function.go`

**Konsep:** Function, parameter, multiple return values

```bash
go run function.go
```

**Output:**
```
Hasil: 8
Hasil bagi: 5
```

**Yang Dipelajari:**
- Function dengan parameter
- Return value
- **Multiple return values** (hasil + error)
- Error handling dengan `nil`

📌 **PENTING:** Multiple return adalah ciri khas Go untuk error handling!

---

### 5️⃣ Slice → `array.go`

**Konsep:** Slice (array dinamis), append, loop dengan range

```bash
go run array.go
```

**Output:**
```
1
5
Index 0: 1
Index 1: 2
Index 2: 3
Index 3: 4
Index 4: 5
```

**Yang Dipelajari:**
- Slice vs Array
- `append()` untuk menambah elemen
- `len()` untuk mendapat panjang
- Loop dengan `range`

💡 **Tips:** Gunakan slice, bukan array! Slice lebih flexible.

---

### 6️⃣ Error Handling → `error-handling.go`

**Konsep:** Error handling ala Go (tanpa try-catch)

```bash
go run error-handling.go
```

**Output (jika file tidak ada):**
```
Error cuy, ini penjelasanya: open data.txt: no such file or directory
```

**Yang Dipelajari:**
- Error sebagai return value
- Check error dengan `if err != nil`
- `defer` untuk cleanup (menutup file)

📌 **KRUSIAL:** Go tidak pakai try-catch. Error adalah nilai yang dikembalikan!

---

### 7️⃣ Struct → `struct.go`

**Konsep:** Struct sebagai pengganti class

```bash
go run struct.go
```

**Output:**
```
Budi
```

**Yang Dipelajari:**
- Definisi struct
- Create instance
- Akses dan update field
- Struct adalah fondasi OOP di Go

---

### 8️⃣ Interface → `interfce.go`

**Konsep:** Interface untuk polymorphism

```bash
go run interfce.go
```

**Output:**
```
Woof!
Meow!
```

**Yang Dipelajari:**
- Interface mendefinisikan behavior
- Implicit implementation (tidak perlu `implements`)
- Polymorphism di Go
- Function yang menerima interface

💡 **Konsep Inti:** Interface membuat kode flexible dan testable!

---

### 9️⃣ Goroutine & Channel → `Goroutine-channel.go`

**Konsep:** Concurrency dengan goroutine dan channel

```bash
go run Goroutine-channel.go
```

**Output (contoh):**
```
Worker 1 mulai job 1
Worker 2 mulai job 2
Worker 3 mulai job 3
Worker 1 mulai job 4
Worker 2 mulai job 5
Hasil: 2
Hasil: 4
Hasil: 6
Hasil: 8
Hasil: 10
```

**Yang Dipelajari:**
- Goroutine dengan keyword `go`
- Channel untuk komunikasi antar goroutine
- Buffered channel
- `close()` channel
- Worker pool pattern

🔥 **Ini alasan utama orang pakai Go!** Concurrency yang mudah.

---

### 🔟 HTTP Server Dasar → `backend.go`

**Konsep:** Web server dengan `net/http` bawaan

```bash
go run backend.go
```

**Test dengan browser atau curl:**
```bash
curl http://localhost:8090/
# Output: Hello Server

curl http://localhost:8090/api
# Output: {"message": "Hello API"}
```

**Yang Dipelajari:**
- `http.HandleFunc()` untuk routing
- Handler function
- ResponseWriter & Request
- Set header
- Listen & Serve

---

### 1️⃣1️⃣ REST API dengan Gin → `API.go`

**Konsep:** REST API dengan framework Gin

**Install dependency terlebih dahulu:**
```bash
go mod init golang-tutorial
go get -u github.com/gin-gonic/gin
```

**Jalankan:**
```bash
go run API.go
```

**Test API:**

**1. GET all users:**
```bash
curl http://localhost:8090/users
```
Output:
```json
[
  {"id":1,"name":"Budi"},
  {"id":2,"name":"Ani"}
]
```

**2. GET user by ID:**
```bash
curl http://localhost:8090/users/1
```
Output:
```json
{"message":"User ID 1"}
```

**3. POST create user:**
```bash
curl -X POST http://localhost:8090/users \
  -H "Content-Type: application/json" \
  -d '{"id":3,"name":"Citra"}'
```
Output:
```json
{"id":3,"name":"Citra"}
```

**Yang Dipelajari:**
- Gin framework
- RESTful routing (GET, POST)
- JSON binding
- Route parameters (`:id`)
- HTTP status codes
- Error handling di API

---

## 🛠️ Setup & Instalasi

### Prasyarat
- Go versi 1.21+ ([Download](https://golang.org/dl/))
- Text editor (VS Code, GoLand, vim, dll)

### Cek Instalasi
```bash
go version
```

### Clone Repository
```bash
git clone <repository-url>
cd golang-tutorial
```

### Install Dependencies (untuk API.go)
```bash
go mod init golang-tutorial
go mod tidy
```

---

## 📚 Cara Menggunakan Tutorial Ini

### Metode 1: Jalankan Satu per Satu
```bash
go run main.go
go run variabel.go
go run loop.go
# ... dan seterusnya
```

### Metode 2: Build ke Binary
```bash
# Build file tertentu
go build main.go
./main

# Build dengan nama custom
go build -o myapp variabel.go
./myapp
```

### Metode 3: Eksperimen
- Buka file dengan editor
- Modifikasi kode
- Jalankan ulang dengan `go run`
- Lihat perubahan output

---

## 🎓 Tips Belajar

### ✅ DO (Lakukan):
1. **Ikuti urutan** - jangan skip!
2. **Ketik ulang kode** - jangan copy-paste
3. **Eksperimen** - ubah nilai dan lihat hasilnya
4. **Buat catatan** - tulis apa yang dipahami
5. **Praktek setiap hari** - minimal 1 jam

### ❌ DON'T (Jangan):
1. Lompat ke framework sebelum paham dasar
2. Hanya membaca tanpa praktek
3. Copy-paste tanpa memahami
4. Skip error handling
5. Belajar terlalu cepat

---

## 🐛 Troubleshooting

### Error: `package not found`
```bash
go mod init golang-tutorial
go mod tidy
```

### Error: `port already in use`
Ganti port di `backend.go` dan `API.go`:
```go
// Dari :8090 ke :8091
http.ListenAndServe(":8091", nil)
```

### Error: `go: command not found`
Install Go terlebih dahulu dari [golang.org](https://golang.org/dl/)

---

## 📖 Konsep Penting Go

### Ciri Khas Go:
- ✅ **Statically typed** - tapi dengan type inference
- ✅ **Compiled** - jadi binary, bukan interpreter
- ✅ **Garbage collected** - otomatis memory management
- ✅ **Concurrency built-in** - goroutine & channel
- ✅ **Simple syntax** - mudah dibaca
- ✅ **Fast compilation** - build cepat
- ✅ **No classes** - pakai struct + interface

### Yang TIDAK Ada di Go:
- ❌ Class inheritance
- ❌ While loop (semua pakai `for`)
- ❌ Try-catch (pakai multiple return)
- ❌ Ternary operator (`? :`)
- ❌ Semicolon wajib

---

## 🚀 Next Steps

Setelah selesai tutorial ini:

### Level Intermediate:
1. **Database:** Belajar GORM atau database/sql
2. **Testing:** Unit test dengan `testing` package
3. **Middleware:** Custom middleware di Gin
4. **JWT Auth:** Implementasi authentication
5. **Docker:** Containerize aplikasi Go

### Project Ideas:
1. **Todo API** - CRUD sederhana
2. **URL Shortener** - seperti bit.ly
3. **Blog API** - dengan kategori dan tags
4. **Chat App** - dengan WebSocket
5. **Monitoring Tool** - cek status website

---

## 📚 Resources Tambahan

### Dokumentasi:
- [Official Go Docs](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

### Tutorial:
- [Go Tour](https://go.dev/tour/)
- [Go Wiki](https://github.com/golang/go/wiki)

### Komunitas:
- [Golang Indonesia - Telegram](https://t.me/golangID)
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

---

## 🏆 Checklist Progress

Tandai setiap file yang sudah kamu pelajari:

```
□ main.go - Hello World
□ variabel.go - Variabel & Tipe Data
□ loop.go - Kondisi & Loop
□ function.go - Function
□ array.go - Slice
□ error-handling.go - Error Handling
□ struct.go - Struct
□ interfce.go - Interface
□ Goroutine-channel.go - Concurrency
□ backend.go - HTTP Server
□ API.go - REST API dengan Gin
```

---

## 🤝 Kontribusi

Ingin menambah contoh atau memperbaiki tutorial?
1. Fork repository ini
2. Buat branch baru
3. Commit perubahan
4. Push dan buat Pull Request

---

## 📝 Lisensi

Repository ini dibuat untuk tujuan edukasi.

---

## 💬 Feedback

Punya pertanyaan atau saran? Buat issue di repository ini!

---

**Selamat Belajar Go! 🚀**

*"Simplicity is the ultimate sophistication."* - Go Philosophy
