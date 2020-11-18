// Copyright (c) 2013-2020 KIDTSUNAMI
// Author: alex@kidtsunami.com
//

// TODO: EIDR language codes
//
// language[-extended_language][-script][-region][-variant]
//  2-3         2-3                 4        3        <=8

package iso

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

// ISO 639-1:2002 Language codes
// https://en.wikipedia.org/wiki/ISO_639-1
var (
	ISO_639_1_2002_CODES []string = []string{
		"aa", // Afar
		"ab", // Abkhazian
		"ae", // Avestan
		"af", // Afrikaans
		"ak", // Akan
		"am", // Amharic
		"an", // Aragonese
		"ar", // Arabic
		"as", // Assamese
		"av", // Avaric
		"ay", // Aymara
		"az", // Azerbaijani
		"ba", // Bashkir
		"be", // Belarusian
		"bg", // Bulgarian
		"bh", // Bihari languages
		"bi", // Bislama
		"bm", // Bambara
		"bn", // Bengali
		"bo", // Tibetan
		"br", // Breton
		"bs", // Bosnian
		"ca", // Catalan; Valencian
		"ce", // Chechen
		"ch", // Chamorro
		"co", // Corsican
		"cr", // Cree
		"cs", // Czech
		"cu", // Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic
		"cv", // Chuvash
		"cy", // Welsh
		"da", // Danish
		"de", // German
		"dv", // Divehi; Dhivehi; Maldivian
		"dz", // Dzongkha
		"ee", // Ewe
		"el", // Greek, Modern (1453-)
		"en", // English
		"eo", // Esperanto
		"es", // Spanish; Castilian
		"et", // Estonian
		"eu", // Basque
		"fa", // Persian
		"ff", // Fulah
		"fi", // Finnish
		"fj", // Fijian
		"fo", // Faroese
		"fr", // French
		"fy", // Western Frisian
		"ga", // Irish
		"gd", // Gaelic; Scottish
		"gl", // Galician
		"gn", // Guarani
		"gu", // Gujarati
		"gv", // Manx
		"ha", // Hausa
		"he", // Hebrew
		"hi", // Hindi
		"ho", // Hiri Motu
		"hr", // Croatian
		"ht", // Haitian; Haitian Creole
		"hu", // Hungarian
		"hy", // Armenian
		"hz", // Herero
		"ia", // Interlingua (International Auxiliary Language Association)
		"id", // Indonesian
		"ie", // Interlingue; Occidental
		"ig", // Igbo
		"ii", // Sichuan Yi; Nuosu
		"ik", // Inupiaq
		"io", // Ido
		"is", // Icelandic
		"it", // Italian
		"iu", // Inuktitut
		"ja", // Japanese
		"jv", // Javanese
		"ka", // Georgian
		"kg", // Kongo
		"ki", // Kikuyu; Gikuyu
		"kj", // Kuanyama; Kwanyama
		"kk", // Kazakh
		"kl", // Kalaallisut; Greenlandic
		"km", // Central
		"kn", // Kannada
		"ko", // Korean
		"kr", // Kanuri
		"ks", // Kashmiri
		"ku", // Kurdish
		"kv", // Komi
		"kw", // Cornish
		"ky", // Kirghiz; Kyrgyz
		"la", // Latin
		"lb", // Luxembourgish; Letzeburgesch
		"lg", // Ganda
		"li", // Limburgan; Limburger; Limburgish
		"ln", // Lingala
		"lo", // Lao
		"lt", // Lithuanian
		"lu", // Luba-Katanga
		"lv", // Latvian
		"mg", // Malagasy
		"mh", // Marshallese
		"mi", // Maori
		"mk", // Macedonian
		"ml", // Malayalam
		"mn", // Mongolian
		"mr", // Marathi
		"ms", // Malay
		"mt", // Maltese
		"my", // Burmese
		"na", // Nauru
		"nb", // Bokmål, Norwegian; Norwegian Bokmål
		"nd", // Ndebele, North; North Ndebele
		"ne", // Nepali
		"ng", // Ndonga
		"nl", // Dutch; Flemish
		"nn", // Norwegian Nynorsk; Nynorsk, Norwegian
		"no", // Norwegian
		"nr", // Ndebele, South; South Ndebele
		"nv", // Navajo; Navaho
		"ny", // Chichewa; Chewa; Nyanja
		"oc", // Occitan (post 1500)
		"oj", // Ojibwa
		"om", // Oromo
		"or", // Oriya
		"os", // Ossetian; Ossetic
		"pa", // Panjabi; Punjabi
		"pi", // Pali
		"pl", // Polish
		"ps", // Pushto; Pashto
		"pt", // Portuguese
		"qu", // Quechua
		"rm", // Romansh
		"rn", // Rundi
		"ro", // Romanian; Moldavian; Moldovan
		"ru", // Russian
		"rw", // Kinyarwanda
		"sa", // Sanskrit
		"sc", // Sardinian
		"sd", // Sindhi
		"se", // Northern Sami
		"sg", // Sango
		"si", // Sinhala; Sinhalese
		"sk", // Slovak
		"sl", // Slovenian
		"sm", // Samoan
		"sn", // Shona
		"so", // Somali
		"sq", // Albanian
		"sr", // Serbian
		"ss", // Swati
		"st", // Sotho, Southern
		"su", // Sundanese
		"sv", // Swedish
		"sw", // Swahili
		"ta", // Tamil
		"te", // Telugu
		"tg", // Tajik
		"th", // Thai
		"ti", // Tigrinya
		"tk", // Turkmen
		"tl", // Tagalog
		"tn", // Tswana
		"to", // Tonga (Tonga Islands)
		"tr", // Turkish
		"ts", // Tsonga
		"tt", // Tatar
		"tw", // Twi
		"ty", // Tahitian
		"ug", // Uighur; Uyghur
		"uk", // Ukrainian
		"ur", // Urdu
		"uz", // Uzbek
		"ve", // Venda
		"vi", // Vietnamese
		"vo", // Volapük
		"wa", // Walloon
		"wo", // Wolof
		"xh", // Xhosa
		"yi", // Yiddish
		"yo", // Yoruba
		"za", // Zhuang; Chuang
		"zh", // Chinese
		"zu", // Zulu
	}
)

// ISO 639-2/B-1998 Language codes (updated 2014-03-18)
// https://www.loc.gov/standards/iso639-2/php/code_list.php
// https://en.wikipedia.org/wiki/ISO_639-2
var (
	ISO_639_2B_1998_CODES []string = []string{
		"alb", // Albanian
		"arm", // Armenian
		"baq", // Basque
		"bur", // Burmese
		"chi", // Chinese
		"cze", // Czech
		"dut", // Dutch; Flemish
		"fre", // French
		"geo", // Georgian
		"ger", // German
		"gre", // Modern Greek (1453–)
		"ice", // Icelandic
		"mac", // Macedonian
		"mao", // Maori
		"may", // Malay
		"per", // Persian
		"rum", // Romanian; Moldavian; Moldovan
		"slo", // Slovak
		"tib", // Tibetan
		"wel", // Welsh
	}
)

// ISO 639-2/T-1998 Language codes (updated 2014-03-18)
// https://www.loc.gov/standards/iso639-2/php/code_list.php
// https://en.wikipedia.org/wiki/ISO_639-2
var (
	ISO_639_2T_1998_CODES []string = []string{
		"aar", // Afar
		"abk", // Abkhazian
		"ace", // Achinese
		"ach", // Acoli
		"ada", // Adangme
		"ady", // Adyghe; Adygei
		"afa", // Afro-Asiatic languages
		"afh", // Afrihili
		"afr", // Afrikaans
		"ain", // Ainu
		"aka", // Akan
		"akk", // Akkadian
		"ale", // Aleut
		"alg", // Algonquian languages
		"alt", // Southern Altai
		"amh", // Amharic
		"ang", // English, Old (ca.450-1100)
		"anp", // Angika
		"apa", // Apache languages
		"ara", // Arabic
		"arc", // Official Aramaic (700-300 BCE); Imperial Aramaic (700-300 BCE)
		"arg", // Aragonese
		"arn", // Mapudungun; Mapuche	mapudungun;
		"arp", // Arapaho
		"art", // Artificial languages
		"arw", // Arawak
		"asm", // Assamese
		"ast", // Asturian; Bable; Leonese;
		"ath", // Athapascan languages
		"aus", // Australian languages
		"ava", // Avaric
		"ave", // Avestan
		"awa", // Awadhi
		"aym", // Aymara
		"aze", // Azerbaijani
		"bad", // Banda languages
		"bai", // Bamileke languages
		"bak", // Bashkir
		"bal", // Baluchi
		"bam", // Bambara
		"ban", // Balinese
		"bas", // Basa
		"bat", // Baltic languages
		"bej", // Beja; Bedawiyet
		"bel", // Belarusian
		"bem", // Bemba
		"ben", // Bengali
		"ber", // Berber languages
		"bho", // Bhojpuri
		"bih", // Bihari languages
		"bik", // Bikol
		"bin", // Bini; Edo
		"bis", // Bislama
		"bla", // Siksika
		"bnt", // Bantu languages
		"bod", // Tibetan
		"bos", // Bosnian
		"bra", // Braj
		"bre", // Breton
		"btk", // Batak languages
		"bua", // Buriat
		"bug", // Buginese
		"bul", // Bulgarian
		"byn", // Blin; Bilin
		"cad", // Caddo
		"cai", // Central American Indian languages
		"car", // Galibi Carib
		"cat", // Catalan; Valencian
		"cau", // Caucasian languages
		"ceb", // Cebuano
		"cel", // Celtic languages
		"ces", // Czech
		"cha", // Chamorro
		"chb", // Chibcha
		"che", // Chechen
		"chg", // Chagatai
		"chk", // Chuukese
		"chm", // Mari
		"chn", // Chinook
		"cho", // Choctaw
		"chp", // Chipewyan; Dene Suline
		"chr", // Cherokee
		"chu", // Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic
		"chv", // Chuvash
		"chy", // Cheyenne
		"cmc", // Chamic languages
		"cop", // Coptic
		"cor", // Cornish
		"cos", // Corsican
		"cpe", // Creoles and pidgins, English based
		"cpf", // Creoles and pidgins, French-based
		"cpp", // Creoles and pidgins, Portuguese-based
		"cre", // Cree
		"crh", // Crimean Tatar; Crimean Turkish
		"crp", // Creoles and pidgins
		"csb", // Kashubian
		"cus", // Cushitic languages
		"cym", // Welsh
		"dak", // Dakota
		"dan", // Danish
		"dar", // Dargwa
		"day", // Land Dayak languages
		"del", // Delaware
		"den", // Slave (Athapascan)
		"deu", // German
		"dgr", // Dogrib
		"din", // Dinka
		"div", // Divehi; Dhivehi; Maldivian
		"doi", // Dogri
		"dra", // Dravidian languages
		"dsb", // Lower Sorbian
		"dua", // Duala
		"dum", // Dutch, Middle (ca.1050-1350)
		"dyu", // Dyula
		"dzo", // Dzongkha
		"efi", // Efik
		"egy", // Egyptian (Ancient)
		"eka", // Ekajuk
		"ell", // Greek, Modern (1453-)
		"elx", // Elamite
		"eng", // English
		"enm", // English, Middle (1100-1500)
		"epo", // Esperanto
		"est", // Estonian
		"eus", // Basque
		"ewe", // Ewe
		"ewo", // Ewondo
		"fan", // Fang
		"fao", // Faroese
		"fas", // Persian
		"fat", // Fanti
		"fij", // Fijian
		"fil", // Filipino; Pilipino
		"fin", // Finnish
		"fiu", // Finno-Ugrian languages
		"fon", // Fon
		"fra", // French
		"frm", // French, Middle (ca.1400-1600)
		"fro", // French, Old (842-ca.1400)
		"frr", // Northern Frisian
		"frs", // Eastern Frisian
		"fry", // Western Frisian
		"ful", // Fulah
		"fur", // Friulian
		"gaa", // Ga
		"gay", // Gayo
		"gba", // Gbaya
		"gem", // Germanic languages
		"gez", // Geez
		"gil", // Gilbertese
		"gla", // Gaelic; Scottish
		"gle", // Irish
		"glg", // Galician
		"glv", // Manx
		"gmh", // German, Middle High (ca.1050-1500)
		"goh", // German, Old High (ca.750-1050)
		"gon", // Gondi
		"gor", // Gorontalo
		"got", // Gothic
		"grb", // Grebo
		"grc", // Greek, Ancient (to 1453)
		"grn", // Guarani
		"gsw", // Swiss German; Alemannic; Alsatian
		"guj", // Gujarati
		"gwi", // Gwich'in
		"hai", // Haida
		"hat", // Haitian; Haitian Creole
		"hau", // Hausa
		"haw", // Hawaiian
		"heb", // Hebrew
		"her", // Herero
		"hil", // Hiligaynon
		"him", // Himachali languages; Western Pahari languages
		"hin", // Hindi
		"hit", // Hittite
		"hmn", // Hmong; Mong
		"hmo", // Hiri Motu
		"hrv", // Croatian
		"hsb", // Upper Sorbian
		"hun", // Hungarian
		"hup", // Hupa
		"hye", // Armenian
		"iba", // Iban
		"ibo", // Igbo
		"ido", // Ido
		"iii", // Sichuan Yi; Nuosu
		"ijo", // Ijo languages
		"iku", // Inuktitut
		"ile", // Interlingue; Occidental
		"ilo", // Iloko
		"ina", // Interlingua (International Auxiliary Language Association)
		"inc", // Indic languages
		"ind", // Indonesian
		"ine", // Indo-European languages
		"inh", // Ingush
		"ipk", // Inupiaq
		"ira", // Iranian languages
		"iro", // Iroquoian languages
		"isl", // Icelandic
		"isl", // Icelandic
		"ita", // Italian
		"jav", // Javanese
		"jbo", // Lojban
		"jpn", // Japanese
		"jpr", // Judeo-Persian
		"jrb", // Judeo-Arabic
		"kaa", // Kara-Kalpak
		"kab", // Kabyle
		"kac", // Kachin; Jingpho
		"kal", // Kalaallisut; Greenlandic
		"kam", // Kamba
		"kan", // Kannada
		"kar", // Karen languages
		"kas", // Kashmiri
		"kat", // Georgian
		"kau", // Kanuri
		"kaw", // Kawi
		"kaz", // Kazakh
		"kbd", // Kabardian
		"kha", // Khasi
		"khi", // Khoisan
		"khm", // Central
		"kho", // Khotanese; Sakan
		"kik", // Kikuyu; Gikuyu
		"kin", // Kinyarwanda
		"kir", // Kirghiz; Kyrgyz
		"kmb", // Kimbundu
		"kok", // Konkani
		"kom", // Komi
		"kon", // Kongo
		"kor", // Korean
		"kos", // Kosraean
		"kpe", // Kpelle
		"krc", // Karachay-Balkar
		"krl", // Karelian
		"kro", // Kru languages
		"kru", // Kurukh
		"kua", // Kuanyama; Kwanyama
		"kum", // Kumyk
		"kur", // Kurdish
		"kut", // Kutenai
		"lad", // Ladino
		"lah", // Lahnda
		"lam", // Lamba
		"lao", // Lao
		"lat", // Latin
		"lav", // Latvian
		"lez", // Lezghian
		"lim", // Limburgan; Limburger; Limburgish
		"lin", // Lingala
		"lit", // Lithuanian
		"lol", // Mongo
		"loz", // Lozi
		"ltz", // Luxembourgish; Letzeburgesch
		"lua", // Luba-Lulua
		"lub", // Luba-Katanga
		"lug", // Ganda
		"lui", // Luiseno
		"lun", // Lunda
		"luo", // Luo (Kenya and Tanzania)
		"lus", // Lushai
		"mad", // Madurese
		"mag", // Magahi
		"mah", // Marshallese
		"mai", // Maithili
		"mak", // Makasar
		"mal", // Malayalam
		"man", // Mandingo
		"map", // Austronesian languages
		"mar", // Marathi
		"mas", // Masai
		"mdf", // Moksha
		"mdr", // Mandar
		"men", // Mende
		"mga", // Irish, Middle (900-1200)
		"mic", // Mi'kmaq; Micmac
		"min", // Minangkabau
		"mis", // Uncoded languages
		"mkd", // Macedonian
		"mkh", // Mon-Khmer languages
		"mlg", // Malagasy
		"mlt", // Maltese
		"mnc", // Manchu
		"mni", // Manipuri
		"mno", // Manobo languages
		"moh", // Mohawk
		"mon", // Mongolian
		"mos", // Mossi
		"mri", // Maori
		"msa", // Malay
		"mul", // Multiple languages
		"mun", // Munda languages
		"mus", // Creek
		"mwl", // Mirandese
		"mwr", // Marwari
		"mya", // Burmese
		"myn", // Mayan languages
		"myv", // Erzya
		"nah", // Nahuatl languages
		"nai", // North American Indian languages
		"nap", // Neapolitan
		"nau", // Nauru
		"nav", // Navajo; Navaho
		"nbl", // Ndebele, South; South Ndebele
		"nde", // Ndebele, North; North Ndebele
		"ndo", // Ndonga
		"nds", // Low German; Low Saxon; German, Low; Saxon, Low
		"nep", // Nepali
		"new", // Nepal Bhasa; Newari
		"nia", // Nias
		"nic", // Niger-Kordofanian languages
		"niu", // Niuean
		"nld", // Dutch; Flemish
		"nno", // Norwegian Nynorsk; Nynorsk, Norwegian
		"nob", // Bokmål, Norwegian; Norwegian Bokmål
		"nog", // Nogai
		"non", // Norse, Old
		"nor", // Norwegian
		"nqo", // N'Ko
		"nso", // Pedi; Sepedi; Northern Sotho
		"nub", // Nubian languages
		"nwc", // Classical Newari; Old Newari; Classical Nepal Bhasa
		"nya", // Chichewa; Chewa; Nyanja
		"nym", // Nyamwezi
		"nyn", // Nyankole
		"nyo", // Nyoro
		"nzi", // Nzima
		"oci", // Occitan (post 1500)
		"oji", // Ojibwa
		"ori", // Oriya
		"orm", // Oromo
		"osa", // Osage
		"oss", // Ossetian; Ossetic
		"ota", // Turkish, Ottoman (1500-1928)
		"oto", // Otomian languages
		"paa", // Papuan languages
		"pag", // Pangasinan
		"pal", // Pahlavi
		"pam", // Pampanga; Kapampangan
		"pan", // Panjabi; Punjabi
		"pap", // Papiamento
		"pau", // Palauan
		"peo", // Persian, Old (ca.600-400 B.C.)
		"phi", // Philippine languages
		"phn", // Phoenician
		"pli", // Pali
		"pol", // Polish
		"pon", // Pohnpeian
		"por", // Portuguese
		"pra", // Prakrit languages
		"pro", // Provençal, Old (to 1500);Occitan, Old (to 1500)
		"pus", // Pushto; Pashto
		"qaa", // Reserved for local use
		"que", // Quechua
		"raj", // Rajasthani
		"rap", // Rapanui
		"rar", // Rarotongan; Cook Islands Maori
		"roa", // Romance languages
		"roh", // Romansh
		"rom", // Romany
		"ron", // Romanian; Moldavian; Moldovan
		"run", // Rundi
		"rup", // Aromanian; Arumanian; Macedo-Romanian
		"rus", // Russian
		"sad", // Sandawe
		"sag", // Sango
		"sah", // Yakut
		"sai", // South American Indian languages
		"sal", // Salishan languages
		"sam", // Samaritan Aramaic
		"san", // Sanskrit
		"sas", // Sasak
		"sat", // Santali
		"scn", // Sicilian
		"sco", // Scots
		"sel", // Selkup
		"sem", // Semitic languages
		"sga", // Irish, Old (to 900)
		"sgn", // Sign Languages
		"shn", // Shan
		"sid", // Sidamo
		"sin", // Sinhala; Sinhalese
		"sio", // Siouan languages
		"sit", // Sino-Tibetan languages
		"sla", // Slavic languages
		"slk", // Slovak
		"slv", // Slovenian
		"sma", // Southern Sami
		"sme", // Northern Sami
		"smi", // Sami languages
		"smj", // Lule Sami
		"smn", // Inari Sami
		"smo", // Samoan
		"sms", // Skolt Sami
		"sna", // Shona
		"snd", // Sindhi
		"snk", // Soninke
		"sog", // Sogdian
		"som", // Somali
		"son", // Songhai languages
		"sot", // Sotho, Southern
		"spa", // Spanish; Castilian
		"sqi", // Albanian
		"srd", // Sardinian
		"srn", // Sranan Tongo
		"srp", // Serbian
		"srr", // Serer
		"ssa", // Nilo-Saharan languages
		"ssw", // Swati
		"suk", // Sukuma
		"sun", // Sundanese
		"sus", // Susu
		"sux", // Sumerian
		"swa", // Swahili
		"swe", // Swedish
		"syc", // Classical Syriac
		"syr", // Syriac
		"tah", // Tahitian
		"tai", // Tai languages
		"tam", // Tamil
		"tat", // Tatar
		"tel", // Telugu
		"tem", // Timne
		"ter", // Tereno
		"tet", // Tetum
		"tgk", // Tajik
		"tgl", // Tagalog
		"tha", // Thai
		"tig", // Tigre
		"tir", // Tigrinya
		"tiv", // Tiv
		"tkl", // Tokelau
		"tlh", // Klingon; tlhIngan-Hol
		"tli", // Tlingit
		"tmh", // Tamashek
		"tog", // Tonga (Nyasa)
		"ton", // Tonga (Tonga Islands)
		"tpi", // Tok Pisin
		"tsi", // Tsimshian
		"tsn", // Tswana
		"tso", // Tsonga
		"tuk", // Turkmen
		"tum", // Tumbuka
		"tup", // Tupi
		"tur", // Turkish
		"tut", // Altaic
		"tvl", // Tuvalu
		"twi", // Twi
		"tyv", // Tuvinian
		"udm", // Udmurt
		"uga", // Ugaritic
		"uig", // Uighur; Uyghur
		"ukr", // Ukrainian
		"umb", // Umbundu
		"und", // Undetermined
		"urd", // Urdu
		"uzb", // Uzbek
		"vai", // Vai
		"ven", // Venda
		"vie", // Vietnamese
		"vol", // Volapük
		"vot", // Votic
		"wak", // Wakashan languages
		"wal", // Wolaitta; Wolaytta
		"war", // Waray
		"was", // Washo
		"wen", // Sorbian languages
		"wln", // Walloon
		"wol", // Wolof
		"xal", // Kalmyk; Oirat
		"xho", // Xhosa
		"yao", // Yao
		"yap", // Yapese
		"yid", // Yiddish
		"yor", // Yoruba
		"ypk", // Yupik languages
		"zap", // Zapotec
		"zbl", // Blissymbols; Blissymbolics; Bliss
		"zen", // Zenaga
		"zgh", // Standard Moroccan Tamazight
		"zha", // Zhuang; Chuang
		"zho", // Chinese
		"znd", // Zande languages
		"zul", // Zulu
		"zun", // Zuni
		"zxx", // No linguistic content; Not applicable
		"zza", // Zaza; Dimili; Dimli; Kirdki; Kirmanjki; Zazaki
	}
)

var ISO_639_2B_TO_2T_MAP map[string]string = map[string]string{
	"alb": "sqi", // Albanian
	"arm": "hye", // Armenian
	"baq": "eus", // Basque
	"bur": "mya", // Burmese
	"chi": "zho", // Chinese
	"cze": "ces", // Czech
	"dut": "nld", // Dutch; Flemish
	"fre": "fra", // French
	"geo": "kat", // Georgian
	"ger": "deu", // German
	"gre": "ell", // Modern Greek (1453–)
	"ice": "isl", // Icelandic
	"mac": "mkd", // Macedonian
	"mao": "mri", // Maori
	"may": "msa", // Malay
	"per": "fas", // Persian
	"rum": "ron", // Romanian; Moldavian; Moldovan
	"slo": "slk", // Slovak
	"tib": "bod", // Tibetan
	"wel": "cym", // Welsh
}

var ISO_639_1_TO_2T_MAP map[string]string = map[string]string{
	"aa": "aar", // Afar
	"ab": "abk", // Abkhazian
	"ae": "ave", // Avestan
	"af": "afr", // Afrikaans
	"ak": "aka", // Akan
	"am": "amh", // Amharic
	"an": "arg", // Aragonese
	"ar": "ara", // Arabic
	"as": "asm", // Assamese
	"av": "ava", // Avaric
	"ay": "aym", // Aymara
	"az": "aze", // Azerbaijani
	"ba": "bak", // Bashkir
	"be": "bel", // Belarusian
	"bg": "bul", // Bulgarian
	"bh": "bih", // Bihari languages
	"bi": "bis", // Bislama
	"bm": "bam", // Bambara
	"bn": "ben", // Bengali
	"bo": "bod", // Tibetan
	"br": "bre", // Breton
	"bs": "bos", // Bosnian
	"ca": "cat", // Catalan; Valencian
	"ce": "che", // Chechen
	"ch": "cha", // Chamorro
	"co": "cos", // Corsican
	"cr": "cre", // Cree
	"cs": "ces", // Czech
	"cu": "chu", // Church Slavic; Old Slavonic; Church Slavonic; Old Bulgarian; Old Church Slavonic
	"cv": "chv", // Chuvash
	"cy": "cym", // Welsh
	"da": "dan", // Danish
	"de": "deu", // German
	"dv": "div", // Divehi; Dhivehi; Maldivian
	"dz": "dzo", // Dzongkha
	"ee": "ewe", // Ewe
	"el": "ell", // Greek, Modern (1453-)
	"en": "eng", // English
	"eo": "epo", // Esperanto
	"es": "spa", // Spanish; Castilian
	"et": "est", // Estonian
	"eu": "eus", // Basque
	"fa": "fas", // Persian
	"ff": "ful", // Fulah
	"fi": "fin", // Finnish
	"fj": "fij", // Fijian
	"fo": "fao", // Faroese
	"fr": "fra", // French
	"fy": "fry", // Western Frisian
	"ga": "gle", // Irish
	"gd": "gla", // Gaelic; Scottish
	"gl": "glg", // Galician
	"gn": "grn", // Guarani
	"gu": "guj", // Gujarati
	"gv": "glv", // Manx
	"ha": "hau", // Hausa
	"he": "heb", // Hebrew
	"hi": "hin", // Hindi
	"ho": "hmo", // Hiri Motu
	"hr": "hrv", // Croatian
	"ht": "hat", // Haitian; Haitian Creole
	"hu": "hun", // Hungarian
	"hy": "hye", // Armenian
	"hz": "her", // Herero
	"ia": "ina", // Interlingua (International Auxiliary Language Association)
	"id": "ind", // Indonesian
	"ie": "ile", // Interlingue; Occidental
	"ig": "ibo", // Igbo
	"ii": "iii", // Sichuan Yi; Nuosu
	"ik": "ipk", // Inupiaq
	"io": "ido", // Ido
	"is": "isl", // Icelandic
	"it": "ita", // Italian
	"iu": "iku", // Inuktitut
	"ja": "jpn", // Japanese
	"jv": "jav", // Javanese
	"ka": "kat", // Georgian
	"kg": "kon", // Kongo
	"ki": "kik", // Kikuyu; Gikuyu
	"kj": "kua", // Kuanyama; Kwanyama
	"kk": "kaz", // Kazakh
	"kl": "kal", // Kalaallisut; Greenlandic
	"km": "khm", // Central
	"kn": "kan", // Kannada
	"ko": "kor", // Korean
	"kr": "kau", // Kanuri
	"ks": "kas", // Kashmiri
	"ku": "kur", // Kurdish
	"kv": "kom", // Komi
	"kw": "cor", // Cornish
	"ky": "kir", // Kirghiz; Kyrgyz
	"la": "lat", // Latin
	"lb": "ltz", // Luxembourgish; Letzeburgesch
	"lg": "lug", // Ganda
	"li": "lim", // Limburgan; Limburger; Limburgish
	"ln": "lin", // Lingala
	"lo": "lao", // Lao
	"lt": "lit", // Lithuanian
	"lu": "lub", // Luba-Katanga
	"lv": "lav", // Latvian
	"mg": "mlg", // Malagasy
	"mh": "mah", // Marshallese
	"mi": "mri", // Maori
	"mk": "mkd", // Macedonian
	"ml": "mal", // Malayalam
	"mn": "mon", // Mongolian
	"mr": "mar", // Marathi
	"ms": "msa", // Malay
	"mt": "mlt", // Maltese
	"my": "mya", // Burmese
	"na": "nau", // Nauru
	"nb": "nob", // Bokmål, Norwegian; Norwegian Bokmål
	"nd": "nde", // Ndebele, North; North Ndebele
	"ne": "nep", // Nepali
	"ng": "ndo", // Ndonga
	"nl": "nld", // Dutch; Flemish
	"nn": "nno", // Norwegian Nynorsk; Nynorsk, Norwegian
	"no": "nor", // Norwegian
	"nr": "nbl", // Ndebele, South; South Ndebele
	"nv": "nav", // Navajo; Navaho
	"ny": "nya", // Chichewa; Chewa; Nyanja
	"oc": "oci", // Occitan (post 1500)
	"oj": "oji", // Ojibwa
	"om": "orm", // Oromo
	"or": "ori", // Oriya
	"os": "oss", // Ossetian; Ossetic
	"pa": "pan", // Panjabi; Punjabi
	"pi": "pli", // Pali
	"pl": "pol", // Polish
	"ps": "pus", // Pushto; Pashto
	"pt": "por", // Portuguese
	"qu": "que", // Quechua
	"rm": "roh", // Romansh
	"rn": "run", // Rundi
	"ro": "ron", // Romanian; Moldavian; Moldovan
	"ru": "rus", // Russian
	"rw": "kin", // Kinyarwanda
	"sa": "san", // Sanskrit
	"sc": "srd", // Sardinian
	"sd": "snd", // Sindhi
	"se": "sme", // Northern Sami
	"sg": "sag", // Sango
	"si": "sin", // Sinhala; Sinhalese
	"sk": "slk", // Slovak
	"sl": "slv", // Slovenian
	"sm": "smo", // Samoan
	"sn": "sna", // Shona
	"so": "som", // Somali
	"sq": "sqi", // Albanian
	"sr": "srp", // Serbian
	"ss": "ssw", // Swati
	"st": "sot", // Sotho, Southern
	"su": "sun", // Sundanese
	"sv": "swe", // Swedish
	"sw": "swa", // Swahili
	"ta": "tam", // Tamil
	"te": "tel", // Telugu
	"tg": "tgk", // Tajik
	"th": "tha", // Thai
	"ti": "tir", // Tigrinya
	"tk": "tuk", // Turkmen
	"tl": "tgl", // Tagalog
	"tn": "tsn", // Tswana
	"to": "ton", // Tonga (Tonga Islands)
	"tr": "tur", // Turkish
	"ts": "tso", // Tsonga
	"tt": "tat", // Tatar
	"tw": "twi", // Twi
	"ty": "tah", // Tahitian
	"ug": "uig", // Uighur; Uyghur
	"uk": "ukr", // Ukrainian
	"ur": "urd", // Urdu
	"uz": "uzb", // Uzbek
	"ve": "ven", // Venda
	"vi": "vie", // Vietnamese
	"vo": "vol", // Volapük
	"wa": "wln", // Walloon
	"wo": "wol", // Wolof
	"xh": "xho", // Xhosa
	"yi": "yid", // Yiddish
	"yo": "yor", // Yoruba
	"za": "zha", // Zhuang; Chuang
	"zh": "zho", // Chinese
	"zu": "zul", // Zulu
}

type Language string

const (
	LanguageUndefined Language = ""
)

func (l Language) String() string {
	return string(l)
}

func ParseLanguage(l string) Language {
	l = strings.ToLower(l)
	switch len(l) {
	case 2:
		// parse as ISO 639-1 code and convert to ISO 639-2
		for _, x := range ISO_639_1_2002_CODES {
			if x == l {
				if v, ok := ISO_639_1_TO_2T_MAP[x]; ok {
					return Language(v)
				}
			}
		}
	case 3:
		// parse as ISO 639-2/T language code
		for _, x := range ISO_639_2T_1998_CODES {
			if x == l {
				return Language(l)
			}
		}
		// parse as ISO 639-2/B
		for _, x := range ISO_639_2B_1998_CODES {
			if x == l {
				if v, ok := ISO_639_2B_TO_2T_MAP[x]; ok {
					return Language(v)
				}
			}
		}
	}
	return LanguageUndefined
}

func (r Language) IsValid() bool {
	return r != LanguageUndefined
}

// Text/JSON conversion
func (r Language) MarshalText() ([]byte, error) {
	return []byte(r), nil
}

func (r *Language) UnmarshalText(data []byte) error {
	rr := ParseLanguage(string(data))
	if !rr.IsValid() {
		return fmt.Errorf("iso: invalid ISO 639-1/2BT:2002 language code '%s'", string(data))
	}
	*r = rr
	return nil
}

// SQL conversion
func (r *Language) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*r = ParseLanguage(v)
	case []byte:
		*r = ParseLanguage(string(v))
	}
	if !(*r).IsValid() {
		return fmt.Errorf("iso: invalid ISO 639-1/2BT:2002 language code '%v'", value)
	}
	return nil
}

func (r Language) Value() (driver.Value, error) {
	return string(r), nil
}
