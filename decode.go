package gfwatch

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func reverseHost(host string) string {
	r := []rune(host)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	if len(r) > 0 && r[0] == '.' {
		r = r[1:]
	}

	return string(r)
}

func (gfw *GfWatch) Decode(r io.Reader) error {

	var (
		line            string
		lineArray       []string
		domainType      string
		domainTypeArray []string
		flag            []bool
		flagBase        bool
		reverse         string
		i               int
		typ             int
		err             error
	)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line = scanner.Text()
		lineArray = strings.Split(line, "|")
		if len(lineArray) != 3 {
			continue
		}

		domainType = lineArray[1][1:len(lineArray[1])-1]
		domainTypeArray = strings.Split(domainType, ", ")
		flag = make([]bool, DomainTypeMax)
		for i = 0; i < len(domainTypeArray); i++ {
			typ, err = strconv.Atoi(domainTypeArray[i])
			if err != nil {
				continue
			}
			if typ >= DomainTypeMax || typ < 0 {
				continue
			}
			flag[typ] = true
		}

		flagBase = false
		if flag[BaseDomainS] {
			gfw.suffixTrie.Set(strings.Split(lineArray[2]+"*", ""), struct{}{})
			flagBase = true
		} else if flag[BaseDomainDS] {
			gfw.suffixTrie.Set(strings.Split(lineArray[2]+".*", ""), struct{}{})
			flagBase = true
		}

		if flag[SBaseDomain] {
			reverse = reverseHost(lineArray[2])
			gfw.prefixTrie.Set(strings.Split(reverse+"*", ""), struct{}{})
			flagBase = true
		} else if flag[SDBaseDomain] {
			reverse = reverseHost(lineArray[2])
			gfw.prefixTrie.Set(strings.Split(reverse+".*", ""), struct{}{})
			flagBase = true
		}
		if flagBase {
			gfw.gfMap[lineArray[2]] = struct{}{}
		} else {
			gfw.gfMap[lineArray[0]] = struct{}{}
		}
	}

	return nil
}

func FromFile(path string) (*GfWatch, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gfw := New()
	if err := gfw.Decode(f); err != nil {
		return nil, err
	}
	return gfw, nil
}