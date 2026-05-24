package extractor

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type bookmarkContent struct {
	Type     string            `json:"type"`
	Name     string            `json:"name"`
	URL      string            `json:"url,omitempty"`
	Children []bookmarkContent `json:"children,omitempty"`
}

type bookmarkStructure struct {
	Roots struct {
		BookmarkBar struct {
			Children []bookmarkContent `json:"children"`
		} `json:"bookmark_bar"`
	} `json:"roots"`
}

type brave struct{}

func NewBraveExtractor() Extractor {
	return brave{}
}

func (b brave) Name() string {
	return "Brave"
}

func (b brave) Extract() ([]Bookmark, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(homeDir, "Library", "Application Support", "BraveSoftware", "Brave-Browser", "Default", "Bookmarks")

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var bookmarks bookmarkStructure
	if err := json.Unmarshal(content, &bookmarks); err != nil {
		return nil, err
	}

	return iterateBookmarks(bookmarks.Roots.BookmarkBar.Children), nil
}

func iterateBookmarks(children []bookmarkContent) []Bookmark {
	var result []Bookmark

	for _, child := range children {
		switch child.Type {
		case "url":
			result = append(result, Bookmark{
				Title: child.Name,
				URL:   child.URL,
			})
		case "folder":
			result = append(result, iterateBookmarks(child.Children)...)
		}
	}

	return result
}
