// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import pb "corp/fif/inte/customers/internal/core/repository/find_customer/proto"
import mock "github.com/stretchr/testify/mock"

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// FindCustomer provides a mock function with given fields: documentNumber, country, documentType
func (_m *CustomerRepository) FindCustomer(documentNumber string, country string, documentType string) (*pb.Customer, error) {
	ret := _m.Called(documentNumber, country, documentType)

	var r0 *pb.Customer
	if rf, ok := ret.Get(0).(func(string, string, string) *pb.Customer); ok {
		r0 = rf(documentNumber, country, documentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(documentNumber, country, documentType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}