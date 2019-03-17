// Copyright 2018 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package ui // import "miniflux.app/ui"

import (
	"miniflux.app/config"
	"miniflux.app/oauth2"
)

func getOAuth2Manager(cfg *config.Config) *oauth2.Manager {
	return oauth2.NewManager(
		cfg.OAuth2ClientID(),
		cfg.OAuth2ClientSecret(),
		cfg.OAuth2RedirectURL(),
	)
}
