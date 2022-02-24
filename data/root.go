package data

type MapImageResponse struct {
	Lat     float64 `json:"lat"`
	Long    float64 `json:"long"`
	Zoom    int     `json:"zoom"`
	Image   string  `json:"image"`   // The map image (in base64 encoded data uri format)
	Version string  `json:"version"` // Service version
}
