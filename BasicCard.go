package conversationapi

type BasicCard struct {
	Title               string               `json:"title"`
	Subtitle            string               `json:"subtitle"`
	FormattedText       string               `json:"formattedText"`
	Image               *Image               `json:"image,omitempty"`
	Buttons             *[]Button            `json:"buttons,omitempty"`
	ImageDisplayOptions *ImageDisplayOptions `json:"imageDisplayOptions,omitempty"`
}
