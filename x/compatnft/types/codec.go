package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "compatnft/IssueDenom", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "compatnft/MintNFT", nil)
	cdc.RegisterConcrete(&MsgEditNFT{}, "compatnft/EditNFT", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "compatnft/TransferNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "compatnft/BurnNFT", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnNFT{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
