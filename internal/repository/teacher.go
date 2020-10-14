package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
)

const traceNameTeacher = "teacher repository"

type Teacher struct {
	db        db.Database
	traceName string
}

func NewTeacher() *Teacher {
	return &Teacher{
		db:        Db,
		traceName: "teacher repository",
	}
}

func (t *Teacher) Create(i *tracer.Infos, email string) error {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO teacher(user_id) select id from users where email = ?  "
	_, err := t.db.Insert(i, q, email)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return err
	}
	return nil

}

func (t *Teacher) Edit(i *tracer.Infos, email string, formation string, specialization string, time int) error {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	q := "" +
		"UPDATE teacher t " +
		"    inner join users u on u.id = t.user_id " +
		"SET " +
		"    formation = ?, " +
		"    specialization = ?, " +
		"    working_time = ? " +
		"WHERE " +
		"      u.email = ?;"
	_, err := t.db.Update(i, q, formation, specialization, time, email)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return err
	}
	return nil
}

func (t *Teacher) GetTeacherByPhoneticName(i *tracer.Infos, name string, offset, limit int) ([]map[string]interface{}, error) {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"       u.first_name, " +
		"       u.last_name, " +
		"       u.nickname," +
		"		u.photo_path," +
		"		u.description, " +
		"       t.formation " +
		"FROM users u " +
		"INNER JOIN teacher t on u.id = t.user_id  " +
		"WHERE " +
		"      u.register = 1 AND " +
		"      (soundex(u.first_name) = soundex(?) OR " +
		"       soundex(u.last_name) = soundex(?)) OR " +
		"		? LIKE CONCAT('%', u.first_name, '%') OR " +
		"		? LIKE CONCAT('%', u.last_name, '%') " +
		"LIMIT ? " +
		"OFFSET ?"
	result, err := t.db.Fetch(i, q, name, name, name, name, limit, offset)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}
