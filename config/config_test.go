package config

import "testing"

func TestParsingResource(t *testing.T) {
	yaml := `
resources:
    comments:
        fields:
            body:
                type: text
                label: Your comment
                placeholder: Type your comment
                required: true
`
	app, err := ParseAppConfig([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	if len(app.Resources) != 1 {
		t.Fatalf("failed to parse `comments` resource")
	}

	comments := app.Resources["comments"]
	if len(comments.Fields) != 1 {
		t.Fatalf("failed to parse fields in the `comments` resource")
	}

	body := comments.Fields["body"]
	if body.Type != "text" {
		t.Fatalf("failed to parse field type as `text`")
	}
	if body.Label != "Your comment" {
		t.Fatalf("failed to parse field label")
	}
	if body.Placeholder != "Type your comment" {
		t.Fatalf("failed to parse field placeholder")
	}
	if !body.Required {
		t.Fatalf("failed to parse field as required")
	}
}

func TestParseHTTPServer(t *testing.T) {
	yaml := `
http:
    listen: 127.0.0.1:8080
`
	app, err := ParseAppConfig([]byte(yaml))
	if err != nil {
		t.Fatal(err)
	}

	if app.HTTP.Listen != "127.0.0.1:8080" {
		t.Fatalf("failed to parse application HTTP listen address: %v", app.HTTP.Listen)
	}
}
