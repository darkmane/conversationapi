package conversationapi

type User struct {
	UserId      string        `json:"userId"`
	Profile     *UserProfile  `json:"profile"`
	AccessToken string        `json:"accessToken"`
	Permissions *[]Permission `json:"permissions"`
	Locale      string        `json:"locale"`
	LastSeen    string        `json:"lastSeen"`
	UserStorage string        `json:"userStorage"`
}
