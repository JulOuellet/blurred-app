package inbox

import (
	"testing"
	"time"

	"github.com/JulOuellet/blurred-app/internal/domains/events"
)

func stageEvents(t *testing.T, days ...string) []events.EventModel {
	t.Helper()
	evts := make([]events.EventModel, len(days))
	for i, day := range days {
		date, err := time.Parse("2006-01-02", day)
		if err != nil {
			t.Fatalf("bad date %q: %v", day, err)
		}
		evts[i] = events.EventModel{
			Name: "Stage " + string(rune('1'+i)),
			Date: &date,
		}
	}
	return evts
}

func TestMatchEventByDate(t *testing.T) {
	// Stages on July 4, 5 and 7 (rest day on the 6th).
	evts := stageEvents(t, "2026-07-04", "2026-07-05", "2026-07-07")
	paris, _ := time.LoadLocation("Europe/Paris")

	tests := []struct {
		name        string
		publishedAt time.Time
		wantIdx     int // -1 = no match
	}{
		{"evening of the stage", time.Date(2026, 7, 4, 19, 0, 0, 0, paris), 0},
		{"just after midnight", time.Date(2026, 7, 5, 0, 30, 0, 0, paris), 0},
		{"next morning", time.Date(2026, 7, 5, 9, 0, 0, 0, paris), 0},
		{"second stage evening", time.Date(2026, 7, 5, 22, 0, 0, 0, paris), 1},
		{"rest day evening maps to nothing", time.Date(2026, 7, 6, 20, 0, 0, 0, paris), -1},
		{"before the race", time.Date(2026, 7, 1, 12, 0, 0, 0, paris), -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := matchEventByDate(evts, tt.publishedAt)
			if tt.wantIdx == -1 {
				if got != nil {
					t.Fatalf("expected no match, got %q", got.Name)
				}
				return
			}
			if got == nil {
				t.Fatalf("expected %q, got no match", evts[tt.wantIdx].Name)
			}
			if got.Name != evts[tt.wantIdx].Name {
				t.Fatalf("expected %q, got %q", evts[tt.wantIdx].Name, got.Name)
			}
		})
	}
}

func TestMatchEventByDateNilDates(t *testing.T) {
	evts := []events.EventModel{{Name: "Stage 1"}, {Name: "Stage 2"}}
	if got := matchEventByDate(evts, time.Now()); got != nil {
		t.Fatalf("expected no match for events without dates, got %q", got.Name)
	}
}

func TestMatchEventByNumber(t *testing.T) {
	date := time.Date(2026, 7, 4, 0, 0, 0, 0, time.UTC)
	named := []events.EventModel{
		{Name: "Stage 1", Date: &date},
		{Name: "Stage 2", Date: &date},
		{Name: "Étape 10", Date: &date},
	}

	if got := matchEventByNumber(named, 2); got == nil || got.Name != "Stage 2" {
		t.Fatalf("expected Stage 2, got %v", got)
	}
	if got := matchEventByNumber(named, 10); got == nil || got.Name != "Étape 10" {
		t.Fatalf("expected Étape 10 via french name, got %v", got)
	}

	unnamed := []events.EventModel{
		{Name: "Opening day", Date: &date},
		{Name: "Queen stage", Date: &date},
	}
	if got := matchEventByNumber(unnamed, 2); got == nil || got.Name != "Queen stage" {
		t.Fatalf("expected positional fallback to Queen stage, got %v", got)
	}
	if got := matchEventByNumber(unnamed, 5); got != nil {
		t.Fatalf("expected no match for out-of-range number, got %q", got.Name)
	}
}
