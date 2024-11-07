package memory

import (
	"fmt"
	"github.com/duongbm/go-ddd-example/aggregate"
	"github.com/duongbm/go-ddd-example/domain/customer"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (m *MemoryRepository) Get(uuid uuid.UUID) (aggregate.Customer, error) {
	if cus, ok := m.customers[uuid]; ok {
		return cus, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (m *MemoryRepository) Add(c aggregate.Customer) error {
	if m.customers == nil {
		m.Lock()
		m.customers = make(map[uuid.UUID]aggregate.Customer)
		m.Unlock()
	}
	// make sure customer is already in repository
	if _, ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}

func (m *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	m.Lock()
	m.customers[c.GetID()] = c
	m.Unlock()
	return nil
}
