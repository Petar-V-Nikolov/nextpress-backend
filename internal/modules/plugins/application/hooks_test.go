package application

import (
	"context"
	"testing"
)

type countingHook struct {
	before, after *int
}

func (c countingHook) BeforePostSave(_ context.Context, _, _ string) error {
	*c.before++
	return nil
}

func (c countingHook) AfterPostSave(_ context.Context, _, _ string) error {
	*c.after++
	return nil
}

func TestHookRegistry_ChainInvokesInOrder(t *testing.T) {
	var beforeCount, afterCount int
	r := NewHookRegistry()
	r.RegisterPostHooks(countingHook{before: &beforeCount, after: &afterCount})
	r.RegisterPostHooks(countingHook{before: &beforeCount, after: &afterCount})

	ctx := context.Background()
	if err := r.BeforePostSave(ctx, "id", "slug"); err != nil {
		t.Fatal(err)
	}
	if beforeCount != 2 {
		t.Fatalf("expected 2 before calls, got %d", beforeCount)
	}
	if err := r.AfterPostSave(ctx, "id", "slug"); err != nil {
		t.Fatal(err)
	}
	if afterCount != 2 {
		t.Fatalf("expected 2 after calls, got %d", afterCount)
	}
}
