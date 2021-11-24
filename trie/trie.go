package trie

import "errors"

type trieChildren map[string]*Trie

type Trie struct {
	Entry      interface{}
	SplatEntry interface{}
	Children   trieChildren
}

func NewTrie() *Trie {
	return &Trie{
		Children: make(trieChildren),
	}
}

func (t *Trie) Get(path []string) (entry interface{}, ok bool) {
	if len(path) == 0 {
		return t.getentry()
	}

	key := path[0]
	newpath := path[1:]

	res, ok := t.Children[key]
	if ok {
		entry, ok = res.Get(newpath)
	}

	if entry == nil && t.SplatEntry != nil {
		entry = t.SplatEntry
		ok = true
	}

	return
}

func (t *Trie) getAll() (entry []interface{}) {

	if value, ok := t.getentry(); ok {
		entry = append(entry, value)
	}

	for _, child := range t.Children {
		childEntry := child.getAll()
		entry = append(entry, childEntry...)
	}
	return
}

func (t *Trie) GetAll(path []string) (entry []interface{}, err error) {

	if len(path) == 0 {
		if value, ok := t.getentry(); ok {
			entry = append(entry, value)
		}
		return
	}

	if path[0] == "*" {
		if len(path) != 1 {
			return entry, errors.New("* should be last element")
		}
		for _, child := range t.Children {
			childEntry := child.getAll()
			entry = append(entry, childEntry...)
		}
		return
	}

	key := path[0]
	newpath := path[1:]

	res, ok := t.Children[key]
	if ok {
		if entry, err = res.GetAll(newpath); err != nil {
			return
		}
	}

	if t.SplatEntry != nil {
		entry = append(entry, t.SplatEntry)
	}

	return
}

func (t *Trie) Set(path []string, value interface{}) error {
	if len(path) == 0 {
		t.setentry(value)
		return nil
	}

	if path[0] == "*" {
		if len(path) != 1 {
			return errors.New("* should be last element")
		}
		t.SplatEntry = value
	}

	key := path[0]
	newpath := path[1:]

	res, ok := t.Children[key]
	if !ok {
		res = NewTrie()
		t.Children[key] = res
	}

	return res.Set(newpath, value)
}

func (t *Trie) Delete(path []string) (empty bool) {

	if len(path) == 0 {
		t.setentry(nil)
		return len(t.Children) == 0
	}

	if path[0] == "*" {
		if len(path) != 1 {
			return false
		}
		t.Children = make(trieChildren)
		return t.Entry == nil
	}

	key := path[0]
	newpath := path[1:]

	res, ok := t.Children[key]
	if ok {
		if res.Delete(newpath) {
			delete(t.Children, key)
		}
	}

	return t.Entry == nil && len(t.Children) == 0
}

func (t *Trie) setentry(value interface{}) {
	t.Entry = value
}

func (t *Trie) getentry() (entry interface{}, ok bool) {
	return t.Entry, t.Entry != nil
}
