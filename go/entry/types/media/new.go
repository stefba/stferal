package media

import (
	"fmt"
	"stferal/go/entry/helper"
	"stferal/go/entry/types/media/text"
)

func NewMedia(path string) (interface{}, error) {
	switch helper.FileType(path) {
	case "text":
		return text.NewText(path)
	}
	return nil, fmt.Errorf("invalid entry type: %v", path)
}
