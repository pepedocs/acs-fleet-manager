// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sso

import (
	"context"
	"github.com/stackrox/acs-fleet-manager/pkg/api"
	"github.com/stackrox/acs-fleet-manager/pkg/client/keycloak"
	"github.com/stackrox/acs-fleet-manager/pkg/errors"
	"sync"
)

// Ensure, that KeycloakServiceMock does implement KeycloakService.
// If this is not the case, regenerate this file with moq.
var _ KeycloakService = &KeycloakServiceMock{}

// KeycloakServiceMock is a mock implementation of KeycloakService.
//
// 	func TestSomethingThatUsesKeycloakService(t *testing.T) {
//
// 		// make and configure a mocked KeycloakService
// 		mockedKeycloakService := &KeycloakServiceMock{
// 			CreateServiceAccountFunc: func(serviceAccountRequest *api.ServiceAccountRequest, ctx context.Context) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the CreateServiceAccount method")
// 			},
// 			CreateServiceAccountInternalFunc: func(request CompleteServiceAccountRequest) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the CreateServiceAccountInternal method")
// 			},
// 			DeRegisterAcsFleetshardOperatorServiceAccountFunc: func(agentClusterId string) *errors.ServiceError {
// 				panic("mock out the DeRegisterAcsFleetshardOperatorServiceAccount method")
// 			},
// 			DeleteServiceAccountFunc: func(ctx context.Context, clientId string) *errors.ServiceError {
// 				panic("mock out the DeleteServiceAccount method")
// 			},
// 			DeleteServiceAccountInternalFunc: func(clientId string) *errors.ServiceError {
// 				panic("mock out the DeleteServiceAccountInternal method")
// 			},
// 			GetAcsClientSecretFunc: func(clientId string) (string, *errors.ServiceError) {
// 				panic("mock out the GetAcsClientSecret method")
// 			},
// 			GetConfigFunc: func() *keycloak.KeycloakConfig {
// 				panic("mock out the GetConfig method")
// 			},
// 			GetRealmConfigFunc: func() *keycloak.KeycloakRealmConfig {
// 				panic("mock out the GetRealmConfig method")
// 			},
// 			GetServiceAccountByClientIdFunc: func(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the GetServiceAccountByClientId method")
// 			},
// 			GetServiceAccountByIdFunc: func(ctx context.Context, id string) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the GetServiceAccountById method")
// 			},
// 			ListServiceAccFunc: func(ctx context.Context, first int, max int) ([]api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the ListServiceAcc method")
// 			},
// 			RegisterAcsFleetshardOperatorServiceAccountFunc: func(agentClusterId string) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the RegisterAcsFleetshardOperatorServiceAccount method")
// 			},
// 			ResetServiceAccountCredentialsFunc: func(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError) {
// 				panic("mock out the ResetServiceAccountCredentials method")
// 			},
// 		}
//
// 		// use mockedKeycloakService in code that requires KeycloakService
// 		// and then make assertions.
//
// 	}
type KeycloakServiceMock struct {
	// CreateServiceAccountFunc mocks the CreateServiceAccount method.
	CreateServiceAccountFunc func(serviceAccountRequest *api.ServiceAccountRequest, ctx context.Context) (*api.ServiceAccount, *errors.ServiceError)

	// CreateServiceAccountInternalFunc mocks the CreateServiceAccountInternal method.
	CreateServiceAccountInternalFunc func(request CompleteServiceAccountRequest) (*api.ServiceAccount, *errors.ServiceError)

	// DeRegisterAcsFleetshardOperatorServiceAccountFunc mocks the DeRegisterAcsFleetshardOperatorServiceAccount method.
	DeRegisterAcsFleetshardOperatorServiceAccountFunc func(agentClusterId string) *errors.ServiceError

	// DeleteServiceAccountFunc mocks the DeleteServiceAccount method.
	DeleteServiceAccountFunc func(ctx context.Context, clientId string) *errors.ServiceError

	// DeleteServiceAccountInternalFunc mocks the DeleteServiceAccountInternal method.
	DeleteServiceAccountInternalFunc func(clientId string) *errors.ServiceError

	// GetAcsClientSecretFunc mocks the GetAcsClientSecret method.
	GetAcsClientSecretFunc func(clientId string) (string, *errors.ServiceError)

	// GetConfigFunc mocks the GetConfig method.
	GetConfigFunc func() *keycloak.KeycloakConfig

	// GetRealmConfigFunc mocks the GetRealmConfig method.
	GetRealmConfigFunc func() *keycloak.KeycloakRealmConfig

	// GetServiceAccountByClientIdFunc mocks the GetServiceAccountByClientId method.
	GetServiceAccountByClientIdFunc func(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError)

	// GetServiceAccountByIdFunc mocks the GetServiceAccountById method.
	GetServiceAccountByIdFunc func(ctx context.Context, id string) (*api.ServiceAccount, *errors.ServiceError)

	// ListServiceAccFunc mocks the ListServiceAcc method.
	ListServiceAccFunc func(ctx context.Context, first int, max int) ([]api.ServiceAccount, *errors.ServiceError)

	// RegisterAcsFleetshardOperatorServiceAccountFunc mocks the RegisterAcsFleetshardOperatorServiceAccount method.
	RegisterAcsFleetshardOperatorServiceAccountFunc func(agentClusterId string) (*api.ServiceAccount, *errors.ServiceError)

	// ResetServiceAccountCredentialsFunc mocks the ResetServiceAccountCredentials method.
	ResetServiceAccountCredentialsFunc func(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError)

	// calls tracks calls to the methods.
	calls struct {
		// CreateServiceAccount holds details about calls to the CreateServiceAccount method.
		CreateServiceAccount []struct {
			// ServiceAccountRequest is the serviceAccountRequest argument value.
			ServiceAccountRequest *api.ServiceAccountRequest
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// CreateServiceAccountInternal holds details about calls to the CreateServiceAccountInternal method.
		CreateServiceAccountInternal []struct {
			// Request is the request argument value.
			Request CompleteServiceAccountRequest
		}
		// DeRegisterAcsFleetshardOperatorServiceAccount holds details about calls to the DeRegisterAcsFleetshardOperatorServiceAccount method.
		DeRegisterAcsFleetshardOperatorServiceAccount []struct {
			// AgentClusterId is the agentClusterId argument value.
			AgentClusterId string
		}
		// DeleteServiceAccount holds details about calls to the DeleteServiceAccount method.
		DeleteServiceAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ClientId is the clientId argument value.
			ClientId string
		}
		// DeleteServiceAccountInternal holds details about calls to the DeleteServiceAccountInternal method.
		DeleteServiceAccountInternal []struct {
			// ClientId is the clientId argument value.
			ClientId string
		}
		// GetAcsClientSecret holds details about calls to the GetAcsClientSecret method.
		GetAcsClientSecret []struct {
			// ClientId is the clientId argument value.
			ClientId string
		}
		// GetConfig holds details about calls to the GetConfig method.
		GetConfig []struct {
		}
		// GetRealmConfig holds details about calls to the GetRealmConfig method.
		GetRealmConfig []struct {
		}
		// GetServiceAccountByClientId holds details about calls to the GetServiceAccountByClientId method.
		GetServiceAccountByClientId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ClientId is the clientId argument value.
			ClientId string
		}
		// GetServiceAccountById holds details about calls to the GetServiceAccountById method.
		GetServiceAccountById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// ListServiceAcc holds details about calls to the ListServiceAcc method.
		ListServiceAcc []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// First is the first argument value.
			First int
			// Max is the max argument value.
			Max int
		}
		// RegisterAcsFleetshardOperatorServiceAccount holds details about calls to the RegisterAcsFleetshardOperatorServiceAccount method.
		RegisterAcsFleetshardOperatorServiceAccount []struct {
			// AgentClusterId is the agentClusterId argument value.
			AgentClusterId string
		}
		// ResetServiceAccountCredentials holds details about calls to the ResetServiceAccountCredentials method.
		ResetServiceAccountCredentials []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ClientId is the clientId argument value.
			ClientId string
		}
	}
	lockCreateServiceAccount                          sync.RWMutex
	lockCreateServiceAccountInternal                  sync.RWMutex
	lockDeRegisterAcsFleetshardOperatorServiceAccount sync.RWMutex
	lockDeleteServiceAccount                          sync.RWMutex
	lockDeleteServiceAccountInternal                  sync.RWMutex
	lockGetAcsClientSecret                            sync.RWMutex
	lockGetConfig                                     sync.RWMutex
	lockGetRealmConfig                                sync.RWMutex
	lockGetServiceAccountByClientId                   sync.RWMutex
	lockGetServiceAccountById                         sync.RWMutex
	lockListServiceAcc                                sync.RWMutex
	lockRegisterAcsFleetshardOperatorServiceAccount   sync.RWMutex
	lockResetServiceAccountCredentials                sync.RWMutex
}

// CreateServiceAccount calls CreateServiceAccountFunc.
func (mock *KeycloakServiceMock) CreateServiceAccount(serviceAccountRequest *api.ServiceAccountRequest, ctx context.Context) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.CreateServiceAccountFunc == nil {
		panic("KeycloakServiceMock.CreateServiceAccountFunc: method is nil but KeycloakService.CreateServiceAccount was just called")
	}
	callInfo := struct {
		ServiceAccountRequest *api.ServiceAccountRequest
		Ctx                   context.Context
	}{
		ServiceAccountRequest: serviceAccountRequest,
		Ctx:                   ctx,
	}
	mock.lockCreateServiceAccount.Lock()
	mock.calls.CreateServiceAccount = append(mock.calls.CreateServiceAccount, callInfo)
	mock.lockCreateServiceAccount.Unlock()
	return mock.CreateServiceAccountFunc(serviceAccountRequest, ctx)
}

// CreateServiceAccountCalls gets all the calls that were made to CreateServiceAccount.
// Check the length with:
//     len(mockedKeycloakService.CreateServiceAccountCalls())
func (mock *KeycloakServiceMock) CreateServiceAccountCalls() []struct {
	ServiceAccountRequest *api.ServiceAccountRequest
	Ctx                   context.Context
} {
	var calls []struct {
		ServiceAccountRequest *api.ServiceAccountRequest
		Ctx                   context.Context
	}
	mock.lockCreateServiceAccount.RLock()
	calls = mock.calls.CreateServiceAccount
	mock.lockCreateServiceAccount.RUnlock()
	return calls
}

// CreateServiceAccountInternal calls CreateServiceAccountInternalFunc.
func (mock *KeycloakServiceMock) CreateServiceAccountInternal(request CompleteServiceAccountRequest) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.CreateServiceAccountInternalFunc == nil {
		panic("KeycloakServiceMock.CreateServiceAccountInternalFunc: method is nil but KeycloakService.CreateServiceAccountInternal was just called")
	}
	callInfo := struct {
		Request CompleteServiceAccountRequest
	}{
		Request: request,
	}
	mock.lockCreateServiceAccountInternal.Lock()
	mock.calls.CreateServiceAccountInternal = append(mock.calls.CreateServiceAccountInternal, callInfo)
	mock.lockCreateServiceAccountInternal.Unlock()
	return mock.CreateServiceAccountInternalFunc(request)
}

// CreateServiceAccountInternalCalls gets all the calls that were made to CreateServiceAccountInternal.
// Check the length with:
//     len(mockedKeycloakService.CreateServiceAccountInternalCalls())
func (mock *KeycloakServiceMock) CreateServiceAccountInternalCalls() []struct {
	Request CompleteServiceAccountRequest
} {
	var calls []struct {
		Request CompleteServiceAccountRequest
	}
	mock.lockCreateServiceAccountInternal.RLock()
	calls = mock.calls.CreateServiceAccountInternal
	mock.lockCreateServiceAccountInternal.RUnlock()
	return calls
}

// DeRegisterAcsFleetshardOperatorServiceAccount calls DeRegisterAcsFleetshardOperatorServiceAccountFunc.
func (mock *KeycloakServiceMock) DeRegisterAcsFleetshardOperatorServiceAccount(agentClusterId string) *errors.ServiceError {
	if mock.DeRegisterAcsFleetshardOperatorServiceAccountFunc == nil {
		panic("KeycloakServiceMock.DeRegisterAcsFleetshardOperatorServiceAccountFunc: method is nil but KeycloakService.DeRegisterAcsFleetshardOperatorServiceAccount was just called")
	}
	callInfo := struct {
		AgentClusterId string
	}{
		AgentClusterId: agentClusterId,
	}
	mock.lockDeRegisterAcsFleetshardOperatorServiceAccount.Lock()
	mock.calls.DeRegisterAcsFleetshardOperatorServiceAccount = append(mock.calls.DeRegisterAcsFleetshardOperatorServiceAccount, callInfo)
	mock.lockDeRegisterAcsFleetshardOperatorServiceAccount.Unlock()
	return mock.DeRegisterAcsFleetshardOperatorServiceAccountFunc(agentClusterId)
}

// DeRegisterAcsFleetshardOperatorServiceAccountCalls gets all the calls that were made to DeRegisterAcsFleetshardOperatorServiceAccount.
// Check the length with:
//     len(mockedKeycloakService.DeRegisterAcsFleetshardOperatorServiceAccountCalls())
func (mock *KeycloakServiceMock) DeRegisterAcsFleetshardOperatorServiceAccountCalls() []struct {
	AgentClusterId string
} {
	var calls []struct {
		AgentClusterId string
	}
	mock.lockDeRegisterAcsFleetshardOperatorServiceAccount.RLock()
	calls = mock.calls.DeRegisterAcsFleetshardOperatorServiceAccount
	mock.lockDeRegisterAcsFleetshardOperatorServiceAccount.RUnlock()
	return calls
}

// DeleteServiceAccount calls DeleteServiceAccountFunc.
func (mock *KeycloakServiceMock) DeleteServiceAccount(ctx context.Context, clientId string) *errors.ServiceError {
	if mock.DeleteServiceAccountFunc == nil {
		panic("KeycloakServiceMock.DeleteServiceAccountFunc: method is nil but KeycloakService.DeleteServiceAccount was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ClientId string
	}{
		Ctx:      ctx,
		ClientId: clientId,
	}
	mock.lockDeleteServiceAccount.Lock()
	mock.calls.DeleteServiceAccount = append(mock.calls.DeleteServiceAccount, callInfo)
	mock.lockDeleteServiceAccount.Unlock()
	return mock.DeleteServiceAccountFunc(ctx, clientId)
}

// DeleteServiceAccountCalls gets all the calls that were made to DeleteServiceAccount.
// Check the length with:
//     len(mockedKeycloakService.DeleteServiceAccountCalls())
func (mock *KeycloakServiceMock) DeleteServiceAccountCalls() []struct {
	Ctx      context.Context
	ClientId string
} {
	var calls []struct {
		Ctx      context.Context
		ClientId string
	}
	mock.lockDeleteServiceAccount.RLock()
	calls = mock.calls.DeleteServiceAccount
	mock.lockDeleteServiceAccount.RUnlock()
	return calls
}

// DeleteServiceAccountInternal calls DeleteServiceAccountInternalFunc.
func (mock *KeycloakServiceMock) DeleteServiceAccountInternal(clientId string) *errors.ServiceError {
	if mock.DeleteServiceAccountInternalFunc == nil {
		panic("KeycloakServiceMock.DeleteServiceAccountInternalFunc: method is nil but KeycloakService.DeleteServiceAccountInternal was just called")
	}
	callInfo := struct {
		ClientId string
	}{
		ClientId: clientId,
	}
	mock.lockDeleteServiceAccountInternal.Lock()
	mock.calls.DeleteServiceAccountInternal = append(mock.calls.DeleteServiceAccountInternal, callInfo)
	mock.lockDeleteServiceAccountInternal.Unlock()
	return mock.DeleteServiceAccountInternalFunc(clientId)
}

// DeleteServiceAccountInternalCalls gets all the calls that were made to DeleteServiceAccountInternal.
// Check the length with:
//     len(mockedKeycloakService.DeleteServiceAccountInternalCalls())
func (mock *KeycloakServiceMock) DeleteServiceAccountInternalCalls() []struct {
	ClientId string
} {
	var calls []struct {
		ClientId string
	}
	mock.lockDeleteServiceAccountInternal.RLock()
	calls = mock.calls.DeleteServiceAccountInternal
	mock.lockDeleteServiceAccountInternal.RUnlock()
	return calls
}

// GetAcsClientSecret calls GetAcsClientSecretFunc.
func (mock *KeycloakServiceMock) GetAcsClientSecret(clientId string) (string, *errors.ServiceError) {
	if mock.GetAcsClientSecretFunc == nil {
		panic("KeycloakServiceMock.GetAcsClientSecretFunc: method is nil but KeycloakService.GetAcsClientSecret was just called")
	}
	callInfo := struct {
		ClientId string
	}{
		ClientId: clientId,
	}
	mock.lockGetAcsClientSecret.Lock()
	mock.calls.GetAcsClientSecret = append(mock.calls.GetAcsClientSecret, callInfo)
	mock.lockGetAcsClientSecret.Unlock()
	return mock.GetAcsClientSecretFunc(clientId)
}

// GetAcsClientSecretCalls gets all the calls that were made to GetAcsClientSecret.
// Check the length with:
//     len(mockedKeycloakService.GetAcsClientSecretCalls())
func (mock *KeycloakServiceMock) GetAcsClientSecretCalls() []struct {
	ClientId string
} {
	var calls []struct {
		ClientId string
	}
	mock.lockGetAcsClientSecret.RLock()
	calls = mock.calls.GetAcsClientSecret
	mock.lockGetAcsClientSecret.RUnlock()
	return calls
}

// GetConfig calls GetConfigFunc.
func (mock *KeycloakServiceMock) GetConfig() *keycloak.KeycloakConfig {
	if mock.GetConfigFunc == nil {
		panic("KeycloakServiceMock.GetConfigFunc: method is nil but KeycloakService.GetConfig was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConfig.Lock()
	mock.calls.GetConfig = append(mock.calls.GetConfig, callInfo)
	mock.lockGetConfig.Unlock()
	return mock.GetConfigFunc()
}

// GetConfigCalls gets all the calls that were made to GetConfig.
// Check the length with:
//     len(mockedKeycloakService.GetConfigCalls())
func (mock *KeycloakServiceMock) GetConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConfig.RLock()
	calls = mock.calls.GetConfig
	mock.lockGetConfig.RUnlock()
	return calls
}

// GetRealmConfig calls GetRealmConfigFunc.
func (mock *KeycloakServiceMock) GetRealmConfig() *keycloak.KeycloakRealmConfig {
	if mock.GetRealmConfigFunc == nil {
		panic("KeycloakServiceMock.GetRealmConfigFunc: method is nil but KeycloakService.GetRealmConfig was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetRealmConfig.Lock()
	mock.calls.GetRealmConfig = append(mock.calls.GetRealmConfig, callInfo)
	mock.lockGetRealmConfig.Unlock()
	return mock.GetRealmConfigFunc()
}

// GetRealmConfigCalls gets all the calls that were made to GetRealmConfig.
// Check the length with:
//     len(mockedKeycloakService.GetRealmConfigCalls())
func (mock *KeycloakServiceMock) GetRealmConfigCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetRealmConfig.RLock()
	calls = mock.calls.GetRealmConfig
	mock.lockGetRealmConfig.RUnlock()
	return calls
}

// GetServiceAccountByClientId calls GetServiceAccountByClientIdFunc.
func (mock *KeycloakServiceMock) GetServiceAccountByClientId(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.GetServiceAccountByClientIdFunc == nil {
		panic("KeycloakServiceMock.GetServiceAccountByClientIdFunc: method is nil but KeycloakService.GetServiceAccountByClientId was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ClientId string
	}{
		Ctx:      ctx,
		ClientId: clientId,
	}
	mock.lockGetServiceAccountByClientId.Lock()
	mock.calls.GetServiceAccountByClientId = append(mock.calls.GetServiceAccountByClientId, callInfo)
	mock.lockGetServiceAccountByClientId.Unlock()
	return mock.GetServiceAccountByClientIdFunc(ctx, clientId)
}

// GetServiceAccountByClientIdCalls gets all the calls that were made to GetServiceAccountByClientId.
// Check the length with:
//     len(mockedKeycloakService.GetServiceAccountByClientIdCalls())
func (mock *KeycloakServiceMock) GetServiceAccountByClientIdCalls() []struct {
	Ctx      context.Context
	ClientId string
} {
	var calls []struct {
		Ctx      context.Context
		ClientId string
	}
	mock.lockGetServiceAccountByClientId.RLock()
	calls = mock.calls.GetServiceAccountByClientId
	mock.lockGetServiceAccountByClientId.RUnlock()
	return calls
}

// GetServiceAccountById calls GetServiceAccountByIdFunc.
func (mock *KeycloakServiceMock) GetServiceAccountById(ctx context.Context, id string) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.GetServiceAccountByIdFunc == nil {
		panic("KeycloakServiceMock.GetServiceAccountByIdFunc: method is nil but KeycloakService.GetServiceAccountById was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetServiceAccountById.Lock()
	mock.calls.GetServiceAccountById = append(mock.calls.GetServiceAccountById, callInfo)
	mock.lockGetServiceAccountById.Unlock()
	return mock.GetServiceAccountByIdFunc(ctx, id)
}

// GetServiceAccountByIdCalls gets all the calls that were made to GetServiceAccountById.
// Check the length with:
//     len(mockedKeycloakService.GetServiceAccountByIdCalls())
func (mock *KeycloakServiceMock) GetServiceAccountByIdCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetServiceAccountById.RLock()
	calls = mock.calls.GetServiceAccountById
	mock.lockGetServiceAccountById.RUnlock()
	return calls
}

// ListServiceAcc calls ListServiceAccFunc.
func (mock *KeycloakServiceMock) ListServiceAcc(ctx context.Context, first int, max int) ([]api.ServiceAccount, *errors.ServiceError) {
	if mock.ListServiceAccFunc == nil {
		panic("KeycloakServiceMock.ListServiceAccFunc: method is nil but KeycloakService.ListServiceAcc was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		First int
		Max   int
	}{
		Ctx:   ctx,
		First: first,
		Max:   max,
	}
	mock.lockListServiceAcc.Lock()
	mock.calls.ListServiceAcc = append(mock.calls.ListServiceAcc, callInfo)
	mock.lockListServiceAcc.Unlock()
	return mock.ListServiceAccFunc(ctx, first, max)
}

// ListServiceAccCalls gets all the calls that were made to ListServiceAcc.
// Check the length with:
//     len(mockedKeycloakService.ListServiceAccCalls())
func (mock *KeycloakServiceMock) ListServiceAccCalls() []struct {
	Ctx   context.Context
	First int
	Max   int
} {
	var calls []struct {
		Ctx   context.Context
		First int
		Max   int
	}
	mock.lockListServiceAcc.RLock()
	calls = mock.calls.ListServiceAcc
	mock.lockListServiceAcc.RUnlock()
	return calls
}

// RegisterAcsFleetshardOperatorServiceAccount calls RegisterAcsFleetshardOperatorServiceAccountFunc.
func (mock *KeycloakServiceMock) RegisterAcsFleetshardOperatorServiceAccount(agentClusterId string) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.RegisterAcsFleetshardOperatorServiceAccountFunc == nil {
		panic("KeycloakServiceMock.RegisterAcsFleetshardOperatorServiceAccountFunc: method is nil but KeycloakService.RegisterAcsFleetshardOperatorServiceAccount was just called")
	}
	callInfo := struct {
		AgentClusterId string
	}{
		AgentClusterId: agentClusterId,
	}
	mock.lockRegisterAcsFleetshardOperatorServiceAccount.Lock()
	mock.calls.RegisterAcsFleetshardOperatorServiceAccount = append(mock.calls.RegisterAcsFleetshardOperatorServiceAccount, callInfo)
	mock.lockRegisterAcsFleetshardOperatorServiceAccount.Unlock()
	return mock.RegisterAcsFleetshardOperatorServiceAccountFunc(agentClusterId)
}

// RegisterAcsFleetshardOperatorServiceAccountCalls gets all the calls that were made to RegisterAcsFleetshardOperatorServiceAccount.
// Check the length with:
//     len(mockedKeycloakService.RegisterAcsFleetshardOperatorServiceAccountCalls())
func (mock *KeycloakServiceMock) RegisterAcsFleetshardOperatorServiceAccountCalls() []struct {
	AgentClusterId string
} {
	var calls []struct {
		AgentClusterId string
	}
	mock.lockRegisterAcsFleetshardOperatorServiceAccount.RLock()
	calls = mock.calls.RegisterAcsFleetshardOperatorServiceAccount
	mock.lockRegisterAcsFleetshardOperatorServiceAccount.RUnlock()
	return calls
}

// ResetServiceAccountCredentials calls ResetServiceAccountCredentialsFunc.
func (mock *KeycloakServiceMock) ResetServiceAccountCredentials(ctx context.Context, clientId string) (*api.ServiceAccount, *errors.ServiceError) {
	if mock.ResetServiceAccountCredentialsFunc == nil {
		panic("KeycloakServiceMock.ResetServiceAccountCredentialsFunc: method is nil but KeycloakService.ResetServiceAccountCredentials was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		ClientId string
	}{
		Ctx:      ctx,
		ClientId: clientId,
	}
	mock.lockResetServiceAccountCredentials.Lock()
	mock.calls.ResetServiceAccountCredentials = append(mock.calls.ResetServiceAccountCredentials, callInfo)
	mock.lockResetServiceAccountCredentials.Unlock()
	return mock.ResetServiceAccountCredentialsFunc(ctx, clientId)
}

// ResetServiceAccountCredentialsCalls gets all the calls that were made to ResetServiceAccountCredentials.
// Check the length with:
//     len(mockedKeycloakService.ResetServiceAccountCredentialsCalls())
func (mock *KeycloakServiceMock) ResetServiceAccountCredentialsCalls() []struct {
	Ctx      context.Context
	ClientId string
} {
	var calls []struct {
		Ctx      context.Context
		ClientId string
	}
	mock.lockResetServiceAccountCredentials.RLock()
	calls = mock.calls.ResetServiceAccountCredentials
	mock.lockResetServiceAccountCredentials.RUnlock()
	return calls
}