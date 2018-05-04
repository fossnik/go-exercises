// programmatically returns sardonic responses to text input
package bob

// import string trimming
import (
	"strings"
)

// remark is input
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if isallcaps(remark) && containssomeletters(remark){
		if lastcharacterisquestionmark(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	}

	if lastcharacterisquestionmark(remark) {
		return "Sure."
	}

	if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever.";
}

func isallcaps(remark string) bool {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	for _, r:= range remark {
		if strings.Contains(lowercase, string(r)) {
			return false
		}
	}

	return true
}

func lastcharacterisquestionmark(remark string) bool {
	if len(remark) > 1 {
		return strings.Contains(remark[len(remark)-1:], "?")
	}
	return false
}

func containssomeletters(remark string) bool {
	return strings.ToLower(remark) != strings.ToUpper(remark)
}