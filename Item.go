package conversationapi

type Item struct {
	SimpleResponse     *SimpleResponse     `json:"simpleResponse,omitempty"`
	BasicCard          *BasicCard          `json:"basicCard,omitempty"`
	StructuredResponse *StructuredResponse `json:"structuredResponse,omitempty"`
}
