package RemoveSubFoldersfromtheFilesystem

import (
	"path/filepath"
	"sort"
)

// SortBy the given folders by length.
type SortBy []string

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func removeSubfolders(folder []string) []string {
	sort.Sort(SortBy(folder))

	baseFolder := map[string]bool{}
	for i, f := range folder {
		for f != "/" {
			if baseFolder[f] == true {
				break
			}
			f = filepath.Dir(f)
		}
		if baseFolder[f] != true {
			baseFolder[folder[i]] = true
		}
	}

	bfolder := []string{}
	for bf := range baseFolder {
		bfolder = append(bfolder, bf)
	}
	return bfolder
}
