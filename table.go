package cassango

import (
	"fmt"

	"github.com/fatih/structs"
)

func (s *Session) CreateTable(m interface{}) *Session {
	if s.Error != nil {
		return s
	}

	i := structs.New(m)
	for _, f := range i.Fields() {
		fmt.Println(f.Name())
		fmt.Println(f.Kind())
		fmt.Println(f.Tag("cassango"))
	}

	query := "CREATE TABLE users (email TEXT, name TEXT, PRIMARY KEY (email));"

	return s.Excute(query)
}

func (s *Session) DropTable(m interface{}) *Session {
	if s.Error != nil {
		return s
	}

	tn := getTableName(m)
	query := "DROP TABLE IF EXISTS " + tn + ";"

	return s.Excute(query)
}

func getTableName(m interface{}) string {
	tn := fmt.Sprintf("%v", callMethod(m, "TableName"))
	if tn != "" {
		return tn
	}

	// TODO snake
	return structs.Name(m)
}

var columnTypes = map[string]string{
	"string": "TEXT",
}
