package cassango

import (
	"github.com/gocql/gocql"
)

type Cluster struct {
	Hosts    []string
	Keyspace string
	Username string
	Password string
}

type Session struct {
	session   *gocql.Session
	Error     error
	Debug     bool
	LogFormat string
}

func (c *Cluster) Open() (Session, error) {
	cluster := gocql.NewCluster(c.Hosts...)
	cluster.Keyspace = c.Keyspace

	session, err := cluster.CreateSession()

	//if err == nil {
	//	defer session.Close()
	//}

	s := Session{
		session: session,
	}

	return s, err
}
