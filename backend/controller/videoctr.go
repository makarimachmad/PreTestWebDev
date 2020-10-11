package video

import (
    "context"
    "fmt"
    "PreTestSertifikasiWebDEV/backend/config"
    "PreTestSertifikasiWebDEV/backend/models"
    "log"
    "errors"
    "database/sql"
)
 
const   table = "kamutube"
// ambil semua data dari database
func GetAll(ctx context.Context) ([]models.Video, error) {
 
    var videos []models.Video
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v", table)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var vd models.Video
 
        if err = rowQuery.Scan(
			&vd.Id,
            &vd.Src,
            &vd.Title,
			&vd.Artist,
			&vd.Sifat,
			&vd.Deskripsi);
			err != nil && sql.ErrNoRows != nil{
            return nil, err
        } 
        videos = append(videos, vd)
    }
 
    return videos, nil
}

// menambahkan data ke database
func Tambah(ctx context.Context, vd models.Video) error {
	db, err := config.Mysql()

	if err != nil {
		log.Fatal("gagal terhubung ke MySQL", err)
	}

    queryText := fmt.Sprintf("INSERT INTO %v (src, title, artist, sifat, deskripsi) values('%v','%v','%v','%v','%v')",
        table,
		vd.Src,
		vd.Title,
		vd.Artist,
		vd.Sifat,
		vd.Deskripsi,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// memperbarui data
func Update(ctx context.Context, vd models.Video) error {
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("Gagal terhubung MySQL", err)
    }

    queryText := fmt.Sprintf("UPDATE %v set src = '%s', title ='%s', artist = '%s', sifat = '%s', deskripsi = '%s' where id = '%d'",
        table,
        vd.Src,
        vd.Title,
		vd.Artist,
		vd.Sifat,
		vd.Deskripsi,
        vd.Id,
    )
    fmt.Println(queryText)
 
    _, err = db.ExecContext(ctx, queryText)
 
    if err != nil {
        return err
    }
 
    return nil
}

// hapus data
func Hapus(ctx context.Context, vd models.Video) error {
 
    db, err := config.Mysql()
 
    if err != nil {
        log.Fatal("gagal terhubung ke MySQL", err)
    }
 
    queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, vd.Id)
 
    s, err := db.ExecContext(ctx, queryText)
 
    if err != nil && err != sql.ErrNoRows {
        return err
    }
 
    check, err := s.RowsAffected()
    fmt.Println(check)
    if check == 0 {
        return errors.New("id tidak ada gan..")
    }
 
    return nil
}