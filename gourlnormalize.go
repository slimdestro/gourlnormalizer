package normalizer

import (
	"fmt"
	"net/url"
	"strings"
	"golang.org/x/net/idna"
	"regexp"
)

var (
	Ports = map[string]int{
		"http":  80,
		"https": 443,
		"ftp":   21,
	}
)

func CountUniqueNormalizedUrls(urls []string) int { 
	output := make(map[string]int)

	if len(urls) == 0{
		return 0
	}

	for _, x := range urls{	
		normalizedText, _ := NormalizeURL(x)
		output[normalizedText] = 1
	}

	return len(output)
}

func CountUniqueNormalizedUrlsPerTopLevelDomain(urls []string) map[string]int {  
	output := make(map[string]int)

	if len(urls) == 0{
		return output
	}

	for _, x := range urls{  
		output[fetchTLD(x)] = output[fetchTLD(x)] + 1	 
	}

	return output
}

// extracts tld from url
func fetchTLD(domain string) string {
	pattern, _ := regexp.Compile(`[^.]*\.[^.]{2,3}(?:\.[^.]{2,3})?$`) 
	replacer := strings.NewReplacer("http://","", "https://", "")

	return replacer.Replace(pattern.FindString(domain))
}


// Separate metrhod so that it can be used by both functions
func NormalizeURL(s string) (string, error) {
	s = strings.TrimSpace(s)

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	if u.Scheme == "" { 
		u, err = url.Parse("http://" + s)
		if err != nil {
			return s, err
		}
	}


	if strings.HasPrefix(s, "//") {
		s = "http:" + s
	}

	p, ok := Ports[u.Scheme]
	if ok {
		u.Host = strings.TrimSuffix(u.Host, fmt.Sprintf(":%d", p))
	}

	got, err := idna.ToUnicode(u.Host)
	if err != nil {
		return got, err
	} else {
		u.Host = got
	}

	u.Host = strings.TrimPrefix(u.Host, "www.")

	v := u.Query()
	u.RawQuery = v.Encode()
	u.RawQuery, _ = url.QueryUnescape(u.RawQuery)

	h := u.String()
	h = strings.TrimSuffix(h, "/")

	return h, nil
}
