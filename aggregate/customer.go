package aggregate

import (
	"errors"
	"github.com/duongbm/go-ddd-example/entity"
	"github.com/duongbm/go-ddd-example/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has an invalid name")
)

type Customer struct {
	// person is the root entity of customer
	// which mean person.ID is main id for customer
	person      *entity.Person
	products    []*entity.Item
	transaction []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate that name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
