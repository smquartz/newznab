package entry

// Content describes the content of an newznab entry; that is, what is actually
// in the files
type Content interface {
	// actual "release" (has specific meaning), if applicable
	Release() Release
	// categories of the content
	Categories() []Category
	// title of the content
	Title() string
}

// GenericContent describes information about content that is common to all
// content types.  It is intended to be embedded by Content implementations
type GenericContent struct {
	// actual "release" (has specific meaning), if applicable
	release    Release
	categories []Category
}

// Release returns the "release" (has specific meaning) for specific content
func (g GenericContent) Release() Release {
	return g.release
}

// Categories returns the newznab categories that content belongs to
func (g GenericContent) Categories() []Category {
	return g.categories
}
