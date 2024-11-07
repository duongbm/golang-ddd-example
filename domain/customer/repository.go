package customer

import (
	"errors"
	"github.com/duongbm/go-ddd-example/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed to update customer")
)

type CustomersRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
