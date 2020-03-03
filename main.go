//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	ap "github.com/moemoe89/practicing-solr-golang/api"
	conf "github.com/moemoe89/practicing-solr-golang/config"

	"github.com/moemoe89/practicing-solr-golang/routers"
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
