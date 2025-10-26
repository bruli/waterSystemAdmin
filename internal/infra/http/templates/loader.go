package templates

import (
	"embed"
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/flosch/pongo2"
)

type EmbedLoader struct {
	fs       embed.FS
	basePath string
	ext      string
}

func NewEmbedLoader(fs embed.FS, basePath string) pongo2.TemplateLoader {
	return &EmbedLoader{
		fs:       fs,
		basePath: basePath,
		ext:      ".html",
	}
}

func (l *EmbedLoader) Abs(base, name string) string {
	if path.IsAbs(name) {
		return name
	}
	return path.Join(path.Dir(base), name)
}

func (l *EmbedLoader) Get(name string) (io.Reader, error) {
	fullPath := path.Join(l.basePath, name)
	data, err := l.fs.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("plantilla no trobada: %s", fullPath)
	}
	return strings.NewReader(string(data)), nil
}
