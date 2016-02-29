package fzp

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

type Fzp struct {
	XMLName         xml.Name    `xml:"module"`
	FritzingVersion string      `xml:"fritzingVersion,attr"`
	ModuleId        string      `xml:"moduleId,attr"`
	ReferenceFile   string      `xml:"referenceFile,attr"`
	Version         string      `xml:"version"`
	Title           string      `xml:"title"`
	Description     string      `xml:"description"`
	Author          string      `xml:"author"`
	Date            string      `xml:"date"`
	URL             string      `xml:"url"`
	Label           string      `xml:"label"`
	//Taxonomy        string      `xml:"taxonomy"`
	//Family          string      `xml:"family"`
	//Variant         string      `xml:"variant"`
	Tags            []string    `xml:"tags>tag"`
	Properties      []Property  `xml:"properties>property"`
	Views           Views       `xml:"views"`
	Connectors      []Connector `xml:"connectors>connector"`
	Buses           []Bus       `xml:"buses>bus"`
}

const FileExtensionFzp = ".fzp"

func IsFileFzp(src string) bool {
	if filepath.Ext(src) == FileExtensionFzp {
		return true
	} else {
		return false
	}
}

// ReadFzp and decode xml data
func ReadFzp(src string) (Fzp, error) {
	fzpData := Fzp{}
	// read
	fzpBytes, err := ioutil.ReadFile(src)
	if err != nil {
		return fzpData, err
	}
	// decode XML
	err = xml.Unmarshal(fzpBytes, &fzpData)
	if err != nil {
		return fzpData, err
	}
	return fzpData, nil
}

func (f *Fzp) CheckFritzingVersion() error {
	if f.FritzingVersion == "" {
		return errors.New("fritzingVersion undefined")
	}
	return nil
}

func (f *Fzp) CheckModuleId() error {
	if f.ModuleId == "" {
		return errors.New("moduleId undefined")
	}
	return nil
}

// TODO: is the referenceFile required?

func (f *Fzp) CheckVersion() error {
	if f.Version == "" {
		return errors.New("version undefined")
	}
	return nil
}

func (f *Fzp) CheckTitle() error {
	if f.Title == "" {
		return errors.New("title undefined")
	}
	return nil
}

func (f *Fzp) CheckDescription() error {
	if f.Description == "" {
		return errors.New("description undefined")
	}
	return nil
}

func (f *Fzp) CheckAuthor() error {
	if f.Author == "" {
		return errors.New("author undefined")
	}
	return nil
}

// Check Date ?
// Check URL ?
// Check Label ?
// Check Taxonomy ?
// Check Family ?
// Check Variant ?

func (f *Fzp) CheckTags() (error, int) {
	countBrokenTags := 0

	if len(f.Tags) != 0 {
		for _, tag := range f.Tags {
			if tag == "" {
				countBrokenTags++
			}
		}
	}

	if countBrokenTags == 0 {
		return nil, countBrokenTags
	} else {
		errMsg := strconv.Itoa(countBrokenTags) + " tag value/s undefined"
		return errors.New(errMsg), countBrokenTags
	}
}

func (f *Fzp) CheckProperties() error {
	if len(f.Properties) != 0 {
		for _, property := range f.Properties {
			if err := property.CheckName(); err != nil {
				return err
			}
			if err := property.CheckValue(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Fzp) CheckViews() error {
	// TODO: ...
	return nil
}

func (f *Fzp) CheckConnectors() error {
	if len(f.Connectors) != 0 {
		for _, connector := range f.Connectors {
			if err := connector.Check(); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *Fzp) CheckBuses() error {
	if len(f.Buses) != 0 {
		for _, bus := range f.Buses {
			if err := bus.CheckId(); err != nil {
				return err
			}
		}
	}
	return nil
}

// check all the data
func (f *Fzp) Check() []error {
	var errList []error

	if err := f.CheckFritzingVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckModuleId(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckVersion(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckTitle(); err != nil {
		errList = append(errList, err)
	}
	errTags, _ := f.CheckTags()
	if errTags != nil {
		errList = append(errList, errTags)
	}
	if err := f.CheckProperties(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckViews(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckConnectors(); err != nil {
		errList = append(errList, err)
	}
	if err := f.CheckBuses(); err != nil {
		errList = append(errList, err)
	}

	return errList
}
