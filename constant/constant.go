package constant

var Limit = map[int]int{
	20: 20,
	60: 60,
}

const (
	LimitMore = "https://data.bmkg.go.id/gempaterkini.xml"
	LimitLess = "https://data.bmkg.go.id/gempadirasakan.xml"

	MaxLimit = 60
)

func GetLimit(number int) int {
	val, ok := Limit[number]
	if !ok {
		return 0
	}

	return val
}

func Check(number int) bool {
	_, ok := Limit[number]
	if !ok {
		return ok
	}

	return ok
}
