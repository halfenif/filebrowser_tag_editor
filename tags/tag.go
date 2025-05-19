package tags

import (
	"fmt"
	"log"

	"github.com/bogem/id3v2"
	"github.com/filebrowser/filebrowser/v2/files"
)

type TagInfo struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetTagInfo(file *files.FileInfo) (*TagInfo, error) {

	tag, err := id3v2.Open(file.RealPath(), id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal(err)
	}
	defer tag.Close()

	fmt.Println(tag.Artist())

	return &TagInfo{
		Artist: tag.Artist(),
		Title:  tag.Title(),
	}, nil
}
