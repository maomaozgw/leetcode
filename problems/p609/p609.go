// https://leetcode.com/problems/find-duplicate-file-in-system/

package p609

import (
	"path/filepath"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	var tmp = map[string][]string{}
	for _, item := range paths {
		parse(item, tmp)
	}
	var result [][]string
	for _, v := range tmp {
		if len(v) > 1 {
			result = append(result, v)
		}
	}
	return result
}

func parse(pathInfo string, contentMap map[string][]string) {
	const (
		dirStart         = 1
		fileNameStart    = 2
		fileContentStart = 3
		spliter          = byte(' ')
		contentStart     = byte('(')
		contentEnd       = byte(')')
	)
	var (
		currentStat     = dirStart
		directory       strings.Builder
		currentFileName strings.Builder
		currentContent  strings.Builder
	)
	for i := 0; i < len(pathInfo); i++ {
		c := pathInfo[i]
		switch currentStat {
		case dirStart:
			if c == spliter {
				currentStat = fileNameStart
				continue
			}
			directory.WriteByte(c)
		case fileNameStart:
			if c == contentStart {
				currentStat = fileContentStart
				continue
			}
			currentFileName.WriteByte(c)
		case fileContentStart:
			if c == contentEnd {
				contentMap[currentContent.String()] = append(contentMap[currentContent.String()], filepath.Join(directory.String(), currentFileName.String()))
				currentContent.Reset()
				currentFileName.Reset()
				i++
				currentStat = fileNameStart
				continue
			}
			currentContent.WriteByte(c)
		}
	}
}
