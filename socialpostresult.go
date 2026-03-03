// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/DayMoonDevelopment/post-for-me-go/internal/apijson"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/apiquery"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/requestconfig"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/param"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/respjson"
)

// SocialPostResultService contains methods and other services that help with
// interacting with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSocialPostResultService] method instead.
type SocialPostResultService struct {
	Options []option.RequestOption
}

// NewSocialPostResultService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSocialPostResultService(opts ...option.RequestOption) (r SocialPostResultService) {
	r = SocialPostResultService{}
	r.Options = opts
	return
}

// Get post result by ID
func (r *SocialPostResultService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialPostResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-post-results/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a paginated result for post results based on the applied filters
func (r *SocialPostResultService) List(ctx context.Context, query SocialPostResultListParams, opts ...option.RequestOption) (res *SocialPostResultListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/social-post-results"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SocialPostResult struct {
	// The unique identifier of the post result
	ID string `json:"id" api:"required"`
	// Detailed logs from the post
	Details any `json:"details" api:"required"`
	// Error message if the post failed
	Error any `json:"error" api:"required"`
	// Platform-specific data
	PlatformData SocialPostResultPlatformData `json:"platform_data" api:"required"`
	// The ID of the associated post
	PostID string `json:"post_id" api:"required"`
	// The ID of the associated social account
	SocialAccountID string `json:"social_account_id" api:"required"`
	// Indicates if the post was successful
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Details         respjson.Field
		Error           respjson.Field
		PlatformData    respjson.Field
		PostID          respjson.Field
		SocialAccountID respjson.Field
		Success         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostResult) RawJSON() string { return r.JSON.raw }
func (r *SocialPostResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Platform-specific data
type SocialPostResultPlatformData struct {
	// Platform-specific ID
	ID string `json:"id"`
	// URL of the posted content
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostResultPlatformData) RawJSON() string { return r.JSON.raw }
func (r *SocialPostResultPlatformData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostResultListResponse struct {
	Data []SocialPostResult               `json:"data" api:"required"`
	Meta SocialPostResultListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostResultListResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialPostResultListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostResultListResponseMeta struct {
	// Maximum number of items returned.
	Limit float64 `json:"limit" api:"required"`
	// URL to the next page of results, or null if none.
	Next string `json:"next" api:"required"`
	// Number of items skipped.
	Offset float64 `json:"offset" api:"required"`
	// Total number of items available.
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		Next        respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostResultListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SocialPostResultListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostResultListParams struct {
	// Number of items to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Filter by platform(s). Multiple values imply OR logic (e.g.,
	// ?platform=x&platform=facebook).
	Platform []string `query:"platform,omitzero" json:"-"`
	// Filter by post IDs. Multiple values imply OR logic (e.g.,
	// ?post_id=123&post_id=456).
	PostID []string `query:"post_id,omitzero" json:"-"`
	// Filter by social account ID(s). Multiple values imply OR logic (e.g.,
	// ?social_account_id=123&social_account_id=456).
	SocialAccountID []string `query:"social_account_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SocialPostResultListParams]'s query parameters as
// `url.Values`.
func (r SocialPostResultListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
