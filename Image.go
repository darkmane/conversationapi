package conversationapi

type Image struct {
	Url               string `json:"url"`
	AccessibilityText string `json:"accessibilityText"`
	Height            int64  `json:"height"`
	Width             int64  `json:"width"`
}
