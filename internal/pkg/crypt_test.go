package pkg

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCrypt(t *testing.T) {
	tests := []struct {
		name    string
		pass    string
	}{
		{
			"[Crypt]: Compare 1",
			"testepass1",
		},
		{
			"[Crypt]: Compare 2",
			"Rjsh!@092B",
		},
		{
			"[Crypt]: Compare 2",
			"2345678912348763249172345hhasdfbgfdu",
		},
	}

	for _, tt := range tests{
		hashed, err := Crypt(tt.pass)
		assert.NoError(t, err)
		err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(tt.pass))
		assert.NoError(t, err)
	}
}