// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/common/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/lifecycle"
	persistence "github.com/hyperledger/fabric/core/chaincode/persistence/intf"
)

type SCCFunctions struct {
	ApproveChaincodeDefinitionForOrgStub        func(string, string, *lifecycle.ChaincodeDefinition, persistence.PackageID, lifecycle.ReadableState, lifecycle.ReadWritableState) error
	approveChaincodeDefinitionForOrgMutex       sync.RWMutex
	approveChaincodeDefinitionForOrgArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 persistence.PackageID
		arg5 lifecycle.ReadableState
		arg6 lifecycle.ReadWritableState
	}
	approveChaincodeDefinitionForOrgReturns struct {
		result1 error
	}
	approveChaincodeDefinitionForOrgReturnsOnCall map[int]struct {
		result1 error
	}
	CommitChaincodeDefinitionStub        func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) ([]bool, error)
	commitChaincodeDefinitionMutex       sync.RWMutex
	commitChaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}
	commitChaincodeDefinitionReturns struct {
		result1 []bool
		result2 error
	}
	commitChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 []bool
		result2 error
	}
	InstallChaincodeStub        func([]byte) (*chaincode.InstalledChaincode, error)
	installChaincodeMutex       sync.RWMutex
	installChaincodeArgsForCall []struct {
		arg1 []byte
	}
	installChaincodeReturns struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	installChaincodeReturnsOnCall map[int]struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	QueryApprovalStatusStub        func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) ([]bool, error)
	queryApprovalStatusMutex       sync.RWMutex
	queryApprovalStatusArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}
	queryApprovalStatusReturns struct {
		result1 []bool
		result2 error
	}
	queryApprovalStatusReturnsOnCall map[int]struct {
		result1 []bool
		result2 error
	}
	QueryChaincodeDefinitionStub        func(string, lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error)
	queryChaincodeDefinitionMutex       sync.RWMutex
	queryChaincodeDefinitionArgsForCall []struct {
		arg1 string
		arg2 lifecycle.ReadableState
	}
	queryChaincodeDefinitionReturns struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}
	queryChaincodeDefinitionReturnsOnCall map[int]struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}
	QueryInstalledChaincodeStub        func(persistence.PackageID) (*chaincode.InstalledChaincode, error)
	queryInstalledChaincodeMutex       sync.RWMutex
	queryInstalledChaincodeArgsForCall []struct {
		arg1 persistence.PackageID
	}
	queryInstalledChaincodeReturns struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	queryInstalledChaincodeReturnsOnCall map[int]struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}
	QueryInstalledChaincodesStub        func() ([]chaincode.InstalledChaincode, error)
	queryInstalledChaincodesMutex       sync.RWMutex
	queryInstalledChaincodesArgsForCall []struct {
	}
	queryInstalledChaincodesReturns struct {
		result1 []chaincode.InstalledChaincode
		result2 error
	}
	queryInstalledChaincodesReturnsOnCall map[int]struct {
		result1 []chaincode.InstalledChaincode
		result2 error
	}
	QueryNamespaceDefinitionsStub        func(lifecycle.RangeableState) (map[string]string, error)
	queryNamespaceDefinitionsMutex       sync.RWMutex
	queryNamespaceDefinitionsArgsForCall []struct {
		arg1 lifecycle.RangeableState
	}
	queryNamespaceDefinitionsReturns struct {
		result1 map[string]string
		result2 error
	}
	queryNamespaceDefinitionsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrg(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 persistence.PackageID, arg5 lifecycle.ReadableState, arg6 lifecycle.ReadWritableState) error {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	ret, specificReturn := fake.approveChaincodeDefinitionForOrgReturnsOnCall[len(fake.approveChaincodeDefinitionForOrgArgsForCall)]
	fake.approveChaincodeDefinitionForOrgArgsForCall = append(fake.approveChaincodeDefinitionForOrgArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 persistence.PackageID
		arg5 lifecycle.ReadableState
		arg6 lifecycle.ReadWritableState
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("ApproveChaincodeDefinitionForOrg", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	if fake.ApproveChaincodeDefinitionForOrgStub != nil {
		return fake.ApproveChaincodeDefinitionForOrgStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.approveChaincodeDefinitionForOrgReturns
	return fakeReturns.result1
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgCallCount() int {
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	return len(fake.approveChaincodeDefinitionForOrgArgsForCall)
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, persistence.PackageID, lifecycle.ReadableState, lifecycle.ReadWritableState) error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = stub
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, persistence.PackageID, lifecycle.ReadableState, lifecycle.ReadWritableState) {
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	argsForCall := fake.approveChaincodeDefinitionForOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgReturns(result1 error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = nil
	fake.approveChaincodeDefinitionForOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *SCCFunctions) ApproveChaincodeDefinitionForOrgReturnsOnCall(i int, result1 error) {
	fake.approveChaincodeDefinitionForOrgMutex.Lock()
	defer fake.approveChaincodeDefinitionForOrgMutex.Unlock()
	fake.ApproveChaincodeDefinitionForOrgStub = nil
	if fake.approveChaincodeDefinitionForOrgReturnsOnCall == nil {
		fake.approveChaincodeDefinitionForOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.approveChaincodeDefinitionForOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SCCFunctions) CommitChaincodeDefinition(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 lifecycle.ReadWritableState, arg5 []lifecycle.OpaqueState) ([]bool, error) {
	var arg5Copy []lifecycle.OpaqueState
	if arg5 != nil {
		arg5Copy = make([]lifecycle.OpaqueState, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.commitChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.commitChaincodeDefinitionReturnsOnCall[len(fake.commitChaincodeDefinitionArgsForCall)]
	fake.commitChaincodeDefinitionArgsForCall = append(fake.commitChaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.recordInvocation("CommitChaincodeDefinition", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.commitChaincodeDefinitionMutex.Unlock()
	if fake.CommitChaincodeDefinitionStub != nil {
		return fake.CommitChaincodeDefinitionStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.commitChaincodeDefinitionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) CommitChaincodeDefinitionCallCount() int {
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	return len(fake.commitChaincodeDefinitionArgsForCall)
}

func (fake *SCCFunctions) CommitChaincodeDefinitionCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) ([]bool, error)) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = stub
}

func (fake *SCCFunctions) CommitChaincodeDefinitionArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) {
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.commitChaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SCCFunctions) CommitChaincodeDefinitionReturns(result1 []bool, result2 error) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = nil
	fake.commitChaincodeDefinitionReturns = struct {
		result1 []bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) CommitChaincodeDefinitionReturnsOnCall(i int, result1 []bool, result2 error) {
	fake.commitChaincodeDefinitionMutex.Lock()
	defer fake.commitChaincodeDefinitionMutex.Unlock()
	fake.CommitChaincodeDefinitionStub = nil
	if fake.commitChaincodeDefinitionReturnsOnCall == nil {
		fake.commitChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 []bool
			result2 error
		})
	}
	fake.commitChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 []bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) InstallChaincode(arg1 []byte) (*chaincode.InstalledChaincode, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.installChaincodeMutex.Lock()
	ret, specificReturn := fake.installChaincodeReturnsOnCall[len(fake.installChaincodeArgsForCall)]
	fake.installChaincodeArgsForCall = append(fake.installChaincodeArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("InstallChaincode", []interface{}{arg1Copy})
	fake.installChaincodeMutex.Unlock()
	if fake.InstallChaincodeStub != nil {
		return fake.InstallChaincodeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.installChaincodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) InstallChaincodeCallCount() int {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	return len(fake.installChaincodeArgsForCall)
}

func (fake *SCCFunctions) InstallChaincodeCalls(stub func([]byte) (*chaincode.InstalledChaincode, error)) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = stub
}

func (fake *SCCFunctions) InstallChaincodeArgsForCall(i int) []byte {
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	argsForCall := fake.installChaincodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) InstallChaincodeReturns(result1 *chaincode.InstalledChaincode, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	fake.installChaincodeReturns = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) InstallChaincodeReturnsOnCall(i int, result1 *chaincode.InstalledChaincode, result2 error) {
	fake.installChaincodeMutex.Lock()
	defer fake.installChaincodeMutex.Unlock()
	fake.InstallChaincodeStub = nil
	if fake.installChaincodeReturnsOnCall == nil {
		fake.installChaincodeReturnsOnCall = make(map[int]struct {
			result1 *chaincode.InstalledChaincode
			result2 error
		})
	}
	fake.installChaincodeReturnsOnCall[i] = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryApprovalStatus(arg1 string, arg2 string, arg3 *lifecycle.ChaincodeDefinition, arg4 lifecycle.ReadWritableState, arg5 []lifecycle.OpaqueState) ([]bool, error) {
	var arg5Copy []lifecycle.OpaqueState
	if arg5 != nil {
		arg5Copy = make([]lifecycle.OpaqueState, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.queryApprovalStatusMutex.Lock()
	ret, specificReturn := fake.queryApprovalStatusReturnsOnCall[len(fake.queryApprovalStatusArgsForCall)]
	fake.queryApprovalStatusArgsForCall = append(fake.queryApprovalStatusArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 *lifecycle.ChaincodeDefinition
		arg4 lifecycle.ReadWritableState
		arg5 []lifecycle.OpaqueState
	}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.recordInvocation("QueryApprovalStatus", []interface{}{arg1, arg2, arg3, arg4, arg5Copy})
	fake.queryApprovalStatusMutex.Unlock()
	if fake.QueryApprovalStatusStub != nil {
		return fake.QueryApprovalStatusStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryApprovalStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryApprovalStatusCallCount() int {
	fake.queryApprovalStatusMutex.RLock()
	defer fake.queryApprovalStatusMutex.RUnlock()
	return len(fake.queryApprovalStatusArgsForCall)
}

func (fake *SCCFunctions) QueryApprovalStatusCalls(stub func(string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) ([]bool, error)) {
	fake.queryApprovalStatusMutex.Lock()
	defer fake.queryApprovalStatusMutex.Unlock()
	fake.QueryApprovalStatusStub = stub
}

func (fake *SCCFunctions) QueryApprovalStatusArgsForCall(i int) (string, string, *lifecycle.ChaincodeDefinition, lifecycle.ReadWritableState, []lifecycle.OpaqueState) {
	fake.queryApprovalStatusMutex.RLock()
	defer fake.queryApprovalStatusMutex.RUnlock()
	argsForCall := fake.queryApprovalStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SCCFunctions) QueryApprovalStatusReturns(result1 []bool, result2 error) {
	fake.queryApprovalStatusMutex.Lock()
	defer fake.queryApprovalStatusMutex.Unlock()
	fake.QueryApprovalStatusStub = nil
	fake.queryApprovalStatusReturns = struct {
		result1 []bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryApprovalStatusReturnsOnCall(i int, result1 []bool, result2 error) {
	fake.queryApprovalStatusMutex.Lock()
	defer fake.queryApprovalStatusMutex.Unlock()
	fake.QueryApprovalStatusStub = nil
	if fake.queryApprovalStatusReturnsOnCall == nil {
		fake.queryApprovalStatusReturnsOnCall = make(map[int]struct {
			result1 []bool
			result2 error
		})
	}
	fake.queryApprovalStatusReturnsOnCall[i] = struct {
		result1 []bool
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryChaincodeDefinition(arg1 string, arg2 lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.queryChaincodeDefinitionReturnsOnCall[len(fake.queryChaincodeDefinitionArgsForCall)]
	fake.queryChaincodeDefinitionArgsForCall = append(fake.queryChaincodeDefinitionArgsForCall, struct {
		arg1 string
		arg2 lifecycle.ReadableState
	}{arg1, arg2})
	fake.recordInvocation("QueryChaincodeDefinition", []interface{}{arg1, arg2})
	fake.queryChaincodeDefinitionMutex.Unlock()
	if fake.QueryChaincodeDefinitionStub != nil {
		return fake.QueryChaincodeDefinitionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryChaincodeDefinitionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryChaincodeDefinitionCallCount() int {
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	return len(fake.queryChaincodeDefinitionArgsForCall)
}

func (fake *SCCFunctions) QueryChaincodeDefinitionCalls(stub func(string, lifecycle.ReadableState) (*lifecycle.ChaincodeDefinition, error)) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = stub
}

func (fake *SCCFunctions) QueryChaincodeDefinitionArgsForCall(i int) (string, lifecycle.ReadableState) {
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	argsForCall := fake.queryChaincodeDefinitionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *SCCFunctions) QueryChaincodeDefinitionReturns(result1 *lifecycle.ChaincodeDefinition, result2 error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = nil
	fake.queryChaincodeDefinitionReturns = struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryChaincodeDefinitionReturnsOnCall(i int, result1 *lifecycle.ChaincodeDefinition, result2 error) {
	fake.queryChaincodeDefinitionMutex.Lock()
	defer fake.queryChaincodeDefinitionMutex.Unlock()
	fake.QueryChaincodeDefinitionStub = nil
	if fake.queryChaincodeDefinitionReturnsOnCall == nil {
		fake.queryChaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 *lifecycle.ChaincodeDefinition
			result2 error
		})
	}
	fake.queryChaincodeDefinitionReturnsOnCall[i] = struct {
		result1 *lifecycle.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincode(arg1 persistence.PackageID) (*chaincode.InstalledChaincode, error) {
	fake.queryInstalledChaincodeMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodeReturnsOnCall[len(fake.queryInstalledChaincodeArgsForCall)]
	fake.queryInstalledChaincodeArgsForCall = append(fake.queryInstalledChaincodeArgsForCall, struct {
		arg1 persistence.PackageID
	}{arg1})
	fake.recordInvocation("QueryInstalledChaincode", []interface{}{arg1})
	fake.queryInstalledChaincodeMutex.Unlock()
	if fake.QueryInstalledChaincodeStub != nil {
		return fake.QueryInstalledChaincodeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInstalledChaincodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryInstalledChaincodeCallCount() int {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	return len(fake.queryInstalledChaincodeArgsForCall)
}

func (fake *SCCFunctions) QueryInstalledChaincodeCalls(stub func(persistence.PackageID) (*chaincode.InstalledChaincode, error)) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = stub
}

func (fake *SCCFunctions) QueryInstalledChaincodeArgsForCall(i int) persistence.PackageID {
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	argsForCall := fake.queryInstalledChaincodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturns(result1 *chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	fake.queryInstalledChaincodeReturns = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodeReturnsOnCall(i int, result1 *chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodeMutex.Lock()
	defer fake.queryInstalledChaincodeMutex.Unlock()
	fake.QueryInstalledChaincodeStub = nil
	if fake.queryInstalledChaincodeReturnsOnCall == nil {
		fake.queryInstalledChaincodeReturnsOnCall = make(map[int]struct {
			result1 *chaincode.InstalledChaincode
			result2 error
		})
	}
	fake.queryInstalledChaincodeReturnsOnCall[i] = struct {
		result1 *chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodes() ([]chaincode.InstalledChaincode, error) {
	fake.queryInstalledChaincodesMutex.Lock()
	ret, specificReturn := fake.queryInstalledChaincodesReturnsOnCall[len(fake.queryInstalledChaincodesArgsForCall)]
	fake.queryInstalledChaincodesArgsForCall = append(fake.queryInstalledChaincodesArgsForCall, struct {
	}{})
	fake.recordInvocation("QueryInstalledChaincodes", []interface{}{})
	fake.queryInstalledChaincodesMutex.Unlock()
	if fake.QueryInstalledChaincodesStub != nil {
		return fake.QueryInstalledChaincodesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryInstalledChaincodesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryInstalledChaincodesCallCount() int {
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	return len(fake.queryInstalledChaincodesArgsForCall)
}

func (fake *SCCFunctions) QueryInstalledChaincodesCalls(stub func() ([]chaincode.InstalledChaincode, error)) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = stub
}

func (fake *SCCFunctions) QueryInstalledChaincodesReturns(result1 []chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	fake.queryInstalledChaincodesReturns = struct {
		result1 []chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryInstalledChaincodesReturnsOnCall(i int, result1 []chaincode.InstalledChaincode, result2 error) {
	fake.queryInstalledChaincodesMutex.Lock()
	defer fake.queryInstalledChaincodesMutex.Unlock()
	fake.QueryInstalledChaincodesStub = nil
	if fake.queryInstalledChaincodesReturnsOnCall == nil {
		fake.queryInstalledChaincodesReturnsOnCall = make(map[int]struct {
			result1 []chaincode.InstalledChaincode
			result2 error
		})
	}
	fake.queryInstalledChaincodesReturnsOnCall[i] = struct {
		result1 []chaincode.InstalledChaincode
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryNamespaceDefinitions(arg1 lifecycle.RangeableState) (map[string]string, error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	ret, specificReturn := fake.queryNamespaceDefinitionsReturnsOnCall[len(fake.queryNamespaceDefinitionsArgsForCall)]
	fake.queryNamespaceDefinitionsArgsForCall = append(fake.queryNamespaceDefinitionsArgsForCall, struct {
		arg1 lifecycle.RangeableState
	}{arg1})
	fake.recordInvocation("QueryNamespaceDefinitions", []interface{}{arg1})
	fake.queryNamespaceDefinitionsMutex.Unlock()
	if fake.QueryNamespaceDefinitionsStub != nil {
		return fake.QueryNamespaceDefinitionsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.queryNamespaceDefinitionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsCallCount() int {
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	return len(fake.queryNamespaceDefinitionsArgsForCall)
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsCalls(stub func(lifecycle.RangeableState) (map[string]string, error)) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = stub
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsArgsForCall(i int) lifecycle.RangeableState {
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	argsForCall := fake.queryNamespaceDefinitionsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsReturns(result1 map[string]string, result2 error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = nil
	fake.queryNamespaceDefinitionsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) QueryNamespaceDefinitionsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.queryNamespaceDefinitionsMutex.Lock()
	defer fake.queryNamespaceDefinitionsMutex.Unlock()
	fake.QueryNamespaceDefinitionsStub = nil
	if fake.queryNamespaceDefinitionsReturnsOnCall == nil {
		fake.queryNamespaceDefinitionsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.queryNamespaceDefinitionsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *SCCFunctions) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.approveChaincodeDefinitionForOrgMutex.RLock()
	defer fake.approveChaincodeDefinitionForOrgMutex.RUnlock()
	fake.commitChaincodeDefinitionMutex.RLock()
	defer fake.commitChaincodeDefinitionMutex.RUnlock()
	fake.installChaincodeMutex.RLock()
	defer fake.installChaincodeMutex.RUnlock()
	fake.queryApprovalStatusMutex.RLock()
	defer fake.queryApprovalStatusMutex.RUnlock()
	fake.queryChaincodeDefinitionMutex.RLock()
	defer fake.queryChaincodeDefinitionMutex.RUnlock()
	fake.queryInstalledChaincodeMutex.RLock()
	defer fake.queryInstalledChaincodeMutex.RUnlock()
	fake.queryInstalledChaincodesMutex.RLock()
	defer fake.queryInstalledChaincodesMutex.RUnlock()
	fake.queryNamespaceDefinitionsMutex.RLock()
	defer fake.queryNamespaceDefinitionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SCCFunctions) recordInvocation(key string, args []interface{}) {
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