package entities

type RequestLocation struct {
	Satellites []InfoSatellite `json:"satellites"`
}

type InfoSatellite struct {
	Name     string   `json:"name"`
	Distance float32  `json:"distance"`
	Message  []string `json:"message"`
}
