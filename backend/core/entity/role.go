package entity

import (
	"fmt"

	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type RoleCode int

const (
	RoleInvalid = iota
	RoleSelf
	RoleGroup
	RoleGlobal
)

var (
	roleNames = [...]string{
		"",
		"self",
		"group",
		"global",
	}
	roleIndexes = map[string]RoleCode{
		roleNames[RoleInvalid]: RoleInvalid,
		roleNames[RoleSelf]:    RoleSelf,
		roleNames[RoleGroup]:   RoleGroup,
		roleNames[RoleGlobal]:  RoleGlobal,
	}
)

func RoleCodeFromString(name string) RoleCode {
	roleCode, exists := roleIndexes[name]
	if !exists {
		log.Warning(fmt.Sprintf("%s is not a valid role", name))
	}
	return roleCode
}

func (r RoleCode) String() string {
	return roleNames[r]
}
