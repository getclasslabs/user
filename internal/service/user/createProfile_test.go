package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_removeAccents(t *testing.T) {
	tests := []struct {
		name   string
		s   string
		want   string
	}{
		{
			"[removeAccent] Name with accent",
			"André.Gâspar21",
			"Andre.Gaspar21",
		},
		{
			"[removeAccent] Name without accent",
			"Andre.Gaspar21",
			"Andre.Gaspar21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{}
			got := p.removeAccents(tt.s)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestProfile_createNick(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
		Nickname  string
	}
	tests := []struct {
		name   string
		fields fields
		want string
	}{
		{
			"[createNick] Name 1",
			fields{
				FirstName: "André",
				LastName: "Gaspar",
			},
			"andre.gaspar",
		},
		{
			"[createNick] Name 2",
			fields{
				FirstName: "Matheus",
				LastName: "Melo",
			},
			"matheus.melo",
		},
		{
			"[createNick] Name 3",
			fields{
				FirstName: "VarÍlão",
				LastName: "Truta",
			},
			"varilao.truta",
		},
		{
			"[createNick] Name 4",
			fields{
				FirstName: "Ana Julia",
				LastName: "Gaspar",
			},
			"anajulia.gaspar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Profile{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Nickname:  tt.fields.Nickname,
			}
			p.createNick()
			fmt.Println(p.Nickname)
			assert.Contains(t, p.Nickname, tt.want)
		})
	}
}