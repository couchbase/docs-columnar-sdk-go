package main

import (
	"context"
	"fmt"

	"github.com/couchbase/gocbcolumnar"
)

func queries() {
	cluster, err := cbcolumnar.NewCluster(
		connStr,
		cbcolumnar.NewCredential(username, password),
	)
	handleErr(err)

	scopeLevelQuery(context.Background(), cluster)
	clusterLevelQuery(context.Background(), cluster)
	positionalParamQuery(context.Background(), cluster)
	namedParamQuery(context.Background(), cluster)
	bufferResults(context.Background(), cluster)

	err = cluster.Close()
	handleErr(err)
}

func scopeLevelQuery(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::scopeLevelQuery[]
	scope := cluster.Database("my_database").Scope("my_scope")
	result, err := scope.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::scopeLevelQuery[]
}

func clusterLevelQuery(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::clusterLevelQuery[]
	result, err := cluster.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err = row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}
	// end::clusterLevelQuery[]
}

func positionalParamQuery(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::positionalParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select ?=1",
		cbcolumnar.NewQueryOptions().SetPositionalParameters([]interface{}{1}),
	)
	handleErr(err)
	// end::positionalParamQuery[]

	handleResult(result)
}

func namedParamQuery(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::namedParamQuery[]
	result, err := cluster.ExecuteQuery(
		ctx,
		"select $foo=1",
		cbcolumnar.NewQueryOptions().SetNamedParameters(map[string]interface{}{"foo": 1}),
	)
	handleErr(err)
	// end::namedParamQuery[]

	handleResult(result)
}

func handleResult(result *cbcolumnar.QueryResult) {
	// tag::handleResults[]
	for row := result.NextRow(); row != nil; row = result.NextRow() {
		var content map[string]int

		err := row.ContentAs(&content)
		handleErr(err)

		fmt.Printf("Got row content: %v", content)
	}

	if err := result.Err(); err != nil {
		handleErr(err)
	}
	// end::handleResults[]
}

func metadata(result *cbcolumnar.QueryResult) {
	// tag::metadata[]
	meta, err := result.MetaData()
	handleErr(err)

	fmt.Printf("Got meta: %v", meta)
	// end::metadata[]
}

func bufferResults(ctx context.Context, cluster *cbcolumnar.Cluster) {
	// tag::bufferResults[]
	result, err := cluster.ExecuteQuery(ctx, "select 1")
	handleErr(err)

	rows, meta, err := cbcolumnar.BufferQueryResult[map[string]int](result)
	handleErr(err)

	for _, row := range rows {
		fmt.Printf("Got row content: %v", row)
	}

	fmt.Printf("Got meta: %v", meta)
	// end::bufferResults[]
}
