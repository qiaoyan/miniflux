// Copyright 2018 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.


package ui // import "miniflux.app/ui"

import (
	"net/http"

	"miniflux.app/v2/internal/http/response/json"
)

func (h *handler) showAppleAppSiteAssociation(w http.ResponseWriter, r *http.Request) {
	type siteIdentifiers struct {
		Apps []string `json:"apps"`
	}

	type credentials struct {
		Webcredentials            siteIdentifiers            `json:"webcredentials"`
	}

	// siteIdentifiers := &siteIdentifiers{
	// 	Apps: []string{
	// 		"V6CP2Z4H8Q.com.NewMobileWay.Geed",
	// 	},
	// }

	result := &credentials{
		Webcredentials: siteIdentifiers{
			Apps: []string{
				"V6CP2Z4H8Q.com.NewMobileWay.Geed",
			},
		},
	}

	json.OK(w, r, result)
}
