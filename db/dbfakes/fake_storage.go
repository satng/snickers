// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/snickers/snickers/db"
	"github.com/snickers/snickers/types"
)

type FakeStorage struct {
	StorePresetStub        func(types.Preset) (types.Preset, error)
	storePresetMutex       sync.RWMutex
	storePresetArgsForCall []struct {
		arg1 types.Preset
	}
	storePresetReturns struct {
		result1 types.Preset
		result2 error
	}
	RetrievePresetStub        func(string) (types.Preset, error)
	retrievePresetMutex       sync.RWMutex
	retrievePresetArgsForCall []struct {
		arg1 string
	}
	retrievePresetReturns struct {
		result1 types.Preset
		result2 error
	}
	UpdatePresetStub        func(string, types.Preset) (types.Preset, error)
	updatePresetMutex       sync.RWMutex
	updatePresetArgsForCall []struct {
		arg1 string
		arg2 types.Preset
	}
	updatePresetReturns struct {
		result1 types.Preset
		result2 error
	}
	GetPresetsStub        func() ([]types.Preset, error)
	getPresetsMutex       sync.RWMutex
	getPresetsArgsForCall []struct{}
	getPresetsReturns     struct {
		result1 []types.Preset
		result2 error
	}
	DeletePresetStub        func(string) (types.Preset, error)
	deletePresetMutex       sync.RWMutex
	deletePresetArgsForCall []struct {
		arg1 string
	}
	deletePresetReturns struct {
		result1 types.Preset
		result2 error
	}
	StoreJobStub        func(types.Job) (types.Job, error)
	storeJobMutex       sync.RWMutex
	storeJobArgsForCall []struct {
		arg1 types.Job
	}
	storeJobReturns struct {
		result1 types.Job
		result2 error
	}
	RetrieveJobStub        func(string) (types.Job, error)
	retrieveJobMutex       sync.RWMutex
	retrieveJobArgsForCall []struct {
		arg1 string
	}
	retrieveJobReturns struct {
		result1 types.Job
		result2 error
	}
	UpdateJobStub        func(string, types.Job) (types.Job, error)
	updateJobMutex       sync.RWMutex
	updateJobArgsForCall []struct {
		arg1 string
		arg2 types.Job
	}
	updateJobReturns struct {
		result1 types.Job
		result2 error
	}
	GetJobsStub        func() ([]types.Job, error)
	getJobsMutex       sync.RWMutex
	getJobsArgsForCall []struct{}
	getJobsReturns     struct {
		result1 []types.Job
		result2 error
	}
	ClearDatabaseStub        func() error
	clearDatabaseMutex       sync.RWMutex
	clearDatabaseArgsForCall []struct{}
	clearDatabaseReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStorage) StorePreset(arg1 types.Preset) (types.Preset, error) {
	fake.storePresetMutex.Lock()
	fake.storePresetArgsForCall = append(fake.storePresetArgsForCall, struct {
		arg1 types.Preset
	}{arg1})
	fake.recordInvocation("StorePreset", []interface{}{arg1})
	fake.storePresetMutex.Unlock()
	if fake.StorePresetStub != nil {
		return fake.StorePresetStub(arg1)
	} else {
		return fake.storePresetReturns.result1, fake.storePresetReturns.result2
	}
}

func (fake *FakeStorage) StorePresetCallCount() int {
	fake.storePresetMutex.RLock()
	defer fake.storePresetMutex.RUnlock()
	return len(fake.storePresetArgsForCall)
}

func (fake *FakeStorage) StorePresetArgsForCall(i int) types.Preset {
	fake.storePresetMutex.RLock()
	defer fake.storePresetMutex.RUnlock()
	return fake.storePresetArgsForCall[i].arg1
}

func (fake *FakeStorage) StorePresetReturns(result1 types.Preset, result2 error) {
	fake.StorePresetStub = nil
	fake.storePresetReturns = struct {
		result1 types.Preset
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) RetrievePreset(arg1 string) (types.Preset, error) {
	fake.retrievePresetMutex.Lock()
	fake.retrievePresetArgsForCall = append(fake.retrievePresetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RetrievePreset", []interface{}{arg1})
	fake.retrievePresetMutex.Unlock()
	if fake.RetrievePresetStub != nil {
		return fake.RetrievePresetStub(arg1)
	} else {
		return fake.retrievePresetReturns.result1, fake.retrievePresetReturns.result2
	}
}

func (fake *FakeStorage) RetrievePresetCallCount() int {
	fake.retrievePresetMutex.RLock()
	defer fake.retrievePresetMutex.RUnlock()
	return len(fake.retrievePresetArgsForCall)
}

func (fake *FakeStorage) RetrievePresetArgsForCall(i int) string {
	fake.retrievePresetMutex.RLock()
	defer fake.retrievePresetMutex.RUnlock()
	return fake.retrievePresetArgsForCall[i].arg1
}

func (fake *FakeStorage) RetrievePresetReturns(result1 types.Preset, result2 error) {
	fake.RetrievePresetStub = nil
	fake.retrievePresetReturns = struct {
		result1 types.Preset
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) UpdatePreset(arg1 string, arg2 types.Preset) (types.Preset, error) {
	fake.updatePresetMutex.Lock()
	fake.updatePresetArgsForCall = append(fake.updatePresetArgsForCall, struct {
		arg1 string
		arg2 types.Preset
	}{arg1, arg2})
	fake.recordInvocation("UpdatePreset", []interface{}{arg1, arg2})
	fake.updatePresetMutex.Unlock()
	if fake.UpdatePresetStub != nil {
		return fake.UpdatePresetStub(arg1, arg2)
	} else {
		return fake.updatePresetReturns.result1, fake.updatePresetReturns.result2
	}
}

func (fake *FakeStorage) UpdatePresetCallCount() int {
	fake.updatePresetMutex.RLock()
	defer fake.updatePresetMutex.RUnlock()
	return len(fake.updatePresetArgsForCall)
}

func (fake *FakeStorage) UpdatePresetArgsForCall(i int) (string, types.Preset) {
	fake.updatePresetMutex.RLock()
	defer fake.updatePresetMutex.RUnlock()
	return fake.updatePresetArgsForCall[i].arg1, fake.updatePresetArgsForCall[i].arg2
}

func (fake *FakeStorage) UpdatePresetReturns(result1 types.Preset, result2 error) {
	fake.UpdatePresetStub = nil
	fake.updatePresetReturns = struct {
		result1 types.Preset
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) GetPresets() ([]types.Preset, error) {
	fake.getPresetsMutex.Lock()
	fake.getPresetsArgsForCall = append(fake.getPresetsArgsForCall, struct{}{})
	fake.recordInvocation("GetPresets", []interface{}{})
	fake.getPresetsMutex.Unlock()
	if fake.GetPresetsStub != nil {
		return fake.GetPresetsStub()
	} else {
		return fake.getPresetsReturns.result1, fake.getPresetsReturns.result2
	}
}

func (fake *FakeStorage) GetPresetsCallCount() int {
	fake.getPresetsMutex.RLock()
	defer fake.getPresetsMutex.RUnlock()
	return len(fake.getPresetsArgsForCall)
}

func (fake *FakeStorage) GetPresetsReturns(result1 []types.Preset, result2 error) {
	fake.GetPresetsStub = nil
	fake.getPresetsReturns = struct {
		result1 []types.Preset
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) DeletePreset(arg1 string) (types.Preset, error) {
	fake.deletePresetMutex.Lock()
	fake.deletePresetArgsForCall = append(fake.deletePresetArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeletePreset", []interface{}{arg1})
	fake.deletePresetMutex.Unlock()
	if fake.DeletePresetStub != nil {
		return fake.DeletePresetStub(arg1)
	} else {
		return fake.deletePresetReturns.result1, fake.deletePresetReturns.result2
	}
}

func (fake *FakeStorage) DeletePresetCallCount() int {
	fake.deletePresetMutex.RLock()
	defer fake.deletePresetMutex.RUnlock()
	return len(fake.deletePresetArgsForCall)
}

func (fake *FakeStorage) DeletePresetArgsForCall(i int) string {
	fake.deletePresetMutex.RLock()
	defer fake.deletePresetMutex.RUnlock()
	return fake.deletePresetArgsForCall[i].arg1
}

func (fake *FakeStorage) DeletePresetReturns(result1 types.Preset, result2 error) {
	fake.DeletePresetStub = nil
	fake.deletePresetReturns = struct {
		result1 types.Preset
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) StoreJob(arg1 types.Job) (types.Job, error) {
	fake.storeJobMutex.Lock()
	fake.storeJobArgsForCall = append(fake.storeJobArgsForCall, struct {
		arg1 types.Job
	}{arg1})
	fake.recordInvocation("StoreJob", []interface{}{arg1})
	fake.storeJobMutex.Unlock()
	if fake.StoreJobStub != nil {
		return fake.StoreJobStub(arg1)
	} else {
		return fake.storeJobReturns.result1, fake.storeJobReturns.result2
	}
}

func (fake *FakeStorage) StoreJobCallCount() int {
	fake.storeJobMutex.RLock()
	defer fake.storeJobMutex.RUnlock()
	return len(fake.storeJobArgsForCall)
}

func (fake *FakeStorage) StoreJobArgsForCall(i int) types.Job {
	fake.storeJobMutex.RLock()
	defer fake.storeJobMutex.RUnlock()
	return fake.storeJobArgsForCall[i].arg1
}

func (fake *FakeStorage) StoreJobReturns(result1 types.Job, result2 error) {
	fake.StoreJobStub = nil
	fake.storeJobReturns = struct {
		result1 types.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) RetrieveJob(arg1 string) (types.Job, error) {
	fake.retrieveJobMutex.Lock()
	fake.retrieveJobArgsForCall = append(fake.retrieveJobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RetrieveJob", []interface{}{arg1})
	fake.retrieveJobMutex.Unlock()
	if fake.RetrieveJobStub != nil {
		return fake.RetrieveJobStub(arg1)
	} else {
		return fake.retrieveJobReturns.result1, fake.retrieveJobReturns.result2
	}
}

func (fake *FakeStorage) RetrieveJobCallCount() int {
	fake.retrieveJobMutex.RLock()
	defer fake.retrieveJobMutex.RUnlock()
	return len(fake.retrieveJobArgsForCall)
}

func (fake *FakeStorage) RetrieveJobArgsForCall(i int) string {
	fake.retrieveJobMutex.RLock()
	defer fake.retrieveJobMutex.RUnlock()
	return fake.retrieveJobArgsForCall[i].arg1
}

func (fake *FakeStorage) RetrieveJobReturns(result1 types.Job, result2 error) {
	fake.RetrieveJobStub = nil
	fake.retrieveJobReturns = struct {
		result1 types.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) UpdateJob(arg1 string, arg2 types.Job) (types.Job, error) {
	fake.updateJobMutex.Lock()
	fake.updateJobArgsForCall = append(fake.updateJobArgsForCall, struct {
		arg1 string
		arg2 types.Job
	}{arg1, arg2})
	fake.recordInvocation("UpdateJob", []interface{}{arg1, arg2})
	fake.updateJobMutex.Unlock()
	if fake.UpdateJobStub != nil {
		return fake.UpdateJobStub(arg1, arg2)
	} else {
		return fake.updateJobReturns.result1, fake.updateJobReturns.result2
	}
}

func (fake *FakeStorage) UpdateJobCallCount() int {
	fake.updateJobMutex.RLock()
	defer fake.updateJobMutex.RUnlock()
	return len(fake.updateJobArgsForCall)
}

func (fake *FakeStorage) UpdateJobArgsForCall(i int) (string, types.Job) {
	fake.updateJobMutex.RLock()
	defer fake.updateJobMutex.RUnlock()
	return fake.updateJobArgsForCall[i].arg1, fake.updateJobArgsForCall[i].arg2
}

func (fake *FakeStorage) UpdateJobReturns(result1 types.Job, result2 error) {
	fake.UpdateJobStub = nil
	fake.updateJobReturns = struct {
		result1 types.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) GetJobs() ([]types.Job, error) {
	fake.getJobsMutex.Lock()
	fake.getJobsArgsForCall = append(fake.getJobsArgsForCall, struct{}{})
	fake.recordInvocation("GetJobs", []interface{}{})
	fake.getJobsMutex.Unlock()
	if fake.GetJobsStub != nil {
		return fake.GetJobsStub()
	} else {
		return fake.getJobsReturns.result1, fake.getJobsReturns.result2
	}
}

func (fake *FakeStorage) GetJobsCallCount() int {
	fake.getJobsMutex.RLock()
	defer fake.getJobsMutex.RUnlock()
	return len(fake.getJobsArgsForCall)
}

func (fake *FakeStorage) GetJobsReturns(result1 []types.Job, result2 error) {
	fake.GetJobsStub = nil
	fake.getJobsReturns = struct {
		result1 []types.Job
		result2 error
	}{result1, result2}
}

func (fake *FakeStorage) ClearDatabase() error {
	fake.clearDatabaseMutex.Lock()
	fake.clearDatabaseArgsForCall = append(fake.clearDatabaseArgsForCall, struct{}{})
	fake.recordInvocation("ClearDatabase", []interface{}{})
	fake.clearDatabaseMutex.Unlock()
	if fake.ClearDatabaseStub != nil {
		return fake.ClearDatabaseStub()
	} else {
		return fake.clearDatabaseReturns.result1
	}
}

func (fake *FakeStorage) ClearDatabaseCallCount() int {
	fake.clearDatabaseMutex.RLock()
	defer fake.clearDatabaseMutex.RUnlock()
	return len(fake.clearDatabaseArgsForCall)
}

func (fake *FakeStorage) ClearDatabaseReturns(result1 error) {
	fake.ClearDatabaseStub = nil
	fake.clearDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStorage) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.storePresetMutex.RLock()
	defer fake.storePresetMutex.RUnlock()
	fake.retrievePresetMutex.RLock()
	defer fake.retrievePresetMutex.RUnlock()
	fake.updatePresetMutex.RLock()
	defer fake.updatePresetMutex.RUnlock()
	fake.getPresetsMutex.RLock()
	defer fake.getPresetsMutex.RUnlock()
	fake.deletePresetMutex.RLock()
	defer fake.deletePresetMutex.RUnlock()
	fake.storeJobMutex.RLock()
	defer fake.storeJobMutex.RUnlock()
	fake.retrieveJobMutex.RLock()
	defer fake.retrieveJobMutex.RUnlock()
	fake.updateJobMutex.RLock()
	defer fake.updateJobMutex.RUnlock()
	fake.getJobsMutex.RLock()
	defer fake.getJobsMutex.RUnlock()
	fake.clearDatabaseMutex.RLock()
	defer fake.clearDatabaseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStorage) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.Storage = new(FakeStorage)
