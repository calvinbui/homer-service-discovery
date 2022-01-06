package homer

// created at commit https://github.com/bastienwirtz/homer/blob/66eace9e95d1962b437154b95e8f206d0da658ec/docs/configuration.md

type Config struct {
	Title             string      `yaml:"title,omitempty"`
	Subtitle          string      `yaml:"subtitle,omitempty"`
	Logo              string      `yaml:"logo,omitempty"`
	Icon              string      `yaml:"icon,omitempty"`
	Header            bool        `yaml:"header"`
	Footer            interface{} `yaml:"footer,omitempty"`
	Columns           string      `yaml:"columns,omitempty"`
	ConnectivityCheck string      `yaml:"connectivityCheck,omitempty"`
	Stylesheet        []string    `yaml:"stylesheet,omitempty"`
	Theme             string      `yaml:"theme,omitempty"`
	Colors            Colors      `yaml:"colors,omitempty"`
	Message           Message     `yaml:"message,omitempty"`
	Links             []Link      `yaml:"links,omitempty"`
	Services          []Service   `yaml:"services,omitempty"`
}

type Colors struct {
	Light Color `yaml:"light,omitempty"`
	Dark  Color `yaml:"dark,omitempty"`
}

type Color struct {
	HighlightPrimary   string `yaml:"highlight-primary,omitempty"`
	HighlightSecondary string `yaml:"highlight-secondary,omitempty"`
	HighlightHover     string `yaml:"highlight-hover,omitempty"`
	Background         string `yaml:"background,omitempty"`
	CardBackground     string `yaml:"card-background,omitempty"`
	Text               string `yaml:"text,omitempty"`
	TextHeader         string `yaml:"text-header,omitempty"`
	TextTitle          string `yaml:"text-title,omitempty"`
	TextSubtitle       string `yaml:"text-subtitle,omitempty"`
	CardShadow         string `yaml:"card-shadow,omitempty"`
	LinkHover          string `yaml:"link-hover,omitempty"`
	BackgroundImage    string `yaml:"background-image,omitempty"`
}

type Message struct {
	Style            string  `yaml:"style,omitempty"`
	Title            string  `yaml:"title,omitempty"`
	Icon             string  `yaml:"icon,omitempty"`
	Content          string  `yaml:"content,omitempty"`
	Url              string  `yaml:"url,omitempty"`             // optional
	Mapping          Mapping `yaml:"mapping,omitempty"`         // optional
	RefreshInternval int     `yaml:"refreshInterval,omitempty"` // optional
}

type Mapping struct {
	Title   string `yaml:"title,omitempty"`
	Content string `yaml:"content,omitempty"`
}

type Link struct {
	Name   string `yaml:"name,omitempty"`
	Icon   string `yaml:"icon,omitempty"`
	Url    string `yaml:"url,omitempty"`
	Target string `yaml:"target,omitempty"` // optional
}

type Service struct {
	Name     string `yaml:"name,omitempty"`
	Icon     string `yaml:"icon,omitempty"`
	Items    []Item `yaml:"items,omitempty"`
	Priority int    `yaml:"priority,omitempty"` // ours
}

type Item struct {
	Name       string `yaml:"name,omitempty"`
	Logo       string `yaml:"logo,omitempty"`
	Icon       string `yaml:"icon,omitempty"`
	Subtitle   string `yaml:"subtitle,omitempty"`
	Tag        string `yaml:"tag,omitempty"`
	Url        string `yaml:"url,omitempty"`
	Target     string `yaml:"target,omitempty"`
	Tagstyle   string `yaml:"tagstyle,omitempty"`   // optional
	Type       string `yaml:"type,omitempty"`       // optional
	Class      string `yaml:"class,omitempty"`      // optional
	Background string `yaml:"background,omitempty"` // optional
	Priority   int    `yaml:"priority,omitempty"`   // ours
}
