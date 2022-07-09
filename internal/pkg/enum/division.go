package enum

type Division int64

const (
	Finance Division = iota + 1
	IT
	HR 
)

func (d Division) String() string {
	switch d {
	case Finance:
		return "Finance"
	case IT:
		return "Information Technology"
	case HR:
		return "Human Resource"
	}
	return "Unknown"
}
