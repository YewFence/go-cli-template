package main

import "testing"

func TestIsTemplateOrigin(t *testing.T) {
	tests := []struct {
		name      string
		remoteURL string
		want      bool
	}{
		{
			name:      "matches template origin",
			remoteURL: "https://github.com/YewFence/go-cli-template",
			want:      true,
		},
		{
			name:      "matches template origin with git suffix",
			remoteURL: "https://github.com/YewFence/go-cli-template.git\n",
			want:      true,
		},
		{
			name:      "matches template origin with http",
			remoteURL: "http://github.com/YewFence/go-cli-template.git",
			want:      true,
		},
		{
			name:      "matches template origin with ssh",
			remoteURL: "git@github.com:YewFence/go-cli-template.git",
			want:      true,
		},
		{
			name:      "ignores other origin",
			remoteURL: "https://github.com/YewFence/other.git",
			want:      false,
		},
		{
			name:      "ignores other ssh origin",
			remoteURL: "git@github.com:YewFence/other.git",
			want:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isTemplateOrigin(test.remoteURL)
			if got != test.want {
				t.Fatalf("isTemplateOrigin(%q) = %v, want %v", test.remoteURL, got, test.want)
			}
		})
	}
}
