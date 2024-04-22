// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/authentication (interfaces: LDAPClientDialer)
//
// Generated by this command:
//
//	mockgen -package authentication -destination ldap_client_dialer_test.go -mock_names LDAPClientDialer=MockLDAPClientDialer github.com/authelia/authelia/v4/internal/authentication LDAPClientDialer
//

// Package authentication is a generated GoMock package.
package authentication

import (
	reflect "reflect"

	ldap "github.com/go-ldap/ldap/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockLDAPClientDialer is a mock of LDAPClientDialer interface.
type MockLDAPClientDialer struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPClientDialerMockRecorder
}

// MockLDAPClientDialerMockRecorder is the mock recorder for MockLDAPClientDialer.
type MockLDAPClientDialerMockRecorder struct {
	mock *MockLDAPClientDialer
}

// NewMockLDAPClientDialer creates a new mock instance.
func NewMockLDAPClientDialer(ctrl *gomock.Controller) *MockLDAPClientDialer {
	mock := &MockLDAPClientDialer{ctrl: ctrl}
	mock.recorder = &MockLDAPClientDialerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPClientDialer) EXPECT() *MockLDAPClientDialerMockRecorder {
	return m.recorder
}

// DialURL mocks base method.
func (m *MockLDAPClientDialer) DialURL(arg0 string, arg1 ...ldap.DialOpt) (ldap.Client, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DialURL", varargs...)
	ret0, _ := ret[0].(ldap.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialURL indicates an expected call of DialURL.
func (mr *MockLDAPClientDialerMockRecorder) DialURL(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialURL", reflect.TypeOf((*MockLDAPClientDialer)(nil).DialURL), varargs...)
}
