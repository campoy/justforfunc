package anaconda_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

var CONSUMER_KEY = os.Getenv("CONSUMER_KEY")
var CONSUMER_SECRET = os.Getenv("CONSUMER_SECRET")
var ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")

var api *anaconda.TwitterApi

var testBase string

func init() {
	// Initialize api so it can be used even when invidual tests are run in isolation
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api = anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	if CONSUMER_KEY != "" && CONSUMER_SECRET != "" && ACCESS_TOKEN != "" && ACCESS_TOKEN_SECRET != "" {
		return
	}

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	parsed, _ := url.Parse(server.URL)
	testBase = parsed.String()
	api.SetBaseUrl(testBase)

	var endpointElems [][]string
	filepath.Walk("json", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			elems := strings.Split(path, string(os.PathSeparator))[1:]
			endpointElems = append(endpointElems, elems)
		}

		return nil
	})

	for _, elems := range endpointElems {
		endpoint := strings.Replace("/"+path.Join(elems...), "_id_", "?id=", -1)
		filename := filepath.Join(append([]string{"json"}, elems...)...)

		mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			// if one filename is the prefix of another, the prefix will always match
			// check if there is a more specific filename that matches this request

			// create local variable to avoid closing over `filename`
			sourceFilename := filename

			r.ParseForm()
			form := strings.Replace(r.Form.Encode(), "=", "_", -1)
			form = strings.Replace(form, "&", "_", -1)
			specific := sourceFilename + "_" + form
			_, err := os.Stat(specific)
			if err == nil {
				sourceFilename = specific
			} else {
				if err != nil && !os.IsNotExist(err) {
					fmt.Fprintf(w, "error: %s", err)
					return
				}
			}

			f, err := os.Open(sourceFilename)
			if err != nil {
				// either the file does not exist
				// or something is seriously wrong with the testing environment
				fmt.Fprintf(w, "error: %s", err)
			}
			defer f.Close()

			// TODO not a hack
			if sourceFilename == filepath.Join("json", "statuses", "show.json_id_404409873170841600_tweet_mode_extended") {
				bts, err := ioutil.ReadAll(f)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)

				}
				http.Error(w, string(bts), http.StatusNotFound)
				return

			}

			io.Copy(w, f)
		})
	}
}

// Test_TwitterCredentials tests that non-empty Twitter credentials are set
// Without this, all following tests will fail
func Test_TwitterCredentials(t *testing.T) {
	if CONSUMER_KEY == "" || CONSUMER_SECRET == "" || ACCESS_TOKEN == "" || ACCESS_TOKEN_SECRET == "" {
		t.Logf("Using HTTP mock responses (API credentials are invalid: at least one is empty)")
	} else {
		t.Logf("Tests will query the live Twitter API (API credentials are all non-empty)")
	}
}

// Test that creating a TwitterApi client creates a client with non-empty OAuth credentials
func Test_TwitterApi_NewTwitterApi(t *testing.T) {
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	apiLocal := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

	if apiLocal.Credentials == nil {
		t.Fatalf("Twitter Api client has empty (nil) credentials")
	}
}

// Test that the GetSearch function actually works and returns non-empty results
func Test_TwitterApi_GetSearch(t *testing.T) {
	search_result, err := api.GetSearch("golang", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Unless something is seriously wrong, there should be at least two tweets
	if len(search_result.Statuses) < 2 {
		t.Fatalf("Expected 2 or more tweets, and found %d", len(search_result.Statuses))
	}

	// Check that at least one tweet is non-empty
	for _, tweet := range search_result.Statuses {
		if tweet.Text != "" {
			return
		}
		fmt.Print(tweet.Text)
	}

	t.Fatalf("All %d tweets had empty text", len(search_result.Statuses))
}

// Test that a valid user can be fetched
// and that unmarshalling works properly
func Test_GetUser(t *testing.T) {
	const username = "chimeracoder"

	users, err := api.GetUsersLookup(username, nil)
	if err != nil {
		t.Fatalf("GetUsersLookup returned error: %s", err.Error())
	}

	if len(users) != 1 {
		t.Fatalf("Expected one user and received %d", len(users))
	}

	// If all attributes are equal to the zero value for that type,
	// then the original value was not valid
	if reflect.DeepEqual(users[0], anaconda.User{}) {
		t.Fatalf("Received %#v", users[0])
	}
}

func Test_GetFavorites(t *testing.T) {
	v := url.Values{}
	v.Set("screen_name", "chimeracoder")
	favorites, err := api.GetFavorites(v)
	if err != nil {
		t.Fatalf("GetFavorites returned error: %s", err.Error())
	}

	if len(favorites) == 0 {
		t.Fatalf("GetFavorites returned no favorites")
	}

	if reflect.DeepEqual(favorites[0], anaconda.Tweet{}) {
		t.Fatalf("GetFavorites returned %d favorites and the first one was empty", len(favorites))
	}
}

// Test that a valid tweet can be fetched properly
// and that unmarshalling of tweet works without error
func Test_GetTweet(t *testing.T) {
	const tweetId = 303777106620452864
	const tweetText = `golang-syd is in session. Dave Symonds is now talking about API design and protobufs. #golang http://t.co/eSq3ROwu`

	tweet, err := api.GetTweet(tweetId, nil)
	if err != nil {
		t.Fatalf("GetTweet returned error: %s", err.Error())
	}

	if tweet.Text != tweetText {
		t.Fatalf("Tweet %d contained incorrect text. Received: %s", tweetId, tweet.Text)
	}

	// Check the entities
	expectedEntities := anaconda.Entities{Hashtags: []struct {
		Indices []int
		Text    string
	}{struct {
		Indices []int
		Text    string
	}{Indices: []int{86, 93}, Text: "golang"}}, Urls: []struct {
		Indices      []int
		Url          string
		Display_url  string
		Expanded_url string
	}{}, User_mentions: []struct {
		Name        string
		Indices     []int
		Screen_name string
		Id          int64
		Id_str      string
	}{}, Media: []anaconda.EntityMedia{anaconda.EntityMedia{
		Id:              303777106628841472,
		Id_str:          "303777106628841472",
		Media_url:       "http://pbs.twimg.com/media/BDc7q0OCEAAoe2C.jpg",
		Media_url_https: "https://pbs.twimg.com/media/BDc7q0OCEAAoe2C.jpg",
		Url:             "http://t.co/eSq3ROwu",
		Display_url:     "pic.twitter.com/eSq3ROwu",
		Expanded_url:    "http://twitter.com/go_nuts/status/303777106620452864/photo/1",
		Sizes: anaconda.MediaSizes{Medium: anaconda.MediaSize{W: 600,
			H:      450,
			Resize: "fit"},
			Thumb: anaconda.MediaSize{W: 150,
				H:      150,
				Resize: "crop"},
			Small: anaconda.MediaSize{W: 340,
				H:      255,
				Resize: "fit"},
			Large: anaconda.MediaSize{W: 1024,
				H:      768,
				Resize: "fit"}},
		Type: "photo",
		Indices: []int{94,
			114}}}}
	if !reflect.DeepEqual(tweet.Entities, expectedEntities) {
		t.Fatalf("Tweet entities differ")
	}
}

func Test_GetQuotedTweet(t *testing.T) {
	const tweetId = 738567564641599489
	const tweetText = `Well, this has certainly come a long way! https://t.co/QomzRzwcti`
	const quotedID = 284377451625340928
	const quotedText = `Just created gojson - a simple tool for turning JSON data into Go structs! http://t.co/QM6k9AUV #golang`

	tweet, err := api.GetTweet(tweetId, nil)
	if err != nil {
		t.Fatalf("GetTweet returned error: %s", err.Error())
	}

	if tweet.Text != tweetText {
		t.Fatalf("Tweet %d contained incorrect text. Received: %s", tweetId, tweet.Text)
	}

	if tweet.QuotedStatusID != quotedID {
		t.Fatalf("Expected quoted status %d, received %d", quotedID, tweet.QuotedStatusID)
	}

	if tweet.QuotedStatusIdStr != strconv.Itoa(quotedID) {
		t.Fatalf("Expected quoted status %d (as string), received %s", quotedID, tweet.QuotedStatusIdStr)
	}

	if tweet.QuotedStatus.Text != quotedText {
		t.Fatalf("Expected quoted status text %#v, received $#v", quotedText, tweet.QuotedStatus.Text)
	}
}

// This assumes that the current user has at least two pages' worth of followers
func Test_GetFollowersListAll(t *testing.T) {
	result := api.GetFollowersListAll(nil)
	i := 0

	for page := range result {
		if i == 2 {
			return
		}

		if page.Error != nil {
			t.Fatalf("Receved error from GetFollowersListAll: %s", page.Error)
		}

		if page.Followers == nil || len(page.Followers) == 0 {
			t.Fatalf("Received invalid value for page %d of followers: %v", i, page.Followers)
		}
		i++
	}
}

// This assumes that the current user has at least two pages' worth of followers
func Test_GetFollowersIdsAll(t *testing.T) {
	result := api.GetFollowersIdsAll(nil)
	i := 0

	for page := range result {
		if i == 2 {
			return
		}

		if page.Error != nil {
			t.Fatalf("Receved error from GetFollowersIdsAll: %s", page.Error)
		}

		if page.Ids == nil || len(page.Ids) == 0 {
			t.Fatalf("Received invalid value for page %d of followers: %v", i, page.Ids)
		}
		i++
	}
}

// This assumes that the current user has at least two pages' worth of friends
func Test_GetFriendsIdsAll(t *testing.T) {
	result := api.GetFriendsIdsAll(nil)
	i := 0

	for page := range result {
		if i == 2 {
			return
		}

		if page.Error != nil {
			t.Fatalf("Receved error from GetFriendsIdsAll : %s", page.Error)
		}

		if page.Ids == nil || len(page.Ids) == 0 {
			t.Fatalf("Received invalid value for page %d of friends : %v", i, page.Ids)
		}
		i++
	}
}

// Test that setting the delay actually changes the stored delay value
func Test_TwitterApi_SetDelay(t *testing.T) {
	const OLD_DELAY = 1 * time.Second
	const NEW_DELAY = 20 * time.Second
	api.EnableThrottling(OLD_DELAY, 4)

	delay := api.GetDelay()
	if delay != OLD_DELAY {
		t.Fatalf("Expected initial delay to be the default delay (%s)", anaconda.DEFAULT_DELAY.String())
	}

	api.SetDelay(NEW_DELAY)

	if newDelay := api.GetDelay(); newDelay != NEW_DELAY {
		t.Fatalf("Attempted to set delay to %s, but delay is now %s (original delay: %s)", NEW_DELAY, newDelay, delay)
	}
}

func Test_TwitterApi_TwitterErrorDoesNotExist(t *testing.T) {

	// Try fetching a tweet that no longer exists (was deleted)
	const DELETED_TWEET_ID = 404409873170841600

	tweet, err := api.GetTweet(DELETED_TWEET_ID, nil)
	if err == nil {
		t.Fatalf("Expected an error when fetching tweet with id %d but got none - tweet object is %+v", DELETED_TWEET_ID, tweet)
	}

	apiErr, ok := err.(*anaconda.ApiError)
	if !ok {
		t.Fatalf("Expected an *anaconda.ApiError, and received error message %s, (%+v)", err.Error(), err)
	}

	terr, ok := apiErr.Decoded.First().(anaconda.TwitterError)

	if !ok {
		t.Fatalf("TwitterErrorResponse.First() should return value of type TwitterError, not %s", reflect.TypeOf(apiErr.Decoded.First()))
	}

	if code := terr.Code; code != anaconda.TwitterErrorDoesNotExist && code != anaconda.TwitterErrorDoesNotExist2 {
		if code == anaconda.TwitterErrorRateLimitExceeded {
			t.Fatalf("Rate limit exceeded during testing - received error code %d instead of %d", anaconda.TwitterErrorRateLimitExceeded, anaconda.TwitterErrorDoesNotExist)
		} else {

			t.Fatalf("Expected Twitter to return error code %d, and instead received error code %d", anaconda.TwitterErrorDoesNotExist, code)
		}
	}
}

// Test that the client can be used to throttle to an arbitrary duration
func Test_TwitterApi_Throttling(t *testing.T) {
	const MIN_DELAY = 15 * time.Second

	api.EnableThrottling(MIN_DELAY, 5)
	oldDelay := api.GetDelay()
	api.SetDelay(MIN_DELAY)

	now := time.Now()
	_, err := api.GetSearch("golang", nil)
	if err != nil {
		t.Fatalf("GetSearch yielded error %s", err.Error())
	}
	_, err = api.GetSearch("anaconda", nil)
	if err != nil {
		t.Fatalf("GetSearch yielded error %s", err.Error())
	}
	after := time.Now()

	if difference := after.Sub(now); difference < MIN_DELAY {
		t.Fatalf("Expected delay of at least %s. Actual delay: %s", MIN_DELAY.String(), difference.String())
	}

	// Reset the delay to its previous value
	api.SetDelay(oldDelay)
}

func Test_DMScreenName(t *testing.T) {
	to, err := api.GetSelf(url.Values{})
	if err != nil {
		t.Error(err)
	}
	_, err = api.PostDMToScreenName("Test the anaconda lib", to.ScreenName)
	if err != nil {
		t.Error(err)
		return
	}
}
