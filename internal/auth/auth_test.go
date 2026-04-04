package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid ApiKey",
			header:  "ApiKey mysecretkey123",
			want:    "mysecretkey123",
			wantErr: false,
		},
		{
			name:    "missing authorization header",
			header:  "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "wrong scheme (Bearer)",
			header:  "Bearer token123",
			want:    "",
			wantErr: true,
		},
		{
			name:    "malformed header - no scheme",
			header:  "justatoken",
			want:    "",
			wantErr: true,
		},
		{
			name:    "only scheme, no key",
			header:  "ApiKey",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.Header{}
			if tt.header != "" {
				h.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(h)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
			}
		})
	}
}



