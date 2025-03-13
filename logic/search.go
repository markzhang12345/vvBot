package logic

import (
	"math/rand"
	"path/filepath"
	"strings"
	"time"
)

func searchImageByKeyword(keyword string, filenames []string) string {
	const imageDir = "./vvsource"

	keyword = strings.ToLower(strings.TrimSpace(keyword))
	
	// 设置随机选择
	rand.Seed(time.Now().UnixNano())
	
	if keyword == "" {
		randomIndex := rand.Intn(len(filenames))
		return filepath.Join(imageDir, filenames[randomIndex])
	}
	
	// 模糊搜索
	var bestMatches []string
	var partialMatches []string
	
	for _, filename := range filenames {
		lowerName := strings.ToLower(filename)
		
		if lowerName == keyword {
			bestMatches = append(bestMatches, filename)
		} else if strings.Contains(lowerName, keyword) {
			partialMatches = append(partialMatches, filename)
		}
	}
	
	if len(bestMatches) > 0 {
		randomIndex := rand.Intn(len(bestMatches))
		return filepath.Join(imageDir, bestMatches[randomIndex])
	}
	
	if len(partialMatches) > 0 {
		randomIndex := rand.Intn(len(partialMatches))
		return filepath.Join(imageDir, partialMatches[randomIndex])
	}
	
	randomIndex := rand.Intn(len(filenames))
	return filepath.Join(imageDir, filenames[randomIndex])
}