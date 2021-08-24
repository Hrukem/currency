// Package request contain function for send request for other website
package request

import (
	"currency/database"
	"currency/errorhandler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Request function send request to website www.currencylayer.com
// and return result in the form of slice of strings
func Request(date string, currencies string, typeRequest string) ([]string, error) {
	resp, err := http.Get(
		os.Getenv("WEBSITE") +
			 typeRequest +
			"?access_key=" + os.Getenv("API_KEY") +
			"&date=" + date +
			"&currencies=" + currencies,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println("Not close Body in request.Request(), line #30", err)
		}
	}()

	database.InsertLogAPI(currencies, typeRequest, time.Now())
	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	errorhandler.CheckError(err)

	answer := ParseRequest(result)

	return answer, nil
}

// ParseRequest function convert map from answer request into slice of string
func ParseRequest(answer map[string]interface{}) []string {
	curr := answer["quotes"]
	currMap := curr.(map[string]interface{})

	var s []string
	for k, v := range currMap {
		switch i := v.(type) {
		case int:
			s = append(s, k+":"+strconv.Itoa(i))
		case string:
			s = append(s, k+":"+i)
		case float64:
			s = append(s, k+":"+fmt.Sprintf("%f", i))
		default:
			s = append(s, k+":"+"unknown value")
		}
	}
	return s
}
