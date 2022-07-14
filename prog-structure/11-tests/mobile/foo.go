package mobile

const (
	UA = "UA"
	UK = "UK"
)

func ParsePhoneNumber(number, country string) string {
	switch country {
	case UA:
		return "(" + UA + ")" + string(number[3:])
	case UK:
		return "(" + UK + ")" + string(number[3:])
	}
	return ""
}
