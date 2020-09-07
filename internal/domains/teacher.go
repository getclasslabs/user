package domains

import "github.com/getclasslabs/user/internal/repository"

type Teacher struct {
	Domain
	Formation      string
	Specialization string
	WorkingTime    int
}

func (t *Teacher) Edit() error {
	tRepo := repository.NewTeacher()
	err := tRepo.Edit(t.Tracer, t.Email, t.Formation, t.Specialization, t.WorkingTime)
	if err != nil {
		t.Tracer.LogError(err)
		return err
	}
	return nil
}