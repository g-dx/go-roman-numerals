package webscrape

import (
	"net/http"
	"regexp"
	"fmt"
	"errors"
	"strings"
	"io/ioutil"
)

func GetTime(zone string) (s string, err error) {

	// Get content
	resp, err := http.Get("http://tycho.usno.navy.mil/cgi-bin/timer.pl")
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	// Read response
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	// Match zone
	matches := regexp.MustCompile("<BR>(.*? ([A-Z]+))\t").FindAllSubmatch(buf, -1)
	for _, match := range matches {
		if string(match[2]) == zone {
			return strings.TrimSpace(string(match[1])), nil
		}
	}
	return s, errors.New(fmt.Sprintf("No zone matching '%s' found in response", zone))
}
