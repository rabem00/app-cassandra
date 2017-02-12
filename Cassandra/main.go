package Cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Session holds our connection to Cassandra
var Session *gocql.Session

func init() {
	var err error

	cluster := gocql.NewCluster("192.168.64.6")
	cluster.ProtoVersion = 4
	cluster.CQLVersion = "3.0.0"
	cluster.Keyspace = "appdemo"
	cluster.Consistency = gocql.Quorum
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
