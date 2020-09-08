package search

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
)

func Teacher(i *tracer.Infos, name string) (map[string]interface{}, error) {
	tRepo := repository.NewTeacher()
	teachers, err := tRepo.GetTeacherByPhoneticName(i, name)
	if err != nil {
		return nil, err
	}

	return teachers, err
}
