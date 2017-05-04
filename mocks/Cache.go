package mocks

import "github.com/DispatchMe/go-rediscache"
import "github.com/stretchr/testify/mock"

type Cache struct {
	mock.Mock
}

func (_m *Cache) SetInt(key string, val int, opts ...rediscache.SetOption) error {
	ret := _m.Called(key, val, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, ...rediscache.SetOption) error); ok {
		r0 = rf(key, val, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Cache) SetString(key string, val string, opts ...rediscache.SetOption) error {
	ret := _m.Called(key, val, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...rediscache.SetOption) error); ok {
		r0 = rf(key, val, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Cache) SetBool(key string, val bool, opts ...rediscache.SetOption) error {
	ret := _m.Called(key, val, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, ...rediscache.SetOption) error); ok {
		r0 = rf(key, val, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Cache) SetFloat(key string, val float64, opts ...rediscache.SetOption) error {
	ret := _m.Called(key, val, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, ...rediscache.SetOption) error); ok {
		r0 = rf(key, val, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Cache) GetInt(key string) (int, error) {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Cache) GetString(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Cache) GetBool(key string) (bool, error) {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Cache) GetFloat(key string) (float64, error) {
	ret := _m.Called(key)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *Cache) SetJSON(key string, val interface{}, opts ...rediscache.SetOption) error {
	ret := _m.Called(key, val, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, ...rediscache.SetOption) error); ok {
		r0 = rf(key, val, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *Cache) GetJSON(key string, dest interface{}) error {
	ret := _m.Called(key, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
