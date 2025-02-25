// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractDelegationManager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IDelegationManagerTypesQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesQueuedWithdrawalParams struct {
	Strategies           []common.Address
	DepositShares        []*big.Int
	DeprecatedWithdrawer common.Address
}

// IDelegationManagerTypesWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerTypesWithdrawal struct {
	Staker       common.Address
	DelegatedTo  common.Address
	Withdrawer   common.Address
	Nonce        *big.Int
	StartBlock   uint32
	Strategies   []common.Address
	ScaledShares []*big.Int
}

// ISignatureUtilsMixinTypesSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsMixinTypesSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// ContractDelegationManagerMetaData contains all meta data concerning the ContractDelegationManager contract.
var ContractDelegationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_eigenPodManager\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"_MIN_WITHDRAWAL_DELAY\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DELEGATION_APPROVAL_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateDelegationApprovalDigestHash\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateWithdrawalRoot\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"convertToDepositShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeWithdrawalsQueued\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalQueued\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"curDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"beaconChainSlashingFactorDecrease\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegateTo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtilsMixinTypes.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegatedTo\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApprover\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationApproverSaltIsSpent\",\"inputs\":[{\"name\":\"delegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"spent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositScalingFactor\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorsShares\",\"inputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawal\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawalRoots\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQueuedWithdrawals\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"shares\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableSharesInQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawableShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"withdrawableShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseDelegatedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevDepositShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"addedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minWithdrawalDelayBlocks\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyOperatorDetails\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"pending\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManagerTypes.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"depositShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"__deprecated_withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawal\",\"type\":\"tuple\",\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redelegate\",\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newOperatorApproverSig\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtilsMixinTypes.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerAsOperator\",\"inputs\":[{\"name\":\"initDelegationApprover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperatorShares\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"prevMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"undelegate\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"withdrawalRoots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOperatorMetadataURI\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DelegationApproverUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newDelegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositScalingFactorUpdated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"newDepositScalingFactor\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorMetadataURIUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationApprover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesDecreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesIncreased\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"staker\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSharesSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"totalSlashedShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalCompleted\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashingWithdrawalQueued\",\"inputs\":[{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"withdrawal\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIDelegationManagerTypes.Withdrawal\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"scaledShares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"sharesToWithdraw\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerDelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerForceUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakerUndelegated\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FullySlashed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSnapshotOrdering\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActivelyDelegated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAllocationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManagerOrEigenPodManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsCannotUndelegate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"WithdrawalDelayNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotQueued\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawerNotCaller\",\"inputs\":[]}]",
	Bin: "0x61016060405234801561001157600080fd5b506040516164d23803806164d283398101604081905261003091610205565b808084898989878a6001600160a01b03811661005f576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0390811660805293841660a05291831660c052821660e05263ffffffff16610100521661012052610096816100b1565b61014052506100a590506100f8565b50505050505050610398565b600080829050601f815111156100e5578260405163305a27a960e01b81526004016100dc919061033e565b60405180910390fd5b80516100f082610371565b179392505050565b600054610100900460ff16156101605760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100dc565b60005460ff908116146101b1576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101c857600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b60005b838110156101fc5781810151838201526020016101e4565b50506000910152565b600080600080600080600060e0888a03121561022057600080fd5b875161022b816101b3565b602089015190975061023c816101b3565b604089015190965061024d816101b3565b606089015190955061025e816101b3565b608089015190945061026f816101b3565b60a089015190935063ffffffff8116811461028957600080fd5b60c08901519092506001600160401b038111156102a557600080fd5b88016000601f82018b136102b857600080fd5b81516001600160401b038111156102d1576102d16101cb565b604051601f8201601f19908116603f011681016001600160401b03811182821017156102ff576102ff6101cb565b6040528181528382016020018d101561031757600080fd5b6103288260208301602087016101e1565b8092508094505050505092959891949750929550565b602081526000825180602084015261035d8160408501602087016101e1565b601f01601f19169190910160400192915050565b80516020808301519190811015610392576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051610120516101405161603861049a6000396000818161113f01526141ef01526000818161046f015261348c01526000818161077e015281816136a7015281816137f00152613add0152600081816107ce01528181610e4a015281816110130152818161139e015281816115c201528181611a370152818161281c015261442701526000818161049601528181610f90015281816115200152818161179e015281816130830152818161326f015261393a0152600081816103cb01528181610f5e015281816116ed015261391401526000818161064901528181610bd20152818161117d015261260c01526160386000f3fe608060405234801561001057600080fd5b506004361061030c5760003560e01c8063715018a61161019d578063bfae3fd2116100e9578063e4cc3f90116100a2578063f2fde38b1161007c578063f2fde38b1461085c578063f698da251461086f578063fabc1cbc14610877578063fd8aa88d1461088a57600080fd5b8063e4cc3f9014610816578063eea9064b14610829578063f0e0e6761461083c57600080fd5b8063bfae3fd214610761578063c448feb814610774578063c978f7ac146107a8578063ca8aa7c7146107c9578063cd6dc687146107f0578063da8be8641461080357600080fd5b80639104c31911610156578063a178848411610130578063a1788484146106dd578063a33a3433146106fd578063b7f06ebe14610710578063bb45fef21461073357600080fd5b80639104c3191461068f5780639435bb43146106aa57806399f5371b146106bd57600080fd5b8063715018a6146105fe578063778e55f31461060657806378296ec514610631578063886f1195146106445780638da5cb5b1461066b578063900413471461067c57600080fd5b806354fd4d501161025c5780635dd685791161021557806365da1264116101ef57806365da12641461058e57806366d5ba93146105b75780636d70f7ae146105d85780636e174448146105eb57600080fd5b80635dd6857914610547578063601bb36f1461056857806360a0d1ce1461057b57600080fd5b806354fd4d50146104cb578063595c6a67146104e0578063597b36da146104e85780635ac86ab7146104fb5780635c975abb1461051e5780635d975e881461052657600080fd5b806339b70e38116102c95780633e28391d116102a35780633e28391d146104475780634657e26a1461046a5780634665bcda1461049157806354b7c96c146104b857600080fd5b806339b70e38146103c65780633c651cf2146104055780633cdeb5e01461041857600080fd5b806304a4f979146103115780630b9f487a1461034b5780630dd8dd021461035e578063136439dd1461037e57806325df922e146103935780632aa6d888146103b3575b600080fd5b6103387f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad81565b6040519081526020015b60405180910390f35b610338610359366004614e60565b61089d565b61037161036c366004614eff565b610926565b6040516103429190614f40565b61039161038c366004614f78565b610bbd565b005b6103a66103a1366004615118565b610c94565b60405161034291906151cd565b6103916103c1366004615235565b610dfa565b6103ed7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610342565b610391610413366004615297565b610f53565b6103ed6104263660046152dd565b6001600160a01b039081166000908152609960205260409020600101541690565b61045a6104553660046152dd565b6110aa565b6040519015158152602001610342565b6103ed7f000000000000000000000000000000000000000000000000000000000000000081565b6103ed7f000000000000000000000000000000000000000000000000000000000000000081565b6103916104c63660046152fa565b6110ca565b6104d3611138565b6040516103429190615383565b610391611168565b6103386104f6366004615456565b61121a565b61045a61050936600461548a565b606654600160ff9092169190911b9081161490565b606654610338565b610539610534366004614f78565b61124a565b604051610342929190615567565b61055a6105553660046152dd565b611267565b6040516103429291906155dc565b610391610576366004615660565b611393565b6103916105893660046156bc565b611515565b6103ed61059c3660046152dd565b609a602052600090815260409020546001600160a01b031681565b6105ca6105c53660046152dd565b6116c4565b6040516103429291906156fe565b61045a6105e63660046152dd565b6119d1565b6103386105f93660046152fa565b611a0b565b610391611ab9565b6103386106143660046152fa565b609860209081526000928352604080842090915290825290205481565b61039161063f366004615711565b611acb565b6103ed7f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b03166103ed565b6103a661068a366004615765565b611b61565b6103ed73beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6103916106b83660046157b4565b611c3b565b6106d06106cb366004614f78565b611d15565b6040516103429190615857565b6103386106eb3660046152dd565b609f6020526000908152604090205481565b61037161070b36600461586a565b611e36565b61045a61071e366004614f78565b609e6020526000908152604090205460ff1681565b61045a61074136600461595a565b609c60209081526000928352604080842090915290825290205460ff1681565b61033861076f3660046152fa565b611e4e565b60405163ffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610342565b6107bb6107b6366004615765565b611e8b565b604051610342929190615986565b6103ed7f000000000000000000000000000000000000000000000000000000000000000081565b6103916107fe36600461595a565b612123565b6103716108113660046152dd565b612244565b6103916108243660046159a7565b61236f565b61039161083736600461586a565b6123c5565b61084f61084a366004615a20565b612430565b6040516103429190615ad3565b61039161086a3660046152dd565b6124d7565b610338612550565b610391610885366004614f78565b61260a565b6103716108983660046152dd565b612723565b604080517f14bde674c9f64b2ad00eaaee4a8bed1fabef35c7507e3c5b9cfc9436909a2dad60208201526001600160a01b03808616928201929092528187166060820152908516608082015260a0810183905260c0810182905260009061091c9060e00160405160208183030381529060405280519060200120612747565b9695505050505050565b6066546060906001906002908116036109525760405163840a48d560e01b815260040160405180910390fd5b61095a612776565b6000836001600160401b0381111561097457610974614f91565b60405190808252806020026020018201604052801561099d578160200160208202803683370190505b50336000908152609a60205260408120549192506001600160a01b03909116905b85811015610bae578686828181106109d8576109d8615ae6565b90506020028101906109ea9190615afc565b6109f8906020810190615b1c565b9050878783818110610a0c57610a0c615ae6565b9050602002810190610a1e9190615afc565b610a289080615b1c565b905014610a48576040516343714afd60e01b815260040160405180910390fd5b6000610ab433848a8a86818110610a6157610a61615ae6565b9050602002810190610a739190615afc565b610a7d9080615b1c565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506127cf92505050565b9050610b8833848a8a86818110610acd57610acd615ae6565b9050602002810190610adf9190615afc565b610ae99080615b1c565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508e92508d9150889050818110610b2f57610b2f615ae6565b9050602002810190610b419190615afc565b610b4f906020810190615b1c565b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250889250612928915050565b848381518110610b9a57610b9a615ae6565b6020908102919091010152506001016109be565b5050600160c955949350505050565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610c21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c459190615b65565b610c6257604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610c875760405163c61dca5d60e01b815260040160405180910390fd5b610c9082612eb0565b5050565b6001600160a01b038084166000908152609a60205260408120546060921690610cbe8683876127cf565b9050600085516001600160401b03811115610cdb57610cdb614f91565b604051908082528060200260200182016040528015610d04578160200160208202803683370190505b50905060005b8651811015610ded576001600160a01b038816600090815260a260205260408120885182908a9085908110610d4157610d41615ae6565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206040518060200160405290816000820154815250509050610dc7878381518110610d9557610d95615ae6565b6020026020010151858481518110610daf57610daf615ae6565b602002602001015183612eed9092919063ffffffff16565b838381518110610dd957610dd9615ae6565b602090810291909101015250600101610d0a565b50925050505b9392505050565b610e02612776565b610e0b336110aa565b15610e2957604051633bf2b50360e11b815260040160405180910390fd5b604051632b6241f360e11b815233600482015263ffffffff841660248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906356c483e690604401600060405180830381600087803b158015610e9657600080fd5b505af1158015610eaa573d6000803e3d6000fd5b50505050610eb83385612f0c565b610ec23333612f6f565b6040516001600160a01b038516815233907fa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c19060200160405180910390a2336001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908383604051610f3b929190615b82565b60405180910390a2610f4d600160c955565b50505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610fb25750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b610fcf5760405163045206a560e21b815260040160405180910390fd5b610fd7612776565b6001600160a01b038481166000908152609a602052604080822054905163152667d960e31b8152908316600482018190528684166024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa15801561105a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107e9190615bb1565b9050600061108d878784613226565b905061109d83888888888661330c565b505050610f4d600160c955565b6001600160a01b039081166000908152609a602052604090205416151590565b816110d48161344d565b6110f15760405163932d94f760e01b815260040160405180910390fd5b6110f9612776565b611102836119d1565b61111f576040516325ec6c1f60e01b815260040160405180910390fd5b6111298383612f0c565b611133600160c955565b505050565b60606111637f00000000000000000000000000000000000000000000000000000000000000006134fb565b905090565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa1580156111cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f09190615b65565b61120d57604051631d77d47760e21b815260040160405180910390fd5b611218600019612eb0565b565b60008160405160200161122d9190615857565b604051602081830303815290604052805190602001209050919050565b611252614d13565b606061125d8361353a565b9094909350915050565b606080600061127584612723565b8051909150806001600160401b0381111561129257611292614f91565b6040519080825280602002602001820160405280156112cb57816020015b6112b8614d13565b8152602001906001900390816112b05790505b509350806001600160401b038111156112e6576112e6614f91565b60405190808252806020026020018201604052801561131957816020015b60608152602001906001900390816113045790505b50925060005b8181101561138b5761134983828151811061133c5761133c615ae6565b602002602001015161353a565b86838151811061135b5761135b615ae6565b6020026020010186848151811061137457611374615ae6565b60209081029190910101919091525260010161131f565b505050915091565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146113dc576040516323d871a560e01b815260040160405180910390fd5b6113e4612776565b6001600160a01b038085166000908152609860209081526040808320938716835292905290812054611423906001600160401b03808616908516613797565b90506000611433868686866137b0565b905060006114418284615be4565b90506114508760008886613871565b604080516001600160a01b038881168252602082018490528916917fdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30910160405180910390a260006114a1876138ec565b60405163debe1eab60e01b81526001600160a01b038981166004830152602482018590529192509082169063debe1eab90604401600060405180830381600087803b1580156114ef57600080fd5b505af1158015611503573d6000803e3d6000fd5b5050505050505050610f4d600160c955565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461155e57604051633213a66160e21b815260040160405180910390fd5b611566612776565b61156f836110aa565b15611129576001600160a01b038381166000908152609a602052604080822054905163152667d960e31b81529083166004820181905273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06024830152927f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015611609573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162d9190615bb1565b6001600160a01b038616600090815260a26020908152604080832073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac08452825280832081519283019091525481529192506116948661168c6001600160401b0380871690891661395f565b849190613974565b90506116b6848873beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac084613871565b50505050611133600160c955565b6040516394f649dd60e01b81526001600160a01b038281166004830152606091829160009182917f000000000000000000000000000000000000000000000000000000000000000016906394f649dd90602401600060405180830381865afa158015611734573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261175c9190810190615c55565b60405163fe243a1760e01b81526001600160a01b03888116600483015273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac060248301529294509092506000917f0000000000000000000000000000000000000000000000000000000000000000169063fe243a1790604401602060405180830381865afa1580156117e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118099190615d16565b90508060000361181e57509094909350915050565b60008351600161182e9190615be4565b6001600160401b0381111561184557611845614f91565b60405190808252806020026020018201604052801561186e578160200160208202803683370190505b5090506000845160016118819190615be4565b6001600160401b0381111561189857611898614f91565b6040519080825280602002602001820160405280156118c1578160200160208202803683370190505b50905073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0828651815181106118ec576118ec615ae6565b60200260200101906001600160a01b031690816001600160a01b031681525050828186518151811061192057611920615ae6565b60200260200101818152505060005b85518110156119c35785818151811061194a5761194a615ae6565b602002602001015183828151811061196457611964615ae6565b60200260200101906001600160a01b031690816001600160a01b03168152505084818151811061199657611996615ae6565b60200260200101518282815181106119b0576119b0615ae6565b602090810291909101015260010161192f565b509097909650945050505050565b60006001600160a01b03821615801590611a0557506001600160a01b038083166000818152609a6020526040902054909116145b92915050565b60405163152667d960e31b81526001600160a01b038381166004830152828116602483015260009182917f0000000000000000000000000000000000000000000000000000000000000000169063a9333ec890604401602060405180830381865afa158015611a7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aa29190615bb1565b9050611ab184848360006137b0565b949350505050565b611ac1613993565b61121860006139ed565b82611ad58161344d565b611af25760405163932d94f760e01b815260040160405180910390fd5b611afb846119d1565b611b18576040516325ec6c1f60e01b815260040160405180910390fd5b836001600160a01b03167f02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b67080908484604051611b53929190615b82565b60405180910390a250505050565b6060600082516001600160401b03811115611b7e57611b7e614f91565b604051908082528060200260200182016040528015611ba7578160200160208202803683370190505b50905060005b8351811015611c33576001600160a01b03851660009081526098602052604081208551909190869084908110611be557611be5615ae6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002054828281518110611c2057611c20615ae6565b6020908102919091010152600101611bad565b509392505050565b606654600290600490811603611c645760405163840a48d560e01b815260040160405180910390fd5b611c6c612776565b8560005b81811015611d0057611cf8898983818110611c8d57611c8d615ae6565b9050602002810190611c9f9190615d2f565b611ca890615d45565b888884818110611cba57611cba615ae6565b9050602002810190611ccc9190615b1c565b888886818110611cde57611cde615ae6565b9050602002016020810190611cf39190615d51565b613a3f565b600101611c70565b5050611d0c600160c955565b50505050505050565b611d1d614d13565b600082815260a46020908152604091829020825160e08101845281546001600160a01b03908116825260018301548116828501526002830154168185015260038201546060820152600482015463ffffffff1660808201526005820180548551818602810186019096528086529194929360a08601939290830182828015611dce57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611db0575b5050505050815260200160068201805480602002602001604051908101604052809291908181526020018280548015611e2657602002820191906000526020600020905b815481526020019060010190808311611e12575b5050505050815250509050919050565b6060611e4133612244565b9050610df38484846123c5565b6001600160a01b03808316600090815260a260209081526040808320938516835292815282822083519182019093529154825290610df390613e9a565b60608082516001600160401b03811115611ea757611ea7614f91565b604051908082528060200260200182016040528015611ed0578160200160208202803683370190505b50915082516001600160401b03811115611eec57611eec614f91565b604051908082528060200260200182016040528015611f15578160200160208202803683370190505b506001600160a01b038086166000908152609a6020526040812054929350911690611f418683876127cf565b905060005b8551811015612118576000611f73878381518110611f6657611f66615ae6565b60200260200101516138ec565b9050806001600160a01b031663fe243a1789898581518110611f9757611f97615ae6565b60200260200101516040518363ffffffff1660e01b8152600401611fd19291906001600160a01b0392831681529116602082015260400190565b602060405180830381865afa158015611fee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120129190615d16565b85838151811061202457612024615ae6565b602002602001018181525050600060a260008a6001600160a01b03166001600160a01b03168152602001908152602001600020600089858151811061206b5761206b615ae6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060405180602001604052908160008201548152505090506120f18684815181106120bf576120bf615ae6565b60200260200101518585815181106120d9576120d9615ae6565b6020026020010151836139749092919063ffffffff16565b87848151811061210357612103615ae6565b60209081029190910101525050600101611f46565b5050505b9250929050565b600054610100900460ff16158080156121435750600054600160ff909116105b8061215d5750303b15801561215d575060005460ff166001145b6121c55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156121e8576000805461ff0019166101001790555b6121f182612eb0565b6121fa836139ed565b8015611133576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b606061224e612776565b612257826110aa565b6122745760405163a5c7c44560e01b815260040160405180910390fd5b61227d826119d1565b1561229b576040516311ca333560e31b815260040160405180910390fd5b336001600160a01b03831614612355576001600160a01b038083166000908152609a6020526040902054166122cf8161344d565b806122f657506001600160a01b038181166000908152609960205260409020600101541633145b61231357604051631e499a2360e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b03167ff0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a60405160405180910390a3505b61235e82613eba565b905061236a600160c955565b919050565b6066546002906004908116036123985760405163840a48d560e01b815260040160405180910390fd5b6123a0612776565b6123b46123ac86615d45565b858585613a3f565b6123be600160c955565b5050505050565b6123cd612776565b6123d6336110aa565b156123f457604051633bf2b50360e11b815260040160405180910390fd5b6123fd836119d1565b61241a576040516325ec6c1f60e01b815260040160405180910390fd5b61242633848484614124565b6111293384612f6f565b6060600083516001600160401b0381111561244d5761244d614f91565b60405190808252806020026020018201604052801561248057816020015b606081526020019060019003908161246b5790505b50905060005b8451811015611c33576124b28582815181106124a4576124a4615ae6565b602002602001015185611b61565b8282815181106124c4576124c4615ae6565b6020908102919091010152600101612486565b6124df613993565b6001600160a01b0381166125445760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016121bc565b61254d816139ed565b50565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b60209091015260007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea6125be6141e6565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612668573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061268c9190615d6e565b6001600160a01b0316336001600160a01b0316146126bd5760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146126e45760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020015b60405180910390a25050565b6001600160a01b038116600090815260a360205260409020606090611a0590614284565b6000612751612550565b60405161190160f01b602082015260228101919091526042810183905260620161122d565b600260c954036127c85760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016121bc565b600260c955565b6060600082516001600160401b038111156127ec576127ec614f91565b604051908082528060200260200182016040528015612815578160200160208202803683370190505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663547afb8786866040518363ffffffff1660e01b8152600401612868929190615d8b565b600060405180830381865afa158015612885573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526128ad9190810190615daf565b905060005b845181101561291d576128f8878683815181106128d1576128d1615ae6565b60200260200101518484815181106128eb576128eb615ae6565b6020026020010151613226565b83828151811061290a5761290a615ae6565b60209081029190910101526001016128b2565b509095945050505050565b60006001600160a01b038616612951576040516339b190bb60e11b815260040160405180910390fd5b83516000036129735760405163796cc52560e01b815260040160405180910390fd5b600084516001600160401b0381111561298e5761298e614f91565b6040519080825280602002602001820160405280156129b7578160200160208202803683370190505b509050600085516001600160401b038111156129d5576129d5614f91565b6040519080825280602002602001820160405280156129fe578160200160208202803683370190505b50905060005b8651811015612cde576000612a24888381518110611f6657611f66615ae6565b9050600060a260008c6001600160a01b03166001600160a01b0316815260200190815260200160002060008a8581518110612a6157612a61615ae6565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206040518060200160405290816000820154815250509050612acf888481518110612ab557612ab5615ae6565b60200260200101518885815181106120d9576120d9615ae6565b848481518110612ae157612ae1615ae6565b602002602001018181525050612b19888481518110612b0257612b02615ae6565b60200260200101518261429190919063ffffffff16565b858481518110612b2b57612b2b615ae6565b60209081029190910101526001600160a01b038a1615612bc057612b828a8a8581518110612b5b57612b5b615ae6565b6020026020010151878681518110612b7557612b75615ae6565b60200260200101516142a6565b612bc08a8c8b8681518110612b9957612b99615ae6565b6020026020010151878781518110612bb357612bb3615ae6565b6020026020010151613871565b6000826001600160a01b031663724af4238d8c8781518110612be457612be4615ae6565b60200260200101518c8881518110612bfe57612bfe615ae6565b60200260200101516040518463ffffffff1660e01b8152600401612c2493929190615e43565b6020604051808303816000875af1158015612c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c679190615d16565b905080600003612cd0576001600160a01b038c16600090815260a2602052604081208b51612cd092908d9088908110612ca257612ca2615ae6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060009055565b505050806001019050612a04565b506001600160a01b0388166000908152609f60205260408120805491829190612d0683615e67565b919050555060006040518060e001604052808b6001600160a01b031681526020018a6001600160a01b031681526020018b6001600160a01b031681526020018381526020014363ffffffff1681526020018981526020018581525090506000612d6e8261121a565b6000818152609e602090815260408083208054600160ff19909116811790915560a4835292819020865181546001600160a01b03199081166001600160a01b039283161783558885015195830180548216968316969096179095559187015160028201805490951692169190911790925560608501516003830155608085015160048301805463ffffffff191663ffffffff90921691909117905560a085015180519394508593612e259260058501920190614d71565b5060c08201518051612e41916006840191602090910190614dd6565b5050506001600160a01b038b16600090815260a360205260409020612e669082614312565b507f26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30818386604051612e9a93929190615e80565b60405180910390a19a9950505050505050505050565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b6000611ab182612f06612eff87613e9a565b869061431e565b9061431e565b6001600160a01b0382811660008181526099602090815260409182902060010180546001600160a01b0319169486169485179055905192835290917f773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c69101612717565b606654600090600190811603612f985760405163840a48d560e01b815260040160405180910390fd5b600080612fa4856116c4565b915091506000612fb6600086856127cf565b6001600160a01b038781166000818152609a602052604080822080546001600160a01b031916948b16948517905551939450919290917fc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d874330491a360005b8351811015611d0c5773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06001600160a01b031684828151811061304b5761304b615ae6565b60200260200101516001600160a01b0316036131c45760405163a3d75e0960e01b81526001600160a01b0388811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa1580156130cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130f09190615bb1565b9050600060a260008a6001600160a01b03166001600160a01b03168152602001908152602001600020600087858151811061312d5761312d615ae6565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060405180602001604052908160008201548152505090506131a385848151811061318157613181615ae6565b6020026020010151836001600160401b0316836139749092919063ffffffff16565b8584815181106131b5576131b5615ae6565b60200260200101818152505050505b61321e86888684815181106131db576131db615ae6565b602002602001015160008786815181106131f7576131f7615ae6565b602002602001015187878151811061321157613211615ae6565b602002602001015161330c565b600101613012565b600073beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeabf196001600160a01b038416016132fc5760405163a3d75e0960e01b81526001600160a01b0385811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a3d75e0990602401602060405180830381865afa1580156132b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132dc9190615bb1565b90506132f46001600160401b0384811690831661395f565b915050610df3565b506001600160401b031692915050565b8060000361332d57604051630a33bc6960e21b815260040160405180910390fd5b6001600160a01b03808616600090815260a26020908152604080832093881683529290522061335e81858585614333565b6040805160208101909152815481527f8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f908790879061339c90613e9a565b6040516133ab93929190615e43565b60405180910390a16133bc866110aa565b15611d0c576001600160a01b038088166000908152609860209081526040808320938916835292905290812080548592906133f8908490615be4565b92505081905550866001600160a01b03167f1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c87878660405161343c93929190615e43565b60405180910390a250505050505050565b604051631beb2b9760e31b81526001600160a01b038281166004830152336024830152306044830152600080356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303816000875af11580156134d7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a059190615b65565b60606000613508836143b2565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b613542614d13565b600082815260a46020908152604091829020825160e08101845281546001600160a01b0390811682526001830154811682850152600283015416818501526003820154606082810191909152600483015463ffffffff1660808301526005830180548651818702810187019097528087529195929460a086019392908301828280156135f757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116135d9575b505050505081526020016006820180548060200260200160405190810160405280929190818152602001828054801561364f57602002820191906000526020600020905b81548152602001906001019080831161363b575b50505050508152505091508160a00151516001600160401b0381111561367757613677614f91565b6040519080825280602002602001820160405280156136a0578160200160208202803683370190505b50905060007f000000000000000000000000000000000000000000000000000000000000000083608001516136d59190615eab565b905060004363ffffffff168263ffffffff161061370857613703846000015185602001518660a001516127cf565b613720565b613720846000015185602001518660a00151856143da565b905060005b8460a001515181101561138b576137728560c00151828151811061374b5761374b615ae6565b602002602001015183838151811061376557613765615ae6565b602002602001015161450f565b84828151811061378457613784615ae6565b6020908102919091010152600101613725565b60006137a6848385600161451b565b611ab19085615ec7565b6001600160a01b03808516600090815260a560209081526040808320938716835292905290812081906137e290614578565b9050600061384a60016138157f000000000000000000000000000000000000000000000000000000000000000043615eda565b61381f9190615eda565b6001600160a01b03808a16600090815260a560209081526040808320938c1683529290522090614593565b905060006138588284615ec7565b90506138658187876145b0565b98975050505050505050565b6001600160a01b038085166000908152609860209081526040808320938616835292905290812080548392906138a8908490615ec7565b92505081905550836001600160a01b03167f6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd848484604051611b5393929190615e43565b60006001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014613938577f0000000000000000000000000000000000000000000000000000000000000000611a05565b7f000000000000000000000000000000000000000000000000000000000000000092915050565b6000610df38383670de0b6b3a76400006145cf565b6000611ab18261398d61398687613e9a565b869061395f565b9061395f565b6033546001600160a01b031633146112185760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016121bc565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60a0840151518214613a64576040516343714afd60e01b815260040160405180910390fd5b83604001516001600160a01b0316336001600160a01b031614613a9a576040516316110d3560e21b815260040160405180910390fd5b6000613aa58561121a565b6000818152609e602052604090205490915060ff16613ad7576040516387c9d21960e01b815260040160405180910390fd5b606060007f00000000000000000000000000000000000000000000000000000000000000008760800151613b0b9190615eab565b90508063ffffffff164363ffffffff1611613b39576040516378f67ae160e11b815260040160405180910390fd5b613b51876000015188602001518960a00151846143da565b87516001600160a01b0316600090815260a360205260409020909250613b789150836146b9565b50600082815260a46020526040812080546001600160a01b031990811682556001820180548216905560028201805490911690556003810182905560048101805463ffffffff1916905590613bd06005830182614e11565b613bde600683016000614e11565b50506000828152609e602052604090819020805460ff19169055517f1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a0090613c289084815260200190565b60405180910390a185516001600160a01b039081166000908152609a6020526040812054885160a08a01519190931692613c639184906127cf565b905060005b8860a0015151811015613e8f576000613c908a60a001518381518110611f6657611f66615ae6565b90506000613cc78b60c001518481518110613cad57613cad615ae6565b602002602001015187858151811061376557613765615ae6565b905080600003613cd8575050613e87565b8715613dac57816001600160a01b0316632eae418c8c600001518d60a001518681518110613d0857613d08615ae6565b60200260200101518d8d88818110613d2257613d22615ae6565b9050602002016020810190613d3791906152dd565b60405160e085901b6001600160e01b03191681526001600160a01b0393841660048201529183166024830152909116604482015260648101849052608401600060405180830381600087803b158015613d8f57600080fd5b505af1158015613da3573d6000803e3d6000fd5b50505050613e84565b600080836001600160a01b03166350ff72258e600001518f60a001518881518110613dd957613dd9615ae6565b6020026020010151866040518463ffffffff1660e01b8152600401613e0093929190615e43565b60408051808303816000875af1158015613e1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e429190615ef6565b91509150613e81878e600001518f60a001518881518110613e6557613e65615ae6565b602002602001015185858b8b8151811061321157613211615ae6565b50505b50505b600101613c68565b505050505050505050565b805160009015613eab578151611a05565b670de0b6b3a764000092915050565b606654606090600190600290811603613ee65760405163840a48d560e01b815260040160405180910390fd5b6001600160a01b038084166000818152609a602052604080822080546001600160a01b0319811690915590519316928392917ffee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af4467691a3600080613f47866116c4565b915091508151600003613f5c5750505061411e565b81516001600160401b03811115613f7557613f75614f91565b604051908082528060200260200182016040528015613f9e578160200160208202803683370190505b5094506000613fae8785856127cf565b905060005b835181101561411857604080516001808252818301909252600091602080830190803683375050604080516001808252818301909252929350600092915060208083019080368337505060408051600180825281830190925292935060009291506020808301908036833701905050905086848151811061403657614036615ae6565b60200260200101518360008151811061405157614051615ae6565b60200260200101906001600160a01b031690816001600160a01b03168152505085848151811061408357614083615ae6565b60200260200101518260008151811061409e5761409e615ae6565b6020026020010181815250508484815181106140bc576140bc615ae6565b6020026020010151816000815181106140d7576140d7615ae6565b6020026020010181815250506140f08b89858585612928565b8a858151811061410257614102615ae6565b6020908102919091010152505050600101613fb3565b50505050505b50919050565b6001600160a01b03808416600090815260996020526040902060010154168061414d5750610f4d565b6001600160a01b0381166000908152609c6020908152604080832085845290915290205460ff161561419257604051630d4c4c9160e21b815260040160405180910390fd5b6001600160a01b0381166000908152609c602090815260408083208584528252909120805460ff191660011790558301516123be9082906141da90889088908490889061089d565b855160208701516146c5565b606060006142137f00000000000000000000000000000000000000000000000000000000000000006134fb565b90508060008151811061422857614228615ae6565b602001015160f81c60f81b8160018151811061424657614246615ae6565b016020908101516040516001600160f81b03199384169281019290925291909116602182015260220160405160208183030381529060405291505090565b60606000610df383614717565b6000610df361429f84613e9a565b839061395f565b6001600160a01b03808416600090815260a56020908152604080832093861683529290529081206142d690614578565b9050610f4d436142e68484615be4565b6001600160a01b03808816600090815260a560209081526040808320938a168352929052209190614773565b6000610df3838361477e565b6000610df383670de0b6b3a7640000846145cf565b8260000361436057604080516020810190915284548152614359908290612f0690613e9a565b8455610f4d565b60408051602081019091528454815260009061437d908584613974565b9050600061438b8483615be4565b905060006143a784612f066143a0888a615be4565b859061431e565b875550505050505050565b600060ff8216601f811115611a0557604051632cd44ac360e21b815260040160405180910390fd5b6060600083516001600160401b038111156143f7576143f7614f91565b604051908082528060200260200182016040528015614420578160200160208202803683370190505b50905060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166394d7d00c8787876040518463ffffffff1660e01b815260040161447593929190615f1a565b600060405180830381865afa158015614492573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526144ba9190810190615daf565b905060005b8551811015614503576144de888783815181106128d1576128d1615ae6565b8382815181106144f0576144f0615ae6565b60209081029190910101526001016144bf565b50909695505050505050565b6000610df3838361395f565b6000806145298686866145cf565b9050600183600281111561453f5761453f615f54565b14801561455c57506000848061455757614557615f6a565b868809115b1561456f5761456c600182615be4565b90505b95945050505050565b600061458482826147cd565b6001600160e01b031692915050565b60006145a0838383614814565b6001600160e01b03169392505050565b6000611ab16145bf8385615f80565b85906001600160401b031661395f565b6000808060001985870985870292508281108382030391505080600003614609578382816145ff576145ff615f6a565b0492505050610df3565b8084116146505760405162461bcd60e51b81526020600482015260156024820152744d6174683a206d756c446976206f766572666c6f7760581b60448201526064016121bc565b60008486880960026001871981018816978890046003810283188082028403028082028403028082028403028082028403028082028403029081029092039091026000889003889004909101858311909403939093029303949094049190911702949350505050565b6000610df3838361485e565b428110156146e657604051630819bdcd60e01b815260040160405180910390fd5b6146fa6001600160a01b0385168484614951565b610f4d57604051638baa579f60e01b815260040160405180910390fd5b60608160000180548060200260200160405190810160405280929190818152602001828054801561476757602002820191906000526020600020905b815481526020019060010190808311614753575b50505050509050919050565b6111338383836149a8565b60008181526001830160205260408120546147c557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611a05565b506000611a05565b8154600090801561480c576147f5846147e7600184615ec7565b600091825260209091200190565b5464010000000090046001600160e01b0316611ab1565b509092915050565b82546000908161482686868385614ab0565b905080156148545761483d866147e7600184615ec7565b5464010000000090046001600160e01b031661091c565b5091949350505050565b60008181526001830160205260408120548015614947576000614882600183615ec7565b855490915060009061489690600190615ec7565b90508181146148fb5760008660000182815481106148b6576148b6615ae6565b90600052602060002001549050808760000184815481106148d9576148d9615ae6565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061490c5761490c615f9f565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050611a05565b6000915050611a05565b60008060006149608585614b06565b9092509050600081600481111561497957614979615f54565b1480156149975750856001600160a01b0316826001600160a01b0316145b8061091c575061091c868686614b48565b82548015614a615760006149c1856147e7600185615ec7565b60408051808201909152905463ffffffff8082168084526401000000009092046001600160e01b031660208401529192509085161015614a145760405163151b8e3f60e11b815260040160405180910390fd5b805163ffffffff808616911603614a5f5782614a35866147e7600186615ec7565b80546001600160e01b03929092166401000000000263ffffffff9092169190911790555050505050565b505b506040805180820190915263ffffffff92831681526001600160e01b03918216602080830191825285546001810187556000968752952091519051909216640100000000029190921617910155565b60005b81831015611c33576000614ac78484614c34565b60008781526020902090915063ffffffff86169082015463ffffffff161115614af257809250614b00565b614afd816001615be4565b93505b50614ab3565b6000808251604103614b3c5760208301516040840151606085015160001a614b3087828585614c4f565b9450945050505061211c565b5060009050600261211c565b6000806000856001600160a01b0316631626ba7e60e01b8686604051602401614b72929190615fb5565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051614bb09190615fce565b600060405180830381855afa9150503d8060008114614beb576040519150601f19603f3d011682016040523d82523d6000602084013e614bf0565b606091505b5091509150818015614c0457506020815110155b801561091c57508051630b135d3f60e11b90614c299083016020908101908401615d16565b149695505050505050565b6000614c436002848418615fe0565b610df390848416615be4565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115614c865750600090506003614d0a565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015614cda573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614d0357600060019250925050614d0a565b9150600090505b94509492505050565b6040518060e0016040528060006001600160a01b0316815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160008152602001600063ffffffff16815260200160608152602001606081525090565b828054828255906000526020600020908101928215614dc6579160200282015b82811115614dc657825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614d91565b50614dd2929150614e2b565b5090565b828054828255906000526020600020908101928215614dc6579160200282015b82811115614dc6578251825591602001919060010190614df6565b508054600082559060005260206000209081019061254d91905b5b80821115614dd25760008155600101614e2c565b6001600160a01b038116811461254d57600080fd5b803561236a81614e40565b600080600080600060a08688031215614e7857600080fd5b8535614e8381614e40565b94506020860135614e9381614e40565b93506040860135614ea381614e40565b94979396509394606081013594506080013592915050565b60008083601f840112614ecd57600080fd5b5081356001600160401b03811115614ee457600080fd5b6020830191508360208260051b850101111561211c57600080fd5b60008060208385031215614f1257600080fd5b82356001600160401b03811115614f2857600080fd5b614f3485828601614ebb565b90969095509350505050565b602080825282518282018190526000918401906040840190835b8181101561291d578351835260209384019390920191600101614f5a565b600060208284031215614f8a57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b0381118282101715614fc957614fc9614f91565b60405290565b604080519081016001600160401b0381118282101715614fc957614fc9614f91565b604051601f8201601f191681016001600160401b038111828210171561501957615019614f91565b604052919050565b60006001600160401b0382111561503a5761503a614f91565b5060051b60200190565b600082601f83011261505557600080fd5b813561506861506382615021565b614ff1565b8082825260208201915060208360051b86010192508583111561508a57600080fd5b602085015b838110156150b05780356150a281614e40565b83526020928301920161508f565b5095945050505050565b600082601f8301126150cb57600080fd5b81356150d961506382615021565b8082825260208201915060208360051b8601019250858311156150fb57600080fd5b602085015b838110156150b0578035835260209283019201615100565b60008060006060848603121561512d57600080fd5b833561513881614e40565b925060208401356001600160401b0381111561515357600080fd5b61515f86828701615044565b92505060408401356001600160401b0381111561517b57600080fd5b615187868287016150ba565b9150509250925092565b600081518084526020840193506020830160005b828110156151c35781518652602095860195909101906001016151a5565b5093949350505050565b602081526000610df36020830184615191565b803563ffffffff8116811461236a57600080fd5b60008083601f84011261520657600080fd5b5081356001600160401b0381111561521d57600080fd5b60208301915083602082850101111561211c57600080fd5b6000806000806060858703121561524b57600080fd5b843561525681614e40565b9350615264602086016151e0565b925060408501356001600160401b0381111561527f57600080fd5b61528b878288016151f4565b95989497509550505050565b600080600080608085870312156152ad57600080fd5b84356152b881614e40565b935060208501356152c881614e40565b93969395505050506040820135916060013590565b6000602082840312156152ef57600080fd5b8135610df381614e40565b6000806040838503121561530d57600080fd5b823561531881614e40565b9150602083013561532881614e40565b809150509250929050565b60005b8381101561534e578181015183820152602001615336565b50506000910152565b6000815180845261536f816020860160208601615333565b601f01601f19169290920160200192915050565b602081526000610df36020830184615357565b600060e082840312156153a857600080fd5b6153b0614fa7565b90506153bb82614e55565b81526153c960208301614e55565b60208201526153da60408301614e55565b6040820152606082810135908201526153f5608083016151e0565b608082015260a08201356001600160401b0381111561541357600080fd5b61541f84828501615044565b60a08301525060c08201356001600160401b0381111561543e57600080fd5b61544a848285016150ba565b60c08301525092915050565b60006020828403121561546857600080fd5b81356001600160401b0381111561547e57600080fd5b611ab184828501615396565b60006020828403121561549c57600080fd5b813560ff81168114610df357600080fd5b600081518084526020840193506020830160005b828110156151c35781516001600160a01b03168652602095860195909101906001016154c1565b80516001600160a01b03908116835260208083015182169084015260408083015190911690830152606080820151908301526080808201516000916155349085018263ffffffff169052565b5060a082015160e060a085015261554e60e08501826154ad565b905060c083015184820360c086015261456f8282615191565b60408152600061557a60408301856154e8565b828103602084015261456f8185615191565b600082825180855260208501945060208160051b8301016020850160005b8381101561450357601f198584030188526155c6838351615191565b60209889019890935091909101906001016155aa565b6000604082016040835280855180835260608501915060608160051b86010192506020870160005b8281101561563557605f198786030184526156208583516154e8565b94506020938401939190910190600101615604565b50505050828103602084015261456f818561558c565b6001600160401b038116811461254d57600080fd5b6000806000806080858703121561567657600080fd5b843561568181614e40565b9350602085013561569181614e40565b925060408501356156a18161564b565b915060608501356156b18161564b565b939692955090935050565b6000806000606084860312156156d157600080fd5b83356156dc81614e40565b92506020840135915060408401356156f38161564b565b809150509250925092565b60408152600061557a60408301856154ad565b60008060006040848603121561572657600080fd5b833561573181614e40565b925060208401356001600160401b0381111561574c57600080fd5b615758868287016151f4565b9497909650939450505050565b6000806040838503121561577857600080fd5b823561578381614e40565b915060208301356001600160401b0381111561579e57600080fd5b6157aa85828601615044565b9150509250929050565b600080600080600080606087890312156157cd57600080fd5b86356001600160401b038111156157e357600080fd5b6157ef89828a01614ebb565b90975095505060208701356001600160401b0381111561580e57600080fd5b61581a89828a01614ebb565b90955093505060408701356001600160401b0381111561583957600080fd5b61584589828a01614ebb565b979a9699509497509295939492505050565b602081526000610df360208301846154e8565b60008060006060848603121561587f57600080fd5b833561588a81614e40565b925060208401356001600160401b038111156158a557600080fd5b8401604081870312156158b757600080fd5b6158bf614fcf565b81356001600160401b038111156158d557600080fd5b8201601f810188136158e657600080fd5b80356001600160401b038111156158ff576158ff614f91565b615912601f8201601f1916602001614ff1565b81815289602083850101111561592757600080fd5b81602084016020830137600060209282018301528352928301359282019290925293969395505050506040919091013590565b6000806040838503121561596d57600080fd5b823561597881614e40565b946020939093013593505050565b60408152600061557a6040830185615191565b801515811461254d57600080fd5b600080600080606085870312156159bd57600080fd5b84356001600160401b038111156159d357600080fd5b850160e081880312156159e557600080fd5b935060208501356001600160401b03811115615a0057600080fd5b615a0c87828801614ebb565b90945092505060408501356156b181615999565b60008060408385031215615a3357600080fd5b82356001600160401b03811115615a4957600080fd5b8301601f81018513615a5a57600080fd5b8035615a6861506382615021565b8082825260208201915060208360051b850101925087831115615a8a57600080fd5b6020840193505b82841015615ab5578335615aa481614e40565b825260209384019390910190615a91565b945050505060208301356001600160401b0381111561579e57600080fd5b602081526000610df3602083018461558c565b634e487b7160e01b600052603260045260246000fd5b60008235605e19833603018112615b1257600080fd5b9190910192915050565b6000808335601e19843603018112615b3357600080fd5b8301803591506001600160401b03821115615b4d57600080fd5b6020019150600581901b360382131561211c57600080fd5b600060208284031215615b7757600080fd5b8151610df381615999565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b600060208284031215615bc357600080fd5b8151610df38161564b565b634e487b7160e01b600052601160045260246000fd5b80820180821115611a0557611a05615bce565b600082601f830112615c0857600080fd5b8151615c1661506382615021565b8082825260208201915060208360051b860101925085831115615c3857600080fd5b602085015b838110156150b0578051835260209283019201615c3d565b60008060408385031215615c6857600080fd5b82516001600160401b03811115615c7e57600080fd5b8301601f81018513615c8f57600080fd5b8051615c9d61506382615021565b8082825260208201915060208360051b850101925087831115615cbf57600080fd5b6020840193505b82841015615cea578351615cd981614e40565b825260209384019390910190615cc6565b8095505050505060208301516001600160401b03811115615d0a57600080fd5b6157aa85828601615bf7565b600060208284031215615d2857600080fd5b5051919050565b6000823560de19833603018112615b1257600080fd5b6000611a053683615396565b600060208284031215615d6357600080fd5b8135610df381615999565b600060208284031215615d8057600080fd5b8151610df381614e40565b6001600160a01b0383168152604060208201819052600090611ab1908301846154ad565b600060208284031215615dc157600080fd5b81516001600160401b03811115615dd757600080fd5b8201601f81018413615de857600080fd5b8051615df661506382615021565b8082825260208201915060208360051b850101925086831115615e1857600080fd5b6020840193505b8284101561091c578351615e328161564b565b825260209384019390910190615e1f565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060018201615e7957615e79615bce565b5060010190565b838152606060208201526000615e9960608301856154e8565b828103604084015261091c8185615191565b63ffffffff8181168382160190811115611a0557611a05615bce565b81810381811115611a0557611a05615bce565b63ffffffff8281168282160390811115611a0557611a05615bce565b60008060408385031215615f0957600080fd5b505080516020909101519092909150565b6001600160a01b0384168152606060208201819052600090615f3e908301856154ad565b905063ffffffff83166040830152949350505050565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b6001600160401b038281168282160390811115611a0557611a05615bce565b634e487b7160e01b600052603160045260246000fd5b828152604060208201526000611ab16040830184615357565b60008251615b12818460208701615333565b600082615ffd57634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220a3e67ee83e71af501e07aba44e389da37f39cea20af29c7831409a910d87e85364736f6c634300081b0033",
}

// ContractDelegationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractDelegationManagerMetaData.ABI instead.
var ContractDelegationManagerABI = ContractDelegationManagerMetaData.ABI

// ContractDelegationManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractDelegationManagerMetaData.Bin instead.
var ContractDelegationManagerBin = ContractDelegationManagerMetaData.Bin

// DeployContractDelegationManager deploys a new Ethereum contract, binding an instance of ContractDelegationManager to it.
func DeployContractDelegationManager(auth *bind.TransactOpts, backend bind.ContractBackend, _strategyManager common.Address, _eigenPodManager common.Address, _allocationManager common.Address, _pauserRegistry common.Address, _permissionController common.Address, _MIN_WITHDRAWAL_DELAY uint32, _version string) (common.Address, *types.Transaction, *ContractDelegationManager, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractDelegationManagerBin), backend, _strategyManager, _eigenPodManager, _allocationManager, _pauserRegistry, _permissionController, _MIN_WITHDRAWAL_DELAY, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// ContractDelegationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractDelegationManagerMethods interface {
	ContractDelegationManagerCalls
	ContractDelegationManagerTransacts
	ContractDelegationManagerFilters
}

// ContractDelegationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractDelegationManagerCalls interface {
	DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error)

	CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error)

	ConvertToDepositShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error)

	CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error)

	DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error)

	DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error)

	DelegationApproverSaltIsSpent(opts *bind.CallOpts, delegationApprover common.Address, salt [32]byte) (bool, error)

	DepositScalingFactor(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetDepositedShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error)

	GetOperatorShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error)

	GetOperatorsShares(opts *bind.CallOpts, operators []common.Address, strategies []common.Address) ([][]*big.Int, error)

	GetQueuedWithdrawal(opts *bind.CallOpts, withdrawalRoot [32]byte) (struct {
		Withdrawal IDelegationManagerTypesWithdrawal
		Shares     []*big.Int
	}, error)

	GetQueuedWithdrawalRoots(opts *bind.CallOpts, staker common.Address) ([][32]byte, error)

	GetQueuedWithdrawals(opts *bind.CallOpts, staker common.Address) (struct {
		Withdrawals []IDelegationManagerTypesWithdrawal
		Shares      [][]*big.Int
	}, error)

	GetSlashableSharesInQueue(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error)

	GetWithdrawableShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) (struct {
		WithdrawableShares []*big.Int
		DepositShares      []*big.Int
	}, error)

	IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error)

	IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error)

	MinWithdrawalDelayBlocks(opts *bind.CallOpts) (uint32, error)

	OperatorShares(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PendingWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error)

	PermissionController(opts *bind.CallOpts) (common.Address, error)

	QueuedWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)

	Version(opts *bind.CallOpts) (string, error)
}

// ContractDelegationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractDelegationManagerTransacts interface {
	CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error)

	CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error)

	DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error)

	DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyOperatorDetails(opts *bind.TransactOpts, operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error)

	Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error)

	RegisterAsOperator(opts *bind.TransactOpts, initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SlashOperatorShares(opts *bind.TransactOpts, operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateOperatorMetadataURI(opts *bind.TransactOpts, operator common.Address, metadataURI string) (*types.Transaction, error)
}

// ContractDelegationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractDelegationManagerFilters interface {
	FilterDelegationApproverUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerDelegationApproverUpdatedIterator, error)
	WatchDelegationApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDelegationApproverUpdated, operator []common.Address) (event.Subscription, error)
	ParseDelegationApproverUpdated(log types.Log) (*ContractDelegationManagerDelegationApproverUpdated, error)

	FilterDepositScalingFactorUpdated(opts *bind.FilterOpts) (*ContractDelegationManagerDepositScalingFactorUpdatedIterator, error)
	WatchDepositScalingFactorUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDepositScalingFactorUpdated) (event.Subscription, error)
	ParseDepositScalingFactorUpdated(log types.Log) (*ContractDelegationManagerDepositScalingFactorUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error)

	FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error)
	WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error)
	ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error)

	FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error)
	WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error)
	ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error)

	FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error)
	WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error)

	FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error)
	WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error)

	FilterOperatorSharesSlashed(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesSlashedIterator, error)
	WatchOperatorSharesSlashed(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesSlashed, operator []common.Address) (event.Subscription, error)
	ParseOperatorSharesSlashed(log types.Log) (*ContractDelegationManagerOperatorSharesSlashed, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error)

	FilterSlashingWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalCompletedIterator, error)
	WatchSlashingWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalCompleted) (event.Subscription, error)
	ParseSlashingWithdrawalCompleted(log types.Log) (*ContractDelegationManagerSlashingWithdrawalCompleted, error)

	FilterSlashingWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalQueuedIterator, error)
	WatchSlashingWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalQueued) (event.Subscription, error)
	ParseSlashingWithdrawalQueued(log types.Log) (*ContractDelegationManagerSlashingWithdrawalQueued, error)

	FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error)
	WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error)

	FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error)
	WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error)

	FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error)
	WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error)
	ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error)
}

// ContractDelegationManager is an auto generated Go binding around an Ethereum contract.
type ContractDelegationManager struct {
	ContractDelegationManagerCaller     // Read-only binding to the contract
	ContractDelegationManagerTransactor // Write-only binding to the contract
	ContractDelegationManagerFilterer   // Log filterer for contract events
}

// ContractDelegationManager implements the ContractDelegationManagerMethods interface.
var _ ContractDelegationManagerMethods = (*ContractDelegationManager)(nil)

// ContractDelegationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerCaller implements the ContractDelegationManagerCalls interface.
var _ ContractDelegationManagerCalls = (*ContractDelegationManagerCaller)(nil)

// ContractDelegationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerTransactor implements the ContractDelegationManagerTransacts interface.
var _ ContractDelegationManagerTransacts = (*ContractDelegationManagerTransactor)(nil)

// ContractDelegationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractDelegationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDelegationManagerFilterer implements the ContractDelegationManagerFilters interface.
var _ ContractDelegationManagerFilters = (*ContractDelegationManagerFilterer)(nil)

// ContractDelegationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDelegationManagerSession struct {
	Contract     *ContractDelegationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractDelegationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDelegationManagerCallerSession struct {
	Contract *ContractDelegationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractDelegationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDelegationManagerTransactorSession struct {
	Contract     *ContractDelegationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractDelegationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDelegationManagerRaw struct {
	Contract *ContractDelegationManager // Generic contract binding to access the raw methods on
}

// ContractDelegationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDelegationManagerCallerRaw struct {
	Contract *ContractDelegationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDelegationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDelegationManagerTransactorRaw struct {
	Contract *ContractDelegationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDelegationManager creates a new instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManager(address common.Address, backend bind.ContractBackend) (*ContractDelegationManager, error) {
	contract, err := bindContractDelegationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManager{ContractDelegationManagerCaller: ContractDelegationManagerCaller{contract: contract}, ContractDelegationManagerTransactor: ContractDelegationManagerTransactor{contract: contract}, ContractDelegationManagerFilterer: ContractDelegationManagerFilterer{contract: contract}}, nil
}

// NewContractDelegationManagerCaller creates a new read-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractDelegationManagerCaller, error) {
	contract, err := bindContractDelegationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerCaller{contract: contract}, nil
}

// NewContractDelegationManagerTransactor creates a new write-only instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDelegationManagerTransactor, error) {
	contract, err := bindContractDelegationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerTransactor{contract: contract}, nil
}

// NewContractDelegationManagerFilterer creates a new log filterer instance of ContractDelegationManager, bound to a specific deployed contract.
func NewContractDelegationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractDelegationManagerFilterer, error) {
	contract, err := bindContractDelegationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerFilterer{contract: contract}, nil
}

// bindContractDelegationManager binds a generic wrapper to an already deployed contract.
func bindContractDelegationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractDelegationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.ContractDelegationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ContractDelegationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDelegationManager *ContractDelegationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDelegationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDelegationManager *ContractDelegationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DELEGATIONAPPROVALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "DELEGATION_APPROVAL_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// DELEGATIONAPPROVALTYPEHASH is a free data retrieval call binding the contract method 0x04a4f979.
//
// Solidity: function DELEGATION_APPROVAL_TYPEHASH() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DELEGATIONAPPROVALTYPEHASH() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DELEGATIONAPPROVALTYPEHASH(&_ContractDelegationManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) AllocationManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.AllocationManager(&_ContractDelegationManager.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) AllocationManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.AllocationManager(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractDelegationManager.Contract.BeaconChainETHStrategy(&_ContractDelegationManager.CallOpts)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateDelegationApprovalDigestHash(opts *bind.CallOpts, staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateDelegationApprovalDigestHash", staker, operator, approver, approverSalt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateDelegationApprovalDigestHash is a free data retrieval call binding the contract method 0x0b9f487a.
//
// Solidity: function calculateDelegationApprovalDigestHash(address staker, address operator, address approver, bytes32 approverSalt, uint256 expiry) view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateDelegationApprovalDigestHash(staker common.Address, operator common.Address, approver common.Address, approverSalt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateDelegationApprovalDigestHash(&_ContractDelegationManager.CallOpts, staker, operator, approver, approverSalt, expiry)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CalculateWithdrawalRoot(opts *bind.CallOpts, withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "calculateWithdrawalRoot", withdrawal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// CalculateWithdrawalRoot is a free data retrieval call binding the contract method 0x597b36da.
//
// Solidity: function calculateWithdrawalRoot((address,address,address,uint256,uint32,address[],uint256[]) withdrawal) pure returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CalculateWithdrawalRoot(withdrawal IDelegationManagerTypesWithdrawal) ([32]byte, error) {
	return _ContractDelegationManager.Contract.CalculateWithdrawalRoot(&_ContractDelegationManager.CallOpts, withdrawal)
}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) ConvertToDepositShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "convertToDepositShares", staker, strategies, withdrawableShares)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) ConvertToDepositShares(staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.ConvertToDepositShares(&_ContractDelegationManager.CallOpts, staker, strategies, withdrawableShares)
}

// ConvertToDepositShares is a free data retrieval call binding the contract method 0x25df922e.
//
// Solidity: function convertToDepositShares(address staker, address[] strategies, uint256[] withdrawableShares) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) ConvertToDepositShares(staker common.Address, strategies []common.Address, withdrawableShares []*big.Int) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.ConvertToDepositShares(&_ContractDelegationManager.CallOpts, staker, strategies, withdrawableShares)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerCaller) CumulativeWithdrawalsQueued(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "cumulativeWithdrawalsQueued", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, staker)
}

// CumulativeWithdrawalsQueued is a free data retrieval call binding the contract method 0xa1788484.
//
// Solidity: function cumulativeWithdrawalsQueued(address staker) view returns(uint256 totalQueued)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) CumulativeWithdrawalsQueued(staker common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.CumulativeWithdrawalsQueued(&_ContractDelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegatedTo(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegatedTo", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, staker)
}

// DelegatedTo is a free data retrieval call binding the contract method 0x65da1264.
//
// Solidity: function delegatedTo(address staker) view returns(address operator)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegatedTo(staker common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegatedTo(&_ContractDelegationManager.CallOpts, staker)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApprover(opts *bind.CallOpts, operator common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApprover", operator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApprover is a free data retrieval call binding the contract method 0x3cdeb5e0.
//
// Solidity: function delegationApprover(address operator) view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApprover(operator common.Address) (common.Address, error) {
	return _ContractDelegationManager.Contract.DelegationApprover(&_ContractDelegationManager.CallOpts, operator)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DelegationApproverSaltIsSpent(opts *bind.CallOpts, delegationApprover common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "delegationApproverSaltIsSpent", delegationApprover, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, delegationApprover, salt)
}

// DelegationApproverSaltIsSpent is a free data retrieval call binding the contract method 0xbb45fef2.
//
// Solidity: function delegationApproverSaltIsSpent(address delegationApprover, bytes32 salt) view returns(bool spent)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DelegationApproverSaltIsSpent(delegationApprover common.Address, salt [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.DelegationApproverSaltIsSpent(&_ContractDelegationManager.CallOpts, delegationApprover, salt)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DepositScalingFactor(opts *bind.CallOpts, staker common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "depositScalingFactor", staker, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.DepositScalingFactor(&_ContractDelegationManager.CallOpts, staker, strategy)
}

// DepositScalingFactor is a free data retrieval call binding the contract method 0xbfae3fd2.
//
// Solidity: function depositScalingFactor(address staker, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DepositScalingFactor(staker common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.DepositScalingFactor(&_ContractDelegationManager.CallOpts, staker, strategy)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractDelegationManager.Contract.DomainSeparator(&_ContractDelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// EigenPodManager is a free data retrieval call binding the contract method 0x4665bcda.
//
// Solidity: function eigenPodManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) EigenPodManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.EigenPodManager(&_ContractDelegationManager.CallOpts)
}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetDepositedShares(opts *bind.CallOpts, staker common.Address) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getDepositedShares", staker)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDepositedShares(&_ContractDelegationManager.CallOpts, staker)
}

// GetDepositedShares is a free data retrieval call binding the contract method 0x66d5ba93.
//
// Solidity: function getDepositedShares(address staker) view returns(address[], uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetDepositedShares(staker common.Address) ([]common.Address, []*big.Int, error) {
	return _ContractDelegationManager.Contract.GetDepositedShares(&_ContractDelegationManager.CallOpts, staker)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetOperatorShares(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getOperatorShares", operator, strategies)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorShares(&_ContractDelegationManager.CallOpts, operator, strategies)
}

// GetOperatorShares is a free data retrieval call binding the contract method 0x90041347.
//
// Solidity: function getOperatorShares(address operator, address[] strategies) view returns(uint256[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetOperatorShares(operator common.Address, strategies []common.Address) ([]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorShares(&_ContractDelegationManager.CallOpts, operator, strategies)
}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetOperatorsShares(opts *bind.CallOpts, operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getOperatorsShares", operators, strategies)

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorsShares(&_ContractDelegationManager.CallOpts, operators, strategies)
}

// GetOperatorsShares is a free data retrieval call binding the contract method 0xf0e0e676.
//
// Solidity: function getOperatorsShares(address[] operators, address[] strategies) view returns(uint256[][])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetOperatorsShares(operators []common.Address, strategies []common.Address) ([][]*big.Int, error) {
	return _ContractDelegationManager.Contract.GetOperatorsShares(&_ContractDelegationManager.CallOpts, operators, strategies)
}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetQueuedWithdrawal(opts *bind.CallOpts, withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getQueuedWithdrawal", withdrawalRoot)

	outstruct := new(struct {
		Withdrawal IDelegationManagerTypesWithdrawal
		Shares     []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Withdrawal = *abi.ConvertType(out[0], new(IDelegationManagerTypesWithdrawal)).(*IDelegationManagerTypesWithdrawal)
	outstruct.Shares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawal(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// GetQueuedWithdrawal is a free data retrieval call binding the contract method 0x5d975e88.
//
// Solidity: function getQueuedWithdrawal(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] shares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetQueuedWithdrawal(withdrawalRoot [32]byte) (struct {
	Withdrawal IDelegationManagerTypesWithdrawal
	Shares     []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawal(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetQueuedWithdrawalRoots(opts *bind.CallOpts, staker common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getQueuedWithdrawalRoots", staker)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerSession) GetQueuedWithdrawalRoots(staker common.Address) ([][32]byte, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawalRoots(&_ContractDelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawalRoots is a free data retrieval call binding the contract method 0xfd8aa88d.
//
// Solidity: function getQueuedWithdrawalRoots(address staker) view returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetQueuedWithdrawalRoots(staker common.Address) ([][32]byte, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawalRoots(&_ContractDelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetQueuedWithdrawals(opts *bind.CallOpts, staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getQueuedWithdrawals", staker)

	outstruct := new(struct {
		Withdrawals []IDelegationManagerTypesWithdrawal
		Shares      [][]*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Withdrawals = *abi.ConvertType(out[0], new([]IDelegationManagerTypesWithdrawal)).(*[]IDelegationManagerTypesWithdrawal)
	outstruct.Shares = *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return *outstruct, err

}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawals(&_ContractDelegationManager.CallOpts, staker)
}

// GetQueuedWithdrawals is a free data retrieval call binding the contract method 0x5dd68579.
//
// Solidity: function getQueuedWithdrawals(address staker) view returns((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, uint256[][] shares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetQueuedWithdrawals(staker common.Address) (struct {
	Withdrawals []IDelegationManagerTypesWithdrawal
	Shares      [][]*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetQueuedWithdrawals(&_ContractDelegationManager.CallOpts, staker)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetSlashableSharesInQueue(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getSlashableSharesInQueue", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.GetSlashableSharesInQueue(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// GetSlashableSharesInQueue is a free data retrieval call binding the contract method 0x6e174448.
//
// Solidity: function getSlashableSharesInQueue(address operator, address strategy) view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetSlashableSharesInQueue(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.GetSlashableSharesInQueue(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) GetWithdrawableShares(opts *bind.CallOpts, staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "getWithdrawableShares", staker, strategies)

	outstruct := new(struct {
		WithdrawableShares []*big.Int
		DepositShares      []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawableShares = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.DepositShares = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetWithdrawableShares(&_ContractDelegationManager.CallOpts, staker, strategies)
}

// GetWithdrawableShares is a free data retrieval call binding the contract method 0xc978f7ac.
//
// Solidity: function getWithdrawableShares(address staker, address[] strategies) view returns(uint256[] withdrawableShares, uint256[] depositShares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) GetWithdrawableShares(staker common.Address, strategies []common.Address) (struct {
	WithdrawableShares []*big.Int
	DepositShares      []*big.Int
}, error) {
	return _ContractDelegationManager.Contract.GetWithdrawableShares(&_ContractDelegationManager.CallOpts, staker, strategies)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsDelegated(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isDelegated", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsDelegated is a free data retrieval call binding the contract method 0x3e28391d.
//
// Solidity: function isDelegated(address staker) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsDelegated(staker common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsDelegated(&_ContractDelegationManager.CallOpts, staker)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) IsOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "isOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address operator) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) IsOperator(operator common.Address) (bool, error) {
	return _ContractDelegationManager.Contract.IsOperator(&_ContractDelegationManager.CallOpts, operator)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerCaller) MinWithdrawalDelayBlocks(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "minWithdrawalDelayBlocks")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _ContractDelegationManager.Contract.MinWithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// MinWithdrawalDelayBlocks is a free data retrieval call binding the contract method 0xc448feb8.
//
// Solidity: function minWithdrawalDelayBlocks() view returns(uint32)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) MinWithdrawalDelayBlocks() (uint32, error) {
	return _ContractDelegationManager.Contract.MinWithdrawalDelayBlocks(&_ContractDelegationManager.CallOpts)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerCaller) OperatorShares(opts *bind.CallOpts, operator common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "operatorShares", operator, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// OperatorShares is a free data retrieval call binding the contract method 0x778e55f3.
//
// Solidity: function operatorShares(address operator, address strategy) view returns(uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) OperatorShares(operator common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractDelegationManager.Contract.OperatorShares(&_ContractDelegationManager.CallOpts, operator, strategy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Owner() (common.Address, error) {
	return _ContractDelegationManager.Contract.Owner(&_ContractDelegationManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractDelegationManager.Contract.Paused(&_ContractDelegationManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractDelegationManager.Contract.Paused0(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractDelegationManager.Contract.PauserRegistry(&_ContractDelegationManager.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PendingWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "pendingWithdrawals", withdrawalRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xb7f06ebe.
//
// Solidity: function pendingWithdrawals(bytes32 withdrawalRoot) view returns(bool pending)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PendingWithdrawals(withdrawalRoot [32]byte) (bool, error) {
	return _ContractDelegationManager.Contract.PendingWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) PermissionController() (common.Address, error) {
	return _ContractDelegationManager.Contract.PermissionController(&_ContractDelegationManager.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) PermissionController() (common.Address, error) {
	return _ContractDelegationManager.Contract.PermissionController(&_ContractDelegationManager.CallOpts)
}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerCaller) QueuedWithdrawals(opts *bind.CallOpts, withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "queuedWithdrawals", withdrawalRoot)

	if err != nil {
		return *new(IDelegationManagerTypesWithdrawal), err
	}

	out0 := *abi.ConvertType(out[0], new(IDelegationManagerTypesWithdrawal)).(*IDelegationManagerTypesWithdrawal)

	return out0, err

}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _ContractDelegationManager.Contract.QueuedWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// QueuedWithdrawals is a free data retrieval call binding the contract method 0x99f5371b.
//
// Solidity: function queuedWithdrawals(bytes32 withdrawalRoot) view returns((address,address,address,uint256,uint32,address[],uint256[]) withdrawal)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) QueuedWithdrawals(withdrawalRoot [32]byte) (IDelegationManagerTypesWithdrawal, error) {
	return _ContractDelegationManager.Contract.QueuedWithdrawals(&_ContractDelegationManager.CallOpts, withdrawalRoot)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) StrategyManager() (common.Address, error) {
	return _ContractDelegationManager.Contract.StrategyManager(&_ContractDelegationManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractDelegationManager *ContractDelegationManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractDelegationManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractDelegationManager *ContractDelegationManagerSession) Version() (string, error) {
	return _ContractDelegationManager.Contract.Version(&_ContractDelegationManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractDelegationManager *ContractDelegationManagerCallerSession) Version() (string, error) {
	return _ContractDelegationManager.Contract.Version(&_ContractDelegationManager.CallOpts)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawal(opts *bind.TransactOpts, withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawal", withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawal is a paid mutator transaction binding the contract method 0xe4cc3f90.
//
// Solidity: function completeQueuedWithdrawal((address,address,address,uint256,uint32,address[],uint256[]) withdrawal, address[] tokens, bool receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawal(withdrawal IDelegationManagerTypesWithdrawal, tokens []common.Address, receiveAsTokens bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawal(&_ContractDelegationManager.TransactOpts, withdrawal, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x9435bb43.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, bool[] receiveAsTokens) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerTypesWithdrawal, tokens [][]common.Address, receiveAsTokens []bool) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.CompleteQueuedWithdrawals(&_ContractDelegationManager.TransactOpts, withdrawals, tokens, receiveAsTokens)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DecreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "decreaseDelegatedShares", staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DecreaseDelegatedShares is a paid mutator transaction binding the contract method 0x60a0d1ce.
//
// Solidity: function decreaseDelegatedShares(address staker, uint256 curDepositShares, uint64 beaconChainSlashingFactorDecrease) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DecreaseDelegatedShares(staker common.Address, curDepositShares *big.Int, beaconChainSlashingFactorDecrease uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DecreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, curDepositShares, beaconChainSlashingFactorDecrease)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) DelegateTo(opts *bind.TransactOpts, operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "delegateTo", operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// DelegateTo is a paid mutator transaction binding the contract method 0xeea9064b.
//
// Solidity: function delegateTo(address operator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) DelegateTo(operator common.Address, approverSignatureAndExpiry ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.DelegateTo(&_ContractDelegationManager.TransactOpts, operator, approverSignatureAndExpiry, approverSalt)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) IncreaseDelegatedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "increaseDelegatedShares", staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// IncreaseDelegatedShares is a paid mutator transaction binding the contract method 0x3c651cf2.
//
// Solidity: function increaseDelegatedShares(address staker, address strategy, uint256 prevDepositShares, uint256 addedShares) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) IncreaseDelegatedShares(staker common.Address, strategy common.Address, prevDepositShares *big.Int, addedShares *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.IncreaseDelegatedShares(&_ContractDelegationManager.TransactOpts, staker, strategy, prevDepositShares, addedShares)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Initialize(&_ContractDelegationManager.TransactOpts, initialOwner, initialPausedStatus)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) ModifyOperatorDetails(opts *bind.TransactOpts, operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "modifyOperatorDetails", operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, operator, newDelegationApprover)
}

// ModifyOperatorDetails is a paid mutator transaction binding the contract method 0x54b7c96c.
//
// Solidity: function modifyOperatorDetails(address operator, address newDelegationApprover) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) ModifyOperatorDetails(operator common.Address, newDelegationApprover common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.ModifyOperatorDetails(&_ContractDelegationManager.TransactOpts, operator, newDelegationApprover)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Pause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.PauseAll(&_ContractDelegationManager.TransactOpts)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactor) QueueWithdrawals(opts *bind.TransactOpts, params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "queueWithdrawals", params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, params)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] params) returns(bytes32[])
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) QueueWithdrawals(params []IDelegationManagerTypesQueuedWithdrawalParams) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.QueueWithdrawals(&_ContractDelegationManager.TransactOpts, params)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Redelegate(opts *bind.TransactOpts, newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "redelegate", newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Redelegate(&_ContractDelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// Redelegate is a paid mutator transaction binding the contract method 0xa33a3433.
//
// Solidity: function redelegate(address newOperator, (bytes,uint256) newOperatorApproverSig, bytes32 approverSalt) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Redelegate(newOperator common.Address, newOperatorApproverSig ISignatureUtilsMixinTypesSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Redelegate(&_ContractDelegationManager.TransactOpts, newOperator, newOperatorApproverSig, approverSalt)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RegisterAsOperator(opts *bind.TransactOpts, initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "registerAsOperator", initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
}

// RegisterAsOperator is a paid mutator transaction binding the contract method 0x2aa6d888.
//
// Solidity: function registerAsOperator(address initDelegationApprover, uint32 allocationDelay, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RegisterAsOperator(initDelegationApprover common.Address, allocationDelay uint32, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RegisterAsOperator(&_ContractDelegationManager.TransactOpts, initDelegationApprover, allocationDelay, metadataURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.RenounceOwnership(&_ContractDelegationManager.TransactOpts)
}

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x601bb36f.
//
// Solidity: function slashOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) SlashOperatorShares(opts *bind.TransactOpts, operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "slashOperatorShares", operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x601bb36f.
//
// Solidity: function slashOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) SlashOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SlashOperatorShares(&_ContractDelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// SlashOperatorShares is a paid mutator transaction binding the contract method 0x601bb36f.
//
// Solidity: function slashOperatorShares(address operator, address strategy, uint64 prevMaxMagnitude, uint64 newMaxMagnitude) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) SlashOperatorShares(operator common.Address, strategy common.Address, prevMaxMagnitude uint64, newMaxMagnitude uint64) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.SlashOperatorShares(&_ContractDelegationManager.TransactOpts, operator, strategy, prevMaxMagnitude, newMaxMagnitude)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.TransferOwnership(&_ContractDelegationManager.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Undelegate(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "undelegate", staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Undelegate is a paid mutator transaction binding the contract method 0xda8be864.
//
// Solidity: function undelegate(address staker) returns(bytes32[] withdrawalRoots)
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Undelegate(staker common.Address) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Undelegate(&_ContractDelegationManager.TransactOpts, staker)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.Unpause(&_ContractDelegationManager.TransactOpts, newPausedStatus)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactor) UpdateOperatorMetadataURI(opts *bind.TransactOpts, operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.contract.Transact(opts, "updateOperatorMetadataURI", operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, operator, metadataURI)
}

// UpdateOperatorMetadataURI is a paid mutator transaction binding the contract method 0x78296ec5.
//
// Solidity: function updateOperatorMetadataURI(address operator, string metadataURI) returns()
func (_ContractDelegationManager *ContractDelegationManagerTransactorSession) UpdateOperatorMetadataURI(operator common.Address, metadataURI string) (*types.Transaction, error) {
	return _ContractDelegationManager.Contract.UpdateOperatorMetadataURI(&_ContractDelegationManager.TransactOpts, operator, metadataURI)
}

// ContractDelegationManagerDelegationApproverUpdatedIterator is returned from FilterDelegationApproverUpdated and is used to iterate over the raw logs and unpacked data for DelegationApproverUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerDelegationApproverUpdatedIterator struct {
	Event *ContractDelegationManagerDelegationApproverUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerDelegationApproverUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerDelegationApproverUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerDelegationApproverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerDelegationApproverUpdated represents a DelegationApproverUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerDelegationApproverUpdated struct {
	Operator              common.Address
	NewDelegationApprover common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDelegationApproverUpdated is a free log retrieval operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterDelegationApproverUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerDelegationApproverUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerDelegationApproverUpdatedIterator{contract: _ContractDelegationManager.contract, event: "DelegationApproverUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationApproverUpdated is a free log subscription operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchDelegationApproverUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDelegationApproverUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "DelegationApproverUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerDelegationApproverUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegationApproverUpdated is a log parse operation binding the contract event 0x773b54c04d756fcc5e678111f7d730de3be98192000799eee3d63716055a87c6.
//
// Solidity: event DelegationApproverUpdated(address indexed operator, address newDelegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseDelegationApproverUpdated(log types.Log) (*ContractDelegationManagerDelegationApproverUpdated, error) {
	event := new(ContractDelegationManagerDelegationApproverUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "DelegationApproverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerDepositScalingFactorUpdatedIterator is returned from FilterDepositScalingFactorUpdated and is used to iterate over the raw logs and unpacked data for DepositScalingFactorUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerDepositScalingFactorUpdatedIterator struct {
	Event *ContractDelegationManagerDepositScalingFactorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerDepositScalingFactorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerDepositScalingFactorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerDepositScalingFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerDepositScalingFactorUpdated represents a DepositScalingFactorUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerDepositScalingFactorUpdated struct {
	Staker                  common.Address
	Strategy                common.Address
	NewDepositScalingFactor *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDepositScalingFactorUpdated is a free log retrieval operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterDepositScalingFactorUpdated(opts *bind.FilterOpts) (*ContractDelegationManagerDepositScalingFactorUpdatedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerDepositScalingFactorUpdatedIterator{contract: _ContractDelegationManager.contract, event: "DepositScalingFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchDepositScalingFactorUpdated is a free log subscription operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchDepositScalingFactorUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerDepositScalingFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "DepositScalingFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerDepositScalingFactorUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositScalingFactorUpdated is a log parse operation binding the contract event 0x8be932bac54561f27260f95463d9b8ab37e06b2842e5ee2404157cc13df6eb8f.
//
// Solidity: event DepositScalingFactorUpdated(address staker, address strategy, uint256 newDepositScalingFactor)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseDepositScalingFactorUpdated(log types.Log) (*ContractDelegationManagerDepositScalingFactorUpdated, error) {
	event := new(ContractDelegationManagerDepositScalingFactorUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "DepositScalingFactorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitializedIterator struct {
	Event *ContractDelegationManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerInitialized represents a Initialized event raised by the ContractDelegationManager contract.
type ContractDelegationManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractDelegationManagerInitializedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerInitializedIterator{contract: _ContractDelegationManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerInitialized)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseInitialized(log types.Log) (*ContractDelegationManagerInitialized, error) {
	event := new(ContractDelegationManagerInitialized)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorMetadataURIUpdatedIterator is returned from FilterOperatorMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for OperatorMetadataURIUpdated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdatedIterator struct {
	Event *ContractDelegationManagerOperatorMetadataURIUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOperatorMetadataURIUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorMetadataURIUpdated represents a OperatorMetadataURIUpdated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorMetadataURIUpdated struct {
	Operator    common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorMetadataURIUpdated is a free log retrieval operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorMetadataURIUpdated(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorMetadataURIUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorMetadataURIUpdatedIterator{contract: _ContractDelegationManager.contract, event: "OperatorMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMetadataURIUpdated is a free log subscription operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorMetadataURIUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorMetadataURIUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorMetadataURIUpdated is a log parse operation binding the contract event 0x02a919ed0e2acad1dd90f17ef2fa4ae5462ee1339170034a8531cca4b6708090.
//
// Solidity: event OperatorMetadataURIUpdated(address indexed operator, string metadataURI)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorMetadataURIUpdated(log types.Log) (*ContractDelegationManagerOperatorMetadataURIUpdated, error) {
	event := new(ContractDelegationManagerOperatorMetadataURIUpdated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegisteredIterator struct {
	Event *ContractDelegationManagerOperatorRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOperatorRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorRegistered represents a OperatorRegistered event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorRegistered struct {
	Operator           common.Address
	DelegationApprover common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorRegisteredIterator{contract: _ContractDelegationManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorRegistered)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorRegistered is a log parse operation binding the contract event 0xa453db612af59e5521d6ab9284dc3e2d06af286eb1b1b7b771fce4716c19f2c1.
//
// Solidity: event OperatorRegistered(address indexed operator, address delegationApprover)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorRegistered(log types.Log) (*ContractDelegationManagerOperatorRegistered, error) {
	event := new(ContractDelegationManagerOperatorRegistered)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesDecreasedIterator is returned from FilterOperatorSharesDecreased and is used to iterate over the raw logs and unpacked data for OperatorSharesDecreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesDecreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOperatorSharesDecreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesDecreased represents a OperatorSharesDecreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesDecreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesDecreased is a free log retrieval operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesDecreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesDecreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesDecreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesDecreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesDecreased is a free log subscription operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesDecreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesDecreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesDecreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesDecreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSharesDecreased is a log parse operation binding the contract event 0x6909600037b75d7b4733aedd815442b5ec018a827751c832aaff64eba5d6d2dd.
//
// Solidity: event OperatorSharesDecreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesDecreased(log types.Log) (*ContractDelegationManagerOperatorSharesDecreased, error) {
	event := new(ContractDelegationManagerOperatorSharesDecreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesIncreasedIterator is returned from FilterOperatorSharesIncreased and is used to iterate over the raw logs and unpacked data for OperatorSharesIncreased events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreasedIterator struct {
	Event *ContractDelegationManagerOperatorSharesIncreased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOperatorSharesIncreased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesIncreased represents a OperatorSharesIncreased event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesIncreased struct {
	Operator common.Address
	Staker   common.Address
	Strategy common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesIncreased is a free log retrieval operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesIncreased(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesIncreasedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesIncreasedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesIncreased", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesIncreased is a free log subscription operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesIncreased(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesIncreased, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesIncreased", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesIncreased)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSharesIncreased is a log parse operation binding the contract event 0x1ec042c965e2edd7107b51188ee0f383e22e76179041ab3a9d18ff151405166c.
//
// Solidity: event OperatorSharesIncreased(address indexed operator, address staker, address strategy, uint256 shares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesIncreased(log types.Log) (*ContractDelegationManagerOperatorSharesIncreased, error) {
	event := new(ContractDelegationManagerOperatorSharesIncreased)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOperatorSharesSlashedIterator is returned from FilterOperatorSharesSlashed and is used to iterate over the raw logs and unpacked data for OperatorSharesSlashed events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesSlashedIterator struct {
	Event *ContractDelegationManagerOperatorSharesSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOperatorSharesSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOperatorSharesSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOperatorSharesSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOperatorSharesSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOperatorSharesSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOperatorSharesSlashed represents a OperatorSharesSlashed event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOperatorSharesSlashed struct {
	Operator           common.Address
	Strategy           common.Address
	TotalSlashedShares *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOperatorSharesSlashed is a free log retrieval operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOperatorSharesSlashed(opts *bind.FilterOpts, operator []common.Address) (*ContractDelegationManagerOperatorSharesSlashedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OperatorSharesSlashed", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOperatorSharesSlashedIterator{contract: _ContractDelegationManager.contract, event: "OperatorSharesSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSharesSlashed is a free log subscription operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOperatorSharesSlashed(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOperatorSharesSlashed, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OperatorSharesSlashed", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOperatorSharesSlashed)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorSharesSlashed is a log parse operation binding the contract event 0xdd611f4ef63f4385f1756c86ce1f1f389a9013ba6fa07daba8528291bc2d3c30.
//
// Solidity: event OperatorSharesSlashed(address indexed operator, address strategy, uint256 totalSlashedShares)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOperatorSharesSlashed(log types.Log) (*ContractDelegationManagerOperatorSharesSlashed, error) {
	event := new(ContractDelegationManagerOperatorSharesSlashed)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OperatorSharesSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferredIterator struct {
	Event *ContractDelegationManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractDelegationManager contract.
type ContractDelegationManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractDelegationManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerOwnershipTransferredIterator{contract: _ContractDelegationManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerOwnershipTransferred)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractDelegationManagerOwnershipTransferred, error) {
	event := new(ContractDelegationManagerOwnershipTransferred)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerPausedIterator struct {
	Event *ContractDelegationManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerPaused represents a Paused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerPausedIterator{contract: _ContractDelegationManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerPaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParsePaused(log types.Log) (*ContractDelegationManagerPaused, error) {
	event := new(ContractDelegationManagerPaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerSlashingWithdrawalCompletedIterator is returned from FilterSlashingWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalCompleted events raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalCompletedIterator struct {
	Event *ContractDelegationManagerSlashingWithdrawalCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerSlashingWithdrawalCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerSlashingWithdrawalCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerSlashingWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerSlashingWithdrawalCompleted represents a SlashingWithdrawalCompleted event raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalCompleted struct {
	WithdrawalRoot [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalCompleted is a free log retrieval operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterSlashingWithdrawalCompleted(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalCompletedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerSlashingWithdrawalCompletedIterator{contract: _ContractDelegationManager.contract, event: "SlashingWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalCompleted is a free log subscription operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchSlashingWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalCompleted) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerSlashingWithdrawalCompleted)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashingWithdrawalCompleted is a log parse operation binding the contract event 0x1f40400889274ed07b24845e5054a87a0cab969eb1277aafe61ae352e7c32a00.
//
// Solidity: event SlashingWithdrawalCompleted(bytes32 withdrawalRoot)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseSlashingWithdrawalCompleted(log types.Log) (*ContractDelegationManagerSlashingWithdrawalCompleted, error) {
	event := new(ContractDelegationManagerSlashingWithdrawalCompleted)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerSlashingWithdrawalQueuedIterator is returned from FilterSlashingWithdrawalQueued and is used to iterate over the raw logs and unpacked data for SlashingWithdrawalQueued events raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalQueuedIterator struct {
	Event *ContractDelegationManagerSlashingWithdrawalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerSlashingWithdrawalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerSlashingWithdrawalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerSlashingWithdrawalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerSlashingWithdrawalQueued represents a SlashingWithdrawalQueued event raised by the ContractDelegationManager contract.
type ContractDelegationManagerSlashingWithdrawalQueued struct {
	WithdrawalRoot   [32]byte
	Withdrawal       IDelegationManagerTypesWithdrawal
	SharesToWithdraw []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashingWithdrawalQueued is a free log retrieval operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterSlashingWithdrawalQueued(opts *bind.FilterOpts) (*ContractDelegationManagerSlashingWithdrawalQueuedIterator, error) {

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerSlashingWithdrawalQueuedIterator{contract: _ContractDelegationManager.contract, event: "SlashingWithdrawalQueued", logs: logs, sub: sub}, nil
}

// WatchSlashingWithdrawalQueued is a free log subscription operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchSlashingWithdrawalQueued(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerSlashingWithdrawalQueued) (event.Subscription, error) {

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "SlashingWithdrawalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerSlashingWithdrawalQueued)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashingWithdrawalQueued is a log parse operation binding the contract event 0x26b2aae26516e8719ef50ea2f6831a2efbd4e37dccdf0f6936b27bc08e793e30.
//
// Solidity: event SlashingWithdrawalQueued(bytes32 withdrawalRoot, (address,address,address,uint256,uint32,address[],uint256[]) withdrawal, uint256[] sharesToWithdraw)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseSlashingWithdrawalQueued(log types.Log) (*ContractDelegationManagerSlashingWithdrawalQueued, error) {
	event := new(ContractDelegationManagerSlashingWithdrawalQueued)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "SlashingWithdrawalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerDelegatedIterator is returned from FilterStakerDelegated and is used to iterate over the raw logs and unpacked data for StakerDelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegatedIterator struct {
	Event *ContractDelegationManagerStakerDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerStakerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerStakerDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerStakerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerDelegated represents a StakerDelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerDelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerDelegated is a free log retrieval operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerDelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerDelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerDelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerDelegated", logs: logs, sub: sub}, nil
}

// WatchStakerDelegated is a free log subscription operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerDelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerDelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerDelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerDelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerDelegated is a log parse operation binding the contract event 0xc3ee9f2e5fda98e8066a1f745b2df9285f416fe98cf2559cd21484b3d8743304.
//
// Solidity: event StakerDelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerDelegated(log types.Log) (*ContractDelegationManagerStakerDelegated, error) {
	event := new(ContractDelegationManagerStakerDelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerForceUndelegatedIterator is returned from FilterStakerForceUndelegated and is used to iterate over the raw logs and unpacked data for StakerForceUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerForceUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerForceUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerStakerForceUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerForceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerForceUndelegated represents a StakerForceUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerForceUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerForceUndelegated is a free log retrieval operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerForceUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerForceUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerForceUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerForceUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerForceUndelegated is a free log subscription operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerForceUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerForceUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerForceUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerForceUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerForceUndelegated is a log parse operation binding the contract event 0xf0eddf07e6ea14f388b47e1e94a0f464ecbd9eed4171130e0fc0e99fb4030a8a.
//
// Solidity: event StakerForceUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerForceUndelegated(log types.Log) (*ContractDelegationManagerStakerForceUndelegated, error) {
	event := new(ContractDelegationManagerStakerForceUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerForceUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerStakerUndelegatedIterator is returned from FilterStakerUndelegated and is used to iterate over the raw logs and unpacked data for StakerUndelegated events raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegatedIterator struct {
	Event *ContractDelegationManagerStakerUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerStakerUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerStakerUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerStakerUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerStakerUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerStakerUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerStakerUndelegated represents a StakerUndelegated event raised by the ContractDelegationManager contract.
type ContractDelegationManagerStakerUndelegated struct {
	Staker   common.Address
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakerUndelegated is a free log retrieval operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterStakerUndelegated(opts *bind.FilterOpts, staker []common.Address, operator []common.Address) (*ContractDelegationManagerStakerUndelegatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerStakerUndelegatedIterator{contract: _ContractDelegationManager.contract, event: "StakerUndelegated", logs: logs, sub: sub}, nil
}

// WatchStakerUndelegated is a free log subscription operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchStakerUndelegated(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerStakerUndelegated, staker []common.Address, operator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "StakerUndelegated", stakerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerStakerUndelegated)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakerUndelegated is a log parse operation binding the contract event 0xfee30966a256b71e14bc0ebfc94315e28ef4a97a7131a9e2b7a310a73af44676.
//
// Solidity: event StakerUndelegated(address indexed staker, address indexed operator)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseStakerUndelegated(log types.Log) (*ContractDelegationManagerStakerUndelegated, error) {
	event := new(ContractDelegationManagerStakerUndelegated)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "StakerUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelegationManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpausedIterator struct {
	Event *ContractDelegationManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelegationManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelegationManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelegationManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelegationManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelegationManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelegationManagerUnpaused represents a Unpaused event raised by the ContractDelegationManager contract.
type ContractDelegationManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractDelegationManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelegationManagerUnpausedIterator{contract: _ContractDelegationManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractDelegationManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractDelegationManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelegationManagerUnpaused)
				if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractDelegationManager *ContractDelegationManagerFilterer) ParseUnpaused(log types.Log) (*ContractDelegationManagerUnpaused, error) {
	event := new(ContractDelegationManagerUnpaused)
	if err := _ContractDelegationManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
