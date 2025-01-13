package model

import "regexp"

var (
	regKanaOnly = regexp.MustCompile("^[\u30A0-\u30FF]+$")
)
