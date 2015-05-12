package pantopoda

// Config holds configuration options for our spider.
type Config struct {
	Verbose bool
	// Routines : Number of concurrent crawl routines to use.
	Routines int
	// PauseBetweenRequests : Pause length between reqquests in ms (per routine)
	PauseBetweenRequests int
	// Http timeout in seconds
	HttpTimeout int
	// Maximum number of links to follow from a single page
	MaxLinksPerPage int
}

// DefaultConfig is the default configuration.
var DefaultConfig Config = Config{
	Routines:             5,
	PauseBetweenRequests: 100,
	Verbose:              true,
	HttpTimeout:          10,
	MaxLinksPerPage:      500,
}
