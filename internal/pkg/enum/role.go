package enum

type Role     int64

const (
	Admin Role = iota + 1
	User
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case User:
		return "User"
	}
	return "Unknown"
}
