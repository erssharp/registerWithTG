package entities

import(
	"github.com/google/uuid"
)

type Employee struct {
	Id uuid.UUID `json:"id"`
	CompanyId uuid.UUID `json:"company_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Rating int `json:"rating"`
	RoleId uuid.UUID `json:"role_id"`
}
