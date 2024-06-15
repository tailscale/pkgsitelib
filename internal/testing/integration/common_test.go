// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package integration

import (
	"github.com/tailscale/pkgsitelib/internal/postgres"
	"github.com/tailscale/pkgsitelib/internal/proxy/proxytest"
)

var (
	testDB      *postgres.DB
	testModules []*proxytest.Module
)
