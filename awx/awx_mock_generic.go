package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
)

type mockGeneric[t any] struct {
	mock.Mock
}

func (mockAwx mockGeneric[t]) GetByID(id int, params map[string]string) (*t, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*t), args.Error(1)
}

func (mockAwx mockGeneric[t]) List(params map[string]string) ([]*t, *awx.ResultsList[t], error) {
	args := mockAwx.Called(params)

	var resultList *awx.ResultsList[t]
	if resultListCast, ok := args.Get(1).(*awx.ResultsList[t]); ok {
		resultList = resultListCast
	}

	return args.Get(0).([]*t), resultList, args.Error(2)
}

func (mockAwx mockGeneric[t]) Create(data map[string]interface{}, params map[string]string) (*t, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*t), args.Error(1)
}

func (mockAwx mockGeneric[t]) Update(id int, data map[string]interface{}, params map[string]string) (*t, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*t), args.Error(1)
}

func (mockAwx mockGeneric[t]) Delete(id int) (*t, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*t), args.Error(1)
}
