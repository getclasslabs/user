package review

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
)

func Get(i *tracer.Infos, teacherId string) (map[string]interface{}, error) {
	i.TraceIt("getting reviews")
	defer i.Span.Finish()

	rRepo := repository.NewReview()

	reviews, err := rRepo.GetReviewsByTeacherId(i, teacherId)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	average, err := rRepo.GetReviewAverageByTeacherId(i, teacherId)
	if err != nil {
		i.LogError(err)
		return nil, err
	}

	return map[string]interface{}{
		"average": average["average"],
		"reviews": reviews,
	}, nil
}
