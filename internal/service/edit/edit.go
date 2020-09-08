package edit

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/domains"
	"github.com/getclasslabs/user/internal/service/userService"
)

type Edit struct {
	userService.Profile
	Twitter     string `json:"twitter,omitempty"`
	Facebook    string `json:"facebook,omitempty"`
	Instagram   string `json:"instagram,omitempty"`
	Description string `json:"description,omitempty"`
	Telephone   string `json:"telephone,omitempty"`
	Address     string `json:"address,omitempty"`

	//teacher
	Formation      string `json:"formation,omitempty"`
	Specialization string `json:"specialization,omitempty"`
	WorkingTime    int    `json:"workingTime,omitempty"`
}

func (e *Edit) Do(i *tracer.Infos, email string) error {
	i.TraceIt("editing profile service")
	defer i.Span.Finish()

	if err := e.editCommonInfo(i, email); err != nil {
		return err
	}

	if err := e.editTeacherInfo(i, email); err != nil {
		return err
	}

	return nil
}

func (e *Edit) editCommonInfo(i *tracer.Infos, email string) error {
	u := domains.User{
		Domain: domains.Domain{
			Tracer: i,
			Email:  email,
		},
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		BirthDate:   e.BirthDate,
		Gender:      e.Gender,
		Nickname:    e.Nickname,
		Twitter:     e.Twitter,
		Facebook:    e.Facebook,
		Instagram:   e.Instagram,
		Description: e.Description,
		Telephone:   e.Telephone,
		Address:     e.Address,
	}

	err := u.Edit()
	if err != nil {
		return err
	}
	return nil
}

func (e *Edit) editTeacherInfo(i *tracer.Infos, email string) error {
	t := domains.Teacher{
		Domain: domains.Domain{
			Tracer: i,
			Email:  email,
		},
		Formation:      e.Formation,
		Specialization: e.Specialization,
		WorkingTime:    e.WorkingTime,
	}

	err := t.Edit()
	if err != nil {
		return err
	}

	return nil
}
