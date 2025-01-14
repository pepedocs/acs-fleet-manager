// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cloudprovider

import (
	"context"
	"github.com/stackrox/acs-fleet-manager/fleetshard/pkg/central/postgres"
	"sync"
)

// Ensure, that DBClientMock does implement DBClient.
// If this is not the case, regenerate this file with moq.
var _ DBClient = &DBClientMock{}

// DBClientMock is a mock implementation of DBClient.
//
//	func TestSomethingThatUsesDBClient(t *testing.T) {
//
//		// make and configure a mocked DBClient
//		mockedDBClient := &DBClientMock{
//			EnsureDBDeprovisionedFunc: func(databaseID string, skipFinalSnapshot bool) error {
//				panic("mock out the EnsureDBDeprovisioned method")
//			},
//			EnsureDBProvisionedFunc: func(ctx context.Context, databaseID string, passwordSecretName string, isTestInstance bool) error {
//				panic("mock out the EnsureDBProvisioned method")
//			},
//			GetAccountQuotasFunc: func(ctx context.Context) (AccountQuotas, error) {
//				panic("mock out the GetAccountQuotas method")
//			},
//			GetDBConnectionFunc: func(databaseID string) (postgres.DBConnection, error) {
//				panic("mock out the GetDBConnection method")
//			},
//		}
//
//		// use mockedDBClient in code that requires DBClient
//		// and then make assertions.
//
//	}
type DBClientMock struct {
	// EnsureDBDeprovisionedFunc mocks the EnsureDBDeprovisioned method.
	EnsureDBDeprovisionedFunc func(databaseID string, skipFinalSnapshot bool) error

	// EnsureDBProvisionedFunc mocks the EnsureDBProvisioned method.
	EnsureDBProvisionedFunc func(ctx context.Context, databaseID string, passwordSecretName string, isTestInstance bool) error

	// GetAccountQuotasFunc mocks the GetAccountQuotas method.
	GetAccountQuotasFunc func(ctx context.Context) (AccountQuotas, error)

	// GetDBConnectionFunc mocks the GetDBConnection method.
	GetDBConnectionFunc func(databaseID string) (postgres.DBConnection, error)

	// calls tracks calls to the methods.
	calls struct {
		// EnsureDBDeprovisioned holds details about calls to the EnsureDBDeprovisioned method.
		EnsureDBDeprovisioned []struct {
			// DatabaseID is the databaseID argument value.
			DatabaseID string
			// SkipFinalSnapshot is the skipFinalSnapshot argument value.
			SkipFinalSnapshot bool
		}
		// EnsureDBProvisioned holds details about calls to the EnsureDBProvisioned method.
		EnsureDBProvisioned []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DatabaseID is the databaseID argument value.
			DatabaseID string
			// PasswordSecretName is the passwordSecretName argument value.
			PasswordSecretName string
			// IsTestInstance is the isTestInstance argument value.
			IsTestInstance bool
		}
		// GetAccountQuotas holds details about calls to the GetAccountQuotas method.
		GetAccountQuotas []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetDBConnection holds details about calls to the GetDBConnection method.
		GetDBConnection []struct {
			// DatabaseID is the databaseID argument value.
			DatabaseID string
		}
	}
	lockEnsureDBDeprovisioned sync.RWMutex
	lockEnsureDBProvisioned   sync.RWMutex
	lockGetAccountQuotas      sync.RWMutex
	lockGetDBConnection       sync.RWMutex
}

// EnsureDBDeprovisioned calls EnsureDBDeprovisionedFunc.
func (mock *DBClientMock) EnsureDBDeprovisioned(databaseID string, skipFinalSnapshot bool) error {
	if mock.EnsureDBDeprovisionedFunc == nil {
		panic("DBClientMock.EnsureDBDeprovisionedFunc: method is nil but DBClient.EnsureDBDeprovisioned was just called")
	}
	callInfo := struct {
		DatabaseID        string
		SkipFinalSnapshot bool
	}{
		DatabaseID:        databaseID,
		SkipFinalSnapshot: skipFinalSnapshot,
	}
	mock.lockEnsureDBDeprovisioned.Lock()
	mock.calls.EnsureDBDeprovisioned = append(mock.calls.EnsureDBDeprovisioned, callInfo)
	mock.lockEnsureDBDeprovisioned.Unlock()
	return mock.EnsureDBDeprovisionedFunc(databaseID, skipFinalSnapshot)
}

// EnsureDBDeprovisionedCalls gets all the calls that were made to EnsureDBDeprovisioned.
// Check the length with:
//
//	len(mockedDBClient.EnsureDBDeprovisionedCalls())
func (mock *DBClientMock) EnsureDBDeprovisionedCalls() []struct {
	DatabaseID        string
	SkipFinalSnapshot bool
} {
	var calls []struct {
		DatabaseID        string
		SkipFinalSnapshot bool
	}
	mock.lockEnsureDBDeprovisioned.RLock()
	calls = mock.calls.EnsureDBDeprovisioned
	mock.lockEnsureDBDeprovisioned.RUnlock()
	return calls
}

// EnsureDBProvisioned calls EnsureDBProvisionedFunc.
func (mock *DBClientMock) EnsureDBProvisioned(ctx context.Context, databaseID string, passwordSecretName string, isTestInstance bool) error {
	if mock.EnsureDBProvisionedFunc == nil {
		panic("DBClientMock.EnsureDBProvisionedFunc: method is nil but DBClient.EnsureDBProvisioned was just called")
	}
	callInfo := struct {
		Ctx                context.Context
		DatabaseID         string
		PasswordSecretName string
		IsTestInstance     bool
	}{
		Ctx:                ctx,
		DatabaseID:         databaseID,
		PasswordSecretName: passwordSecretName,
		IsTestInstance:     isTestInstance,
	}
	mock.lockEnsureDBProvisioned.Lock()
	mock.calls.EnsureDBProvisioned = append(mock.calls.EnsureDBProvisioned, callInfo)
	mock.lockEnsureDBProvisioned.Unlock()
	return mock.EnsureDBProvisionedFunc(ctx, databaseID, passwordSecretName, isTestInstance)
}

// EnsureDBProvisionedCalls gets all the calls that were made to EnsureDBProvisioned.
// Check the length with:
//
//	len(mockedDBClient.EnsureDBProvisionedCalls())
func (mock *DBClientMock) EnsureDBProvisionedCalls() []struct {
	Ctx                context.Context
	DatabaseID         string
	PasswordSecretName string
	IsTestInstance     bool
} {
	var calls []struct {
		Ctx                context.Context
		DatabaseID         string
		PasswordSecretName string
		IsTestInstance     bool
	}
	mock.lockEnsureDBProvisioned.RLock()
	calls = mock.calls.EnsureDBProvisioned
	mock.lockEnsureDBProvisioned.RUnlock()
	return calls
}

// GetAccountQuotas calls GetAccountQuotasFunc.
func (mock *DBClientMock) GetAccountQuotas(ctx context.Context) (AccountQuotas, error) {
	if mock.GetAccountQuotasFunc == nil {
		panic("DBClientMock.GetAccountQuotasFunc: method is nil but DBClient.GetAccountQuotas was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAccountQuotas.Lock()
	mock.calls.GetAccountQuotas = append(mock.calls.GetAccountQuotas, callInfo)
	mock.lockGetAccountQuotas.Unlock()
	return mock.GetAccountQuotasFunc(ctx)
}

// GetAccountQuotasCalls gets all the calls that were made to GetAccountQuotas.
// Check the length with:
//
//	len(mockedDBClient.GetAccountQuotasCalls())
func (mock *DBClientMock) GetAccountQuotasCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAccountQuotas.RLock()
	calls = mock.calls.GetAccountQuotas
	mock.lockGetAccountQuotas.RUnlock()
	return calls
}

// GetDBConnection calls GetDBConnectionFunc.
func (mock *DBClientMock) GetDBConnection(databaseID string) (postgres.DBConnection, error) {
	if mock.GetDBConnectionFunc == nil {
		panic("DBClientMock.GetDBConnectionFunc: method is nil but DBClient.GetDBConnection was just called")
	}
	callInfo := struct {
		DatabaseID string
	}{
		DatabaseID: databaseID,
	}
	mock.lockGetDBConnection.Lock()
	mock.calls.GetDBConnection = append(mock.calls.GetDBConnection, callInfo)
	mock.lockGetDBConnection.Unlock()
	return mock.GetDBConnectionFunc(databaseID)
}

// GetDBConnectionCalls gets all the calls that were made to GetDBConnection.
// Check the length with:
//
//	len(mockedDBClient.GetDBConnectionCalls())
func (mock *DBClientMock) GetDBConnectionCalls() []struct {
	DatabaseID string
} {
	var calls []struct {
		DatabaseID string
	}
	mock.lockGetDBConnection.RLock()
	calls = mock.calls.GetDBConnection
	mock.lockGetDBConnection.RUnlock()
	return calls
}
