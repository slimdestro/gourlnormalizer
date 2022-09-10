# (+)go-url-normalize[RFC 3986]
#### Normalize URL | Count unique normalized URL | Count unique normalized URL per TLD

[![N|Solid](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/60px-Go_Logo_Blue.svg.png)](https://dev.to/slimdestro)
 
## Installation

Install the package by doing go get:

```sh
go get github.com/slimdestro/gourlnormalizer
```

Three exported methods that you can use in your modules are:

```sh
NormalizeURL(url string) 
CountUniqueNormalizedUrls(url []string)  
CountUniqueNormalizedUrlsPerTopLevelDomain(url []string)   
```

## Example

```sh
package main 

import (
	"fmt"
	"github.com/slimdestro/gourlnormalizer"
)
func main() { 		
	/* calls to NormalizeURL() */
	url := "https://example.com?b=2&a=1"
	normalizedUrl,_ := normalizer.NormalizeURL(url)
	fmt.Println(normalizedUrl)
        // compare both url to see what has changed

	/* calls to CountUniqueUrls() */
	urls__ := []string{"https://example.com?a=1&b=2", "https://example.com?b=2&a=1"}
	fmt.Println(normalizer.CountUniqueNormalizedUrls(urls__))
	
	/* calls to CountUniqueUrlsPerTopLevelDomain() */
	urls___ := []string{"https://tempo.com", "https://example.com", "https://subdomain.example.com"}
	fmt.Println(normalizer.CountUniqueNormalizedUrlsPerTopLevelDomain(urls___))
}
 
```


## Author

[slimdestro(Mukul Mishra)](https://linktr.ee/slimdestro)
