package search

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/config"
	"github.com/getclasslabs/user/internal/repository"
	"strconv"
)

func Teacher(i *tracer.Infos, name, page string) ([]map[string]interface{}, error) {
	tRepo := repository.NewTeacher()

	limit := config.Config.SearchLimit

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		pageNumber = 1
	}
	offset := (pageNumber - 1) * limit

	teachers, err := tRepo.GetTeacherByPhoneticName(i, name, offset, limit)
	if err != nil {
		return nil, err
	}

	return teachers, err
}
