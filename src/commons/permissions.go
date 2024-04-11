package commons

import "golang.org/x/exp/slices"

var Permissions = []string{
	"PERM_ROLE_ADD",
	"PERM_ROLE_MOD",
	"PERM_ROLE_GET",
	"PERM_ROLE_DEL",
}

func ValidatePermissions(permissions []string) bool {
	if len(permissions) == 1 && permissions[0] == "*" {
		return true
	}
	for _, permission := range permissions {
		if !slices.Contains(Permissions, permission) {
			return false
		}
	}
	return true
}
