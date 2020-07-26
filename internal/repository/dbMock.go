package repository

import (
	"database/sql"
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
)

type Mock struct{
	
}

func (m Mock) Connect(config db.Config) {
	panic("implement me")
}

func (m Mock) Insert(infos *tracer.Infos, s string, i ...interface{}) (sql.Result, error) {
	panic("implement me")
}

func (m Mock) Update(infos *tracer.Infos, s string, i ...interface{}) (sql.Result, error) {
	panic("implement me")
}

func (m Mock) Get(infos *tracer.Infos, s string, i ...interface{}) (map[string]interface{}, error) {
	panic("implement me")
}

func (m Mock) Fetch(infos *tracer.Infos, s string, i ...interface{}) ([]map[string]interface{}, error) {
	panic("implement me")
}

