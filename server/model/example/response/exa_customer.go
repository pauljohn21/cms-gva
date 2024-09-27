package response

import "github.com/pauljohn21/cms-gva/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
