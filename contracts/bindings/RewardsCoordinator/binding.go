// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractRewardsCoordinator

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

// IRewardsCoordinatorTypesDistributionRoot is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesDistributionRoot struct {
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Disabled                       bool
}

// IRewardsCoordinatorTypesEarnerTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesEarnerTreeMerkleLeaf struct {
	Earner          common.Address
	EarnerTokenRoot [32]byte
}

// IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	OperatorRewards          []IRewardsCoordinatorTypesOperatorReward
	StartTimestamp           uint32
	Duration                 uint32
	Description              string
}

// IRewardsCoordinatorTypesOperatorReward is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesOperatorReward struct {
	Operator common.Address
	Amount   *big.Int
}

// IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams struct {
	DelegationManager          common.Address
	StrategyManager            common.Address
	AllocationManager          common.Address
	PauserRegistry             common.Address
	PermissionController       common.Address
	CALCULATIONINTERVALSECONDS uint32
	MAXREWARDSDURATION         uint32
	MAXRETROACTIVELENGTH       uint32
	MAXFUTURELENGTH            uint32
	GENESISREWARDSTIMESTAMP    uint32
	Version                    string
}

// IRewardsCoordinatorTypesRewardsMerkleClaim is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsMerkleClaim struct {
	RootIndex       uint32
	EarnerIndex     uint32
	EarnerTreeProof []byte
	EarnerLeaf      IRewardsCoordinatorTypesEarnerTreeMerkleLeaf
	TokenIndices    []uint32
	TokenTreeProofs [][]byte
	TokenLeaves     []IRewardsCoordinatorTypesTokenTreeMerkleLeaf
}

// IRewardsCoordinatorTypesRewardsSubmission is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesRewardsSubmission struct {
	StrategiesAndMultipliers []IRewardsCoordinatorTypesStrategyAndMultiplier
	Token                    common.Address
	Amount                   *big.Int
	StartTimestamp           uint32
	Duration                 uint32
}

// IRewardsCoordinatorTypesStrategyAndMultiplier is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesStrategyAndMultiplier struct {
	Strategy   common.Address
	Multiplier *big.Int
}

// IRewardsCoordinatorTypesTokenTreeMerkleLeaf is an auto generated low-level Go binding around an user-defined struct.
type IRewardsCoordinatorTypesTokenTreeMerkleLeaf struct {
	Token              common.Address
	CumulativeEarnings *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs common.Address
	Id  uint32
}

// ContractRewardsCoordinatorMetaData contains all meta data concerning the ContractRewardsCoordinator contract.
var ContractRewardsCoordinatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsCoordinatorConstructorParams\",\"components\":[{\"name\":\"delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"permissionController\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"},{\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_REWARDS_DURATION\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_RETROACTIVE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"MAX_FUTURE_LENGTH\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CALCULATION_INTERVAL_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GENESIS_REWARDS_TIMESTAMP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUTURE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_RETROACTIVE_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REWARDS_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activationDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateEarnerLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"calculateTokenLeafHash\",\"inputs\":[{\"name\":\"leaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"checkClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAVSRewardsSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedAVSRewardsSubmission\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOperatorDirectedOperatorSetRewardsSubmission\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operatorDirectedRewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllEarners\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRewardsForAllSubmission\",\"inputs\":[{\"name\":\"rewardsSubmissions\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission[]\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeClaimed\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"outputs\":[{\"name\":\"totalClaimed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currRewardsCalculationEndTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultOperatorSplitBips\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableRoot\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentClaimableDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentDistributionRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootAtIndex\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.DistributionRoot\",\"components\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"disabled\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDistributionRootsLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRootIndexFromHash\",\"inputs\":[{\"name\":\"rootHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_defaultSplitBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedAVSRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperatorDirectedOperatorSetRewardsSubmissionHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllEarnersHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRewardsSubmissionForAllHash\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permissionController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPermissionController\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processClaim\",\"inputs\":[{\"name\":\"claim\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processClaims\",\"inputs\":[{\"name\":\"claims\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.RewardsMerkleClaim[]\",\"components\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"earnerTreeProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"earnerLeaf\",\"type\":\"tuple\",\"internalType\":\"structIRewardsCoordinatorTypes.EarnerTreeMerkleLeaf\",\"components\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"earnerTokenRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"tokenIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"tokenTreeProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"tokenLeaves\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.TokenTreeMerkleLeaf[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"cumulativeEarnings\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardsUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setActivationDelay\",\"inputs\":[{\"name\":\"_activationDelay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClaimerFor\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultOperatorSplit\",\"inputs\":[{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorAVSSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorPISplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOperatorSetSplit\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"split\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsForAllSubmitter\",\"inputs\":[{\"name\":\"_submitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_newValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardsUpdater\",\"inputs\":[{\"name\":\"_rewardsUpdater\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submissionNonce\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"submitRoot\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ActivationDelaySet\",\"inputs\":[{\"name\":\"oldActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"newActivationDelay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimerForSet\",\"inputs\":[{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldClaimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultOperatorSplitBipsSet\",\"inputs\":[{\"name\":\"oldDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newDefaultOperatorSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootDisabled\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DistributionRootSubmitted\",\"inputs\":[{\"name\":\"rootIndex\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsCalculationEndTimestamp\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorAVSSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedAVSRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDirectedOperatorSetRewardsSubmissionCreated\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorDirectedRewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"operatorDirectedRewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.OperatorDirectedRewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"operatorRewards\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.OperatorReward[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorPISplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorPISplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetSplitBipsSet\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"activatedAt\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"oldOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"newOperatorSetSplitBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"earner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsForAllSubmitterSet\",\"inputs\":[{\"name\":\"rewardsForAllSubmitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"newValue\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllCreated\",\"inputs\":[{\"name\":\"submitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsSubmissionForAllEarnersCreated\",\"inputs\":[{\"name\":\"tokenHopper\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"submissionNonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rewardsSubmissionHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewardsSubmission\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIRewardsCoordinatorTypes.RewardsSubmission\",\"components\":[{\"name\":\"strategiesAndMultipliers\",\"type\":\"tuple[]\",\"internalType\":\"structIRewardsCoordinatorTypes.StrategyAndMultiplier[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"multiplier\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"duration\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsUpdaterSet\",\"inputs\":[{\"name\":\"oldRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRewardsUpdater\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AmountExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountIsZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DurationExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EarningsNotGreaterThanClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCalculationIntervalSecondsRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDurationRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEarnerLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidGenesisRewardsTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPermissions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRootIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStartTimestampRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTokenLeafIndex\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewRootMustBeForNewCalculatedPeriod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorsNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreviousSplitPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsEndTimestampNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootAlreadyActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RootNotActivated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SplitExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInFuture\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartTimestampTooFarInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategiesNotInAscendingOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StrategyNotWhitelisted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"SubmissionNotRetroactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedCaller\",\"inputs\":[]}]",
	Bin: "0x6101e060405234801561001157600080fd5b50604051614ec9380380614ec983398101604081905261003091610377565b610140810151608082015182516020840151604085015160a086015160c087015160e08801516101008901516101208a015160608b01516001600160a01b03811661008e576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b03166080526100a48582610497565b63ffffffff16156100c857604051630e06bd3160e01b815260040160405180910390fd5b6100d56201518086610497565b63ffffffff16156100f95760405163223c7b3960e11b815260040160405180910390fd5b6001600160a01b0397881660a05295871660c05293861660e05263ffffffff9283166101005290821661012052811661014052908116610160521661018052166101a05261014681610159565b6101c052506101536101a0565b50610527565b600080829050601f8151111561018d578260405163305a27a960e01b815260040161018491906104cd565b60405180910390fd5b805161019882610500565b179392505050565b600054610100900460ff16156102085760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608401610184565b60005460ff90811614610259576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b634e487b7160e01b600052604160045260246000fd5b60405161016081016001600160401b03811182821017156102945761029461025b565b60405290565b80516001600160a01b03811681146102b157600080fd5b919050565b805163ffffffff811681146102b157600080fd5b60005b838110156102e55781810151838201526020016102cd565b50506000910152565b600082601f8301126102ff57600080fd5b81516001600160401b038111156103185761031861025b565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103465761034661025b565b60405281815283820160200185101561035e57600080fd5b61036f8260208301602087016102ca565b949350505050565b60006020828403121561038957600080fd5b81516001600160401b0381111561039f57600080fd5b820161016081850312156103b257600080fd5b6103ba610271565b6103c38261029a565b81526103d16020830161029a565b60208201526103e26040830161029a565b60408201526103f36060830161029a565b60608201526104046080830161029a565b608082015261041560a083016102b6565b60a082015261042660c083016102b6565b60c082015261043760e083016102b6565b60e082015261044961010083016102b6565b61010082015261045c61012083016102b6565b6101208201526101408201516001600160401b0381111561047c57600080fd5b610488868285016102ee565b61014083015250949350505050565b600063ffffffff8316806104bb57634e487b7160e01b600052601260045260246000fd5b8063ffffffff84160691505092915050565b60208152600082518060208401526104ec8160408501602087016102ca565b601f01601f19169190910160400192915050565b80516020808301519190811015610521576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516101c0516148b361061660003960006114f101526000818161060501526128480152600081816104b8015261364b0152600081816104060152612c94015260008181610567015261360901526000818161089701526135190152600081816107d60152818161356901526135b70152600081816108ec01528181610b1701528181611e7d015261213101526000818161058e01526136e801526000818161095f0152611dea01526000818161076801528181610f0501528181611543015261241b01526148b36000f3fe608060405234801561001057600080fd5b50600436106103c45760003560e01c8063886f1195116101ff578063de02e5031161011a578063f6efbb59116100ad578063fabc1cbc1161007c578063fabc1cbc14610a4f578063fbf1e2c114610a62578063fce36c7d14610a75578063ff9f6cce14610a8857600080fd5b8063f6efbb5914610a03578063f74e8eac14610a16578063f8cd844814610a29578063f96abf2e14610a3c57600080fd5b8063ed71e6a2116100e9578063ed71e6a214610981578063f22cef85146109af578063f2f07ab4146109c2578063f2fde38b146109f057600080fd5b8063de02e50314610921578063e063f81f14610934578063e810ce2114610947578063ea4d3c9b1461095a57600080fd5b8063a50a1d9c11610192578063bf21a8aa11610161578063bf21a8aa14610892578063c46db606146108b9578063ca8aa7c7146108e7578063dcbb03b31461090e57600080fd5b8063a50a1d9c1461081e578063aebd8bae14610831578063b3dbb0e01461085f578063bb7e451f1461087257600080fd5b80639cb9a5fa116101ce5780639cb9a5fa146107be5780639d45c281146107d15780639de4b35f146107f8578063a0169ddd1461080b57600080fd5b8063886f1195146107635780638da5cb5b1461078a5780639104c3191461079b5780639be3d4e4146107b657600080fd5b80634596021c116102ef5780635c975abb11610282578063715018a611610251578063715018a6146107155780637b8f8b051461071d578063863cb9a914610725578063865c69531461073857600080fd5b80635c975abb146106b75780635e9d8348146106bf57806363f6a798146106d25780636d21117e146106e757600080fd5b806354fd4d50116102be57806354fd4d501461066457806358baaa3e14610679578063595c6a671461068c5780635ac86ab71461069457600080fd5b80634596021c146105ed5780634657e26a146106005780634b943960146106275780634d18cc351461064d57600080fd5b8063149bc8721161036757806339b70e381161033657806339b70e38146105895780633a8c0786146105b05780633ccc861d146105c75780633efe1db6146105da57600080fd5b8063149bc872146104ed5780632b9f64a41461050e57806336af41fa1461054f57806337838ed01461056257600080fd5b80630e9a53cf116103a35780630e9a53cf146104525780630eb38345146104a0578063131433b4146104b3578063136439dd146104da57600080fd5b806218572c146103c957806304a0c502146104015780630ca298991461043d575b600080fd5b6103ec6103d7366004613c41565b60d16020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b6104287f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103f8565b61045061044b366004613cc1565b610a9b565b005b61045a610d6a565b6040516103f8919060006080820190508251825263ffffffff602084015116602083015263ffffffff604084015116604083015260608301511515606083015292915050565b6104506104ae366004613d22565b610e6e565b6104287f000000000000000000000000000000000000000000000000000000000000000081565b6104506104e8366004613d5b565b610ef0565b6105006104fb366004613d74565b610fc7565b6040519081526020016103f8565b61053761051c366004613c41565b60cc602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016103f8565b61045061055d366004613d90565b61103d565b6104287f000000000000000000000000000000000000000000000000000000000000000081565b6105377f000000000000000000000000000000000000000000000000000000000000000081565b60cb5461042890600160a01b900463ffffffff1681565b6104506105d5366004613de4565b6111d1565b6104506105e8366004613e3e565b611216565b6104506105fb366004613e6a565b61140c565b6105377f000000000000000000000000000000000000000000000000000000000000000081565b61063a610635366004613c41565b61148e565b60405161ffff90911681526020016103f8565b60cb5461042890600160c01b900463ffffffff1681565b61066c6114ea565b6040516103f89190613ee4565b610450610687366004613f17565b61151a565b61045061152e565b6103ec6106a2366004613f32565b606654600160ff9092169190911b9081161490565b606654610500565b6103ec6106cd366004613f55565b6115e0565b60cb5461063a90600160e01b900461ffff1681565b6103ec6106f5366004613f89565b60cf60209081526000928352604080842090915290825290205460ff1681565b61045061166d565b60ca54610500565b610450610733366004613c41565b61167f565b610500610746366004613fb5565b60cd60209081526000928352604080842090915290825290205481565b6105377f000000000000000000000000000000000000000000000000000000000000000081565b6033546001600160a01b0316610537565b61053773beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b61045a611690565b6104506107cc366004613fe3565b61172e565b6104287f000000000000000000000000000000000000000000000000000000000000000081565b61063a61080636600461401e565b6118cc565b610450610819366004613c41565b611951565b61045061082c36600461405d565b61195c565b6103ec61083f366004613f89565b60d260209081526000928352604080842090915290825290205460ff1681565b61045061086d366004614078565b61196d565b610500610880366004613c41565b60ce6020526000908152604090205481565b6104287f000000000000000000000000000000000000000000000000000000000000000081565b6103ec6108c7366004613f89565b60d060209081526000928352604080842090915290825290205460ff1681565b6105377f000000000000000000000000000000000000000000000000000000000000000081565b61045061091c3660046140a4565b611ab9565b61045a61092f366004613d5b565b611c27565b61063a610942366004613fb5565b611cb9565b610428610955366004613d5b565b611d1f565b6105377f000000000000000000000000000000000000000000000000000000000000000081565b6103ec61098f366004613f89565b60d360209081526000928352604080842090915290825290205460ff1681565b6104506109bd366004613fb5565b611da4565b6103ec6109d0366004613f89565b60d760209081526000928352604080842090915290825290205460ff1681565b6104506109fe366004613c41565b611f13565b610450610a113660046140eb565b611f8e565b610450610a2436600461414d565b6120c9565b610500610a37366004613d74565b6122b5565b610450610a4a366004613f17565b6122c6565b610450610a5d366004613d5b565b612419565b60cb54610537906001600160a01b031681565b610450610a83366004613d90565b612531565b610450610a96366004613d90565b612685565b60665460099061020090811603610ac55760405163840a48d560e01b815260040160405180910390fd5b610ad26020850185613c41565b610adb81612809565b610af85760405163932d94f760e01b815260040160405180910390fd5b610b006128b7565b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc75890610b4c9088906004016141bb565b602060405180830381865afa158015610b69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8d91906141c9565b610baa57604051631fb1705560e21b815260040160405180910390fd5b60005b83811015610d585736858583818110610bc857610bc86141e6565b9050602002810190610bda91906141fc565b9050600060ce81610bee60208b018b613c41565b6001600160a01b03168152602080820192909252604001600090812054925090610c1a908a018a613c41565b8284604051602001610c2e93929190614434565b6040516020818303038152906040528051906020012090506000610c5184612910565b9050600160d76000610c6660208e018e613c41565b6001600160a01b03168152602080820192909252604090810160009081208682529092529020805460ff1916911515919091179055610ca683600161447a565b60ce6000610cb760208e018e613c41565b6001600160a01b03166001600160a01b031681526020019081526020016000208190555081336001600160a01b03167ffff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc7568c8688604051610d199392919061448d565b60405180910390a3610d48333083610d376040890160208a01613c41565b6001600160a01b0316929190612b01565b505060019092019150610bad9050565b50610d636001609755565b5050505050565b60408051608081018252600080825260208201819052918101829052606081019190915260ca545b8015610e4557600060ca610da76001846144b3565b81548110610db757610db76141e6565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161580156060830181905291925090610e275750806040015163ffffffff164210155b15610e325792915050565b5080610e3d816144c6565b915050610d92565b505060408051608081018252600080825260208201819052918101829052606081019190915290565b610e76612b6c565b6001600160a01b038216600081815260d1602052604080822054905160ff9091169284151592841515927f4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c9190a4506001600160a01b0391909116600090815260d160205260409020805460ff1916911515919091179055565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610f54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f7891906141c9565b610f9557604051631d77d47760e21b815260040160405180910390fd5b6066548181168114610fba5760405163c61dca5d60e01b815260040160405180910390fd5b610fc382612bc6565b5050565b600080610fd76020840184613c41565b83602001356040516020016110209392919060f89390931b6001600160f81b031916835260609190911b6bffffffffffffffffffffffff19166001830152601582015260350190565b604051602081830303815290604052805190602001209050919050565b6066546001906002908116036110665760405163840a48d560e01b815260040160405180910390fd5b33600090815260d1602052604090205460ff1661109657604051635c427cd960e01b815260040160405180910390fd5b61109e6128b7565b60005b828110156111c157368484838181106110bc576110bc6141e6565b90506020028101906110ce91906144dd565b33600081815260ce602090815260408083205490519495509391926110f99290918591879101614572565b60405160208183030381529060405280519060200120905061111a83612c03565b33600090815260d0602090815260408083208484529091529020805460ff1916600190811790915561114d90839061447a565b33600081815260ce602052604090819020929092559051829184917f51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf27048290611195908890614599565b60405180910390a46111b6333060408601803590610d379060208901613c41565b5050506001016110a1565b506111cc6001609755565b505050565b6066546002906004908116036111fa5760405163840a48d560e01b815260040160405180910390fd5b6112026128b7565b61120c8383612cef565b6111cc6001609755565b60665460039060089081160361123f5760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b0316331461126a57604051635c427cd960e01b815260040160405180910390fd5b60cb5463ffffffff600160c01b90910481169083161161129d57604051631ca7e50b60e21b815260040160405180910390fd5b428263ffffffff16106112c3576040516306957c9160e11b815260040160405180910390fd5b60ca5460cb546000906112e390600160a01b900463ffffffff16426145ac565b6040805160808101825287815263ffffffff878116602080840182815286841685870181815260006060880181815260ca8054600181018255925297517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee160029092029182015592517f42d72674974f694b5f5159593243114d38a5c39c89d6b62fee061ff523240ee290930180549151975193871667ffffffffffffffff1990921691909117600160201b978716979097029690961760ff60401b1916600160401b921515929092029190911790945560cb805463ffffffff60c01b1916600160c01b840217905593519283529394508892908616917fecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08910160405180910390a45050505050565b6066546002906004908116036114355760405163840a48d560e01b815260040160405180910390fd5b61143d6128b7565b60005b8381101561147d5761147585858381811061145d5761145d6141e6565b905060200281019061146f91906145c8565b84612cef565b600101611440565b506114886001609755565b50505050565b6001600160a01b038116600090815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff16908201526114e490612f81565b92915050565b60606115157f0000000000000000000000000000000000000000000000000000000000000000612ff3565b905090565b611522612b6c565b61152b81613032565b50565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015611592573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b691906141c9565b6115d357604051631d77d47760e21b815260040160405180910390fd5b6115de600019612bc6565b565b60006116658260ca6115f56020830183613f17565b63ffffffff168154811061160b5761160b6141e6565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff16151560608201526130a3565b506001919050565b611675612b6c565b6115de6000613247565b611687612b6c565b61152b81613299565b60408051608081018252600080825260208201819052918101829052606081019190915260ca80546116c4906001906144b3565b815481106116d4576116d46141e6565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff1615156060820152919050565b6066546005906020908116036117575760405163840a48d560e01b815260040160405180910390fd5b8361176181612809565b61177e5760405163932d94f760e01b815260040160405180910390fd5b6117866128b7565b60005b83811015610d5857368585838181106117a4576117a46141e6565b90506020028101906117b691906141fc565b6001600160a01b038816600090815260ce60209081526040808320549051939450926117e8918b918591879101614434565b604051602081830303815290604052805190602001209050600061180b84612910565b6001600160a01b038b16600090815260d3602090815260408083208684529091529020805460ff1916600190811790915590915061184a90849061447a565b6001600160a01b038b16600081815260ce60205260409081902092909255905183919033907ffc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e09061189e9088908a906145de565b60405180910390a46118bc333083610d376040890160208a01613c41565b5050600190920191506117899050565b6001600160a01b038216600090815260d66020526040812061194a90826119006118fb368790038701876145f7565b6132f5565b815260208082019290925260409081016000208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff1690820152612f81565b9392505050565b33610fc3818361335a565b611964612b6c565b61152b816133be565b6066546007906080908116036119965760405163840a48d560e01b815260040160405180910390fd5b826119a081612809565b6119bd5760405163932d94f760e01b815260040160405180910390fd5b60cb546000906119da90600160a01b900463ffffffff16426145ac565b6001600160a01b038616600090815260d5602090815260408083208151606081018352905461ffff80821683526201000082041693820193909352600160201b90920463ffffffff169082015291925090611a3490612f81565b6001600160a01b038716600090815260d560205260409020909150611a5a908684613429565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388169133917fd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f749181900360600190a3505050505050565b606654600690604090811603611ae25760405163840a48d560e01b815260040160405180910390fd5b83611aec81612809565b611b095760405163932d94f760e01b815260040160405180910390fd5b60cb54600090611b2690600160a01b900463ffffffff16426145ac565b6001600160a01b03878116600090815260d460209081526040808320938a1683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff1692810192909252919250611b8e90612f81565b6001600160a01b03808916600090815260d460209081526040808320938b16835292905220909150611bc1908684613429565b6040805163ffffffff8416815261ffff838116602083015287168183015290516001600160a01b0388811692908a169133917f48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934919081900360600190a450505050505050565b60408051608081018252600080825260208201819052918101829052606081019190915260ca8281548110611c5e57611c5e6141e6565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff161515606082015292915050565b6001600160a01b03828116600090815260d46020908152604080832093851683529281528282208351606081018552905461ffff80821683526201000082041692820192909252600160201b90910463ffffffff16928101929092529061194a90612f81565b60ca546000905b63ffffffff811615611d8a578260ca611d40600184614663565b63ffffffff1681548110611d5657611d566141e6565b90600052602060002090600202016000015403611d785761194a600182614663565b80611d828161467f565b915050611d26565b5060405163504570e360e01b815260040160405180910390fd5b81611dae81612809565b611dcb5760405163932d94f760e01b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa158015611e31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e5591906141c9565b80611eec575060405163ba1a84e560e01b81526001600160a01b0384811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063ba1a84e590602401602060405180830381865afa158015611ec6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611eea919061469f565b115b611f095760405163fb494ea160e01b815260040160405180910390fd5b6111cc838361335a565b611f1b612b6c565b6001600160a01b038116611f855760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61152b81613247565b600054610100900460ff1615808015611fae5750600054600160ff909116105b80611fc85750303b158015611fc8575060005460ff166001145b61202b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401611f7c565b6000805460ff19166001179055801561204e576000805461ff0019166101001790555b61205785612bc6565b61206086613247565b61206984613299565b61207283613032565b61207b826133be565b80156120c1576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b606654600890610100908116036120f35760405163840a48d560e01b815260040160405180910390fd5b836120fd81612809565b61211a5760405163932d94f760e01b815260040160405180910390fd5b6040516304c1b8eb60e31b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063260dc758906121669087906004016141bb565b602060405180830381865afa158015612183573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121a791906141c9565b6121c457604051631fb1705560e21b815260040160405180910390fd5b60cb546000906121e190600160a01b900463ffffffff16426145ac565b6001600160a01b038716600090815260d6602052604081209192509061221490826119006118fb368b90038b018b6145f7565b6001600160a01b038816600090815260d66020526040812091925061225b91906122466118fb368b90038b018b6145f7565b81526020019081526020016000208684613429565b866001600160a01b0316336001600160a01b03167f14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f18885858a6040516122a494939291906146b8565b60405180910390a350505050505050565b60006001610fd76020840184613c41565b6066546003906008908116036122ef5760405163840a48d560e01b815260040160405180910390fd5b60cb546001600160a01b0316331461231a57604051635c427cd960e01b815260040160405180910390fd5b60ca5463ffffffff831610612342576040516394a8d38960e01b815260040160405180910390fd5b600060ca8363ffffffff168154811061235d5761235d6141e6565b906000526020600020906002020190508060010160089054906101000a900460ff161561239d57604051631b14174b60e01b815260040160405180910390fd5b6001810154600160201b900463ffffffff1642106123ce57604051630c36f66560e21b815260040160405180910390fd5b60018101805460ff60401b1916600160401b17905560405163ffffffff8416907fd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e90600090a2505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612477573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061249b91906146ec565b6001600160a01b0316336001600160a01b0316146124cc5760405163794821ff60e01b815260040160405180910390fd5b606654801982198116146124f35760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200160405180910390a25050565b60665460009060019081160361255a5760405163840a48d560e01b815260040160405180910390fd5b6125626128b7565b60005b828110156111c15736848483818110612580576125806141e6565b905060200281019061259291906144dd565b33600081815260ce602090815260408083205490519495509391926125bd9290918591879101614572565b6040516020818303038152906040528051906020012090506125de83612c03565b33600090815260cf602090815260408083208484529091529020805460ff1916600190811790915561261190839061447a565b33600081815260ce602052604090819020929092559051829184917f450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e628190612659908890614599565b60405180910390a461267a333060408601803590610d379060208901613c41565b505050600101612565565b6066546004906010908116036126ae5760405163840a48d560e01b815260040160405180910390fd5b33600090815260d1602052604090205460ff166126de57604051635c427cd960e01b815260040160405180910390fd5b6126e66128b7565b60005b828110156111c15736848483818110612704576127046141e6565b905060200281019061271691906144dd565b33600081815260ce602090815260408083205490519495509391926127419290918591879101614572565b60405160208183030381529060405280519060200120905061276283612c03565b33600090815260d2602090815260408083208484529091529020805460ff1916600190811790915561279590839061447a565b33600081815260ce602052604090819020929092559051829184917f5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b906127dd908890614599565b60405180910390a46127fe333060408601803590610d379060208901613c41565b5050506001016126e9565b604051631beb2b9760e31b81526001600160a01b038281166004830152336024830152306044830152600080356001600160e01b0319166064840152917f00000000000000000000000000000000000000000000000000000000000000009091169063df595cb8906084016020604051808303816000875af1158015612893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e491906141c9565b6002609754036129095760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611f7c565b6002609755565b600061294461291f8380614709565b61292f6080860160608701613f17565b61293f60a0870160808801613f17565b6134f9565b60006129536040840184614709565b9050116129735760405163796cc52560e01b815260040160405180910390fd5b4261298460a0840160808501613f17565b6129946080850160608601613f17565b61299e91906145ac565b63ffffffff16106129c25760405163150358a160e21b815260040160405180910390fd5b60008060005b6129d56040860186614709565b9050811015612ac857366129ec6040870187614709565b838181106129fc576129fc6141e6565b60400291909101915060009050612a166020830183613c41565b6001600160a01b031603612a3d57604051630863a45360e11b815260040160405180910390fd5b612a4a6020820182613c41565b6001600160a01b0316836001600160a01b031610612a7b576040516310fb47f160e31b815260040160405180910390fd5b6000816020013511612aa0576040516310eb483f60e21b815260040160405180910390fd5b612aad6020820182613c41565b9250612abd60208201358561447a565b9350506001016129c8565b506f4b3b4ca85a86c47a098a223fffffffff821115612afa5760405163070b5a6f60e21b815260040160405180910390fd5b5092915050565b6040516001600160a01b03808516602483015283166044820152606481018290526114889085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526137d5565b6033546001600160a01b031633146115de5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611f7c565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b612c30612c108280614709565b612c206080850160608601613f17565b61293f60a0860160808701613f17565b6000816040013511612c55576040516310eb483f60e21b815260040160405180910390fd5b6f4b3b4ca85a86c47a098a223fffffffff81604001351115612c8a5760405163070b5a6f60e21b815260040160405180910390fd5b612cba63ffffffff7f0000000000000000000000000000000000000000000000000000000000000000164261447a565b612cca6080830160608401613f17565b63ffffffff16111561152b57604051637ee2b44360e01b815260040160405180910390fd5b600060ca612d006020850185613f17565b63ffffffff1681548110612d1657612d166141e6565b600091825260209182902060408051608081018252600293909302909101805483526001015463ffffffff80821694840194909452600160201b810490931690820152600160401b90910460ff16151560608201529050612d7783826130a3565b6000612d896080850160608601613c41565b6001600160a01b03808216600090815260cc60205260409020549192501680612daf5750805b336001600160a01b03821614612dd857604051635c427cd960e01b815260040160405180910390fd5b60005b612de860a0870187614752565b90508110156120c15736612dff60e0880188614709565b83818110612e0f57612e0f6141e6565b6001600160a01b038716600090815260cd602090815260408083209302949094019450929091508290612e4490850185613c41565b6001600160a01b03166001600160a01b0316815260200190815260200160002054905080826020013511612e8b5760405163aa385e8160e01b815260040160405180910390fd5b6000612e9b8260208501356144b3565b6001600160a01b038716600090815260cd60209081526040822092935085018035929190612ec99087613c41565b6001600160a01b0316815260208082019290925260400160002091909155612f0b9089908390612efb90870187613c41565b6001600160a01b031691906138aa565b86516001600160a01b03808a1691878216918916907f9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce3190612f4f6020890189613c41565b604080519283526001600160a01b039091166020830152810186905260600160405180910390a4505050600101612ddb565b6000816040015163ffffffff1660001480612fb55750815161ffff908116148015612fb55750816040015163ffffffff1642105b15612fcd57505060cb54600160e01b900461ffff1690565b816040015163ffffffff16421015612fe65781516114e4565b506020015190565b919050565b60606000613000836138da565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b60cb546040805163ffffffff600160a01b9093048316815291831660208301527faf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3910160405180910390a160cb805463ffffffff909216600160a01b0263ffffffff60a01b19909216919091179055565b8060600151156130c657604051631b14174b60e01b815260040160405180910390fd5b806040015163ffffffff164210156130f157604051631437a2bb60e31b815260040160405180910390fd5b6130fe60c0830183614752565b905061310d60a0840184614752565b90501461312d576040516343714afd60e01b815260040160405180910390fd5b61313a60e0830183614709565b905061314960c0840184614752565b905014613169576040516343714afd60e01b815260040160405180910390fd5b80516131959061317f6040850160208601613f17565b61318c604086018661479b565b86606001613902565b60005b6131a560a0840184614752565b90508110156111cc5761323f60808401356131c360a0860186614752565b848181106131d3576131d36141e6565b90506020020160208101906131e89190613f17565b6131f560c0870187614752565b85818110613205576132056141e6565b9050602002810190613217919061479b565b61322460e0890189614709565b87818110613234576132346141e6565b9050604002016139a8565b600101613198565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60cb546040516001600160a01b038084169216907f237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb90600090a360cb80546001600160a01b0319166001600160a01b0392909216919091179055565b60008160000151826020015163ffffffff1660405160200161334292919060609290921b6bffffffffffffffffffffffff1916825260a01b6001600160a01b031916601482015260200190565b6040516020818303038152906040526114e4906147e1565b6001600160a01b03808316600081815260cc602052604080822080548686166001600160a01b0319821681179092559151919094169392849290917fbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca3129190a4505050565b60cb546040805161ffff600160e01b9093048316815291831660208301527fe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e910160405180910390a160cb805461ffff909216600160e01b0261ffff60e01b19909216919091179055565b61271061ffff831611156134505760405163891c63df60e01b815260040160405180910390fd5b8254600160201b900463ffffffff16421161347e57604051637b1e25c560e01b815260040160405180910390fd5b8254600160201b900463ffffffff166000036134a657825461ffff191661ffff1783556134bd565b825462010000810461ffff1661ffff199091161783555b825463ffffffff909116600160201b0267ffffffff000000001961ffff90931662010000029290921667ffffffffffff00001990911617179055565b826135175760405163796cc52560e01b815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16111561356457604051630dd0b9f560e21b815260040160405180910390fd5b61358e7f00000000000000000000000000000000000000000000000000000000000000008261481b565b63ffffffff16156135b25760405163ee66470560e01b815260040160405180910390fd5b6135dc7f00000000000000000000000000000000000000000000000000000000000000008361481b565b63ffffffff161561360057604051633c1a94f160e21b815260040160405180910390fd5b8163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff164261363991906144b3565b1115801561367357508163ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1611155b6136905760405163041aa75760e11b815260040160405180910390fd5b6000805b848110156120c15760008686838181106136b0576136b06141e6565b6136c69260206040909202019081019150613c41565b60405163198f077960e21b81526001600160a01b0380831660048301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063663c1de490602401602060405180830381865afa158015613731573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061375591906141c9565b8061377c57506001600160a01b03811673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0145b61379957604051632efd965160e11b815260040160405180910390fd5b806001600160a01b0316836001600160a01b0316106137cb5760405163dfad9ca160e01b815260040160405180910390fd5b9150600101613694565b600061382a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166139e79092919063ffffffff16565b905080516000148061384b57508080602001905181019061384b91906141c9565b6111cc5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401611f7c565b6040516001600160a01b0383166024820152604481018290526111cc90849063a9059cbb60e01b90606401612b35565b600060ff8216601f8111156114e457604051632cd44ac360e21b815260040160405180910390fd5b61390d602083614843565b6001901b8463ffffffff16106139355760405162c6c39d60e71b815260040160405180910390fd5b600061394082610fc7565b905061398b84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a92508591505063ffffffff89166139fe565b6120c1576040516369ca16c960e01b815260040160405180910390fd5b6139b3602083614843565b6001901b8463ffffffff16106139dc5760405163054ff4df60e51b815260040160405180910390fd5b6000613940826122b5565b60606139f68484600085613a16565b949350505050565b600083613a0c868585613af1565b1495945050505050565b606082471015613a775760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401611f7c565b600080866001600160a01b03168587604051613a939190614857565b60006040518083038185875af1925050503d8060008114613ad0576040519150601f19603f3d011682016040523d82523d6000602084013e613ad5565b606091505b5091509150613ae687838387613b8e565b979650505050505050565b600060208451613b019190614869565b15613b1f576040516313717da960e21b815260040160405180910390fd5b8260205b85518111613b8557613b36600285614869565b600003613b5a57816000528086015160205260406000209150600284049350613b73565b8086015160005281602052604060002091506002840493505b613b7e60208261447a565b9050613b23565b50949350505050565b60608315613bfd578251600003613bf6576001600160a01b0385163b613bf65760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401611f7c565b50816139f6565b6139f68383815115613c125781518083602001fd5b8060405162461bcd60e51b8152600401611f7c9190613ee4565b6001600160a01b038116811461152b57600080fd5b600060208284031215613c5357600080fd5b813561194a81613c2c565b600060408284031215613c7057600080fd5b50919050565b60008083601f840112613c8857600080fd5b5081356001600160401b03811115613c9f57600080fd5b6020830191508360208260051b8501011115613cba57600080fd5b9250929050565b600080600060608486031215613cd657600080fd5b613ce08585613c5e565b925060408401356001600160401b03811115613cfb57600080fd5b613d0786828701613c76565b9497909650939450505050565b801515811461152b57600080fd5b60008060408385031215613d3557600080fd5b8235613d4081613c2c565b91506020830135613d5081613d14565b809150509250929050565b600060208284031215613d6d57600080fd5b5035919050565b600060408284031215613d8657600080fd5b61194a8383613c5e565b60008060208385031215613da357600080fd5b82356001600160401b03811115613db957600080fd5b613dc585828601613c76565b90969095509350505050565b60006101008284031215613c7057600080fd5b60008060408385031215613df757600080fd5b82356001600160401b03811115613e0d57600080fd5b613e1985828601613dd1565b9250506020830135613d5081613c2c565b803563ffffffff81168114612fee57600080fd5b60008060408385031215613e5157600080fd5b82359150613e6160208401613e2a565b90509250929050565b600080600060408486031215613e7f57600080fd5b83356001600160401b03811115613e9557600080fd5b613ea186828701613c76565b9094509250506020840135613eb581613c2c565b809150509250925092565b60005b83811015613edb578181015183820152602001613ec3565b50506000910152565b6020815260008251806020840152613f03816040850160208701613ec0565b601f01601f19169190910160400192915050565b600060208284031215613f2957600080fd5b61194a82613e2a565b600060208284031215613f4457600080fd5b813560ff8116811461194a57600080fd5b600060208284031215613f6757600080fd5b81356001600160401b03811115613f7d57600080fd5b6139f684828501613dd1565b60008060408385031215613f9c57600080fd5b8235613fa781613c2c565b946020939093013593505050565b60008060408385031215613fc857600080fd5b8235613fd381613c2c565b91506020830135613d5081613c2c565b600080600060408486031215613ff857600080fd5b833561400381613c2c565b925060208401356001600160401b03811115613cfb57600080fd5b6000806060838503121561403157600080fd5b823561403c81613c2c565b9150613e618460208501613c5e565b803561ffff81168114612fee57600080fd5b60006020828403121561406f57600080fd5b61194a8261404b565b6000806040838503121561408b57600080fd5b823561409681613c2c565b9150613e616020840161404b565b6000806000606084860312156140b957600080fd5b83356140c481613c2c565b925060208401356140d481613c2c565b91506140e26040850161404b565b90509250925092565b600080600080600060a0868803121561410357600080fd5b853561410e81613c2c565b945060208601359350604086013561412581613c2c565b925061413360608701613e2a565b91506141416080870161404b565b90509295509295909350565b60008060006080848603121561416257600080fd5b833561416d81613c2c565b925061417c8560208601613c5e565b91506140e26060850161404b565b803561419581613c2c565b6001600160a01b0316825263ffffffff6141b160208301613e2a565b1660208301525050565b604081016114e4828461418a565b6000602082840312156141db57600080fd5b815161194a81613d14565b634e487b7160e01b600052603260045260246000fd5b6000823560be1983360301811261421257600080fd5b9190910192915050565b6000808335601e1984360301811261423357600080fd5b83016020810192503590506001600160401b0381111561425257600080fd5b8060061b3603821315613cba57600080fd5b81835260208301925060008160005b848110156142ca57813561428681613c2c565b6001600160a01b0316865260208201356bffffffffffffffffffffffff81168082146142b157600080fd5b6020880152506040958601959190910190600101614273565b5093949350505050565b6000808335601e198436030181126142eb57600080fd5b83016020810192503590506001600160401b0381111561430a57600080fd5b803603821315613cba57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600061434e828361421c565b60c0855261436060c086018284614264565b915050602083013561437181613c2c565b6001600160a01b0316602085015261438c604084018461421c565b85830360408701528083529091600091906020015b818310156143dc5783356143b481613c2c565b6001600160a01b031681526020848101359082015260409384019360019390930192016143a1565b6143e860608701613e2a565b63ffffffff81166060890152935061440260808701613e2a565b63ffffffff81166080890152935061441d60a08701876142d4565b9450925086810360a0880152613ae6818585614319565b60018060a01b038416815282602082015260606040820152600061445b6060830184614342565b95945050505050565b634e487b7160e01b600052601160045260246000fd5b808201808211156114e4576114e4614464565b614497818561418a565b82604082015260806060820152600061445b6080830184614342565b818103818111156114e4576114e4614464565b6000816144d5576144d5614464565b506000190190565b60008235609e1983360301811261421257600080fd5b60006144ff828361421c565b60a0855261451160a086018284614264565b915050602083013561452281613c2c565b6001600160a01b031660208501526040838101359085015263ffffffff61454b60608501613e2a565b16606085015263ffffffff61456260808501613e2a565b1660808501528091505092915050565b60018060a01b038416815282602082015260606040820152600061445b60608301846144f3565b60208152600061194a60208301846144f3565b63ffffffff81811683821601908111156114e4576114e4614464565b6000823560fe1983360301811261421257600080fd5b8281526040602082015260006139f66040830184614342565b6000604082840312801561460a57600080fd5b50604080519081016001600160401b038111828210171561463b57634e487b7160e01b600052604160045260246000fd5b604052823561464981613c2c565b815261465760208401613e2a565b60208201529392505050565b63ffffffff82811682821603908111156114e4576114e4614464565b600063ffffffff82168061469557614695614464565b6000190192915050565b6000602082840312156146b157600080fd5b5051919050565b60a081016146c6828761418a565b63ffffffff94909416604082015261ffff92831660608201529116608090910152919050565b6000602082840312156146fe57600080fd5b815161194a81613c2c565b6000808335601e1984360301811261472057600080fd5b8301803591506001600160401b0382111561473a57600080fd5b6020019150600681901b3603821315613cba57600080fd5b6000808335601e1984360301811261476957600080fd5b8301803591506001600160401b0382111561478357600080fd5b6020019150600581901b3603821315613cba57600080fd5b6000808335601e198436030181126147b257600080fd5b8301803591506001600160401b038211156147cc57600080fd5b602001915036819003821315613cba57600080fd5b80516020808301519190811015613c705760001960209190910360031b1b16919050565b634e487b7160e01b600052601260045260246000fd5b600063ffffffff83168061483157614831614805565b8063ffffffff84160691505092915050565b60008261485257614852614805565b500490565b60008251614212818460208701613ec0565b60008261487857614878614805565b50069056fea2646970667358221220036b4ae61df7766da491998f092bd174758335f223a2cfc0b5ce481d4262e9f164736f6c634300081b0033",
}

// ContractRewardsCoordinatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRewardsCoordinatorMetaData.ABI instead.
var ContractRewardsCoordinatorABI = ContractRewardsCoordinatorMetaData.ABI

// ContractRewardsCoordinatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRewardsCoordinatorMetaData.Bin instead.
var ContractRewardsCoordinatorBin = ContractRewardsCoordinatorMetaData.Bin

// DeployContractRewardsCoordinator deploys a new Ethereum contract, binding an instance of ContractRewardsCoordinator to it.
func DeployContractRewardsCoordinator(auth *bind.TransactOpts, backend bind.ContractBackend, params IRewardsCoordinatorTypesRewardsCoordinatorConstructorParams) (common.Address, *types.Transaction, *ContractRewardsCoordinator, error) {
	parsed, err := ContractRewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRewardsCoordinatorBin), backend, params)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRewardsCoordinator{ContractRewardsCoordinatorCaller: ContractRewardsCoordinatorCaller{contract: contract}, ContractRewardsCoordinatorTransactor: ContractRewardsCoordinatorTransactor{contract: contract}, ContractRewardsCoordinatorFilterer: ContractRewardsCoordinatorFilterer{contract: contract}}, nil
}

// ContractRewardsCoordinatorMethods is an auto generated interface around an Ethereum contract.
type ContractRewardsCoordinatorMethods interface {
	ContractRewardsCoordinatorCalls
	ContractRewardsCoordinatorTransacts
	ContractRewardsCoordinatorFilters
}

// ContractRewardsCoordinatorCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractRewardsCoordinatorCalls interface {
	CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error)

	GENESISREWARDSTIMESTAMP(opts *bind.CallOpts) (uint32, error)

	MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error)

	MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error)

	MAXREWARDSDURATION(opts *bind.CallOpts) (uint32, error)

	ActivationDelay(opts *bind.CallOpts) (uint32, error)

	AllocationManager(opts *bind.CallOpts) (common.Address, error)

	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error)

	CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error)

	CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error)

	ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error)

	CumulativeClaimed(opts *bind.CallOpts, earner common.Address, token common.Address) (*big.Int, error)

	CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error)

	DefaultOperatorSplitBips(opts *bind.CallOpts) (uint16, error)

	DelegationManager(opts *bind.CallOpts) (common.Address, error)

	GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error)

	GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error)

	GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error)

	GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error)

	GetOperatorAVSSplit(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error)

	GetOperatorPISplit(opts *bind.CallOpts, operator common.Address) (uint16, error)

	GetOperatorSetSplit(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (uint16, error)

	GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error)

	IsAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error)

	IsOperatorDirectedAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error)

	IsOperatorDirectedOperatorSetRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error)

	IsRewardsForAllSubmitter(opts *bind.CallOpts, submitter common.Address) (bool, error)

	IsRewardsSubmissionForAllEarnersHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error)

	IsRewardsSubmissionForAllHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PermissionController(opts *bind.CallOpts) (common.Address, error)

	RewardsUpdater(opts *bind.CallOpts) (common.Address, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)

	SubmissionNonce(opts *bind.CallOpts, avs common.Address) (*big.Int, error)

	Version(opts *bind.CallOpts) (string, error)
}

// ContractRewardsCoordinatorTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractRewardsCoordinatorTransacts interface {
	CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error)

	CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error)

	CreateOperatorDirectedOperatorSetRewardsSubmission(opts *bind.TransactOpts, operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error)

	CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error)

	CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error)

	DisableRoot(opts *bind.TransactOpts, rootIndex uint32) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error)

	ProcessClaims(opts *bind.TransactOpts, claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error)

	SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error)

	SetClaimerFor0(opts *bind.TransactOpts, earner common.Address, claimer common.Address) (*types.Transaction, error)

	SetDefaultOperatorSplit(opts *bind.TransactOpts, split uint16) (*types.Transaction, error)

	SetOperatorAVSSplit(opts *bind.TransactOpts, operator common.Address, avs common.Address, split uint16) (*types.Transaction, error)

	SetOperatorPISplit(opts *bind.TransactOpts, operator common.Address, split uint16) (*types.Transaction, error)

	SetOperatorSetSplit(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error)

	SetRewardsForAllSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error)

	SetRewardsUpdater(opts *bind.TransactOpts, _rewardsUpdater common.Address) (*types.Transaction, error)

	SubmitRoot(opts *bind.TransactOpts, root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)
}

// ContractRewardsCoordinatorFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractRewardsCoordinatorFilters interface {
	FilterAVSRewardsSubmissionCreated(opts *bind.FilterOpts, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator, error)
	WatchAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorAVSRewardsSubmissionCreated, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseAVSRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorAVSRewardsSubmissionCreated, error)

	FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractRewardsCoordinatorActivationDelaySetIterator, error)
	WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorActivationDelaySet) (event.Subscription, error)
	ParseActivationDelaySet(log types.Log) (*ContractRewardsCoordinatorActivationDelaySet, error)

	FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*ContractRewardsCoordinatorClaimerForSetIterator, error)
	WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error)
	ParseClaimerForSet(log types.Log) (*ContractRewardsCoordinatorClaimerForSet, error)

	FilterDefaultOperatorSplitBipsSet(opts *bind.FilterOpts) (*ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator, error)
	WatchDefaultOperatorSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDefaultOperatorSplitBipsSet) (event.Subscription, error)
	ParseDefaultOperatorSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorDefaultOperatorSplitBipsSet, error)

	FilterDistributionRootDisabled(opts *bind.FilterOpts, rootIndex []uint32) (*ContractRewardsCoordinatorDistributionRootDisabledIterator, error)
	WatchDistributionRootDisabled(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDistributionRootDisabled, rootIndex []uint32) (event.Subscription, error)
	ParseDistributionRootDisabled(log types.Log) (*ContractRewardsCoordinatorDistributionRootDisabled, error)

	FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (*ContractRewardsCoordinatorDistributionRootSubmittedIterator, error)
	WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (event.Subscription, error)
	ParseDistributionRootSubmitted(log types.Log) (*ContractRewardsCoordinatorDistributionRootSubmitted, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractRewardsCoordinatorInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractRewardsCoordinatorInitialized, error)

	FilterOperatorAVSSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address, avs []common.Address) (*ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator, error)
	WatchOperatorAVSSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorAVSSplitBipsSet, caller []common.Address, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorAVSSplitBipsSet, error)

	FilterOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator, error)
	WatchOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseOperatorDirectedAVSRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, error)

	FilterOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator, error)
	WatchOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseOperatorDirectedOperatorSetRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, error)

	FilterOperatorPISplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractRewardsCoordinatorOperatorPISplitBipsSetIterator, error)
	WatchOperatorPISplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorPISplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error)
	ParseOperatorPISplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorPISplitBipsSet, error)

	FilterOperatorSetSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator, error)
	WatchOperatorSetSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorSetSplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error)
	ParseOperatorSetSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorSetSplitBipsSet, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRewardsCoordinatorOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractRewardsCoordinatorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRewardsCoordinatorPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractRewardsCoordinatorPaused, error)

	FilterRewardsClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*ContractRewardsCoordinatorRewardsClaimedIterator, error)
	WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error)
	ParseRewardsClaimed(log types.Log) (*ContractRewardsCoordinatorRewardsClaimed, error)

	FilterRewardsForAllSubmitterSet(opts *bind.FilterOpts, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (*ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator, error)
	WatchRewardsForAllSubmitterSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsForAllSubmitterSet, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error)
	ParseRewardsForAllSubmitterSet(log types.Log) (*ContractRewardsCoordinatorRewardsForAllSubmitterSet, error)

	FilterRewardsSubmissionForAllCreated(opts *bind.FilterOpts, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator, error)
	WatchRewardsSubmissionForAllCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsSubmissionForAllCreated, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseRewardsSubmissionForAllCreated(log types.Log) (*ContractRewardsCoordinatorRewardsSubmissionForAllCreated, error)

	FilterRewardsSubmissionForAllEarnersCreated(opts *bind.FilterOpts, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator, error)
	WatchRewardsSubmissionForAllEarnersCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error)
	ParseRewardsSubmissionForAllEarnersCreated(log types.Log) (*ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, error)

	FilterRewardsUpdaterSet(opts *bind.FilterOpts, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (*ContractRewardsCoordinatorRewardsUpdaterSetIterator, error)
	WatchRewardsUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsUpdaterSet, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (event.Subscription, error)
	ParseRewardsUpdaterSet(log types.Log) (*ContractRewardsCoordinatorRewardsUpdaterSet, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRewardsCoordinatorUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractRewardsCoordinatorUnpaused, error)
}

// ContractRewardsCoordinator is an auto generated Go binding around an Ethereum contract.
type ContractRewardsCoordinator struct {
	ContractRewardsCoordinatorCaller     // Read-only binding to the contract
	ContractRewardsCoordinatorTransactor // Write-only binding to the contract
	ContractRewardsCoordinatorFilterer   // Log filterer for contract events
}

// ContractRewardsCoordinator implements the ContractRewardsCoordinatorMethods interface.
var _ ContractRewardsCoordinatorMethods = (*ContractRewardsCoordinator)(nil)

// ContractRewardsCoordinatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRewardsCoordinatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRewardsCoordinatorCaller implements the ContractRewardsCoordinatorCalls interface.
var _ ContractRewardsCoordinatorCalls = (*ContractRewardsCoordinatorCaller)(nil)

// ContractRewardsCoordinatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRewardsCoordinatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRewardsCoordinatorTransactor implements the ContractRewardsCoordinatorTransacts interface.
var _ ContractRewardsCoordinatorTransacts = (*ContractRewardsCoordinatorTransactor)(nil)

// ContractRewardsCoordinatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRewardsCoordinatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRewardsCoordinatorFilterer implements the ContractRewardsCoordinatorFilters interface.
var _ ContractRewardsCoordinatorFilters = (*ContractRewardsCoordinatorFilterer)(nil)

// ContractRewardsCoordinatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRewardsCoordinatorSession struct {
	Contract     *ContractRewardsCoordinator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractRewardsCoordinatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRewardsCoordinatorCallerSession struct {
	Contract *ContractRewardsCoordinatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractRewardsCoordinatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRewardsCoordinatorTransactorSession struct {
	Contract     *ContractRewardsCoordinatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractRewardsCoordinatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRewardsCoordinatorRaw struct {
	Contract *ContractRewardsCoordinator // Generic contract binding to access the raw methods on
}

// ContractRewardsCoordinatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRewardsCoordinatorCallerRaw struct {
	Contract *ContractRewardsCoordinatorCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRewardsCoordinatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRewardsCoordinatorTransactorRaw struct {
	Contract *ContractRewardsCoordinatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRewardsCoordinator creates a new instance of ContractRewardsCoordinator, bound to a specific deployed contract.
func NewContractRewardsCoordinator(address common.Address, backend bind.ContractBackend) (*ContractRewardsCoordinator, error) {
	contract, err := bindContractRewardsCoordinator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinator{ContractRewardsCoordinatorCaller: ContractRewardsCoordinatorCaller{contract: contract}, ContractRewardsCoordinatorTransactor: ContractRewardsCoordinatorTransactor{contract: contract}, ContractRewardsCoordinatorFilterer: ContractRewardsCoordinatorFilterer{contract: contract}}, nil
}

// NewContractRewardsCoordinatorCaller creates a new read-only instance of ContractRewardsCoordinator, bound to a specific deployed contract.
func NewContractRewardsCoordinatorCaller(address common.Address, caller bind.ContractCaller) (*ContractRewardsCoordinatorCaller, error) {
	contract, err := bindContractRewardsCoordinator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorCaller{contract: contract}, nil
}

// NewContractRewardsCoordinatorTransactor creates a new write-only instance of ContractRewardsCoordinator, bound to a specific deployed contract.
func NewContractRewardsCoordinatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRewardsCoordinatorTransactor, error) {
	contract, err := bindContractRewardsCoordinator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorTransactor{contract: contract}, nil
}

// NewContractRewardsCoordinatorFilterer creates a new log filterer instance of ContractRewardsCoordinator, bound to a specific deployed contract.
func NewContractRewardsCoordinatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRewardsCoordinatorFilterer, error) {
	contract, err := bindContractRewardsCoordinator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorFilterer{contract: contract}, nil
}

// bindContractRewardsCoordinator binds a generic wrapper to an already deployed contract.
func bindContractRewardsCoordinator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractRewardsCoordinatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRewardsCoordinator.Contract.ContractRewardsCoordinatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ContractRewardsCoordinatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ContractRewardsCoordinatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRewardsCoordinator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.contract.Transact(opts, method, params...)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CALCULATIONINTERVALSECONDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "CALCULATION_INTERVAL_SECONDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_ContractRewardsCoordinator.CallOpts)
}

// CALCULATIONINTERVALSECONDS is a free data retrieval call binding the contract method 0x9d45c281.
//
// Solidity: function CALCULATION_INTERVAL_SECONDS() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CALCULATIONINTERVALSECONDS() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.CALCULATIONINTERVALSECONDS(&_ContractRewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GENESISREWARDSTIMESTAMP(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "GENESIS_REWARDS_TIMESTAMP")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_ContractRewardsCoordinator.CallOpts)
}

// GENESISREWARDSTIMESTAMP is a free data retrieval call binding the contract method 0x131433b4.
//
// Solidity: function GENESIS_REWARDS_TIMESTAMP() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GENESISREWARDSTIMESTAMP() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.GENESISREWARDSTIMESTAMP(&_ContractRewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) MAXFUTURELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "MAX_FUTURE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) MAXFUTURELENGTH() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXFUTURELENGTH(&_ContractRewardsCoordinator.CallOpts)
}

// MAXFUTURELENGTH is a free data retrieval call binding the contract method 0x04a0c502.
//
// Solidity: function MAX_FUTURE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) MAXFUTURELENGTH() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXFUTURELENGTH(&_ContractRewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) MAXRETROACTIVELENGTH(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "MAX_RETROACTIVE_LENGTH")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractRewardsCoordinator.CallOpts)
}

// MAXRETROACTIVELENGTH is a free data retrieval call binding the contract method 0x37838ed0.
//
// Solidity: function MAX_RETROACTIVE_LENGTH() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) MAXRETROACTIVELENGTH() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXRETROACTIVELENGTH(&_ContractRewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) MAXREWARDSDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "MAX_REWARDS_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) MAXREWARDSDURATION() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXREWARDSDURATION(&_ContractRewardsCoordinator.CallOpts)
}

// MAXREWARDSDURATION is a free data retrieval call binding the contract method 0xbf21a8aa.
//
// Solidity: function MAX_REWARDS_DURATION() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) MAXREWARDSDURATION() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.MAXREWARDSDURATION(&_ContractRewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) ActivationDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "activationDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) ActivationDelay() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.ActivationDelay(&_ContractRewardsCoordinator.CallOpts)
}

// ActivationDelay is a free data retrieval call binding the contract method 0x3a8c0786.
//
// Solidity: function activationDelay() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) ActivationDelay() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.ActivationDelay(&_ContractRewardsCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) AllocationManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.AllocationManager(&_ContractRewardsCoordinator.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) AllocationManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.AllocationManager(&_ContractRewardsCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.BeaconChainETHStrategy(&_ContractRewardsCoordinator.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.BeaconChainETHStrategy(&_ContractRewardsCoordinator.CallOpts)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CalculateEarnerLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "calculateEarnerLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractRewardsCoordinator.Contract.CalculateEarnerLeafHash(&_ContractRewardsCoordinator.CallOpts, leaf)
}

// CalculateEarnerLeafHash is a free data retrieval call binding the contract method 0x149bc872.
//
// Solidity: function calculateEarnerLeafHash((address,bytes32) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CalculateEarnerLeafHash(leaf IRewardsCoordinatorTypesEarnerTreeMerkleLeaf) ([32]byte, error) {
	return _ContractRewardsCoordinator.Contract.CalculateEarnerLeafHash(&_ContractRewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CalculateTokenLeafHash(opts *bind.CallOpts, leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "calculateTokenLeafHash", leaf)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractRewardsCoordinator.Contract.CalculateTokenLeafHash(&_ContractRewardsCoordinator.CallOpts, leaf)
}

// CalculateTokenLeafHash is a free data retrieval call binding the contract method 0xf8cd8448.
//
// Solidity: function calculateTokenLeafHash((address,uint256) leaf) pure returns(bytes32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CalculateTokenLeafHash(leaf IRewardsCoordinatorTypesTokenTreeMerkleLeaf) ([32]byte, error) {
	return _ContractRewardsCoordinator.Contract.CalculateTokenLeafHash(&_ContractRewardsCoordinator.CallOpts, leaf)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CheckClaim(opts *bind.CallOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "checkClaim", claim)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CheckClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
	return _ContractRewardsCoordinator.Contract.CheckClaim(&_ContractRewardsCoordinator.CallOpts, claim)
}

// CheckClaim is a free data retrieval call binding the contract method 0x5e9d8348.
//
// Solidity: function checkClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CheckClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim) (bool, error) {
	return _ContractRewardsCoordinator.Contract.CheckClaim(&_ContractRewardsCoordinator.CallOpts, claim)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) ClaimerFor(opts *bind.CallOpts, earner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "claimerFor", earner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.ClaimerFor(&_ContractRewardsCoordinator.CallOpts, earner)
}

// ClaimerFor is a free data retrieval call binding the contract method 0x2b9f64a4.
//
// Solidity: function claimerFor(address earner) view returns(address claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) ClaimerFor(earner common.Address) (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.ClaimerFor(&_ContractRewardsCoordinator.CallOpts, earner)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CumulativeClaimed(opts *bind.CallOpts, earner common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "cumulativeClaimed", earner, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CumulativeClaimed(earner common.Address, token common.Address) (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.CumulativeClaimed(&_ContractRewardsCoordinator.CallOpts, earner, token)
}

// CumulativeClaimed is a free data retrieval call binding the contract method 0x865c6953.
//
// Solidity: function cumulativeClaimed(address earner, address token) view returns(uint256 totalClaimed)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CumulativeClaimed(earner common.Address, token common.Address) (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.CumulativeClaimed(&_ContractRewardsCoordinator.CallOpts, earner, token)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) CurrRewardsCalculationEndTimestamp(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "currRewardsCalculationEndTimestamp")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_ContractRewardsCoordinator.CallOpts)
}

// CurrRewardsCalculationEndTimestamp is a free data retrieval call binding the contract method 0x4d18cc35.
//
// Solidity: function currRewardsCalculationEndTimestamp() view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) CurrRewardsCalculationEndTimestamp() (uint32, error) {
	return _ContractRewardsCoordinator.Contract.CurrRewardsCalculationEndTimestamp(&_ContractRewardsCoordinator.CallOpts)
}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) DefaultOperatorSplitBips(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "defaultOperatorSplitBips")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) DefaultOperatorSplitBips() (uint16, error) {
	return _ContractRewardsCoordinator.Contract.DefaultOperatorSplitBips(&_ContractRewardsCoordinator.CallOpts)
}

// DefaultOperatorSplitBips is a free data retrieval call binding the contract method 0x63f6a798.
//
// Solidity: function defaultOperatorSplitBips() view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) DefaultOperatorSplitBips() (uint16, error) {
	return _ContractRewardsCoordinator.Contract.DefaultOperatorSplitBips(&_ContractRewardsCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) DelegationManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.DelegationManager(&_ContractRewardsCoordinator.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) DelegationManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.DelegationManager(&_ContractRewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetCurrentClaimableDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getCurrentClaimableDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_ContractRewardsCoordinator.CallOpts)
}

// GetCurrentClaimableDistributionRoot is a free data retrieval call binding the contract method 0x0e9a53cf.
//
// Solidity: function getCurrentClaimableDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetCurrentClaimableDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetCurrentClaimableDistributionRoot(&_ContractRewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetCurrentDistributionRoot(opts *bind.CallOpts) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getCurrentDistributionRoot")

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetCurrentDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetCurrentDistributionRoot(&_ContractRewardsCoordinator.CallOpts)
}

// GetCurrentDistributionRoot is a free data retrieval call binding the contract method 0x9be3d4e4.
//
// Solidity: function getCurrentDistributionRoot() view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetCurrentDistributionRoot() (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetCurrentDistributionRoot(&_ContractRewardsCoordinator.CallOpts)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetDistributionRootAtIndex(opts *bind.CallOpts, index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getDistributionRootAtIndex", index)

	if err != nil {
		return *new(IRewardsCoordinatorTypesDistributionRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(IRewardsCoordinatorTypesDistributionRoot)).(*IRewardsCoordinatorTypesDistributionRoot)

	return out0, err

}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetDistributionRootAtIndex(&_ContractRewardsCoordinator.CallOpts, index)
}

// GetDistributionRootAtIndex is a free data retrieval call binding the contract method 0xde02e503.
//
// Solidity: function getDistributionRootAtIndex(uint256 index) view returns((bytes32,uint32,uint32,bool))
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetDistributionRootAtIndex(index *big.Int) (IRewardsCoordinatorTypesDistributionRoot, error) {
	return _ContractRewardsCoordinator.Contract.GetDistributionRootAtIndex(&_ContractRewardsCoordinator.CallOpts, index)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetDistributionRootsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getDistributionRootsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetDistributionRootsLength() (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.GetDistributionRootsLength(&_ContractRewardsCoordinator.CallOpts)
}

// GetDistributionRootsLength is a free data retrieval call binding the contract method 0x7b8f8b05.
//
// Solidity: function getDistributionRootsLength() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetDistributionRootsLength() (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.GetDistributionRootsLength(&_ContractRewardsCoordinator.CallOpts)
}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetOperatorAVSSplit(opts *bind.CallOpts, operator common.Address, avs common.Address) (uint16, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getOperatorAVSSplit", operator, avs)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorAVSSplit(&_ContractRewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorAVSSplit is a free data retrieval call binding the contract method 0xe063f81f.
//
// Solidity: function getOperatorAVSSplit(address operator, address avs) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetOperatorAVSSplit(operator common.Address, avs common.Address) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorAVSSplit(&_ContractRewardsCoordinator.CallOpts, operator, avs)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetOperatorPISplit(opts *bind.CallOpts, operator common.Address) (uint16, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getOperatorPISplit", operator)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorPISplit(&_ContractRewardsCoordinator.CallOpts, operator)
}

// GetOperatorPISplit is a free data retrieval call binding the contract method 0x4b943960.
//
// Solidity: function getOperatorPISplit(address operator) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetOperatorPISplit(operator common.Address) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorPISplit(&_ContractRewardsCoordinator.CallOpts, operator)
}

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetOperatorSetSplit(opts *bind.CallOpts, operator common.Address, operatorSet OperatorSet) (uint16, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getOperatorSetSplit", operator, operatorSet)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetOperatorSetSplit(operator common.Address, operatorSet OperatorSet) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorSetSplit(&_ContractRewardsCoordinator.CallOpts, operator, operatorSet)
}

// GetOperatorSetSplit is a free data retrieval call binding the contract method 0x9de4b35f.
//
// Solidity: function getOperatorSetSplit(address operator, (address,uint32) operatorSet) view returns(uint16)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetOperatorSetSplit(operator common.Address, operatorSet OperatorSet) (uint16, error) {
	return _ContractRewardsCoordinator.Contract.GetOperatorSetSplit(&_ContractRewardsCoordinator.CallOpts, operator, operatorSet)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) GetRootIndexFromHash(opts *bind.CallOpts, rootHash [32]byte) (uint32, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "getRootIndexFromHash", rootHash)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractRewardsCoordinator.Contract.GetRootIndexFromHash(&_ContractRewardsCoordinator.CallOpts, rootHash)
}

// GetRootIndexFromHash is a free data retrieval call binding the contract method 0xe810ce21.
//
// Solidity: function getRootIndexFromHash(bytes32 rootHash) view returns(uint32)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) GetRootIndexFromHash(rootHash [32]byte) (uint32, error) {
	return _ContractRewardsCoordinator.Contract.GetRootIndexFromHash(&_ContractRewardsCoordinator.CallOpts, rootHash)
}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isAVSRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0x6d21117e.
//
// Solidity: function isAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsAVSRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsOperatorDirectedAVSRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isOperatorDirectedAVSRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsOperatorDirectedAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedAVSRewardsSubmissionHash is a free data retrieval call binding the contract method 0xed71e6a2.
//
// Solidity: function isOperatorDirectedAVSRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsOperatorDirectedAVSRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsOperatorDirectedAVSRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsOperatorDirectedOperatorSetRewardsSubmissionHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isOperatorDirectedOperatorSetRewardsSubmissionHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsOperatorDirectedOperatorSetRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsOperatorDirectedOperatorSetRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsOperatorDirectedOperatorSetRewardsSubmissionHash is a free data retrieval call binding the contract method 0xf2f07ab4.
//
// Solidity: function isOperatorDirectedOperatorSetRewardsSubmissionHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsOperatorDirectedOperatorSetRewardsSubmissionHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsOperatorDirectedOperatorSetRewardsSubmissionHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsRewardsForAllSubmitter(opts *bind.CallOpts, submitter common.Address) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isRewardsForAllSubmitter", submitter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsRewardsForAllSubmitter(submitter common.Address) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_ContractRewardsCoordinator.CallOpts, submitter)
}

// IsRewardsForAllSubmitter is a free data retrieval call binding the contract method 0x0018572c.
//
// Solidity: function isRewardsForAllSubmitter(address submitter) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsRewardsForAllSubmitter(submitter common.Address) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsForAllSubmitter(&_ContractRewardsCoordinator.CallOpts, submitter)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsRewardsSubmissionForAllEarnersHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllEarnersHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsRewardsSubmissionForAllEarnersHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllEarnersHash is a free data retrieval call binding the contract method 0xaebd8bae.
//
// Solidity: function isRewardsSubmissionForAllEarnersHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsRewardsSubmissionForAllEarnersHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsSubmissionForAllEarnersHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) IsRewardsSubmissionForAllHash(opts *bind.CallOpts, avs common.Address, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "isRewardsSubmissionForAllHash", avs, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) IsRewardsSubmissionForAllHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// IsRewardsSubmissionForAllHash is a free data retrieval call binding the contract method 0xc46db606.
//
// Solidity: function isRewardsSubmissionForAllHash(address avs, bytes32 hash) view returns(bool valid)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) IsRewardsSubmissionForAllHash(avs common.Address, hash [32]byte) (bool, error) {
	return _ContractRewardsCoordinator.Contract.IsRewardsSubmissionForAllHash(&_ContractRewardsCoordinator.CallOpts, avs, hash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Owner() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.Owner(&_ContractRewardsCoordinator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) Owner() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.Owner(&_ContractRewardsCoordinator.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Paused(index uint8) (bool, error) {
	return _ContractRewardsCoordinator.Contract.Paused(&_ContractRewardsCoordinator.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) Paused(index uint8) (bool, error) {
	return _ContractRewardsCoordinator.Contract.Paused(&_ContractRewardsCoordinator.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Paused0() (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.Paused0(&_ContractRewardsCoordinator.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) Paused0() (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.Paused0(&_ContractRewardsCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) PauserRegistry() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.PauserRegistry(&_ContractRewardsCoordinator.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.PauserRegistry(&_ContractRewardsCoordinator.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) PermissionController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "permissionController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) PermissionController() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.PermissionController(&_ContractRewardsCoordinator.CallOpts)
}

// PermissionController is a free data retrieval call binding the contract method 0x4657e26a.
//
// Solidity: function permissionController() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) PermissionController() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.PermissionController(&_ContractRewardsCoordinator.CallOpts)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) RewardsUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "rewardsUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) RewardsUpdater() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.RewardsUpdater(&_ContractRewardsCoordinator.CallOpts)
}

// RewardsUpdater is a free data retrieval call binding the contract method 0xfbf1e2c1.
//
// Solidity: function rewardsUpdater() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) RewardsUpdater() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.RewardsUpdater(&_ContractRewardsCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) StrategyManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.StrategyManager(&_ContractRewardsCoordinator.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) StrategyManager() (common.Address, error) {
	return _ContractRewardsCoordinator.Contract.StrategyManager(&_ContractRewardsCoordinator.CallOpts)
}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) SubmissionNonce(opts *bind.CallOpts, avs common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "submissionNonce", avs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SubmissionNonce(avs common.Address) (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.SubmissionNonce(&_ContractRewardsCoordinator.CallOpts, avs)
}

// SubmissionNonce is a free data retrieval call binding the contract method 0xbb7e451f.
//
// Solidity: function submissionNonce(address avs) view returns(uint256 nonce)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) SubmissionNonce(avs common.Address) (*big.Int, error) {
	return _ContractRewardsCoordinator.Contract.SubmissionNonce(&_ContractRewardsCoordinator.CallOpts, avs)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractRewardsCoordinator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Version() (string, error) {
	return _ContractRewardsCoordinator.Contract.Version(&_ContractRewardsCoordinator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorCallerSession) Version() (string, error) {
	return _ContractRewardsCoordinator.Contract.Version(&_ContractRewardsCoordinator.CallOpts)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) CreateAVSRewardsSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "createAVSRewardsSubmission", rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateAVSRewardsSubmission is a paid mutator transaction binding the contract method 0xfce36c7d.
//
// Solidity: function createAVSRewardsSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) CreateAVSRewardsSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateAVSRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) CreateOperatorDirectedAVSRewardsSubmission(opts *bind.TransactOpts, avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "createOperatorDirectedAVSRewardsSubmission", avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedAVSRewardsSubmission is a paid mutator transaction binding the contract method 0x9cb9a5fa.
//
// Solidity: function createOperatorDirectedAVSRewardsSubmission(address avs, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) CreateOperatorDirectedAVSRewardsSubmission(avs common.Address, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateOperatorDirectedAVSRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, avs, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) CreateOperatorDirectedOperatorSetRewardsSubmission(opts *bind.TransactOpts, operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "createOperatorDirectedOperatorSetRewardsSubmission", operatorSet, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CreateOperatorDirectedOperatorSetRewardsSubmission(operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateOperatorDirectedOperatorSetRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, operatorSet, operatorDirectedRewardsSubmissions)
}

// CreateOperatorDirectedOperatorSetRewardsSubmission is a paid mutator transaction binding the contract method 0x0ca29899.
//
// Solidity: function createOperatorDirectedOperatorSetRewardsSubmission((address,uint32) operatorSet, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string)[] operatorDirectedRewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) CreateOperatorDirectedOperatorSetRewardsSubmission(operatorSet OperatorSet, operatorDirectedRewardsSubmissions []IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateOperatorDirectedOperatorSetRewardsSubmission(&_ContractRewardsCoordinator.TransactOpts, operatorSet, operatorDirectedRewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) CreateRewardsForAllEarners(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "createRewardsForAllEarners", rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateRewardsForAllEarners(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllEarners is a paid mutator transaction binding the contract method 0xff9f6cce.
//
// Solidity: function createRewardsForAllEarners(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) CreateRewardsForAllEarners(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateRewardsForAllEarners(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) CreateRewardsForAllSubmission(opts *bind.TransactOpts, rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "createRewardsForAllSubmission", rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// CreateRewardsForAllSubmission is a paid mutator transaction binding the contract method 0x36af41fa.
//
// Solidity: function createRewardsForAllSubmission(((address,uint96)[],address,uint256,uint32,uint32)[] rewardsSubmissions) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) CreateRewardsForAllSubmission(rewardsSubmissions []IRewardsCoordinatorTypesRewardsSubmission) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.CreateRewardsForAllSubmission(&_ContractRewardsCoordinator.TransactOpts, rewardsSubmissions)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) DisableRoot(opts *bind.TransactOpts, rootIndex uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "disableRoot", rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.DisableRoot(&_ContractRewardsCoordinator.TransactOpts, rootIndex)
}

// DisableRoot is a paid mutator transaction binding the contract method 0xf96abf2e.
//
// Solidity: function disableRoot(uint32 rootIndex) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) DisableRoot(rootIndex uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.DisableRoot(&_ContractRewardsCoordinator.TransactOpts, rootIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Initialize(&_ContractRewardsCoordinator.TransactOpts, initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
}

// Initialize is a paid mutator transaction binding the contract method 0xf6efbb59.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus, address _rewardsUpdater, uint32 _activationDelay, uint16 _defaultSplitBips) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int, _rewardsUpdater common.Address, _activationDelay uint32, _defaultSplitBips uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Initialize(&_ContractRewardsCoordinator.TransactOpts, initialOwner, initialPausedStatus, _rewardsUpdater, _activationDelay, _defaultSplitBips)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Pause(&_ContractRewardsCoordinator.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Pause(&_ContractRewardsCoordinator.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.PauseAll(&_ContractRewardsCoordinator.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.PauseAll(&_ContractRewardsCoordinator.TransactOpts)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) ProcessClaim(opts *bind.TransactOpts, claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "processClaim", claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) ProcessClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ProcessClaim(&_ContractRewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaim is a paid mutator transaction binding the contract method 0x3ccc861d.
//
// Solidity: function processClaim((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[]) claim, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) ProcessClaim(claim IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ProcessClaim(&_ContractRewardsCoordinator.TransactOpts, claim, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) ProcessClaims(opts *bind.TransactOpts, claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "processClaims", claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) ProcessClaims(claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ProcessClaims(&_ContractRewardsCoordinator.TransactOpts, claims, recipient)
}

// ProcessClaims is a paid mutator transaction binding the contract method 0x4596021c.
//
// Solidity: function processClaims((uint32,uint32,bytes,(address,bytes32),uint32[],bytes[],(address,uint256)[])[] claims, address recipient) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) ProcessClaims(claims []IRewardsCoordinatorTypesRewardsMerkleClaim, recipient common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.ProcessClaims(&_ContractRewardsCoordinator.TransactOpts, claims, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.RenounceOwnership(&_ContractRewardsCoordinator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.RenounceOwnership(&_ContractRewardsCoordinator.TransactOpts)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetActivationDelay(opts *bind.TransactOpts, _activationDelay uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setActivationDelay", _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetActivationDelay(&_ContractRewardsCoordinator.TransactOpts, _activationDelay)
}

// SetActivationDelay is a paid mutator transaction binding the contract method 0x58baaa3e.
//
// Solidity: function setActivationDelay(uint32 _activationDelay) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetActivationDelay(_activationDelay uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetActivationDelay(&_ContractRewardsCoordinator.TransactOpts, _activationDelay)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetClaimerFor(opts *bind.TransactOpts, claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setClaimerFor", claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetClaimerFor(&_ContractRewardsCoordinator.TransactOpts, claimer)
}

// SetClaimerFor is a paid mutator transaction binding the contract method 0xa0169ddd.
//
// Solidity: function setClaimerFor(address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetClaimerFor(claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetClaimerFor(&_ContractRewardsCoordinator.TransactOpts, claimer)
}

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetClaimerFor0(opts *bind.TransactOpts, earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setClaimerFor0", earner, claimer)
}

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetClaimerFor0(earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetClaimerFor0(&_ContractRewardsCoordinator.TransactOpts, earner, claimer)
}

// SetClaimerFor0 is a paid mutator transaction binding the contract method 0xf22cef85.
//
// Solidity: function setClaimerFor(address earner, address claimer) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetClaimerFor0(earner common.Address, claimer common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetClaimerFor0(&_ContractRewardsCoordinator.TransactOpts, earner, claimer)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetDefaultOperatorSplit(opts *bind.TransactOpts, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setDefaultOperatorSplit", split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetDefaultOperatorSplit(&_ContractRewardsCoordinator.TransactOpts, split)
}

// SetDefaultOperatorSplit is a paid mutator transaction binding the contract method 0xa50a1d9c.
//
// Solidity: function setDefaultOperatorSplit(uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetDefaultOperatorSplit(split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetDefaultOperatorSplit(&_ContractRewardsCoordinator.TransactOpts, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetOperatorAVSSplit(opts *bind.TransactOpts, operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setOperatorAVSSplit", operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorAVSSplit(&_ContractRewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorAVSSplit is a paid mutator transaction binding the contract method 0xdcbb03b3.
//
// Solidity: function setOperatorAVSSplit(address operator, address avs, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetOperatorAVSSplit(operator common.Address, avs common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorAVSSplit(&_ContractRewardsCoordinator.TransactOpts, operator, avs, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetOperatorPISplit(opts *bind.TransactOpts, operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setOperatorPISplit", operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorPISplit(&_ContractRewardsCoordinator.TransactOpts, operator, split)
}

// SetOperatorPISplit is a paid mutator transaction binding the contract method 0xb3dbb0e0.
//
// Solidity: function setOperatorPISplit(address operator, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetOperatorPISplit(operator common.Address, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorPISplit(&_ContractRewardsCoordinator.TransactOpts, operator, split)
}

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetOperatorSetSplit(opts *bind.TransactOpts, operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setOperatorSetSplit", operator, operatorSet, split)
}

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetOperatorSetSplit(operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorSetSplit(&_ContractRewardsCoordinator.TransactOpts, operator, operatorSet, split)
}

// SetOperatorSetSplit is a paid mutator transaction binding the contract method 0xf74e8eac.
//
// Solidity: function setOperatorSetSplit(address operator, (address,uint32) operatorSet, uint16 split) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetOperatorSetSplit(operator common.Address, operatorSet OperatorSet, split uint16) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetOperatorSetSplit(&_ContractRewardsCoordinator.TransactOpts, operator, operatorSet, split)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetRewardsForAllSubmitter(opts *bind.TransactOpts, _submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setRewardsForAllSubmitter", _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_ContractRewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsForAllSubmitter is a paid mutator transaction binding the contract method 0x0eb38345.
//
// Solidity: function setRewardsForAllSubmitter(address _submitter, bool _newValue) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetRewardsForAllSubmitter(_submitter common.Address, _newValue bool) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetRewardsForAllSubmitter(&_ContractRewardsCoordinator.TransactOpts, _submitter, _newValue)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SetRewardsUpdater(opts *bind.TransactOpts, _rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "setRewardsUpdater", _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetRewardsUpdater(&_ContractRewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SetRewardsUpdater is a paid mutator transaction binding the contract method 0x863cb9a9.
//
// Solidity: function setRewardsUpdater(address _rewardsUpdater) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SetRewardsUpdater(_rewardsUpdater common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SetRewardsUpdater(&_ContractRewardsCoordinator.TransactOpts, _rewardsUpdater)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) SubmitRoot(opts *bind.TransactOpts, root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "submitRoot", root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SubmitRoot(&_ContractRewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// SubmitRoot is a paid mutator transaction binding the contract method 0x3efe1db6.
//
// Solidity: function submitRoot(bytes32 root, uint32 rewardsCalculationEndTimestamp) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) SubmitRoot(root [32]byte, rewardsCalculationEndTimestamp uint32) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.SubmitRoot(&_ContractRewardsCoordinator.TransactOpts, root, rewardsCalculationEndTimestamp)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.TransferOwnership(&_ContractRewardsCoordinator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.TransferOwnership(&_ContractRewardsCoordinator.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Unpause(&_ContractRewardsCoordinator.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractRewardsCoordinator.Contract.Unpause(&_ContractRewardsCoordinator.TransactOpts, newPausedStatus)
}

// ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator is returned from FilterAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for AVSRewardsSubmissionCreated events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator struct {
	Event *ContractRewardsCoordinatorAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorAVSRewardsSubmissionCreated)
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
		it.Event = new(ContractRewardsCoordinatorAVSRewardsSubmissionCreated)
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
func (it *ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorAVSRewardsSubmissionCreated represents a AVSRewardsSubmissionCreated event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorAVSRewardsSubmissionCreated struct {
	Avs                   common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterAVSRewardsSubmissionCreated(opts *bind.FilterOpts, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorAVSRewardsSubmissionCreatedIterator{contract: _ContractRewardsCoordinator.contract, event: "AVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorAVSRewardsSubmissionCreated, avs []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "AVSRewardsSubmissionCreated", avsRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorAVSRewardsSubmissionCreated)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
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

// ParseAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0x450a367a380c4e339e5ae7340c8464ef27af7781ad9945cfe8abd828f89e6281.
//
// Solidity: event AVSRewardsSubmissionCreated(address indexed avs, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseAVSRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorAVSRewardsSubmissionCreated, error) {
	event := new(ContractRewardsCoordinatorAVSRewardsSubmissionCreated)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "AVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorActivationDelaySetIterator is returned from FilterActivationDelaySet and is used to iterate over the raw logs and unpacked data for ActivationDelaySet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorActivationDelaySetIterator struct {
	Event *ContractRewardsCoordinatorActivationDelaySet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorActivationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorActivationDelaySet)
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
		it.Event = new(ContractRewardsCoordinatorActivationDelaySet)
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
func (it *ContractRewardsCoordinatorActivationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorActivationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorActivationDelaySet represents a ActivationDelaySet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorActivationDelaySet struct {
	OldActivationDelay uint32
	NewActivationDelay uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterActivationDelaySet is a free log retrieval operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterActivationDelaySet(opts *bind.FilterOpts) (*ContractRewardsCoordinatorActivationDelaySetIterator, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorActivationDelaySetIterator{contract: _ContractRewardsCoordinator.contract, event: "ActivationDelaySet", logs: logs, sub: sub}, nil
}

// WatchActivationDelaySet is a free log subscription operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchActivationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorActivationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "ActivationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorActivationDelaySet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
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

// ParseActivationDelaySet is a log parse operation binding the contract event 0xaf557c6c02c208794817a705609cfa935f827312a1adfdd26494b6b95dd2b4b3.
//
// Solidity: event ActivationDelaySet(uint32 oldActivationDelay, uint32 newActivationDelay)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseActivationDelaySet(log types.Log) (*ContractRewardsCoordinatorActivationDelaySet, error) {
	event := new(ContractRewardsCoordinatorActivationDelaySet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "ActivationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorClaimerForSetIterator is returned from FilterClaimerForSet and is used to iterate over the raw logs and unpacked data for ClaimerForSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorClaimerForSetIterator struct {
	Event *ContractRewardsCoordinatorClaimerForSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorClaimerForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorClaimerForSet)
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
		it.Event = new(ContractRewardsCoordinatorClaimerForSet)
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
func (it *ContractRewardsCoordinatorClaimerForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorClaimerForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorClaimerForSet represents a ClaimerForSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorClaimerForSet struct {
	Earner     common.Address
	OldClaimer common.Address
	Claimer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimerForSet is a free log retrieval operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterClaimerForSet(opts *bind.FilterOpts, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (*ContractRewardsCoordinatorClaimerForSetIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorClaimerForSetIterator{contract: _ContractRewardsCoordinator.contract, event: "ClaimerForSet", logs: logs, sub: sub}, nil
}

// WatchClaimerForSet is a free log subscription operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchClaimerForSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorClaimerForSet, earner []common.Address, oldClaimer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var oldClaimerRule []interface{}
	for _, oldClaimerItem := range oldClaimer {
		oldClaimerRule = append(oldClaimerRule, oldClaimerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "ClaimerForSet", earnerRule, oldClaimerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorClaimerForSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
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

// ParseClaimerForSet is a log parse operation binding the contract event 0xbab947934d42e0ad206f25c9cab18b5bb6ae144acfb00f40b4e3aa59590ca312.
//
// Solidity: event ClaimerForSet(address indexed earner, address indexed oldClaimer, address indexed claimer)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseClaimerForSet(log types.Log) (*ContractRewardsCoordinatorClaimerForSet, error) {
	event := new(ContractRewardsCoordinatorClaimerForSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "ClaimerForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator is returned from FilterDefaultOperatorSplitBipsSet and is used to iterate over the raw logs and unpacked data for DefaultOperatorSplitBipsSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator struct {
	Event *ContractRewardsCoordinatorDefaultOperatorSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorDefaultOperatorSplitBipsSet)
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
		it.Event = new(ContractRewardsCoordinatorDefaultOperatorSplitBipsSet)
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
func (it *ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorDefaultOperatorSplitBipsSet represents a DefaultOperatorSplitBipsSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDefaultOperatorSplitBipsSet struct {
	OldDefaultOperatorSplitBips uint16
	NewDefaultOperatorSplitBips uint16
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterDefaultOperatorSplitBipsSet is a free log retrieval operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterDefaultOperatorSplitBipsSet(opts *bind.FilterOpts) (*ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorDefaultOperatorSplitBipsSetIterator{contract: _ContractRewardsCoordinator.contract, event: "DefaultOperatorSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchDefaultOperatorSplitBipsSet is a free log subscription operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchDefaultOperatorSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDefaultOperatorSplitBipsSet) (event.Subscription, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "DefaultOperatorSplitBipsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorDefaultOperatorSplitBipsSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
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

// ParseDefaultOperatorSplitBipsSet is a log parse operation binding the contract event 0xe6cd4edfdcc1f6d130ab35f73d72378f3a642944fb4ee5bd84b7807a81ea1c4e.
//
// Solidity: event DefaultOperatorSplitBipsSet(uint16 oldDefaultOperatorSplitBips, uint16 newDefaultOperatorSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseDefaultOperatorSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorDefaultOperatorSplitBipsSet, error) {
	event := new(ContractRewardsCoordinatorDefaultOperatorSplitBipsSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DefaultOperatorSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorDistributionRootDisabledIterator is returned from FilterDistributionRootDisabled and is used to iterate over the raw logs and unpacked data for DistributionRootDisabled events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDistributionRootDisabledIterator struct {
	Event *ContractRewardsCoordinatorDistributionRootDisabled // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorDistributionRootDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorDistributionRootDisabled)
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
		it.Event = new(ContractRewardsCoordinatorDistributionRootDisabled)
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
func (it *ContractRewardsCoordinatorDistributionRootDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorDistributionRootDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorDistributionRootDisabled represents a DistributionRootDisabled event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDistributionRootDisabled struct {
	RootIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootDisabled is a free log retrieval operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterDistributionRootDisabled(opts *bind.FilterOpts, rootIndex []uint32) (*ContractRewardsCoordinatorDistributionRootDisabledIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorDistributionRootDisabledIterator{contract: _ContractRewardsCoordinator.contract, event: "DistributionRootDisabled", logs: logs, sub: sub}, nil
}

// WatchDistributionRootDisabled is a free log subscription operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchDistributionRootDisabled(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDistributionRootDisabled, rootIndex []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "DistributionRootDisabled", rootIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorDistributionRootDisabled)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
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

// ParseDistributionRootDisabled is a log parse operation binding the contract event 0xd850e6e5dfa497b72661fa73df2923464eaed9dc2ff1d3cb82bccbfeabe5c41e.
//
// Solidity: event DistributionRootDisabled(uint32 indexed rootIndex)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseDistributionRootDisabled(log types.Log) (*ContractRewardsCoordinatorDistributionRootDisabled, error) {
	event := new(ContractRewardsCoordinatorDistributionRootDisabled)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DistributionRootDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorDistributionRootSubmittedIterator is returned from FilterDistributionRootSubmitted and is used to iterate over the raw logs and unpacked data for DistributionRootSubmitted events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDistributionRootSubmittedIterator struct {
	Event *ContractRewardsCoordinatorDistributionRootSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorDistributionRootSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorDistributionRootSubmitted)
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
		it.Event = new(ContractRewardsCoordinatorDistributionRootSubmitted)
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
func (it *ContractRewardsCoordinatorDistributionRootSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorDistributionRootSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorDistributionRootSubmitted represents a DistributionRootSubmitted event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorDistributionRootSubmitted struct {
	RootIndex                      uint32
	Root                           [32]byte
	RewardsCalculationEndTimestamp uint32
	ActivatedAt                    uint32
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterDistributionRootSubmitted is a free log retrieval operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterDistributionRootSubmitted(opts *bind.FilterOpts, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (*ContractRewardsCoordinatorDistributionRootSubmittedIterator, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorDistributionRootSubmittedIterator{contract: _ContractRewardsCoordinator.contract, event: "DistributionRootSubmitted", logs: logs, sub: sub}, nil
}

// WatchDistributionRootSubmitted is a free log subscription operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchDistributionRootSubmitted(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorDistributionRootSubmitted, rootIndex []uint32, root [][32]byte, rewardsCalculationEndTimestamp []uint32) (event.Subscription, error) {

	var rootIndexRule []interface{}
	for _, rootIndexItem := range rootIndex {
		rootIndexRule = append(rootIndexRule, rootIndexItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}
	var rewardsCalculationEndTimestampRule []interface{}
	for _, rewardsCalculationEndTimestampItem := range rewardsCalculationEndTimestamp {
		rewardsCalculationEndTimestampRule = append(rewardsCalculationEndTimestampRule, rewardsCalculationEndTimestampItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "DistributionRootSubmitted", rootIndexRule, rootRule, rewardsCalculationEndTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorDistributionRootSubmitted)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
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

// ParseDistributionRootSubmitted is a log parse operation binding the contract event 0xecd866c3c158fa00bf34d803d5f6023000b57080bcb48af004c2b4b46b3afd08.
//
// Solidity: event DistributionRootSubmitted(uint32 indexed rootIndex, bytes32 indexed root, uint32 indexed rewardsCalculationEndTimestamp, uint32 activatedAt)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseDistributionRootSubmitted(log types.Log) (*ContractRewardsCoordinatorDistributionRootSubmitted, error) {
	event := new(ContractRewardsCoordinatorDistributionRootSubmitted)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "DistributionRootSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorInitializedIterator struct {
	Event *ContractRewardsCoordinatorInitialized // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorInitialized)
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
		it.Event = new(ContractRewardsCoordinatorInitialized)
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
func (it *ContractRewardsCoordinatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorInitialized represents a Initialized event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractRewardsCoordinatorInitializedIterator, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorInitializedIterator{contract: _ContractRewardsCoordinator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorInitialized)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseInitialized(log types.Log) (*ContractRewardsCoordinatorInitialized, error) {
	event := new(ContractRewardsCoordinatorInitialized)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator is returned from FilterOperatorAVSSplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorAVSSplitBipsSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator struct {
	Event *ContractRewardsCoordinatorOperatorAVSSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOperatorAVSSplitBipsSet)
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
		it.Event = new(ContractRewardsCoordinatorOperatorAVSSplitBipsSet)
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
func (it *ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOperatorAVSSplitBipsSet represents a OperatorAVSSplitBipsSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorAVSSplitBipsSet struct {
	Caller                  common.Address
	Operator                common.Address
	Avs                     common.Address
	ActivatedAt             uint32
	OldOperatorAVSSplitBips uint16
	NewOperatorAVSSplitBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSSplitBipsSet is a free log retrieval operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOperatorAVSSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address, avs []common.Address) (*ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOperatorAVSSplitBipsSetIterator{contract: _ContractRewardsCoordinator.contract, event: "OperatorAVSSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSSplitBipsSet is a free log subscription operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOperatorAVSSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorAVSSplitBipsSet, caller []common.Address, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OperatorAVSSplitBipsSet", callerRule, operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOperatorAVSSplitBipsSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
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

// ParseOperatorAVSSplitBipsSet is a log parse operation binding the contract event 0x48e198b6ae357e529204ee53a8e514c470ff77d9cc8e4f7207f8b5d490ae6934.
//
// Solidity: event OperatorAVSSplitBipsSet(address indexed caller, address indexed operator, address indexed avs, uint32 activatedAt, uint16 oldOperatorAVSSplitBips, uint16 newOperatorAVSSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOperatorAVSSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorAVSSplitBipsSet, error) {
	event := new(ContractRewardsCoordinatorOperatorAVSSplitBipsSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorAVSSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator is returned from FilterOperatorDirectedAVSRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for OperatorDirectedAVSRewardsSubmissionCreated events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator struct {
	Event *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
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
		it.Event = new(ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
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
func (it *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated represents a OperatorDirectedAVSRewardsSubmissionCreated event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated struct {
	Caller                                common.Address
	Avs                                   common.Address
	OperatorDirectedRewardsSubmissionHash [32]byte
	SubmissionNonce                       *big.Int
	OperatorDirectedRewardsSubmission     IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterOperatorDirectedAVSRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreatedIterator{contract: _ContractRewardsCoordinator.contract, event: "OperatorDirectedAVSRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorDirectedAVSRewardsSubmissionCreated is a free log subscription operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOperatorDirectedAVSRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, caller []common.Address, avs []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OperatorDirectedAVSRewardsSubmissionCreated", callerRule, avsRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
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

// ParseOperatorDirectedAVSRewardsSubmissionCreated is a log parse operation binding the contract event 0xfc8888bffd711da60bc5092b33f677d81896fe80ecc677b84cfab8184462b6e0.
//
// Solidity: event OperatorDirectedAVSRewardsSubmissionCreated(address indexed caller, address indexed avs, bytes32 indexed operatorDirectedRewardsSubmissionHash, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOperatorDirectedAVSRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated, error) {
	event := new(ContractRewardsCoordinatorOperatorDirectedAVSRewardsSubmissionCreated)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedAVSRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator is returned from FilterOperatorDirectedOperatorSetRewardsSubmissionCreated and is used to iterate over the raw logs and unpacked data for OperatorDirectedOperatorSetRewardsSubmissionCreated events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator struct {
	Event *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
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
		it.Event = new(ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
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
func (it *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated represents a OperatorDirectedOperatorSetRewardsSubmissionCreated event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated struct {
	Caller                                common.Address
	OperatorDirectedRewardsSubmissionHash [32]byte
	OperatorSet                           OperatorSet
	SubmissionNonce                       *big.Int
	OperatorDirectedRewardsSubmission     IRewardsCoordinatorTypesOperatorDirectedRewardsSubmission
	Raw                                   types.Log // Blockchain specific contextual infos
}

// FilterOperatorDirectedOperatorSetRewardsSubmissionCreated is a free log retrieval operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.FilterOpts, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OperatorDirectedOperatorSetRewardsSubmissionCreated", callerRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreatedIterator{contract: _ContractRewardsCoordinator.contract, event: "OperatorDirectedOperatorSetRewardsSubmissionCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorDirectedOperatorSetRewardsSubmissionCreated is a free log subscription operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOperatorDirectedOperatorSetRewardsSubmissionCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, caller []common.Address, operatorDirectedRewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorDirectedRewardsSubmissionHashRule []interface{}
	for _, operatorDirectedRewardsSubmissionHashItem := range operatorDirectedRewardsSubmissionHash {
		operatorDirectedRewardsSubmissionHashRule = append(operatorDirectedRewardsSubmissionHashRule, operatorDirectedRewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OperatorDirectedOperatorSetRewardsSubmissionCreated", callerRule, operatorDirectedRewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedOperatorSetRewardsSubmissionCreated", log); err != nil {
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

// ParseOperatorDirectedOperatorSetRewardsSubmissionCreated is a log parse operation binding the contract event 0xfff0759ccb371dfb5691798724e70b4fa61cb3bfe730a33ac19fb86a48efc756.
//
// Solidity: event OperatorDirectedOperatorSetRewardsSubmissionCreated(address indexed caller, bytes32 indexed operatorDirectedRewardsSubmissionHash, (address,uint32) operatorSet, uint256 submissionNonce, ((address,uint96)[],address,(address,uint256)[],uint32,uint32,string) operatorDirectedRewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOperatorDirectedOperatorSetRewardsSubmissionCreated(log types.Log) (*ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated, error) {
	event := new(ContractRewardsCoordinatorOperatorDirectedOperatorSetRewardsSubmissionCreated)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorDirectedOperatorSetRewardsSubmissionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOperatorPISplitBipsSetIterator is returned from FilterOperatorPISplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorPISplitBipsSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorPISplitBipsSetIterator struct {
	Event *ContractRewardsCoordinatorOperatorPISplitBipsSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOperatorPISplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOperatorPISplitBipsSet)
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
		it.Event = new(ContractRewardsCoordinatorOperatorPISplitBipsSet)
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
func (it *ContractRewardsCoordinatorOperatorPISplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOperatorPISplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOperatorPISplitBipsSet represents a OperatorPISplitBipsSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorPISplitBipsSet struct {
	Caller                 common.Address
	Operator               common.Address
	ActivatedAt            uint32
	OldOperatorPISplitBips uint16
	NewOperatorPISplitBips uint16
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterOperatorPISplitBipsSet is a free log retrieval operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOperatorPISplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractRewardsCoordinatorOperatorPISplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOperatorPISplitBipsSetIterator{contract: _ContractRewardsCoordinator.contract, event: "OperatorPISplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorPISplitBipsSet is a free log subscription operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOperatorPISplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorPISplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OperatorPISplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOperatorPISplitBipsSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
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

// ParseOperatorPISplitBipsSet is a log parse operation binding the contract event 0xd1e028bd664486a46ad26040e999cd2d22e1e9a094ee6afe19fcf64678f16f74.
//
// Solidity: event OperatorPISplitBipsSet(address indexed caller, address indexed operator, uint32 activatedAt, uint16 oldOperatorPISplitBips, uint16 newOperatorPISplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOperatorPISplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorPISplitBipsSet, error) {
	event := new(ContractRewardsCoordinatorOperatorPISplitBipsSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorPISplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator is returned from FilterOperatorSetSplitBipsSet and is used to iterate over the raw logs and unpacked data for OperatorSetSplitBipsSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator struct {
	Event *ContractRewardsCoordinatorOperatorSetSplitBipsSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOperatorSetSplitBipsSet)
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
		it.Event = new(ContractRewardsCoordinatorOperatorSetSplitBipsSet)
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
func (it *ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOperatorSetSplitBipsSet represents a OperatorSetSplitBipsSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOperatorSetSplitBipsSet struct {
	Caller                  common.Address
	Operator                common.Address
	OperatorSet             OperatorSet
	ActivatedAt             uint32
	OldOperatorSetSplitBips uint16
	NewOperatorSetSplitBips uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetSplitBipsSet is a free log retrieval operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOperatorSetSplitBipsSet(opts *bind.FilterOpts, caller []common.Address, operator []common.Address) (*ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OperatorSetSplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOperatorSetSplitBipsSetIterator{contract: _ContractRewardsCoordinator.contract, event: "OperatorSetSplitBipsSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSetSplitBipsSet is a free log subscription operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOperatorSetSplitBipsSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOperatorSetSplitBipsSet, caller []common.Address, operator []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OperatorSetSplitBipsSet", callerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOperatorSetSplitBipsSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorSetSplitBipsSet", log); err != nil {
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

// ParseOperatorSetSplitBipsSet is a log parse operation binding the contract event 0x14918b3834ab6752eb2e1b489b6663a67810efb5f56f3944a97ede8ecf1fd9f1.
//
// Solidity: event OperatorSetSplitBipsSet(address indexed caller, address indexed operator, (address,uint32) operatorSet, uint32 activatedAt, uint16 oldOperatorSetSplitBips, uint16 newOperatorSetSplitBips)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOperatorSetSplitBipsSet(log types.Log) (*ContractRewardsCoordinatorOperatorSetSplitBipsSet, error) {
	event := new(ContractRewardsCoordinatorOperatorSetSplitBipsSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OperatorSetSplitBipsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOwnershipTransferredIterator struct {
	Event *ContractRewardsCoordinatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorOwnershipTransferred)
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
		it.Event = new(ContractRewardsCoordinatorOwnershipTransferred)
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
func (it *ContractRewardsCoordinatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorOwnershipTransferred represents a OwnershipTransferred event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractRewardsCoordinatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorOwnershipTransferredIterator{contract: _ContractRewardsCoordinator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorOwnershipTransferred)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseOwnershipTransferred(log types.Log) (*ContractRewardsCoordinatorOwnershipTransferred, error) {
	event := new(ContractRewardsCoordinatorOwnershipTransferred)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorPausedIterator struct {
	Event *ContractRewardsCoordinatorPaused // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorPaused)
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
		it.Event = new(ContractRewardsCoordinatorPaused)
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
func (it *ContractRewardsCoordinatorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorPaused represents a Paused event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractRewardsCoordinatorPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorPausedIterator{contract: _ContractRewardsCoordinator.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorPaused)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParsePaused(log types.Log) (*ContractRewardsCoordinatorPaused, error) {
	event := new(ContractRewardsCoordinatorPaused)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsClaimedIterator struct {
	Event *ContractRewardsCoordinatorRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorRewardsClaimed)
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
		it.Event = new(ContractRewardsCoordinatorRewardsClaimed)
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
func (it *ContractRewardsCoordinatorRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorRewardsClaimed represents a RewardsClaimed event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsClaimed struct {
	Root          [32]byte
	Earner        common.Address
	Claimer       common.Address
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, earner []common.Address, claimer []common.Address, recipient []common.Address) (*ContractRewardsCoordinatorRewardsClaimedIterator, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorRewardsClaimedIterator{contract: _ContractRewardsCoordinator.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsClaimed, earner []common.Address, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var earnerRule []interface{}
	for _, earnerItem := range earner {
		earnerRule = append(earnerRule, earnerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "RewardsClaimed", earnerRule, claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorRewardsClaimed)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9543dbd55580842586a951f0386e24d68a5df99ae29e3b216588b45fd684ce31.
//
// Solidity: event RewardsClaimed(bytes32 root, address indexed earner, address indexed claimer, address indexed recipient, address token, uint256 claimedAmount)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseRewardsClaimed(log types.Log) (*ContractRewardsCoordinatorRewardsClaimed, error) {
	event := new(ContractRewardsCoordinatorRewardsClaimed)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator is returned from FilterRewardsForAllSubmitterSet and is used to iterate over the raw logs and unpacked data for RewardsForAllSubmitterSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator struct {
	Event *ContractRewardsCoordinatorRewardsForAllSubmitterSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorRewardsForAllSubmitterSet)
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
		it.Event = new(ContractRewardsCoordinatorRewardsForAllSubmitterSet)
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
func (it *ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorRewardsForAllSubmitterSet represents a RewardsForAllSubmitterSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsForAllSubmitterSet struct {
	RewardsForAllSubmitter common.Address
	OldValue               bool
	NewValue               bool
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterRewardsForAllSubmitterSet is a free log retrieval operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterRewardsForAllSubmitterSet(opts *bind.FilterOpts, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (*ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorRewardsForAllSubmitterSetIterator{contract: _ContractRewardsCoordinator.contract, event: "RewardsForAllSubmitterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsForAllSubmitterSet is a free log subscription operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchRewardsForAllSubmitterSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsForAllSubmitterSet, rewardsForAllSubmitter []common.Address, oldValue []bool, newValue []bool) (event.Subscription, error) {

	var rewardsForAllSubmitterRule []interface{}
	for _, rewardsForAllSubmitterItem := range rewardsForAllSubmitter {
		rewardsForAllSubmitterRule = append(rewardsForAllSubmitterRule, rewardsForAllSubmitterItem)
	}
	var oldValueRule []interface{}
	for _, oldValueItem := range oldValue {
		oldValueRule = append(oldValueRule, oldValueItem)
	}
	var newValueRule []interface{}
	for _, newValueItem := range newValue {
		newValueRule = append(newValueRule, newValueItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "RewardsForAllSubmitterSet", rewardsForAllSubmitterRule, oldValueRule, newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorRewardsForAllSubmitterSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
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

// ParseRewardsForAllSubmitterSet is a log parse operation binding the contract event 0x4de6293e668df1398422e1def12118052c1539a03cbfedc145895d48d7685f1c.
//
// Solidity: event RewardsForAllSubmitterSet(address indexed rewardsForAllSubmitter, bool indexed oldValue, bool indexed newValue)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseRewardsForAllSubmitterSet(log types.Log) (*ContractRewardsCoordinatorRewardsForAllSubmitterSet, error) {
	event := new(ContractRewardsCoordinatorRewardsForAllSubmitterSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsForAllSubmitterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator is returned from FilterRewardsSubmissionForAllCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllCreated events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator struct {
	Event *ContractRewardsCoordinatorRewardsSubmissionForAllCreated // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorRewardsSubmissionForAllCreated)
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
		it.Event = new(ContractRewardsCoordinatorRewardsSubmissionForAllCreated)
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
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorRewardsSubmissionForAllCreated represents a RewardsSubmissionForAllCreated event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsSubmissionForAllCreated struct {
	Submitter             common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllCreated is a free log retrieval operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterRewardsSubmissionForAllCreated(opts *bind.FilterOpts, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorRewardsSubmissionForAllCreatedIterator{contract: _ContractRewardsCoordinator.contract, event: "RewardsSubmissionForAllCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllCreated is a free log subscription operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchRewardsSubmissionForAllCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsSubmissionForAllCreated, submitter []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllCreated", submitterRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorRewardsSubmissionForAllCreated)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
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

// ParseRewardsSubmissionForAllCreated is a log parse operation binding the contract event 0x51088b8c89628df3a8174002c2a034d0152fce6af8415d651b2a4734bf270482.
//
// Solidity: event RewardsSubmissionForAllCreated(address indexed submitter, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseRewardsSubmissionForAllCreated(log types.Log) (*ContractRewardsCoordinatorRewardsSubmissionForAllCreated, error) {
	event := new(ContractRewardsCoordinatorRewardsSubmissionForAllCreated)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator is returned from FilterRewardsSubmissionForAllEarnersCreated and is used to iterate over the raw logs and unpacked data for RewardsSubmissionForAllEarnersCreated events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator struct {
	Event *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
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
		it.Event = new(ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
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
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated represents a RewardsSubmissionForAllEarnersCreated event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated struct {
	TokenHopper           common.Address
	SubmissionNonce       *big.Int
	RewardsSubmissionHash [32]byte
	RewardsSubmission     IRewardsCoordinatorTypesRewardsSubmission
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSubmissionForAllEarnersCreated is a free log retrieval operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterRewardsSubmissionForAllEarnersCreated(opts *bind.FilterOpts, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (*ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreatedIterator{contract: _ContractRewardsCoordinator.contract, event: "RewardsSubmissionForAllEarnersCreated", logs: logs, sub: sub}, nil
}

// WatchRewardsSubmissionForAllEarnersCreated is a free log subscription operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchRewardsSubmissionForAllEarnersCreated(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, tokenHopper []common.Address, submissionNonce []*big.Int, rewardsSubmissionHash [][32]byte) (event.Subscription, error) {

	var tokenHopperRule []interface{}
	for _, tokenHopperItem := range tokenHopper {
		tokenHopperRule = append(tokenHopperRule, tokenHopperItem)
	}
	var submissionNonceRule []interface{}
	for _, submissionNonceItem := range submissionNonce {
		submissionNonceRule = append(submissionNonceRule, submissionNonceItem)
	}
	var rewardsSubmissionHashRule []interface{}
	for _, rewardsSubmissionHashItem := range rewardsSubmissionHash {
		rewardsSubmissionHashRule = append(rewardsSubmissionHashRule, rewardsSubmissionHashItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "RewardsSubmissionForAllEarnersCreated", tokenHopperRule, submissionNonceRule, rewardsSubmissionHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
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

// ParseRewardsSubmissionForAllEarnersCreated is a log parse operation binding the contract event 0x5251b6fdefcb5d81144e735f69ea4c695fd43b0289ca53dc075033f5fc80068b.
//
// Solidity: event RewardsSubmissionForAllEarnersCreated(address indexed tokenHopper, uint256 indexed submissionNonce, bytes32 indexed rewardsSubmissionHash, ((address,uint96)[],address,uint256,uint32,uint32) rewardsSubmission)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseRewardsSubmissionForAllEarnersCreated(log types.Log) (*ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated, error) {
	event := new(ContractRewardsCoordinatorRewardsSubmissionForAllEarnersCreated)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsSubmissionForAllEarnersCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorRewardsUpdaterSetIterator is returned from FilterRewardsUpdaterSet and is used to iterate over the raw logs and unpacked data for RewardsUpdaterSet events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsUpdaterSetIterator struct {
	Event *ContractRewardsCoordinatorRewardsUpdaterSet // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorRewardsUpdaterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorRewardsUpdaterSet)
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
		it.Event = new(ContractRewardsCoordinatorRewardsUpdaterSet)
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
func (it *ContractRewardsCoordinatorRewardsUpdaterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorRewardsUpdaterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorRewardsUpdaterSet represents a RewardsUpdaterSet event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorRewardsUpdaterSet struct {
	OldRewardsUpdater common.Address
	NewRewardsUpdater common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsUpdaterSet is a free log retrieval operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterRewardsUpdaterSet(opts *bind.FilterOpts, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (*ContractRewardsCoordinatorRewardsUpdaterSetIterator, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorRewardsUpdaterSetIterator{contract: _ContractRewardsCoordinator.contract, event: "RewardsUpdaterSet", logs: logs, sub: sub}, nil
}

// WatchRewardsUpdaterSet is a free log subscription operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchRewardsUpdaterSet(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorRewardsUpdaterSet, oldRewardsUpdater []common.Address, newRewardsUpdater []common.Address) (event.Subscription, error) {

	var oldRewardsUpdaterRule []interface{}
	for _, oldRewardsUpdaterItem := range oldRewardsUpdater {
		oldRewardsUpdaterRule = append(oldRewardsUpdaterRule, oldRewardsUpdaterItem)
	}
	var newRewardsUpdaterRule []interface{}
	for _, newRewardsUpdaterItem := range newRewardsUpdater {
		newRewardsUpdaterRule = append(newRewardsUpdaterRule, newRewardsUpdaterItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "RewardsUpdaterSet", oldRewardsUpdaterRule, newRewardsUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorRewardsUpdaterSet)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
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

// ParseRewardsUpdaterSet is a log parse operation binding the contract event 0x237b82f438d75fc568ebab484b75b01d9287b9e98b490b7c23221623b6705dbb.
//
// Solidity: event RewardsUpdaterSet(address indexed oldRewardsUpdater, address indexed newRewardsUpdater)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseRewardsUpdaterSet(log types.Log) (*ContractRewardsCoordinatorRewardsUpdaterSet, error) {
	event := new(ContractRewardsCoordinatorRewardsUpdaterSet)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "RewardsUpdaterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsCoordinatorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorUnpausedIterator struct {
	Event *ContractRewardsCoordinatorUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractRewardsCoordinatorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsCoordinatorUnpaused)
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
		it.Event = new(ContractRewardsCoordinatorUnpaused)
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
func (it *ContractRewardsCoordinatorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsCoordinatorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsCoordinatorUnpaused represents a Unpaused event raised by the ContractRewardsCoordinator contract.
type ContractRewardsCoordinatorUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractRewardsCoordinatorUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsCoordinatorUnpausedIterator{contract: _ContractRewardsCoordinator.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractRewardsCoordinatorUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractRewardsCoordinator.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsCoordinatorUnpaused)
				if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractRewardsCoordinator *ContractRewardsCoordinatorFilterer) ParseUnpaused(log types.Log) (*ContractRewardsCoordinatorUnpaused, error) {
	event := new(ContractRewardsCoordinatorUnpaused)
	if err := _ContractRewardsCoordinator.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
