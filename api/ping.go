//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package api

import (
	"net/http"
	"time"

	"github.com/martini-contrib/render"
)

var starTime = time.Now()

// Ping will handle the ping endpoint
func Ping(r render.Render) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := starTime.In(loc)

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
	}
	r.JSON(http.StatusOK, res)
}
