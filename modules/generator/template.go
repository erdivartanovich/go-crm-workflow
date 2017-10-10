package generator

var (
	serviceTempl = `package %s
type %s struct {

}

type SearchAdapter struct {

}

func ({{owner}} {{ownerType}}) Browse(adapter *SearchAdapter) ([]{{entityType}}, error) {
    return
}

func ({{owner}} {{ownerType}}) Read({{entity}} {{entityType}}) (%s, error) {

}

func ({{owner}} {{ownerType}}) Edit() (%s, error) {

}

func ({{owner}} {{ownerType}}) Add() (%s, error) {

}

func ({{owner}} {{ownerType}}) Delete() (%s, error) {

}
`

	entityTempl = `package %s

import (
    "time"
)

type %s struct {
    ID          uint64    %s
    UserID      uint64    %s
    Name        string    %s
    CreatedAt   time.Time %s
	UpdatedAt   time.Time %s
	DeletedAt   time.Time
}
`
	repositoryTempl = `package %s
`
)
