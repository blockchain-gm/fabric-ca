// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	amcl "github.com/hyperledger/fabric-amcl/amcl"
	FP256BN "github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	sm2 "github.com/tjfoc/gmsm/sm2"

	idemix "github.com/hyperledger/fabric/idemix"

	mock "github.com/stretchr/testify/mock"
)

// Lib is an autogenerated mock type for the Lib type
type Lib struct {
	mock.Mock
}

// CreateCRI provides a mock function with given fields: key, unrevokedHandles, epoch, alg, rng
func (_m *Lib) CreateCRI(key *sm2.PrivateKey, unrevokedHandles []*FP256BN.BIG, epoch int, alg idemix.RevocationAlgorithm, rng *amcl.RAND) (*idemix.CredentialRevocationInformation, error) {
	ret := _m.Called(key, unrevokedHandles, epoch, alg, rng)

	var r0 *idemix.CredentialRevocationInformation
	if rf, ok := ret.Get(0).(func(*sm2.PrivateKey, []*FP256BN.BIG, int, idemix.RevocationAlgorithm, *amcl.RAND) *idemix.CredentialRevocationInformation); ok {
		r0 = rf(key, unrevokedHandles, epoch, alg, rng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.CredentialRevocationInformation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*sm2.PrivateKey, []*FP256BN.BIG, int, idemix.RevocationAlgorithm, *amcl.RAND) error); ok {
		r1 = rf(key, unrevokedHandles, epoch, alg, rng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateLongTermRevocationKey provides a mock function with given fields:
func (_m *Lib) GenerateLongTermRevocationKey() (*sm2.PrivateKey, error) {
	ret := _m.Called()

	var r0 *sm2.PrivateKey
	if rf, ok := ret.Get(0).(func() *sm2.PrivateKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sm2.PrivateKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRand provides a mock function with given fields:
func (_m *Lib) GetRand() (*amcl.RAND, error) {
	ret := _m.Called()

	var r0 *amcl.RAND
	if rf, ok := ret.Get(0).(func() *amcl.RAND); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amcl.RAND)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCredential provides a mock function with given fields: key, m, attrs, rng
func (_m *Lib) NewCredential(key *idemix.IssuerKey, m *idemix.CredRequest, attrs []*FP256BN.BIG, rng *amcl.RAND) (*idemix.Credential, error) {
	ret := _m.Called(key, m, attrs, rng)

	var r0 *idemix.Credential
	if rf, ok := ret.Get(0).(func(*idemix.IssuerKey, *idemix.CredRequest, []*FP256BN.BIG, *amcl.RAND) *idemix.Credential); ok {
		r0 = rf(key, m, attrs, rng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.Credential)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*idemix.IssuerKey, *idemix.CredRequest, []*FP256BN.BIG, *amcl.RAND) error); ok {
		r1 = rf(key, m, attrs, rng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIssuerKey provides a mock function with given fields: AttributeNames, rng
func (_m *Lib) NewIssuerKey(AttributeNames []string, rng *amcl.RAND) (*idemix.IssuerKey, error) {
	ret := _m.Called(AttributeNames, rng)

	var r0 *idemix.IssuerKey
	if rf, ok := ret.Get(0).(func([]string, *amcl.RAND) *idemix.IssuerKey); ok {
		r0 = rf(AttributeNames, rng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*idemix.IssuerKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, *amcl.RAND) error); ok {
		r1 = rf(AttributeNames, rng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandModOrder provides a mock function with given fields: rng
func (_m *Lib) RandModOrder(rng *amcl.RAND) (*FP256BN.BIG, error) {
	ret := _m.Called(rng)

	var r0 *FP256BN.BIG
	if rf, ok := ret.Get(0).(func(*amcl.RAND) *FP256BN.BIG); ok {
		r0 = rf(rng)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FP256BN.BIG)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*amcl.RAND) error); ok {
		r1 = rf(rng)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
