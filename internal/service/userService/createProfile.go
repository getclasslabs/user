package userService

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/domains"
	"github.com/getclasslabs/user/internal/repository"
	"golang.org/x/text/runes"
	"math/rand"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Profile struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	BirthDate string `json:"birthDate,omitempty"`
	Gender    int    `json:"gender,omitempty"`
	Register  int    `json:"register,omitempty"` //0 to student an 1 to teacher
	Nickname  string `json:"nickname,omitempty"`
}

func (p *Profile) Do(i *tracer.Infos, email string) error {
	i.TraceIt("creating profile")
	defer i.Span.Finish()

	p.createNick()

	uRepo := repository.NewUser()

	err := uRepo.SaveProfile(i, email, p.Register, p.Gender, p.FirstName, p.LastName, p.BirthDate, p.Nickname)
	if err != nil {
		i.LogError(err)
		return err
	}

	if p.Register == domains.StudentRegister {
		s := repository.NewStudent()
		err = s.Create(i, email)
		if err != nil {
			i.LogError(err)
			return err
		}
		return nil
	}

	t := repository.NewTeacher()
	err = t.Create(i, email)
	if err != nil {
		i.LogError(err)
		return err
	}
	return nil
}

func (p *Profile) createNick() {
	fName := strings.Replace(p.FirstName, " ", "", -1)
	lName := strings.Replace(p.LastName, " ", "", -1)

	p.Nickname = strings.ToLower(p.removeAccents(fName + "." + lName + strconv.Itoa(rand.Intn(89)+10)))
}

func (p *Profile) removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}
