create untuk pasien
POST http://localhost:8080/pasien
{
    "nama": "Eka",
    "alamat": "Jl. Kamboja No. 5",
    "tanggal_lahir": "1985-05-05",
    "keluhan": "Demam tinggi",
    "nomor_telepon": "08123456793"
}

create untuk dokter
POST http://localhost:8080/dokter
{
    "nama": "Dr. Joko",
    "spesialis": "Kulit",
    "nomor_telepon": "08123456798"
}

create untuk kunjungan
POST http://localhost:8080/kunjungan
{
    "id_pasien": 1,
    "id_dokter": 2,
    "tanggal_kunjungan": "2024-06-05T00:00:00Z",
    "keluhan": "Demam tinggi"
}

update untuk pasien,di akhir itu id
PUT http://localhost:8080/pasien/1
{
    "nama": "Eka",
    "alamat": "Jl. Kamboja No. 5",
    "tanggal_lahir": "1985-05-05",
    "keluhan": "Demam tinggi",
    "nomor_telepon": "08123456793"
}

update untuk dokter,di akhir itu id
PUT http://localhost:8080/dokter/1
{
    "nama": "Dr. Joko",
    "spesialis": "Kulit",
    "nomor_telepon": "08123456798"
}

update untuk kunjungan,di akhir itu id
PUT http://localhost:8080/kunjungan/1
{
    "id_pasien": 1,
    "id_dokter": 2,
    "tanggal_kunjungan": "2024-06-05T00:00:00Z",
    "keluhan": "Demam tinggi"
}

untuk read menggunkan GET
http://localhost:8080/pasien
http://localhost:8080/dokter
http://localhost:8080/kunjungan

dan untuk hapus/delete menggunakan DELETE
http://localhost:8080/pasien/1
http://localhost:8080/dokter/1
http://localhost:8080/kunjungan/1
