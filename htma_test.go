package htma

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestHTMLRendering(t *testing.T) {
	// Component to be tested
	component := HTML().LangAttr("en-US").AddChild(
		Head().AddChild(
			Style().Text("/* styles to be added in the next listing */"),
		),
		Body().AddChild(
			Div().ClassAttr("grid").AddChild(
				Div().ClassAttr("a").Text("a"),
				Div().ClassAttr("b").Text("b"),
				Div().ClassAttr("c").Text("c"),
				Div().ClassAttr("d").Text("d"),
				Div().ClassAttr("e").Text("e"),
				Div().ClassAttr("f").Text("f"),
			),
		),
	)

	// Render the component
	var renderedHTML strings.Builder
	err := component.RenderStream(&renderedHTML)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}

	// Golden file path
	goldenFile := filepath.Join("testdata", "TestHTMLRendering.golden")

	// If the update flag is set, write to the golden file
	if *update {
		err := os.MkdirAll(filepath.Dir(goldenFile), 0755)
		if err != nil {
			t.Fatalf("failed to create testdata directory: %v", err)
		}
		err = ioutil.WriteFile(goldenFile, []byte(renderedHTML.String()), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
	}

	// Read the golden file
	expectedHTML, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	// Compare the rendered HTML with the golden file
	if renderedHTML.String() != string(expectedHTML) {
		t.Errorf("rendered HTML does not match golden file.\nGot:\n%s\n\nWant:\n%s", renderedHTML.String(), string(expectedHTML))
	}
}

func TestMixedContentRendering(t *testing.T) {
	// Component to be tested, demonstrating mixed text and element children
	component := Main().ClassAttr("main tile").AddChild(
		H1().Text("Team collaboration done right"),
		P().AddChild(
			Content("Thousands of teams from all over the world turn to "),
			B().Text("Ink"),
			Content(" to communicate and get things done."),
		),
	)

	// Render the component
	var renderedHTML strings.Builder
	err := component.RenderStream(&renderedHTML)
	if err != nil {
		t.Fatalf("failed to render component: %v", err)
	}

	// Golden file path
	goldenFile := filepath.Join("testdata", "TestMixedContentRendering.golden")

	// If the update flag is set, write to the golden file
	if *update {
		err := os.MkdirAll(filepath.Dir(goldenFile), 0755)
		if err != nil {
			t.Fatalf("failed to create testdata directory: %v", err)
		}
		err = ioutil.WriteFile(goldenFile, []byte(renderedHTML.String()), 0644)
		if err != nil {
			t.Fatalf("failed to update golden file: %v", err)
		}
	}

	// Read the golden file
	expectedHTML, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	// Compare the rendered HTML with the golden file
	if renderedHTML.String() != string(expectedHTML) {
		t.Errorf("rendered HTML does not match golden file.\nGot:\n%s\n\nWant:\n%s", renderedHTML.String(), string(expectedHTML))
	}
}
