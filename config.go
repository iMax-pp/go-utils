// Copyright (c) 2014 Maxime SIMON. All rights reserved.

package utils

import (
	"bufio"
	"os"
	"strings"
)

func LoadConfig(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	props := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		// Ignore lines without "="
		if !strings.Contains(line, "=") {
			continue
		}
		entry := strings.SplitN(line, "=", 2)
		key := strings.TrimSpace(entry[0])
		val := strings.TrimSpace(entry[1])
		props[key] = val
	}

	return props, scanner.Err()
}
