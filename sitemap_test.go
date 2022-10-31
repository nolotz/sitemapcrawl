package sitemapcrawl_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nolotz/sitemapcrawl"
)

func TestResolve(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/sitemap.xml":
			writer.Write([]byte(`<sitemapindex><sitemap><loc>http://` + request.Host + `/sitemap_2.xml</loc></sitemap></sitemapindex>`))
			return
		case "/sitemap_2.xml":
			writer.Write([]byte(`<urlset><url><loc>TEST</loc></url></urlset>`))
		default:
			writer.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()

	sitemapResolver := sitemapcrawl.NewResolver()

	urls, err := sitemapResolver.Resolve(context.Background(), srv.URL+"/sitemap.xml")
	if err != nil {
		panic(err)
	}

	assert.Len(t, urls, 1)
	assert.Equal(t, "TEST", urls[0].Loc)
}
