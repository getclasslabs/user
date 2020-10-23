package review

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
)

type Create struct {
	Comment    	string `json:"comment"`
	Value	  	float64 `json:"value"`
	Student	   	int `json:"student"`
	Teacher		int `json:"teacher"`
}

func (r *Create) Create(i *tracer.Infos) error {
	i.TraceIt("review")

	rRepo := repository.NewReview()
	err := rRepo.SaveReview(i, r.Comment, r.Value, r.Student, r.Teacher)
	if err != nil {
		return err
	}

	return nil
}
