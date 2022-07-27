package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/logrusorgru/aurora"
)

// func GetPubKeyByArmor(pubKeyArmor string) cryptotypes.PubKey {
// 	pub, _, err := crypto.UnarmorPubKeyBytes(pubKeyArmor)
// 	if err != nil {
// 		LogGw(err)
// 	}

// 	return pub
// }

func GetAddrByPrivKeyArmor(priKeyArmor string) sdk.AccAddress {
	return GetAddrByPrivKey(GetPriKeyByArmor(priKeyArmor))
}

func GetPriKeyByArmor(priKeyArmor string) cryptotypes.PrivKey {
	keyOwnerPw := GetConfig().Get("keyOwnerPw")
	priv, _, err := crypto.UnarmorDecryptPrivKey(priKeyArmor, keyOwnerPw)
	if err != nil {
		LogGw(err)
	}

	return priv
}

func GetAddrByPrivKey(priv cryptotypes.PrivKey) sdk.AccAddress{
	gwAdd, err := sdk.AccAddressFromHex(priv.PubKey().Address().String())
	if err != nil {
		LogGw(err)
	}

	return gwAdd
}

func ConvertConfigParam(str string) []string {
	var strList []string
	if strings.Contains(str, "mnemonic") {
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
		strList[0] = strings.Replace(strList[0], " ", "", -1)
		strList[1] = strings.TrimRight(strList[1], " ")
	} else if strings.Contains(str, "http") {
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
		strList[1] = strings.Join(strList[1:], ":")
	} else {
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "\r", "", -1)
		str = strings.Replace(str, "\n", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, "\"", "", -1)
		strList = strings.Split(str, ":")
	}
	
	return strList
}

func ToString(value interface{}, defaultValue string) string {
	str := strings.TrimSpace(fmt.Sprintf("%v", value))
	if str == "" {
		return defaultValue
	} else {
		return str
	}
}

func FromStringToUint64(value string) uint64 {
	number, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		LogGw(err)
	}

	return number
}

func FromStringToInt64(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		LogGw(err)
	}

	return number
}

func JsonUnmarshal(jsonStruct interface{}, jsonFilePath string) interface{} {
	jsonData, err := os.Open(jsonFilePath) 
	if err != nil {
		LogGw(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonData)
	jsonStruct = JsonUnmarshalData(jsonStruct, byteValue)

	return jsonStruct
}

func JsonUnmarshalData(jsonStruct interface{}, byteValue []byte) interface{} {
	json.Unmarshal(byteValue, &jsonStruct)

	return jsonStruct
}

func JsonMarshal(jsonData interface{}, jsonFilePath string) {
	byteData, err := JsonMarshalData(jsonData)
	if err != nil {
		LogGw(err)
	}
	err = ioutil.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
	if err != nil {		
		LogGw(err)
		path := strings.Split(jsonFilePath, "/")
		pathPop := path[:len(path)-1]
		filePath := strings.Join(pathPop, "/")		

		err := os.Mkdir(filePath, 0755)
		if err != nil {
			LogGw(err)
		}
		err = ioutil.WriteFile(jsonFilePath, byteData, os.FileMode(0644))
	}
}

func JsonMarshalData(jsonData interface{}) ([]byte, error) {
	byteData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		LogGw(err)
	}

	return byteData, err
}

func LogGw(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.White("Gateway    ").String() + str)
}

func LogHttpServer(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.Blue("HTTPServer ").String() + str)
}

func LogHttpClient(log ...interface{}) {
	str := ToString(log, "")
	fmt.Println(aurora.Green("HTTPClient ").String() + str)
}