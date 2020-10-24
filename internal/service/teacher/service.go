package teacher

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
)

func GetTeacherById(i *tracer.Infos, id int) (map[string]interface{}, error){
	i.TraceIt("getting teacher")
	defer i.Span.Finish()

	tRepo := repository.NewTeacher()
	return tRepo.GetTeacherById(i, id)

}

