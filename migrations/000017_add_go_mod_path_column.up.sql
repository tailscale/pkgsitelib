-- Copyright 2020 The Go Authors. All rights reserved.
-- Use of this source code is governed by a BSD-style
-- license that can be found in the LICENSE file.

BEGIN;

ALTER TABLE module_version_states ADD COLUMN go_mod_path text;

COMMENT ON COLUMN module_version_states.go_mod_path IS
'COLUMN go_mod_path holds the module path from the go.mod file.';

END;
