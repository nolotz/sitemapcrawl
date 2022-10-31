package sitemapcrawl

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"strings"
)

// Resolver fetch index sitemap and resolves all urls
type Resolver struct {
	HTTP *http.Client
}

func NewResolver() *Resolver {
	return &Resolver{
		HTTP: http.DefaultClient,
	}
}

func (r *Resolver) Resolve(ctx context.Context, sitemap string) ([]URL, error) {
	docBody, err := r.requestDocument(ctx, sitemap)
	if err != nil {
		return nil, err
	}

	root, err := extractRoot(docBody)
	if err != nil {
		return nil, err
	}

	switch root.XMLName.Local {
	case indexXMLNameLocal:
		return r.resolveIndex(ctx, root.Sitemap)
	case urlSetXMLNameLocal:
		return root.URL, nil
	}

	return nil, newUnexpectedTypeError(root.XMLName.Local)
}

func (r *Resolver) resolveIndex(ctx context.Context, indices []URL) ([]URL, error) {
	urls := make([]URL, 0, len(indices))

	for _, url := range indices {
		root, err := r.getRoot(ctx, url.Loc)
		if err != nil {
			return nil, err
		}

		if root.XMLName.Local != urlSetXMLNameLocal {
			return nil, newUnexpectedTypeError(root.XMLName.Local)
		}

		urls = append(urls, root.URL...)
	}

	return urls, nil
}

func (r *Resolver) getRoot(ctx context.Context, sitemap string) (*root, error) {
	docBody, err := r.requestDocument(ctx, sitemap)
	if err != nil {
		return nil, err
	}

	return extractRoot(docBody)
}

func (r *Resolver) requestDocument(ctx context.Context, sitemap string) ([]byte, error) {
	req, err := http.NewRequest("GET", sitemap, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")

	req = req.WithContext(ctx)
	resp, err := r.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if !resp.Uncompressed && strings.HasSuffix(sitemap, ".gz") {
		reader, err := gzip.NewReader(bytes.NewReader(respBody))
		if err != nil {
			return nil, err
		}

		respBody, err = io.ReadAll(reader)
		if err != nil {
			return nil, err
		}
	}

	return respBody, nil
}

type root struct {
	XMLName xml.Name
	Sitemap []URL `xml:"sitemap,omitempty"`
	URL     []URL `xml:"url,omitempty"`
}

func extractRoot(data []byte) (*root, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.Strict = false

	var rt root
	err := decoder.Decode(&rt)
	if err != nil {
		return nil, err
	}

	return &rt, nil
}
