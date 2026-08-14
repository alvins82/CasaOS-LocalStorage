package v2

import "testing"

func TestInitializeDataLayoutDefersDefaultDirectoriesForMergerFS(t *testing.T) {
	tests := []struct {
		name            string
		mergerFSEnabled bool
		want            []string
	}{
		{name: "enabled", mergerFSEnabled: true, want: []string{"restore"}},
		{name: "disabled", mergerFSEnabled: false, want: []string{"defaults"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var events []string
			InitializeDataLayout(
				tt.mergerFSEnabled,
				func() { events = append(events, "restore") },
				func() { events = append(events, "defaults") },
			)
			if len(events) != len(tt.want) {
				t.Fatalf("InitializeDataLayout() events = %#v, want %#v", events, tt.want)
			}
			for i := range events {
				if events[i] != tt.want[i] {
					t.Fatalf("InitializeDataLayout() events = %#v, want %#v", events, tt.want)
				}
			}
		})
	}
}
