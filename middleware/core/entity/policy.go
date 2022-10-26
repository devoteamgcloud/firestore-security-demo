package entity

import (
	"fmt"

	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type PolicyCode int

const (
	PolicyInvalid = iota
	PolicyPrivate
	PolicyPublic
	PolicyRestricted
)

var (
	policyNames = [...]string{
		"",
		"private",
		"public",
		"restricted",
	}
	policyIndexes = map[string]PolicyCode{
		policyNames[PolicyInvalid]:    PolicyInvalid,
		policyNames[PolicyPrivate]:    PolicyPrivate,
		policyNames[PolicyPublic]:     PolicyPublic,
		policyNames[PolicyRestricted]: PolicyRestricted,
	}
)

func PolicyCodeFromString(name string) PolicyCode {
	policyCode, exists := policyIndexes[name]
	if !exists {
		log.Warning(fmt.Sprintf("%s is not a valid polict", name))
	}
	return policyCode
}

func (p PolicyCode) String() string {
	return policyNames[p]
}
