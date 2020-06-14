//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/go-solr-edo/api"
	conf "github.com/moemoe89/go-solr-edo/config"

	"github.com/moemoe89/go-solr-edo/routers"
)

func main() {
	client, err := conf.InitSolr()
	if err != nil {
		panic(err)
	}

	log := conf.InitLog()

	repo := ap.NewSolrRepository(client)
	svc := ap.NewService(log, repo)

	app := routers.GetRouter(log, svc)
	app.RunOnAddr(":" + conf.Configuration.Port)
}
