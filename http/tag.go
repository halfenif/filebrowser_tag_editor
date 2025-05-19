package http

import (
	"fmt"
	"net/http"

	"github.com/filebrowser/filebrowser/v2/files"
	"github.com/filebrowser/filebrowser/v2/tags"
)

var tagHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	// if !d.user.Perm.Rename {
	// 	return http.StatusAccepted, nil
	// }

	file, err := files.NewFileInfo(&files.FileOptions{
		Fs:         d.user.Fs,
		Path:       r.URL.Path,
		Modify:     d.user.Perm.Modify,
		Expand:     true,
		ReadHeader: d.server.TypeDetectionByHeader,
		Checker:    d,
		Content:    true,
	})
	if err != nil {
		fmt.Println("tagHandler.withUser.files.NewFileInfo", err)
		return errToStatus(err), err
	}

	// fmt.Println("-------------------------------------------")
	// fmt.Println("http.Request", r)
	// fmt.Println("-------------------------------------------")
	// fmt.Println("data", d)
	fmt.Println("-------------------------------------------")
	fmt.Println("file", file)
	fmt.Println("-------------------------------------------")

	tag, err := tags.GetTagInfo(file)
	if err != nil {
		fmt.Println("tags.GetTagInfo", err)
		return http.StatusInternalServerError, err
	}
	fmt.Println("tag", tag)

	return renderJSON(w, r, tag)
})

// func tagFileHandler(w http.ResponseWriter, r *http.Request, file *files.FileInfo) (int, error) {
// 	// if !files.IsSupportedTag(file.Name) {
// 	// 	return http.StatusBadRequest, nil
// 	// }

// 	fd, err := file.Fs.Open(file.Path)
// 	if err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	defer fd.Close()

// 	setContentDisposition(w, r, file)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Cache-Control", "private")

// 	if _, err := w.Write([]byte(file.Name)); err != nil {
// 		return http.StatusInternalServerError, err
// 	}

// 	return http.StatusOK, nil
// }
