package main

import (
	"encoding/json"
	"regexp"
)

var regex = regexp.MustCompile(`data-vue2='(.*?)'`)

func ParsePageContent(content string) ([]Thread, error) {
	matches := regex.FindAllStringSubmatch(content, -1)

	items := make([]Thread, 0)

	for _, match := range matches {
		jsonStr := match[1]

		jsonDict := struct {
			Name string `json:"name"`
		}{}
		err := json.Unmarshal([]byte(jsonStr), &jsonDict)
		if err != nil {
			return nil, err
		}

		if jsonDict.Name != "ThreadMainListItemNormalizer" {
			continue
		}

		item := ThreadMainListItemNormalizer{}
		err = json.Unmarshal([]byte(jsonStr), &item)
		if err != nil {
			return nil, err
		}

		items = append(items, item.Props.Thread)

	}

	return items, nil
}
