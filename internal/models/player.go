package models

type Player struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Team          string `json:"team"`
	ImageURL      string `json:"image_url"`
	Championships int    `json:"championships"`
	MVP           int    `json:"mvp"`
	FinalsMVP     int    `json:"finals_mvp"`
	DPOY          int    `json:"dpoy"`
	ROTY          int    `json:"roty"`
}
