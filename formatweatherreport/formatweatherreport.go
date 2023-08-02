package formatweatherreport

import (
	"fmt"
	"strings"
)

type ReportData struct {
	CountryAbbrev string
	State         string
	City          string
}

func ValidateInput(data ReportData) error {
	err := CheckCountryAbbrev(data.CountryAbbrev)
	if err != nil {
		return err
	}

	if data.CountryAbbrev == "us" {
		err := GetUSCity(data.City)
		if err != nil {
			return err
		}
	}

	if data.CountryAbbrev == "" && data.City == "" {
		return fmt.Errorf("cannot process request due to invalid input")
	}

	return nil
}

func FormatDataForReport(data *ReportData) error {
	data.CountryAbbrev = strings.TrimSpace(data.CountryAbbrev)
	data.State = strings.TrimSpace(data.State)
	data.City = strings.TrimSpace(data.City)

	data.City = strings.ReplaceAll(data.City, " ", "-")

	return nil
}

func CheckCountryAbbrev(country string) error {
	countries := []string{"ad", "ae", "af", "ag", "ai", "al", "am", "ao", "aq", "ar",
		"as", "at", "au", "aw", "ax", "az", "ba", "bb", "bd", "be", "bf", "bg", "bh", "bi",
		"bj", "bl", "bm", "bn", "bo", "bq", "br", "bs", "bt", "bv", "bw", "by", "bz", "ca",
		"cc", "cd", "cf", "cg", "ch", "ci", "ck", "cl", "cm", "cn", "co", "cr", "cu", "cv",
		"cw", "cx", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "ee", "eg", "eh",
		"er", "es", "et", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gb", "gd", "ge", "gf",
		"gg", "gh", "gi", "gl", "gm", "gn", "gp", "gq", "gr", "gs", "gt", "gu", "gw", "gy",
		"hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il", "im", "in", "io", "iq", "ir",
		"is", "it", "je", "jm", "jo", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr",
		"kw", "ky", "kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv", "ly",
		"ma", "mc", "md", "me", "mf", "mg", "mh", "mk", "ml", "mm", "mn", "mo", "mp", "mq",
		"mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz", "na", "nc", "ne", "nf", "ng",
		"ni", "nl", "no", "np", "nr", "nu", "nz", "om", "pa", "pe", "pf", "pg", "ph", "pk",
		"pl", "pm", "pn", "pr", "ps", "pt", "pw", "py", "qa", "re", "ro", "rs", "ru", "rw",
		"sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk", "sl", "sm", "sn", "so",
		"sr", "ss", "st", "sv", "sx", "sy", "sz", "tc", "td", "tf", "tg", "th", "tj", "tk",
		"tl", "tm", "tn", "to", "tr", "tt", "tv", "tw", "tz", "ua", "ug", "um", "us", "uy",
		"uz", "va", "vc", "ve", "vg", "vi", "vn", "vu", "wf", "ws", "ye", "yt", "za", "zm",
		"zw"}

	for _, selectedCountry := range countries {
		if selectedCountry == country {
			return nil
		}
	}

	return fmt.Errorf("country doesn't match the two letter abbrevation for countries in the world")
}

func GetUSCity(state string) error {
	states := []string{"al", "ak", "az", "ar", "ca", "co", "ct", "de", "fl", "ga", "hi",
		"id", "il", "in", "ia", "ks", "ky", "la", "me", "md", "ma", "mi", "mn", "ms", "mo",
		"mt", "ne", "nv", "nh", "nj", "nm", "ny", "nc", "nd", "oh", "ok", "or", "pa", "ri",
		"sc", "sd", "tn", "tx", "ut", "vt", "va", "wa", "wi", "wv", "wy", "as", "dc", "gu",
		"mp", "pr", "vi"}

	for _, aState := range states {
		if aState == state {
			return nil
		}
	}

	return fmt.Errorf("state does not exist")
}
