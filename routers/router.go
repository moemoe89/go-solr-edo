//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	"github.com/moemoe89/practicing-solr-golang/api/api_struct/form"
	ap "github.com/moemoe89/practicing-solr-golang/api"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/sirupsen/logrus"
)

// GetRouter will create a variable that represent the martini.Martini
func GetRouter(log *logrus.Entry, svc ap.Service) *martini.Martini {
	m := martini.New()
	m.Use(martini.Recovery())
	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	ctrl := ap.NewCtrl(log, svc)

	r := martini.NewRouter()
	r.Get(`/`, ap.Ping)
	r.Get(`/ping`, ap.Ping)
	r.Post(`/create`, binding.Bind(form.UserForm{}), ctrl.Create)
	r.Get(`/select`, ctrl.Select)
	r.Delete(`/delete/:id`, ctrl.Delete)

	m.Action(r.Handle)

	return m
}
