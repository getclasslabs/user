package edit

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
	"github.com/getclasslabs/user/internal/service/user"
)

type Edit struct {
	user.Profile
	Twitter string `json:"twitter,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Description string `json:"description,omitempty"`
	Telephone string `json:"telephone,omitempty"`
	Address string `json:"address,omitempty"`

	//teacher
	Formation string `json:"formation,omitempty"`
	Specialization string `json:"specialization,omitempty"`
	WorkingTime int `json:"workingTime,omitempty"`
}

func (e *Edit) Do(i *tracer.Infos, email string) error {
	i.TraceIt("editing profile service")
	defer i.Span.Finish()

	if err := e.editCommonInfo(i, email); err != nil {
		i.LogError(err)
		return err
	}

	if err := e.editTeacherInfo(i, email); err != nil {
		i.LogError(err)
		return err
	}

	return nil
}


func (e *Edit) editCommonInfo(i *tracer.Infos, email string) error {
	uRepo := repository.NewUser()

	err := uRepo.Edit(i,
		email,
		e.Nickname,
		string(e.Gender),
		e.FirstName,
		e.LastName,
		e.BirthDate,
		e.Twitter,
		e.Facebook,
		e.Instagram,
		e.Description,
		e.Telephone,
		e.Address)
	if err != nil {
		i.LogError(err)
		return err
	}

	return nil
}


func (e *Edit) editTeacherInfo(i *tracer.Infos, email string) error {
	tRepo := repository.NewTeacher()

	err := tRepo.Edit(i,
		email,
		e.Formation,
		e.Specialization,
		e.WorkingTime)

	if err != nil {
		i.LogError(err)
		return err
	}

	return nil
}
