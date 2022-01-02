package katas

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Service interface {
	DoSomething(int) (bool, error)
}

type RealService struct{}

func (m *RealService) DoSomething(number int) (bool, error) {
	return false, errors.New("Really expensive 3rd party call")
}

func UnderTest(service Service, number int) (bool, error) {
	return service.DoSomething(number)
}

type MyTestDoubleObject struct {
	mock.Mock
}

func (m *MyTestDoubleObject) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func TestRealServiceIsFaliling(t *testing.T) {
	_, err := UnderTest(new(RealService), 42)
	assert.Equal(t, "Really expensive 3rd party call", err.Error())
}

type MyStub struct{}

func (m *MyStub) DoSomething(number int) (bool, error) {
	return true, nil
}

func TestMyStubIsReturningTrue(t *testing.T) {
	result, err := UnderTest(new(MyStub), 42)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestStubIsReturningTure(t *testing.T) {
	stub := new(MyTestDoubleObject)

	stub.On("DoSomething", mock.Anything).Return(true, nil)

	result, err := UnderTest(stub, 42)
	assert.True(t, result)
	assert.Nil(t, err)
}

func TestMockIsCalledWith42(t *testing.T) {
	stub := new(MyTestDoubleObject)
	stub.On("DoSomething", 42).Return(false, nil)

	UnderTest(stub, 42)

	stub.AssertExpectations(t)
}

func TestMockIsFailingWhenNotCalledWith42(t *testing.T) {
	stub := new(MyTestDoubleObject)
	stub.On("DoSomething", 42).Return(false, nil)

	UnderTest(stub, 13)

	stub.AssertExpectations(t)
}
