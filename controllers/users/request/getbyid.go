package request

import "github.com/iskandardevan/book-library/business/users"

type GetByID struct {
	Id uint `json:"id"`
}

func ToDomainById(getbyid GetByID) users.Domain {
	return users.Domain{
		Id: getbyid.Id,
	}
}