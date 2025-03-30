package gojikan

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dmji/go-jikan/internal/hooks"
	"github.com/dmji/go-jikan/internal/utils"
	"github.com/dmji/go-jikan/retry"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Jikan REST API
	"https://api.jikan.moe/v4",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Jikan API: [Jikan](https://jikan.moe) is an **Unofficial** MyAnimeList API.
// It scrapes the website to satisfy the need for a complete API - which MyAnimeList lacks.
//
// # Information
//
// ⚡ Jikan is powered by its awesome backers - 🙏 [Become a backer](https://www.patreon.com/jikan)
//
// ## Rate Limiting
//
// | Duration | Requests |
// |----|----|
// | Daily | **Unlimited** |
// | Per Minute | 60 requests |
// | Per Second | 3 requests |
//
// Note: It's still possible to get rate limited from MyAnimeList.net instead.
//
// ## JSON Notes
// - Any property (except arrays or objects) whose value does not exist or is undetermined, will be `null`.
// - Any array or object property whose value does not exist or is undetermined, will be empty.
// - Any `score` property whose value does not exist or is undetermined, will be `0`.
// - All dates and timestamps are returned in [ISO8601](https://en.wikipedia.org/wiki/ISO_8601) format and in UTC timezone
//
// ## Caching
// By **CACHING**, we refer to the data parsed from MyAnimeList which is stored temporarily on our servers to provide better API performance.
//
// All requests are cached for **24 hours**.
//
// The following response headers will detail cache information.
//
// | Header | Remarks |
// | ---- | ---- |
// | `Expires` | Cache expiry date |
// | `Last-Modified` | Cache set date |
// | `X-Request-Fingerprint` | Unique request fingerprint (only for cachable requests, not queries) |
//
// Note: `X-Request-Fingerprint` will only be available on single resource requests and their child endpoints. e.g `/anime/1`, `/anime/1/relations`.
// They won't be available on pages which perform queries, like /anime, or /top/anime, etc.
//
// ## Allowed HTTP(s) requests
//
// **Jikan REST API does not provide authenticated requests for MyAnimeList.** This means you can not use it to update your anime/manga list.
// Only GET requests are supported which return READ-ONLY data.
//
// ## HTTP Responses
//
// All error responses are accompanied by a JSON Error response.
//
// | Exception | HTTP Status | Remarks |
// | ---- | ---- | ---- |
// | N/A | `200 - OK` | The request was successful |
// | N/A | `304 - Not Modified` | You have the latest data (Cache Validation response) |
// | `BadRequestException`,`ValidationException` | `400 - Bad Request` | You've made an invalid request. Recheck documentation |
// | `BadResponseException` | `404 - Not Found` | The resource was not found or MyAnimeList responded with a `404` |
// | `BadRequestException` | `405 - Method Not Allowed` | Requested Method is not supported for resource. Only `GET` requests are allowed |
// | `RateLimitException` | `429 - Too Many Request` | You are being rate limited by Jikan or MyAnimeList is rate-limiting our servers (specified in the error response) |
// | `UpstreamException`,`ParserException`,etc. | `500 - Internal Server Error` | Something didn't work. Try again later. If you see an error response with a `report_url` URL, please click on it to open an auto-generated GitHub issue |
// | `ServiceUnavailableException` | `503 - Service Unavailable` | In most cases this is intentionally done if the service is down for maintenance. |
//
// ## JSON Error Response
//
// ```json
//
//	{
//	    "status": 500,
//	    "type": "InternalException",
//	    "message": "Exception Message",
//	    "error": "Exception Trace",
//	    "report_url": "https://github.com..."
//	 }
//
// ```
//
// | Property | Remarks |
// | ---- | ---- |
// | `status` | Returned HTTP Status Code |
// | `type` | Thrown Exception |
// | `message` | Human-readable error message |
// | `error` | Error response and trace from the API |
// | `report_url` | Clicking this would redirect you to a generated GitHub issue |
//
// ## Cache Validation
//
// - All requests return a `ETag` header which is an MD5 hash of the response
// - You can use this hash to verify if there's new or updated content by suppliying it as the value for the `If-None-Match` in your next request header
// - You will get a HTTP `304 - Not Modified` response if the content has not changed
// - If the content has changed, you'll get a HTTP `200 - OK` response with the updated JSON response
//
// ![Cache Validation](https://i.imgur.com/925ozVn.png 'Cache Validation')
//
// ## Disclaimer
//
// - Jikan is not affiliated with MyAnimeList.net.
// - Jikan is a free, open-source API. Please use it responsibly.
//
// ----
//
// By using the API, you are agreeing to Jikan's [terms of use](https://jikan.moe/terms) policy.
//
// [v3 Documentation](https://jikan.docs.apiary.io/) - [Wrappers/SDKs](https://github.com/jikan-me/jikan#wrappers) - [Report an issue](https://github.com/jikan-me/jikan-rest/issues/new) - [Host your own server](https://github.com/jikan-me/jikan-rest)
//
// https://jikan.moe - About
type Jikan struct {
	Anime           *Anime
	Characters      *Characters
	Clubs           *Clubs
	Genres          *Genres
	Magazines       *Magazines
	Manga           *Manga
	People          *People
	Producers       *Producers
	Random          *Random
	Recommendations *Recommendations
	Reviews         *Reviews
	Schedules       *Schedules
	Users           *Users
	Seasons         *Seasons
	Top             *Top
	Watch           *Watch

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Jikan)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Jikan) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Jikan) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Jikan) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Jikan) {
		sdk.sdkConfiguration.Client = client
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Jikan) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Jikan) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Jikan {
	sdk := &Jikan{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "4.0.0",
			SDKVersion:        "0.1.0",
			GenVersion:        "2.563.0",
			UserAgent:         "speakeasy-sdk/go 0.1.0 2.563.0 4.0.0 github.com/dmji/go-jikan",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Anime = newAnime(sdk.sdkConfiguration)

	sdk.Characters = newCharacters(sdk.sdkConfiguration)

	sdk.Clubs = newClubs(sdk.sdkConfiguration)

	sdk.Genres = newGenres(sdk.sdkConfiguration)

	sdk.Magazines = newMagazines(sdk.sdkConfiguration)

	sdk.Manga = newManga(sdk.sdkConfiguration)

	sdk.People = newPeople(sdk.sdkConfiguration)

	sdk.Producers = newProducers(sdk.sdkConfiguration)

	sdk.Random = newRandom(sdk.sdkConfiguration)

	sdk.Recommendations = newRecommendations(sdk.sdkConfiguration)

	sdk.Reviews = newReviews(sdk.sdkConfiguration)

	sdk.Schedules = newSchedules(sdk.sdkConfiguration)

	sdk.Users = newUsers(sdk.sdkConfiguration)

	sdk.Seasons = newSeasons(sdk.sdkConfiguration)

	sdk.Top = newTop(sdk.sdkConfiguration)

	sdk.Watch = newWatch(sdk.sdkConfiguration)

	return sdk
}
