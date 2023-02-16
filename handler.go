package reacthandler

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Handler struct {
	builtFile fs.FS
}

func NewHandler(static embed.FS) *Handler {
	f, err := fs.Sub(static, "build")
	if err != nil {
		panic(err)
	}
	return &Handler{builtFile: f}
}

const DefaultIndexHTML = "index.html"

// handleStatic http handler
func (h *Handler) handleStatic(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = DefaultIndexHTML
	}
	file, err := h.builtFile.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// reopen index.html
		file, _ = h.builtFile.Open(DefaultIndexHTML)
	}
	contentType := mime.TypeByExtension(filepath.Ext(path))
	w.Header().Set("Content-Type", contentType)
	stat, err := file.Stat()
	if err == nil && stat.Size() > 0 {
		w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	}
	io.Copy(w, file)
}
