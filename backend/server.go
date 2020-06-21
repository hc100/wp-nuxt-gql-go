package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph/generated"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const defaultDataSource = "foo:bar@tcp(localhost:3306)/baz?charset=utf8&parseTime=True&loc=Local"
const defaultPort = "8080"

func main() {
	dataSource := os.Getenv("WP_DATASOURCE")
	if dataSource == "" {
		dataSource = defaultDataSource
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	defer func() {
		if db != nil {
			if err := db.Close(); err != nil {
				panic(err)
			}
		}
	}()
	db.LogMode(true)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
