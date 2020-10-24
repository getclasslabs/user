package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/config"
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

func (t *Teacher) GetTeacherByPhoneticName(i *tracer.Infos, name string, page int) ([]map[string]interface{}, error) {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	limit := config.Config.SearchLimit
	offset := (page - 1) * limit

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
		"       soundex(u.last_name) = soundex(?) OR " +
		"		soundex(CONCAT(u.first_name, ' ', u.last_name)) = soundex(?)) " +
		"LIMIT ? " +
		"OFFSET ?"
	result, err := t.db.Fetch(i, q, name, name, name, limit, offset)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}


func (t *Teacher) GetNextPageTeacher(i *tracer.Infos, name string) (map[string]interface{}, error) {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"       count(u.id) as count " +
		"FROM users u " +
		"INNER JOIN teacher t on u.id = t.user_id  " +
		"WHERE " +
		"      u.register = 1 AND " +
		"      (soundex(u.first_name) = soundex(?) OR " +
		"       soundex(u.last_name) = soundex(?) OR " +
		"		soundex(CONCAT(u.first_name, ' ', u.last_name)) = soundex(?)) "
	result, err := t.db.Get(i, q, name, name, name)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}

func (t *Teacher) GetTeacherById(i *tracer.Infos, id int) (map[string]interface{}, error) {
	i.TraceIt(t.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	u.first_name, " +
		"	u.last_name, " +
		"	u.nickname," +
		"	u.photo_path," +
		"	u.description, " +
		"	t.formation " +
		"FROM users u " +
		"INNER JOIN teacher t on u.id = t.user_id  " +
		"WHERE " +
		"    t.id = ?"
	result, err := t.db.Get(i, q, id)

	if err != nil {
		err := customerror.NewDbError(t, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}