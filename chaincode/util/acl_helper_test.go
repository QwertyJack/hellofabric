//
// cert_checker_test.go
// Copyright (C) 2018 jack <jack@HP-WorkStation>
//
// Distributed under terms of the MIT license.
//

package util

import (
	"testing"
)

func TestIsAdmin(t *testing.T) {
	data := map[string]bool{
		"admin":     true,
		"Admin@org": true,
		"jack":      false,
		"adm":       false,
		"adMIN":     false,
		"Admin":     false,
	}

	for k, v := range data {
		if IsAdmin(k) != v {
			t.Error(k)
		}
	}
}
