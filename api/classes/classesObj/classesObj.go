package classesObj

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Class struct {
	ClassId   string `json:"classId"`
	ClassName string `json:"className"`
}

type Result struct {
	Result bool    `json:"result"`
	Body   []Class `json:"body"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:F_2324@a@tcp(172.28.0.2:3306)/pastQuestion")
	if err != nil {
		panic(err)
	}
}

func (class *Class) AddClass() (result bool) {
	num := 0
	statement := `SELECT COUNT(*) FROM classes WHERE classId=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}
	err = stmt.QueryRow(class.ClassId).Scan(&num)

	if num == 0 {
		statement = `INSERT INTO classes (classId, className) VALUES (?, ?)`
		stmt, err = db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		_, err = stmt.Exec(class.ClassId, class.ClassName)
		if err != nil {
			result = false
			return
		}
		result = true
		return
	}
	result = false
	return
}

func (class *Class) ModifyClass() (result bool) {
	num := 0
	statement := "SELECT COUNT(*) FROM classes WHERE classId=?"
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Printf("%v", err)
		result = false
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(class.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}

	if num != 0 {
		statement = "UPDATE classes SET className=? where classId=?"
		stmt, err = db.Prepare(statement)
		if err != nil {
			result = false
			return
		}

		_, err = stmt.Exec(class.ClassName, class.ClassId)
		if err != nil {
			result = false
			return
		}

		result = true
		return
	}
	result = false
	return
}

func (class *Class) DeleteClass() (result bool) {
	num := 0
	statement := "SELECT COUNT(*) FROM classes WHERE classId=?"
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		result = false
		return
	}

	err = stmt.QueryRow(class.ClassId).Scan(&num)
	if err != nil {
		result = false
		return
	}

	if num == 1 {
		statement = "DELETE FROM classes WHERE classId=?"
		stmt, err := db.Prepare(statement)
		defer stmt.Close()
		if err != nil {
			result = false
			return
		}
		_, err = stmt.Exec(class.ClassId)
		if err != nil {
			result = false
			return
		}
		result = true
		return
	}
	result = false
	return
}

func GetClass(className string) (classes []Class, ok bool) {
	statement := `SELECT classId, className
					FROM classes
					WHERE className=?`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return
	}

	rows, _ := stmt.Query(className)
	for rows.Next() {
		class := new(Class)
		err = rows.Scan(&class.ClassId, &class.ClassName)
		if err != nil {
			return
		}
		classes = append(classes, *class)
	}
	ok = true
	return
}

func GetHomeClass() (classes []Class, ok bool) {

	statement := `SELECT classId, className
					FROM classes
					ORDER BY classId
					LIMIT 10`
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		fmt.Printf("%v\n\n\n\n", err)
		return
	}
	rows, _ := stmt.Query()
	for rows.Next() {
		class := new(Class)
		err = rows.Scan(&class.ClassId, &class.ClassName)
		if err != nil {
			return
		}
		classes = append(classes, *class)
	}
	ok = true
	return
}
