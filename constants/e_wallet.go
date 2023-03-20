package constants

import "fmt"

const Dana = "dana"
const GoPay = "gopay"
const LinkAja = "linkaja"
const LinkAjaApp = "linkaja_app"
const Ovo = "ovo"
const ShopeePay = "shopeepay"
const ShopeePayApp = "shopeepay_app"
const Qris = "qris"

func GetAllEWallet() map[string]string {
	mapWallet := map[string]string{
		Dana:         "Dana",
		GoPay:        "GoPay",
		LinkAja:      "LinkAja",
		LinkAjaApp:   "LinkAja App",
		Ovo:          "OVO",
		ShopeePay:    "ShopeePay",
		ShopeePayApp: "ShopeePay",
		Qris:         "QRIS",
	}
	return mapWallet
}

func GetWalletLabel(walletCode string) (string, error) {
	allWallet := GetAllEWallet()
	if _, ok := allWallet[walletCode]; ok {
		return allWallet[walletCode], nil
	}

	// https://docs.flip.id/?php#destination-bank
	return "", fmt.Errorf("e-wallet not set, please check flip documentation for any e-wallet update")
}
