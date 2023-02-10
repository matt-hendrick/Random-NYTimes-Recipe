package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SiteMapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	SiteMaps []SiteMap `xml:"sitemap"`
}

type SiteMap struct {
	XMLName xml.Name `xml:"sitemap"`
	Loc     string   `xml:"loc"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func main() {
	var urlList []string
	if siteMapUrlArr, err := getSiteMapURLs(); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
	} else {
		fmt.Println("SiteMap Count is: ", len(siteMapUrlArr))
		for i, url := range siteMapUrlArr {
			fmt.Println("Iteration: ", i+1)
			time.Sleep(5 * time.Second)
			if xmlBytes, err := getXML(url); err != nil {
				fmt.Printf("Failed to get XML: %v", err)
			} else {

				var urlSet URLSet

				if err := xml.Unmarshal(xmlBytes, &urlSet); err != nil {
					fmt.Println("Error")
				}
				fmt.Println("UrlSet count is: ", len(urlSet.Urls))
				for _, url := range urlSet.Urls {
					urlList = append(urlList, url.Loc)
				}

			}
		}
		file, _ := json.Marshal(urlList)

		_ = ioutil.WriteFile("export.json", file, 0644)
	}
}

func getSiteMapURLs() ([]string, error) {
	var siteMapUrlArr []string
	if xmlBytes, err := getXML("https://www.nytimes.com/sitemaps/new/cooking.xml.gz"); err != nil {
		fmt.Printf("Failed to get XML: %v", err)
		return siteMapUrlArr, err
	} else {
		var siteMapIndex SiteMapIndex

		if err := xml.Unmarshal(xmlBytes, &siteMapIndex); err != nil {
			fmt.Println("Error")
		}
		for _, sitemap := range siteMapIndex.SiteMaps {
			siteMapUrlArr = append(siteMapUrlArr, sitemap.Loc)
		}
		return siteMapUrlArr, err
	}
}

// pull from here https://gist.github.com/james2doyle/e2f05b5756e4ee46848a8d987405f152
func getXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("read body: %v", err)
	}

	return data, nil
}
