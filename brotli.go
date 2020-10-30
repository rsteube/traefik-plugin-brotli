// Package brotli contains brotli compression middleware
package traefik_plugin_brotli

import (
	"context"
	"log"
	"net/http"

	"github.com/rsteube/traefik-plugin-brotli/pkg/brotlihandler"
)

// Config the plugin configuration.
type Config struct {
	Test string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type Brotli struct {
	next http.Handler
	name string
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Brotli{
		next: brotlihandler.CompressHandler(next),
		name: name,
	}, nil
}

func (e *Brotli) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	log.Println("plugin executed")
	e.next.ServeHTTP(rw, req)
}
