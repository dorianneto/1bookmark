package extractor

import "time"

type Bookmark struct {
	URL         string
	Title       string
	FolderPath  string   // slash-joined breadcrumb, e.g. "Dev/Go/Libraries"
	Tags        []string // from Pocket/Instapaper tags, or derived from folder
	Description string
	SavedAt     time.Time // zero if unknown
	ReadAt      time.Time // zero if never read / unknown
	IsRead      bool
}

type Extractor interface {
	Name() string
	Extract() ([]Bookmark, error)
}
