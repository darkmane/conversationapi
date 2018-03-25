package conversationapi

type BasicCard struct {
	Title               string               `json:"title"`
	Subtitle            string               `json:"subtitle"`
	FormattedText       string               `json:"formattedText"`
	Image               *Image               `json:"image"`
	Buttons             *[]Button            `json:"buttons"`
	ImageDisplayOptions *ImageDisplayOptions `json:"imageDisplayOptions"`
}
