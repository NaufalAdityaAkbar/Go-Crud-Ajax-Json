package mahasiswacontroll

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/jeypc/uas_semester4/entities"
	mahasiswamodel "github.com/jeypc/uas_semester4/models"
)

var mahasiswaModel = mahasiswamodel.New()
//Aturan proses program
func Index(w http.ResponseWriter, r *http.Request){
	//Memanggil data untuk ditampilkan di index.html
	data := map[string]interface{}{
		"data": template.HTML( GetData()),
	}


	//Menghubungkan -> di arahkan kedalam index.html
	temp, _:= template.ParseFiles("views/mahasiswa/index.html")
	temp. Execute(w, data)
}

//Proses Meminta data
func GetData() string{
	buffer := &bytes.Buffer{}

	temp, _:= template.New("data.html").Funcs(template.FuncMap{
		
		//mengembalikan fungsi integer
		"increment": func(a, b int) int{
			return a + b
		},
		//untuk data yang akan ditampilkan kedalam html
	}).ParseFiles("views/mahasiswa/data.html")
	

	var mahasiswa []entities.Mahasiswa
	err := mahasiswaModel.FindAll(&mahasiswa)
	if err != nil{
		panic(err)
	}
	data := map[string]interface{}{
		"mahasiswa": mahasiswa,
	}
	//mengeksekusi data.html
	temp.ExecuteTemplate(buffer, "data.html", data)

	return buffer.String()
}
// Pencarian berdasarkan nama
func Cari(w http.ResponseWriter, r *http.Request) {
    // Parse form untuk mendapatkan nilai query dari URL
    r.ParseForm()
    nama := r.Form.Get("query")

    // Jika input pencarian kosong, tampilkan halaman pencarian
    if nama == "" {
        tmpl, err := template.ParseFiles("views/mahasiswa/cari.html")
        if err != nil {
            ResponseError(w, http.StatusInternalServerError, err.Error())
            return
        }
        err = tmpl.Execute(w, nil)
        if err != nil {
            ResponseError(w, http.StatusInternalServerError, err.Error())
            return
        }
        return
    }

    // Lakukan pencarian berdasarkan nama
    var mahasiswa []entities.Mahasiswa
    err := mahasiswaModel.SearchByName(nama, &mahasiswa)
    if err != nil {
        ResponseError(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Data untuk diberikan kepada template HTML
    data := map[string]interface{}{
        "mahasiswa": mahasiswa,
    }

    // Parse template HTML dan kirimkan sebagai respons
    tmpl, err := template.ParseFiles("views/mahasiswa/data.html")
    if err != nil {
        ResponseError(w, http.StatusInternalServerError, err.Error())
        return
    }

    // Eksekusi template dengan data yang diberikan dan kirimkan ke client
    err = tmpl.Execute(w, data)
    if err != nil {
        ResponseError(w, http.StatusInternalServerError, err.Error())
        return
    }
}


//Memanggil form inputan
func GetForm(w http.ResponseWriter, r *http.Request){

	//Pemanggilan data edit
	queryString := r.URL.Query()

	//convert jadi int
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil{

		data = map[string]interface{}{
			"title":"Add Data Mahasiswa",
		} 

		}else {
			var mahasiswa entities.Mahasiswa
			err := mahasiswaModel.Find(id, &mahasiswa)
			if err != nil{
				panic(err)
			}

		data = map[string]interface{}{
			"title":"Ubah Data Mahasiswa",
			"mahasiswa":mahasiswa,
		}
	}


	temp, _:= template.ParseFiles("views/mahasiswa/formpopup.html")
	temp. Execute(w, data)
}

//Post
func Tambah(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{

		r.ParseForm()
		var mahasiswa entities.Mahasiswa

		mahasiswa.Namalengkap = r.Form.Get("nama_lengkap")
		mahasiswa.JenisKelamin = r.Form.Get("jenis_kelamin")
		mahasiswa.TempatLahir = r.Form.Get("tempat_lahir")
		mahasiswa.TanggalLahir = r.Form.Get("tanggal_lahir")
		mahasiswa.Alamat = r.Form.Get("alamat")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil{
			//insert
			err := mahasiswaModel.Create(&mahasiswa)
			if err != nil{
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"Message": "Data berhasil ditambahkan",
				//Menarik GetData 
				"data": template.HTML(GetData()),
			} 
		}else{
		//update
		mahasiswa.Id = id
			err := mahasiswaModel.Update(mahasiswa)
			if err != nil{
				ResponseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{
				"Message": "Data berhasil Ubah",
				//Menarik GetData 
				"data": template.HTML(GetData()),
			} 
		}
		ResponseJson(w, http.StatusOK, data)
	}

}


//Delete
func Hapus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)
	
	if err != nil{
		panic(err)
	}

	//memanggil method
	err = mahasiswaModel.Delete(id)
	
	if err != nil{
		panic(err)
	}
	// Buat instance dari entitas mahasiswa

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	// Kirim respons berhasil
	ResponseJson(w, http.StatusOK, data)
}


//handle error
func ResponseError(w http.ResponseWriter, code int, message string){
	ResponseJson(w, code, map[string]string{"error": message})
}

//problem error notification
func ResponseJson(w http.ResponseWriter, code int, playload interface{}){
	response, _ := json.Marshal(playload)
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}