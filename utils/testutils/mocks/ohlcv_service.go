// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	mock "github.com/stretchr/testify/mock"

	types "github.com/Proofsuite/amp-matching-engine/types"
	"github.com/Proofsuite/amp-matching-engine/ws"
)

// OHLCVService is an autogenerated mock type for the OHLCVService type
type OHLCVService struct {
	mock.Mock
}

// GetOHLCV provides a mock function with given fields: p, duration, unit, timeInterval
func (_m *OHLCVService) GetOHLCV(p []types.PairAddresses, duration int64, unit string, timeInterval ...int64) ([]*types.Tick, error) {
	_va := make([]interface{}, len(timeInterval))
	for _i := range timeInterval {
		_va[_i] = timeInterval[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, p, duration, unit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*types.Tick
	if rf, ok := ret.Get(0).(func([]types.PairAddresses, int64, string, ...int64) []*types.Tick); ok {
		r0 = rf(p, duration, unit, timeInterval...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tick)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.PairAddresses, int64, string, ...int64) error); ok {
		r1 = rf(p, duration, unit, timeInterval...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: conn, bt, qt, p
func (_m *OHLCVService) Subscribe(conn *ws.Conn, bt common.Address, qt common.Address, p *types.Params) {
	_m.Called(conn, bt, qt, p)
}

// Unsubscribe provides a mock function with given fields: conn, bt, qt, p
func (_m *OHLCVService) Unsubscribe(conn *ws.Conn, bt common.Address, qt common.Address, p *types.Params) {
	_m.Called(conn, bt, qt, p)
}
