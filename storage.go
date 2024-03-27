package storage

import (
	"fmt"
	"github.com/fanky5g/ponzu/config"
	"github.com/fanky5g/ponzu/driver"
	"net/http"
	"path/filepath"
)

type client struct {
	s          driver.StaticFileSystemInterface
	storageDir string
}

func New(dir string) (driver.StorageClientInterface, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("failed to get config: %v", err)
	}

	if dir == "" {
		dir = filepath.Join(cfg.Paths.DataDir, "uploads")
	}

	s, err := NewLocalStaticFileSystem(http.Dir(dir))
	if err != nil {
		return nil, err
	}

	return &client{s: s, storageDir: dir}, nil
}
