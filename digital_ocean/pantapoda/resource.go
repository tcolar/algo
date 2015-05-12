package pantopoda

import (
	"net/url"
)

// PendingResource represents a resource identifed but yet to be crawled
type PendingResource struct {
	Url  *url.URL
	Type ResourceType
}

// Resource is a resource that has been processed / crawled
type Resource struct {
	Status    string // response status code
	Type      ResourceType
	Resources []string // Resources embedded
}

// ResourceType defines what kind of resource a URL points to
type ResourceType string

const (
	UnknownResource ResourceType = "Unknown"
	PageResource                 = "Page"  // A page, will be furtehr parsed
	AssetResource                = "Asset" // An asset, no further arding
)
