package loopring

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/katzenpost/core/crypto/eddsa"
)

var BASE_API_ENDPOINT = "https://api3.loopring.io/api/v3"

var ErrNotImplemented = errors.New("not implemented")
var ErrMethodNotSupported = func(method string) error {
	err := fmt.Errorf("method '%s' is not supported. Please use GET, POST, PUT, or DELETE", method)
	return err
}

type QueryParamPair struct {
	QueryParam string
	Value      string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ConstructSignatureBase(httpMethod string, apiPath string, params ...QueryParamPair) string {
	// Make sure http method is supported by Loopring
	acceptedMethods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	if !contains(acceptedMethods, httpMethod) {
		log.Fatal(ErrMethodNotSupported(httpMethod))
	}

	// Initialize signatureBase to a new string
	var signatureBase = ""

	// Append the API's HTTP method to signatureBase.
	signatureBase += httpMethod

	// Append '＆' to signatureBase
	signatureBase += "&"

	// Append percent-encoded full URL path (without ? or any query parameters) to signatureBase
	signatureBase += url.QueryEscape(BASE_API_ENDPOINT + apiPath)

	// Append '&' to signatureBase
	signatureBase += "&"

	// Initialize parameterString to an empty string
	var parameterString = ""

	// GET/DELETE requests
	if httpMethod == http.MethodGet || httpMethod == http.MethodDelete {
		// Sort query parameters in ascending order lexicographically
		sort.Slice(params, func(i, j int) bool {
			return params[i].QueryParam < params[j].QueryParam
		})

		// Append key name to parameterString
		// Append an '=' to parameterString;
		// Append value to parameterString;
		// Append a '&' if there are more key value pairs.
		for i, s := range params {
			parameterString += s.QueryParam
			parameterString += "="
			parameterString += s.Value
			if i != len(params)-1 {
				parameterString += "&"
			}
		}
	}

	// POST/PUT Requests
	/*
		if httpMethod == http.MethodPost || httpMethod == http.MethodPut {
			parameterString += requestBody
		}
	*/

	signatureBase += url.QueryEscape(parameterString)

	return signatureBase
}

func FromBase16(base16 string) []byte {
	parsedHex := strings.Replace(base16, "0x", "", -1)
	i, ok := new(big.Int).SetString(parsedHex, 16)
	if !ok {
		log.Fatalln("trying to convert from base16 a bad number: ", base16)
	}
	return i.Bytes()
}

type UserData struct {
	address  string
	password string
}

func consumeData() UserData {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	address := os.Getenv("LOOPRING_ACCOUNT_ADDRESS")
	password := os.Getenv("LOOPRING_ACCOUNT_PASSWORD")

	userData := UserData{}
	userData.address = address
	userData.password = password

	return userData
}

func GenerateLoopringKeyPair() *eddsa.PrivateKey {
	var userData UserData = consumeData()

	var hashedPassword = crypto.Keccak256Hash([]byte(userData.password))
	r := strings.NewReader("LOOPRING" + strings.ToLower(userData.address) + hashedPassword.String())

	eddsaKeyPair, err := eddsa.NewKeypair(r)
	if err != nil {
		panic(err)
	}

	return eddsaKeyPair
}
