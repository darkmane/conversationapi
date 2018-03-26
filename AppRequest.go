package conversationapi

type AppRequest struct {
	User         *User         `json:"user,omitempty"`
	Device       *Device       `json:"device,omitempty"`
	Surface      *Surface      `json:"surface,omitempty"`
	Conversation *Conversation `json:"conversation,omitempty"`
	Inputs       *[]Input      `json:"inputs,omitempty"`
	IsInSandbox  bool          `json:"isInSandbox"`
}
