package user

import (
	"github.com/getclasslabs/go-tools/pkg/tracer"
	"github.com/getclasslabs/user/internal/repository"
	"golang.org/x/text/runes"
	"math/rand"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)


const (
	student = 0
	teacher = 1
)

type Profile struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	Gender int `json:"gender"`
	Register int `json:"register"` //0 to student an 1 to teacher
	Nickname string
}

func (p *Profile) Do(i *tracer.Infos, email string) error {
	i.TraceIt("creating profile")
	defer i.Span.Finish()

	p.createNick()

	uRepo := repository.NewUser()

	err := uRepo.SaveProfile(i, email, p.Register, p.Gender, p.FirstName, p.LastName, p.BirthDate, p.Nickname)
	if err != nil{
		i.LogError(err)
		return err
	}

	if p.Register == student {
		s := repository.NewStudent()
		err = s.Create(i, email)
		if err != nil{
			i.LogError(err)
			return err
		}
		return nil
	}

	t := repository.NewTeacher()
	err = t.Create(i, email)
	if err != nil{
		i.LogError(err)
		return err
	}
	return nil
}

func (p *Profile) createNick(){
	fName := strings.Replace(p.FirstName, " ", "", -1)
	lName := strings.Replace(p.LastName, " ", "", -1)

	p.Nickname = strings.ToLower(p.removeAccents(fName + "." + lName + strconv.Itoa(rand.Intn(89) + 10)))
}

func (p *Profile) removeAccents(s string) string  {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}