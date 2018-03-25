package conversationapi

type AppRequest struct {
	User         *User         `json:"user"`
	Device       *Device       `json:"device"`
	Surface      *Surface      `json:"surface"`
	Conversation *Conversation `json:"conversation"`
	Inputs       *[]Input      `json:"inputs"`
	IsInSandbox  bool          `json:"isInSandbox"`
}
