package homer

import (
	"reflect"
	"testing"
)

func Test_unmarshalConfig(t *testing.T) {
	type args struct {
		contents []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				contents: []byte(`title: "Testing Golang"`),
			},
			want: Config{
				Title: "Testing Golang",
			},
			wantErr: false,
		},
		{
			name: "Full",
			args: args{
				contents: []byte(fullString),
			},
			want:    fullConfig,
			wantErr: false,
		},
		{
			name: "Wrong key",
			args: args{
				contents: []byte(`services: false`),
			},
			wantErr: true,
		},
		{
			name: "Non-existing key",
			args: args{
				contents: []byte(`thisIsMyRandomKey: okayGoodForYou`),
			},
			wantErr: false,
		},
		{
			name: "Empty",
			args: args{
				contents: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unmarshalConfig(tt.args.contents)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

var fullString = `
title: "App dashboard"
subtitle: "Homer"
logo: "assets/logo.png"

header: true
footer: '<p>Created with <span class="has-text-danger">❤️</span> with <a href="https://bulma.io/">bulma</a>, <a href="https://vuejs.org/">vuejs</a> & <a href="https://fontawesome.com/">font awesome</a> // Fork me on <a href="https://github.com/bastienwirtz/homer"><i class="fab fa-github-alt"></i></a></p>'

columns: "3"
connectivityCheck: true

theme: default

colors:
  light:
    highlight-primary: "#3367d6"
    highlight-secondary: "#4285f4"
    highlight-hover: "#5a95f5"
    background: "#f5f5f5"
    card-background: "#ffffff"
    text: "#363636"
    text-header: "#424242"
    text-title: "#303030"
    text-subtitle: "#424242"
    card-shadow: rgba(0, 0, 0, 0.1)
    link-hover: "#363636"
    background-image: "assets/your/light/bg.png"
  dark:
    highlight-primary: "#3367d6"
    highlight-secondary: "#4285f4"
    highlight-hover: "#5a95f5"
    background: "#131313"
    card-background: "#2b2b2b"
    text: "#eaeaea"
    text-header: "#ffffff"
    text-title: "#fafafa"
    text-subtitle: "#f5f5f5"
    card-shadow: rgba(0, 0, 0, 0.4)
    link-hover: "#ffdd57"
    background-image: "assets/your/dark/bg.png"

message:
  style: "is-warning"
  title: "Optional message!"
  icon: "fa fa-exclamation-triangle"
  content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

links:
  - name: "Link 1"
    icon: "fab fa-github"
    url: "https://github.com/bastienwirtz/homer"
    target: "_blank"
  - name: "link 2"
    icon: "fas fa-book"
    url: "https://github.com/bastienwirtz/homer"
  - name: "Second Page"
    icon: "fas fa-file-alt"
    url: "#page2"

services:
  - name: "Application"
    icon: "fas fa-code-branch"
    items:
      - name: "Awesome app"
        logo: "assets/tools/sample.png"
        subtitle: "Bookmark example"
        tag: "app"
        url: "https://www.reddit.com/r/selfhosted/"
        target: "_blank"
      - name: "Another one"
        logo: "assets/tools/sample2.png"
        subtitle: "Another application"
        tag: "app"
        tagstyle: "is-success"
        url: "#"
  - name: "Other group"
    icon: "fas fa-heartbeat"
    items:
      - name: "Pi-hole"
        logo: "assets/tools/sample.png"
        tag: "other"
        url: "http://192.168.0.151/admin"
        type: "PiHole"
        target: "_blank"
`

var fullConfig = Config{
	Title:    "App dashboard",
	Subtitle: "Homer",
	Logo:     "assets/logo.png",

	Header: "true",
	Footer: "<p>Created with <span class=\"has-text-danger\">❤️</span> with <a href=\"https://bulma.io/\">bulma</a>, <a href=\"https://vuejs.org/\">vuejs</a> & <a href=\"https://fontawesome.com/\">font awesome</a> // Fork me on <a href=\"https://github.com/bastienwirtz/homer\"><i class=\"fab fa-github-alt\"></i></a></p>",

	Columns:           "3",
	ConnectivityCheck: "true",

	Theme: "default",

	Colors: Colors{
		Light: Color{
			HighlightPrimary:   "#3367d6",
			HighlightSecondary: "#4285f4",
			HighlightHover:     "#5a95f5",
			Background:         "#f5f5f5",
			CardBackground:     "#ffffff",
			Text:               "#363636",
			TextHeader:         "#424242",
			TextTitle:          "#303030",
			TextSubtitle:       "#424242",
			CardShadow:         "rgba(0, 0, 0, 0.1)",
			LinkHover:          "#363636",
			BackgroundImage:    "assets/your/light/bg.png",
		},
		Dark: Color{
			HighlightPrimary:   "#3367d6",
			HighlightSecondary: "#4285f4",
			HighlightHover:     "#5a95f5",
			Background:         "#131313",
			CardBackground:     "#2b2b2b",
			Text:               "#eaeaea",
			TextHeader:         "#ffffff",
			TextTitle:          "#fafafa",
			TextSubtitle:       "#f5f5f5",
			CardShadow:         "rgba(0, 0, 0, 0.4)",
			LinkHover:          "#ffdd57",
			BackgroundImage:    "assets/your/dark/bg.png",
		},
	},

	Message: Message{
		Style:   "is-warning",
		Title:   "Optional message!",
		Icon:    "fa fa-exclamation-triangle",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	},

	Links: []Link{
		{
			Name:   "Link 1",
			Icon:   "fab fa-github",
			Url:    "https://github.com/bastienwirtz/homer",
			Target: "_blank",
		},
		{
			Name: "link 2",
			Icon: "fas fa-book",
			Url:  "https://github.com/bastienwirtz/homer",
		},
		{
			Name: "Second Page",
			Icon: "fas fa-file-alt",
			Url:  "#page2",
		},
	},

	Services: []Service{
		{
			Name: "Application",
			Icon: "fas fa-code-branch",
			Items: []Item{
				{
					Name:     "Awesome app",
					Logo:     "assets/tools/sample.png",
					Subtitle: "Bookmark example",
					Tag:      "app",
					Url:      "https://www.reddit.com/r/selfhosted/",
					Target:   "_blank",
				},
				{
					Name:     "Another one",
					Logo:     "assets/tools/sample2.png",
					Subtitle: "Another application",
					Tag:      "app",
					TagStyle: "is-success",
					Url:      "#",
				},
			},
		},
		{
			Name: "Other group",
			Icon: "fas fa-heartbeat",
			Items: []Item{
				{
					Name:   "Pi-hole",
					Logo:   "assets/tools/sample.png",
					Tag:    "other",
					Url:    "http://192.168.0.151/admin",
					Type:   "PiHole",
					Target: "_blank",
				},
			},
		},
	},
}
