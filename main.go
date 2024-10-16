package main

import (
	"cpoc/dbcrawler/config"
	"cpoc/dbcrawler/trace"
)

func main() {

	conf := config.NewConfig("dbcrawler.conf")
	trace.SearchPidList(conf)

}
