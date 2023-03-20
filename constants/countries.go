package constants

const CountryCodeSingapore = "SGD"
const CountryCodeMalaysia = "MYS"
const CountryCodeAustralia = "AUS"
const CountryCodeJapan = "JPN"
const CountryCodeTurkey = "TUR"
const CountryCodeUnitedKingdom = "GBR"
const CountryCodeThailand = "THA"
const CountryCodeSouthKorea = "KOR"
const CountryCodeHongKong = "HKG"
const CountryCodeChina = "CHN"
const CountryCodeAndorra = "AND"
const CountryCodeAustria = "AUT"
const CountryCodeBelgium = "BEL"
const CountryCodeBulgaria = "BGR"
const CountryCodeCroatia = "HRV"
const CountryCodeCyprus = "CYP"
const CountryCodeCzechRepublic = "CZE"
const CountryCodeDenmark = "DNK"
const CountryCodeEstonia = "EST"
const CountryCodeFinland = "FIN"
const CountryCodeFrance = "FRA"
const CountryCodeGermany = "DEU"
const CountryCodeGreece = "GRC"
const CountryCodeHungary = "HUN"
const CountryCodeIceland = "ISL"
const CountryCodeIreland = "IRL"
const CountryCodeItaly = "ITA"
const CountryCodeLatvia = "LVA"
const CountryCodeLiechtenstein = "LIE"
const CountryCodeLithuania = "LTU"
const CountryCodeLuxembourg = "LUX"
const CountryCodeMalta = "MLT"
const CountryCodeMonaco = "MCO"
const CountryCodeNetherlands = "NLD"
const CountryCodeNorway = "NOR"
const CountryCodePoland = "POL"
const CountryCodePortugal = "PRT"
const CountryCodeRomania = "ROU"
const CountryCodeSanMarino = "SMR"
const CountryCodeSlovakia = "SVK"
const CountryCodeSlovenia = "SVN"
const CountryCodeSpain = "ESP"
const CountryCodeSweden = "SWE"
const CountryCodeSwitzerland = "CHE"
const CountryCodeVaticanCityState = "VAT"
const CountryCodeSaudiArabia = "SAU"
const CountryCodeUnitedArabEmirates = "ARE"

const CurrencyCodeEuro = "EUR"
const CurrencyCodeSingapore = "SGD"
const CurrencyCodeMalaysia = "MYR"
const CurrencyCodeAustralia = "AUD"
const CurrencyCodeJapan = "JPY"
const CurrencyCodeTurkey = "TRY"
const CurrencyCodeUnitedKingdom = "GBP"
const CurrencyCodeThailand = "THB"
const CurrencyCodeSouthKorea = "KRW"
const CurrencyCodeHongKongHKD = "HKD"
const CurrencyCodeHongKongCNY = CurrencyCodeChina
const CurrencyCodeChina = "CNY"
const CurrencyCodeAndorra = CurrencyCodeEuro
const CurrencyCodeAustria = CurrencyCodeEuro
const CurrencyCodeBelgium = CurrencyCodeEuro
const CurrencyCodeBulgaria = CurrencyCodeEuro
const CurrencyCodeCroatia = CurrencyCodeEuro
const CurrencyCodeCyprus = CurrencyCodeEuro
const CurrencyCodeCzechRepublic = CurrencyCodeEuro
const CurrencyCodeDenmark = CurrencyCodeEuro
const CurrencyCodeEstonia = CurrencyCodeEuro
const CurrencyCodeFinland = CurrencyCodeEuro
const CurrencyCodeFrance = CurrencyCodeEuro
const CurrencyCodeGermany = CurrencyCodeEuro
const CurrencyCodeGreece = CurrencyCodeEuro
const CurrencyCodeHungary = CurrencyCodeEuro
const CurrencyCodeIceland = CurrencyCodeEuro
const CurrencyCodeIreland = CurrencyCodeEuro
const CurrencyCodeItaly = CurrencyCodeEuro
const CurrencyCodeLatvia = CurrencyCodeEuro
const CurrencyCodeLiechtenstein = CurrencyCodeEuro
const CurrencyCodeLithuania = CurrencyCodeEuro
const CurrencyCodeLuxembourg = CurrencyCodeEuro
const CurrencyCodeMalta = CurrencyCodeEuro
const CurrencyCodeMonaco = CurrencyCodeEuro
const CurrencyCodeNetherlands = CurrencyCodeEuro
const CurrencyCodeNorway = CurrencyCodeEuro
const CurrencyCodePoland = CurrencyCodeEuro
const CurrencyCodePortugal = CurrencyCodeEuro
const CurrencyCodeRomania = CurrencyCodeEuro
const CurrencyCodeSanMarino = CurrencyCodeEuro
const CurrencyCodeSlovakia = CurrencyCodeEuro
const CurrencyCodeSlovenia = CurrencyCodeEuro
const CurrencyCodeSpain = CurrencyCodeEuro
const CurrencyCodeSweden = CurrencyCodeEuro
const CurrencyCodeSwitzerland = CurrencyCodeEuro
const CurrencyCodeVaticanCityState = CurrencyCodeEuro
const CurrencyCodeSaudiArabia = "SAR"
const CurrencyCodeUnitedArabEmirates = "AED"

type DestinationCountry struct {
	CountryName  string
	CountryCode  string
	CurrencyCode string
}

func GetAllDestinationCountry(countryCode string) map[string]DestinationCountry {
	countries := map[string]DestinationCountry{
		CountryCodeSingapore: {
			CountryName:  "Singapore",
			CountryCode:  CountryCodeSingapore,
			CurrencyCode: CurrencyCodeSingapore,
		},
		CountryCodeMalaysia: {
			CountryName:  "Malaysia",
			CountryCode:  CountryCodeMalaysia,
			CurrencyCode: CurrencyCodeMalaysia,
		},
		CountryCodeAustralia: {
			CountryName:  "Malaysia",
			CountryCode:  CountryCodeAustralia,
			CurrencyCode: CurrencyCodeAustralia,
		},
		CountryCodeJapan: {
			CountryName:  "Japan",
			CountryCode:  CountryCodeJapan,
			CurrencyCode: CurrencyCodeJapan,
		},
		CountryCodeTurkey: {
			CountryName:  "Turkey",
			CountryCode:  CountryCodeTurkey,
			CurrencyCode: CurrencyCodeTurkey,
		},
		CountryCodeUnitedKingdom: {
			CountryName:  "United Kingdom",
			CountryCode:  CountryCodeUnitedKingdom,
			CurrencyCode: CurrencyCodeUnitedKingdom,
		},
		CountryCodeThailand: {
			CountryName:  "Thailand",
			CountryCode:  CountryCodeThailand,
			CurrencyCode: CurrencyCodeThailand,
		},
		CountryCodeSouthKorea: {
			CountryName:  "South Korea",
			CountryCode:  CountryCodeSouthKorea,
			CurrencyCode: CurrencyCodeSouthKorea,
		},
		CountryCodeHongKong: {
			CountryName:  "HongKong",
			CountryCode:  CountryCodeHongKong,
			CurrencyCode: CurrencyCodeHongKongHKD,
		},
		CountryCodeChina: {
			CountryName:  "China",
			CountryCode:  CountryCodeChina,
			CurrencyCode: CurrencyCodeChina,
		},
		CountryCodeAndorra: {
			CountryName:  "Andorra",
			CountryCode:  CountryCodeAndorra,
			CurrencyCode: CurrencyCodeAndorra,
		},
		CountryCodeAustria: {
			CountryName:  "Austria",
			CountryCode:  CountryCodeAustria,
			CurrencyCode: CurrencyCodeAustria,
		},
		CountryCodeBelgium: {
			CountryName:  "Belgium",
			CountryCode:  CountryCodeBelgium,
			CurrencyCode: CurrencyCodeBelgium,
		},
		CountryCodeBulgaria: {
			CountryName:  "Bulgaria",
			CountryCode:  CountryCodeBulgaria,
			CurrencyCode: CurrencyCodeBulgaria,
		},
		CountryCodeCroatia: {
			CountryName:  "Croatia",
			CountryCode:  CountryCodeCroatia,
			CurrencyCode: CurrencyCodeCroatia,
		},
		CountryCodeCyprus: {
			CountryName:  "Cyprus",
			CountryCode:  CountryCodeCyprus,
			CurrencyCode: CurrencyCodeCyprus,
		},
		CountryCodeCzechRepublic: {
			CountryName:  "Czech Republic",
			CountryCode:  CountryCodeCzechRepublic,
			CurrencyCode: CurrencyCodeCzechRepublic,
		},
		CountryCodeDenmark: {
			CountryName:  "Denmark",
			CountryCode:  CountryCodeDenmark,
			CurrencyCode: CurrencyCodeDenmark,
		},
		CountryCodeEstonia: {
			CountryName:  "Estonia",
			CountryCode:  CountryCodeEstonia,
			CurrencyCode: CurrencyCodeEstonia,
		},
		CountryCodeFinland: {
			CountryName:  "Finland",
			CountryCode:  CountryCodeFinland,
			CurrencyCode: CurrencyCodeFinland,
		},
		CountryCodeFrance: {
			CountryName:  "France",
			CountryCode:  CountryCodeFrance,
			CurrencyCode: CurrencyCodeFrance,
		},
		CountryCodeGermany: {
			CountryName:  "Germany",
			CountryCode:  CountryCodeGermany,
			CurrencyCode: CurrencyCodeGermany,
		},
		CountryCodeGreece: {
			CountryName:  "Greece",
			CountryCode:  CountryCodeGreece,
			CurrencyCode: CurrencyCodeGreece,
		},
		CountryCodeHungary: {
			CountryName:  "Hungary",
			CountryCode:  CountryCodeHungary,
			CurrencyCode: CurrencyCodeHungary,
		},
		CountryCodeIceland: {
			CountryName:  "Iceland",
			CountryCode:  CountryCodeIceland,
			CurrencyCode: CurrencyCodeIceland,
		},
		CountryCodeIreland: {
			CountryName:  "Ireland",
			CountryCode:  CountryCodeIreland,
			CurrencyCode: CurrencyCodeIreland,
		},
		CountryCodeItaly: {
			CountryName:  "Italy",
			CountryCode:  CountryCodeItaly,
			CurrencyCode: CurrencyCodeItaly,
		},
		CountryCodeLatvia: {
			CountryName:  "Latvia",
			CountryCode:  CountryCodeLatvia,
			CurrencyCode: CurrencyCodeLatvia,
		},
		CountryCodeLiechtenstein: {
			CountryName:  "Liechtenstein",
			CountryCode:  CountryCodeLiechtenstein,
			CurrencyCode: CurrencyCodeLiechtenstein,
		},
		CountryCodeLithuania: {
			CountryName:  "Lithuania",
			CountryCode:  CountryCodeLithuania,
			CurrencyCode: CurrencyCodeLithuania,
		},
		CountryCodeLuxembourg: {
			CountryName:  "Luxembourg",
			CountryCode:  CountryCodeLuxembourg,
			CurrencyCode: CurrencyCodeLuxembourg,
		},
		CountryCodeMalta: {
			CountryName:  "Malta",
			CountryCode:  CountryCodeMalta,
			CurrencyCode: CurrencyCodeMalta,
		},
		CountryCodeMonaco: {
			CountryName:  "Monaco",
			CountryCode:  CountryCodeMonaco,
			CurrencyCode: CurrencyCodeMonaco,
		},
		CountryCodeNetherlands: {
			CountryName:  "Netherlands",
			CountryCode:  CountryCodeNetherlands,
			CurrencyCode: CurrencyCodeNetherlands,
		},
		CountryCodeNorway: {
			CountryName:  "Norway",
			CountryCode:  CountryCodeNorway,
			CurrencyCode: CurrencyCodeNorway,
		},
		CountryCodePoland: {
			CountryName:  "Poland",
			CountryCode:  CountryCodePoland,
			CurrencyCode: CurrencyCodePoland,
		},
		CountryCodePortugal: {
			CountryName:  "Portugal",
			CountryCode:  CountryCodePortugal,
			CurrencyCode: CurrencyCodePortugal,
		},
		CountryCodeRomania: {
			CountryName:  "Romania",
			CountryCode:  CountryCodeRomania,
			CurrencyCode: CurrencyCodeRomania,
		},
		CountryCodeSanMarino: {
			CountryName:  "San Marino",
			CountryCode:  CountryCodeSanMarino,
			CurrencyCode: CurrencyCodeSanMarino,
		},
		CountryCodeSlovakia: {
			CountryName:  "Slovakia",
			CountryCode:  CountryCodeSlovakia,
			CurrencyCode: CurrencyCodeSlovakia,
		},
		CountryCodeSlovenia: {
			CountryName:  "Slovenia",
			CountryCode:  CountryCodeSlovenia,
			CurrencyCode: CurrencyCodeSlovenia,
		},
		CountryCodeSpain: {
			CountryName:  "Spain",
			CountryCode:  CountryCodeSpain,
			CurrencyCode: CurrencyCodeSpain,
		},
		CountryCodeSweden: {
			CountryName:  "Sweden",
			CountryCode:  CountryCodeSweden,
			CurrencyCode: CurrencyCodeSweden,
		},
		CountryCodeSwitzerland: {
			CountryName:  "Switzerland",
			CountryCode:  CountryCodeSwitzerland,
			CurrencyCode: CurrencyCodeSwitzerland,
		},
		CountryCodeVaticanCityState: {
			CountryName:  "Vatican City State",
			CountryCode:  CountryCodeVaticanCityState,
			CurrencyCode: CurrencyCodeVaticanCityState,
		},
		CountryCodeSaudiArabia: {
			CountryName:  "Saudi Arabia",
			CountryCode:  CountryCodeSaudiArabia,
			CurrencyCode: CurrencyCodeSaudiArabia,
		},
		CountryCodeUnitedArabEmirates: {
			CountryName:  "United Arab Emirates",
			CountryCode:  CountryCodeUnitedArabEmirates,
			CurrencyCode: CurrencyCodeUnitedArabEmirates,
		},
	}

	if countryCode != "" {
		if country, ok := countries[countryCode]; !ok {
			return nil
		} else {
			return map[string]DestinationCountry{
				country.CountryCode: country,
			}
		}
	}

	return countries
}
