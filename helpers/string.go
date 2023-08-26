package helpers

import (
	"regexp"
	"strings"
)

// IsValidEmail : email
func IsValidEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)
}

// ToSlug : string
func ToSlug(s string) string {
	s = strings.ToLower(s)
	charReplace := map[string]string{
		"[ḁͦ|ѧ|á|à|ả|ã|ạ|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ|α|ɐ|ą|ä|ค|a̫]": "a",
		"[b̥ͦ|ɞ|b̫|в]": "b",
		"[c̥ͦ|¢|ċ|c̫]": "c",
		"[d̥ͦ|đ|d̫|Ԁ]": "d",
		"[e̥ͦ|é|è|ẻ|ẽ|ẹ|ê|ế|ề|ể|ễ|ệ|є|ĕ|e̫]": "e",
		"[f̥ͦ|ƒ|f̫]":         "f",
		"[g̥ͦ|ɢ|ğ|g̫]":       "g",
		"[h̥ͦ|ђ|h̫|һ]":       "h",
		"[i̥ͦ|í|ì|ỉ|ĩ|ị|i̫]": "i",
		"[נ|ĵ|j̫|j̥ͦ]":       "j",
		"[ʞ|k̫|ҡ|k̥ͦ]":       "k",
		"[ℓ|l̫|ʟ|l̥ͦ]":       "l",
		"[ή|ń|n̫|ṅ|n̥ͦ]":     "n",
		"[ﾶ|m̫|ṃ|m̥ͦ]":       "m",
		"[ó|ò|õ|ọ|ỏ|ô|ố|ồ|ổ|ỗ|ộ|ơ|ớ|ờ|ỡ|ợ|ở|ö|o̫|ȏ|o̥ͦ]": "o",
		"[p̫|ƿ|p̥ͦ]":     "p",
		"[q̫|զ|q̥ͦ]":     "q",
		"[ŕ|я|r̫|r̥ͦ]":   "r",
		"[ś|ƨ|s̫|ṡ|s̥ͦ]": "s",
		"[ŧ|ʇ|t̫|ṭ|t̥ͦ]": "t",
		"[ú|ù|ũ|ủ|ụ|ư|ứ|ừ|ữ|ử|ự|u̫|ȗ|u̥ͦ]": "u",
		"[v̫|ṿ|v̥ͦ]":           "v",
		"[x̫|×|x̥ͦ]":           "x",
		"[ý|ỳ|ỷ|ỹ|ỵ|y̫|ʏ|y̥ͦ]": "y",
		"[z̫|ẓ|z̥ͦ]":           "z",
		"[w̫|ẇ|w̥ͦ]":           "w",
	}

	for r, c := range charReplace {
		reg, _ := regexp.Compile(r)
		s = reg.ReplaceAllString(s, c)
	}

	reg, _ := regexp.Compile("[^a-zA-Z\\d]+")
	s = reg.ReplaceAllString(s, "-")

	return strings.Trim(s, "-")
}
