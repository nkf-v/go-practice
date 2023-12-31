package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain.OrderRepository -o ./mocks\order_repository_minimock.go -n OrderRepositoryMock

import (
	"context"
	mm_domain "route256/checkout/internal/domain"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// OrderRepositoryMock implements domain.OrderRepository
type OrderRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, user int64, items []*mm_domain.CartItem) (err error)
	inspectFuncCreate   func(ctx context.Context, user int64, items []*mm_domain.CartItem)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mOrderRepositoryMockCreate
}

// NewOrderRepositoryMock returns a mock for domain.OrderRepository
func NewOrderRepositoryMock(t minimock.Tester) *OrderRepositoryMock {
	m := &OrderRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mOrderRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*OrderRepositoryMockCreateParams{}

	return m
}

type mOrderRepositoryMockCreate struct {
	mock               *OrderRepositoryMock
	defaultExpectation *OrderRepositoryMockCreateExpectation
	expectations       []*OrderRepositoryMockCreateExpectation

	callArgs []*OrderRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// OrderRepositoryMockCreateExpectation specifies expectation struct of the OrderRepository.Create
type OrderRepositoryMockCreateExpectation struct {
	mock    *OrderRepositoryMock
	params  *OrderRepositoryMockCreateParams
	results *OrderRepositoryMockCreateResults
	Counter uint64
}

// OrderRepositoryMockCreateParams contains parameters of the OrderRepository.Create
type OrderRepositoryMockCreateParams struct {
	ctx   context.Context
	user  int64
	items []*mm_domain.CartItem
}

// OrderRepositoryMockCreateResults contains results of the OrderRepository.Create
type OrderRepositoryMockCreateResults struct {
	err error
}

// Expect sets up expected params for OrderRepository.Create
func (mmCreate *mOrderRepositoryMockCreate) Expect(ctx context.Context, user int64, items []*mm_domain.CartItem) *mOrderRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("OrderRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &OrderRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &OrderRepositoryMockCreateParams{ctx, user, items}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the OrderRepository.Create
func (mmCreate *mOrderRepositoryMockCreate) Inspect(f func(ctx context.Context, user int64, items []*mm_domain.CartItem)) *mOrderRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for OrderRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by OrderRepository.Create
func (mmCreate *mOrderRepositoryMockCreate) Return(err error) *OrderRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("OrderRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &OrderRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &OrderRepositoryMockCreateResults{err}
	return mmCreate.mock
}

// Set uses given function f to mock the OrderRepository.Create method
func (mmCreate *mOrderRepositoryMockCreate) Set(f func(ctx context.Context, user int64, items []*mm_domain.CartItem) (err error)) *OrderRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the OrderRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the OrderRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the OrderRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mOrderRepositoryMockCreate) When(ctx context.Context, user int64, items []*mm_domain.CartItem) *OrderRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("OrderRepositoryMock.Create mock is already set by Set")
	}

	expectation := &OrderRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &OrderRepositoryMockCreateParams{ctx, user, items},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up OrderRepository.Create return parameters for the expectation previously defined by the When method
func (e *OrderRepositoryMockCreateExpectation) Then(err error) *OrderRepositoryMock {
	e.results = &OrderRepositoryMockCreateResults{err}
	return e.mock
}

// Create implements domain.OrderRepository
func (mmCreate *OrderRepositoryMock) Create(ctx context.Context, user int64, items []*mm_domain.CartItem) (err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, user, items)
	}

	mm_params := &OrderRepositoryMockCreateParams{ctx, user, items}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := OrderRepositoryMockCreateParams{ctx, user, items}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("OrderRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the OrderRepositoryMock.Create")
		}
		return (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, user, items)
	}
	mmCreate.t.Fatalf("Unexpected call to OrderRepositoryMock.Create. %v %v %v", ctx, user, items)
	return
}

// CreateAfterCounter returns a count of finished OrderRepositoryMock.Create invocations
func (mmCreate *OrderRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of OrderRepositoryMock.Create invocations
func (mmCreate *OrderRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to OrderRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mOrderRepositoryMockCreate) Calls() []*OrderRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*OrderRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *OrderRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *OrderRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to OrderRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to OrderRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to OrderRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to OrderRepositoryMock.Create")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *OrderRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *OrderRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *OrderRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone()
}
