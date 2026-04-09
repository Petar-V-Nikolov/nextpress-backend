package domain

// CorePostsPersistence is the subset of Repository used by core post CRUD and taxonomy assignment.
type CorePostsPersistence interface {
	PostReader
	PostWriter
	PostTaxonomyWriter
}
