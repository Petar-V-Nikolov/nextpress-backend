package domain

// Repository is the composed persistence port for the posts module.
// Prefer depending on smaller interfaces (PostReader, PostSEOStore, …) in new code.
type Repository interface {
	PostReader
	PostWriter
	PostTaxonomyWriter
	PostSEOStore
	PostMetricsStore
	PostFeaturedImageStore
	PostSeriesLinkStore
	PostCoauthorsStore
	PostGalleryStore
	PostChangelogStore
	PostSyndicationStore
	PostTranslationsStore
	SeriesRepository
	TranslationGroupRepository
}
