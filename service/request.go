package service

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sport-test/entity"
)

func RequestNewsInformation() (entity.NewListInformation, error) {
	var newsInformation entity.NewListInformation
	if xmlBytes, err := getXML(entity.URL_SPORT_MANY); err != nil {
		log.Printf("Failed to get XML: %v", err)
		return newsInformation, err
	} else {
		err = xml.Unmarshal(xmlBytes, &newsInformation)
		if err != nil {
			log.Printf("Failed to Parser XML: %v", err)
			return newsInformation, err
		}
	}
	return newsInformation, nil
}
func RequestNewsDetails(idExternal string) (entity.NewsArticleInformation, error) {
	var newsArticle entity.NewsArticleInformation
	if xmlBytes, err := getXML(entity.URL_SPORT_ONE + idExternal); err != nil {
		log.Printf("Failed to get XML: %v", err)
		return newsArticle, err
	} else {
		err = xml.Unmarshal(xmlBytes, &newsArticle)
		if err != nil {
			log.Printf("Failed to Parser XML: %v", err)
			return newsArticle, err
		}
	}
	return newsArticle, nil
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
