package search

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/config"
	"github.com/getclasslabs/user/internal/repository"
	"strconv"
)

func Teacher(i *tracer.Infos, name, page string) (map[string]interface{}, error) {
	tRepo := repository.NewTeacher()
	limit := config.Config.SearchLimit

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 0 {
		pageNumber = 1
	}

	teachers, err := tRepo.GetTeacherByPhoneticName(i, name, pageNumber)
	if err != nil {
		return nil, err
	}

	next, err := tRepo.GetNextPageTeacher(i, name)
	if err != nil {
		return nil, err
	}

	hasNextCount := (pageNumber * limit) + 1
	var hasNext bool
	if len(next) > 0 && next["count"].(int64) >= int64(hasNextCount)  {
		hasNext = true
	} else {
		hasNext = false
	}

	return map[string]interface{}{
		"next": hasNext,
		"results": teachers,
	}, err
}
