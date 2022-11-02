
>**Todo List App**
Aplikasi yang umumnya digunakan untuk memelihara tugas sehari-hari atau membuat daftar semua yang harus dilakukan, dengan urutan prioritas tugas tertinggi hingga terendah. Sangat membantu dalam merencanakan jadwal harian.

> Pada applikasi ini terdiri dari 2 table, yaitu task dan subtask, dimana subtask merupakan bagian dari task. Kita bisa menambahkan task dan subtask.

## Requirement
Sebelum Menjelakan aplikasi pastikan telah  memiliki aplikasi berikut di local anda : 
### Teknologi dan library yang wajib digunakan:
| Teknologi   | Version | Link |
| ----------- | ---------------- | ------------------- |
| Golang      | v1.18 or later   | [Go Download](https://go.dev/dl)  |
| Go Echo Framework     | v4     | [Echo Installation](https://echo.labstack.com/guide/#installation) | 
| GORM | v2 | [GORM Installation](https://gorm.io/docs/#Install) |
| PostgreSQL | v13 or later | [PostgreSQL Download](https://www.postgresql.org/download/) |
## Cara Menjalankan: 
clone repository : 
```
git clone https://github.com/hafis915/todolist.git
```
jalankan :
```
go run main.go
```
pada folder conf terdapat file app.ini, pada file tersebut ubah configurasi yang ada dengan configurasi yang dibutuhkan.
** Untuk Path upload local , Jika ingin mengikuti config file yang sudah ada. tambahkan directory dengan nama assets pada root file**


## List API:

Buat API Todo List dengan kriteria sebagai berikut :
1. Menampilkan data list tanpa sub list ( dengan dan tanpa pagination ).
```
route tanpa pagination: http://localhost:8000/tasks
route dengan pagination: http://localhost:8000/tasks?page=1&perpage=100
Method : GET
```
2. Menampilkan data list beserta dengan sub list nya jika ada.
```
route  : http://localhost:8000/tasksubtask
Method : GET
```
3. Menampilkan data sub list by list id.
```
route : http://localhost:8000/subtask/:taskid
method : GET
```
4. Menambahkan data list.
```
route : http://localhost:8000/task
method : POST
untuk field file digunakan saat melakukan hit dengan FormData.
```
**Body :**
|field| data type | isRequired|
|-----|-----------|-----------|
|title|string|true
|description|string|true|
|file|file|false|

5. Menambahkan data sub list untuk spesifik list.
```
route : http://localhost:8000/subtask
method : POST
```
**Body :**
|field| data type | isRequired|
|-----|-----------|-----------|
|title|string|true
|description|string|true|
|task_id|string|true|
|file|file|false


6. Mengubah data list/sub list dengan kritera input diatas.

**Subtask**
```
route mengubah data task : http://localhost:8000/subtask/:id
Method : PUT
```
Body :
|field| data type | isRequired|
|-----|-----------|-----------|
|title|string|false
|description|string|false|
|task_id|string|false|
|status|integer|false|
|file|file|false

**Task**
```
route mengubah data task : http://localhost:8000/task/:id
Method : PUT
```
|field| data type | isRequired|
|-----|-----------|-----------|
|title|string|false
|description|string|false|
|status|integer|false|
|file|file|false

7. Menghapus data list/sub list.

**task**
```
route : http://localhost:8000/task/:id
method: Delete
```
**subtask**
```
route : http://localhost:8000/subtask/:id
method: Delete
```
