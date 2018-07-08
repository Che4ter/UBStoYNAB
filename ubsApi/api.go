package ubsApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const BASE_URL = "https://mobs-ch1.ubs.com"
const SECOND_URL = "https://mokp-ch.ubs.com"

var UBSCookies []*http.Cookie
var UBSCookieJar *cookiejar.Jar
var USBContractNumber string
var USBRequestKey string

//do init before all others
func initData(contractNumber string) {
	UBSCookies = nil
	//var err error;
	UBSCookieJar, _ = cookiejar.New(nil)

	USBContractNumber = contractNumber
}

//login
func GetAuthenticatorChallenge(contractNumber string) (string){
	initData(contractNumber)

	requestURL := "/ClientWorkbenchSystem/MobileInterfaceFoundation/V1/MobileInterfaceFoundation/getNavigation/8.1.0/de?login"

	client := &http.Client{
		Jar:           UBSCookieJar,
	}

	v := url.Values{}
	v.Set("isiwebmethod", "authenticate")
	v.Set("isiwebuserid", USBContractNumber)
	v.Set("isiwebargs", "login&language=de&_="+strconv.FormatInt(NowAsUnixMilli(), 10))

	req, err := http.NewRequest("POST", BASE_URL + requestURL, bytes.NewBufferString(v.Encode()))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE; Navajo=")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	UBSCookies = UBSCookieJar.Cookies(req.URL)

	var responseDTO ChallengeResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	var challenge1, challenge2, challenge3 string
	for _, element := range responseDTO.Gui.GuiElem {
		if(element.Name == "errorcode" && element.Label != ""){
			fmt.Println(element.Label)
			os.Exit(-1)
		}
		if element.Name == "challenge1" {
			challenge1 = element.Value
		}else if element.Name == "challenge2" {
			challenge2 = element.Value
		}else if element.Name == "challenge3" {
			challenge3 = element.Value
		}
	}

	return challenge1 + " " + challenge2 + " " + challenge3
}

func SendAuthenticatorChallengeResponse(response1 string, response2 string, response3 string, response4 string) bool {
	requestURL := "/ClientWorkbenchSystem/MobileInterfaceFoundation/V1/MobileInterfaceFoundation/getNavigation/8.1.0/de?login"

	client := &http.Client{
		CheckRedirect: nil,
		Jar:           UBSCookieJar,
	}

	v := url.Values{}
	v.Set("authid", USBContractNumber)
	v.Set("isiwebuserid", USBContractNumber)
	v.Set("isiwebuserid_check", USBContractNumber)
	v.Set("authenticate", "authenticate")
	v.Set("response1", response1)
	v.Set("response2", response2)
	v.Set("response3", response3)
	v.Set("response4", response4)

	v.Set("isiwebargs", "login&language=de&_="+strconv.FormatInt(NowAsUnixMilli(), 10))

	req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBufferString(v.Encode()))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	UBSCookies = UBSCookieJar.Cookies(req.URL)

	var responseDTO ChallengeCompleteResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	if responseDTO.Status == "OK" {
		USBRequestKey = responseDTO.RequestKey
		return true
	} else {
		return false
	}
}

//normal accounts
func GetAvailableAccounts() []CashAccounts {

	requestURL := "/ClientWorkbenchSystem/MobileInterfaceAssetView/V1/MobileInterfaceAssetView/getCashAccountOverview"

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	request := GetAccountsRequest{0, 30, "alias", "de", USBRequestKey}

	jsonString, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBuffer(jsonString))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var responseDTO GetAccountsResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	return responseDTO.CashAccounts
}

func GetAccountTransactions(accountId string, numRecords int, startDate string) []CashAccountTrxes {

	requestURL := "/ClientWorkbenchSystem/MobileInterfaceAssetView/V1/MobileInterfaceAssetView/getCashAccountTrx"

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	request := AccountBalanceRequest{accountId, 0, numRecords, true, false, "de", startDate, USBRequestKey}

	jsonString, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBuffer(jsonString))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var responseDTO AccountBalanceResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	return responseDTO.CashAccountTrxes
}

//credit cards
func GetAvailableCreditCardAccounts() []CreditCardAccounts {
	requestURL := "/ClientWorkbenchSystem/MobileInterfaceCards/V1/MobileInterfaceCards/getCreditCardAccountOverview"

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	request := CreditCardOverviewRequest{false, 0, 15, "de", USBRequestKey}

	jsonString, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBuffer(jsonString))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var responseDTO CreditCardOverviewResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	return responseDTO.CreditCardAccounts
}

func GetAvailableCreditCards(accountId string) []CreditCards {
	requestURL := "/ClientWorkbenchSystem/MobileInterfaceCards/V1/MobileInterfaceCards/getCreditCardAccountData"

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	request := CreditCardDetailsRequest{accountId, 0, 15, false, "de", USBRequestKey}

	jsonString, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBuffer(jsonString))
	req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
	req.Header.Add("content-type", "application/json")
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var responseDTO CreditCardDetailsResponse
	if err := json.Unmarshal(data, &responseDTO); err != nil {
		panic(err)
	}

	return responseDTO.CreditCards
}

func GetCardTransactions(accountId string, numRecords int, startDate string, endDate string) ([]CardTransactions, []CreditCardAccountTransactions) {

	requestURL := "/ClientWorkbenchSystem/MobileInterfaceCards/V1/MobileInterfaceCards/getInvoiceDetails"

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	hasMore := true
	var creditCardTransactions []CardTransactions
	var accountTransactions []CreditCardAccountTransactions
	for hasMore {
		request := CreditCardTransactionsRequest{accountId, "CHF", startDate, endDate, 0, 150, 0, 150, false, "de", USBRequestKey}

		jsonString, err := json.Marshal(request)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", BASE_URL+requestURL, bytes.NewBuffer(jsonString))
		req.Header.Add("cookie", "NavLB_MOBS=mobs-ch1.ubs.com; bzelang=en-DE;")
		req.Header.Add("content-type", "application/json")
		if err != nil {
			log.Println(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		defer req.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		var responseDTO CreditCardTransactionsResponse
		if err := json.Unmarshal(data, &responseDTO); err != nil {
			panic(err)
		}

		if !responseDTO.HasMoreCardTransactions {
			hasMore = false
		} else {
			startDate = responseDTO.CardTransactions[0].TransactionDate
		}

		creditCardTransactions = append(creditCardTransactions, responseDTO.CardTransactions...)
		accountTransactions = append(accountTransactions, responseDTO.AccountTransactions...)
	}

	return creditCardTransactions, accountTransactions
}


//others
func PingRequest(contractNumber string) {
	requestURL := "/auth/ismobilekeyavailable?login"

	v := url.Values{}
	v.Set("isiwebuserid", contractNumber)
	v.Set("language", "de")

	s := v.Encode()
	fmt.Printf("v.Encode(): %v\n", s)

	req, err := http.NewRequest("POST", SECOND_URL+requestURL, strings.NewReader(s))
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		Jar: UBSCookieJar,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("http.Do() error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() error: %v\n", err)
		return
	}

	UBSCookies = UBSCookieJar.Cookies(req.URL)

	fmt.Println(string(data))

	var response PingResponse
	if err := json.Unmarshal(data, &response); err != nil {
		panic(err)
	}

	fmt.Println(response)
}
