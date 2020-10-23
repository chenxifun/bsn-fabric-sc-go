package chaincode

import "math/big"

//nft input and out

type Input struct {
	ABIEncoded   string   `json:"abi_encoded"`
	To           []byte   `json:"to"`
	AmountToMint *big.Int `json:"amount_to_mint"`
	MetaID       string   `json:"meta_id"`
	SetPrice     *big.Int `json:"set_price"`
	IsForSale    bool     `json:"is_for_sale"`
}

type Output struct {
	NftID string `json:"nft_id"`
}
