package main

import (
	"time"

	"github.com/couchbaselabs/gocbcolumnar"
)

func connecting() {
	connStr := "couchbases://..."
	username := "..."
	password := "..."

	// tag::connecting[]
	cluster, err := cbcolumnar.NewCluster(
		connStr,
		cbcolumnar.NewCredential(username, password),
		// The third parameter is optional.
		// This example sets the default server query timeout to 3 minutes,
		// that is the timeout value sent to the query server.
		cbcolumnar.NewClusterOptions().SetTimeoutOptions(
			cbcolumnar.NewTimeoutOptions().SetServerQueryTimeout(3*time.Minute),
		),
	)
	handleErr(err)
	// end::connecting[]

	err = cluster.Close()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
