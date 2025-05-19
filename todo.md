[-] Login 
 - Register user

[-] Login as (user)
 - Cek Harga + kirim *limit 5 tempat
 - Cek Paket

[-] Login as (kurir)
 - Lihat Assigned Paket
     - Update Status Paket

[-] Login as (admin)
 - Tambah Kurir (CRUD)
 - Assign Paket Ke Kurir

// Struct Tempat
    - city
    - longtitude
    - langtitude

// Struct User
    - username
    - password
    - role

// Struct Paket
    - tipe
    - berat
    - harga
    - route -> []array of kota
    - status -> []array of string
    - createdAt

//user
- tambah paket [X]
- cek paket    [X]

//admin
- tambah paket []
- lihat paket  [-]
- update paket []
- hapus paket  []

- tambah kurir []
- lihat kurir  []
- update kurir []
- hapus kurir  []

- assign paket ke kurir []

//kurir
- lihat assigned kurir []
- lihat paket          []

//system
- hitung harga         []