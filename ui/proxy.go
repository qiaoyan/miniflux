// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package ui // import "miniflux.app/ui"

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"miniflux.app/crypto"
	"miniflux.app/http/client"
	"miniflux.app/http/request"
	"miniflux.app/http/response"
	"miniflux.app/http/response/html"
)

func (h *handler) imageProxy(w http.ResponseWriter, r *http.Request) {
	// If we receive a "If-None-Match" header, we assume the image is already stored in browser cache.
	if r.Header.Get("If-None-Match") != "" {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	encodedURL := request.RouteStringParam(r, "encodedURL")
	if encodedURL == "" {
		html.BadRequest(w, r, errors.New("No URL provided"))
		return
	}

	decodedURL, err := base64.URLEncoding.DecodeString(encodedURL)
	if err != nil {
		html.BadRequest(w, r, errors.New("Unable to decode this URL"))
		return
	}

	clt := client.New(string(decodedURL))
	resp, err := clt.Get()
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	if resp.HasServerFailure() {
		html.NotFound(w, r)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	etag := crypto.HashFromBytes(body)

	response.New(w ,r).WithCaching(etag, 72*time.Hour, func(b *response.Builder) {
		b.WithHeader("Content-Type", resp.ContentType)
		b.WithBody(body)
		b.WithoutCompression()
		b.Write()
	})
}
