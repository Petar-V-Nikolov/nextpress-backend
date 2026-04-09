package extensions

import "context"

// TranslationGroupRepository manages translation_groups.
type TranslationGroupRepository interface {
	CreateTranslationGroup(ctx context.Context, id string) error
	FindTranslationGroup(ctx context.Context, id string) (bool, error)
	DeleteTranslationGroup(ctx context.Context, id string) error
}
