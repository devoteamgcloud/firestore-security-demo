package entity

import (
	"fmt"

	"github.com/kaan-devoteam/one-click-deploy-demo/log"
)

type DepartmentCode int

const (
	DepartmentInvalid = iota
	DepartmentData
	DepartmentML
	DepartmentInfra
)

var (
	departmentNames = [...]string{
		"",
		"data",
		"ml",
		"infra",
	}
	departmentIndexes = map[string]DepartmentCode{
		departmentNames[DepartmentInvalid]: DepartmentInvalid,
		departmentNames[DepartmentData]:    DepartmentData,
		departmentNames[DepartmentML]:      DepartmentML,
		departmentNames[DepartmentInfra]:   DepartmentInfra,
	}
)

func DepartmentCodeFromString(name string) DepartmentCode {
	departmentCode, exists := departmentIndexes[name]
	if !exists {
		log.Warning(fmt.Sprintf("%s is not a valid department", name))
	}
	return departmentCode
}

func (d DepartmentCode) String() string {
	return departmentNames[d]
}
