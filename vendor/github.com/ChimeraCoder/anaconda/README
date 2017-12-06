Anaconda
====================

[![Build Status](https://travis-ci.org/ChimeraCoder/anaconda.svg?branch=master)](https://travis-ci.org/ChimeraCoder/anaconda) [![Build Status](https://ci.appveyor.com/api/projects/status/63pi6csod8bps80i/branch/master?svg=true)](https://ci.appveyor.com/project/ChimeraCoder/anaconda/branch/master) [![GoDoc](https://godoc.org/github.com/ChimeraCoder/anaconda?status.svg)](https://godoc.org/github.com/ChimeraCoder/anaconda)

Anaconda is a simple, transparent Go package for accessing version 1.1 of the Twitter API. 

Successful API queries return native Go structs that can be used immediately, with no need for type assertions.



Examples
-------------

### Authentication

If you already have the access token (and secret) for your user (Twitter provides this for your own account on the developer portal), creating the client is simple:

````go
anaconda.SetConsumerKey("your-consumer-key")
anaconda.SetConsumerSecret("your-consumer-secret")
api := anaconda.NewTwitterApi("your-access-token", "your-access-token-secret")
````

### Queries

Queries are conducted using a pointer to an authenticated `TwitterApi` struct. In v1.1 of Twitter's API, all requests should be authenticated.

````go
searchResult, _ := api.GetSearch("golang", nil)
for _ , tweet := range searchResult.Statuses {
    fmt.Println(tweet.Text)
}   
````
Certain endpoints allow separate optional parameter; if desired, these can be passed as the final parameter. 

````go
//Perhaps we want 30 values instead of the default 15
v := url.Values{}
v.Set("count", "30")
result, err := api.GetSearch("golang", v)
````

(Remember that `url.Values` is equivalent to a `map[string][]string`, if you find that more convenient notation when specifying values). Otherwise, `nil` suffices.



Endpoints
------------

Anaconda implements most of the endpoints defined in the [Twitter API documentation](https://dev.twitter.com/docs/api/1.1). For clarity, in most cases, the function name is simply the name of the HTTP method and the endpoint (e.g., the endpoint `GET /friendships/incoming` is provided by the function `GetFriendshipsIncoming`).

In a few cases, a shortened form has been chosen to make life easier (for example, retweeting is simply the function `Retweet`)



Error Handling, Rate Limiting, and Throttling
---------------------------------

### Error Handling

Twitter errors are returned as an `ApiError`, which satisfies the `error` interface and can be treated as a vanilla `error`. However, it also contains the additional information returned by the Twitter API that may be useful in deciding how to proceed after encountering an error.


If you make queries too quickly, you may bump against Twitter's [rate limits](https://dev.twitter.com/docs/rate-limiting/1.1). If this happens, `anaconda` automatically retries the query when the rate limit resets, using the `X-Rate-Limit-Reset` header that Twitter provides to determine how long to wait.

In other words, users of the `anaconda` library should not need to handle rate limiting errors themselves; this is handled seamlessly behind-the-scenes. If an error is returned by a function, another form of error must have occurred (which can be checked by using the fields provided by the `ApiError` struct).


(If desired, this feature can be turned off by calling `ReturnRateLimitError(true)`.)


### Throttling

Anaconda now supports automatic client-side throttling of queries to avoid hitting the Twitter rate-limit.

This is currently *off* by default; however, it may be turned on by default in future versions of the library, as the implementation is improved.


To set a delay between queries, use the `SetDelay` method:

````go
    api.SetDelay(10 * time.Second)
````

Delays are set specific to each `TwitterApi` struct, so queries that use different users' access credentials are completely independent.


To turn off automatic throttling, set the delay to `0`:

````go
    api.SetDelay(0 * time.Second)
````

### Query Queue Persistence

If your code creates a NewTwitterApi in a regularly called function, you'll need to call `.Close()` on the API struct to clear the queryQueue and allow the goroutine to exit. Otherwise you could see goroutine and therefor heap memory leaks in long-running applications.

### Google App Engine

Since Google App Engine doesn't make the standard `http.Transport` available, it's necessary to tell Anaconda to use a different client context.

````go
	api = anaconda.NewTwitterApi("", "")
	c := appengine.NewContext(r)
	api.HttpClient.Transport = &urlfetch.Transport{Context: c}
````


License
-----------
Anaconda is free software licensed under the MIT/X11 license. Details provided in the LICENSE file.
