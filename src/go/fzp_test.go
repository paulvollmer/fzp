package fzp

import (
	"testing"
)

func Test_IsFileFzp(t *testing.T) {
	result := IsFileFzp("data.fzp")
	if !result {
		t.Error("IsFileFzp broken")
	}
	result = IsFileFzp("data.notafzp")
	if result {
		t.Error("IsFileFzp broken")
	}
}

func Test_ReadFzp_Ok(t *testing.T) {
	f, err := ReadFzp("../../template.fzp")
	if err != nil {
		t.Error("Fzp.ReadFzp broken")
	}

	errCheck := f.Check()
	if errCheck != nil {
		t.Error("Fzp.Check broken:", errCheck)
	}
}

func Test_ReadFzp_Failed(t *testing.T) {
	_, err := ReadFzp("../not.found")
	if err == nil {
		t.Error("Fzp.ReadFzp (that doesn't exists) broken")
	}
}

func Test_ReadFzp_CheckTags(t *testing.T) {
	// fake data
	fzpData := Fzp{}
	fzpData.Tags = append(fzpData.Tags, "")
	// was an error returned?
	err, _ := fzpData.CheckTags()
	if err == nil {
		t.Error("Fzp.CheckTags broken")
	}
}
