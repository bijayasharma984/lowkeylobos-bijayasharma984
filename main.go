package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Items        []Metadata `json:"Items"`
	Count        int        `json:"Count"`
	ScannedCount int        `json:"ScannedCount"`
}

type Metadata struct {
	ResourceType     string   `json:"resourceType"`
	ItemType         string   `json:"itemType"`
	Title            string   `json:"title,omitempty"`
	Description      string   `json:"description"`
	MetadataLanguage string   `json:"metadataLanguage,omitempty"`
	Genres           []string `json:"genres,omitempty"`
	Seasons          []struct {
		SeasonNumber int    `json:"seasonNumber"`
		ResourceType string `json:"resourceType"`
	} `json:"seasons,omitempty"`
	ParentalRating string `json:"parentalRating,omitempty"`
	Augmentation   struct {
		Constraints struct {
			IsRecordable bool `json:"isRecordable"`
		} `json:"constraints"`
	} `json:"augmentation,omitempty"`
	Categories         []string `json:"categories,omitempty"`
	ProgramAttribution string   `json:"programAttribution,omitempty"`
	ResourceID         string   `json:"resourceId"`
	ContentType        string   `json:"contentType,omitempty"`
	DisplayTitle       string   `json:"displayTitle,omitempty"`
	OriginalAirDate    string   `json:"originalAirDate,omitempty"`
	ReleaseYear        int      `json:"releaseYear,omitempty"`
	MajorChannelNumber int      `json:"majorChannelNumber,omitempty"`
	Name               string   `json:"name,omitempty"`
	CcID               string   `json:"ccId,omitempty"`
	MarketID           string   `json:"marketId,omitempty"`
}

var dummyData []Metadata

func setDummyData(data string) []Metadata {
	if err := json.Unmarshal([]byte(data), &dummyData); err != nil {
		panic(err)
	}
	return dummyData
}

var dummyDataString = string(`[
  {
    "resourceType": "SERIES",
    "itemType": "VIDEO_CONTENT",
    "title": "Magic of Disney's Animal Kingdom",
    "description": "An all-access pass to explore Disney's Animal Kingdom Theme Park, Disney's Animal Kingdom Lodge and The Seas With Nemo & Friends at EPCOT, paying tribute to an array of more than 5,000 animals and their animal care experts who keep things running.",
    "metadataLanguage": "en",
    "genres": [
      "Animals",
      "Documentary",
      "Nature"
    ],
    "seasons": [
      {
        "seasonNumber": 1,
        "resourceType": "SEASON"
      }
    ],
    "parentalRating": "TVPG",
    "augmentation": {
      "constraints": {
        "isRecordable": true
      }
    },
    "categories": [
      "TVShow"
    ],
    "programAttribution": "Nat Geo WILD",
    "resourceId": "d68347e7-f18a-45dc-86bc-f677caa90adb"
  },
  {
    "resourceType": "CONTENT",
    "itemType": "VIDEO_CONTENT",
    "contentType": "MOVIE",
    "title": "Magic Mike XXL",
    "displayTitle": "Magic Mike XXL",
    "description": "It's been three years since Mike Lane's (Channing Tatum) retirement from stripping, but the former dancer misses the excitement and feeling of being on stage. Most of all, though, he misses the best friends that he ever had, the crew known as the Kings of Tampa. Opportunity comes knocking when the guys look him up as they travel to Myrtle Beach, S.C., for a stripper convention. With the promise of outrageous fun, a reinvigorated Mike can't resist the chance to join in and take it off again.",
    "metadataLanguage": "en",
    "parentalRating": "R",
    "genres": [
      "Comedy drama"
    ],
    "originalAirDate": "2015-07-01",
    "releaseYear": 2015,
    "categories": [
      "Movies"
    ],
    "resourceId": "3b36bf5a-5b3a-442b-b0ba-7ea48bee7a1d"
  },
  {
    "resourceType": "CHANNEL",
    "itemType": "VIDEO_CHANNEL",
    "majorChannelNumber": 431,
    "name": "Peru Magico",
    "description": "Peru Magico es un canal 100% peruano, las 24 horas del dia, con la mejor programacion de noticias, entretenimiento, deportes y futbol de los canales PlusTV, CMD y Canal N.",
    "ccId": "71104",
    "marketId": "0",
    "resourceId": "3b931c74-ee6e-49e6-9b2e-3cf41e7fcbdd"
  }
]`)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the LOWKEYLOBOS!")
}

func BasicAuthentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		fmt.Println("username: ", user)
		fmt.Println("password: ", pass)

		if !ok || !checkAuth(user, pass) {
			w.Header().Set("WWW-Authentikate", `Basic realm="Please enter username and password for this site"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized.\n"))
			return
		}

		handler(w, r)
	}
}

func checkAuth(username string, password string) bool {
	return username == "lowkeylobos2022" && password == "bijaya-sharma"
}

func GetAllMetadata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	w.Header().Set("Content-Type", "application/json")
	getAllMetadataFromService()
	json.NewEncoder(w).Encode(dummyData)
}

// Demo how we can call internal service
func getAllMetadataFromService() []Metadata {
	resp, err := http.Get("https://7ir47r3nq5.execute-api.us-east-2.amazonaws.com/metadata")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var response Response
	json.Unmarshal(bodyBytes, &response)

	if len(response.Items) > 0 {

		// for demo purpose only
		dummyData = response.Items
	}

	return dummyData
}

func GetMetaData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["resourceId"]

	for _, metadata := range dummyData {
		if metadata.ResourceID == key {
			fmt.Println()
			json.NewEncoder(w).Encode(metadata)
		}
	}
}

func CreateOrUpdateMetadata(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var metadata Metadata
	json.Unmarshal(reqBody, &metadata)

	dummyData = append(dummyData, metadata)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dummyData)
}

func DeleteMetadata(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["resourceId"]

	for index, metadata := range dummyData {
		if metadata.ResourceID == id {
			dummyData = append(dummyData[:index], dummyData[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	myRouter.Handle("/metadata", BasicAuthentication(http.HandlerFunc(GetAllMetadata))).Methods("GET")
	myRouter.HandleFunc("/metadata/{resourceId}", BasicAuthentication(http.HandlerFunc(GetMetaData))).Methods("GET")
	myRouter.HandleFunc("/metadata", BasicAuthentication(http.HandlerFunc(CreateOrUpdateMetadata))).Methods("POST")
	myRouter.HandleFunc("/metadata/{resourceId}", BasicAuthentication(http.HandlerFunc(DeleteMetadata))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	setDummyData(dummyDataString)
	handleRequests()
}
