package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	DBUsername = "root" // Ganti dengan username MySQL Anda
	DBPassword = ""     // Ganti dengan password MySQL Anda
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "tes2" // Ganti dengan nama database Anda
)

var db *sql.DB

type User struct {
	ID        int
	Username  string
	Password  string
	Email     string
	Name      string
	CreatedAt string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Membaca isi file HTML
	htmlContent, err := ioutil.ReadFile("login.html")
	if err != nil {
		http.Error(w, "Gagal membaca file HTML", http.StatusInternalServerError)
		return
	}

	// Menetapkan header Content-Type
	w.Header().Set("Content-Type", "text/html")

	// Menulis isi file HTML ke ResponseWriter
	w.Write(htmlContent)
}

func main() {
	// Inisialisasi koneksi ke database
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
	var err error
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inisialisasi HTTP server dan routing
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/index", indexHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.FormValue("username"), r.FormValue("password"))
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Cari user berdasarkan username
		user, err := findUserByUsername(username)
		if err != nil {
			http.Error(w, "Gagal melakukan login: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Verifikasi password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Username atau password salah", http.StatusUnauthorized)
			return
		}

		// Berhasil login
		w.Write([]byte("Login berhasil!"))
	} else {
		http.Error(w, "Metode request tidak diizinkan", http.StatusMethodNotAllowed)
	}
}

func findUserByUsername(username string) (User, error) {
	var user User
	query := "SELECT * FROM author WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Name, &user.CreatedAt)

	if err != nil {
		return user, err
	}

	return user, nil
}
