package conversationapi

type UserProfile struct {
	DisplayName string `json:"displayName"`
	GivenName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
}
