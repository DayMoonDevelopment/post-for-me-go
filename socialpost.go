// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/post-for-me-go/internal/apijson"
	"github.com/stainless-sdks/post-for-me-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/post-for-me-go/internal/encoding/json"
	"github.com/stainless-sdks/post-for-me-go/internal/requestconfig"
	"github.com/stainless-sdks/post-for-me-go/option"
	"github.com/stainless-sdks/post-for-me-go/packages/param"
	"github.com/stainless-sdks/post-for-me-go/packages/respjson"
)

// SocialPostService contains methods and other services that help with interacting
// with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSocialPostService] method instead.
type SocialPostService struct {
	Options []option.RequestOption
}

// NewSocialPostService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSocialPostService(opts ...option.RequestOption) (r SocialPostService) {
	r = SocialPostService{}
	r.Options = opts
	return
}

// Create Post
func (r *SocialPostService) New(ctx context.Context, body SocialPostNewParams, opts ...option.RequestOption) (res *SocialPost, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/social-posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Post by ID
func (r *SocialPostService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialPost, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Post
func (r *SocialPostService) Update(ctx context.Context, id string, body SocialPostUpdateParams, opts ...option.RequestOption) (res *SocialPost, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get a paginated result for posts based on the applied filters
func (r *SocialPostService) List(ctx context.Context, query SocialPostListParams, opts ...option.RequestOption) (res *SocialPostListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/social-posts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Post
func (r *SocialPostService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialPostDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-posts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The properties Caption, SocialAccounts are required.
type CreateSocialPostParam struct {
	// Caption text for the post
	Caption string `json:"caption,required"`
	// Array of social account IDs for posting
	SocialAccounts []string `json:"social_accounts,omitzero,required"`
	// Array of social account IDs for posting
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// If isDraft is set then the post will not be processed
	IsDraft param.Opt[bool] `json:"isDraft,omitzero"`
	// Scheduled date and time for the post, setting to null or undefined will post
	// instantly
	ScheduledAt param.Opt[time.Time] `json:"scheduled_at,omitzero" format:"date-time"`
	// Account-specific configurations for the post
	AccountConfigurations []CreateSocialPostAccountConfigurationParam `json:"account_configurations,omitzero"`
	// Array of media URLs associated with the post
	Media []CreateSocialPostMediaParam `json:"media,omitzero"`
	// Platform-specific configurations for the post
	PlatformConfigurations CreateSocialPostPlatformConfigurationsParam `json:"platform_configurations,omitzero"`
	paramObj
}

func (r CreateSocialPostParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Configuration, SocialAccountID are required.
type CreateSocialPostAccountConfigurationParam struct {
	// Configuration for the social account
	Configuration CreateSocialPostAccountConfigurationConfigurationParam `json:"configuration,omitzero,required"`
	// ID of the social account, you want to apply the configuration to
	SocialAccountID string `json:"social_account_id,required"`
	paramObj
}

func (r CreateSocialPostAccountConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostAccountConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostAccountConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the social account
type CreateSocialPostAccountConfigurationConfigurationParam struct {
	// Allow comments on TikTok
	AllowComment param.Opt[bool] `json:"allow_comment,omitzero"`
	// Allow duets on TikTok
	AllowDuet param.Opt[bool] `json:"allow_duet,omitzero"`
	// Allow stitch on TikTok
	AllowStitch param.Opt[bool] `json:"allow_stitch,omitzero"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent param.Opt[bool] `json:"disclose_branded_content,omitzero"`
	// Disclose your brand on TikTok
	DiscloseYourBrand param.Opt[bool] `json:"disclose_your_brand,omitzero"`
	// Pinterest post link
	Link param.Opt[string] `json:"link,omitzero"`
	// Sets the privacy status for TikTok (private, public)
	PrivacyStatus param.Opt[string] `json:"privacy_status,omitzero"`
	// Overrides the `title` from the post
	Title param.Opt[string] `json:"title,omitzero"`
	// Pinterest board IDs
	BoardIDs []string `json:"board_ids,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Threads post location
	//
	// Any of "reels", "timeline".
	Location string `json:"location,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostAccountConfigurationConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostAccountConfigurationConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostAccountConfigurationConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateSocialPostAccountConfigurationConfigurationParam](
		"location", "reels", "timeline",
	)
}

// The property URL is required.
type CreateSocialPostMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r CreateSocialPostMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Platform-specific configurations for the post
type CreateSocialPostPlatformConfigurationsParam struct {
	// Bluesky configuration
	Bluesky CreateSocialPostPlatformConfigurationsBlueskyParam `json:"bluesky,omitzero"`
	// Facebook configuration
	Facebook CreateSocialPostPlatformConfigurationsFacebookParam `json:"facebook,omitzero"`
	// Instagram configuration
	Instagram CreateSocialPostPlatformConfigurationsInstagramParam `json:"instagram,omitzero"`
	// LinkedIn configuration
	Linkedin CreateSocialPostPlatformConfigurationsLinkedinParam `json:"linkedin,omitzero"`
	// Pinterest configuration
	Pinterest CreateSocialPostPlatformConfigurationsPinterestParam `json:"pinterest,omitzero"`
	// Threads configuration
	Threads CreateSocialPostPlatformConfigurationsThreadsParam `json:"threads,omitzero"`
	// Twitter configuration
	X CreateSocialPostPlatformConfigurationsXParam `json:"x,omitzero"`
	// YouTube configuration
	Youtube CreateSocialPostPlatformConfigurationsYoutubeParam `json:"youtube,omitzero"`
	// TikTok configuration
	Tiktok TiktokConfigurationParam `json:"tiktok,omitzero"`
	// TikTok configuration
	TiktokBusiness TiktokConfigurationParam `json:"tiktok_business,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bluesky configuration
type CreateSocialPostPlatformConfigurationsBlueskyParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsBlueskyParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsBlueskyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsBlueskyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Facebook configuration
type CreateSocialPostPlatformConfigurationsFacebookParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsFacebookParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsFacebookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsFacebookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instagram configuration
type CreateSocialPostPlatformConfigurationsInstagramParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsInstagramParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsInstagramParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsInstagramParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LinkedIn configuration
type CreateSocialPostPlatformConfigurationsLinkedinParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsLinkedinParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsLinkedinParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsLinkedinParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pinterest configuration
type CreateSocialPostPlatformConfigurationsPinterestParam struct {
	// Pinterest post link
	Link param.Opt[string] `json:"link,omitzero"`
	// Pinterest board IDs
	BoardIDs []string `json:"board_ids,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsPinterestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsPinterestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsPinterestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threads configuration
type CreateSocialPostPlatformConfigurationsThreadsParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Threads post location
	//
	// Any of "reels", "timeline".
	Location string `json:"location,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsThreadsParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsThreadsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsThreadsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CreateSocialPostPlatformConfigurationsThreadsParam](
		"location", "reels", "timeline",
	)
}

// Twitter configuration
type CreateSocialPostPlatformConfigurationsXParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsXParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsXParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsXParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// YouTube configuration
type CreateSocialPostPlatformConfigurationsYoutubeParam struct {
	// Overrides the `title` from the post
	Title param.Opt[string] `json:"title,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r CreateSocialPostPlatformConfigurationsYoutubeParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPlatformConfigurationsYoutubeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPlatformConfigurationsYoutubeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPost struct {
	// Unique identifier of the post
	ID string `json:"id,required"`
	// Account-specific configurations for the post
	AccountConfigurations []any `json:"account_configurations,required"`
	// Caption text for the post
	Caption string `json:"caption,required"`
	// Timestamp when the post was created
	CreatedAt string `json:"created_at,required"`
	// Provided unique identifier of the post
	ExternalID string `json:"external_id,required"`
	// Array of media URLs associated with the post
	Media any `json:"media,required"`
	// Platform-specific configurations for the post
	PlatformConfigurations any `json:"platform_configurations,required"`
	// Scheduled date and time for the post
	ScheduledAt string `json:"scheduled_at,required"`
	// Array of social account IDs for posting
	SocialAccounts []string `json:"social_accounts,required"`
	// Current status of the post: draft, processed, scheduled, or processing
	//
	// Any of "draft", "scheduled", "processing", "processed".
	Status SocialPostStatus `json:"status,required"`
	// Timestamp when the post was last updated
	UpdatedAt string `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		AccountConfigurations  respjson.Field
		Caption                respjson.Field
		CreatedAt              respjson.Field
		ExternalID             respjson.Field
		Media                  respjson.Field
		PlatformConfigurations respjson.Field
		ScheduledAt            respjson.Field
		SocialAccounts         respjson.Field
		Status                 respjson.Field
		UpdatedAt              respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPost) RawJSON() string { return r.JSON.raw }
func (r *SocialPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the post: draft, processed, scheduled, or processing
type SocialPostStatus string

const (
	SocialPostStatusDraft      SocialPostStatus = "draft"
	SocialPostStatusScheduled  SocialPostStatus = "scheduled"
	SocialPostStatusProcessing SocialPostStatus = "processing"
	SocialPostStatusProcessed  SocialPostStatus = "processed"
)

type TiktokConfigurationParam struct {
	// Allow comments on TikTok
	AllowComment param.Opt[bool] `json:"allow_comment,omitzero"`
	// Allow duets on TikTok
	AllowDuet param.Opt[bool] `json:"allow_duet,omitzero"`
	// Allow stitch on TikTok
	AllowStitch param.Opt[bool] `json:"allow_stitch,omitzero"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent param.Opt[bool] `json:"disclose_branded_content,omitzero"`
	// Disclose your brand on TikTok
	DiscloseYourBrand param.Opt[bool] `json:"disclose_your_brand,omitzero"`
	// Sets the privacy status for TikTok (private, public)
	PrivacyStatus param.Opt[string] `json:"privacy_status,omitzero"`
	// Overrides the `title` from the post
	Title param.Opt[string] `json:"title,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	paramObj
}

func (r TiktokConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow TiktokConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TiktokConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostListResponse struct {
	Data []SocialPost               `json:"data,required"`
	Meta SocialPostListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostListResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialPostListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostListResponseMeta struct {
	// Maximum number of items returned.
	Limit float64 `json:"limit,required"`
	// URL to the next page of results, or null if none.
	Next string `json:"next,required"`
	// Number of items skipped.
	Offset float64 `json:"offset,required"`
	// Total number of items available.
	Total float64 `json:"total,required"`
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
func (r SocialPostListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SocialPostListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostDeleteResponse struct {
	// Whether or not the entity was deleted
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialPostDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostNewParams struct {
	CreateSocialPost CreateSocialPostParam
	paramObj
}

func (r SocialPostNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSocialPost)
}
func (r *SocialPostNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateSocialPost)
}

type SocialPostUpdateParams struct {
	CreateSocialPost CreateSocialPostParam
	paramObj
}

func (r SocialPostUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSocialPost)
}
func (r *SocialPostUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateSocialPost)
}

type SocialPostListParams struct {
	// Number of items to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Filter by external ID. Multiple values imply OR logic.
	ExternalID []string `query:"external_id,omitzero" json:"-"`
	// Filter by platforms. Multiple values imply OR logic.
	//
	// Any of "bluesky", "facebook", "instagram", "linkedin", "pinterest", "threads",
	// "tiktok", "twitter", "youtube".
	Platform []string `query:"platform,omitzero" json:"-"`
	// Filter by post status. Multiple values imply OR logic.
	//
	// Any of "draft", "scheduled", "processing", "processed".
	Status []string `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SocialPostListParams]'s query parameters as `url.Values`.
func (r SocialPostListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
