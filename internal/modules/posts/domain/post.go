package domain

import (
	"encoding/json"
	"time"
)

type PostID string

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type Post struct {
	ID          PostID
	UUID        *string
	AuthorID    string
	Title       string
	Slug        string
	Subtitle    string
	Excerpt     string
	PostType    string
	Format      string
	Visibility  string
	Locale      string
	Timezone    string
	Content     string // markdown
	Status      Status
	WorkflowStage string
	Revision       int

	ReviewerUserID     *string
	LastEditedByUserID *string
	EditorUserIDs      []string

	Author       *UserSummary
	Reviewer     *UserSummary
	LastEditedBy *UserSummary
	Editors      []UserSummary

	ScheduledPublishAt *time.Time
	PublishedAt        *time.Time
	FirstIndexedAt     *time.Time

	CustomFields json.RawMessage
	Flags        json.RawMessage
	Engagement   json.RawMessage
	Workflow     json.RawMessage

	FeaturedMediaID *string
	FeaturedMediaPublicURL *string
	FeaturedAlt     *string
	FeaturedWidth   *int
	FeaturedHeight  *int
	FeaturedFocalX  *float32
	FeaturedFocalY  *float32
	FeaturedCredit  *string
	FeaturedLicense *string

	SEO     *PostSEO
	Metrics *PostMetrics

	Categories []PostCategory
	Tags       []PostTag
	Series     *PostSeries
	CoAuthors  []UserSummary
	Gallery    []PostGalleryItem
	Syndication []PostSyndication
	Changelog  []PostChangelogEntry
	Translations *PostTranslations

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserSummary struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"displayName"`
	Email       *string `json:"email,omitempty"`
	Role        *string `json:"role,omitempty"`
	AvatarURL   *string `json:"avatarUrl,omitempty"`
}

type PostSEO struct {
	Title          *string
	Description    *string
	CanonicalURL   *string
	Robots         *string
	OGType         *string
	OGImageURL     *string
	TwitterCard    *string
	StructuredData json.RawMessage
	UpdatedAt      time.Time
}

type PostMetrics struct {
	WordCount               int
	CharacterCount          int
	ReadingTimeMinutes      int
	EstReadTimeSeconds      int
	ViewCount               int64
	UniqueVisitors7d        int64
	ScrollDepthAvgPercent   float32
	BounceRatePercent       float32
	AvgTimeOnPageSeconds    int
	CommentCount            int
	LikeCount               int
	ShareCount              int
	BookmarkCount           int
	UpdatedAt               time.Time
}

type PostCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	IsPrimary bool   `json:"isPrimary"`
}

type PostTag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type PostSeries struct {
	ID        string  `json:"id"`
	Title     string  `json:"title"`
	Slug      string  `json:"slug"`
	PartIndex *int    `json:"partIndex,omitempty"`
	PartLabel *string `json:"partLabel,omitempty"`
}

type PostGalleryItem struct {
	ID        string  `json:"id"`
	MediaID   string  `json:"mediaId"`
	URL       *string `json:"url,omitempty"`
	SortOrder int     `json:"sortOrder"`
	Caption   *string `json:"caption,omitempty"`
	Alt       *string `json:"alt,omitempty"`
}

type PostSyndication struct {
	ID       string `json:"id"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Status   string `json:"status"`
}

type PostChangelogEntry struct {
	ID   string       `json:"id"`
	At   time.Time    `json:"at"`
	User *UserSummary `json:"user,omitempty"`
	Note string       `json:"note"`
}

type PostTranslations struct {
	GroupID      *string               `json:"translationGroupId,omitempty"`
	Translations []PostTranslationEntry `json:"translations"`
}

type PostTranslationEntry struct {
	PostID string `json:"postId"`
	Locale string `json:"locale"`
	Slug   string `json:"slug"`
}

