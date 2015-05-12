package pantopoda

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"code.google.com/p/go.net/html"
)

// Fetcher specifies a Fetcher implementations
type Fetcher interface {
	// Return a normalized key for a given url
	UrlKey(u *url.URL) string

	// Fetch a given url and return a resource
	Fetch(*PendingResource) (*Resource, []PendingResource)

	// AllowUrl decides whether a url should be fetched or not
	AllowUrl(*url.URL) bool
}

// DomainFetcher is a Fetcher implementation that only crawls a single domain.
// Exluding subdomains or any other domains.
type DomainFetcher struct {
	Domain string
	Config *Config
}

// AllowUrl : decides wether a given URL should be crawled or not
func (c *DomainFetcher) AllowUrl(u *url.URL) bool {
	if strings.ToLower(u.Host) == strings.ToLower(c.Domain) {
		return true
	}
	return false
}

// Fetch fetches a given resource
// If it's an HTML page then "parse" it to find more links.
func (c *DomainFetcher) Fetch(res *PendingResource) (*Resource, []PendingResource) {
	resType := res.Type
	pending := []PendingResource{}
	resources := []string{}

	status := "500" // In case of issue we wil use a status 500

	if resType == UnknownResource || resType == AssetResource {
		// TODO: maybe we could use mime types ?
		if resType == UnknownResource {
			resType = PageResource
		}
		resp, err := c.head(res.Url.String())
		if err == nil {
			status = resp.Status
			ct := resp.Header["Content-Type"]
			if len(ct) != 0 && !strings.HasPrefix(ct[0], "text/html") {
				resType = AssetResource
			}
		}
	}

	// If an asset, no need to parse it further
	if resType == AssetResource {
		return &Resource{
			Status:    status,
			Type:      resType,
			Resources: resources,
		}, pending
	}

	// Otherwise a page
	u := res.Url
	resp, err := c.get(u.String())
	if err != nil {
		if c.Config.Verbose {
			log.Printf("Failed to fetch %s", u.String())
		}
		return &Resource{Status: status}, []PendingResource{}
	}
	status = resp.Status
	// TODO: If 40x, 50x no point lookng at body ?
	doc, err := html.Parse(resp.Body)
	if err != nil {
		if c.Config.Verbose {
			log.Printf("Failed to read body of %s", u.String())
		}
		return &Resource{Status: status}, []PendingResource{}
	}

	pending = c.findResources(doc, u)
	for _, p := range pending {
		resources = append(resources, c.UrlKey(p.Url))
	}
	return &Resource{
		Status:    status,
		Type:      resType,
		Resources: resources,
	}, pending
}

// findResources looks for more resurces & links in a page body.
func (c *DomainFetcher) findResources(body *html.Node, from *url.URL) []PendingResource {
	pending := []PendingResource{}
	key := c.UrlKey(from)
	var err error
	var to *url.URL
	var typ ResourceType
	nbLinks := 0

	var scan func(*html.Node)
	scan = func(n *html.Node) {
		to = nil
		data := strings.ToLower(n.Data)
		// TODO: Should probably use a visitor patterns as there might be many type of links
		if n.Type == html.ElementNode && data == "a" {
			// Links such as <a href="...">
			for _, attr := range n.Attr {
				if strings.ToLower(attr.Key) == "href" {
					to, err = from.Parse(strings.Trim(attr.Val, " \t\r\n"))
					if err != nil {
						log.Fatal(err)
					}
					typ = UnknownResource
				}
			}
		} else if n.Type == html.ElementNode && data == "link" {
			// Assets such as <link href="...">
			for _, attr := range n.Attr {
				if strings.ToLower(attr.Key) == "href" {
					to, err = from.Parse(strings.Trim(attr.Val, " \t\r\n"))
					if err != nil {
						log.Fatal(err)
					}
					typ = AssetResource
				}
			}
		} else if n.Type == html.ElementNode && (data == "script" || data == "img") {
			// Assets such as <script src="..."> or <img src="...">
			for _, attr := range n.Attr {
				if strings.ToLower(attr.Key) == "src" {
					to, err = from.Parse(strings.Trim(attr.Val, " \t\r\n"))
					if err != nil {
						log.Fatal(err)
					}
					typ = AssetResource
				}
			}
		}

		// If we found a link, add it to pending resources
		if to != nil {
			nbLinks++
			toKey := c.UrlKey(to)
			if c.AllowUrl(to) && toKey != key { // no links to self
				res := PendingResource{
					Url:  to,
					Type: typ,
				}
				pending = append(pending, res)
			}
		}
		for child := n.FirstChild; child != nil; child = child.NextSibling {
			if nbLinks > c.Config.MaxLinksPerPage {
				break
			}
			scan(child)
		}
	}
	scan(body)
	return pending
}

// UrlKey returns a key for what we consider a unique web resource.
func (c *DomainFetcher) UrlKey(u *url.URL) string {
	return strings.ToLower(u.Host) +
		strings.ToLower(u.Path) + u.RawQuery
}

func (c *DomainFetcher) dialTimeout(network, addr string) (net.Conn, error) {
	timeout := time.Duration(c.Config.HttpTimeout) * time.Second
	return net.DialTimeout(network, addr, timeout)
}

// get makes a GET request using our client with a custom User-Agent.
// Note: Go defaut client follows up to 10 redirects
func (c *DomainFetcher) get(url string) (*http.Response, error) {
	transport := http.Transport{
		Dial: c.dialTimeout,
	}

	client := &http.Client{
		Transport: &transport,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	return client.Do(req)
}

// head makes a HEAD request using our client with a custom User-Agent.
// Note: Go default client follows up to 10 redirects
func (c *DomainFetcher) head(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	return client.Do(req)
}
