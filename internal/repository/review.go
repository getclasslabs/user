package repository

import (
	"github.com/getclasslabs/go-tools/pkg/db"
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/customerror"
)

const traceReview = "review repository"

type Review struct {
	db        db.Database
	traceName string
}

func NewReview() *Review {
	return &Review{
		db:        Db,
		traceName: "review repository",
	}
}

func (r *Review) SaveReview(i *tracer.Infos, comment string, value float64, student, teacher int) error {
	i.TraceIt(r.traceName)
	defer i.Span.Finish()

	q := "INSERT INTO reviews(comment, value, student_id, teacher_id) VALUES(?, ?, ?, ?) "
	_, err := r.db.Insert(i, q, comment, value, student, teacher)

	if err != nil {
		err := customerror.NewDbError(r, q, err)
		i.LogError(err)
		return err
	}
	return nil
}

func (r *Review) GetReviewsByTeacherId(i *tracer.Infos, teacherId string) ([]map[string]interface{}, error) {
	i.TraceIt(r.traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"       r.comment, " +
		"       r.value," +
		"		u.first_name," +
		"		u.last_name " +
		"FROM reviews r " +
		"INNER JOIN students s ON r.student_id = s.id " +
		"INNER JOIN users u ON s.user_id = u.id " +
		"WHERE " +
		"      r.teacher_id = ? " +
		"ORDER BY r.id DESC "
	result, err := r.db.Fetch(i, q, teacherId)

	if err != nil {
		err := customerror.NewDbError(r, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}

func (r *Review) GetReviewAverageByTeacherId(i *tracer.Infos, teacherId string) (map[string]interface{}, error) {
	i.TraceIt(traceName)
	defer i.Span.Finish()

	q := "SELECT " +
		"	ROUND(AVG(r.value), 2) as average " +
		"FROM  reviews r " +
		"WHERE " +
		"	r.teacher_id = ? " +
		"GROUP BY r.teacher_id"

	result, err := r.db.Get(i, q, teacherId)

	if err != nil {
		err := customerror.NewDbError(r, q, err)
		i.LogError(err)
		return nil, err
	}
	return result, nil
}