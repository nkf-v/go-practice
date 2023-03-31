package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/loms/internal/domain.OrderItemStockRepository -o ./mocks\order_item_stock_repository_minimock.go -n OrderItemStockRepositoryMock

import (
	"context"
	mm_domain "route256/loms/internal/domain"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// OrderItemStockRepositoryMock implements domain.OrderItemStockRepository
type OrderItemStockRepositoryMock struct {
	t minimock.Tester

	funcDelete          func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error)
	inspectFuncDelete   func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mOrderItemStockRepositoryMockDelete

	funcGetListByOrderID          func(ctx context.Context, orderID int64) (opa1 []*mm_domain.OrderItemStock, err error)
	inspectFuncGetListByOrderID   func(ctx context.Context, orderID int64)
	afterGetListByOrderIDCounter  uint64
	beforeGetListByOrderIDCounter uint64
	GetListByOrderIDMock          mOrderItemStockRepositoryMockGetListByOrderID

	funcSave          func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error)
	inspectFuncSave   func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock)
	afterSaveCounter  uint64
	beforeSaveCounter uint64
	SaveMock          mOrderItemStockRepositoryMockSave
}

// NewOrderItemStockRepositoryMock returns a mock for domain.OrderItemStockRepository
func NewOrderItemStockRepositoryMock(t minimock.Tester) *OrderItemStockRepositoryMock {
	m := &OrderItemStockRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mOrderItemStockRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*OrderItemStockRepositoryMockDeleteParams{}

	m.GetListByOrderIDMock = mOrderItemStockRepositoryMockGetListByOrderID{mock: m}
	m.GetListByOrderIDMock.callArgs = []*OrderItemStockRepositoryMockGetListByOrderIDParams{}

	m.SaveMock = mOrderItemStockRepositoryMockSave{mock: m}
	m.SaveMock.callArgs = []*OrderItemStockRepositoryMockSaveParams{}

	return m
}

type mOrderItemStockRepositoryMockDelete struct {
	mock               *OrderItemStockRepositoryMock
	defaultExpectation *OrderItemStockRepositoryMockDeleteExpectation
	expectations       []*OrderItemStockRepositoryMockDeleteExpectation

	callArgs []*OrderItemStockRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// OrderItemStockRepositoryMockDeleteExpectation specifies expectation struct of the OrderItemStockRepository.Delete
type OrderItemStockRepositoryMockDeleteExpectation struct {
	mock    *OrderItemStockRepositoryMock
	params  *OrderItemStockRepositoryMockDeleteParams
	results *OrderItemStockRepositoryMockDeleteResults
	Counter uint64
}

// OrderItemStockRepositoryMockDeleteParams contains parameters of the OrderItemStockRepository.Delete
type OrderItemStockRepositoryMockDeleteParams struct {
	ctx            context.Context
	orderItemStock *mm_domain.OrderItemStock
}

// OrderItemStockRepositoryMockDeleteResults contains results of the OrderItemStockRepository.Delete
type OrderItemStockRepositoryMockDeleteResults struct {
	err error
}

// Expect sets up expected params for OrderItemStockRepository.Delete
func (mmDelete *mOrderItemStockRepositoryMockDelete) Expect(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) *mOrderItemStockRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("OrderItemStockRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &OrderItemStockRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &OrderItemStockRepositoryMockDeleteParams{ctx, orderItemStock}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the OrderItemStockRepository.Delete
func (mmDelete *mOrderItemStockRepositoryMockDelete) Inspect(f func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock)) *mOrderItemStockRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for OrderItemStockRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by OrderItemStockRepository.Delete
func (mmDelete *mOrderItemStockRepositoryMockDelete) Return(err error) *OrderItemStockRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("OrderItemStockRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &OrderItemStockRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &OrderItemStockRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the OrderItemStockRepository.Delete method
func (mmDelete *mOrderItemStockRepositoryMockDelete) Set(f func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error)) *OrderItemStockRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the OrderItemStockRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the OrderItemStockRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the OrderItemStockRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mOrderItemStockRepositoryMockDelete) When(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) *OrderItemStockRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("OrderItemStockRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &OrderItemStockRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &OrderItemStockRepositoryMockDeleteParams{ctx, orderItemStock},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up OrderItemStockRepository.Delete return parameters for the expectation previously defined by the When method
func (e *OrderItemStockRepositoryMockDeleteExpectation) Then(err error) *OrderItemStockRepositoryMock {
	e.results = &OrderItemStockRepositoryMockDeleteResults{err}
	return e.mock
}

// Delete implements domain.OrderItemStockRepository
func (mmDelete *OrderItemStockRepositoryMock) Delete(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, orderItemStock)
	}

	mm_params := &OrderItemStockRepositoryMockDeleteParams{ctx, orderItemStock}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := OrderItemStockRepositoryMockDeleteParams{ctx, orderItemStock}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("OrderItemStockRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the OrderItemStockRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, orderItemStock)
	}
	mmDelete.t.Fatalf("Unexpected call to OrderItemStockRepositoryMock.Delete. %v %v", ctx, orderItemStock)
	return
}

// DeleteAfterCounter returns a count of finished OrderItemStockRepositoryMock.Delete invocations
func (mmDelete *OrderItemStockRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of OrderItemStockRepositoryMock.Delete invocations
func (mmDelete *OrderItemStockRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to OrderItemStockRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mOrderItemStockRepositoryMockDelete) Calls() []*OrderItemStockRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*OrderItemStockRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *OrderItemStockRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *OrderItemStockRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OrderItemStockRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to OrderItemStockRepositoryMock.Delete")
	}
}

type mOrderItemStockRepositoryMockGetListByOrderID struct {
	mock               *OrderItemStockRepositoryMock
	defaultExpectation *OrderItemStockRepositoryMockGetListByOrderIDExpectation
	expectations       []*OrderItemStockRepositoryMockGetListByOrderIDExpectation

	callArgs []*OrderItemStockRepositoryMockGetListByOrderIDParams
	mutex    sync.RWMutex
}

// OrderItemStockRepositoryMockGetListByOrderIDExpectation specifies expectation struct of the OrderItemStockRepository.GetListByOrderID
type OrderItemStockRepositoryMockGetListByOrderIDExpectation struct {
	mock    *OrderItemStockRepositoryMock
	params  *OrderItemStockRepositoryMockGetListByOrderIDParams
	results *OrderItemStockRepositoryMockGetListByOrderIDResults
	Counter uint64
}

// OrderItemStockRepositoryMockGetListByOrderIDParams contains parameters of the OrderItemStockRepository.GetListByOrderID
type OrderItemStockRepositoryMockGetListByOrderIDParams struct {
	ctx     context.Context
	orderID int64
}

// OrderItemStockRepositoryMockGetListByOrderIDResults contains results of the OrderItemStockRepository.GetListByOrderID
type OrderItemStockRepositoryMockGetListByOrderIDResults struct {
	opa1 []*mm_domain.OrderItemStock
	err  error
}

// Expect sets up expected params for OrderItemStockRepository.GetListByOrderID
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) Expect(ctx context.Context, orderID int64) *mOrderItemStockRepositoryMockGetListByOrderID {
	if mmGetListByOrderID.mock.funcGetListByOrderID != nil {
		mmGetListByOrderID.mock.t.Fatalf("OrderItemStockRepositoryMock.GetListByOrderID mock is already set by Set")
	}

	if mmGetListByOrderID.defaultExpectation == nil {
		mmGetListByOrderID.defaultExpectation = &OrderItemStockRepositoryMockGetListByOrderIDExpectation{}
	}

	mmGetListByOrderID.defaultExpectation.params = &OrderItemStockRepositoryMockGetListByOrderIDParams{ctx, orderID}
	for _, e := range mmGetListByOrderID.expectations {
		if minimock.Equal(e.params, mmGetListByOrderID.defaultExpectation.params) {
			mmGetListByOrderID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetListByOrderID.defaultExpectation.params)
		}
	}

	return mmGetListByOrderID
}

// Inspect accepts an inspector function that has same arguments as the OrderItemStockRepository.GetListByOrderID
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) Inspect(f func(ctx context.Context, orderID int64)) *mOrderItemStockRepositoryMockGetListByOrderID {
	if mmGetListByOrderID.mock.inspectFuncGetListByOrderID != nil {
		mmGetListByOrderID.mock.t.Fatalf("Inspect function is already set for OrderItemStockRepositoryMock.GetListByOrderID")
	}

	mmGetListByOrderID.mock.inspectFuncGetListByOrderID = f

	return mmGetListByOrderID
}

// Return sets up results that will be returned by OrderItemStockRepository.GetListByOrderID
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) Return(opa1 []*mm_domain.OrderItemStock, err error) *OrderItemStockRepositoryMock {
	if mmGetListByOrderID.mock.funcGetListByOrderID != nil {
		mmGetListByOrderID.mock.t.Fatalf("OrderItemStockRepositoryMock.GetListByOrderID mock is already set by Set")
	}

	if mmGetListByOrderID.defaultExpectation == nil {
		mmGetListByOrderID.defaultExpectation = &OrderItemStockRepositoryMockGetListByOrderIDExpectation{mock: mmGetListByOrderID.mock}
	}
	mmGetListByOrderID.defaultExpectation.results = &OrderItemStockRepositoryMockGetListByOrderIDResults{opa1, err}
	return mmGetListByOrderID.mock
}

// Set uses given function f to mock the OrderItemStockRepository.GetListByOrderID method
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) Set(f func(ctx context.Context, orderID int64) (opa1 []*mm_domain.OrderItemStock, err error)) *OrderItemStockRepositoryMock {
	if mmGetListByOrderID.defaultExpectation != nil {
		mmGetListByOrderID.mock.t.Fatalf("Default expectation is already set for the OrderItemStockRepository.GetListByOrderID method")
	}

	if len(mmGetListByOrderID.expectations) > 0 {
		mmGetListByOrderID.mock.t.Fatalf("Some expectations are already set for the OrderItemStockRepository.GetListByOrderID method")
	}

	mmGetListByOrderID.mock.funcGetListByOrderID = f
	return mmGetListByOrderID.mock
}

// When sets expectation for the OrderItemStockRepository.GetListByOrderID which will trigger the result defined by the following
// Then helper
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) When(ctx context.Context, orderID int64) *OrderItemStockRepositoryMockGetListByOrderIDExpectation {
	if mmGetListByOrderID.mock.funcGetListByOrderID != nil {
		mmGetListByOrderID.mock.t.Fatalf("OrderItemStockRepositoryMock.GetListByOrderID mock is already set by Set")
	}

	expectation := &OrderItemStockRepositoryMockGetListByOrderIDExpectation{
		mock:   mmGetListByOrderID.mock,
		params: &OrderItemStockRepositoryMockGetListByOrderIDParams{ctx, orderID},
	}
	mmGetListByOrderID.expectations = append(mmGetListByOrderID.expectations, expectation)
	return expectation
}

// Then sets up OrderItemStockRepository.GetListByOrderID return parameters for the expectation previously defined by the When method
func (e *OrderItemStockRepositoryMockGetListByOrderIDExpectation) Then(opa1 []*mm_domain.OrderItemStock, err error) *OrderItemStockRepositoryMock {
	e.results = &OrderItemStockRepositoryMockGetListByOrderIDResults{opa1, err}
	return e.mock
}

// GetListByOrderID implements domain.OrderItemStockRepository
func (mmGetListByOrderID *OrderItemStockRepositoryMock) GetListByOrderID(ctx context.Context, orderID int64) (opa1 []*mm_domain.OrderItemStock, err error) {
	mm_atomic.AddUint64(&mmGetListByOrderID.beforeGetListByOrderIDCounter, 1)
	defer mm_atomic.AddUint64(&mmGetListByOrderID.afterGetListByOrderIDCounter, 1)

	if mmGetListByOrderID.inspectFuncGetListByOrderID != nil {
		mmGetListByOrderID.inspectFuncGetListByOrderID(ctx, orderID)
	}

	mm_params := &OrderItemStockRepositoryMockGetListByOrderIDParams{ctx, orderID}

	// Record call args
	mmGetListByOrderID.GetListByOrderIDMock.mutex.Lock()
	mmGetListByOrderID.GetListByOrderIDMock.callArgs = append(mmGetListByOrderID.GetListByOrderIDMock.callArgs, mm_params)
	mmGetListByOrderID.GetListByOrderIDMock.mutex.Unlock()

	for _, e := range mmGetListByOrderID.GetListByOrderIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.opa1, e.results.err
		}
	}

	if mmGetListByOrderID.GetListByOrderIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetListByOrderID.GetListByOrderIDMock.defaultExpectation.Counter, 1)
		mm_want := mmGetListByOrderID.GetListByOrderIDMock.defaultExpectation.params
		mm_got := OrderItemStockRepositoryMockGetListByOrderIDParams{ctx, orderID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetListByOrderID.t.Errorf("OrderItemStockRepositoryMock.GetListByOrderID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetListByOrderID.GetListByOrderIDMock.defaultExpectation.results
		if mm_results == nil {
			mmGetListByOrderID.t.Fatal("No results are set for the OrderItemStockRepositoryMock.GetListByOrderID")
		}
		return (*mm_results).opa1, (*mm_results).err
	}
	if mmGetListByOrderID.funcGetListByOrderID != nil {
		return mmGetListByOrderID.funcGetListByOrderID(ctx, orderID)
	}
	mmGetListByOrderID.t.Fatalf("Unexpected call to OrderItemStockRepositoryMock.GetListByOrderID. %v %v", ctx, orderID)
	return
}

// GetListByOrderIDAfterCounter returns a count of finished OrderItemStockRepositoryMock.GetListByOrderID invocations
func (mmGetListByOrderID *OrderItemStockRepositoryMock) GetListByOrderIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetListByOrderID.afterGetListByOrderIDCounter)
}

// GetListByOrderIDBeforeCounter returns a count of OrderItemStockRepositoryMock.GetListByOrderID invocations
func (mmGetListByOrderID *OrderItemStockRepositoryMock) GetListByOrderIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetListByOrderID.beforeGetListByOrderIDCounter)
}

// Calls returns a list of arguments used in each call to OrderItemStockRepositoryMock.GetListByOrderID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetListByOrderID *mOrderItemStockRepositoryMockGetListByOrderID) Calls() []*OrderItemStockRepositoryMockGetListByOrderIDParams {
	mmGetListByOrderID.mutex.RLock()

	argCopy := make([]*OrderItemStockRepositoryMockGetListByOrderIDParams, len(mmGetListByOrderID.callArgs))
	copy(argCopy, mmGetListByOrderID.callArgs)

	mmGetListByOrderID.mutex.RUnlock()

	return argCopy
}

// MinimockGetListByOrderIDDone returns true if the count of the GetListByOrderID invocations corresponds
// the number of defined expectations
func (m *OrderItemStockRepositoryMock) MinimockGetListByOrderIDDone() bool {
	for _, e := range m.GetListByOrderIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListByOrderIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListByOrderIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetListByOrderID != nil && mm_atomic.LoadUint64(&m.afterGetListByOrderIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetListByOrderIDInspect logs each unmet expectation
func (m *OrderItemStockRepositoryMock) MinimockGetListByOrderIDInspect() {
	for _, e := range m.GetListByOrderIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.GetListByOrderID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListByOrderIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListByOrderIDCounter) < 1 {
		if m.GetListByOrderIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OrderItemStockRepositoryMock.GetListByOrderID")
		} else {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.GetListByOrderID with params: %#v", *m.GetListByOrderIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetListByOrderID != nil && mm_atomic.LoadUint64(&m.afterGetListByOrderIDCounter) < 1 {
		m.t.Error("Expected call to OrderItemStockRepositoryMock.GetListByOrderID")
	}
}

type mOrderItemStockRepositoryMockSave struct {
	mock               *OrderItemStockRepositoryMock
	defaultExpectation *OrderItemStockRepositoryMockSaveExpectation
	expectations       []*OrderItemStockRepositoryMockSaveExpectation

	callArgs []*OrderItemStockRepositoryMockSaveParams
	mutex    sync.RWMutex
}

// OrderItemStockRepositoryMockSaveExpectation specifies expectation struct of the OrderItemStockRepository.Save
type OrderItemStockRepositoryMockSaveExpectation struct {
	mock    *OrderItemStockRepositoryMock
	params  *OrderItemStockRepositoryMockSaveParams
	results *OrderItemStockRepositoryMockSaveResults
	Counter uint64
}

// OrderItemStockRepositoryMockSaveParams contains parameters of the OrderItemStockRepository.Save
type OrderItemStockRepositoryMockSaveParams struct {
	ctx            context.Context
	orderItemStock *mm_domain.OrderItemStock
}

// OrderItemStockRepositoryMockSaveResults contains results of the OrderItemStockRepository.Save
type OrderItemStockRepositoryMockSaveResults struct {
	err error
}

// Expect sets up expected params for OrderItemStockRepository.Save
func (mmSave *mOrderItemStockRepositoryMockSave) Expect(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) *mOrderItemStockRepositoryMockSave {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("OrderItemStockRepositoryMock.Save mock is already set by Set")
	}

	if mmSave.defaultExpectation == nil {
		mmSave.defaultExpectation = &OrderItemStockRepositoryMockSaveExpectation{}
	}

	mmSave.defaultExpectation.params = &OrderItemStockRepositoryMockSaveParams{ctx, orderItemStock}
	for _, e := range mmSave.expectations {
		if minimock.Equal(e.params, mmSave.defaultExpectation.params) {
			mmSave.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSave.defaultExpectation.params)
		}
	}

	return mmSave
}

// Inspect accepts an inspector function that has same arguments as the OrderItemStockRepository.Save
func (mmSave *mOrderItemStockRepositoryMockSave) Inspect(f func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock)) *mOrderItemStockRepositoryMockSave {
	if mmSave.mock.inspectFuncSave != nil {
		mmSave.mock.t.Fatalf("Inspect function is already set for OrderItemStockRepositoryMock.Save")
	}

	mmSave.mock.inspectFuncSave = f

	return mmSave
}

// Return sets up results that will be returned by OrderItemStockRepository.Save
func (mmSave *mOrderItemStockRepositoryMockSave) Return(err error) *OrderItemStockRepositoryMock {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("OrderItemStockRepositoryMock.Save mock is already set by Set")
	}

	if mmSave.defaultExpectation == nil {
		mmSave.defaultExpectation = &OrderItemStockRepositoryMockSaveExpectation{mock: mmSave.mock}
	}
	mmSave.defaultExpectation.results = &OrderItemStockRepositoryMockSaveResults{err}
	return mmSave.mock
}

// Set uses given function f to mock the OrderItemStockRepository.Save method
func (mmSave *mOrderItemStockRepositoryMockSave) Set(f func(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error)) *OrderItemStockRepositoryMock {
	if mmSave.defaultExpectation != nil {
		mmSave.mock.t.Fatalf("Default expectation is already set for the OrderItemStockRepository.Save method")
	}

	if len(mmSave.expectations) > 0 {
		mmSave.mock.t.Fatalf("Some expectations are already set for the OrderItemStockRepository.Save method")
	}

	mmSave.mock.funcSave = f
	return mmSave.mock
}

// When sets expectation for the OrderItemStockRepository.Save which will trigger the result defined by the following
// Then helper
func (mmSave *mOrderItemStockRepositoryMockSave) When(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) *OrderItemStockRepositoryMockSaveExpectation {
	if mmSave.mock.funcSave != nil {
		mmSave.mock.t.Fatalf("OrderItemStockRepositoryMock.Save mock is already set by Set")
	}

	expectation := &OrderItemStockRepositoryMockSaveExpectation{
		mock:   mmSave.mock,
		params: &OrderItemStockRepositoryMockSaveParams{ctx, orderItemStock},
	}
	mmSave.expectations = append(mmSave.expectations, expectation)
	return expectation
}

// Then sets up OrderItemStockRepository.Save return parameters for the expectation previously defined by the When method
func (e *OrderItemStockRepositoryMockSaveExpectation) Then(err error) *OrderItemStockRepositoryMock {
	e.results = &OrderItemStockRepositoryMockSaveResults{err}
	return e.mock
}

// Save implements domain.OrderItemStockRepository
func (mmSave *OrderItemStockRepositoryMock) Save(ctx context.Context, orderItemStock *mm_domain.OrderItemStock) (err error) {
	mm_atomic.AddUint64(&mmSave.beforeSaveCounter, 1)
	defer mm_atomic.AddUint64(&mmSave.afterSaveCounter, 1)

	if mmSave.inspectFuncSave != nil {
		mmSave.inspectFuncSave(ctx, orderItemStock)
	}

	mm_params := &OrderItemStockRepositoryMockSaveParams{ctx, orderItemStock}

	// Record call args
	mmSave.SaveMock.mutex.Lock()
	mmSave.SaveMock.callArgs = append(mmSave.SaveMock.callArgs, mm_params)
	mmSave.SaveMock.mutex.Unlock()

	for _, e := range mmSave.SaveMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSave.SaveMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSave.SaveMock.defaultExpectation.Counter, 1)
		mm_want := mmSave.SaveMock.defaultExpectation.params
		mm_got := OrderItemStockRepositoryMockSaveParams{ctx, orderItemStock}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSave.t.Errorf("OrderItemStockRepositoryMock.Save got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSave.SaveMock.defaultExpectation.results
		if mm_results == nil {
			mmSave.t.Fatal("No results are set for the OrderItemStockRepositoryMock.Save")
		}
		return (*mm_results).err
	}
	if mmSave.funcSave != nil {
		return mmSave.funcSave(ctx, orderItemStock)
	}
	mmSave.t.Fatalf("Unexpected call to OrderItemStockRepositoryMock.Save. %v %v", ctx, orderItemStock)
	return
}

// SaveAfterCounter returns a count of finished OrderItemStockRepositoryMock.Save invocations
func (mmSave *OrderItemStockRepositoryMock) SaveAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSave.afterSaveCounter)
}

// SaveBeforeCounter returns a count of OrderItemStockRepositoryMock.Save invocations
func (mmSave *OrderItemStockRepositoryMock) SaveBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSave.beforeSaveCounter)
}

// Calls returns a list of arguments used in each call to OrderItemStockRepositoryMock.Save.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSave *mOrderItemStockRepositoryMockSave) Calls() []*OrderItemStockRepositoryMockSaveParams {
	mmSave.mutex.RLock()

	argCopy := make([]*OrderItemStockRepositoryMockSaveParams, len(mmSave.callArgs))
	copy(argCopy, mmSave.callArgs)

	mmSave.mutex.RUnlock()

	return argCopy
}

// MinimockSaveDone returns true if the count of the Save invocations corresponds
// the number of defined expectations
func (m *OrderItemStockRepositoryMock) MinimockSaveDone() bool {
	for _, e := range m.SaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSave != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		return false
	}
	return true
}

// MinimockSaveInspect logs each unmet expectation
func (m *OrderItemStockRepositoryMock) MinimockSaveInspect() {
	for _, e := range m.SaveMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.Save with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SaveMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		if m.SaveMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OrderItemStockRepositoryMock.Save")
		} else {
			m.t.Errorf("Expected call to OrderItemStockRepositoryMock.Save with params: %#v", *m.SaveMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSave != nil && mm_atomic.LoadUint64(&m.afterSaveCounter) < 1 {
		m.t.Error("Expected call to OrderItemStockRepositoryMock.Save")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OrderItemStockRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteInspect()

		m.MinimockGetListByOrderIDInspect()

		m.MinimockSaveInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OrderItemStockRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *OrderItemStockRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockGetListByOrderIDDone() &&
		m.MinimockSaveDone()
}