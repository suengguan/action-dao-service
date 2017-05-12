package models

type Action struct {
	Id        int64
	Time      int64
	SessionId string
	User      *User
	DevType   int
	Type      int
	Content   string
}
