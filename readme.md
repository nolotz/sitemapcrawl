# sitemapcrawl

Resolve all URL's from sitemap and sitemap index files.

```go
package main

import (
	"context"

	"github.com/nolotz/sitemapcrawl"
)

func MyFlexiFunc() {
	sitemapResolver := sitemapcrawl.NewResolver()
	ctx := context.Background()

	urls, err := sitemapResolver.Resolve(ctx, "https://example.com/sitemap.xml")
	if err != nil {
		panic(err)
	}

	// enjoy bro!
	_ = urls
}
```