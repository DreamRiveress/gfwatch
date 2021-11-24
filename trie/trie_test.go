package trie

import (
	"strings"
	"testing"
)

func TestDNSMatch(t *testing.T) {

	trie := NewTrie()
	if trie == nil {
		t.Error()
	}

	trie.Set(strings.Split("com.baidu.www", "."), "1")
	trie.Set(strings.Split("com.baidu.*", "."), "2")
	trie.Set(strings.Split("co.baidu.weidu", "."), "2")

	_, ok := trie.Get(strings.Split("com.baidu.www", "."))
	if !ok {
		t.Error()
	}

	_, ok = trie.Get(strings.Split("com.baidu.100", "."))
	if !ok {
		t.Error()
	}

	match, _ := trie.Get(strings.Split("com.1baidu.100", "."))
	if match != nil {
		t.Error()
	}
}

func TestGetAll(t *testing.T) {

	trie := NewTrie()
	if trie == nil {
		t.Error()
	}

	trie.Set(strings.Split("yyy.test.com/live/yyytest/*", "/"), "1a")
	trie.Set(strings.Split("yyy.test.com/*", "/"), "ccc")
	trie.Set(strings.Split("yyy.test.com/live/yyst", "/"), "1abc")
	trie.Set(strings.Split("yyy.test.com/live/yyst/zzz", "/"), "1add")
	trie.Set(strings.Split("yyy.test.com/live123/yyst/zzz", "/"), "zzz")

	trie.Set(strings.Split("yyy.test.com/live", "/"), "vvv")
	trie.Set(strings.Split("www.baidu.com/test123", "/"), "www")
	trie.Set(strings.Split("*", "/"), "ddd")

	v, err := trie.GetAll(strings.Split("yyy.test.com/*", "/"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(v) != 7 {
		t.Error()
	}
}

func TestDelete(t *testing.T) {

	trie := NewTrie()
	if trie == nil {
		t.Error()
	}

	trie.Set(strings.Split("yyy.test.com/live/yyytest/*", "/"), "1a")
	trie.Set(strings.Split("yyy.test.com/*", "/"), "ccc")
	trie.Set(strings.Split("yyy.test.com/live/yyst", "/"), "1abc")
	trie.Set(strings.Split("yyy.test.com/live/yyst/zzz", "/"), "1add")
	trie.Set(strings.Split("yyy.test.com/live123/yyst/zzz", "/"), "zzz")

	trie.Set(strings.Split("yyy.test.com/live", "/"), "vvv")
	trie.Set(strings.Split("www.baidu.com/test123", "/"), "www")
	trie.Set(strings.Split("*", "/"), "ddd")

	v, err := trie.GetAll(strings.Split("*", "/"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(v) != 8 {
		t.Error()
	}

	trie.Delete(strings.Split("yyy.test.com/live/*", "/"))

	v, err = trie.GetAll(strings.Split("*", "/"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(v) != 5 {
		t.Error()
	}
}
