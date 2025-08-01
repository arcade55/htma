// Package htma provides a Hypertext Markup Abstraction for generating HTML in pure Go.
package htma

import (
	"fmt"
	"html"
	"io"
	"strings"
)

// Renderable defines types that can render HTML and SSE.
type Renderable interface {
	Render() string
	RenderStream(w io.Writer) error
}

// TextContent represents an escaped plain text node.
type TextContent struct {
	Content string
}

// Render returns the escaped text content.
func (t TextContent) Render() string {
	return html.EscapeString(t.Content)
}

// RenderStream writes the escaped text content to a writer.
func (t TextContent) RenderStream(w io.Writer) error {
	_, err := io.WriteString(w, html.EscapeString(t.Content))
	return err
}

// Content creates an escaped text node that can be a child of another element.
// This is used for creating mixed-content elements (text and other tags).
func Content(content string) TextContent {
	return TextContent{Content: content}
}

// Raw represents an unescaped HTML string.
type Raw struct {
	Content string
}

// Render returns the unescaped content as is.
func (r Raw) Render() string {
	return r.Content
}

// RenderStream writes the unescaped content to a writer.
func (r Raw) RenderStream(w io.Writer) error {
	_, err := io.WriteString(w, r.Content)
	return err
}

// RawContent creates a raw HTML node that will not be escaped.
// Use with caution, as this can open you up to XSS vulnerabilities if used with untrusted content.
func RawContent(content string) Raw {
	return Raw{Content: content}
}

// Attributable defines types that can set attributes.
type Attributable interface {
	Attr(key, value string) Element
}

// Element is the base HTML element, modeling tags, attributes, and children.
type Element struct {
	tag      string
	attrs    map[string]string
	children []Renderable
	text     string
	isVoid   bool
	isRoot   bool // Indicates if this is the root <html> element
}

// newElement creates a generic element (internal).
func newElement(tag string, isVoid bool) Element {
	return Element{
		tag:    tag,
		attrs:  make(map[string]string),
		isVoid: isVoid,
	}
}

// Constructors for HTML Elements (alphabetical order)
func A() Element {
	return newElement("a", false)
}

func Abbr() Element {
	return newElement("abbr", false)
}

func Address() Element {
	return newElement("address", false)
}

func Area() Element {
	return newElement("area", true)
}

func Article() Element {
	return newElement("article", false)
}

func Aside() Element {
	return newElement("aside", false)
}

func Audio() Element {
	return newElement("audio", false)
}

func B() Element {
	return newElement("b", false)
}

func Base() Element {
	return newElement("base", true)
}

func Bdi() Element {
	return newElement("bdi", false)
}

func Bdo() Element {
	return newElement("bdo", false)
}

func Blockquote() Element {
	return newElement("blockquote", false)
}

func Body() Element {
	return newElement("body", false)
}

func Br() Element {
	return newElement("br", true)
}

func Button() Element {
	return newElement("button", false)
}

func Canvas() Element {
	return newElement("canvas", false)
}

func Caption() Element {
	return newElement("caption", false)
}

func Cite() Element {
	return newElement("cite", false)
}

func Code() Element {
	return newElement("code", false)
}

func Col() Element {
	return newElement("col", true)
}

func Colgroup() Element {
	return newElement("colgroup", false)
}

func Data() Element {
	return newElement("data", false)
}

func Datalist() Element {
	return newElement("datalist", false)
}

func Dd() Element {
	return newElement("dd", false)
}

func Del() Element {
	return newElement("del", false)
}

func Details() Element {
	return newElement("details", false)
}

func Dfn() Element {
	return newElement("dfn", false)
}

func Dialog() Element {
	return newElement("dialog", false)
}

func Div() Element {
	return newElement("div", false)
}

func Dl() Element {
	return newElement("dl", false)
}

func Dt() Element {
	return newElement("dt", false)
}

func Em() Element {
	return newElement("em", false)
}

func Embed() Element {
	return newElement("embed", true)
}

func Fieldset() Element {
	return newElement("fieldset", false)
}

func Figcaption() Element {
	return newElement("figcaption", false)
}

func Figure() Element {
	return newElement("figure", false)
}

func Footer() Element {
	return newElement("footer", false)
}

func Form() Element {
	return newElement("form", false)
}

func H1() Element {
	return newElement("h1", false)
}

func H2() Element {
	return newElement("h2", false)
}

func H3() Element {
	return newElement("h3", false)
}

func H4() Element {
	return newElement("h4", false)
}

func H5() Element {
	return newElement("h5", false)
}

func H6() Element {
	return newElement("h6", false)
}

func Head() Element {
	return newElement("head", false)
}

func Header() Element {
	return newElement("header", false)
}

func Hgroup() Element {
	return newElement("hgroup", false)
}

func Hr() Element {
	return newElement("hr", true)
}

func HTML() Element {
	return Element{tag: "html", attrs: make(map[string]string), isRoot: true}
}

func I() Element {
	return newElement("i", false)
}

func Iframe() Element {
	return newElement("iframe", false)
}

func Img() Element {
	return newElement("img", true)
}

func Input() Element {
	return newElement("input", true)
}

func Ins() Element {
	return newElement("ins", false)
}

func Kbd() Element {
	return newElement("kbd", false)
}

func Label() Element {
	return newElement("label", false)
}

func Legend() Element {
	return newElement("legend", false)
}

func Li() Element {
	return newElement("li", false)
}

func Link() Element {
	return newElement("link", true)
}

func Main() Element {
	return newElement("main", false)
}

func Map() Element {
	return newElement("map", false)
}

func Mark() Element {
	return newElement("mark", false)
}

func Math() Element {
	return newElement("math", false)
}

func Menu() Element {
	return newElement("menu", false)
}

func Meta() Element {
	return newElement("meta", true)
}

func Meter() Element {
	return newElement("meter", false)
}

func Nav() Element {
	return newElement("nav", false)
}

func Noscript() Element {
	return newElement("noscript", false)
}

func Object() Element {
	return newElement("object", false)
}

func Ol() Element {
	return newElement("ol", false)
}

func Optgroup() Element {
	return newElement("optgroup", false)
}

func Option() Element {
	return newElement("option", false)
}

func Output() Element {
	return newElement("output", false)
}

func P() Element {
	return newElement("p", false)
}

func Picture() Element {
	return newElement("picture", false)
}

func Pre() Element {
	return newElement("pre", false)
}

func Progress() Element {
	return newElement("progress", false)
}

func Q() Element {
	return newElement("q", false)
}

func Rp() Element {
	return newElement("rp", false)
}

func Rt() Element {
	return newElement("rt", false)
}

func Ruby() Element {
	return newElement("ruby", false)
}

func S() Element {
	return newElement("s", false)
}

func Samp() Element {
	return newElement("samp", false)
}

func Script() Element {
	return newElement("script", false)
}

func Search() Element {
	return newElement("search", false)
}

func Section() Element {
	return newElement("section", false)
}

func Select() Element {
	return newElement("select", false)
}

func Slot() Element {
	return newElement("slot", false)
}

func Small() Element {
	return newElement("small", false)
}

func Source() Element {
	return newElement("source", true)
}

func Span() Element {
	return newElement("span", false)
}

func Strong() Element {
	return newElement("strong", false)
}

func Style() Element {
	return newElement("style", false)
}

func Sub() Element {
	return newElement("sub", false)
}

func Summary() Element {
	return newElement("summary", false)
}

func Sup() Element {
	return newElement("sup", false)
}

func Svg() Element {
	return newElement("svg", false)
}

func Table() Element {
	return newElement("table", false)
}

func Tbody() Element {
	return newElement("tbody", false)
}

func Td() Element {
	return newElement("td", false)
}

func Template() Element {
	return newElement("template", false)
}

func Textarea() Element {
	return newElement("textarea", false)
}

func Tfoot() Element {
	return newElement("tfoot", false)
}

func Th() Element {
	return newElement("th", false)
}

func Thead() Element {
	return newElement("thead", false)
}

func Time() Element {
	return newElement("time", false)
}

func Title(text string) Element {
	return newElement("title", false).Text(text)
}

func Tr() Element {
	return newElement("tr", false)
}

func Track() Element {
	return newElement("track", true)
}

func U() Element {
	return newElement("u", false)
}

func Ul() Element {
	return newElement("ul", false)
}

func Var() Element {
	return newElement("var", false)
}

func Video() Element {
	return newElement("video", false)
}

func Wbr() Element {
	return newElement("wbr", true)
}

// Element Methods (Chainable)
func (e Element) AddChild(children ...Renderable) Element {
	if e.isVoid {
		panic(fmt.Sprintf("cannot add children to void element: <%s>", e.tag))
	}
	e.children = append(e.children, children...)
	return e
}

func (e Element) Text(text string) Element {
	if e.isVoid {
		panic(fmt.Sprintf("cannot add text to void element: <%s>", e.tag))
	}
	e.text = text
	return e
}

func (e Element) IDAttr(id string) Element {
	if strings.ContainsAny(id, " \t\n") {
		panic("invalid ID: " + id)
	}
	e.attrs["id"] = id
	return e
}

func (e Element) ClassAttr(class string) Element {
	classes := strings.Fields(class)
	if len(classes) == 0 {
		return e
	}
	for _, cls := range classes {
		if strings.ContainsAny(cls, " \t\n;") {
			panic("invalid class name: " + cls)
		}
	}
	combined := strings.Join(classes, " ")
	e.attrs["class"] = appendClassInternal(e.attrs["class"], combined)
	return e
}

func (e Element) Classes(classes ...string) Element {
	for _, c := range classes {
		e = e.ClassAttr(c)
	}
	return e
}

func (e Element) StyleAttr(key, value string) Element {
	current := e.attrs["style"]
	if current != "" {
		current += "; "
	}
	e.attrs["style"] = current + fmt.Sprintf("%s: %s", key, value)
	return e
}

func (e Element) Attr(key, value string) Element {
	e.attrs[key] = value
	return e
}

// Global Attribute Methods
func (e Element) AccessKeyAttr(key string) Element {
	return e.Attr("accesskey", key)
}

func (e Element) AutocapitalizeAttr(value string) Element {
	return e.Attr("autocapitalize", value)
}

func (e Element) AutofocusAttr() Element {
	return e.Attr("autofocus", "")
}

func (e Element) ContentEditableAttr(value string) Element {
	return e.Attr("contenteditable", value)
}

func (e Element) DirAttr(value string) Element {
	return e.Attr("dir", value)
}

func (e Element) DraggableAttr(value string) Element {
	return e.Attr("draggable", value)
}

func (e Element) EnterKeyHintAttr(value string) Element {
	return e.Attr("enterkeyhint", value)
}

func (e Element) HiddenAttr() Element {
	return e.Attr("hidden", "")
}

func (e Element) InertAttr() Element {
	return e.Attr("inert", "")
}

func (e Element) InputModeAttr(value string) Element {
	return e.Attr("inputmode", value)
}

func (e Element) IsAttr(value string) Element {
	return e.Attr("is", value)
}

func (e Element) ItemIDAttr(value string) Element {
	return e.Attr("itemid", value)
}

func (e Element) ItemPropAttr(value string) Element {
	return e.Attr("itemprop", value)
}

func (e Element) ItemRefAttr(value string) Element {
	return e.Attr("itemref", value)
}

func (e Element) ItemScopeAttr() Element {
	return e.Attr("itemscope", "")
}

func (e Element) ItemTypeAttr(value string) Element {
	return e.Attr("itemtype", value)
}

func (e Element) LangAttr(value string) Element {
	return e.Attr("lang", value)
}

func (e Element) NonceAttr(value string) Element {
	return e.Attr("nonce", value)
}

func (e Element) PartAttr(value string) Element {
	return e.Attr("part", value)
}

func (e Element) PopoverAttr(value string) Element {
	return e.Attr("popover", value)
}

func (e Element) SlotAttr(value string) Element {
	return e.Attr("slot", value)
}

func (e Element) SpellCheckAttr(value string) Element {
	return e.Attr("spellcheck", value)
}

func (e Element) TabIndexAttr(index int) Element {
	return e.Attr("tabindex", fmt.Sprint(index))
}

func (e Element) TitleAttr(value string) Element {
	return e.Attr("title", value)
}

func (e Element) TranslateAttr(value string) Element {
	return e.Attr("translate", value)
}

// Aria Attributes
func (e Element) AriaLabelAttr(label string) Element {
	return e.Attr("aria-label", label)
}

func (e Element) AriaHiddenAttr(value string) Element {
	return e.Attr("aria-hidden", value)
}

func (e Element) AriaRoleAttr(role string) Element {
	return e.Attr("role", role)
}

// Custom Data Attributes
func (e Element) DataAttr(key, value string) Element {
	return e.Attr("data-"+key, value)
}

// Datastar Directives
func (e Element) DataAttrAttr(value string) Element {
	return e.Attr("data-attr", value)
}

func (e Element) DataBindAttr(value string) Element {
	return e.Attr("data-bind", value)
}

func (e Element) DataClassAttr(value string) Element {
	return e.Attr("data-class", value)
}

func (e Element) DataComputedAttr(value string) Element {
	return e.Attr("data-computed", value)
}

func (e Element) DataEffectAttr(value string) Element {
	return e.Attr("data-effect", value)
}

func (e Element) DataIgnoreAttr(value string) Element {
	return e.Attr("data-ignore", value)
}

func (e Element) DataIgnoreMorphAttr(value string) Element {
	return e.Attr("data-ignore-morph", value)
}

func (e Element) DataIndicatorAttr(value string) Element {
	return e.Attr("data-indicator", value)
}

func (e Element) DataJsonSignalsAttr(value string) Element {
	return e.Attr("data-json-signals", value)
}

func (e Element) DataOnClickAttr(value string) Element {
	return e.Attr("data-on-click", value)
}

func (e Element) DataOnIntersectAttr(value string) Element {
	return e.Attr("data-on-intersect", value)
}

func (e Element) DataOnIntervalAttr(value string) Element {
	return e.Attr("data-on-interval", value)
}

func (e Element) DataOnLoadAttr(value string) Element {
	return e.Attr("data-on-load", value)
}

func (e Element) DataOnSignalPatchAttr(value string) Element {
	return e.Attr("data-on-signal-patch", value)
}

func (e Element) DataOnSignalPatchFilterAttr(value string) Element {
	return e.Attr("data-on-signal-patch-filter", value)
}

func (e Element) DataPreserveAttrAttr(value string) Element {
	return e.Attr("data-preserve-attr", value)
}

func (e Element) DataRefAttr(value string) Element {
	return e.Attr("data-ref", value)
}

func (e Element) DataShowAttr(value string) Element {
	return e.Attr("data-show", value)
}

func (e Element) DataSignalsAttr(value string) Element {
	return e.Attr("data-signals", value)
}

func (e Element) DataStyleAttr(value string) Element {
	return e.Attr("data-style", value)
}

func (e Element) DataTextAttr(value string) Element {
	return e.Attr("data-text", value)
}

// Element-Specific Attribute Methods
func (e Element) AcceptAttr(value string) Element {
	return e.Attr("accept", value)
}

func (e Element) AcceptCharsetAttr(value string) Element {
	return e.Attr("accept-charset", value)
}

func (e Element) ActionAttr(action string) Element {
	return e.Attr("action", action)
}

func (e Element) AltAttr(alt string) Element {
	return e.Attr("alt", alt)
}

func (e Element) AsyncAttr() Element {
	return e.Attr("async", "")
}

func (e Element) AutoPlayAttr() Element {
	return e.Attr("autoplay", "")
}

func (e Element) CharsetAttr(value string) Element {
	return e.Attr("charset", value)
}

func (e Element) CheckedAttr() Element {
	return e.Attr("checked", "")
}

func (e Element) CiteAttr(value string) Element {
	return e.Attr("cite", value)
}

func (e Element) ColsAttr(cols int) Element {
	return e.Attr("cols", fmt.Sprint(cols))
}

func (e Element) ColSpanAttr(span int) Element {
	return e.Attr("colspan", fmt.Sprint(span))
}

func (e Element) ControlsAttr() Element {
	return e.Attr("controls", "")
}

func (e Element) CoordsAttr(value string) Element {
	return e.Attr("coords", value)
}

func (e Element) CrossOriginAttr(value string) Element {
	return e.Attr("crossorigin", value)
}

func (e Element) DateTimeAttr(value string) Element {
	return e.Attr("datetime", value)
}

func (e Element) DefaultAttr() Element {
	return e.Attr("default", "")
}

func (e Element) DeferAttr() Element {
	return e.Attr("defer", "")
}

func (e Element) DisabledAttr() Element {
	return e.Attr("disabled", "")
}

func (e Element) DownloadAttr(value string) Element {
	return e.Attr("download", value)
}

func (e Element) EncTypeAttr(value string) Element {
	return e.Attr("enctype", value)
}

func (e Element) ForAttr(value string) Element {
	return e.Attr("for", value)
}

func (e Element) FormAttr(value string) Element {
	return e.Attr("form", value)
}

func (e Element) FormActionAttr(value string) Element {
	return e.Attr("formaction", value)
}

func (e Element) FormEncTypeAttr(value string) Element {
	return e.Attr("formenctype", value)
}

func (e Element) FormMethodAttr(value string) Element {
	return e.Attr("formmethod", value)
}

func (e Element) FormNoValidateAttr() Element {
	return e.Attr("formnovalidate", "")
}

func (e Element) FormTargetAttr(value string) Element {
	return e.Attr("formtarget", value)
}

func (e Element) HeightAttr(height int) Element {
	return e.Attr("height", fmt.Sprint(height))
}

func (e Element) HrefAttr(href string) Element {
	return e.Attr("href", href)
}

func (e Element) HrefLangAttr(value string) Element {
	return e.Attr("hreflang", value)
}

func (e Element) HttpEquivAttr(value string) Element {
	return e.Attr("http-equiv", value)
}

func (e Element) IntegrityAttr(value string) Element {
	return e.Attr("integrity", value)
}

func (e Element) KindAttr(value string) Element {
	return e.Attr("kind", value)
}

func (e Element) LabelAttr(value string) Element {
	return e.Attr("label", value)
}

func (e Element) ListAttr(value string) Element {
	return e.Attr("list", value)
}

func (e Element) LoopAttr() Element {
	return e.Attr("loop", "")
}

func (e Element) MaxAttr(value string) Element {
	return e.Attr("max", value)
}

func (e Element) MaxLengthAttr(length int) Element {
	return e.Attr("maxlength", fmt.Sprint(length))
}

func (e Element) MediaAttr(value string) Element {
	return e.Attr("media", value)
}

func (e Element) MethodAttr(value string) Element {
	return e.Attr("method", value)
}

func (e Element) MinAttr(value string) Element {
	return e.Attr("min", value)
}

func (e Element) MinLengthAttr(length int) Element {
	return e.Attr("minlength", fmt.Sprint(length))
}

func (e Element) MultipleAttr() Element {
	return e.Attr("multiple", "")
}

func (e Element) MutedAttr() Element {
	return e.Attr("muted", "")
}

func (e Element) NameAttr(value string) Element {
	return e.Attr("name", value)
}

func (e Element) NoValidateAttr() Element {
	return e.Attr("novalidate", "")
}

func (e Element) OpenAttr() Element {
	return e.Attr("open", "")
}

func (e Element) PatternAttr(value string) Element {
	return e.Attr("pattern", value)
}

func (e Element) PlaceholderAttr(value string) Element {
	return e.Attr("placeholder", value)
}

func (e Element) PosterAttr(value string) Element {
	return e.Attr("poster", value)
}

func (e Element) PreloadAttr(value string) Element {
	return e.Attr("preload", value)
}

func (e Element) ReadOnlyAttr() Element {
	return e.Attr("readonly", "")
}

func (e Element) RelAttr(value string) Element {
	return e.Attr("rel", value)
}

func (e Element) RequiredAttr() Element {
	return e.Attr("required", "")
}

func (e Element) ReversedAttr() Element {
	return e.Attr("reversed", "")
}

func (e Element) RowsAttr(rows int) Element {
	return e.Attr("rows", fmt.Sprint(rows))
}

func (e Element) RowSpanAttr(span int) Element {
	return e.Attr("rowspan", fmt.Sprint(span))
}

func (e Element) SandboxAttr(value string) Element {
	return e.Attr("sandbox", value)
}

func (e Element) ScopeAttr(value string) Element {
	return e.Attr("scope", value)
}

func (e Element) SelectedAttr() Element {
	return e.Attr("selected", "")
}

func (e Element) ShapeAttr(value string) Element {
	return e.Attr("shape", value)
}

func (e Element) SizeAttr(size int) Element {
	return e.Attr("size", fmt.Sprint(size))
}

func (e Element) SizesAttr(value string) Element {
	return e.Attr("sizes", value)
}

func (e Element) SpanAttr(span int) Element {
	return e.Attr("span", fmt.Sprint(span))
}

func (e Element) SrcAttr(src string) Element {
	return e.Attr("src", src)
}

func (e Element) SrcDocAttr(value string) Element {
	return e.Attr("srcdoc", value)
}

func (e Element) SrcLangAttr(value string) Element {
	return e.Attr("srclang", value)
}

func (e Element) SrcSetAttr(value string) Element {
	return e.Attr("srcset", value)
}

func (e Element) StartAttr(start int) Element {
	return e.Attr("start", fmt.Sprint(start))
}

func (e Element) StepAttr(value string) Element {
	return e.Attr("step", value)
}

func (e Element) TargetAttr(value string) Element {
	return e.Attr("target", value)
}

func (e Element) TypeAttr(typ string) Element {
	return e.Attr("type", typ)
}

func (e Element) UseMapAttr(value string) Element {
	return e.Attr("usemap", value)
}

func (e Element) ValueAttr(val string) Element {
	return e.Attr("value", val)
}

func (e Element) WidthAttr(width int) Element {
	return e.Attr("width", fmt.Sprint(width))
}

func (e Element) WrapAttr(value string) Element {
	return e.Attr("wrap", value)
}

// Attributable Implementation
func (e Element) AddAttribute(key, value string) Element {
	// Deprecated: Use Attr instead.
	e.attrs[key] = value
	return e
}

// Render Methods for Element
func (e Element) Render() string {
	var b strings.Builder
	e.RenderStream(&b)
	return b.String()
}

func (e Element) RenderStream(w io.Writer) error {
	if e.isRoot {
		if _, err := io.WriteString(w, "<!DOCTYPE html>"); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "<"); err != nil {
		return err
	}
	if _, err := io.WriteString(w, e.tag); err != nil {
		return err
	}
	for k, v := range e.attrs {
		if _, err := fmt.Fprintf(w, ` %s="%s"`, k, escapeInternal(v)); err != nil {
			return err
		}
	}
	if e.isVoid {
		if _, err := io.WriteString(w, ">"); err != nil {
			return err
		}
		return nil
	}
	if _, err := io.WriteString(w, ">"); err != nil {
		return err
	}
	if e.text != "" {
		if _, err := io.WriteString(w, escapeInternal(e.text)); err != nil {
			return err
		}
	}
	for _, c := range e.children {
		if err := c.RenderStream(w); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "</"); err != nil {
		return err
	}
	if _, err := io.WriteString(w, e.tag); err != nil {
		return err
	}
	if _, err := io.WriteString(w, ">"); err != nil {
		return err
	}
	return nil
}

// Helper Functions
func appendClassInternal(existing, newClass string) string {
	if existing == "" {
		return newClass
	}
	return existing + " " + newClass
}

func escapeInternal(s string) string {
	return html.EscapeString(s)
}
