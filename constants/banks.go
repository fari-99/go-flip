package constants

import (
	"fmt"
)

/*
Bank List
https://docs.flip.id/?php#destination-bank
*/

const HardaInter = "harda"
const Anz = "anz"
const AcehSyariah = "aceh"
const AladinSyariah = "aladin"
const AntarDaerah = "antardaerah"
const Artha = "artha"
const Bengkulu = "bengkulu"
const BpdDIY = "daerah_istimewa"
const BpdDiySyariah = "daerah_istimewa_syr"
const BPTNSyariah = "btpn_syr"
const BukopinSyariah = "bukopin_syr"
const BumiArta = "bumi_arta"
const CapitalIndo = "capital"
const BCA = "bca"
const CCBIndo = "ccb"
const CNB = "cnb"
const DanamonSyariah = "danamon"
const DinarIndo = "dinar"
const DKI = "dki"
const DKISyariah = "dki_syr"
const Ganesha = "ganesha"
const IBKIndo = "agris"
const InaPerdana = "ina_perdana"
const IndexSelindo = "index_selindo"
const JagoSyariah = "artos_syr"
const Jambi = "jambi"
const JambiSyariah = "jambi_syr"
const JasaJakarta = "jasa_jakarta"
const JaTeng = "jawa_tengah"
const JaTengSyariah = "jawa_tengah_syr"
const JaTim = "jawa_timur"
const JaTimSyariah = "jawa_timur_syr"
const KalBar = "kalimantan_barat"
const KalBarSyariah = "kalimantan_barat_syr"
const KalSel = "kalimantan_selatan"
const KalSelSyariah = "kalimantan_selatan_syr"
const KalTeng = "kalimantan_tengah"
const KalTimSyariah = "kalimantan_timur_syr"
const KalTim = "kalimantan_timur"
const Lampung = "lampung"
const Maluku = "maluku"
const Mandiri = "mandiri"
const Mantap = "mantap"
const Maspion = "maspion"
const Mayapada = "mayapada"
const MayoraIndo = "mayora"
const Mega = "mega"
const MegaSyariah = "mega_syr"
const MestikaDharma = "mestika_dharma"
const MultiArtaSentosa = "mas"
const Mutiara = "mutiara"
const Nagari = "sumatera_barat"
const NagariSyariah = "sumatera_barat_syr"
const NTBSyariah = "nusa_tenggara_barat"
const NTT = "nusa_tenggara_timur"
const NusaParahyangan = "nusantara_parahyangan"
const OCBC = "ocbc"
const OCBCSyariah = "ocbc_syr"
const BOCLimited = "boc"
const IndiaIndo = "india"
const TokyoMitsubishi = "tokyo"
const Papua = "papua"
const Prima = "prima"
const BRI = "bri"
const RiauKepri = "riau_dan_kepri"
const Sampoerna = "sahabat_sampoerna"
const ShinhanIndo = "shinhan"
const Sinarmas = "sinarmas"
const SinarmasSyariah = "sinarmas_syr"
const Sulselbar = "sulselbar"
const SulselbarSyariah = "sulselbar_syr"
const Sulawesi = "sulawesi"
const Sultra = "sulawesi_tenggara"
const SulutGo = "sulut"
const SumselBabel = "sumsel_dan_babel"
const SumselBabelSyariah = "sumsel_dan_babel_syr"
const Sumut = "sumut"
const SumutSyariah = "sumut_syr"
const VictoriaInternational = "victoria_internasional"
const VictoriaSyariah = "victoria_syr"
const WooriSaudara = "woori"
const BCASyariah = "bca_syr"
const BJB = "bjb"
const BJBSyariah = "bjb_syr"
const BCADigital = "royal"
const BNI = "bni"
const BPDBali = "bali"
const BPDBanten = "banten"
const Eka = "eka"
const BRIAgroniaga = "agroniaga"
const BankSyariahIndo = "bsm"
const BTNSyariah = "btn"
const BTPN = "tabungan_pensiunan_nasional"
const CIMB = "cimb"
const Citibank = "citibank"
const CommonWealth = "commonwealth"
const CTBCIndo = "chinatrust"
const DBSIndo = "dbs"
const HSBCIndo = "hsbc"
const ICBCIndo = "icbc"
const Artos = "artos"
const Hana = "hana"
const MaybankIndo = "bii"
const MaybankSyariah = "bii_syr"
const MNC = "mnc_internasional"
const Muamalat = "muamalat"
const YudhaBhakti = "yudha_bakti"
const Nobu = "nationalnobu"
const Panin = "panin"
const PaninSyariah = "panin_syr"
const Permata = "permata"
const PermataSyariah = "permata_syr"
const QNBIndo = "qnb_kesawan"
const RaboBankInterIndo = "rabobank"
const SBIINdo = "sbi_indonesia"
const SeabankBKE = "kesejahteraan_ekonomi"
const StandardChartered = "standard_chartered"
const UOB = "uob"
const Bukopin = "bukopin"

func GetAllBankLabel() map[string]string {
	mapBank := map[string]string{
		HardaInter:            "Allo Bank/Bank Harda Internasional",
		Anz:                   "ANZ Indonesia",
		AcehSyariah:           "Bank Aceh Syariah",
		AladinSyariah:         "Bank Aladin Syariah",
		AntarDaerah:           "Bank Antardaerah",
		Artha:                 "Bank Artha Graha Internasional",
		Bengkulu:              "Bank Bengkulu",
		BpdDIY:                "Bank BPD DIY",
		BpdDiySyariah:         "Bank BPD DIY Syariah",
		BPTNSyariah:           "Bank BTPN Syariah",
		BukopinSyariah:        "Bank Bukopin Syariah",
		BumiArta:              "Bank Bumi Arta",
		CapitalIndo:           "Bank Capital Indonesia",
		BCA:                   "Bank Central Asia",
		CCBIndo:               "Bank China ruction Bank Indonesia",
		CNB:                   "Bank CNB (Centratama Nasional Bank)",
		DanamonSyariah:        "Bank Danamon &amp; Danamon Syariah",
		DinarIndo:             "Bank Dinar Indonesia",
		DKI:                   "Bank DKI",
		DKISyariah:            "Bank DKI Syariah",
		Ganesha:               "Bank Ganesha",
		IBKIndo:               "Bank IBK Indonesia",
		InaPerdana:            "Bank Ina Perdana",
		IndexSelindo:          "Bank Index Selindo",
		JagoSyariah:           "Bank Jago Syariah",
		Jambi:                 "Bank Jambi",
		JambiSyariah:          "Bank Jambi Syariah",
		JasaJakarta:           "Bank Jasa Jakarta",
		JaTeng:                "Bank Jateng",
		JaTengSyariah:         "Bank Jateng Syariah",
		JaTim:                 "Bank Jatim",
		JaTimSyariah:          "Bank Jatim Syariah",
		KalBar:                "Bank Kalbar",
		KalBarSyariah:         "Bank Kalbar Syariah",
		KalSel:                "Bank Kalsel",
		KalSelSyariah:         "Bank Kalsel Syariah",
		KalTeng:               "Bank Kalteng",
		KalTimSyariah:         "Bank Kaltim Syariah",
		KalTim:                "Bank Kaltimtara",
		Lampung:               "Bank Lampung",
		Maluku:                "Bank Maluku",
		Mandiri:               "Bank Mandiri",
		Mantap:                "Bank MANTAP (Mandiri Taspen)",
		Maspion:               "Bank Maspion Indonesia",
		Mayapada:              "Bank Mayapada",
		MayoraIndo:            "Bank Mayora Indonesia",
		Mega:                  "Bank Mega",
		MegaSyariah:           "Bank Mega Syariah",
		MestikaDharma:         "Bank Mestika Dharma",
		MultiArtaSentosa:      "Bank Multi Arta Sentosa (Bank MAS)",
		Mutiara:               "Bank Mutiara",
		Nagari:                "Bank Nagari",
		NagariSyariah:         "Bank Nagari Syariah",
		NTBSyariah:            "Bank NTB Syariah",
		NTT:                   "Bank NTT",
		NusaParahyangan:       "Bank Nusantara Parahyangan",
		OCBC:                  "Bank OCBC NISP",
		OCBCSyariah:           "Bank OCBC NISP Syariah",
		BOCLimited:            "Bank of China (Hong Kong) Limited",
		IndiaIndo:             "Bank of India Indonesia",
		TokyoMitsubishi:       "Bank of Tokyo Mitsubishi UFJ",
		Papua:                 "Bank Papua",
		Prima:                 "Bank Prima Master",
		BRI:                   "Bank Rakyat Indonesia",
		RiauKepri:             "Bank Riau Kepri",
		Sampoerna:             "Bank Sahabat Sampoerna",
		ShinhanIndo:           "Bank Shinhan Indonesia",
		Sinarmas:              "Bank Sinarmas",
		SinarmasSyariah:       "Bank Sinarmas Syariah",
		Sulselbar:             "Bank Sulselbar",
		SulselbarSyariah:      "Bank Sulselbar Syariah",
		Sulawesi:              "Bank Sulteng",
		Sultra:                "Bank Sultra",
		SulutGo:               "Bank SulutGo",
		SumselBabel:           "Bank Sumsel Babel",
		SumselBabelSyariah:    "Bank Sumsel Babel Syariah",
		Sumut:                 "Bank Sumut",
		SumutSyariah:          "Bank Sumut Syariah",
		VictoriaInternational: "Bank Victoria International",
		VictoriaSyariah:       "Bank Victoria Syariah",
		WooriSaudara:          "Bank Woori Saudara",
		BCASyariah:            "BCA (Bank Central Asia) Syariah",
		BJB:                   "BJB",
		BJBSyariah:            "BJB Syariah",
		BCADigital:            "Blu/BCA Digital",
		BNI:                   "BNI (Bank Negara Indonesia)",
		BPDBali:               "BPD Bali",
		BPDBanten:             "BPD Banten",
		Eka:                   "BPR EKA (Bank Eka)",
		BRIAgroniaga:          "BRI Agroniaga",
		BankSyariahIndo:       "BSI (Bank Syariah Indonesia)",
		BTNSyariah:            "BTN/BTN Syariah",
		BTPN:                  "BTPN/Jenius/BTPN Wow!",
		CIMB:                  "CIMB Niaga &amp; CIMB Niaga Syariah",
		Citibank:              "Citibank",
		CommonWealth:          "Commonwealth Bank",
		CTBCIndo:              "CTBC (Chinatrust) Indonesia",
		DBSIndo:               "DBS Indonesia",
		HSBCIndo:              "HSBC Indonesia",
		ICBCIndo:              "ICBC Indonesia",
		Artos:                 "Jago/Artos",
		Hana:                  "LINE Bank/KEB Hana",
		MaybankIndo:           "Maybank Indonesia",
		MaybankSyariah:        "Maybank Syariah",
		MNC:                   "Motion/MNC Bank",
		Muamalat:              "Muamalat",
		YudhaBhakti:           "Neo Commerce/Yudha Bhakti",
		Nobu:                  "Nobu (Nationalnobu) Bank",
		Panin:                 "Panin Bank",
		PaninSyariah:          "Panin Dubai Syariah",
		Permata:               "Permata",
		PermataSyariah:        "Permata Syariah",
		QNBIndo:               "QNB Indonesia",
		RaboBankInterIndo:     "Rabobank International Indonesia",
		SBIINdo:               "SBI Indonesia",
		SeabankBKE:            "Seabank/Bank BKE",
		StandardChartered:     "Standard Chartered Bank",
		UOB:                   "TMRW/UOB",
		Bukopin:               "Wokee/Bukopin",
	}

	return mapBank
}

func GetLabelBank(bankCode string) (string, error) {
	mapBank := GetAllBankLabel()
	if _, ok := mapBank[bankCode]; ok {
		return mapBank[bankCode], nil
	}

	// https://docs.flip.id/?php#destination-bank
	return "", fmt.Errorf("bank not set, please check flip documentation for any bank update")
}

/*
---------------- INQUIRY STATUS
https://docs.flip.id/?php#bank-account-inquiry
*/

// InquiryStatusPending Inquiry still in process
const InquiryStatusPending = "PENDING"

// InquiryStatusSuccess Inquiry process is complete and bank account number is valid
const InquiryStatusSuccess = "SUCCESS"

// InquiryStatusInvalidAccountNumber
// Inquiry process is complete but the account number is invalid or maybe a virtual account number
const InquiryStatusInvalidAccountNumber = "INVALID_ACCOUNT_NUMBER"

// InquiryStatusSuspectedAccount
// Bank account have been suspected on doing fraud. You still can do a disbursement to this account.
const InquiryStatusSuspectedAccount = "SUSPECTED_ACCOUNT"

// InquiryStatusBlackListed
// Bank account have been confirmed on doing a fraud and therefore is blacklisted.
// You can’t do a disbursement to this account.
const InquiryStatusBlackListed = "BLACK_LISTED"

// InquiryStatusFailed The inquiry process is failed before we get the final status of the inquiry,
// e.g. due to timeout or any other errors from the bank.
// If you get this response, please retry the inquiry to trigger reverification of the account.
const InquiryStatusFailed = "FAILED"

// InquiryStatusClosed
// The inquiry process is complete and the account is valid, but it is closed/inactive so that it cannot receive money.
// You cannot do a disbursement to this account.
const InquiryStatusClosed = "CLOSED"

// GetAllBankInquiryStatus
// this function to get all status inquiry and see if you can do transfer or not
func GetAllBankInquiryStatus() map[string]bool {
	inquiryStatus := map[string]bool{
		InquiryStatusPending:              true,
		InquiryStatusSuccess:              true,
		InquiryStatusSuspectedAccount:     true,
		InquiryStatusInvalidAccountNumber: false,
		InquiryStatusBlackListed:          false,
		InquiryStatusFailed:               false,
		InquiryStatusClosed:               false,
	}

	return inquiryStatus
}

func CanTransferBank(inquiryStatus string) (bool, error) {
	inquiryStatuses := GetAllBankInquiryStatus()
	if value, ok := inquiryStatuses[inquiryStatus]; !ok {
		return false, fmt.Errorf("inquiry status [%s] not found", inquiryStatus)
	} else {
		return value, nil
	}
}

func GetAllBankInquiryStatusLabel() map[string]string {
	inquiryStatus := map[string]string{
		InquiryStatusPending:              "Inquiry still in process",
		InquiryStatusSuccess:              "Inquiry process is complete and bank account number is valid",
		InquiryStatusInvalidAccountNumber: "Inquiry process is complete but the account number is invalid",
		InquiryStatusSuspectedAccount:     "Bank account have been suspected on doing fraud",
		InquiryStatusBlackListed:          "Bank account have been confirmed on doing a fraud and therefore is blacklisted",
		InquiryStatusFailed:               "The inquiry process is failed before we get the final status of the inquiry",
		InquiryStatusClosed:               "The account is valid, but it is closed/inactive so that it cannot receive money",
	}

	return inquiryStatus
}

func GetBankInquiryStatusLabel(inquiryStatus string) (string, error) {
	inquiryStatuses := GetAllBankInquiryStatusLabel()
	if value, ok := inquiryStatuses[inquiryStatus]; !ok {
		return "", fmt.Errorf("inquiry status [%s] not found", inquiryStatus)
	} else {
		return value, nil
	}
}

func GetSenderBank() map[string]string {
	getValue := func(code string, typeCode string) (value string) {
		if typeCode == BankTypeBankAccount || typeCode == BankTypeVirtualAccount {
			value, _ = GetLabelBank(code)
		} else if typeCode == BankTypeWalletAccount {
			value, _ = GetWalletLabel(code)
		}

		return value
	}

	senderBank := map[string]string{
		// Bank
		BNI:             getValue(BNI, BankTypeBankAccount),
		BRI:             getValue(BRI, BankTypeBankAccount),
		BCA:             getValue(BCA, BankTypeBankAccount),
		Mandiri:         getValue(Mandiri, BankTypeBankAccount),
		CIMB:            getValue(CIMB, BankTypeBankAccount),
		BTPN:            getValue(BTPN, BankTypeBankAccount),
		DBSIndo:         getValue(DBSIndo, BankTypeBankAccount),
		Permata:         getValue(Permata, BankTypeBankAccount),
		Muamalat:        getValue(Muamalat, BankTypeBankAccount),
		DanamonSyariah:  getValue(DanamonSyariah, BankTypeBankAccount),
		BankSyariahIndo: getValue(BankSyariahIndo, BankTypeBankAccount),

		// E-Wallet
		Ovo:          getValue(Ovo, BankTypeWalletAccount),
		Qris:         getValue(Qris, BankTypeWalletAccount),
		ShopeePayApp: getValue(ShopeePayApp, BankTypeWalletAccount),
		LinkAja:      getValue(LinkAja, BankTypeWalletAccount),
		LinkAjaApp:   getValue(LinkAjaApp, BankTypeWalletAccount),
		Dana:         getValue(Dana, BankTypeWalletAccount),
	}

	return senderBank
}
