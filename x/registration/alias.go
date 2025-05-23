// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/scrtlabs/SecretNetwork/x/compute/internal/types
// ALIASGEN: github.com/scrtlabs/SecretNetwork/x/compute/internal/keeper
package registration

import (
	"github.com/scrtlabs/SecretNetwork/x/registration/internal/keeper"
	"github.com/scrtlabs/SecretNetwork/x/registration/internal/keeper/enclave"
	"github.com/scrtlabs/SecretNetwork/x/registration/internal/types"
)

const (
	ModuleName                 = types.ModuleName
	StoreKey                   = types.StoreKey
	TStoreKey                  = types.TStoreKey
	QuerierRoute               = types.QuerierRoute
	RouterKey                  = types.RouterKey
	EnclaveSealedData          = types.EnclaveSealedData
	SecretNodeSeedLegacyConfig = types.SecretNodeSeedLegacyConfig
	SecretNodeSeedNewConfig    = types.SecretNodeSeedNewConfig
	SecretNodeCfgFolder        = types.SecretNodeCfgFolder
	EncryptedKeyLength         = types.EncryptedKeyLength
	LegacyEncryptedKeyLength   = types.LegacyEncryptedKeyLength
	AttestationCertPath        = types.AttestationCertPath
	AttestationCombinedPath    = types.AttestationCombinedPath
	IoExchMasterKeyPath        = types.IoExchMasterKeyPath
	LegacyIoMasterCertificate  = types.LegacyIoMasterCertificate
	NodeExchMasterKeyPath      = types.NodeExchMasterKeyPath
	SeedPath                   = types.SeedPath
	MasterIoKeyId              = types.MasterIoKeyId
	MasterNodeKeyId            = types.MasterNodeKeyId
	SeedConfigVersion          = types.SeedConfigVersion
)

var (
	// functions aliases
	RegisterCodec               = types.RegisterLegacyAminoCodec
	RegisterInterfaces          = types.RegisterInterfaces
	ValidateGenesis             = types.ValidateGenesis
	InitGenesis                 = keeper.InitGenesis
	ExportGenesis               = keeper.ExportGenesis
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	GetGenesisStateFromAppState = keeper.GetGenesisStateFromAppState
	IsHexString                 = keeper.IsHexString
	GetApiKey                   = types.GetApiKey
	GetSpid                     = types.GetSpid
	// variable aliases
	ModuleCdc               = types.ModuleCdc
	DefaultCodespace        = types.DefaultCodespace
	ErrAuthenticateFailed   = types.ErrAuthenticateFailed
	ErrSeedInitFailed       = types.ErrSeedInitFailed
	RegistrationStorePrefix = types.RegistrationStorePrefix
	ErrInvalidType          = types.ErrInvalidType
)

type (
	MsgRaAuthenticate    = types.RaAuthenticate
	GenesisState         = types.GenesisState
	Keeper               = keeper.Keeper
	SeedConfig           = types.SeedConfig
	LegacySeedConfig     = types.LegacySeedConfig
	EnclaveApi           = enclave.Api
	MasterKey            = types.MasterKey
	Key                  = types.Key
	RegistrationNodeInfo = types.RegistrationNodeInfo 
)
