package service

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// )

// func TestMain(m *testing.M) {
// 	// os.Exit skips defer calls
// 	// so we need to call another function
// 	code, err := run(m)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	os.Exit(code)
// }

// func run(m *testing.M) (code int, err error) {

// 	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "a", "b", "c", "d"))
// 	if err != nil {
// 		return -1, fmt.Errorf("could not connect to database: %w", err)
// 	}

// 	db.Close()
// 	// truncates all test data after the tests are run
// 	// defer func() {
// 	//     for _, t := range string{"books", "authors"} {
// 	//         _, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
// 	//     }

// 	//     db.Close()
// 	// }()

// 	return m.Run(), nil
// }
