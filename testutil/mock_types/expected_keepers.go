// Code generated by MockGen. DO NOT EDIT.
// Source: x/checkers/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package mock_types

import (
	reflect "reflect"

	types "github.com/alice/checkers/x/leaderboard/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types0.Context, addr types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types0.Context, addr types0.AccAddress) types0.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types0.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockBankEscrowKeeper is a mock of BankEscrowKeeper interface.
type MockBankEscrowKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankEscrowKeeperMockRecorder
}

// MockBankEscrowKeeperMockRecorder is the mock recorder for MockBankEscrowKeeper.
type MockBankEscrowKeeperMockRecorder struct {
	mock *MockBankEscrowKeeper
}

// NewMockBankEscrowKeeper creates a new mock instance.
func NewMockBankEscrowKeeper(ctrl *gomock.Controller) *MockBankEscrowKeeper {
	mock := &MockBankEscrowKeeper{ctrl: ctrl}
	mock.recorder = &MockBankEscrowKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankEscrowKeeper) EXPECT() *MockBankEscrowKeeperMockRecorder {
	return m.recorder
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankEscrowKeeper) SendCoinsFromAccountToModule(ctx types0.Context, senderAddr types0.AccAddress, recipientModule string, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankEscrowKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankEscrowKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankEscrowKeeper) SendCoinsFromModuleToAccount(ctx types0.Context, senderModule string, recipientAddr types0.AccAddress, amt types0.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankEscrowKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankEscrowKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockCheckersLeaderboardKeeper is a mock of CheckersLeaderboardKeeper interface.
type MockCheckersLeaderboardKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckersLeaderboardKeeperMockRecorder
}

// MockCheckersLeaderboardKeeperMockRecorder is the mock recorder for MockCheckersLeaderboardKeeper.
type MockCheckersLeaderboardKeeperMockRecorder struct {
	mock *MockCheckersLeaderboardKeeper
}

// NewMockCheckersLeaderboardKeeper creates a new mock instance.
func NewMockCheckersLeaderboardKeeper(ctrl *gomock.Controller) *MockCheckersLeaderboardKeeper {
	mock := &MockCheckersLeaderboardKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckersLeaderboardKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckersLeaderboardKeeper) EXPECT() *MockCheckersLeaderboardKeeperMockRecorder {
	return m.recorder
}

// MustAddForfeitedGameResultToPlayer mocks base method.
func (m *MockCheckersLeaderboardKeeper) MustAddForfeitedGameResultToPlayer(ctx types0.Context, player types0.AccAddress) types.PlayerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustAddForfeitedGameResultToPlayer", ctx, player)
	ret0, _ := ret[0].(types.PlayerInfo)
	return ret0
}

// MustAddForfeitedGameResultToPlayer indicates an expected call of MustAddForfeitedGameResultToPlayer.
func (mr *MockCheckersLeaderboardKeeperMockRecorder) MustAddForfeitedGameResultToPlayer(ctx, player interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustAddForfeitedGameResultToPlayer", reflect.TypeOf((*MockCheckersLeaderboardKeeper)(nil).MustAddForfeitedGameResultToPlayer), ctx, player)
}

// MustAddLostGameResultToPlayer mocks base method.
func (m *MockCheckersLeaderboardKeeper) MustAddLostGameResultToPlayer(ctx types0.Context, player types0.AccAddress) types.PlayerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustAddLostGameResultToPlayer", ctx, player)
	ret0, _ := ret[0].(types.PlayerInfo)
	return ret0
}

// MustAddLostGameResultToPlayer indicates an expected call of MustAddLostGameResultToPlayer.
func (mr *MockCheckersLeaderboardKeeperMockRecorder) MustAddLostGameResultToPlayer(ctx, player interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustAddLostGameResultToPlayer", reflect.TypeOf((*MockCheckersLeaderboardKeeper)(nil).MustAddLostGameResultToPlayer), ctx, player)
}

// MustAddWonGameResultToPlayer mocks base method.
func (m *MockCheckersLeaderboardKeeper) MustAddWonGameResultToPlayer(ctx types0.Context, player types0.AccAddress) types.PlayerInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustAddWonGameResultToPlayer", ctx, player)
	ret0, _ := ret[0].(types.PlayerInfo)
	return ret0
}

// MustAddWonGameResultToPlayer indicates an expected call of MustAddWonGameResultToPlayer.
func (mr *MockCheckersLeaderboardKeeperMockRecorder) MustAddWonGameResultToPlayer(ctx, player interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustAddWonGameResultToPlayer", reflect.TypeOf((*MockCheckersLeaderboardKeeper)(nil).MustAddWonGameResultToPlayer), ctx, player)
}