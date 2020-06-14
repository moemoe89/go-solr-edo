//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"github.com/moemoe89/go-solr-edo/api/api_struct/form"
	"github.com/moemoe89/go-solr-edo/api/api_struct/model"
	cons "github.com/moemoe89/go-solr-edo/constant"

	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/sirupsen/logrus"
)

// ctrl struct represent the delivery for controller
type ctrl struct {
	log *logrus.Entry
	svc Service
}

// NewCtrl will create an object that represent the ctrl struct
func NewCtrl(log *logrus.Entry, svc Service) *ctrl {
	return &ctrl{log, svc}
}

func (ct *ctrl) Create(req form.UserForm, r render.Render) {
	errs := req.Validate()
	if len(errs) > 0 {
		r.JSON(http.StatusBadRequest, model.NewGenericResponse(http.StatusBadRequest, cons.ERR, errs, nil))
		return
	}

	status, err := ct.svc.Create(req)
	if err != nil {
		r.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusCreated, model.NewGenericResponse(http.StatusCreated, cons.OK, []string{"Data has been saved"}, req))
}

func (ct *ctrl) Select(req *http.Request, r render.Render) {
	key := req.URL.Query().Get("key")
	value := req.URL.Query().Get("value")
	data, status, err := ct.svc.Select(key, value)
	if err != nil {
		r.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been retrieved"}, data))
}

func (ct *ctrl) Delete(params martini.Params, r render.Render) {
	id := params["id"]
	status, err := ct.svc.Delete(id)
	if err != nil {
		r.JSON(status, model.NewGenericResponse(status, cons.ERR, []string{err.Error()}, nil))
		return
	}
	r.JSON(http.StatusOK, model.NewGenericResponse(http.StatusOK, cons.OK, []string{"Data has been deleted"}, nil))
}
