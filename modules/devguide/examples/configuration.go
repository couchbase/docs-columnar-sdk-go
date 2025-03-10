package main

import (
	"time"

	"github.com/couchbase/gocbcolumnar"
)

func configuration() {
	// #tag::configuration[]
	cluster, err := cbcolumnar.NewCluster(
		connStr,
		cbcolumnar.NewCredential(username, password),
		cbcolumnar.NewClusterOptions().
			SetTimeoutOptions(
				cbcolumnar.NewTimeoutOptions().
					SetConnectTimeout(30*time.Second).
					SetQueryTimeout(2*time.Minute),
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
