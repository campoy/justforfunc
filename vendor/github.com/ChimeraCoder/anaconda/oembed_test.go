package anaconda_test

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/ChimeraCoder/anaconda"
)

func TestOEmbed(t *testing.T) {
	// It is the only one that can be tested without auth
	// However, it is still rate-limited
	api := anaconda.NewTwitterApi("", "")
	api.SetBaseUrl(testBase)
	o, err := api.GetOEmbed(url.Values{"id": []string{"99530515043983360"}})
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(o, expectedOEmbed) {
		t.Errorf("Actual OEmbed differs expected:\n%#v\n Got: \n%#v\n", expectedOEmbed, o)
	}
}

var expectedOEmbed anaconda.OEmbed = anaconda.OEmbed{
	Cache_age:     "3153600000",
	Url:           "https://twitter.com/twitter/status/99530515043983360",
	Height:        0,
	Provider_url:  "https://twitter.com",
	Provider_name: "Twitter",
	Author_name:   "Twitter",
	Version:       "1.0",
	Author_url:    "https://twitter.com/twitter",
	Type:          "rich",
	Html: `<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Cool! “<a href="https://twitter.com/tw1tt3rart">@tw1tt3rart</a>: <a href="https://twitter.com/hashtag/TWITTERART?src=hash">#TWITTERART</a> ╱╱╱╱╱╱╱╱ ╱╱╭━━━━╮╱╱╭━━━━╮ ╱╱┃▇┆┆▇┃╱╭┫ⓦⓔⓔⓚ┃ ╱╱┃▽▽▽▽┃━╯┃♡ⓔⓝⓓ┃ ╱╭┫△△△△┣╮╱╰━━━━╯ ╱┃┃┆┆┆┆┃┃╱╱╱╱╱╱ ╱┗┫┆┏┓┆┣┛╱╱╱╱╱”</p>&mdash; Twitter (@twitter) <a href="https://twitter.com/twitter/status/99530515043983360">August 5, 2011</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>`,
	Width: 550,
}
