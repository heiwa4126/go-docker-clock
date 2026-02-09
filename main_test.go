package main

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestTimeHandler_ValidTimezone(t *testing.T) {
	tests := []struct {
		name     string
		timezone string
	}{
		{"UTC", "UTC"},
		{"Asia/Tokyo", "Asia/Tokyo"},
		{"America/New_York", "America/New_York"},
		{"Local", "Local"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/time?tz="+tt.timezone, nil)
			w := httptest.NewRecorder()

			timeHandler(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("expected status 200, got %d", w.Code)
			}

			body := w.Body.String()
			if body == "" {
				t.Error("expected non-empty response body")
			}

			// レスポンスが指定されたフォーマット（layout）に従っているか確認
			// "2006-01-02 15:04:05 MST"
			pattern := `^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \w+$`
			matched, err := regexp.MatchString(pattern, body)
			if err != nil {
				t.Fatalf("regex error: %v", err)
			}
			if !matched {
				t.Errorf("response body doesn't match expected format: %s", body)
			}
		})
	}
}

func TestTimeHandler_InvalidTimezone(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/time?tz=Invalid/Timezone", nil)
	w := httptest.NewRecorder()

	timeHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("expected error message in response body")
	}
}

func TestTimeHandler_EmptyTimezone(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/time", nil)
	w := httptest.NewRecorder()

	timeHandler(w, req)

	// 空文字列のタイムゾーンは"UTC"として扱われるため、成功するはず
	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
