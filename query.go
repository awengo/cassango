package cassango

import (
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

func (s *Session) Excute(q string, bs ...interface{}) *Session {
	s.Error = s.session.Query(q, bs...).Exec()
	//for iter.Scan(&id, &text) {
	//	fmt.Println("Tweet:", id, text)
	//}
	//	if err := iter.Close(); err != nil {
	//		log.Fatal(err)
	//	}

	l := logger{
		format: s.LogFormat,
		Query:  q,
		Error:  s.Error,
	}
	l.printLog()

	return s
}

func (s *Session) Find(m interface{}) *Session {
	if s.Error != nil {
		return s
	}

	v := reflect.ValueOf(m)

	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Multiple
	if v.Kind() == reflect.Slice {
		fmt.Println(m)
	}

	return s
}

func (s *Session) Create(m interface{}) *Session {
	tn := getTableName(m)

	i := structs.New(m)

	query := "INSERT INTO " + tn + " (email, name) VALUES (?, ?);"
	binds := []interface{}{
		i.Field("Email").Value(),
		i.Field("Name").Value(),
	}

	return s.Excute(query, binds...)
}
