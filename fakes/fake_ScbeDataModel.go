// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/local/scbe"
	_ "github.com/mattn/go-sqlite3"
)

type FakeScbeDataModel struct {
	DeleteVolumeStub        func(name string) error
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		name string
	}
	deleteVolumeReturns struct {
		result1 error
	}
	deleteVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	InsertVolumeStub        func(volumeName string, wwn string, fstype string) error
	insertVolumeMutex       sync.RWMutex
	insertVolumeArgsForCall []struct {
		volumeName string
		wwn        string
		fstype     string
	}
	insertVolumeReturns struct {
		result1 error
	}
	insertVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	GetVolumeStub        func(name string) (scbe.ScbeVolume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		name string
	}
	getVolumeReturns struct {
		result1 scbe.ScbeVolume
		result2 bool
		result3 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 scbe.ScbeVolume
		result2 bool
		result3 error
	}
	ListVolumesStub        func() ([]scbe.ScbeVolume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []scbe.ScbeVolume
		result2 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 []scbe.ScbeVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScbeDataModel) DeleteVolume(name string) error {
	fake.deleteVolumeMutex.Lock()
	ret, specificReturn := fake.deleteVolumeReturnsOnCall[len(fake.deleteVolumeArgsForCall)]
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteVolume", []interface{}{name})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteVolumeReturns.result1
}

func (fake *FakeScbeDataModel) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *FakeScbeDataModel) DeleteVolumeArgsForCall(i int) string {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return fake.deleteVolumeArgsForCall[i].name
}

func (fake *FakeScbeDataModel) DeleteVolumeReturns(result1 error) {
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModel) DeleteVolumeReturnsOnCall(i int, result1 error) {
	fake.DeleteVolumeStub = nil
	if fake.deleteVolumeReturnsOnCall == nil {
		fake.deleteVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModel) InsertVolume(volumeName string, wwn string, fstype string) error {
	fake.insertVolumeMutex.Lock()
	ret, specificReturn := fake.insertVolumeReturnsOnCall[len(fake.insertVolumeArgsForCall)]
	fake.insertVolumeArgsForCall = append(fake.insertVolumeArgsForCall, struct {
		volumeName string
		wwn        string
		fstype     string
	}{volumeName, wwn, fstype})
	fake.recordInvocation("InsertVolume", []interface{}{volumeName, wwn, fstype})
	fake.insertVolumeMutex.Unlock()
	if fake.InsertVolumeStub != nil {
		return fake.InsertVolumeStub(volumeName, wwn, fstype)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertVolumeReturns.result1
}

func (fake *FakeScbeDataModel) InsertVolumeCallCount() int {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return len(fake.insertVolumeArgsForCall)
}

func (fake *FakeScbeDataModel) InsertVolumeArgsForCall(i int) (string, string, string) {
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	return fake.insertVolumeArgsForCall[i].volumeName, fake.insertVolumeArgsForCall[i].wwn, fake.insertVolumeArgsForCall[i].fstype
}

func (fake *FakeScbeDataModel) InsertVolumeReturns(result1 error) {
	fake.InsertVolumeStub = nil
	fake.insertVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModel) InsertVolumeReturnsOnCall(i int, result1 error) {
	fake.InsertVolumeStub = nil
	if fake.insertVolumeReturnsOnCall == nil {
		fake.insertVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeScbeDataModel) GetVolume(name string) (scbe.ScbeVolume, bool, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVolume", []interface{}{name})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
}

func (fake *FakeScbeDataModel) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeScbeDataModel) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].name
}

func (fake *FakeScbeDataModel) GetVolumeReturns(result1 scbe.ScbeVolume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 scbe.ScbeVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScbeDataModel) GetVolumeReturnsOnCall(i int, result1 scbe.ScbeVolume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 scbe.ScbeVolume
			result2 bool
			result3 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 scbe.ScbeVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeScbeDataModel) ListVolumes() ([]scbe.ScbeVolume, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
}

func (fake *FakeScbeDataModel) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeScbeDataModel) ListVolumesReturns(result1 []scbe.ScbeVolume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []scbe.ScbeVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeScbeDataModel) ListVolumesReturnsOnCall(i int, result1 []scbe.ScbeVolume, result2 error) {
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 []scbe.ScbeVolume
			result2 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 []scbe.ScbeVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeScbeDataModel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.insertVolumeMutex.RLock()
	defer fake.insertVolumeMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeScbeDataModel) recordInvocation(key string, args []interface{}) {
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

var _ scbe.ScbeDataModel = new(FakeScbeDataModel)
