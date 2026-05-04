package persistence

import (
	"fmt"
	"time"

	"github.com/nextpresskit/backend/pkg/seed/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const demoSeedRows = 100

// SeedDemo inserts demo posts and related rows.
func SeedDemo(tx *gorm.DB) error {
	if err := seedPosts(tx); err != nil {
		return err
	}
	if err := seedPostCategories(tx); err != nil {
		return err
	}
	if err := seedPostTags(tx); err != nil {
		return err
	}
	if err := seedPostSEO(tx); err != nil {
		return err
	}
	if err := seedPostMetrics(tx); err != nil {
		return err
	}
	if err := seedSeries(tx); err != nil {
		return err
	}
	if err := seedPostSeries(tx); err != nil {
		return err
	}
	if err := seedPostCoauthors(tx); err != nil {
		return err
	}
	if err := seedPostGalleryItems(tx); err != nil {
		return err
	}
	if err := seedPostChangelog(tx); err != nil {
		return err
	}
	if err := seedPostSyndication(tx); err != nil {
		return err
	}
	if err := seedTranslationGroups(tx); err != nil {
		return err
	}
	return seedPostTranslations(tx)
}

func seedPosts(tx *gorm.DB) error {
	now := time.Now().UTC()
	for i := 1; i <= demoSeedRows; i++ {
		status := "draft"
		var publishedAt *time.Time
		if i%3 == 0 {
			status = "published"
			t := now.Add(-time.Duration(i) * time.Hour)
			publishedAt = &t
		}
		postID := helpers.SeedUUID(0x0700, i)
		pid := postID
		focalX := float32(0.5)
		focalY := float32(0.5)
		w, h := 1920, 1080
		alt := fmt.Sprintf("Featured alt %03d", i)
		credit := "Seed Generator"
		license := "CC-BY"
		catID := helpers.SeedUUID(0x0400, i)
		mediaID := helpers.SeedUUID(0x0600, i)
		p := Post{
			ID:                 postID,
			UUID:               &pid,
			AuthorID:           helpers.UserPublicIDFromUUID(tx, "users", helpers.SeedUUID(0x0100, i)),
			Title:              fmt.Sprintf("Seed Post %03d", i),
			Slug:               fmt.Sprintf("seed-post-%03d", i),
			Content:            fmt.Sprintf("Seeded content body for post %03d.", i),
			Status:             status,
			PublishedAt:        publishedAt,
			PostType:           "article",
			Format:             "standard",
			Subtitle:           fmt.Sprintf("Subtitle for post %03d", i),
			Excerpt:            fmt.Sprintf("Short excerpt for post %03d", i),
			Visibility:         "public",
			Locale:             "en-US",
			Timezone:           "UTC",
			ScheduledPublishAt: helpers.PtrTime(now.Add(time.Duration(i) * time.Hour)),
			FirstIndexedAt:     &now,
			ReviewerUserID:     helpers.Int64Ptr(helpers.UserPublicIDFromUUID(tx, "users", helpers.SeedUUID(0x0100, ((i)%demoSeedRows)+1))),
			LastEditedByUserID: helpers.Int64Ptr(helpers.UserPublicIDFromUUID(tx, "users", helpers.SeedUUID(0x0100, ((i+1)%demoSeedRows)+1))),
			WorkflowStage:      "draft",
			Revision:           1,
			CustomFields:       helpers.Jf(`{"seed_index": %d}`, i),
			Flags:              helpers.J(`{"featured": false}`),
			Engagement:         helpers.Jf(`{"score": %d}`, i),
			Workflow:           helpers.Jf(`{"state": "draft-%03d"}`, i),
			FeaturedMediaID:    &mediaID,
			FeaturedAlt:        &alt,
			FeaturedWidth:      &w,
			FeaturedHeight:     &h,
			FeaturedFocalX:     &focalX,
			FeaturedFocalY:     &focalY,
			FeaturedCredit:     &credit,
			FeaturedLicense:    &license,
			PrimaryCategoryID:  &catID,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "slug"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"author_id", "title", "content", "status", "published_at", "uuid", "post_type", "format",
				"subtitle", "excerpt", "visibility", "locale", "timezone", "scheduled_publish_at", "first_indexed_at",
				"reviewer_user_id", "last_edited_by_user_id", "workflow_stage", "revision",
				"custom_fields", "flags", "engagement", "workflow",
				"featured_media_id", "featured_alt", "featured_width", "featured_height",
				"featured_focal_x", "featured_focal_y", "featured_credit", "featured_license",
				"primary_category_id", "deleted_at", "updated_at",
			}),
		}).Create(&p).Error; err != nil {
			return fmt.Errorf("posts row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostCategories(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		row := PostCategory{PostID: helpers.SeedUUID(0x0700, i), CategoryID: helpers.SeedUUID(0x0400, i)}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "post_id"}, {Name: "category_id"}},
			DoNothing: true,
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_categories row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostTags(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		row := PostTag{PostID: helpers.SeedUUID(0x0700, i), TagID: helpers.SeedUUID(0x0500, i)}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "post_id"}, {Name: "tag_id"}},
			DoNothing: true,
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_tags row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostSEO(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		t := fmt.Sprintf("SEO Title %03d", i)
		d := fmt.Sprintf("SEO Description %03d", i)
		canon := fmt.Sprintf("https://example.local/seed-post-%03d", i)
		robots := "index,follow"
		ogt := "article"
		ogu := fmt.Sprintf("https://example.local/media/seed-image-%03d.jpg", i)
		tw := "summary_large_image"
		row := PostSEO{
			PostID:         helpers.SeedUUID(0x0700, i),
			Title:          &t,
			Description:    &d,
			CanonicalURL:   &canon,
			Robots:         &robots,
			OGType:         &ogt,
			OGImageURL:     &ogu,
			TwitterCard:    &tw,
			StructuredData: helpers.Jf(`{"seed_index": %d}`, i),
			UpdatedAt:      time.Now().UTC(),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "post_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"title":           row.Title,
				"description":     row.Description,
				"canonical_url":   row.CanonicalURL,
				"robots":          row.Robots,
				"og_type":         row.OGType,
				"og_image_url":    row.OGImageURL,
				"twitter_card":    row.TwitterCard,
				"structured_data": row.StructuredData,
				"updated_at":      time.Now().UTC(),
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_seo row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostMetrics(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		row := PostMetrics{
			PostID:                helpers.SeedUUID(0x0700, i),
			WordCount:             800 + i,
			CharacterCount:        5000 + i*20,
			ReadingTimeMinutes:    5,
			EstReadTimeSeconds:    300,
			ViewCount:             int64(i * 100),
			UniqueVisitors7d:      int64(i * 10),
			ScrollDepthAvgPercent: 60.5,
			BounceRatePercent:     35.0,
			AvgTimeOnPageSeconds:  240,
			CommentCount:          i % 20,
			LikeCount:             i % 25,
			ShareCount:            i % 15,
			BookmarkCount:         i % 10,
			UpdatedAt:             time.Now().UTC(),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "post_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"word_count": row.WordCount, "character_count": row.CharacterCount,
				"reading_time_minutes": row.ReadingTimeMinutes, "est_read_time_seconds": row.EstReadTimeSeconds,
				"view_count": row.ViewCount, "unique_visitors_7d": row.UniqueVisitors7d,
				"scroll_depth_avg_percent": row.ScrollDepthAvgPercent, "bounce_rate_percent": row.BounceRatePercent,
				"avg_time_on_page_seconds": row.AvgTimeOnPageSeconds, "comment_count": row.CommentCount,
				"like_count": row.LikeCount, "share_count": row.ShareCount, "bookmark_count": row.BookmarkCount,
				"updated_at": time.Now().UTC(),
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_metrics row %d: %w", i, err)
		}
	}
	return nil
}

func seedSeries(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		s := Series{
			ID:    helpers.SeedUUID(0x0c00, i),
			Title: fmt.Sprintf("Series %03d", i),
			Slug:  fmt.Sprintf("series-%03d", i),
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "slug"}},
			DoUpdates: clause.Assignments(map[string]any{
				"title": s.Title, "updated_at": time.Now().UTC(),
			}),
		}).Create(&s).Error; err != nil {
			return fmt.Errorf("series row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostSeries(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		pi := i
		lbl := fmt.Sprintf("Part %03d", i)
		row := PostSeries{
			PostID:    helpers.SeedUUID(0x0700, i),
			SeriesID:  helpers.SeedUUID(0x0c00, i),
			PartIndex: &pi,
			PartLabel: &lbl,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "post_id"}, {Name: "series_id"}},
			DoNothing: true,
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_series row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostCoauthors(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		uid := helpers.UserPublicIDFromUUID(tx, "users", helpers.SeedUUID(0x0100, ((i)%demoSeedRows)+1))
		row := PostCoauthor{PostID: helpers.SeedUUID(0x0700, i), UserID: uid, SortOrder: 1}
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "post_id"}, {Name: "user_id"}},
			DoNothing: true,
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_coauthors row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostGalleryItems(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		cap := fmt.Sprintf("Gallery caption %03d", i)
		alt := fmt.Sprintf("Gallery alt %03d", i)
		row := PostGalleryItem{
			ID:        helpers.SeedUUID(0x0d00, i),
			PostID:    helpers.SeedUUID(0x0700, i),
			MediaID:   helpers.SeedUUID(0x0600, i),
			SortOrder: i,
			Caption:   &cap,
			Alt:       &alt,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"post_id": row.PostID, "media_id": row.MediaID, "sort_order": row.SortOrder,
				"caption": row.Caption, "alt": row.Alt,
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_gallery_items row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostChangelog(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		u := helpers.UserPublicIDFromUUID(tx, "users", helpers.SeedUUID(0x0100, ((i+2)%demoSeedRows)+1))
		note := fmt.Sprintf("Seed changelog entry %03d", i)
		row := PostChangelog{
			ID:     helpers.SeedUUID(0x0e00, i),
			PostID: helpers.SeedUUID(0x0700, i),
			At:     time.Now().UTC().Add(-time.Duration(i) * time.Hour),
			UserID: helpers.Int64Ptr(u),
			Note:   note,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"post_id": row.PostID, "at": row.At, "user_id": row.UserID, "note": row.Note,
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_changelog row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostSyndication(tx *gorm.DB) error {
	now := time.Now().UTC()
	for i := 1; i <= demoSeedRows; i++ {
		row := PostSyndication{
			ID:        helpers.SeedUUID(0x0f00, i),
			PostID:    helpers.SeedUUID(0x0700, i),
			Platform:  "medium",
			URL:       fmt.Sprintf("https://medium.example/seed-post-%03d", i),
			Status:    "active",
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"post_id": row.PostID, "platform": row.Platform, "url": row.URL, "status": row.Status,
				"updated_at": time.Now().UTC(),
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_syndication row %d: %w", i, err)
		}
	}
	return nil
}

func seedTranslationGroups(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		row := TranslationGroup{ID: helpers.SeedUUID(0x1000, i), CreatedAt: time.Now().UTC()}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&row).Error; err != nil {
			return fmt.Errorf("translation_groups row %d: %w", i, err)
		}
	}
	return nil
}

func seedPostTranslations(tx *gorm.DB) error {
	for i := 1; i <= demoSeedRows; i++ {
		locale := "en-US"
		if i%2 == 0 {
			locale = "bg-BG"
		}
		row := PostTranslation{
			PostID:  helpers.SeedUUID(0x0700, i),
			GroupID: helpers.SeedUUID(0x1000, i),
			Locale:  locale,
		}
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "post_id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"group_id": row.GroupID, "locale": row.Locale,
			}),
		}).Create(&row).Error; err != nil {
			return fmt.Errorf("post_translations row %d: %w", i, err)
		}
	}
	return nil
}
