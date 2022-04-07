package indexing

import (
	"my-fabric-offchain-storage/indexing/kvstore"
	"my-fabric-offchain-storage/indexing/mpt"
)

type IndexingManager struct {
	//repo -> trie
	tries map[string]mpt.Trie
}

func NewManager() *IndexingManager {
	var man IndexingManager
	man.tries = make(map[string]mpt.Trie)
	return &man
}

func (manager *IndexingManager) GetTrieWithRepoName(repo string) (*mpt.Trie, error) {
	if val, ok := manager.tries[repo]; ok {
		return &val, nil
	} else {
		kv, err := kvstore.NewLevelDB("./test")
		if err != nil {
			return nil, err
		}
		trie := mpt.New(nil, kv)
		manager.tries[repo] = *trie
		return trie, nil
	}
}
