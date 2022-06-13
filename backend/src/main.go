package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/ethclient"

	birden "main/contracts"
)

var birdenAddr = "0xDB9E8D1acc5b422FA3843e9F37643e1b7A595a0c"

// TokenRequest wallet
type TokenRequest struct {
	Wallet  string `json:"wallet"`
	TokenID int    `json:"tokenID"`
}

// TokenResponse token
type TokenResponse struct {
	Token string `json:"token"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/token", TokenHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}

// TokenHandler token
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	var tokenRequest TokenRequest

	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &tokenRequest)
	log.Printf("LoginHandler received: %v\n", tokenRequest)

	client, err := ethclient.Dial("https://alfajores-forno.celo-testnet.org")
	if err != nil {
		log.Fatal(err)
	}

	contract, err := birden.NewBirden(common.HexToAddress(birdenAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	tokens, err := contract.TokensInfo(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	tokenID := big.NewInt(int64(tokenRequest.TokenID))
	tokenRentedBy := ""
	for _, token := range tokens {
		if token.TokenId == tokenID {
			tokenRentedBy = token.RentedBy.String()
		}
	}

	if strings.ToLower(tokenRentedBy) != strings.ToLower(tokenRequest.Wallet) {
		log.Fatal("not owner")
	}

	var tokenResponse TokenResponse
	tokenResponse.Token = getPinataKey()

	jsonResponse, _ := json.Marshal(tokenResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func getPinataKey() string {
	url := "https://managed.mypinata.cloud/api/v1/auth/content/jwt"
	method := "POST"

	payload := strings.NewReader(`{
    "timeoutSeconds": 3600,
    "contentIds": ["bafybeidvt4trjo7kdi6ofyiah6626lzssdc2244zit6lh5vsg4ndtt5hea"]
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("x-api-key", "key")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}
