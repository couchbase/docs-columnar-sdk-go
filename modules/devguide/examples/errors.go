package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/couchbase/gocbcolumnar"
)

func columanarError(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::columanarError[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select 1=1",
	)
	if err != nil {
		var columnarErr cbcolumnar.ColumnarError
		if errors.As(err, &columnarErr) {
			// Do something with this error.
		}

		// This error occurred out of a client-server interaction.
	}
	// end::columanarError[]

	handleResult(result)
}

func queryError(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::queryError[]
	handleQueryError := func(err error) {
		if err != nil {
			var queryErr cbcolumnar.QueryError
			if errors.As(err, &queryErr) {
				fmt.Printf("Error code: %d, error message: %s", queryErr.Code(), queryErr.Message())
				return
			}

			// This error isn't a result of query processing, possibly something like a connection error.
		}
	}

	result, err := cluster.ExecuteQuery(
		ctx,
		"selec 1=1", // Syntax error
	)
	handleQueryError(err)
	// end::queryError[]

	for row := result.NextRow(); row != nil; row = result.NextRow() {
	}

	err = result.Err()
	handleQueryError(err)

	handleResult(result)
}
