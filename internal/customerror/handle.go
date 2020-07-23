package customerror

import "net/http"

func HandleStatus(err CustomError) int {
	switch err.GetErrType() {
	case Default:
		return http.StatusInternalServerError
	case Request:
		return err.GetExtraInfo()["status_code"].(int)
	case Parsing:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
