package mergerfs

import (
	"os"
	"reflect"
	"testing"
)

func TestSourceKeyPrefersModernBranchesKey(t *testing.T) {
	values := map[string]string{
		legacySourceKey: "/mnt/old",
		branchesKey:     "/mnt/new=RW",
	}

	got, err := sourceKey(values)
	if err != nil {
		t.Fatalf("sourceKey() error = %v", err)
	}
	if got != branchesKey {
		t.Fatalf("sourceKey() = %q, want %q", got, branchesKey)
	}
}

func TestSourceKeyFallsBackToLegacyKey(t *testing.T) {
	got, err := sourceKey(map[string]string{legacySourceKey: "/mnt/old"})
	if err != nil {
		t.Fatalf("sourceKey() error = %v", err)
	}
	if got != legacySourceKey {
		t.Fatalf("sourceKey() = %q, want %q", got, legacySourceKey)
	}
}

func TestSourceKeyRejectsUnknownControlInterface(t *testing.T) {
	if _, err := sourceKey(map[string]string{}); err == nil {
		t.Fatal("sourceKey() error = nil, want an error")
	}
}

func TestNormalizeSourcesRemovesModernBranchModes(t *testing.T) {
	got := normalizeSources("/mnt/a=RW:/mnt/b=RO:/mnt/c=NC:/mnt/name=kept")
	want := []string{"/mnt/a", "/mnt/b", "/mnt/c", "/mnt/name=kept"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("normalizeSources() = %#v, want %#v", got, want)
	}
}

func TestDedupeSourcesPreservesOrder(t *testing.T) {
	got := dedupeSources([]string{"/mnt/a", "/mnt/b", "/mnt/a", "/mnt/c"})
	want := []string{"/mnt/a", "/mnt/b", "/mnt/c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("dedupeSources() = %#v, want %#v", got, want)
	}
}

func TestLiveMountSourceControl(t *testing.T) {
	mountPoint := os.Getenv("MERGERFS_TEST_MOUNT")
	if mountPoint == "" {
		t.Skip("MERGERFS_TEST_MOUNT is not set")
	}

	sources, err := GetSource(mountPoint)
	if err != nil {
		t.Fatalf("GetSource() error = %v", err)
	}
	if len(sources) == 0 {
		t.Fatal("GetSource() returned no sources")
	}

	if err := SetSource(mountPoint, sources); err != nil {
		t.Fatalf("SetSource() error = %v", err)
	}

	got, err := GetSource(mountPoint)
	if err != nil {
		t.Fatalf("GetSource() after SetSource() error = %v", err)
	}
	if !reflect.DeepEqual(got, sources) {
		t.Fatalf("GetSource() after SetSource() = %#v, want %#v", got, sources)
	}
}
