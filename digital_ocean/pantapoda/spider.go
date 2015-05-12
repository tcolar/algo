package pantopoda

import (
	"encoding/json"
	"io"
	"log"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

var Version = "0.0.1"

var UserAgent = "Pantapoda " + Version

// Spider represents the spider/crawler itself.
// It is used to crawl some resources.
type Spider struct {
	Config  *Config
	Fetcher Fetcher

	// Hash stores the Craweled data (in memory in this implemtation)
	Hash map[string]Resource

	// queue is a channel that holds the items pending to be crawled
	queue chan *PendingResource
	// pending tracks the number of pending crawling requests so we know when we are done.
	pending int64
	// mutex used to protect concurrent R/W access to the hash
	hashMutex sync.Mutex
}

// NewSpider creates a new spider initialized with the given config.
func NewSpider(config *Config) *Spider {
	return &Spider{
		Hash:   map[string]Resource{},
		Config: config,
		queue:  make(chan *PendingResource, 100),
	}
}

// CrawlDomain starts crawling a single domain. (sudomains excluded)
// For Example given the domain "digitalocean.com" it will crawl digitalocean.com
// but *NOT* foo.digitalocean.com or any other domains
func (s *Spider) CrawlDomain(domain string) {
	s.Fetcher = &DomainFetcher{
		Config: s.Config,
		Domain: domain,
	}
	base, err := url.Parse("http://" + domain + "/")
	if err != nil {
		log.Fatal(err)
	}
	s.Fetcher.UrlKey(base)
	// Start the Fetcher routines
	for i := 0; i != s.Config.Routines; i++ {
		go s.worker(i)
	}
	// Add the domain root as the base URL to start from
	s.enQueue(PendingResource{
		Url:  base,
		Type: PageResource,
	})
	// Wait until all work has been done
	for s.pending > 0 {
		time.Sleep(100 * time.Millisecond)
	}
}

// enQueue queues a new resource to be crawled
func (s *Spider) enQueue(res PendingResource) {
	if s.Config.Verbose {
		log.Printf("Enqueue %s", res.Url)
	}
	atomic.AddInt64(&s.pending, 1)
	s.queue <- &res
}

// worker handles an individual Fetcher worker
// All workers keep working until all work has been completed (s.Pending = 0)
func (s *Spider) worker(id int) {
	delay := time.Duration(s.Config.PauseBetweenRequests) * time.Millisecond
	verbose := s.Config.Verbose
	if verbose {
		log.Printf("Starting worker %d", id)
	}
	for {
		if s.pending == 0 {
			return
		}
		select {
		case res, ok := <-s.queue:
			if ok {
				if verbose {
					log.Printf("Worker %d fetching %s", id, res.Url)
				}
				// process it
				s.processResource(res)
				// decrement counter
				atomic.AddInt64(&s.pending, -1)
				time.Sleep(delay)
			}
		}
	}
}

// Process a web resource
func (s *Spider) processResource(res *PendingResource) {
	u := res.Url

	resource, pending := s.Fetcher.Fetch(res)

	for _, res := range pending {
		key := s.Fetcher.UrlKey(res.Url)
		s.hashMutex.Lock()
		if _, exists := s.Hash[key]; !exists {
			s.Hash[key] = *resource
			// Queue a new resource to be processed
			go s.enQueue(res)
		}
		s.hashMutex.Unlock()
	}
	// Add the page to the map
	s.hashMutex.Lock()
	s.Hash[s.Fetcher.UrlKey(u)] = *resource
	s.hashMutex.Unlock()
}

// ToJson outputs the Crawled data (in Hash) to the writer in Json format.
func (s *Spider) ToJson(w io.Writer) error {
	b, err := json.MarshalIndent(s.Hash, "", "    ")
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
