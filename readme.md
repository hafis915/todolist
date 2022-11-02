
## Moonlay Academy - Backend Test (GOLANG)

>**Todo List App**
Aplikasi yang umumnya digunakan untuk memelihara tugas sehari-hari atau membuat daftar semua yang harus dilakukan, dengan urutan prioritas tugas tertinggi hingga terendah. Sangat membantu dalam merencanakan jadwal harian.

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
|field| data type | isrequired|
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
|field| data type | isrequired|
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
|field| data type | isrequired|
|-----|-----------|-----------|
|title|string|true
|description|string|true|
|task_id|string|true|
|file|file|false

**Task**
```
route mengubah data task : http://localhost:8000/task/:id
Method : PUT
```
|field| data type | isrequired|
|-----|-----------|-----------|
|title|string|true
|description|string|true|
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
