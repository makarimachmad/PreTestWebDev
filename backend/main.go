//menjalankan aksi yang ada pada controller
//sumber belajar 1. https://kodingin.com/golang-koneksi-database-mysql/

package main

import (
	"log"
	"fmt"
	"context"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"PreTestSertifikasiWebDEV/backend/controller"
	"PreTestSertifikasiWebDEV/backend/models"
	"PreTestSertifikasiWebDEV/backend/utils"
)

func main() {

	handleRequest()

	err := http.ListenAndServe(":8090", nil)
	if err != nil{
		log.Fatal(err)
	}
}

//tukang panggil fungsi
func handleRequest(){
	
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "application/json; charset=UTF-8"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})


    router.HandleFunc("/", getVideo).Methods("GET")
	router.HandleFunc("/tambah", tambahVideo).Methods("POST")
	router.HandleFunc("/update", updateVideo).Methods("PUT")
	router.HandleFunc("/update", updateVideo).Methods("PATCH")
	router.HandleFunc("/hapus", hapusVideo).Methods("DELETE")
	

	fmt.Println("Terhubung ke 8090")
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(router))
}

//fungsi middleware cek isian dan method GET untuk ambil data
func getVideo(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        ctx, cancel := context.WithCancel(context.Background())
		
        defer cancel()
		
        videos, err := video.GetAll(ctx)
 
        if err != nil {
            fmt.Println(err)
        }
 
        utils.ResponseJSON(w, videos, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak diijinkan", http.StatusNotFound)
    return
}

//fungsi middleware cek isian dan method POST untuk insert data
func tambahVideo(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
		
		//data yg dikirim harus berupa json
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var vd models.Video
 
        if err := json.NewDecoder(r.Body).Decode(&vd); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        if err := video.Tambah(ctx, vd); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        response := map[string]string{
            "status": "berhasil gan",
        }
 
        utils.ResponseJSON(w, response, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

//memperbarui data dengan fungsi middleware
func updateVideo(w http.ResponseWriter, r *http.Request) {
    if r.Method == "PUT" {
 
        if r.Header.Get("Content-Type") != "application/json" {
            http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
            return
        }
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var vd models.Video
 
        if err := json.NewDecoder(r.Body).Decode(&vd); err != nil {
            utils.ResponseJSON(w, err, http.StatusBadRequest)
            return
        }
 
        fmt.Println(vd)
 
        if err := video.Update(ctx, vd); err != nil {
            utils.ResponseJSON(w, err, http.StatusInternalServerError)
            return
        }
 
        response := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, response, http.StatusCreated)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}

//hapus data
func hapusVideo(w http.ResponseWriter, r *http.Request) {
 
    if r.Method == "DELETE" {
 
        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
 
        var vd models.Video
 
        id := r.URL.Query().Get("id")
 
        if id == "" {
            utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
            return
        }
        vd.Id, _ = strconv.Atoi(id)
 
        if err := video.Hapus(ctx, vd); err != nil {
 
            kesalahan := map[string]string{
                "error": fmt.Sprintf("%v", err),
            }
 
            utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
            return
        }
 
        response := map[string]string{
            "status": "Succesfully",
        }
 
        utils.ResponseJSON(w, response, http.StatusOK)
        return
    }
 
    http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
    return
}