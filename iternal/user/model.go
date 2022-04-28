package user

type DaysInWeek struct {
	DayOfWeek string `json:"dayOfWeek"`
	IsEntered bool   `json:"isEntered"`
	Time      string `json:"time"`
}

type User struct {
	Id              string       `json:"id"`
	Username        string       `json:"username"`
	Email           string       `json:"email"`
	Password        string       `json:"password"`
	Level           string       `json:"level"`
	DaysInRow       string       `json:"daysInRow"`
	DaysInWeek      []DaysInWeek `json:"daysInWeek"`
	DoesSendPushUps bool         `json:"doesSendPushUps"`
	Theme           string       `json:"theme"`
	Language        string       `json:"language"`
	Image           string       `json:"image"`
}
