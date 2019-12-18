package isvalid

import (
	"fmt"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
)

// PubKey checks if it is a valid public key
// Use nnetwork as "Public" for horizonclient.DefaultPublicNetClient
// Use nnetwork as "Testnet" for horizonclient.DefaultTestNetClient
func PubKey(pubkey string, nnetwork string) bool {

	//Checking if stellar address is valid or not
	nclient := horizonclient.DefaultTestNetClient
	if nnetwork == "Public" {
		nclient = horizonclient.DefaultPublicNetClient
	}
	accountRequest := horizonclient.AccountRequest{AccountID: pubkey}

	_, err := nclient.AccountDetail(accountRequest)
	if err != nil {
		fmt.Println("Invalid Public Key, please check and try again!")
		return false
	}

	return true
}

// SecKey checks if it is a valid secret key
func SecKey(seckey string) bool {

	//Checking if stellar secret key is valid or not
	initialchar := string(seckey[0])
	_, err := keypair.Parse(seckey)
	if err != nil || initialchar != "S" {
		fmt.Println("Invalid Secret Key, please check and try again!")
		return false
	}

	return true
}

// TLen checks if it is a valid secret length
func TLen(assetcode string) bool {

	//maximum length is 12 alphanumeric characters
	assetLength := len(assetcode)
	if assetLength > 12 {
		fmt.Println("Invalid Asset Length, maximum length is 12 alphanumeric characters!")
		return false
	}

	return true
}
