package handler

import (
	"fmt"
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HalloHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Selamat datang di perpustakaa Golang"))
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layouts.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"id":   "1",
	// 	"nama": "Belajar Golang Fundamental",
	// }

	// data := entity.Buku{ID: 1, Nama: "Belajar Golang Fundamental", Stock: 3}

	data := []entity.Buku{
		{ID: 1, Nama: "Belajar Golang Fundamental", Stock: 1},
		{ID: 2, Nama: "SiGolang Si Bocah belajar Golanng", Stock: 3},
		{ID: 3, Nama: "GoGrak : Go-React", Stock: 2},
	}
	erro := tmpl.Execute(w, data)
	if erro != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}
}

func BukuHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Daftar Buku Nomor : %d", idNumb)
}

func ListbukuHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Selamat datang di perpustakaa Golang"))
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layouts.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Buku{
		{ID: 1, Nama: "Belajar Golang Fundamental", Stock: 4},
		{ID: 2, Nama: "SiGolang Si Bocah belajar Golanng", Stock: 3},
		{ID: 3, Nama: "GoGrak : Go-React", Stock: 5},
	}
	erro := tmpl.Execute(w, data)
	if erro != nil {
		log.Println(err)
		http.Error(w, "Error is Happening, Keep Calm", http.StatusInternalServerError)
		return
	}
}
