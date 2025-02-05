package web

import (
	"testing"
)

func TestTemplateParser(t *testing.T) {
	tests := []struct {
		correct           []string
		targets           []string
		includeTargetPath bool
	}{
		{
			correct:           []string{"templates/index", "templates/mypage/index", "templates/mypage/header", "components/index"},
			targets:           []string{"templates", "components"},
			includeTargetPath: true,
		},
		{
			correct:           []string{"index", "mypage/index", "mypage/header"},
			targets:           []string{"templates"},
			includeTargetPath: false,
		},
	}
	for _, test := range tests {
		templates, err := ParseTemplate("test", test.includeTargetPath, test.targets...)
		if err != nil {
			t.Errorf("Failed to parse template: %v", err)
		}
		if len(templates.Templates()) != len(test.correct) {
			t.Errorf("Expected %d templates, got %d", len(test.correct), len(templates.Templates()))
		}
		for _, v := range test.correct {
			template := templates.Lookup(v)
			if template == nil {
				t.Errorf("Template %s not found", v)
				continue
			}
			if template.Tree == nil {
				t.Errorf("Template %s has no tree", v)
			}
		}
	}
}
