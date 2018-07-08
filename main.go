package main

import (
	"UBStoYNAB/csvExport"
	"UBStoYNAB/ubsApi"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starte UBS E-Banking crawler...")
	config := loadConfiguration()
	if login(config.ContractNumber) {
		//export normal accounts
		accounts := ubsApi.GetAvailableAccounts()
		for index, element := range accounts {
			fmt.Println("Account ", index)
			fmt.Println("Alias: ", element.Alias)
			fmt.Println("Balance: ", element.Balance)
			fmt.Println("Try to export transactions")

			csvExport.ExportNormalAccountToCSV(ubsApi.GetAccountTransactions(element.ID, 350, "20180101"), element.Alias)
		}

		//export credit cards
		creditCardAccounts := ubsApi.GetAvailableCreditCardAccounts()
		for index, element := range creditCardAccounts {
			fmt.Println("Account ", index)
			fmt.Println("Alias: ", element.Alias)
			fmt.Println("Balance: ", element.Balance)

			creditCards := ubsApi.GetAvailableCreditCards(element.ID)
			for index, card := range creditCards {
				fmt.Println("-->Card ", index)
				fmt.Println("-->Alias: ", card.ProductText)

				fmt.Println("Try to export transactions")
				cardTransactions, accountTransactions := ubsApi.GetCardTransactions(card.ID, 150, "20180301", "20180707")
				csvExport.ExportCreditCardToCSV(cardTransactions, accountTransactions, card.Alias)
			}
		}
	}
}

func login(contractNumber string) bool {
	for i := 0; i < 3; i++ {
		responses := getChallengeInput(ubsApi.GetAuthenticatorChallenge(contractNumber))

		if ubsApi.SendAuthenticatorChallengeResponse(responses[0], responses[1], responses[2], responses[3]) {
			return true
		} else {
			fmt.Println("Antwort wurde nicht akzeptiert! Versuche es erneut..")
		}
	}
	fmt.Println("Warnung: Zu viele falsche Anmeldeversuche können zu einer vorübergehenden Sperrung des E-Bankings führen.")
	return false
}

func getChallengeInput(challenge string) []string {

	fmt.Println("Eingabe - Kartenleser oder Access Card Display")
	fmt.Println(challenge)
	for {
		fmt.Println("\nSicherheitscode(XX XX XX XX):")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		responses := strings.Fields(strings.ToUpper(scanner.Text()))
		if len(responses) == 4 {
			return responses
		} else {
			fmt.Println("Invalid Input")
		}
	}
}

func loadConfiguration() Configuration {
	var config Configuration
	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

type Configuration struct {
	ContractNumber string
}
