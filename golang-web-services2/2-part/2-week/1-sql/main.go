package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Item struct {
	Id          int
	Title       string
	Description string
	Updated     sql.NullString
}

type Handler struct {
	DB   *sql.DB
	Tmpl *template.Template
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	var items []*Item
	rows, err := h.DB.Query("SELECT id, title, updated FROM items")
	__err_panic(err)
	for rows.Next() {
		post := &Item{}
		err = rows.Scan(&post.Id, &post.Title, &post.Updated)
		__err_panic(err)
		items = append(items, post)
	}
	// it is important to close the connection, or it will be a flow
	rows.Close()

	err = h.Tmpl.ExecuteTemplate(w, "index.html", struct {
		Items []*Item
	}{
		Items: items,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := h.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	// validation is skipped for simplicity
	result, err := h.DB.Exec(
		"INSERT INTO items (`title`, `description`) VALUES (?, ?)",
		r.FormValue("title"),
		r.FormValue("description"),
	)
	__err_panic(err)

	affected, err := result.RowsAffected()
	__err_panic(err)
	lastID, err := result.LastInsertId()
	__err_panic(err)

	fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", lastID)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	__err_panic(err)

	post := &Item{}
	// QueryRow will close connect itself
	row := h.DB.QueryRow("SELECT id, title, updated, description FROM items WHERE id = ?", id)

	err = row.Scan(&post.Id, &post.Title, &post.Updated, &post.Description)
	__err_panic(err)

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	__err_panic(err)

	// validation is skipped for simplicity
	result, err := h.DB.Exec(
		"UPDATE items SET "+
			"`title` = ?"+
			",`description` = ?"+
			",`updated` = ?"+
			"WHERE id = ?",
		r.FormValue("title"),
		r.FormValue("description"),
		"rvasily",
		id,
	)
	__err_panic(err)

	affected, err := result.RowsAffected()
	__err_panic(err)

	fmt.Println("Update - RowsAffected", affected)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	__err_panic(err)

	result, err := h.DB.Exec(
		"DELETE FROM items WHERE id = ?",
		id,
	)
	__err_panic(err)

	affected, err := result.RowsAffected()
	__err_panic(err)

	fmt.Println("Delete - RowsAffected", affected)

	w.Header().Set("Content-Type", "application/json")
	resp := `{"affected": ` + strconv.Itoa(int(affected)) + `}`
	w.Write([]byte(resp))
}

func main() {
	// main settings for db
	dsn := "root:0dmen@tcp(localhost:3306)/go-web-services?"
	// coding
	dsn += "&charset=utf8"
	// turn off prepared statements
	// params will be put instantly
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)

	db.SetMaxOpenConns(10)

	err = db.Ping() // here will be first connection to db
	__err_panic(err)

	handlers := &Handler{
		DB:   db,
		Tmpl: template.Must(template.ParseGlob("./crud_templates/*")),
	}

	// authorization is skipped for simplicity
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")

	fmt.Println("starting server at :4000")
	http.ListenAndServe(":4000", r)
}

// не используйте такой код в прошакшене
// ошибка должна всегда явно обрабатываться
func __err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
