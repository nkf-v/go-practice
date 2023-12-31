package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/loms/internal/domain.StockRepository -o ./mocks\stock_repository_minimock.go -n StockRepositoryMock

import (
	"context"
	mm_domain "route256/loms/internal/domain"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// StockRepositoryMock implements domain.StockRepository
type StockRepositoryMock struct {
	t minimock.Tester

	funcGetByWarehouseIDAndSku          func(ctx context.Context, warehouseID int64, sku uint32) (sp1 *mm_domain.Stock, err error)
	inspectFuncGetByWarehouseIDAndSku   func(ctx context.Context, warehouseID int64, sku uint32)
	afterGetByWarehouseIDAndSkuCounter  uint64
	beforeGetByWarehouseIDAndSkuCounter uint64
	GetByWarehouseIDAndSkuMock          mStockRepositoryMockGetByWarehouseIDAndSku

	funcGetListBySKU          func(ctx context.Context, sku uint32) (spa1 []*mm_domain.Stock, err error)
	inspectFuncGetListBySKU   func(ctx context.Context, sku uint32)
	afterGetListBySKUCounter  uint64
	beforeGetListBySKUCounter uint64
	GetListBySKUMock          mStockRepositoryMockGetListBySKU

	funcUpdateCount          func(ctx context.Context, stock *mm_domain.Stock) (err error)
	inspectFuncUpdateCount   func(ctx context.Context, stock *mm_domain.Stock)
	afterUpdateCountCounter  uint64
	beforeUpdateCountCounter uint64
	UpdateCountMock          mStockRepositoryMockUpdateCount
}

// NewStockRepositoryMock returns a mock for domain.StockRepository
func NewStockRepositoryMock(t minimock.Tester) *StockRepositoryMock {
	m := &StockRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetByWarehouseIDAndSkuMock = mStockRepositoryMockGetByWarehouseIDAndSku{mock: m}
	m.GetByWarehouseIDAndSkuMock.callArgs = []*StockRepositoryMockGetByWarehouseIDAndSkuParams{}

	m.GetListBySKUMock = mStockRepositoryMockGetListBySKU{mock: m}
	m.GetListBySKUMock.callArgs = []*StockRepositoryMockGetListBySKUParams{}

	m.UpdateCountMock = mStockRepositoryMockUpdateCount{mock: m}
	m.UpdateCountMock.callArgs = []*StockRepositoryMockUpdateCountParams{}

	return m
}

type mStockRepositoryMockGetByWarehouseIDAndSku struct {
	mock               *StockRepositoryMock
	defaultExpectation *StockRepositoryMockGetByWarehouseIDAndSkuExpectation
	expectations       []*StockRepositoryMockGetByWarehouseIDAndSkuExpectation

	callArgs []*StockRepositoryMockGetByWarehouseIDAndSkuParams
	mutex    sync.RWMutex
}

// StockRepositoryMockGetByWarehouseIDAndSkuExpectation specifies expectation struct of the StockRepository.GetByWarehouseIDAndSku
type StockRepositoryMockGetByWarehouseIDAndSkuExpectation struct {
	mock    *StockRepositoryMock
	params  *StockRepositoryMockGetByWarehouseIDAndSkuParams
	results *StockRepositoryMockGetByWarehouseIDAndSkuResults
	Counter uint64
}

// StockRepositoryMockGetByWarehouseIDAndSkuParams contains parameters of the StockRepository.GetByWarehouseIDAndSku
type StockRepositoryMockGetByWarehouseIDAndSkuParams struct {
	ctx         context.Context
	warehouseID int64
	sku         uint32
}

// StockRepositoryMockGetByWarehouseIDAndSkuResults contains results of the StockRepository.GetByWarehouseIDAndSku
type StockRepositoryMockGetByWarehouseIDAndSkuResults struct {
	sp1 *mm_domain.Stock
	err error
}

// Expect sets up expected params for StockRepository.GetByWarehouseIDAndSku
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) Expect(ctx context.Context, warehouseID int64, sku uint32) *mStockRepositoryMockGetByWarehouseIDAndSku {
	if mmGetByWarehouseIDAndSku.mock.funcGetByWarehouseIDAndSku != nil {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("StockRepositoryMock.GetByWarehouseIDAndSku mock is already set by Set")
	}

	if mmGetByWarehouseIDAndSku.defaultExpectation == nil {
		mmGetByWarehouseIDAndSku.defaultExpectation = &StockRepositoryMockGetByWarehouseIDAndSkuExpectation{}
	}

	mmGetByWarehouseIDAndSku.defaultExpectation.params = &StockRepositoryMockGetByWarehouseIDAndSkuParams{ctx, warehouseID, sku}
	for _, e := range mmGetByWarehouseIDAndSku.expectations {
		if minimock.Equal(e.params, mmGetByWarehouseIDAndSku.defaultExpectation.params) {
			mmGetByWarehouseIDAndSku.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetByWarehouseIDAndSku.defaultExpectation.params)
		}
	}

	return mmGetByWarehouseIDAndSku
}

// Inspect accepts an inspector function that has same arguments as the StockRepository.GetByWarehouseIDAndSku
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) Inspect(f func(ctx context.Context, warehouseID int64, sku uint32)) *mStockRepositoryMockGetByWarehouseIDAndSku {
	if mmGetByWarehouseIDAndSku.mock.inspectFuncGetByWarehouseIDAndSku != nil {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("Inspect function is already set for StockRepositoryMock.GetByWarehouseIDAndSku")
	}

	mmGetByWarehouseIDAndSku.mock.inspectFuncGetByWarehouseIDAndSku = f

	return mmGetByWarehouseIDAndSku
}

// Return sets up results that will be returned by StockRepository.GetByWarehouseIDAndSku
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) Return(sp1 *mm_domain.Stock, err error) *StockRepositoryMock {
	if mmGetByWarehouseIDAndSku.mock.funcGetByWarehouseIDAndSku != nil {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("StockRepositoryMock.GetByWarehouseIDAndSku mock is already set by Set")
	}

	if mmGetByWarehouseIDAndSku.defaultExpectation == nil {
		mmGetByWarehouseIDAndSku.defaultExpectation = &StockRepositoryMockGetByWarehouseIDAndSkuExpectation{mock: mmGetByWarehouseIDAndSku.mock}
	}
	mmGetByWarehouseIDAndSku.defaultExpectation.results = &StockRepositoryMockGetByWarehouseIDAndSkuResults{sp1, err}
	return mmGetByWarehouseIDAndSku.mock
}

// Set uses given function f to mock the StockRepository.GetByWarehouseIDAndSku method
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) Set(f func(ctx context.Context, warehouseID int64, sku uint32) (sp1 *mm_domain.Stock, err error)) *StockRepositoryMock {
	if mmGetByWarehouseIDAndSku.defaultExpectation != nil {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("Default expectation is already set for the StockRepository.GetByWarehouseIDAndSku method")
	}

	if len(mmGetByWarehouseIDAndSku.expectations) > 0 {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("Some expectations are already set for the StockRepository.GetByWarehouseIDAndSku method")
	}

	mmGetByWarehouseIDAndSku.mock.funcGetByWarehouseIDAndSku = f
	return mmGetByWarehouseIDAndSku.mock
}

// When sets expectation for the StockRepository.GetByWarehouseIDAndSku which will trigger the result defined by the following
// Then helper
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) When(ctx context.Context, warehouseID int64, sku uint32) *StockRepositoryMockGetByWarehouseIDAndSkuExpectation {
	if mmGetByWarehouseIDAndSku.mock.funcGetByWarehouseIDAndSku != nil {
		mmGetByWarehouseIDAndSku.mock.t.Fatalf("StockRepositoryMock.GetByWarehouseIDAndSku mock is already set by Set")
	}

	expectation := &StockRepositoryMockGetByWarehouseIDAndSkuExpectation{
		mock:   mmGetByWarehouseIDAndSku.mock,
		params: &StockRepositoryMockGetByWarehouseIDAndSkuParams{ctx, warehouseID, sku},
	}
	mmGetByWarehouseIDAndSku.expectations = append(mmGetByWarehouseIDAndSku.expectations, expectation)
	return expectation
}

// Then sets up StockRepository.GetByWarehouseIDAndSku return parameters for the expectation previously defined by the When method
func (e *StockRepositoryMockGetByWarehouseIDAndSkuExpectation) Then(sp1 *mm_domain.Stock, err error) *StockRepositoryMock {
	e.results = &StockRepositoryMockGetByWarehouseIDAndSkuResults{sp1, err}
	return e.mock
}

// GetByWarehouseIDAndSku implements domain.StockRepository
func (mmGetByWarehouseIDAndSku *StockRepositoryMock) GetByWarehouseIDAndSku(ctx context.Context, warehouseID int64, sku uint32) (sp1 *mm_domain.Stock, err error) {
	mm_atomic.AddUint64(&mmGetByWarehouseIDAndSku.beforeGetByWarehouseIDAndSkuCounter, 1)
	defer mm_atomic.AddUint64(&mmGetByWarehouseIDAndSku.afterGetByWarehouseIDAndSkuCounter, 1)

	if mmGetByWarehouseIDAndSku.inspectFuncGetByWarehouseIDAndSku != nil {
		mmGetByWarehouseIDAndSku.inspectFuncGetByWarehouseIDAndSku(ctx, warehouseID, sku)
	}

	mm_params := &StockRepositoryMockGetByWarehouseIDAndSkuParams{ctx, warehouseID, sku}

	// Record call args
	mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.mutex.Lock()
	mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.callArgs = append(mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.callArgs, mm_params)
	mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.mutex.Unlock()

	for _, e := range mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1, e.results.err
		}
	}

	if mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.defaultExpectation.Counter, 1)
		mm_want := mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.defaultExpectation.params
		mm_got := StockRepositoryMockGetByWarehouseIDAndSkuParams{ctx, warehouseID, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetByWarehouseIDAndSku.t.Errorf("StockRepositoryMock.GetByWarehouseIDAndSku got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetByWarehouseIDAndSku.GetByWarehouseIDAndSkuMock.defaultExpectation.results
		if mm_results == nil {
			mmGetByWarehouseIDAndSku.t.Fatal("No results are set for the StockRepositoryMock.GetByWarehouseIDAndSku")
		}
		return (*mm_results).sp1, (*mm_results).err
	}
	if mmGetByWarehouseIDAndSku.funcGetByWarehouseIDAndSku != nil {
		return mmGetByWarehouseIDAndSku.funcGetByWarehouseIDAndSku(ctx, warehouseID, sku)
	}
	mmGetByWarehouseIDAndSku.t.Fatalf("Unexpected call to StockRepositoryMock.GetByWarehouseIDAndSku. %v %v %v", ctx, warehouseID, sku)
	return
}

// GetByWarehouseIDAndSkuAfterCounter returns a count of finished StockRepositoryMock.GetByWarehouseIDAndSku invocations
func (mmGetByWarehouseIDAndSku *StockRepositoryMock) GetByWarehouseIDAndSkuAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByWarehouseIDAndSku.afterGetByWarehouseIDAndSkuCounter)
}

// GetByWarehouseIDAndSkuBeforeCounter returns a count of StockRepositoryMock.GetByWarehouseIDAndSku invocations
func (mmGetByWarehouseIDAndSku *StockRepositoryMock) GetByWarehouseIDAndSkuBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByWarehouseIDAndSku.beforeGetByWarehouseIDAndSkuCounter)
}

// Calls returns a list of arguments used in each call to StockRepositoryMock.GetByWarehouseIDAndSku.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetByWarehouseIDAndSku *mStockRepositoryMockGetByWarehouseIDAndSku) Calls() []*StockRepositoryMockGetByWarehouseIDAndSkuParams {
	mmGetByWarehouseIDAndSku.mutex.RLock()

	argCopy := make([]*StockRepositoryMockGetByWarehouseIDAndSkuParams, len(mmGetByWarehouseIDAndSku.callArgs))
	copy(argCopy, mmGetByWarehouseIDAndSku.callArgs)

	mmGetByWarehouseIDAndSku.mutex.RUnlock()

	return argCopy
}

// MinimockGetByWarehouseIDAndSkuDone returns true if the count of the GetByWarehouseIDAndSku invocations corresponds
// the number of defined expectations
func (m *StockRepositoryMock) MinimockGetByWarehouseIDAndSkuDone() bool {
	for _, e := range m.GetByWarehouseIDAndSkuMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByWarehouseIDAndSkuMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByWarehouseIDAndSkuCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByWarehouseIDAndSku != nil && mm_atomic.LoadUint64(&m.afterGetByWarehouseIDAndSkuCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetByWarehouseIDAndSkuInspect logs each unmet expectation
func (m *StockRepositoryMock) MinimockGetByWarehouseIDAndSkuInspect() {
	for _, e := range m.GetByWarehouseIDAndSkuMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StockRepositoryMock.GetByWarehouseIDAndSku with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByWarehouseIDAndSkuMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByWarehouseIDAndSkuCounter) < 1 {
		if m.GetByWarehouseIDAndSkuMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StockRepositoryMock.GetByWarehouseIDAndSku")
		} else {
			m.t.Errorf("Expected call to StockRepositoryMock.GetByWarehouseIDAndSku with params: %#v", *m.GetByWarehouseIDAndSkuMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByWarehouseIDAndSku != nil && mm_atomic.LoadUint64(&m.afterGetByWarehouseIDAndSkuCounter) < 1 {
		m.t.Error("Expected call to StockRepositoryMock.GetByWarehouseIDAndSku")
	}
}

type mStockRepositoryMockGetListBySKU struct {
	mock               *StockRepositoryMock
	defaultExpectation *StockRepositoryMockGetListBySKUExpectation
	expectations       []*StockRepositoryMockGetListBySKUExpectation

	callArgs []*StockRepositoryMockGetListBySKUParams
	mutex    sync.RWMutex
}

// StockRepositoryMockGetListBySKUExpectation specifies expectation struct of the StockRepository.GetListBySKU
type StockRepositoryMockGetListBySKUExpectation struct {
	mock    *StockRepositoryMock
	params  *StockRepositoryMockGetListBySKUParams
	results *StockRepositoryMockGetListBySKUResults
	Counter uint64
}

// StockRepositoryMockGetListBySKUParams contains parameters of the StockRepository.GetListBySKU
type StockRepositoryMockGetListBySKUParams struct {
	ctx context.Context
	sku uint32
}

// StockRepositoryMockGetListBySKUResults contains results of the StockRepository.GetListBySKU
type StockRepositoryMockGetListBySKUResults struct {
	spa1 []*mm_domain.Stock
	err  error
}

// Expect sets up expected params for StockRepository.GetListBySKU
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) Expect(ctx context.Context, sku uint32) *mStockRepositoryMockGetListBySKU {
	if mmGetListBySKU.mock.funcGetListBySKU != nil {
		mmGetListBySKU.mock.t.Fatalf("StockRepositoryMock.GetListBySKU mock is already set by Set")
	}

	if mmGetListBySKU.defaultExpectation == nil {
		mmGetListBySKU.defaultExpectation = &StockRepositoryMockGetListBySKUExpectation{}
	}

	mmGetListBySKU.defaultExpectation.params = &StockRepositoryMockGetListBySKUParams{ctx, sku}
	for _, e := range mmGetListBySKU.expectations {
		if minimock.Equal(e.params, mmGetListBySKU.defaultExpectation.params) {
			mmGetListBySKU.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetListBySKU.defaultExpectation.params)
		}
	}

	return mmGetListBySKU
}

// Inspect accepts an inspector function that has same arguments as the StockRepository.GetListBySKU
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) Inspect(f func(ctx context.Context, sku uint32)) *mStockRepositoryMockGetListBySKU {
	if mmGetListBySKU.mock.inspectFuncGetListBySKU != nil {
		mmGetListBySKU.mock.t.Fatalf("Inspect function is already set for StockRepositoryMock.GetListBySKU")
	}

	mmGetListBySKU.mock.inspectFuncGetListBySKU = f

	return mmGetListBySKU
}

// Return sets up results that will be returned by StockRepository.GetListBySKU
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) Return(spa1 []*mm_domain.Stock, err error) *StockRepositoryMock {
	if mmGetListBySKU.mock.funcGetListBySKU != nil {
		mmGetListBySKU.mock.t.Fatalf("StockRepositoryMock.GetListBySKU mock is already set by Set")
	}

	if mmGetListBySKU.defaultExpectation == nil {
		mmGetListBySKU.defaultExpectation = &StockRepositoryMockGetListBySKUExpectation{mock: mmGetListBySKU.mock}
	}
	mmGetListBySKU.defaultExpectation.results = &StockRepositoryMockGetListBySKUResults{spa1, err}
	return mmGetListBySKU.mock
}

// Set uses given function f to mock the StockRepository.GetListBySKU method
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) Set(f func(ctx context.Context, sku uint32) (spa1 []*mm_domain.Stock, err error)) *StockRepositoryMock {
	if mmGetListBySKU.defaultExpectation != nil {
		mmGetListBySKU.mock.t.Fatalf("Default expectation is already set for the StockRepository.GetListBySKU method")
	}

	if len(mmGetListBySKU.expectations) > 0 {
		mmGetListBySKU.mock.t.Fatalf("Some expectations are already set for the StockRepository.GetListBySKU method")
	}

	mmGetListBySKU.mock.funcGetListBySKU = f
	return mmGetListBySKU.mock
}

// When sets expectation for the StockRepository.GetListBySKU which will trigger the result defined by the following
// Then helper
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) When(ctx context.Context, sku uint32) *StockRepositoryMockGetListBySKUExpectation {
	if mmGetListBySKU.mock.funcGetListBySKU != nil {
		mmGetListBySKU.mock.t.Fatalf("StockRepositoryMock.GetListBySKU mock is already set by Set")
	}

	expectation := &StockRepositoryMockGetListBySKUExpectation{
		mock:   mmGetListBySKU.mock,
		params: &StockRepositoryMockGetListBySKUParams{ctx, sku},
	}
	mmGetListBySKU.expectations = append(mmGetListBySKU.expectations, expectation)
	return expectation
}

// Then sets up StockRepository.GetListBySKU return parameters for the expectation previously defined by the When method
func (e *StockRepositoryMockGetListBySKUExpectation) Then(spa1 []*mm_domain.Stock, err error) *StockRepositoryMock {
	e.results = &StockRepositoryMockGetListBySKUResults{spa1, err}
	return e.mock
}

// GetListBySKU implements domain.StockRepository
func (mmGetListBySKU *StockRepositoryMock) GetListBySKU(ctx context.Context, sku uint32) (spa1 []*mm_domain.Stock, err error) {
	mm_atomic.AddUint64(&mmGetListBySKU.beforeGetListBySKUCounter, 1)
	defer mm_atomic.AddUint64(&mmGetListBySKU.afterGetListBySKUCounter, 1)

	if mmGetListBySKU.inspectFuncGetListBySKU != nil {
		mmGetListBySKU.inspectFuncGetListBySKU(ctx, sku)
	}

	mm_params := &StockRepositoryMockGetListBySKUParams{ctx, sku}

	// Record call args
	mmGetListBySKU.GetListBySKUMock.mutex.Lock()
	mmGetListBySKU.GetListBySKUMock.callArgs = append(mmGetListBySKU.GetListBySKUMock.callArgs, mm_params)
	mmGetListBySKU.GetListBySKUMock.mutex.Unlock()

	for _, e := range mmGetListBySKU.GetListBySKUMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.spa1, e.results.err
		}
	}

	if mmGetListBySKU.GetListBySKUMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetListBySKU.GetListBySKUMock.defaultExpectation.Counter, 1)
		mm_want := mmGetListBySKU.GetListBySKUMock.defaultExpectation.params
		mm_got := StockRepositoryMockGetListBySKUParams{ctx, sku}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetListBySKU.t.Errorf("StockRepositoryMock.GetListBySKU got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetListBySKU.GetListBySKUMock.defaultExpectation.results
		if mm_results == nil {
			mmGetListBySKU.t.Fatal("No results are set for the StockRepositoryMock.GetListBySKU")
		}
		return (*mm_results).spa1, (*mm_results).err
	}
	if mmGetListBySKU.funcGetListBySKU != nil {
		return mmGetListBySKU.funcGetListBySKU(ctx, sku)
	}
	mmGetListBySKU.t.Fatalf("Unexpected call to StockRepositoryMock.GetListBySKU. %v %v", ctx, sku)
	return
}

// GetListBySKUAfterCounter returns a count of finished StockRepositoryMock.GetListBySKU invocations
func (mmGetListBySKU *StockRepositoryMock) GetListBySKUAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetListBySKU.afterGetListBySKUCounter)
}

// GetListBySKUBeforeCounter returns a count of StockRepositoryMock.GetListBySKU invocations
func (mmGetListBySKU *StockRepositoryMock) GetListBySKUBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetListBySKU.beforeGetListBySKUCounter)
}

// Calls returns a list of arguments used in each call to StockRepositoryMock.GetListBySKU.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetListBySKU *mStockRepositoryMockGetListBySKU) Calls() []*StockRepositoryMockGetListBySKUParams {
	mmGetListBySKU.mutex.RLock()

	argCopy := make([]*StockRepositoryMockGetListBySKUParams, len(mmGetListBySKU.callArgs))
	copy(argCopy, mmGetListBySKU.callArgs)

	mmGetListBySKU.mutex.RUnlock()

	return argCopy
}

// MinimockGetListBySKUDone returns true if the count of the GetListBySKU invocations corresponds
// the number of defined expectations
func (m *StockRepositoryMock) MinimockGetListBySKUDone() bool {
	for _, e := range m.GetListBySKUMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListBySKUMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListBySKUCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetListBySKU != nil && mm_atomic.LoadUint64(&m.afterGetListBySKUCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetListBySKUInspect logs each unmet expectation
func (m *StockRepositoryMock) MinimockGetListBySKUInspect() {
	for _, e := range m.GetListBySKUMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StockRepositoryMock.GetListBySKU with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListBySKUMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListBySKUCounter) < 1 {
		if m.GetListBySKUMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StockRepositoryMock.GetListBySKU")
		} else {
			m.t.Errorf("Expected call to StockRepositoryMock.GetListBySKU with params: %#v", *m.GetListBySKUMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetListBySKU != nil && mm_atomic.LoadUint64(&m.afterGetListBySKUCounter) < 1 {
		m.t.Error("Expected call to StockRepositoryMock.GetListBySKU")
	}
}

type mStockRepositoryMockUpdateCount struct {
	mock               *StockRepositoryMock
	defaultExpectation *StockRepositoryMockUpdateCountExpectation
	expectations       []*StockRepositoryMockUpdateCountExpectation

	callArgs []*StockRepositoryMockUpdateCountParams
	mutex    sync.RWMutex
}

// StockRepositoryMockUpdateCountExpectation specifies expectation struct of the StockRepository.UpdateCount
type StockRepositoryMockUpdateCountExpectation struct {
	mock    *StockRepositoryMock
	params  *StockRepositoryMockUpdateCountParams
	results *StockRepositoryMockUpdateCountResults
	Counter uint64
}

// StockRepositoryMockUpdateCountParams contains parameters of the StockRepository.UpdateCount
type StockRepositoryMockUpdateCountParams struct {
	ctx   context.Context
	stock *mm_domain.Stock
}

// StockRepositoryMockUpdateCountResults contains results of the StockRepository.UpdateCount
type StockRepositoryMockUpdateCountResults struct {
	err error
}

// Expect sets up expected params for StockRepository.UpdateCount
func (mmUpdateCount *mStockRepositoryMockUpdateCount) Expect(ctx context.Context, stock *mm_domain.Stock) *mStockRepositoryMockUpdateCount {
	if mmUpdateCount.mock.funcUpdateCount != nil {
		mmUpdateCount.mock.t.Fatalf("StockRepositoryMock.UpdateCount mock is already set by Set")
	}

	if mmUpdateCount.defaultExpectation == nil {
		mmUpdateCount.defaultExpectation = &StockRepositoryMockUpdateCountExpectation{}
	}

	mmUpdateCount.defaultExpectation.params = &StockRepositoryMockUpdateCountParams{ctx, stock}
	for _, e := range mmUpdateCount.expectations {
		if minimock.Equal(e.params, mmUpdateCount.defaultExpectation.params) {
			mmUpdateCount.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateCount.defaultExpectation.params)
		}
	}

	return mmUpdateCount
}

// Inspect accepts an inspector function that has same arguments as the StockRepository.UpdateCount
func (mmUpdateCount *mStockRepositoryMockUpdateCount) Inspect(f func(ctx context.Context, stock *mm_domain.Stock)) *mStockRepositoryMockUpdateCount {
	if mmUpdateCount.mock.inspectFuncUpdateCount != nil {
		mmUpdateCount.mock.t.Fatalf("Inspect function is already set for StockRepositoryMock.UpdateCount")
	}

	mmUpdateCount.mock.inspectFuncUpdateCount = f

	return mmUpdateCount
}

// Return sets up results that will be returned by StockRepository.UpdateCount
func (mmUpdateCount *mStockRepositoryMockUpdateCount) Return(err error) *StockRepositoryMock {
	if mmUpdateCount.mock.funcUpdateCount != nil {
		mmUpdateCount.mock.t.Fatalf("StockRepositoryMock.UpdateCount mock is already set by Set")
	}

	if mmUpdateCount.defaultExpectation == nil {
		mmUpdateCount.defaultExpectation = &StockRepositoryMockUpdateCountExpectation{mock: mmUpdateCount.mock}
	}
	mmUpdateCount.defaultExpectation.results = &StockRepositoryMockUpdateCountResults{err}
	return mmUpdateCount.mock
}

// Set uses given function f to mock the StockRepository.UpdateCount method
func (mmUpdateCount *mStockRepositoryMockUpdateCount) Set(f func(ctx context.Context, stock *mm_domain.Stock) (err error)) *StockRepositoryMock {
	if mmUpdateCount.defaultExpectation != nil {
		mmUpdateCount.mock.t.Fatalf("Default expectation is already set for the StockRepository.UpdateCount method")
	}

	if len(mmUpdateCount.expectations) > 0 {
		mmUpdateCount.mock.t.Fatalf("Some expectations are already set for the StockRepository.UpdateCount method")
	}

	mmUpdateCount.mock.funcUpdateCount = f
	return mmUpdateCount.mock
}

// When sets expectation for the StockRepository.UpdateCount which will trigger the result defined by the following
// Then helper
func (mmUpdateCount *mStockRepositoryMockUpdateCount) When(ctx context.Context, stock *mm_domain.Stock) *StockRepositoryMockUpdateCountExpectation {
	if mmUpdateCount.mock.funcUpdateCount != nil {
		mmUpdateCount.mock.t.Fatalf("StockRepositoryMock.UpdateCount mock is already set by Set")
	}

	expectation := &StockRepositoryMockUpdateCountExpectation{
		mock:   mmUpdateCount.mock,
		params: &StockRepositoryMockUpdateCountParams{ctx, stock},
	}
	mmUpdateCount.expectations = append(mmUpdateCount.expectations, expectation)
	return expectation
}

// Then sets up StockRepository.UpdateCount return parameters for the expectation previously defined by the When method
func (e *StockRepositoryMockUpdateCountExpectation) Then(err error) *StockRepositoryMock {
	e.results = &StockRepositoryMockUpdateCountResults{err}
	return e.mock
}

// UpdateCount implements domain.StockRepository
func (mmUpdateCount *StockRepositoryMock) UpdateCount(ctx context.Context, stock *mm_domain.Stock) (err error) {
	mm_atomic.AddUint64(&mmUpdateCount.beforeUpdateCountCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateCount.afterUpdateCountCounter, 1)

	if mmUpdateCount.inspectFuncUpdateCount != nil {
		mmUpdateCount.inspectFuncUpdateCount(ctx, stock)
	}

	mm_params := &StockRepositoryMockUpdateCountParams{ctx, stock}

	// Record call args
	mmUpdateCount.UpdateCountMock.mutex.Lock()
	mmUpdateCount.UpdateCountMock.callArgs = append(mmUpdateCount.UpdateCountMock.callArgs, mm_params)
	mmUpdateCount.UpdateCountMock.mutex.Unlock()

	for _, e := range mmUpdateCount.UpdateCountMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateCount.UpdateCountMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateCount.UpdateCountMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateCount.UpdateCountMock.defaultExpectation.params
		mm_got := StockRepositoryMockUpdateCountParams{ctx, stock}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateCount.t.Errorf("StockRepositoryMock.UpdateCount got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateCount.UpdateCountMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateCount.t.Fatal("No results are set for the StockRepositoryMock.UpdateCount")
		}
		return (*mm_results).err
	}
	if mmUpdateCount.funcUpdateCount != nil {
		return mmUpdateCount.funcUpdateCount(ctx, stock)
	}
	mmUpdateCount.t.Fatalf("Unexpected call to StockRepositoryMock.UpdateCount. %v %v", ctx, stock)
	return
}

// UpdateCountAfterCounter returns a count of finished StockRepositoryMock.UpdateCount invocations
func (mmUpdateCount *StockRepositoryMock) UpdateCountAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateCount.afterUpdateCountCounter)
}

// UpdateCountBeforeCounter returns a count of StockRepositoryMock.UpdateCount invocations
func (mmUpdateCount *StockRepositoryMock) UpdateCountBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateCount.beforeUpdateCountCounter)
}

// Calls returns a list of arguments used in each call to StockRepositoryMock.UpdateCount.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateCount *mStockRepositoryMockUpdateCount) Calls() []*StockRepositoryMockUpdateCountParams {
	mmUpdateCount.mutex.RLock()

	argCopy := make([]*StockRepositoryMockUpdateCountParams, len(mmUpdateCount.callArgs))
	copy(argCopy, mmUpdateCount.callArgs)

	mmUpdateCount.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateCountDone returns true if the count of the UpdateCount invocations corresponds
// the number of defined expectations
func (m *StockRepositoryMock) MinimockUpdateCountDone() bool {
	for _, e := range m.UpdateCountMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateCountMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCountCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateCount != nil && mm_atomic.LoadUint64(&m.afterUpdateCountCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateCountInspect logs each unmet expectation
func (m *StockRepositoryMock) MinimockUpdateCountInspect() {
	for _, e := range m.UpdateCountMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StockRepositoryMock.UpdateCount with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateCountMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCountCounter) < 1 {
		if m.UpdateCountMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StockRepositoryMock.UpdateCount")
		} else {
			m.t.Errorf("Expected call to StockRepositoryMock.UpdateCount with params: %#v", *m.UpdateCountMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateCount != nil && mm_atomic.LoadUint64(&m.afterUpdateCountCounter) < 1 {
		m.t.Error("Expected call to StockRepositoryMock.UpdateCount")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StockRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetByWarehouseIDAndSkuInspect()

		m.MinimockGetListBySKUInspect()

		m.MinimockUpdateCountInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StockRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *StockRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetByWarehouseIDAndSkuDone() &&
		m.MinimockGetListBySKUDone() &&
		m.MinimockUpdateCountDone()
}
