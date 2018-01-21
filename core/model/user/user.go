package user

import (
	"database/sql"
	"fmt"

	"github.com/Dlacreme/httpd/back/qbuilder"
	"github.com/Dlacreme/httpd/back/wdb"
	"github.com/Dlacreme/httpd/back/werror"
)

var (
	table = "User"
)

type User struct {
	UserId   int    `db:"UserId"`
	Email    string `db:"Email"`
	Name     string `db:"Name"`
	Password string `db:"Password"`
	StatusId int8   `db:"StatusId"`
	TypeId   int8   `db:"TypeId"`
}

// FindById will load a user from his id
func FindById(db wdb.Connection, id int) (*User, *werror.Error) {
	res := User{}
	err := db.Get(&res, fmt.Sprintf(`
			SELECT UserId, Email, Name, StatusId, TypeId
			FROM %v 
			WHERE UserId = ?
			LIMIT 1
		`, table),
		id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Error while Find User By Id : %s", err.Error()))
	}

	return &res, nil
}

func FindByCredentials(db wdb.Connection, email string, passwd string) (*User, *werror.Error) {
	res := User{}

	err := db.Get(&res, fmt.Sprintf(`
				SELECT UserId, Email, Name, StatusId, TypeId
				FROM %v 
				WHERE Email = ? AND Password = ?
				LIMIT 1
			`, table),
		email, passwd)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, werror.New(500, fmt.Sprintf("Error while fetch user by cred : %s", err.Error()))
	}

	return &res, nil
}

func CreateMember(db wdb.Connection, email string, name string, password string) (int, *werror.Error) {
	exist, err := isExisting(db, email)

	if err != nil {
		return 0, err
	}

	if exist {
		return 0, werror.New(401, fmt.Sprintf("Email already existing"))
	}

	fields := []string{"Email", "Name", "Password", "TypeId"}
	q := qbuilder.InsertInto(table, fields)
	item := []string{qbuilder.FormatStr(email), qbuilder.FormatStr(name), qbuilder.FormatStr(password), "1"}
	q += qbuilder.AddInto(q, item)

	r, e := db.Exec(q)
	if e != nil {
		return 0, werror.New(500, fmt.Sprintf("Error: %s", e.Error()))
	}
	id, _ := r.LastInsertId()
	return int(id), nil
}

func CreateStaff(db wdb.Connection, email string, name string, password string) (int, *werror.Error) {
	exist, err := isExisting(db, email)

	if err != nil {
		return 0, err
	}

	if exist {
		return 0, werror.New(401, fmt.Sprintf("Email already existing"))
	}
	fields := []string{"Email", "Name", "Password", "TypeId"}
	q := qbuilder.InsertInto(table, fields)
	item := []string{qbuilder.FormatStr(email), qbuilder.FormatStr(name), qbuilder.FormatStr(password), "1"}
	q += qbuilder.AddInto(q, item)

	r, e := db.Exec(q)
	if e != nil {
		return 0, werror.New(500, fmt.Sprintf("Error: %s", e.Error()))
	}
	id, _ := r.LastInsertId()
	return int(id), nil
}

func CreateAdmin(db wdb.Connection, email string, name string, password string) (int, *werror.Error) {
	exist, err := isExisting(db, email)

	if err != nil {
		return 0, err
	}

	if exist {
		return 0, werror.New(401, fmt.Sprintf("Email already existing"))
	}
	fields := []string{"Email", "Name", "Password", "TypeId"}
	q := qbuilder.InsertInto(table, fields)
	item := []string{qbuilder.FormatStr(email), qbuilder.FormatStr(name), qbuilder.FormatStr(password), "1"}
	q += qbuilder.AddInto(q, item)

	r, e := db.Exec(q)
	if e != nil {
		return 0, werror.New(500, fmt.Sprintf("Error: %s", e.Error()))
	}
	id, _ := r.LastInsertId()
	return int(id), nil
}
