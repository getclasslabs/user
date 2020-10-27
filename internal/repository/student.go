package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
)

type Student struct {
	db        db.Database
	traceName string
}

func NewStudent() *Student {
	return &Student{
		db:        Db,
		traceName: "student repository",
	}
}

func (s *Student) Create(i *tracer.Infos, userEmail string) error {
	i.TraceIt(s.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO students(user_id) SELECT id from users where email = ? "
	_, err := s.db.Insert(i, q, userEmail)

	if err != nil {
		err := customerror.NewDbError(s, q, err)
		i.LogError(err)
		return err
	}
	return nil

}
