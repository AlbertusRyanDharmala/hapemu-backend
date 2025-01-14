package model

type DxoMark struct {
	Name    string `json:"name"`
	Camera  int    `json:"mobileScore"`
	Selfie  int    `json:"selfieScore"`
	Audio   int    `json:"audioScore"`
	Display int    `json:"displayScore"`
	Battery int    `json:"batteryScore"`
	Mobile  Mobile `json:"mobile"`
}

type Mobile struct {
	Subscores Subscores `json:"subscores"`
}

type Subscores struct {
	Photo   int `json:"photo"`
	Bokeh   int `json:"bokeh"`
	Preview int `json:"preview"`
	Zoom    int `json:"zoom"`
	Video   int `json:"video"`
}

type DxoMarkRequest struct {
	SortBy string
}
