package mahasiswamodel

import (
	"database/sql"

	"github.com/jeypc/uas_semester4/config"
	"github.com/jeypc/uas_semester4/entities"
)

type MahasiswaModel struct {
	db *sql.DB
}

func New() *MahasiswaModel{
	db, err := config.DBconnection()
	if err != nil{
		panic(err)
	}
	return &MahasiswaModel{db: db}
}
	//Pengambilan data
func (m *MahasiswaModel) FindAll(mahasiswa *[]entities.Mahasiswa) error{
	rows, err := m.db.Query("select * from mahasiswa")
	if err != nil{
		return err
	}

	defer rows.Close()

	for rows.Next(){
		var data entities.Mahasiswa
		rows.Scan(
			&data.Id,
			&data.Namalengkap,
			&data.JenisKelamin,
			&data.TempatLahir,
			&data.TanggalLahir,
			&data.Alamat)
			
			*mahasiswa = append(*mahasiswa, data)
	}
	return nil
	
}

//simpan data atau insert
func (m *MahasiswaModel) Create(mahasiswa *entities.Mahasiswa) error{
	result, err := m.db.Exec("insert into mahasiswa(nama_lengkap, jenis_kelamin, tempat_lahir, tanggal_lahir, alamat)values(?,?,?,?,?)",
		mahasiswa.Namalengkap, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat)
	
	if err != nil{
		return err
	}
	//Menambah id(Auto Increment)
	lastInsertId, _ := result.LastInsertId()
	mahasiswa.Id = lastInsertId
	return nil
	
}

//Pemanggilan data melalu Id
func(m *MahasiswaModel) Find(id int64, mahasiswa *entities.Mahasiswa)error {
	return m.db.QueryRow("select * from mahasiswa where id = ?", id).Scan(
			&mahasiswa.Id,
			&mahasiswa.Namalengkap,
			&mahasiswa.JenisKelamin,
			&mahasiswa.TempatLahir,
			&mahasiswa.TanggalLahir,
			&mahasiswa.Alamat)
			
}
// Update data mahasiswa berdasarkan ID
func (m *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {
    _, err := m.db.Exec("update mahasiswa SET nama_lengkap=?, jenis_kelamin=?, tempat_lahir=?, tanggal_lahir=?, alamat=? where id=?",
        mahasiswa.Namalengkap, mahasiswa.JenisKelamin, mahasiswa.TempatLahir, mahasiswa.TanggalLahir, mahasiswa.Alamat, mahasiswa.Id)
    if err != nil {
        return err
    }
    return nil
}

//Delete data mahasiswa berdasarkan ID
func (m *MahasiswaModel) Delete(id int64) error {
    _, err := m.db.Exec("delete from mahasiswa where id = ?", id)
    if err != nil {
        return err
    }
    return nil
}

// Pencarian data berdasarkan nama
func (m *MahasiswaModel) SearchByName(nama_lengkap string, mahasiswa *[]entities.Mahasiswa) error {
	rows, err := m.db.Query("SELECT * FROM mahasiswa WHERE nama_lengkap LIKE ?", "%"+nama_lengkap+"%")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var data entities.Mahasiswa
        err := rows.Scan(
            &data.Id,
            &data.Namalengkap,
            &data.JenisKelamin,
            &data.TempatLahir,
            &data.TanggalLahir,
            &data.Alamat)
        if err != nil {
            return err
        }

        *mahasiswa = append(*mahasiswa, data)
    }
    return nil
}




