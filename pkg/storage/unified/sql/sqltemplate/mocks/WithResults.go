// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	reflect "reflect"

	mock "github.com/stretchr/testify/mock"

	sqltemplate "github.com/grafana/grafana/pkg/storage/unified/sql/sqltemplate"
)

// WithResults is an autogenerated mock type for the WithResults type
type WithResults[T interface{}] struct {
	mock.Mock
}

type WithResults_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *WithResults[T]) EXPECT() *WithResults_Expecter[T] {
	return &WithResults_Expecter[T]{mock: &_m.Mock}
}

// Arg provides a mock function with given fields: x
func (_m *WithResults[T]) Arg(x interface{}) string {
	ret := _m.Called(x)

	if len(ret) == 0 {
		panic("no return value specified for Arg")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(x)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WithResults_Arg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Arg'
type WithResults_Arg_Call[T interface{}] struct {
	*mock.Call
}

// Arg is a helper method to define mock.On call
//   - x interface{}
func (_e *WithResults_Expecter[T]) Arg(x interface{}) *WithResults_Arg_Call[T] {
	return &WithResults_Arg_Call[T]{Call: _e.mock.On("Arg", x)}
}

func (_c *WithResults_Arg_Call[T]) Run(run func(x interface{})) *WithResults_Arg_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *WithResults_Arg_Call[T]) Return(_a0 string) *WithResults_Arg_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_Arg_Call[T]) RunAndReturn(run func(interface{}) string) *WithResults_Arg_Call[T] {
	_c.Call.Return(run)
	return _c
}

// ArgList provides a mock function with given fields: slice
func (_m *WithResults[T]) ArgList(slice reflect.Value) (string, error) {
	ret := _m.Called(slice)

	if len(ret) == 0 {
		panic("no return value specified for ArgList")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(reflect.Value) (string, error)); ok {
		return rf(slice)
	}
	if rf, ok := ret.Get(0).(func(reflect.Value) string); ok {
		r0 = rf(slice)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(reflect.Value) error); ok {
		r1 = rf(slice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithResults_ArgList_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ArgList'
type WithResults_ArgList_Call[T interface{}] struct {
	*mock.Call
}

// ArgList is a helper method to define mock.On call
//   - slice reflect.Value
func (_e *WithResults_Expecter[T]) ArgList(slice interface{}) *WithResults_ArgList_Call[T] {
	return &WithResults_ArgList_Call[T]{Call: _e.mock.On("ArgList", slice)}
}

func (_c *WithResults_ArgList_Call[T]) Run(run func(slice reflect.Value)) *WithResults_ArgList_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reflect.Value))
	})
	return _c
}

func (_c *WithResults_ArgList_Call[T]) Return(_a0 string, _a1 error) *WithResults_ArgList_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithResults_ArgList_Call[T]) RunAndReturn(run func(reflect.Value) (string, error)) *WithResults_ArgList_Call[T] {
	_c.Call.Return(run)
	return _c
}

// ArgPlaceholder provides a mock function with given fields: argNum
func (_m *WithResults[T]) ArgPlaceholder(argNum int) string {
	ret := _m.Called(argNum)

	if len(ret) == 0 {
		panic("no return value specified for ArgPlaceholder")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(argNum)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WithResults_ArgPlaceholder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ArgPlaceholder'
type WithResults_ArgPlaceholder_Call[T interface{}] struct {
	*mock.Call
}

// ArgPlaceholder is a helper method to define mock.On call
//   - argNum int
func (_e *WithResults_Expecter[T]) ArgPlaceholder(argNum interface{}) *WithResults_ArgPlaceholder_Call[T] {
	return &WithResults_ArgPlaceholder_Call[T]{Call: _e.mock.On("ArgPlaceholder", argNum)}
}

func (_c *WithResults_ArgPlaceholder_Call[T]) Run(run func(argNum int)) *WithResults_ArgPlaceholder_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *WithResults_ArgPlaceholder_Call[T]) Return(_a0 string) *WithResults_ArgPlaceholder_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_ArgPlaceholder_Call[T]) RunAndReturn(run func(int) string) *WithResults_ArgPlaceholder_Call[T] {
	_c.Call.Return(run)
	return _c
}

// CurrentEpoch provides a mock function with given fields:
func (_m *WithResults[T]) CurrentEpoch() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CurrentEpoch")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WithResults_CurrentEpoch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentEpoch'
type WithResults_CurrentEpoch_Call[T interface{}] struct {
	*mock.Call
}

// CurrentEpoch is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) CurrentEpoch() *WithResults_CurrentEpoch_Call[T] {
	return &WithResults_CurrentEpoch_Call[T]{Call: _e.mock.On("CurrentEpoch")}
}

func (_c *WithResults_CurrentEpoch_Call[T]) Run(run func()) *WithResults_CurrentEpoch_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_CurrentEpoch_Call[T]) Return(_a0 string) *WithResults_CurrentEpoch_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_CurrentEpoch_Call[T]) RunAndReturn(run func() string) *WithResults_CurrentEpoch_Call[T] {
	_c.Call.Return(run)
	return _c
}

// DialectName provides a mock function with given fields:
func (_m *WithResults[T]) DialectName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DialectName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WithResults_DialectName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DialectName'
type WithResults_DialectName_Call[T interface{}] struct {
	*mock.Call
}

// DialectName is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) DialectName() *WithResults_DialectName_Call[T] {
	return &WithResults_DialectName_Call[T]{Call: _e.mock.On("DialectName")}
}

func (_c *WithResults_DialectName_Call[T]) Run(run func()) *WithResults_DialectName_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_DialectName_Call[T]) Return(_a0 string) *WithResults_DialectName_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_DialectName_Call[T]) RunAndReturn(run func() string) *WithResults_DialectName_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetArgs provides a mock function with given fields:
func (_m *WithResults[T]) GetArgs() []interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetArgs")
	}

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func() []interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// WithResults_GetArgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetArgs'
type WithResults_GetArgs_Call[T interface{}] struct {
	*mock.Call
}

// GetArgs is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) GetArgs() *WithResults_GetArgs_Call[T] {
	return &WithResults_GetArgs_Call[T]{Call: _e.mock.On("GetArgs")}
}

func (_c *WithResults_GetArgs_Call[T]) Run(run func()) *WithResults_GetArgs_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_GetArgs_Call[T]) Return(_a0 []interface{}) *WithResults_GetArgs_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_GetArgs_Call[T]) RunAndReturn(run func() []interface{}) *WithResults_GetArgs_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetColNames provides a mock function with given fields:
func (_m *WithResults[T]) GetColNames() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetColNames")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// WithResults_GetColNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetColNames'
type WithResults_GetColNames_Call[T interface{}] struct {
	*mock.Call
}

// GetColNames is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) GetColNames() *WithResults_GetColNames_Call[T] {
	return &WithResults_GetColNames_Call[T]{Call: _e.mock.On("GetColNames")}
}

func (_c *WithResults_GetColNames_Call[T]) Run(run func()) *WithResults_GetColNames_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_GetColNames_Call[T]) Return(_a0 []string) *WithResults_GetColNames_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_GetColNames_Call[T]) RunAndReturn(run func() []string) *WithResults_GetColNames_Call[T] {
	_c.Call.Return(run)
	return _c
}

// GetScanDest provides a mock function with given fields:
func (_m *WithResults[T]) GetScanDest() []interface{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetScanDest")
	}

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func() []interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	return r0
}

// WithResults_GetScanDest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetScanDest'
type WithResults_GetScanDest_Call[T interface{}] struct {
	*mock.Call
}

// GetScanDest is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) GetScanDest() *WithResults_GetScanDest_Call[T] {
	return &WithResults_GetScanDest_Call[T]{Call: _e.mock.On("GetScanDest")}
}

func (_c *WithResults_GetScanDest_Call[T]) Run(run func()) *WithResults_GetScanDest_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_GetScanDest_Call[T]) Return(_a0 []interface{}) *WithResults_GetScanDest_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_GetScanDest_Call[T]) RunAndReturn(run func() []interface{}) *WithResults_GetScanDest_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Ident provides a mock function with given fields: _a0
func (_m *WithResults[T]) Ident(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Ident")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithResults_Ident_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ident'
type WithResults_Ident_Call[T interface{}] struct {
	*mock.Call
}

// Ident is a helper method to define mock.On call
//   - _a0 string
func (_e *WithResults_Expecter[T]) Ident(_a0 interface{}) *WithResults_Ident_Call[T] {
	return &WithResults_Ident_Call[T]{Call: _e.mock.On("Ident", _a0)}
}

func (_c *WithResults_Ident_Call[T]) Run(run func(_a0 string)) *WithResults_Ident_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *WithResults_Ident_Call[T]) Return(_a0 string, _a1 error) *WithResults_Ident_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithResults_Ident_Call[T]) RunAndReturn(run func(string) (string, error)) *WithResults_Ident_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Into provides a mock function with given fields: v, colName
func (_m *WithResults[T]) Into(v reflect.Value, colName string) (string, error) {
	ret := _m.Called(v, colName)

	if len(ret) == 0 {
		panic("no return value specified for Into")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(reflect.Value, string) (string, error)); ok {
		return rf(v, colName)
	}
	if rf, ok := ret.Get(0).(func(reflect.Value, string) string); ok {
		r0 = rf(v, colName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(reflect.Value, string) error); ok {
		r1 = rf(v, colName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithResults_Into_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Into'
type WithResults_Into_Call[T interface{}] struct {
	*mock.Call
}

// Into is a helper method to define mock.On call
//   - v reflect.Value
//   - colName string
func (_e *WithResults_Expecter[T]) Into(v interface{}, colName interface{}) *WithResults_Into_Call[T] {
	return &WithResults_Into_Call[T]{Call: _e.mock.On("Into", v, colName)}
}

func (_c *WithResults_Into_Call[T]) Run(run func(v reflect.Value, colName string)) *WithResults_Into_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(reflect.Value), args[1].(string))
	})
	return _c
}

func (_c *WithResults_Into_Call[T]) Return(_a0 string, _a1 error) *WithResults_Into_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithResults_Into_Call[T]) RunAndReturn(run func(reflect.Value, string) (string, error)) *WithResults_Into_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Reset provides a mock function with given fields:
func (_m *WithResults[T]) Reset() {
	_m.Called()
}

// WithResults_Reset_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Reset'
type WithResults_Reset_Call[T interface{}] struct {
	*mock.Call
}

// Reset is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) Reset() *WithResults_Reset_Call[T] {
	return &WithResults_Reset_Call[T]{Call: _e.mock.On("Reset")}
}

func (_c *WithResults_Reset_Call[T]) Run(run func()) *WithResults_Reset_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_Reset_Call[T]) Return() *WithResults_Reset_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *WithResults_Reset_Call[T]) RunAndReturn(run func()) *WithResults_Reset_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Results provides a mock function with given fields:
func (_m *WithResults[T]) Results() (T, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Results")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func() (T, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithResults_Results_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Results'
type WithResults_Results_Call[T interface{}] struct {
	*mock.Call
}

// Results is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) Results() *WithResults_Results_Call[T] {
	return &WithResults_Results_Call[T]{Call: _e.mock.On("Results")}
}

func (_c *WithResults_Results_Call[T]) Run(run func()) *WithResults_Results_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_Results_Call[T]) Return(_a0 T, _a1 error) *WithResults_Results_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithResults_Results_Call[T]) RunAndReturn(run func() (T, error)) *WithResults_Results_Call[T] {
	_c.Call.Return(run)
	return _c
}

// SelectFor provides a mock function with given fields: _a0
func (_m *WithResults[T]) SelectFor(_a0 ...string) (string, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SelectFor")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(...string) (string, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithResults_SelectFor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectFor'
type WithResults_SelectFor_Call[T interface{}] struct {
	*mock.Call
}

// SelectFor is a helper method to define mock.On call
//   - _a0 ...string
func (_e *WithResults_Expecter[T]) SelectFor(_a0 ...interface{}) *WithResults_SelectFor_Call[T] {
	return &WithResults_SelectFor_Call[T]{Call: _e.mock.On("SelectFor",
		append([]interface{}{}, _a0...)...)}
}

func (_c *WithResults_SelectFor_Call[T]) Run(run func(_a0 ...string)) *WithResults_SelectFor_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *WithResults_SelectFor_Call[T]) Return(_a0 string, _a1 error) *WithResults_SelectFor_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WithResults_SelectFor_Call[T]) RunAndReturn(run func(...string) (string, error)) *WithResults_SelectFor_Call[T] {
	_c.Call.Return(run)
	return _c
}

// SetDialect provides a mock function with given fields: _a0
func (_m *WithResults[T]) SetDialect(_a0 sqltemplate.Dialect) {
	_m.Called(_a0)
}

// WithResults_SetDialect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDialect'
type WithResults_SetDialect_Call[T interface{}] struct {
	*mock.Call
}

// SetDialect is a helper method to define mock.On call
//   - _a0 sqltemplate.Dialect
func (_e *WithResults_Expecter[T]) SetDialect(_a0 interface{}) *WithResults_SetDialect_Call[T] {
	return &WithResults_SetDialect_Call[T]{Call: _e.mock.On("SetDialect", _a0)}
}

func (_c *WithResults_SetDialect_Call[T]) Run(run func(_a0 sqltemplate.Dialect)) *WithResults_SetDialect_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(sqltemplate.Dialect))
	})
	return _c
}

func (_c *WithResults_SetDialect_Call[T]) Return() *WithResults_SetDialect_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *WithResults_SetDialect_Call[T]) RunAndReturn(run func(sqltemplate.Dialect)) *WithResults_SetDialect_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields:
func (_m *WithResults[T]) Validate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithResults_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type WithResults_Validate_Call[T interface{}] struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
func (_e *WithResults_Expecter[T]) Validate() *WithResults_Validate_Call[T] {
	return &WithResults_Validate_Call[T]{Call: _e.mock.On("Validate")}
}

func (_c *WithResults_Validate_Call[T]) Run(run func()) *WithResults_Validate_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WithResults_Validate_Call[T]) Return(_a0 error) *WithResults_Validate_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WithResults_Validate_Call[T]) RunAndReturn(run func() error) *WithResults_Validate_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewWithResults creates a new instance of WithResults. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWithResults[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *WithResults[T] {
	mock := &WithResults[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
