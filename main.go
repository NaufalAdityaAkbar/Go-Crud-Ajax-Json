package main

import (
	"net/http"

	"github.com/jeypc/uas_semester4/controllers/mahasiswacontroll"
)

func main() {
	//Menyambungkan ke dalam mahasiswacontrroll
	http.HandleFunc("/", mahasiswacontroll.Index)
	//menarik data baru untuk show form.html
	http.HandleFunc("/mahasiswa/get_formpopup", mahasiswacontroll.GetForm)
	//Menambahkan data dan update
	http.HandleFunc("/mahasiswa/tambah", mahasiswacontroll.Tambah)
	//Hapus data mahasiswa
	http.HandleFunc("/mahasiswa/hapus", mahasiswacontroll.Hapus)
	
	http.HandleFunc("/mahasiswa/cari", mahasiswacontroll.Cari)

	http.HandleFunc("/mahasiswa/get_data", mahasiswacontroll.Index)

	http.ListenAndServe(":1818", nil)
}