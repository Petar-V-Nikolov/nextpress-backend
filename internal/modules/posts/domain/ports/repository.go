package ports

import (
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/extensions"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/metrics"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/relations"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/seo"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/series"
)

// Repository is the composed persistence port for the posts module.
// Prefer depending on smaller interfaces (PostReader, PostSEOStore, …) in new code.
type Repository interface {
	PostReader
	PostWriter
	relations.PostTaxonomyWriter
	seo.PostSEOStore
	metrics.PostMetricsStore
	relations.PostFeaturedImageStore
	relations.PostSeriesLinkStore
	relations.PostCoauthorsStore
	extensions.PostGalleryStore
	extensions.PostChangelogStore
	extensions.PostSyndicationStore
	extensions.PostTranslationsStore
	series.SeriesRepository
	extensions.TranslationGroupRepository
}
