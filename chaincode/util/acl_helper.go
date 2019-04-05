//
// cert_checker.go
// Copyright (C) 2018 jack <jack@HP-WorkStation>
//
// Distributed under terms of the MIT license.
//

package util

// IsAdmin checks whether a username string is admin or not
func IsAdmin(cert string) bool {
	if cert == "admin" {
		return true
	} else if len(cert) < 6 {
		return false
	} else {
		return cert[:6] == "Admin@"
	}
}

// Contains check where list contains elem
func Contains(list []string, elem string) bool {
	for _, a := range list {
		if a == elem {
			return true
		}
	}
	return false
}
