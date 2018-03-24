package entry

// Release describes a "release", with specific meaning
// this will be moved to its own package at some point for release name parsing
type Release struct {
	// the release group
	Group string
	// the release name
	Name string
}
