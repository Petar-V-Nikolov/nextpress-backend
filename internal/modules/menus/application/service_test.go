package application

import (
	"context"
	"errors"
	"testing"

	menuDomain "github.com/Petar-V-Nikolov/nextpress-backend/internal/modules/menus/domain"
)

type menusRepoStub struct {
	menus map[menuDomain.MenuID]*menuDomain.Menu
}

func (s *menusRepoStub) CreateMenu(_ context.Context, m *menuDomain.Menu) error {
	if s.menus == nil {
		s.menus = map[menuDomain.MenuID]*menuDomain.Menu{}
	}
	for _, v := range s.menus {
		if v.Slug == m.Slug {
			return errors.New("duplicate")
		}
	}
	cp := *m
	s.menus[m.ID] = &cp
	return nil
}
func (s *menusRepoStub) ListMenus(_ context.Context, _, _ int) ([]menuDomain.Menu, error) { return nil, nil }
func (s *menusRepoStub) FindMenuByID(_ context.Context, id menuDomain.MenuID) (*menuDomain.Menu, error) {
	return s.menus[id], nil
}
func (s *menusRepoStub) FindMenuBySlug(_ context.Context, slug string) (*menuDomain.Menu, error) {
	for _, v := range s.menus {
		if v.Slug == slug {
			return v, nil
		}
	}
	return nil, nil
}
func (s *menusRepoStub) UpdateMenu(_ context.Context, m *menuDomain.Menu) error {
	s.menus[m.ID] = m
	return nil
}
func (s *menusRepoStub) DeleteMenu(_ context.Context, id menuDomain.MenuID) error {
	delete(s.menus, id)
	return nil
}
func (s *menusRepoStub) ListMenuItems(_ context.Context, _ menuDomain.MenuID) ([]menuDomain.MenuItem, error) {
	return []menuDomain.MenuItem{}, nil
}
func (s *menusRepoStub) ReplaceMenuItems(_ context.Context, _ menuDomain.MenuID, _ []menuDomain.MenuItem) error {
	return nil
}

func TestCreateMenu_NormalizesSlug(t *testing.T) {
	repo := &menusRepoStub{menus: map[menuDomain.MenuID]*menuDomain.Menu{}}
	svc := NewService(repo)

	m, err := svc.CreateMenu(context.Background(), "Main", " Main Nav ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m.Slug != "main-nav" {
		t.Fatalf("expected normalized slug main-nav, got %q", m.Slug)
	}
}

func TestReplaceItems_InvalidURLInput(t *testing.T) {
	repo := &menusRepoStub{
		menus: map[menuDomain.MenuID]*menuDomain.Menu{
			"menu-1": {ID: "menu-1", Name: "Main", Slug: "main"},
		},
	}
	svc := NewService(repo)

	err := svc.ReplaceItems(context.Background(), "menu-1", []ItemInput{
		{Label: "Broken", ItemType: "url"},
	})
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestGetMenu_NotFound(t *testing.T) {
	svc := NewService(&menusRepoStub{menus: map[menuDomain.MenuID]*menuDomain.Menu{}})
	_, err := svc.GetMenu(context.Background(), "missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

