// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tailscale/pkgsitelib/pkg/config"
)

func TestErrorReporting(t *testing.T) {
	tests := []struct {
		code         int
		bypassHeader string
		wantReports  int
	}{
		{500, "", 1},
		{404, "", 0},
		{200, "", 0},
		{503, "", 0},
		{550, "", 0},
		{500, "true", 0}, // set bypass header
		{500, "false", 1},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d,%s", test.code, test.bypassHeader), func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if test.bypassHeader != "" {
					w.Header().Set(config.BypassErrorReportingHeader, test.bypassHeader)
				}
				w.WriteHeader(test.code)
			})
			fr := &fakeReporter{}
			mw := ErrorReporting(fr)
			ts := httptest.NewServer(mw(handler))
			resp, err := http.Get(ts.URL)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if got := fr.reports; got != test.wantReports {
				t.Errorf("Got %d reports, want %d", got, test.wantReports)
			}
		})
	}
}

type fakeReporter struct {
	reports int
}

func (f *fakeReporter) Report(error, *http.Request, []byte) {
	f.reports++
}
