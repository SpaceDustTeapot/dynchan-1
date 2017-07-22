package main

type Specification struct {
	Port               uint16 `default:"8080"`
	DatabaseDialect    string `default:"sqlite3" split_words:"true"`
	DatabaseConnection string `default:"dynchan2.sqlite" split_words:"true"`
}
