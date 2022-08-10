package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sport-test/entity"
)

func RequestNewsInformation() {
	var newsInformation entity.NewsInformation
	if xmlBytes, err := getXML("https://www.wearehullcity.co.uk/api/incrowd/getnewlistinformation?count=50"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Fatal(xmlBytes)
		xml.Unmarshal(xmlBytes, &newsInformation)
	}
}
func RequestNewsDetails(idExternal string) {
	var newsInformation entity.NewsInformation
	if xmlBytes, err := getXML("https://www.wearehullcity.co.uk/api/incrowd/getnewsarticleinformation?id=" + idExternal); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		xml.Unmarshal(xmlBytes, &newsInformation)
	}
}

func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}
