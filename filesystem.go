package storage

import (
	"github.com/fanky5g/ponzu/driver"
	"net/http"
	"os"
)

func NewLocalStaticFileSystem(dir http.Dir) (driver.StaticFileSystemInterface, error) {
	return justFilesFilesystem{dir}, nil
}

// the code below removes the open directory listing when accessing a URL which
// normally would point to a directory. code from golang-nuts mailing list:
// https://groups.google.com/d/msg/golang-nuts/bStLPdIVM6w/hidTJgDZpHcJ
// credit: Brad Fitzpatrick (c) 2012

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
