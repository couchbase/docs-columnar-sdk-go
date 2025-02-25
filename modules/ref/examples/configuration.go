package main

import (
	"time"

	cbcolumnar "github.com/couchbaselabs/gocbcolumnar"
)

func main() {
	connStr := "couchbases://..."
	username := "..."
	password := "..."

	// #tag::configuration[]
	cluster, err := cbcolumnar.NewCluster(
		connStr,
		cbcolumnar.NewCredential(username, password),
		cbcolumnar.NewClusterOptions().
			SetTimeoutOptions(
				cbcolumnar.NewTimeoutOptions().
					SetConnectTimeout(30*time.Second).
					SetServerQueryTimeout(2*time.Minute),
			).
			SetSecurityOptions(cbcolumnar.NewSecurityOptions().
				SetCipherSuites([]string{"MY_APPROVED_CIPHER_SUITE"}),
			),
	)
	handleErr(err)
	// #end::configuration[]

	err = cluster.Close()
	handleErr(err)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
