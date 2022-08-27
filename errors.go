package arvanvod

import (
	"errors"
	"net/http"
)

var (
	ErrorUnknown                  = errors.New("unknown error")
	ErrorUnauthenticated          = errors.New("unauthenticated")
	ErrorEntityOrSubDomainIsTaken = errors.New("unproccesable entity")
	ErrorNotFound                 = errors.New("not found")
)

func getErrorByStatus(status int) error {
	err := ErrorUnknown
	switch status {
	case http.StatusUnprocessableEntity:
		err = ErrorEntityOrSubDomainIsTaken
	case http.StatusUnauthorized:
		err = ErrorUnauthenticated
	case http.StatusNotFound:
		err = ErrorNotFound
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		err = nil
	}

	return err
}
