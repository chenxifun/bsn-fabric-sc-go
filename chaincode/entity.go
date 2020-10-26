package chaincode

//nft input and out

type Input struct {
	ABIEncoded   string `json:"abi_encoded,omitempty"`
	To           string `json:"to"`
	AmountToMint string `json:"amount_to_mint"`
	MetaID       string `json:"meta_id"`
	SetPrice     string `json:"set_price"`
	IsForSale    bool   `json:"is_for_sale"`
}

type Output struct {
	NftID string `json:"nft_id"`
}
