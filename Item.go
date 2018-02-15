package conversationapi

type Item struct {
	SimpleResponse     SimpleResponse     `json:"simpleResponse"`
	BasicCard          BasicCard          `json:"basicCard"`
	StructuredResponse StructuredResponse `json:"structuredResponse"`
}
