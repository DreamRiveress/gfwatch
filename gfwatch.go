package gfwatch

import (
)

/*
type DomainType byte
type gfWatchMap map[string]struct{}

const (
	DomainCensoredAs DomainType = 0  // Rule 0 - domain censored as is
	BaseDomainDS     DomainType = 1  // Rule 1 - blocked by regex: base_censored_domain.*
	BaseDomainS      DomainType = 2  // Rule 2 - blocked by regex: base_censored_domain*
	SDBaseDomain     DomainType = 3  // Rule 3 - blocked by regex: *.base_censored_domain
	SBaseDomain      DomainType = 4  // Rule 4 - blocked by regex: *base_censored_domain
	SDBaseDomainDS   DomainType = 5  // Rule 5 - blocked by regex: *.base_censored_domain.*
	SDBaseDomainS    DomainType = 6  // Rule 6 - blocked by regex: *.base_censored_domain*
	SBaseDomainDS    DomainType = 7  // Rule 7 - blocked by regex: *base_censored_domain.*
	SBaseDomainS     DomainType = 8  // Rule 8 - blocked by regex: *base_censored_domain*
	DomainTypeMax               = 9
)*/

type GfWatch struct {
	gfMap map[string]struct{}
}

func New() *GfWatch {
	return &GfWatch{
		gfMap: make(map[string]struct{}),
	}
}

func (gfw *GfWatch) IsForbidden(domain string) bool {
	_, found := gfw.gfMap[domain]
	if !found {
		return false
	}
	return true
}
