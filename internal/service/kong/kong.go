package kong

import "net/http"

type Kong struct {

}

func CreateUser(){
	http.Post(main.config)
}