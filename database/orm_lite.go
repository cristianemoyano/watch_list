package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "./foo.db")
	CheckDBErr(err)
	return db
}

func CheckDBErr(err error) {
	if err != nil {
		panic(err)
	}
}

func RunMigration() {
	db := GetDBConnection()
	stmt, _ := db.Prepare(
		`CREATE TABLE IF NOT EXISTS movies (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			type TEXT NOT NULL
		);`)
	stmt.Exec()

	stmt, _ = db.Prepare("INSERT INTO movies(name, type) values(?,?)")

	stmt.Exec("Monster Inc.", "pelicula")

	stmt.Exec("Friends", "serie")

	stmt.Exec("Todo sobre el asado", "documental")
}

func GetAllMovies() {
	db := GetDBConnection()
	// query
	rows, _ := db.Query("SELECT * FROM movies")

	var id int
	var name string
	var tipo string

	for rows.Next() {
		rows.Scan(&id, &name, &tipo)
		fmt.Printf("#%v - %v - type: %v\n", id, name, tipo)
	}

	rows.Close() //good habit to close
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	CheckDBErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	CheckDBErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	CheckDBErr(err)

	id, err := res.LastInsertId()
	CheckDBErr(err)

	fmt.Println(id)
	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckDBErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	CheckDBErr(err)

	affect, err := res.RowsAffected()
	CheckDBErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM userinfo")
	CheckDBErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		CheckDBErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	rows.Close() //good habit to close

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	CheckDBErr(err)

	res, err = stmt.Exec(id)
	CheckDBErr(err)

	affect, err = res.RowsAffected()
	CheckDBErr(err)

	fmt.Println(affect)

	db.Close()

}
