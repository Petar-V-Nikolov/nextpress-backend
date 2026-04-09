package ports

import (
	"context"

	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/ident"
	"github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/posts/domain/model"
)

// PostWriter persists core post rows.
type PostWriter interface {
	Create(ctx context.Context, post *model.Post) error
	Update(ctx context.Context, post *model.Post) error
	Delete(ctx context.Context, id ident.PostID) error
}
