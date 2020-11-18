// Copyright (c) 2013-2020 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

package iso

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strings"
)

// ISO 4217 Currency codes (2016-12-17)
// https://en.wikipedia.org/wiki/ISO_4217
//
var (
	ISO_4217_CURRENCY_CODES []string = []string{
		"AED", // United Arab Emirates dirham
		"AFN", // Afghan afghani
		"ALL", // Albanian lek
		"AMD", // Armenian dram
		"ANG", // Netherlands Antillean guilder
		"AOA", // Angolan kwanza
		"ARS", // Argentine peso
		"AUD", // Australian dollar
		"AWG", // Aruban florin
		"AZN", // Azerbaijani manat
		"BAM", // Bosnia and Herzegovina convertible mark
		"BBD", // Barbados dollar
		"BDT", // Bangladeshi taka
		"BGN", // Bulgarian lev
		"BHD", // Bahraini dinar
		"BIF", // Burundian franc
		"BMD", // Bermudian dollar
		"BND", // Brunei dollar
		"BOB", // Boliviano
		"BOV", // Bolivian Mvdol
		"BRL", // Brazilian real
		"BSD", // Bahamian dollar
		"BTN", // Bhutanese ngultrum
		"BWP", // Botswana pula
		"BYN", // New Belarusian ruble
		"BYR", // Belarusian ruble
		"BZD", // Belize dollar
		"CAD", // Canadian dollar
		"CDF", // Congolese franc
		"CHE", // WIR Euro (complementary currency) Switzerland
		"CHF", // Swiss franc
		"CHW", // WIR Franc (complementary currency) Switzerland
		"CLF", // Unidad de Fomento
		"CLP", // Chilean peso
		"CNY", // Chinese yuan
		"COP", // Colombian peso
		"COU", // Unidad de Valor Real (UVR)
		"CRC", // Costa Rican colon
		"CUC", // Cuban convertible peso
		"CUP", // Cuban peso
		"CVE", // Cape Verde escudo
		"CZK", // Czech koruna
		"DJF", // Djiboutian franc
		"DKK", // Danish krone
		"DOP", // Dominican peso
		"DZD", // Algerian dinar
		"EGP", // Egyptian pound
		"ERN", // Eritrean nakfa
		"ETB", // Ethiopian birr
		"EUR", // Euro
		"FJD", // Fiji dollar
		"FKP", // Falkland Islands pound
		"GBP", // Pound sterling
		"GEL", // Georgian lari
		"GHS", // Ghanaian cedi
		"GIP", // Gibraltar pound
		"GMD", // Gambian dalasi
		"GNF", // Guinean franc
		"GTQ", // Guatemalan quetzal
		"GYD", // Guyanese dollar
		"HKD", // Hong Kong dollar
		"HNL", // Honduran lempira
		"HRK", // Croatian kuna
		"HTG", // Haitian gourde
		"HUF", // Hungarian forint
		"IDR", // Indonesian rupiah
		"ILS", // Israeli new shekel
		"INR", // Indian rupee
		"IQD", // Iraqi dinar
		"IRR", // Iranian rial
		"ISK", // Icelandic króna
		"JMD", // Jamaican dollar
		"JOD", // Jordanian dinar
		"JPY", // Japanese yen
		"KES", // Kenyan shilling
		"KGS", // Kyrgyzstani som
		"KHR", // Cambodian riel
		"KMF", // Comoro franc
		"KPW", // North Korean won
		"KRW", // South Korean won
		"KWD", // Kuwaiti dinar
		"KYD", // Cayman Islands dollar
		"KZT", // Kazakhstani tenge
		"LAK", // Lao kip
		"LBP", // Lebanese pound
		"LKR", // Sri Lankan rupee
		"LRD", // Liberian dollar
		"LSL", // Lesotho loti
		"LYD", // Libyan dinar
		"MAD", // Moroccan dirham
		"MDL", // Moldovan leu
		"MGA", // Malagasy ariary
		"MKD", // Macedonian denar
		"MMK", // Myanmar kyat
		"MNT", // Mongolian tögrög
		"MOP", // Macanese pataca
		"MRO", // Mauritanian ouguiya
		"MUR", // Mauritian rupee
		"MVR", // Maldivian rufiyaa
		"MWK", // Malawian kwacha
		"MXN", // Mexican peso
		"MXV", // Mexican Unidad de Inversion (UDI)
		"MYR", // Malaysian ringgit
		"MZN", // Mozambican metical
		"NAD", // Namibian dollar
		"NGN", // Nigerian naira
		"NIO", // Nicaraguan córdoba
		"NOK", // Norwegian krone
		"NPR", // Nepalese rupee
		"NZD", // New Zealand dollar
		"OMR", // Omani rial
		"PAB", // Panamanian balboa
		"PEN", // Peruvian Sol
		"PGK", // Papua New Guinean kina
		"PHP", // Philippine peso
		"PKR", // Pakistani rupee
		"PLN", // Polish złoty
		"PYG", // Paraguayan guaraní
		"QAR", // Qatari riyal
		"RON", // Romanian leu
		"RSD", // Serbian dinar
		"RUB", // Russian ruble
		"RWF", // Rwandan franc
		"SAR", // Saudi riyal
		"SBD", // Solomon Islands dollar
		"SCR", // Seychelles rupee
		"SDG", // Sudanese pound
		"SEK", // Swedish krona/kronor
		"SGD", // Singapore dollar
		"SHP", // Saint Helena pound
		"SLL", // Sierra Leonean leone
		"SOS", // Somali shilling
		"SRD", // Surinamese dollar
		"SSP", // South Sudanese pound
		"STD", // São Tomé and Príncipe dobra
		"SVC", // Salvadoran colón
		"SYP", // Syrian pound
		"SZL", // Swazi lilangeni
		"THB", // Thai baht
		"TJS", // Tajikistani somoni
		"TMT", // Turkmenistani manat
		"TND", // Tunisian dinar
		"TOP", // Tongan paʻanga
		"TRY", // Turkish lira
		"TTD", // Trinidad and Tobago dollar
		"TWD", // New Taiwan dollar
		"TZS", // Tanzanian shilling
		"UAH", // Ukrainian hryvnia
		"UGX", // Ugandan shilling
		"USD", // United States dollar
		"USN", // United States dollar (next day) (funds code)
		"UYI", // Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)
		"UYU", // Uruguayan peso
		"UZS", // Uzbekistan som
		"VEF", // Venezuelan bolívar
		"VND", // Vietnamese dong
		"VUV", // Vanuatu vatu
		"WST", // Samoan tala
		"XAF", // CFA franc BEAC
		"XAG", // Silver (one troy ounce)
		"XAU", // Gold (one troy ounce)
		"XBA", // European Composite Unit (EURCO) (bond market unit)
		"XBB", // European Monetary Unit (E.M.U.-6) (bond market unit)
		"XBC", // European Unit of Account 9 (E.U.A.-9) (bond market unit)
		"XBD", // European Unit of Account 17 (E.U.A.-17) (bond market unit)
		"XCD", // East Caribbean dollar
		"XDR", // Special drawing rights, International Monetary Fund
		"XOF", // CFA franc BCEAO
		"XPD", // Palladium (one troy ounce)
		"XPF", // CFP franc (franc Pacifique)
		"XPT", // Platinum (one troy ounce)
		"XSU", // SUCRE, Unified System for Regional Compensation
		"XTS", // Code reserved for testing purposes
		"XUA", // ADB Unit of Account
		"XXX", // No currency
		"YER", // Yemeni rial
		"ZAR", // South African rand
		"ZMW", // Zambian kwacha
		"ZWL", // Zimbabwean dollar A/10
	}
)

type currency struct {
	IsoNumeric         int
	Name               string
	Symbol             string
	SymbolFirst        bool
	AlternateSymbols   []string
	ThousandsSeparator string
	DecimalMark        string
	SubUnit            string
	SubUnitToUnit      int64
	SubUnitPrecision   int
	HTMLEntity         string
}

var currencies = map[string]currency{
	"AED": currency{784, "United Arab Emirates Dirham", "د.إ", true, []string{"DH", "Dhs"}, ",", ".", "Fils", 100, 2, ""},
	"AFN": currency{971, "Afghan Afghani", "؋", false, []string{"Af", "Afs"}, ",", ".", "Pul", 100, 2, ""},
	"ALL": currency{8, "Albanian Lek", "L", false, []string{"Lek"}, ",", ".", "Qintar", 100, 2, ""},
	"AMD": currency{51, "Armenian Dram", "դր.", false, []string{"dram"}, ",", ".", "Luma", 100, 2, ""},
	"ANG": currency{532, "Netherlands Antillean Gulden", "ƒ", true, []string{"NAƒ", "NAf", "f"}, ".", ",", "Cent", 100, 2, "&#x0192;"},
	"AOA": currency{973, "Angolan Kwanza", "Kz", false, []string{}, ",", ".", "Cêntimo", 100, 2, ""},
	"ARS": currency{32, "Argentine Peso", "$", true, []string{"$m/n", "m$n"}, ".", ",", "Centavo", 100, 2, "&#x20B1;"},
	"AUD": currency{36, "Australian Dollar", "$", true, []string{"A$"}, ",", ".", "Cent", 100, 2, "$"},
	"AWG": currency{533, "Aruban Florin", "ƒ", false, []string{"Afl"}, ",", ".", "Cent", 100, 2, "&#x0192;"},
	"AZN": currency{944, "Azerbaijani Manat", "₼", true, []string{"m", "man"}, ",", ".", "Qəpik", 100, 2, ""},
	"BAM": currency{977, "Bosnia and Herzegovina Convertible Mark", "КМ", true, []string{"KM"}, ",", ".", "Fening", 100, 2, ""},
	"BBD": currency{52, "Barbadian Dollar", "$", false, []string{"Bds$"}, ",", ".", "Cent", 100, 2, "$"},
	"BDT": currency{50, "Bangladeshi Taka", "৳", true, []string{"Tk"}, ",", ".", "Paisa", 100, 2, ""},
	"BGN": currency{975, "Bulgarian Lev", "лв", false, []string{"lev", "leva", "лев", "лева"}, ",", ".", "Stotinka", 100, 2, ""},
	"BHD": currency{48, "Bahraini Dinar", "ب.د", true, []string{"BD"}, ",", ".", "Fils", 1000, 2, ""},
	"BIF": currency{108, "Burundian Franc", "Fr", false, []string{"FBu"}, ",", ".", "Centime", 100, 2, ""},
	"BMD": currency{60, "Bermudian Dollar", "$", true, []string{"BD$"}, ",", ".", "Cent", 100, 2, "$"},
	"BND": currency{96, "Brunei Dollar", "$", true, []string{"B$"}, ",", ".", "Sen", 100, 2, "$"},
	"BOB": currency{68, "Bolivian Boliviano", "Bs.", true, []string{"Bs"}, ",", ".", "Centavo", 100, 2, ""},
	"BRL": currency{986, "Brazilian Real", "R$", true, []string{}, ".", ",", "Centavo", 100, 2, "R$"},
	"BSD": currency{44, "Bahamian Dollar", "$", true, []string{"B$"}, ",", ".", "Cent", 100, 2, "$"},
	"BTC": currency{0, "Bitcoin", "B⃦", true, []string{}, ",", ".", "Satoshi", 100000000, 8, ""},
	"BTN": currency{64, "Bhutanese Ngultrum", "Nu.", false, []string{"Nu"}, ",", ".", "Chertrum", 100, 2, ""},
	"BWP": currency{72, "Botswana Pula", "P", true, []string{}, ",", ".", "Thebe", 100, 2, ""},
	"BYR": currency{974, "Belarusian Ruble", "Br", false, []string{""}, ",", ".", "Kapyeyka", 100, 2, ""},
	"BZD": currency{84, "Belize Dollar", "$", true, []string{"BZ$"}, ",", ".", "Cent", 100, 2, "$"},
	"CAD": currency{124, "Canadian Dollar", "$", true, []string{"C$", "CAD$"}, ",", ".", "Cent", 100, 2, "$"},
	"CDF": currency{976, "Congolese Franc", "Fr", false, []string{"FC"}, ",", ".", "Centime", 100, 2, ""},
	"CHF": currency{756, "Swiss Franc", "Fr", true, []string{"SFr", "CHF"}, ",", ".", "Rappen", 100, 2, ""},
	"CLF": currency{990, "Unidad de Fomento", "UF", true, []string{}, ".", ",", "Peso", 1, 0, "&#x20B1;"},
	"CLP": currency{152, "Chilean Peso", "$", true, []string{}, ".", ",", "Peso", 100, 2, "&#36;"},
	"CNY": currency{156, "Chinese Renminbi Yuan", "¥", true, []string{"CN¥", "元", "CN元"}, ",", ".", "Fen", 100, 2, "￥"},
	"COP": currency{170, "Colombian Peso", "$", true, []string{"COL$"}, ".", ",", "Centavo", 100, 2, "&#x20B1;"},
	"CRC": currency{188, "Costa Rican Colón", "₡", true, []string{"¢"}, ".", ",", "Céntimo", 100, 2, "&#x20A1;"},
	"CUC": currency{931, "Cuban Convertible Peso", "$", false, []string{"CUC$"}, ",", ".", "Centavo", 100, 2, ""},
	"CUP": currency{192, "Cuban Peso", "$", true, []string{"$MN"}, ",", ".", "Centavo", 100, 2, "&#x20B1;"},
	"CVE": currency{132, "Cape Verdean Escudo", "$", false, []string{"Esc"}, ",", ".", "Centavo", 100, 2, ""},
	"CZK": currency{203, "Czech Koruna", "Kč", false, []string{}, ".", ",", "Haléř", 100, 2, ""},
	"DJF": currency{262, "Djiboutian Franc", "Fdj", false, []string{}, ",", ".", "Centime", 100, 2, ""},
	"DKK": currency{208, "Danish Krone", "kr", false, []string{",-"}, ".", ",", "Øre", 100, 2, ""},
	"DOP": currency{214, "Dominican Peso", "$", true, []string{"RD$"}, ",", ".", "Centavo", 100, 2, "&#x20B1;"},
	"DZD": currency{12, "Algerian Dinar", "د.ج", false, []string{"DA"}, ",", ".", "Centime", 100, 2, ""},
	"EEK": currency{233, "Estonian Kroon", "KR", false, []string{}, ",", ".", "Sent", 100, 2, ""},
	"EGP": currency{818, "Egyptian Pound", "ج.م", true, []string{"LE", "E£", "L.E."}, ",", ".", "Piastre", 100, 2, "&#x00A3;"},
	"ERN": currency{232, "Eritrean Nakfa", "Nfk", false, []string{}, ",", ".", "Cent", 100, 2, ""},
	"ETB": currency{230, "Ethiopian Birr", "Br", false, []string{}, ",", ".", "Santim", 100, 2, ""},
	"EUR": currency{978, "Euro", "€", false, []string{}, ".", ",", "Cent", 100, 2, "&#x20AC;"},
	"FJD": currency{242, "Fijian Dollar", "$", false, []string{"FJ$"}, ",", ".", "Cent", 100, 2, "$"},
	"FKP": currency{238, "Falkland Pound", "£", false, []string{"FK£"}, ",", ".", "Penny", 100, 2, "&#x00A3;"},
	"GBP": currency{826, "British Pound", "£", true, []string{}, ",", ".", "Penny", 100, 2, "&#x00A3;"},
	"GEL": currency{981, "Georgian Lari", "ლ", false, []string{"lari"}, ",", ".", "Tetri", 100, 2, ""},
	"GHS": currency{936, "Ghanaian Cedi", "₵", true, []string{"GH¢", "GH₵"}, ",", ".", "Pesewa", 100, 2, "&#x20B5;"},
	"GIP": currency{292, "Gibraltar Pound", "£", true, []string{}, ",", ".", "Penny", 100, 2, "&#x00A3;"},
	"GMD": currency{270, "Gambian Dalasi", "D", false, []string{}, ",", ".", "Butut", 100, 2, ""},
	"GNF": currency{324, "Guinean Franc", "Fr", false, []string{"FG", "GFr"}, ",", ".", "Centime", 100, 2, ""},
	"GTQ": currency{320, "Guatemalan Quetzal", "Q", true, []string{}, ",", ".", "Centavo", 100, 2, ""},
	"GYD": currency{328, "Guyanese Dollar", "$", false, []string{"G$"}, ",", ".", "Cent", 100, 2, "$"},
	"HKD": currency{344, "Hong Kong Dollar", "$", true, []string{"HK$"}, ",", ".", "Cent", 100, 2, "$"},
	"HNL": currency{340, "Honduran Lempira", "L", true, []string{}, ",", ".", "Centavo", 100, 2, ""},
	"HRK": currency{191, "Croatian Kuna", "kn", true, []string{}, ".", ",", "Lipa", 100, 2, ""},
	"HTG": currency{332, "Haitian Gourde", "G", false, []string{}, ",", ".", "Centime", 100, 2, ""},
	"HUF": currency{348, "Hungarian Forint", "Ft", false, []string{}, ".", ",", "Fillér", 100, 2, ""},
	"IDR": currency{360, "Indonesian Rupiah", "Rp", true, []string{}, ".", ",", "Sen", 100, 2, ""},
	"ILS": currency{376, "Israeli New Sheqel", "₪", true, []string{"ש״ח", "NIS"}, ",", ".", "Agora", 100, 2, "&#x20AA;"},
	"INR": currency{356, "Indian Rupee", "₹", true, []string{"Rs", "৳", "૱", "௹", "रु", "₨"}, ",", ".", "Paisa", 100, 2, "&#x20b9;"},
	"IQD": currency{368, "Iraqi Dinar", "ع.د", false, []string{}, ",", ".", "Fils", 1000, 3, ""},
	"IRR": currency{364, "Iranian Rial", "﷼", true, []string{}, ",", ".", "Dinar", 100, 2, "&#xFDFC;"},
	"ISK": currency{352, "Icelandic Króna", "kr", true, []string{"Íkr"}, ".", ",", "Eyrir", 100, 2, ""},
	"JEP": currency{0, "Jersey Pound", "£", true, []string{}, ",", ".", "Penny", 100, 2, "&#x00A3;"},
	"JMD": currency{388, "Jamaican Dollar", "$", true, []string{"J$"}, ",", ".", "Cent", 100, 2, "$"},
	"JOD": currency{400, "Jordanian Dinar", "د.ا", true, []string{"JD"}, ",", ".", "Piastre", 100, 2, ""},
	"JPY": currency{392, "Japanese Yen", "¥", true, []string{"円", "圓"}, ",", ".", "", 1, 0, "&#x00A5;"},
	"KES": currency{404, "Kenyan Shilling", "KSh", true, []string{"Sh"}, ",", ".", "Cent", 100, 2, ""},
	"KGS": currency{417, "Kyrgyzstani Som", "som", false, []string{"сом"}, ",", ".", "Tyiyn", 100, 2, ""},
	"KHR": currency{116, "Cambodian Riel", "៛", false, []string{}, ",", ".", "Sen", 100, 2, "&#x17DB;"},
	"KMF": currency{174, "Comorian Franc", "Fr", false, []string{"CF"}, ",", ".", "Centime", 100, 2, ""},
	"KPW": currency{408, "North Korean Won", "₩", false, []string{}, ",", ".", "Chŏn", 100, 2, "&#x20A9;"},
	"KRW": currency{410, "South Korean Won", "₩", true, []string{}, ",", ".", "", 100, 2, "&#x20A9;"},
	"KWD": currency{414, "Kuwaiti Dinar", "د.ك", true, []string{"K.D."}, ",", ".", "Fils", 1000, 3, ""},
	"KYD": currency{136, "Cayman Islands Dollar", "$", true, []string{"CI$"}, ",", ".", "Cent", 100, 2, "$"},
	"KZT": currency{398, "Kazakhstani Tenge", "〒", false, []string{}, ",", ".", "Tiyn", 100, 2, ""},
	"LAK": currency{418, "Lao Kip", "₭", false, []string{"₭N"}, ",", ".", "Att", 100, 2, "&#x20AD;"},
	"LBP": currency{422, "Lebanese Pound", "ل.ل", true, []string{"£", "L£"}, ",", ".", "Piastre", 100, 2, "&#x00A3;"},
	"LKR": currency{144, "Sri Lankan Rupee", "₨", false, []string{"රු", "ரூ", "SLRs", "/-"}, ",", ".", "Cent", 100, 2, "&#x0BF9;"},
	"LRD": currency{430, "Liberian Dollar", "$", false, []string{"L$"}, ",", ".", "Cent", 100, 2, "$"},
	"LSL": currency{426, "Lesotho Loti", "L", false, []string{"M"}, ",", ".", "Sente", 100, 2, ""},
	"LTL": currency{440, "Lithuanian Litas", "Lt", false, []string{}, ",", ".", "Centas", 100, 2, ""},
	"LVL": currency{428, "Latvian Lats", "Ls", true, []string{}, ",", ".", "Santīms", 100, 2, ""},
	"LYD": currency{434, "Libyan Dinar", "ل.د", false, []string{"LD"}, ",", ".", "Dirham", 1000, 3, ""},
	"MAD": currency{504, "Moroccan Dirham", "د.م.", false, []string{}, ",", ".", "Centime", 100, 2, ""},
	"MDL": currency{498, "Moldovan Leu", "L", false, []string{"lei"}, ",", ".", "Ban", 100, 2, ""},
	"MGA": currency{969, "Malagasy Ariary", "Ar", true, []string{}, ",", ".", "Iraimbilanja", 5, 2, ""},
	"MKD": currency{807, "Macedonian Denar", "ден", false, []string{}, ",", ".", "Deni", 100, 2, ""},
	"MMK": currency{104, "Myanmar Kyat", "K", false, []string{}, ",", ".", "Pya", 100, 2, ""},
	"MNT": currency{496, "Mongolian Tögrög", "₮", false, []string{}, ",", ".", "Möngö", 100, 2, "&#x20AE;"},
	"MOP": currency{446, "Macanese Pataca", "P", false, []string{"MOP$"}, ",", ".", "Avo", 100, 2, ""},
	"MRO": currency{478, "Mauritanian Ouguiya", "UM", false, []string{}, ",", ".", "Khoums", 5, 2, ""},
	"MTL": currency{470, "Maltese Lira", "₤", true, []string{"Lm"}, ",", ".", "Cent", 100, 2, "&#x00A3;"},
	"MUR": currency{480, "Mauritian Rupee", "₨", true, []string{}, ",", ".", "Cent", 100, 2, "&#x20A8;"},
	"MVR": currency{462, "Maldivian Rufiyaa", "MVR", false, []string{"MRF", "Rf", "/-", "ރ"}, ",", ".", "Laari", 100, 2, ""},
	"MWK": currency{454, "Malawian Kwacha", "MK", false, []string{}, ",", ".", "Tambala", 100, 2, ""},
	"MXN": currency{484, "Mexican Peso", "$", true, []string{"MEX$"}, ",", ".", "Centavo", 100, 2, "$"},
	"MYR": currency{458, "Malaysian Ringgit", "RM", true, []string{}, ",", ".", "Sen", 100, 2, ""},
	"MZN": currency{943, "Mozambican Metical", "MTn", true, []string{"MZN"}, ".", ",", "Centavo", 100, 2, ""},
	"NAD": currency{516, "Namibian Dollar", "$", false, []string{"N$"}, ",", ".", "Cent", 100, 2, "$"},
	"NGN": currency{566, "Nigerian Naira", "₦", true, []string{}, ",", ".", "Kobo", 100, 2, "&#x20A6;"},
	"NIO": currency{558, "Nicaraguan Córdoba", "C$", false, []string{}, ",", ".", "Centavo", 100, 2, ""},
	"NOK": currency{578, "Norwegian Krone", "kr", false, []string{",-"}, ".", ",", "Øre", 100, 2, "kr"},
	"NPR": currency{524, "Nepalese Rupee", "₨", true, []string{"Rs", "रू"}, ",", ".", "Paisa", 100, 2, "&#x20A8;"},
	"NZD": currency{554, "New Zealand Dollar", "$", true, []string{"NZ$"}, ",", ".", "Cent", 100, 2, "$"},
	"OMR": currency{512, "Omani Rial", "ر.ع.", true, []string{}, ",", ".", "Baisa", 1000, 3, "&#xFDFC;"},
	"PAB": currency{590, "Panamanian Balboa", "B/.", false, []string{}, ",", ".", "Centésimo", 100, 2, ""},
	"PEN": currency{604, "Peruvian Nuevo Sol", "S/.", true, []string{}, ",", ".", "Céntimo", 100, 2, "S/."},
	"PGK": currency{598, "Papua New Guinean Kina", "K", false, []string{}, ",", ".", "Toea", 100, 2, ""},
	"PHP": currency{608, "Philippine Peso", "₱", true, []string{"PHP", "PhP", "P"}, ",", ".", "Centavo", 100, 2, "&#x20B1;"},
	"PKR": currency{586, "Pakistani Rupee", "₨", true, []string{"Rs"}, ",", ".", "Paisa", 100, 2, "&#x20A8;"},
	"PLN": currency{985, "Polish Złoty", "zł", false, []string{}, " ", ",", "Grosz", 100, 2, "&#322;"},
	"PYG": currency{600, "Paraguayan Guaraní", "₲", true, []string{}, ",", ".", "Céntimo", 100, 2, "&#x20B2;"},
	"QAR": currency{634, "Qatari Riyal", "ر.ق", false, []string{"QR"}, ",", ".", "Dirham", 100, 2, "&#xFDFC;"},
	"RON": currency{946, "Romanian Leu", "Lei", true, []string{}, ".", ",", "Bani", 100, 2, ""},
	"RSD": currency{941, "Serbian Dinar", "РСД", true, []string{"RSD", "din", "дин"}, ",", ".", "Para", 100, 2, ""},
	"RUB": currency{643, "Russian Ruble", "₽", false, []string{"руб.", "р."}, ".", ",", "Kopeck", 100, 2, "&#x20BD;"},
	"RWF": currency{646, "Rwandan Franc", "FRw", false, []string{"RF", "R₣"}, ",", ".", "Centime", 100, 2, ""},
	"SAR": currency{682, "Saudi Riyal", "ر.س", true, []string{"SR", "﷼"}, ",", ".", "Hallallah", 100, 2, "&#xFDFC;"},
	"SBD": currency{90, "Solomon Islands Dollar", "$", false, []string{"SI$"}, ",", ".", "Cent", 100, 2, "$"},
	"SCR": currency{690, "Seychellois Rupee", "₨", false, []string{"SRe", "SR"}, ",", ".", "Cent", 100, 2, "&#x20A8;"},
	"SDG": currency{938, "Sudanese Pound", "£", true, []string{}, ",", ".", "Piastre", 100, 2, ""},
	"SEK": currency{752, "Swedish Krona", "kr", false, []string{":-"}, " ", ",", "Öre", 100, 2, ""},
	"SGD": currency{702, "Singapore Dollar", "$", true, []string{"S$"}, ",", ".", "Cent", 100, 2, "$"},
	"SHP": currency{654, "Saint Helenian Pound", "£", false, []string{}, ",", ".", "Penny", 100, 2, "&#x00A3;"},
	"SKK": currency{703, "Slovak Koruna", "Sk", true, []string{}, ",", ".", "Halier", 100, 2, ""},
	"SLL": currency{694, "Sierra Leonean Leone", "Le", false, []string{}, ",", ".", "Cent", 100, 2, ""},
	"SOS": currency{706, "Somali Shilling", "Sh", false, []string{"Sh.So"}, ",", ".", "Cent", 100, 2, ""},
	"SRD": currency{968, "Surinamese Dollar", "$", false, []string{}, ",", ".", "Cent", 100, 2, ""},
	"SSP": currency{728, "South Sudanese Pound", "£", false, []string{}, ",", ".", "piaster", 100, 2, "&#x00A3;"},
	"STD": currency{678, "São Tomé and Príncipe Dobra", "Db", false, []string{}, ",", ".", "Cêntimo", 100, 2, ""},
	"SVC": currency{222, "Salvadoran Colón", "₡", true, []string{"¢"}, ",", ".", "Centavo", 100, 2, "&#x20A1;"},
	"SYP": currency{760, "Syrian Pound", "£S", false, []string{"£", "ل.س", "LS", "الليرة السورية"}, ",", ".", "Piastre", 100, 2, "&#x00A3;"},
	"SZL": currency{748, "Swazi Lilangeni", "L", true, []string{"E"}, ",", ".", "Cent", 100, 2, ""},
	"THB": currency{764, "Thai Baht", "฿", true, []string{}, ",", ".", "Satang", 100, 2, "&#x0E3F;"},
	"TJS": currency{972, "Tajikistani Somoni", "ЅМ", false, []string{}, ",", ".", "Diram", 100, 2, ""},
	"TMT": currency{934, "Turkmenistani Manat", "T", false, []string{}, ",", ".", "Tenge", 100, 2, ""},
	"TND": currency{788, "Tunisian Dinar", "د.ت", false, []string{"TD", "DT"}, ",", ".", "Millime", 1000, 3, ""},
	"TOP": currency{776, "Tongan Paʻanga", "T$", true, []string{"PT"}, ",", ".", "Seniti", 100, 2, ""},
	"TRY": currency{949, "Turkish Lira", "₺", false, []string{"TL"}, ".", ",", "kuruş", 100, 2, ""},
	"TTD": currency{780, "Trinidad and Tobago Dollar", "$", false, []string{"TT$"}, ",", ".", "Cent", 100, 2, "$"},
	"TWD": currency{901, "New Taiwan Dollar", "$", true, []string{"NT$"}, ",", ".", "Cent", 100, 2, "$"},
	"TZS": currency{834, "Tanzanian Shilling", "Sh", true, []string{}, ",", ".", "Cent", 100, 2, ""},
	"UAH": currency{980, "Ukrainian Hryvnia", "₴", false, []string{}, ",", ".", "Kopiyka", 100, 2, "&#x20B4;"},
	"UGX": currency{800, "Ugandan Shilling", "USh", false, []string{}, ",", ".", "Cent", 100, 2, ""},
	"USD": currency{840, "United States Dollar", "$", true, []string{"US$"}, ",", ".", "Cent", 100, 2, "$"},
	"UYU": currency{858, "Uruguayan Peso", "$", true, []string{"$U"}, ".", ",", "Centésimo", 100, 2, "&#x20B1;"},
	"UZS": currency{860, "Uzbekistani Som", "", false, []string{}, ",", ".", "Tiyin", 100, 2, ""},
	"VEF": currency{937, "Venezuelan Bolívar", "Bs F", true, []string{"Bs.F", "Bs"}, ".", ",", "Céntimo", 100, 2, ""},
	"VND": currency{704, "Vietnamese Đồng", "₫", true, []string{}, ".", ",", "Hào", 1, 0, "&#x20AB;"},
	"VUV": currency{548, "Vanuatu Vatu", "Vt", true, []string{}, ",", ".", "", 1, 0, ""},
	"WST": currency{882, "Samoan Tala", "T", false, []string{"WS$", "SAT", "ST"}, ",", ".", "Sene", 100, 2, ""},
	"XAF": currency{950, "Central African Cfa Franc", "Fr", false, []string{"FCFA"}, ",", ".", "Centime", 100, 2, ""},
	"XAG": currency{961, "Silver (Troy Ounce)", "oz t", false, []string{}, ",", ".", "oz", 1, 0, ""},
	"XAU": currency{959, "Gold (Troy Ounce)", "oz t", false, []string{}, ",", ".", "oz", 1, 0, ""},
	"XCD": currency{951, "East Caribbean Dollar", "$", true, []string{"EC$"}, ",", ".", "Cent", 100, 2, "$"},
	"XDR": currency{960, "Special Drawing Rights", "SDR", false, []string{"XDR"}, ",", ".", "", 1, 0, "$"},
	"XOF": currency{952, "West African Cfa Franc", "Fr", false, []string{"CFA"}, ",", ".", "Centime", 100, 2, ""},
	"XPF": currency{953, "Cfp Franc", "Fr", false, []string{"F"}, ",", ".", "Centime", 100, 2, ""},
	"YER": currency{886, "Yemeni Rial", "﷼", false, []string{}, ",", ".", "Fils", 100, 2, "&#xFDFC;"},
	"ZAR": currency{710, "South African Rand", "R", true, []string{}, ",", ".", "Cent", 100, 2, "&#x0052;"},
	"ZMK": currency{894, "Zambian Kwacha", "ZK", false, []string{}, ",", ".", "Ngwee", 100, 2, ""},
	"ZMW": currency{967, "Zambian Kwacha", "ZK", false, []string{}, ",", ".", "Ngwee", 100, 2, ""},
	"ZWD": currency{716, "Zimbabwean Dollar", "$", true, []string{"Z$"}, ",", ".", "Cent", 100, 2, "$"},
	"ZWL": currency{932, "Zimbabwean Dollar", "$", true, []string{"Z$"}, ",", ".", "Cent", 100, 2, "$"},
	"ZWN": currency{942, "Zimbabwean Dollar", "$", true, []string{"Z$"}, ",", ".", "Cent", 100, 2, "$"},
	"ZWR": currency{935, "Zimbabwean Dollar", "$", true, []string{"Z$"}, ",", ".", "Cent", 100, 2, "$"},
}

type Currency string

const (
	CurrencyUndefined Currency = ""
)

func ParseCurrency(c string) Currency {
	c = strings.ToUpper(c)
	for _, x := range ISO_4217_CURRENCY_CODES {
		if x == c {
			return Currency(c)
		}
	}
	return CurrencyUndefined
}

func (r Currency) IsValid() bool {
	return r != CurrencyUndefined
}

// Text/JSON conversion
func (r Currency) MarshalText() ([]byte, error) {
	return []byte(r), nil
}

func (r *Currency) UnmarshalText(data []byte) error {
	rr := ParseCurrency(string(data))
	if !rr.IsValid() {
		return fmt.Errorf("iso: invalid ISO currency code '%s'", string(data))
	}
	*r = rr
	return nil
}

// SQL conversion
func (r *Currency) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*r = ParseCurrency(v)
	case []byte:
		*r = ParseCurrency(string(v))
	}
	if !(*r).IsValid() {
		return fmt.Errorf("iso: invalid ISO currency code '%v'", value)
	}
	return nil
}

func (r Currency) Value() (driver.Value, error) {
	return string(r), nil
}

func (c Currency) Symbol() string {
	if cc, ok := currencies[string(c)]; ok {
		return cc.Symbol
	}
	return string(c)
}

type CurrencyOptions struct {
	WithCents              bool //  true,
	WithCurrency           bool //  false,
	WithSymbol             bool //  true,
	WithSymbolSpace        bool //  false,
	WithThousandsSeparator bool //  true,
}

func NewCurrencyOptions() *CurrencyOptions {
	return &CurrencyOptions{
		WithCents:              true,
		WithCurrency:           false,
		WithSymbol:             true,
		WithSymbolSpace:        false,
		WithThousandsSeparator: true,
	}
}

func (o *CurrencyOptions) Cents(f bool) *CurrencyOptions {
	o.WithCents = f
	return o
}

func (o *CurrencyOptions) Currency(f bool) *CurrencyOptions {
	o.WithCurrency = f
	return o
}

func (o *CurrencyOptions) Symbol(f bool) *CurrencyOptions {
	o.WithSymbol = f
	return o
}

func (o *CurrencyOptions) Space(f bool) *CurrencyOptions {
	o.WithSymbolSpace = f
	return o
}

func (o *CurrencyOptions) Separator(f bool) *CurrencyOptions {
	o.WithThousandsSeparator = f
	return o
}

// Format returns a formatted price string according to currency rules and options
func (c Currency) Format(val float64, opts *CurrencyOptions) (result string) {
	if opts == nil {
		opts = NewCurrencyOptions()
	}

	cc, ok := currencies[string(c)]
	if !ok {
		cc = currency{
			IsoNumeric:         0,
			Name:               string(c),
			Symbol:             string(c),
			SymbolFirst:        false,
			AlternateSymbols:   nil,
			ThousandsSeparator: ",",
			DecimalMark:        ".",
			SubUnit:            "",
			SubUnitToUnit:      100,
			SubUnitPrecision:   2,
			HTMLEntity:         "",
		}
	}

	sign := ""
	if math.Signbit(val) {
		sign = "-"
	}
	integer, fractional := cc.splitValue(math.Abs(val))

	if opts.WithThousandsSeparator {
		result = cc.separateThousands(integer)
	} else {
		result = integer
	}

	if opts.WithCents && cc.SubUnit != "" {
		result = fmt.Sprintf("%s%s%s%s", sign, result, cc.DecimalMark, fractional)
	}

	if opts.WithSymbol {
		result = cc.addSymbol(result, opts)
	}

	if opts.WithCurrency {
		result = fmt.Sprintf("%s %s", result, string(c))
	}

	return result
}

func (c currency) addSymbol(result string, opts *CurrencyOptions) string {
	var space string

	if opts.WithSymbolSpace {
		space = " "
	}

	if c.SymbolFirst {
		result = fmt.Sprintf("%s%s%s", c.Symbol, space, result)
	} else {
		result = fmt.Sprintf("%s%s%s", result, space, c.Symbol)
	}

	return result
}

func (c currency) separateThousands(value string) string {
	chunks := len(value) / 3

	if chunks == 0 {
		return value
	}

	if partial := math.Mod(float64(len(value)), 3); partial > 0 {
		chunks++
	}

	result := make([]string, chunks)

	for i := chunks - 1; i >= 0; i-- {
		if i == 0 {
			result[i] = value
			break
		}

		chunk := value[len(value)-3:]
		value = strings.TrimSuffix(value, chunk)
		result[i] = chunk
	}

	return strings.Join(result, c.ThousandsSeparator)
}

func (c currency) splitValue(val float64) (integer, fractional string) {
	i, f := math.Modf(val)

	integer = fmt.Sprintf("%.0f", i)
	if c.SubUnitPrecision > 0 {
		fractional = fmt.Sprintf("%.[2]*[1]f", f, c.SubUnitPrecision)[2:]
	}

	return
}
