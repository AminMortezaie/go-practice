package pkgquery

import (
	"encoding/json"
	"io"
	"net/http"
)

/**
The structure has two string fields: Name and Version. The struct tags
`json:"name"` and "json:"version"` indicate the key identifiers for the
corresponding fields in the JSON data. Now that we have defined the data
structure, we can deserialize the JSON data for the packages into a slice of
pkgData objects.
**/

type pkgData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func fetchPackageData(url string) ([]pkgData, error) {
	var packages []pkgData
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "application/json" {
		return packages, nil
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return packages, err
	}
	err = json.Unmarshal(data, &packages)
	return packages, err
}
