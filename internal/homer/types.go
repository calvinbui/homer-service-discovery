package homer

// created at commit https://github.com/bastienwirtz/homer/blob/66eace9e95d1962b437154b95e8f206d0da658ec/docs/configuration.md

type Config struct {
	Title             string    `yaml:"title"`
	Subtitle          string    `yaml:"subtitle"`
	Logo              string    `yaml:"logo"`
	Header            string    `yaml:"header"`
	Footer            string    `yaml:"footer"`
	Columns           string    `yaml:"columns"`
	ConnectivityCheck string    `yaml:"connectivityCheck"`
	Stylesheet        []string  `yaml:"stylesheet"`
	Theme             string    `yaml:"theme"`
	Colors            Colors    `yaml:"colors"`
	Message           Message   `yaml:"message"`
	Links             []Link    `yaml:"links"`
	Services          []Service `yaml:"services"`
}

type Colors struct {
	Light Color `yaml:"light"`
	Dark  Color `yaml:"dark"`
}

type Color struct {
	HighlightPrimary   string `yaml:"highlight-primary"`
	HighlightSecondary string `yaml:"highlight-secondary"`
	HighlightHover     string `yaml:"highlight-hover"`
	Background         string `yaml:"background"`
	CardBackground     string `yaml:"card-background"`
	Text               string `yaml:"text"`
	TextHeader         string `yaml:"text-header"`
	TextTitle          string `yaml:"text-title"`
	TextSubtitle       string `yaml:"text-subtitle"`
	CardShadow         string `yaml:"card-shadow"`
	LinkHover          string `yaml:"link-hover"`
	BackgroundImage    string `yaml:"background-image"`
}

type Message struct {
	Style            string  `yaml:"style"`
	Title            string  `yaml:"title"`
	Icon             string  `yaml:"icon"`
	Content          string  `yaml:"content"`
	Url              string  `yaml:"url"`             // optional
	Mapping          Mapping `yaml:"mapping"`         // optional
	RefreshInternval int     `yaml:"refreshInterval"` // optional
}

type Mapping struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

type Link struct {
	Name   string `yaml:"name"`
	Icon   string `yaml:"icon"`
	Url    string `yaml:"url"`
	Target string `yaml:"target"` // optional
}

type Service struct {
	Name  string `yaml:"name"`
	Icon  string `yaml:"icon"`
	Items []Item `yaml:"items"`
}

type Item struct {
	Name       string `yaml:"name"`
	Logo       string `yaml:"logo"`
	Icon       string `yaml:"icon"`
	Subtitle   string `yaml:"subtitle"`
	Tag        string `yaml:"tag"`
	Url        string `yaml:"url"`
	Target     string `yaml:"target"`
	TagStyle   string `yaml:"tagstyle"`   // optional
	Type       string `yaml:"type"`       // optional
	Class      string `yaml:"class"`      // optional
	Background string `yaml:"background"` // optional
}
