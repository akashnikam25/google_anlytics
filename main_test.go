package main

import (
	"testing"
)

func Test_decodeData(t *testing.T) {
	data, err := decodeData("eehhe")
	if err != nil {
		t.Fatal(err)
	} else if data.SiteID != "my-site-id-here" {
		t.Errorf("expected 'my-site-id-here' got %s",data.SiteID)
	}
}
