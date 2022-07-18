// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	entity "sizesure/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// SizesuresUseCase is an autogenerated mock type for the SizesuresUseCase type
type SizesuresUseCase struct {
	mock.Mock
}

// AddSizesures provides a mock function with given fields: _a0
func (_m *SizesuresUseCase) AddSizesures(_a0 entity.Seizure) (entity.SeizureResp, error) {
	ret := _m.Called(_a0)

	var r0 entity.SeizureResp
	if rf, ok := ret.Get(0).(func(entity.Seizure) entity.SeizureResp); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(entity.SeizureResp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Seizure) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
