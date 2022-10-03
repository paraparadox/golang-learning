package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var (
	db *sql.DB
)

var loginFormTmpl = `
<html>
	<body>
	<form action="/login" method="post">
		Login: <input name="login"/>
		Password: <input name="password"/>
		<input type="submit" value="login"/>
	</form>
	</body>
</html>
`

func main() {
	dsn := "root:0dmen@tcp(localhost:3306)/go-web-services?&charset=utf8&interpolateParams=true"

	var err error
	db, err := sql.Open("mysql", dsn)
	PanicOnErr(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(loginFormTmpl))
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var (
			id          int
			login, body string
		)

		inputLogin := r.FormValue("login")
		body += fmt.Sprintln("inputLogin:", inputLogin)

		// The wrong way. Don't do like this!
		query := fmt.Sprintf("SELECT id, login FROM users WHERE login = '%s' LIMIT 1", inputLogin)

		body += fmt.Sprintln("Sprint query:", query)

		row := db.QueryRow(query)
		err := row.Scan(&id, &login)

		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Sprint case: Not found")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Sprint case: id:", id, "login", login)
		}

		// The right way. Use placeholders
		row = db.QueryRow("SELECT id, login FROM users WHERE login = ? LIMIT 1", inputLogin)
		err = row.Scan(&id, &login)
		if err == sql.ErrNoRows {
			body += fmt.Sprintln("Placeholders case: Not found")
		} else {
			PanicOnErr(err)
			body += fmt.Sprintln("Placeholders id:", id, "login:", login)
		}

		w.Write([]byte(body))
	})

	fmt.Println("starting server at :4000")
	http.ListenAndServe(":4000", nil)
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
