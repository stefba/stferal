package server

import (
	"stferal/go/entry"
	"stferal/go/entry/types/tree"
	"fmt"
)

type SectionTree struct {
	Public, Private map[string]*tree.Tree
}

type SectionEntries struct {
	Public, Private map[string]entry.Entries
}

func (s *SectionTree) Local(local bool) map[string]*tree.Tree {
	if local {
		return s.Private
	}
	return s.Public
}

func (s *SectionEntries) Local(local bool) map[string]entry.Entries {
	if local {
		return s.Private
	}
	return s.Public
}


// trees

func (s *Server) LoadTrees() error {
	sections := []string{
		"index",
		"graph",
		"about",
		"extra",
	}

	trees := map[string]*SectionTree{}
	recents := map[string]*SectionEntries{}

	for _, section := range sections {
		t, err := tree.ReadTree(s.Paths.Data+"/"+section, nil)
		if err != nil {
			return err
		}

		trees[section] = &SectionTree{
			Private: makeLangs(t),
			Public:  makeLangs(t.Public()),
		}

		recents[section] = &SectionEntries{
			Private: serializeLangs(trees[section].Private),
			Public:  serializeLangs(trees[section].Public),
		}

		fmt.Printf("%v: lang de %v\n", section, len(recents[section].Public["de"]))
		fmt.Printf("%v: lang en %v\n", section, len(recents[section].Public["en"]))
	}

	s.Trees = trees
	s.Recents = recents

	return nil
}

func makeLangs(t *tree.Tree) map[string]*tree.Tree {
	return map[string]*tree.Tree{
		"de": t,
		"en": t.Lang("en"),
	}
}

func serializeLangs(langMap map[string]*tree.Tree) map[string]entry.Entries {
	return map[string]entry.Entries{
		"de": serialize(langMap["de"]),
		"en": serialize(langMap["en"]),
	}
}

func serialize(t *tree.Tree) entry.Entries {
	if t.Section() == "graph" {
		return t.TraverseEntriesReverse()
	}
	return t.TraverseEntries().Exclude().Desc()
}

