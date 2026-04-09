package application

import (
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/extensions"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/metrics"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/ports"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/relations"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/seo"
)

// PostSubresourceStores groups narrow persistence ports for post sub-resources.
type PostSubresourceStores struct {
	Reader       ports.PostReader
	SEO          seo.PostSEOStore
	Metrics      metrics.PostMetricsStore
	Featured     relations.PostFeaturedImageStore
	SeriesLink   relations.PostSeriesLinkStore
	Coauthors    relations.PostCoauthorsStore
	Gallery      extensions.PostGalleryStore
	Changelog    extensions.PostChangelogStore
	Syndication  extensions.PostSyndicationStore
	Translations extensions.PostTranslationsStore
}

// PostSubresourceStoresFrom adapts a full Repository into store fields (same concrete value).
func PostSubresourceStoresFrom(r ports.Repository) PostSubresourceStores {
	return PostSubresourceStores{
		Reader:       r,
		SEO:          r,
		Metrics:      r,
		Featured:     r,
		SeriesLink:   r,
		Coauthors:    r,
		Gallery:      r,
		Changelog:    r,
		Syndication:  r,
		Translations: r,
	}
}
