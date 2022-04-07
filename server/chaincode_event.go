package server

import "path/filepath"

type ChainCodeEvent struct {
	Repository string `json:"repository"`
	HashDigest string `json:"hashDigest"`
}

func (event *ChainCodeEvent) getRelativePathInFileSystem() string {
	return filepath.Join(event.Repository, event.HashDigest)
}

func getSubFolderPrefix2Chars(hashDigest string) string {
	prefix := hashDigest[:2]
	return filepath.Join(prefix, hashDigest)
}
