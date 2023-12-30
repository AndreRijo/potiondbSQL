package sql

import (
	"strings"
)

//TODO: Support special character % for contains: a few TPC-H queries use this.
//Also missing the "in"

type Function interface {
	GetFunType() FunType //Mostly just to identify
	GetFunName() string
}

type FunType int

type StartsWithFun struct{}
type EndsWithFun struct{}
type ContainsFun struct{}

func (f StartsWithFun) GetFunType() FunType { return STARTS_WITH_FUN }
func (f EndsWithFun) GetFunType() FunType   { return ENDS_WITH_FUN }
func (f ContainsFun) GetFunType() FunType   { return CONTAINS_FUN }
func (f StartsWithFun) GetFunName() string  { return "startsWith" }
func (f EndsWithFun) GetFunName() string    { return "endsWith" }
func (f ContainsFun) GetFunName() string    { return "contains" }

func (f StartsWithFun) ApplyFun(myString, start string) bool {
	return strings.HasPrefix(myString, start)
}

func (f EndsWithFun) ApplyFun(myString, end string) bool {
	return strings.HasSuffix(myString, end)
}

func (f ContainsFun) ApplyFun(myString, part string) bool {
	return strings.Contains(myString, part)
}

func MakeFunFromName(funcName string) Function {
	funcName = strings.ToLower(funcName)
	if funcName == "startswith" || funcName == "hasprefix" {
		return StartsWithFun{}
	}
	if funcName == "endswith" || funcName == "hassuffix" {
		return EndsWithFun{}
	}
	if funcName == "contains" || funcName == "has" {
		return ContainsFun{}
	}
	return nil
}

const (
	STARTS_WITH_FUN = 1
	ENDS_WITH_FUN   = 2
	CONTAINS_FUN    = 3
)
