// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractAVSDirectory

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

// ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractAVSDirectoryMetaData contains all meta data concerning the ContractAVSDirectory contract.
var ContractAVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegation\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPERATOR_SET_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avsOperatorStatus\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isSpent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtilsMixinTypes.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectoryTypes.OperatorAVSRegistrationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToAVS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegisteredToEigenLayer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x60e060405234801561001057600080fd5b506040516119ea3803806119ea83398101604081905261002f916101df565b808084846001600160a01b03811661005a576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b039081166080521660a0526100758161008b565b60c0525061008390506100d2565b505050610314565b600080829050601f815111156100bf578260405163305a27a960e01b81526004016100b691906102ba565b60405180910390fd5b80516100ca826102ed565b179392505050565b600054610100900460ff161561013a5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b60648201526084016100b6565b60005460ff9081161461018b576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146101a257600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b60005b838110156101d65781810151838201526020016101be565b50506000910152565b6000806000606084860312156101f457600080fd5b83516101ff8161018d565b60208501519093506102108161018d565b60408501519092506001600160401b0381111561022c57600080fd5b8401601f8101861361023d57600080fd5b80516001600160401b03811115610256576102566101a5565b604051601f8201601f19908116603f011681016001600160401b0381118282101715610284576102846101a5565b60405281815282820160200188101561029c57600080fd5b6102ad8260208301602086016101bb565b8093505050509250925092565b60208152600082518060208401526102d98160408501602087016101bb565b601f01601f19169190910160400192915050565b8051602080830151919081101561030e576000198160200360031b1b821691505b50919050565b60805160a05160c051611684610366600039600081816104ca0152610ec1015260008181610368015261069f01526000818161024401528181610401015281816105080152610be201526116846000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063a1060c88116100c3578063dce974b91161007c578063dce974b91461033c578063df5cf72314610363578063ec76f4421461038a578063f2fde38b146103be578063f698da25146103d1578063fabc1cbc146103d957600080fd5b8063a1060c88146102a2578063a364f4da146102b5578063a98fb355146102c8578063c825fe68146102db578063cd6dc68714610302578063d79aceab1461031557600080fd5b80635ac86ab7116101155780635ac86ab7146102025780635c975abb14610225578063715018a614610237578063886f11951461023f5780638da5cb5b1461027e5780639926ee7d1461028f57600080fd5b8063136439dd14610152578063374823b51461016757806349075da3146101aa57806354fd4d50146101e5578063595c6a67146101fa575b600080fd5b6101656101603660046111da565b6103ec565b005b610195610175366004611208565b609960209081526000928352604080842090915290825290205460ff1681565b60405190151581526020015b60405180910390f35b6101d86101b8366004611234565b609860209081526000928352604080842090915290825290205460ff1681565b6040516101a19190611283565b6101ed6104c3565b6040516101a191906112fb565b6101656104f3565b610195610210366004611315565b606654600160ff9092169190911b9081161490565b6066545b6040519081526020016101a1565b6101656105a5565b6102667f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016101a1565b6033546001600160a01b0316610266565b61016561029d3660046113a8565b6105b7565b6102296102b036600461149d565b6107dc565b6101656102c33660046114e3565b61085c565b6101656102d6366004611500565b610944565b6102297f809c5ac049c45b7a7f050a20f00c16cf63797efbf8b1eb8d749fdfa39ff8f92981565b610165610310366004611208565b61098b565b6102297fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd81565b6102297f4ee65f64218c67b68da66fd0db16560040a6b973290b9e71912d661ee53fe49581565b6102667f000000000000000000000000000000000000000000000000000000000000000081565b6101656103983660046111da565b33600090815260996020908152604080832093835292905220805460ff19166001179055565b6101656103cc3660046114e3565b610aad565b610229610b26565b6101656103e73660046111da565b610be0565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610450573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104749190611574565b61049157604051631d77d47760e21b815260040160405180910390fd5b60665481811681146104b65760405163c61dca5d60e01b815260040160405180910390fd5b6104bf82610cf1565b5050565b60606104ee7f0000000000000000000000000000000000000000000000000000000000000000610d2e565b905090565b60405163237dfb4760e11b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906346fbf68e90602401602060405180830381865afa158015610557573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057b9190611574565b61059857604051631d77d47760e21b815260040160405180910390fd5b6105a3600019610cf1565b565b6105ad610d6d565b6105a36000610dc7565b6066546000906001908116036105e05760405163840a48d560e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038816845290915290205460ff16600181111561061a5761061a61126d565b0361063857604051631aa528bb60e11b815260040160405180910390fd5b6001600160a01b038316600090815260996020908152604080832085830151845290915290205460ff161561068057604051630d4c4c9160e21b815260040160405180910390fd5b6040516336b87bd760e11b81526001600160a01b0384811660048301527f00000000000000000000000000000000000000000000000000000000000000001690636d70f7ae90602401602060405180830381865afa1580156106e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070a9190611574565b61072757604051639f88c8af60e01b815260040160405180910390fd5b61074b8361073f8533866020015187604001516107dc565b84516040860151610e19565b6001600160a01b038316600081815260996020908152604080832086830151845282528083208054600160ff19918216811790925533808652609885528386208787529094529382902080549094168117909355519092917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41916107cf9190611283565b60405180910390a3505050565b604080517fda2c89bafdd34776a2b8bb9c83c82f419e20cc8c67207f70edd58249b92661bd60208201526001600160a01b038087169282019290925290841660608201526080810183905260a081018290526000906108539060c00160405160208183030381529060405280519060200120610e71565b95945050505050565b6066546000906001908116036108855760405163840a48d560e01b815260040160405180910390fd5b60013360009081526098602090815260408083206001600160a01b038716845290915290205460ff1660018111156108bf576108bf61126d565b146108dd576040516352df45c960e01b815260040160405180910390fd5b3360008181526098602090815260408083206001600160a01b0387168085529252808320805460ff191690555190917ff0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41916109389190611283565b60405180910390a35050565b336001600160a01b03167fa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713838360405161097f929190611596565b60405180910390a25050565b600054610100900460ff16158080156109ab5750600054600160ff909116105b806109c55750303b1580156109c5575060005460ff166001145b610a2d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610a50576000805461ff0019166101001790555b610a5982610cf1565b610a6283610dc7565b8015610aa8576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610ab5610d6d565b6001600160a01b038116610b1a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610a24565b610b2381610dc7565b50565b60408051808201909152600a81526922b4b3b2b72630bcb2b960b11b60209091015260007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f7f71b625cfad44bac63b13dba07f2e1d6084ee04b6f8752101ece6126d584ee6ea610b94610eb8565b805160209182012060408051928301949094529281019190915260608101919091524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c3e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c6291906115c5565b6001600160a01b0316336001600160a01b031614610c935760405163794821ff60e01b815260040160405180910390fd5b60665480198219811614610cba5760405163c61dca5d60e01b815260040160405180910390fd5b606682905560405182815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c9060200161097f565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a250565b60606000610d3b83610f56565b604080516020808252818301909252919250600091906020820181803683375050509182525060208101929092525090565b6033546001600160a01b031633146105a35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a24565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b42811015610e3a57604051630819bdcd60e01b815260040160405180910390fd5b610e4e6001600160a01b0385168484610f84565b610e6b57604051638baa579f60e01b815260040160405180910390fd5b50505050565b6000610e7b610b26565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60606000610ee57f0000000000000000000000000000000000000000000000000000000000000000610d2e565b905080600081518110610efa57610efa6115e2565b602001015160f81c60f81b81600181518110610f1857610f186115e2565b016020908101516040516001600160f81b03199384169281019290925291909116602182015260220160405160208183030381529060405291505090565b600060ff8216601f811115610f7e57604051632cd44ac360e21b815260040160405180910390fd5b92915050565b6000806000610f938585610fe5565b90925090506000816004811115610fac57610fac61126d565b148015610fca5750856001600160a01b0316826001600160a01b0316145b80610fdb5750610fdb86868661102a565b9695505050505050565b600080825160410361101b5760208301516040840151606085015160001a61100f87828585611116565b94509450505050611023565b506000905060025b9250929050565b6000806000856001600160a01b0316631626ba7e60e01b86866040516024016110549291906115f8565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516110929190611619565b600060405180830381855afa9150503d80600081146110cd576040519150601f19603f3d011682016040523d82523d6000602084013e6110d2565b606091505b50915091508180156110e657506020815110155b8015610fdb57508051630b135d3f60e11b9061110b9083016020908101908401611635565b149695505050505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561114d57506000905060036111d1565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156111a1573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111ca576000600192509250506111d1565b9150600090505b94509492505050565b6000602082840312156111ec57600080fd5b5035919050565b6001600160a01b0381168114610b2357600080fd5b6000806040838503121561121b57600080fd5b8235611226816111f3565b946020939093013593505050565b6000806040838503121561124757600080fd5b8235611252816111f3565b91506020830135611262816111f3565b809150509250929050565b634e487b7160e01b600052602160045260246000fd5b60208101600283106112a557634e487b7160e01b600052602160045260246000fd5b91905290565b60005b838110156112c65781810151838201526020016112ae565b50506000910152565b600081518084526112e78160208601602086016112ab565b601f01601f19169290920160200192915050565b60208152600061130e60208301846112cf565b9392505050565b60006020828403121561132757600080fd5b813560ff8116811461130e57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561137157611371611338565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156113a0576113a0611338565b604052919050565b600080604083850312156113bb57600080fd5b82356113c6816111f3565b9150602083013567ffffffffffffffff8111156113e257600080fd5b8301606081860312156113f457600080fd5b6113fc61134e565b813567ffffffffffffffff81111561141357600080fd5b8201601f8101871361142457600080fd5b803567ffffffffffffffff81111561143e5761143e611338565b611451601f8201601f1916602001611377565b81815288602083850101111561146657600080fd5b8160208401602083013760006020928201830152835283810135908301525060409182013591810191909152919491935090915050565b600080600080608085870312156114b357600080fd5b84356114be816111f3565b935060208501356114ce816111f3565b93969395505050506040820135916060013590565b6000602082840312156114f557600080fd5b813561130e816111f3565b6000806020838503121561151357600080fd5b823567ffffffffffffffff81111561152a57600080fd5b8301601f8101851361153b57600080fd5b803567ffffffffffffffff81111561155257600080fd5b85602082840101111561156457600080fd5b6020919091019590945092505050565b60006020828403121561158657600080fd5b8151801515811461130e57600080fd5b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000602082840312156115d757600080fd5b815161130e816111f3565b634e487b7160e01b600052603260045260246000fd5b82815260406020820152600061161160408301846112cf565b949350505050565b6000825161162b8184602087016112ab565b9190910192915050565b60006020828403121561164757600080fd5b505191905056fea26469706673582212202e9d4d1ed4c366b7f53b5a680d5cce0a49dff056c2be5de36bf0a36dd6f43e3264736f6c634300081b0033",
}

// ContractAVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAVSDirectoryMetaData.ABI instead.
var ContractAVSDirectoryABI = ContractAVSDirectoryMetaData.ABI

// ContractAVSDirectoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAVSDirectoryMetaData.Bin instead.
var ContractAVSDirectoryBin = ContractAVSDirectoryMetaData.Bin

// DeployContractAVSDirectory deploys a new Ethereum contract, binding an instance of ContractAVSDirectory to it.
func DeployContractAVSDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, _delegation common.Address, _pauserRegistry common.Address, _version string) (common.Address, *types.Transaction, *ContractAVSDirectory, error) {
	parsed, err := ContractAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAVSDirectoryBin), backend, _delegation, _pauserRegistry, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAVSDirectory{ContractAVSDirectoryCaller: ContractAVSDirectoryCaller{contract: contract}, ContractAVSDirectoryTransactor: ContractAVSDirectoryTransactor{contract: contract}, ContractAVSDirectoryFilterer: ContractAVSDirectoryFilterer{contract: contract}}, nil
}

// ContractAVSDirectoryMethods is an auto generated interface around an Ethereum contract.
type ContractAVSDirectoryMethods interface {
	ContractAVSDirectoryCalls
	ContractAVSDirectoryTransacts
	ContractAVSDirectoryFilters
}

// ContractAVSDirectoryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractAVSDirectoryCalls interface {
	OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	AvsOperatorStatus(opts *bind.CallOpts, avs common.Address, operator common.Address) (uint8, error)

	CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	Version(opts *bind.CallOpts) (string, error)
}

// ContractAVSDirectoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractAVSDirectoryTransacts interface {
	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractAVSDirectoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractAVSDirectoryFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractAVSDirectoryAVSMetadataURIUpdated, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractAVSDirectoryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractAVSDirectoryInitialized, error)

	FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error)
	WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSDirectoryOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractAVSDirectoryOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractAVSDirectoryPaused, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractAVSDirectoryUnpaused, error)
}

// ContractAVSDirectory is an auto generated Go binding around an Ethereum contract.
type ContractAVSDirectory struct {
	ContractAVSDirectoryCaller     // Read-only binding to the contract
	ContractAVSDirectoryTransactor // Write-only binding to the contract
	ContractAVSDirectoryFilterer   // Log filterer for contract events
}

// ContractAVSDirectory implements the ContractAVSDirectoryMethods interface.
var _ ContractAVSDirectoryMethods = (*ContractAVSDirectory)(nil)

// ContractAVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryCaller implements the ContractAVSDirectoryCalls interface.
var _ ContractAVSDirectoryCalls = (*ContractAVSDirectoryCaller)(nil)

// ContractAVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryTransactor implements the ContractAVSDirectoryTransacts interface.
var _ ContractAVSDirectoryTransacts = (*ContractAVSDirectoryTransactor)(nil)

// ContractAVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAVSDirectoryFilterer implements the ContractAVSDirectoryFilters interface.
var _ ContractAVSDirectoryFilters = (*ContractAVSDirectoryFilterer)(nil)

// ContractAVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAVSDirectorySession struct {
	Contract     *ContractAVSDirectory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContractAVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAVSDirectoryCallerSession struct {
	Contract *ContractAVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ContractAVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAVSDirectoryTransactorSession struct {
	Contract     *ContractAVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractAVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAVSDirectoryRaw struct {
	Contract *ContractAVSDirectory // Generic contract binding to access the raw methods on
}

// ContractAVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAVSDirectoryCallerRaw struct {
	Contract *ContractAVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAVSDirectoryTransactorRaw struct {
	Contract *ContractAVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAVSDirectory creates a new instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectory(address common.Address, backend bind.ContractBackend) (*ContractAVSDirectory, error) {
	contract, err := bindContractAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectory{ContractAVSDirectoryCaller: ContractAVSDirectoryCaller{contract: contract}, ContractAVSDirectoryTransactor: ContractAVSDirectoryTransactor{contract: contract}, ContractAVSDirectoryFilterer: ContractAVSDirectoryFilterer{contract: contract}}, nil
}

// NewContractAVSDirectoryCaller creates a new read-only instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*ContractAVSDirectoryCaller, error) {
	contract, err := bindContractAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryCaller{contract: contract}, nil
}

// NewContractAVSDirectoryTransactor creates a new write-only instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAVSDirectoryTransactor, error) {
	contract, err := bindContractAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryTransactor{contract: contract}, nil
}

// NewContractAVSDirectoryFilterer creates a new log filterer instance of ContractAVSDirectory, bound to a specific deployed contract.
func NewContractAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAVSDirectoryFilterer, error) {
	contract, err := bindContractAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryFilterer{contract: contract}, nil
}

// bindContractAVSDirectory binds a generic wrapper to an already deployed contract.
func bindContractAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSDirectory *ContractAVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.ContractAVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAVSDirectory *ContractAVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORSETFORCEDEREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETFORCEDEREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xdce974b9.
//
// Solidity: function OPERATOR_SET_FORCE_DEREGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORSETFORCEDEREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETFORCEDEREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OPERATORSETREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "OPERATOR_SET_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// OPERATORSETREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xc825fe68.
//
// Solidity: function OPERATOR_SET_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OPERATORSETREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.OPERATORSETREGISTRATIONTYPEHASH(&_ContractAVSDirectory.CallOpts)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) AvsOperatorStatus(opts *bind.CallOpts, avs common.Address, operator common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "avsOperatorStatus", avs, operator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectorySession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _ContractAVSDirectory.Contract.AvsOperatorStatus(&_ContractAVSDirectory.CallOpts, avs, operator)
}

// AvsOperatorStatus is a free data retrieval call binding the contract method 0x49075da3.
//
// Solidity: function avsOperatorStatus(address avs, address operator) view returns(uint8)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) AvsOperatorStatus(avs common.Address, operator common.Address) (uint8, error) {
	return _ContractAVSDirectory.Contract.AvsOperatorStatus(&_ContractAVSDirectory.CallOpts, avs, operator)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Delegation() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Delegation(&_ContractAVSDirectory.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Delegation() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Delegation(&_ContractAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.DomainSeparator(&_ContractAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractAVSDirectory.Contract.DomainSeparator(&_ContractAVSDirectory.CallOpts)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectorySession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractAVSDirectory.CallOpts, operator, salt)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool isSpent)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractAVSDirectory.CallOpts, operator, salt)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Owner() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Owner(&_ContractAVSDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Owner() (common.Address, error) {
	return _ContractAVSDirectory.Contract.Owner(&_ContractAVSDirectory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Paused(index uint8) (bool, error) {
	return _ContractAVSDirectory.Contract.Paused(&_ContractAVSDirectory.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Paused(index uint8) (bool, error) {
	return _ContractAVSDirectory.Contract.Paused(&_ContractAVSDirectory.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Paused0() (*big.Int, error) {
	return _ContractAVSDirectory.Contract.Paused0(&_ContractAVSDirectory.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Paused0() (*big.Int, error) {
	return _ContractAVSDirectory.Contract.Paused0(&_ContractAVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectorySession) PauserRegistry() (common.Address, error) {
	return _ContractAVSDirectory.Contract.PauserRegistry(&_ContractAVSDirectory.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractAVSDirectory.Contract.PauserRegistry(&_ContractAVSDirectory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractAVSDirectory *ContractAVSDirectoryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractAVSDirectory.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractAVSDirectory *ContractAVSDirectorySession) Version() (string, error) {
	return _ContractAVSDirectory.Contract.Version(&_ContractAVSDirectory.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ContractAVSDirectory *ContractAVSDirectoryCallerSession) Version() (string, error) {
	return _ContractAVSDirectory.Contract.Version(&_ContractAVSDirectory.CallOpts)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CancelSalt(&_ContractAVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.CancelSalt(&_ContractAVSDirectory.TransactOpts, salt)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractAVSDirectory.TransactOpts, operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "initialize", initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Initialize(&_ContractAVSDirectory.TransactOpts, initialOwner, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address initialOwner, uint256 initialPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Initialize(initialOwner common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Initialize(&_ContractAVSDirectory.TransactOpts, initialOwner, initialPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Pause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Pause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) PauseAll() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.PauseAll(&_ContractAVSDirectory.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.PauseAll(&_ContractAVSDirectory.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsMixinTypesSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RenounceOwnership(&_ContractAVSDirectory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.RenounceOwnership(&_ContractAVSDirectory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.TransferOwnership(&_ContractAVSDirectory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.TransferOwnership(&_ContractAVSDirectory.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Unpause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.Unpause(&_ContractAVSDirectory.TransactOpts, newPausedStatus)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractAVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractAVSDirectory *ContractAVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractAVSDirectory.TransactOpts, metadataURI)
}

// ContractAVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *ContractAVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryAVSMetadataURIUpdated)
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
		it.Event = new(ContractAVSDirectoryAVSMetadataURIUpdated)
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
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractAVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryAVSMetadataURIUpdatedIterator{contract: _ContractAVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryAVSMetadataURIUpdated)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractAVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(ContractAVSDirectoryAVSMetadataURIUpdated)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryInitializedIterator struct {
	Event *ContractAVSDirectoryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryInitialized)
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
		it.Event = new(ContractAVSDirectoryInitialized)
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
func (it *ContractAVSDirectoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryInitialized represents a Initialized event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractAVSDirectoryInitializedIterator, error) {

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryInitializedIterator{contract: _ContractAVSDirectory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryInitialized)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseInitialized(log types.Log) (*ContractAVSDirectoryInitialized, error) {
	event := new(ContractAVSDirectoryInitialized)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _ContractAVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(ContractAVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOwnershipTransferredIterator struct {
	Event *ContractAVSDirectoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryOwnershipTransferred)
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
		it.Event = new(ContractAVSDirectoryOwnershipTransferred)
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
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAVSDirectoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryOwnershipTransferredIterator{contract: _ContractAVSDirectory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryOwnershipTransferred)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAVSDirectoryOwnershipTransferred, error) {
	event := new(ContractAVSDirectoryOwnershipTransferred)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPausedIterator struct {
	Event *ContractAVSDirectoryPaused // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryPaused)
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
		it.Event = new(ContractAVSDirectoryPaused)
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
func (it *ContractAVSDirectoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryPaused represents a Paused event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryPausedIterator{contract: _ContractAVSDirectory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryPaused)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParsePaused(log types.Log) (*ContractAVSDirectoryPaused, error) {
	event := new(ContractAVSDirectoryPaused)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAVSDirectoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryUnpausedIterator struct {
	Event *ContractAVSDirectoryUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAVSDirectoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAVSDirectoryUnpaused)
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
		it.Event = new(ContractAVSDirectoryUnpaused)
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
func (it *ContractAVSDirectoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAVSDirectoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAVSDirectoryUnpaused represents a Unpaused event raised by the ContractAVSDirectory contract.
type ContractAVSDirectoryUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractAVSDirectoryUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractAVSDirectoryUnpausedIterator{contract: _ContractAVSDirectory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAVSDirectoryUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractAVSDirectory.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAVSDirectoryUnpaused)
				if err := _ContractAVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractAVSDirectory *ContractAVSDirectoryFilterer) ParseUnpaused(log types.Log) (*ContractAVSDirectoryUnpaused, error) {
	event := new(ContractAVSDirectoryUnpaused)
	if err := _ContractAVSDirectory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
