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

	"github.com/DayMoonDevelopment/post-for-me-go/internal/apijson"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/apiquery"
	shimjson "github.com/DayMoonDevelopment/post-for-me-go/internal/encoding/json"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/requestconfig"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/param"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/respjson"
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

type BlueskyConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []BlueskyConfigurationDtoMedia `json:"media,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlueskyConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *BlueskyConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BlueskyConfigurationDto to a BlueskyConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BlueskyConfigurationDtoParam.Overrides()
func (r BlueskyConfigurationDto) ToParam() BlueskyConfigurationDtoParam {
	return param.Override[BlueskyConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type BlueskyConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BlueskyConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *BlueskyConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlueskyConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []BlueskyConfigurationDtoMediaParam `json:"media,omitzero"`
	paramObj
}

func (r BlueskyConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow BlueskyConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlueskyConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type BlueskyConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r BlueskyConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow BlueskyConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BlueskyConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	PlatformConfigurations PlatformConfigurationsDtoParam `json:"platform_configurations,omitzero"`
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
	// Will automatically add music to photo posts on TikTok
	AutoAddMusic param.Opt[bool] `json:"auto_add_music,omitzero"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent param.Opt[bool] `json:"disclose_branded_content,omitzero"`
	// Disclose your brand on TikTok
	DiscloseYourBrand param.Opt[bool] `json:"disclose_your_brand,omitzero"`
	// Flag content as AI generated on TikTok
	IsAIGenerated param.Opt[bool] `json:"is_ai_generated,omitzero"`
	// Will create a draft upload to TikTok, posting will need to be completed from
	// within the app
	IsDraft param.Opt[bool] `json:"is_draft,omitzero"`
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
	// Overrides the `media` from the post
	Media []string `json:"media,omitzero"`
	// Post placement for Facebook/Instagram/Threads
	//
	// Any of "reels", "timeline", "stories".
	Placement string `json:"placement,omitzero"`
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
		"placement", "reels", "timeline", "stories",
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

type FacebookConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []FacebookConfigurationDtoMedia `json:"media,nullable"`
	// Facebook post placement
	//
	// Any of "reels", "stories", "timeline".
	Placement FacebookConfigurationDtoPlacement `json:"placement,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		Placement   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FacebookConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *FacebookConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FacebookConfigurationDto to a
// FacebookConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FacebookConfigurationDtoParam.Overrides()
func (r FacebookConfigurationDto) ToParam() FacebookConfigurationDtoParam {
	return param.Override[FacebookConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type FacebookConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FacebookConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *FacebookConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Facebook post placement
type FacebookConfigurationDtoPlacement string

const (
	FacebookConfigurationDtoPlacementReels    FacebookConfigurationDtoPlacement = "reels"
	FacebookConfigurationDtoPlacementStories  FacebookConfigurationDtoPlacement = "stories"
	FacebookConfigurationDtoPlacementTimeline FacebookConfigurationDtoPlacement = "timeline"
)

type FacebookConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []FacebookConfigurationDtoMediaParam `json:"media,omitzero"`
	// Facebook post placement
	//
	// Any of "reels", "stories", "timeline".
	Placement FacebookConfigurationDtoPlacement `json:"placement,omitzero"`
	paramObj
}

func (r FacebookConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow FacebookConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FacebookConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type FacebookConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r FacebookConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow FacebookConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FacebookConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InstagramConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Instagram usernames to be tagged as a collaborator
	Collaborators []string `json:"collaborators,nullable"`
	// Overrides the `media` from the post
	Media []InstagramConfigurationDtoMedia `json:"media,nullable"`
	// Instagram post placement
	//
	// Any of "reels", "stories", "timeline".
	Placement InstagramConfigurationDtoPlacement `json:"placement,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption       respjson.Field
		Collaborators respjson.Field
		Media         respjson.Field
		Placement     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstagramConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *InstagramConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InstagramConfigurationDto to a
// InstagramConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InstagramConfigurationDtoParam.Overrides()
func (r InstagramConfigurationDto) ToParam() InstagramConfigurationDtoParam {
	return param.Override[InstagramConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type InstagramConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InstagramConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *InstagramConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Instagram post placement
type InstagramConfigurationDtoPlacement string

const (
	InstagramConfigurationDtoPlacementReels    InstagramConfigurationDtoPlacement = "reels"
	InstagramConfigurationDtoPlacementStories  InstagramConfigurationDtoPlacement = "stories"
	InstagramConfigurationDtoPlacementTimeline InstagramConfigurationDtoPlacement = "timeline"
)

type InstagramConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Instagram usernames to be tagged as a collaborator
	Collaborators []string `json:"collaborators,omitzero"`
	// Overrides the `media` from the post
	Media []InstagramConfigurationDtoMediaParam `json:"media,omitzero"`
	// Instagram post placement
	//
	// Any of "reels", "stories", "timeline".
	Placement InstagramConfigurationDtoPlacement `json:"placement,omitzero"`
	paramObj
}

func (r InstagramConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow InstagramConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstagramConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type InstagramConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r InstagramConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow InstagramConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InstagramConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedinConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []LinkedinConfigurationDtoMedia `json:"media,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedinConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *LinkedinConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LinkedinConfigurationDto to a
// LinkedinConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LinkedinConfigurationDtoParam.Overrides()
func (r LinkedinConfigurationDto) ToParam() LinkedinConfigurationDtoParam {
	return param.Override[LinkedinConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type LinkedinConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LinkedinConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *LinkedinConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LinkedinConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []LinkedinConfigurationDtoMediaParam `json:"media,omitzero"`
	paramObj
}

func (r LinkedinConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedinConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedinConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type LinkedinConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r LinkedinConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow LinkedinConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LinkedinConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PinterestConfigurationDto struct {
	// Pinterest board IDs
	BoardIDs []string `json:"board_ids,nullable"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Pinterest post link
	Link string `json:"link,nullable"`
	// Overrides the `media` from the post
	Media []PinterestConfigurationDtoMedia `json:"media,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BoardIDs    respjson.Field
		Caption     respjson.Field
		Link        respjson.Field
		Media       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PinterestConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *PinterestConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PinterestConfigurationDto to a
// PinterestConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PinterestConfigurationDtoParam.Overrides()
func (r PinterestConfigurationDto) ToParam() PinterestConfigurationDtoParam {
	return param.Override[PinterestConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type PinterestConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PinterestConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *PinterestConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PinterestConfigurationDtoParam struct {
	// Pinterest post link
	Link param.Opt[string] `json:"link,omitzero"`
	// Pinterest board IDs
	BoardIDs []string `json:"board_ids,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []PinterestConfigurationDtoMediaParam `json:"media,omitzero"`
	paramObj
}

func (r PinterestConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow PinterestConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PinterestConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type PinterestConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r PinterestConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow PinterestConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PinterestConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformConfigurationsDto struct {
	// Bluesky configuration
	Bluesky BlueskyConfigurationDto `json:"bluesky,nullable"`
	// Facebook configuration
	Facebook FacebookConfigurationDto `json:"facebook,nullable"`
	// Instagram configuration
	Instagram InstagramConfigurationDto `json:"instagram,nullable"`
	// LinkedIn configuration
	Linkedin LinkedinConfigurationDto `json:"linkedin,nullable"`
	// Pinterest configuration
	Pinterest PinterestConfigurationDto `json:"pinterest,nullable"`
	// Threads configuration
	Threads ThreadsConfigurationDto `json:"threads,nullable"`
	// TikTok configuration
	Tiktok TiktokConfiguration `json:"tiktok,nullable"`
	// TikTok configuration
	TiktokBusiness TiktokConfiguration `json:"tiktok_business,nullable"`
	// Twitter configuration
	X TwitterConfigurationDto `json:"x,nullable"`
	// YouTube configuration
	Youtube YoutubeConfigurationDto `json:"youtube,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bluesky        respjson.Field
		Facebook       respjson.Field
		Instagram      respjson.Field
		Linkedin       respjson.Field
		Pinterest      respjson.Field
		Threads        respjson.Field
		Tiktok         respjson.Field
		TiktokBusiness respjson.Field
		X              respjson.Field
		Youtube        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformConfigurationsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformConfigurationsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlatformConfigurationsDto to a
// PlatformConfigurationsDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlatformConfigurationsDtoParam.Overrides()
func (r PlatformConfigurationsDto) ToParam() PlatformConfigurationsDtoParam {
	return param.Override[PlatformConfigurationsDtoParam](json.RawMessage(r.RawJSON()))
}

type PlatformConfigurationsDtoParam struct {
	// Bluesky configuration
	Bluesky BlueskyConfigurationDtoParam `json:"bluesky,omitzero"`
	// Facebook configuration
	Facebook FacebookConfigurationDtoParam `json:"facebook,omitzero"`
	// Instagram configuration
	Instagram InstagramConfigurationDtoParam `json:"instagram,omitzero"`
	// LinkedIn configuration
	Linkedin LinkedinConfigurationDtoParam `json:"linkedin,omitzero"`
	// Pinterest configuration
	Pinterest PinterestConfigurationDtoParam `json:"pinterest,omitzero"`
	// Threads configuration
	Threads ThreadsConfigurationDtoParam `json:"threads,omitzero"`
	// TikTok configuration
	Tiktok TiktokConfigurationParam `json:"tiktok,omitzero"`
	// TikTok configuration
	TiktokBusiness TiktokConfigurationParam `json:"tiktok_business,omitzero"`
	// Twitter configuration
	X TwitterConfigurationDtoParam `json:"x,omitzero"`
	// YouTube configuration
	Youtube YoutubeConfigurationDtoParam `json:"youtube,omitzero"`
	paramObj
}

func (r PlatformConfigurationsDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow PlatformConfigurationsDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlatformConfigurationsDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPost struct {
	// Unique identifier of the post
	ID string `json:"id,required"`
	// Account-specific configurations for the post
	AccountConfigurations []SocialPostAccountConfiguration `json:"account_configurations,required"`
	// Caption text for the post
	Caption string `json:"caption,required"`
	// Timestamp when the post was created
	CreatedAt string `json:"created_at,required"`
	// Provided unique identifier of the post
	ExternalID string `json:"external_id,required"`
	// Array of media URLs associated with the post
	Media []SocialPostMedia `json:"media,required"`
	// Platform-specific configurations for the post
	PlatformConfigurations PlatformConfigurationsDto `json:"platform_configurations,required"`
	// Scheduled date and time for the post
	ScheduledAt string `json:"scheduled_at,required"`
	// Array of social account IDs for posting
	SocialAccounts []SocialAccount `json:"social_accounts,required"`
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

type SocialPostAccountConfiguration struct {
	// Configuration for the social account
	Configuration SocialPostAccountConfigurationConfiguration `json:"configuration,required"`
	// ID of the social account, you want to apply the configuration to
	SocialAccountID string `json:"social_account_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Configuration   respjson.Field
		SocialAccountID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostAccountConfiguration) RawJSON() string { return r.JSON.raw }
func (r *SocialPostAccountConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the social account
type SocialPostAccountConfigurationConfiguration struct {
	// Allow comments on TikTok
	AllowComment bool `json:"allow_comment,nullable"`
	// Allow duets on TikTok
	AllowDuet bool `json:"allow_duet,nullable"`
	// Allow stitch on TikTok
	AllowStitch bool `json:"allow_stitch,nullable"`
	// Will automatically add music to photo posts on TikTok
	AutoAddMusic bool `json:"auto_add_music,nullable"`
	// Pinterest board IDs
	BoardIDs []string `json:"board_ids,nullable"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent bool `json:"disclose_branded_content,nullable"`
	// Disclose your brand on TikTok
	DiscloseYourBrand bool `json:"disclose_your_brand,nullable"`
	// Flag content as AI generated on TikTok
	IsAIGenerated bool `json:"is_ai_generated,nullable"`
	// Will create a draft upload to TikTok, posting will need to be completed from
	// within the app
	IsDraft bool `json:"is_draft,nullable"`
	// Pinterest post link
	Link string `json:"link,nullable"`
	// Overrides the `media` from the post
	Media []string `json:"media,nullable"`
	// Post placement for Facebook/Instagram/Threads
	//
	// Any of "reels", "timeline", "stories".
	Placement string `json:"placement,nullable"`
	// Sets the privacy status for TikTok (private, public)
	PrivacyStatus string `json:"privacy_status,nullable"`
	// Overrides the `title` from the post
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowComment           respjson.Field
		AllowDuet              respjson.Field
		AllowStitch            respjson.Field
		AutoAddMusic           respjson.Field
		BoardIDs               respjson.Field
		Caption                respjson.Field
		DiscloseBrandedContent respjson.Field
		DiscloseYourBrand      respjson.Field
		IsAIGenerated          respjson.Field
		IsDraft                respjson.Field
		Link                   respjson.Field
		Media                  respjson.Field
		Placement              respjson.Field
		PrivacyStatus          respjson.Field
		Title                  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostAccountConfigurationConfiguration) RawJSON() string { return r.JSON.raw }
func (r *SocialPostAccountConfigurationConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostMedia) RawJSON() string { return r.JSON.raw }
func (r *SocialPostMedia) UnmarshalJSON(data []byte) error {
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

type ThreadsConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []ThreadsConfigurationDtoMedia `json:"media,nullable"`
	// Threads post placement
	//
	// Any of "reels", "timeline".
	Placement ThreadsConfigurationDtoPlacement `json:"placement,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		Placement   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThreadsConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *ThreadsConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ThreadsConfigurationDto to a ThreadsConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ThreadsConfigurationDtoParam.Overrides()
func (r ThreadsConfigurationDto) ToParam() ThreadsConfigurationDtoParam {
	return param.Override[ThreadsConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type ThreadsConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ThreadsConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *ThreadsConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Threads post placement
type ThreadsConfigurationDtoPlacement string

const (
	ThreadsConfigurationDtoPlacementReels    ThreadsConfigurationDtoPlacement = "reels"
	ThreadsConfigurationDtoPlacementTimeline ThreadsConfigurationDtoPlacement = "timeline"
)

type ThreadsConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []ThreadsConfigurationDtoMediaParam `json:"media,omitzero"`
	// Threads post placement
	//
	// Any of "reels", "timeline".
	Placement ThreadsConfigurationDtoPlacement `json:"placement,omitzero"`
	paramObj
}

func (r ThreadsConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow ThreadsConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThreadsConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type ThreadsConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r ThreadsConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow ThreadsConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ThreadsConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TiktokConfiguration struct {
	// Allow comments on TikTok
	AllowComment bool `json:"allow_comment,nullable"`
	// Allow duets on TikTok
	AllowDuet bool `json:"allow_duet,nullable"`
	// Allow stitch on TikTok
	AllowStitch bool `json:"allow_stitch,nullable"`
	// Will automatically add music to photo posts
	AutoAddMusic bool `json:"auto_add_music,nullable"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent bool `json:"disclose_branded_content,nullable"`
	// Disclose your brand on TikTok
	DiscloseYourBrand bool `json:"disclose_your_brand,nullable"`
	// Flag content as AI generated on TikTok
	IsAIGenerated bool `json:"is_ai_generated,nullable"`
	// Will create a draft upload to TikTok, posting will need to be completed from
	// within the app
	IsDraft bool `json:"is_draft,nullable"`
	// Overrides the `media` from the post
	Media []TiktokConfigurationMedia `json:"media,nullable"`
	// Sets the privacy status for TikTok (private, public)
	PrivacyStatus string `json:"privacy_status,nullable"`
	// Overrides the `title` from the post
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowComment           respjson.Field
		AllowDuet              respjson.Field
		AllowStitch            respjson.Field
		AutoAddMusic           respjson.Field
		Caption                respjson.Field
		DiscloseBrandedContent respjson.Field
		DiscloseYourBrand      respjson.Field
		IsAIGenerated          respjson.Field
		IsDraft                respjson.Field
		Media                  respjson.Field
		PrivacyStatus          respjson.Field
		Title                  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TiktokConfiguration) RawJSON() string { return r.JSON.raw }
func (r *TiktokConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TiktokConfiguration to a TiktokConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TiktokConfigurationParam.Overrides()
func (r TiktokConfiguration) ToParam() TiktokConfigurationParam {
	return param.Override[TiktokConfigurationParam](json.RawMessage(r.RawJSON()))
}

type TiktokConfigurationMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TiktokConfigurationMedia) RawJSON() string { return r.JSON.raw }
func (r *TiktokConfigurationMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TiktokConfigurationParam struct {
	// Allow comments on TikTok
	AllowComment param.Opt[bool] `json:"allow_comment,omitzero"`
	// Allow duets on TikTok
	AllowDuet param.Opt[bool] `json:"allow_duet,omitzero"`
	// Allow stitch on TikTok
	AllowStitch param.Opt[bool] `json:"allow_stitch,omitzero"`
	// Will automatically add music to photo posts
	AutoAddMusic param.Opt[bool] `json:"auto_add_music,omitzero"`
	// Disclose branded content on TikTok
	DiscloseBrandedContent param.Opt[bool] `json:"disclose_branded_content,omitzero"`
	// Disclose your brand on TikTok
	DiscloseYourBrand param.Opt[bool] `json:"disclose_your_brand,omitzero"`
	// Flag content as AI generated on TikTok
	IsAIGenerated param.Opt[bool] `json:"is_ai_generated,omitzero"`
	// Will create a draft upload to TikTok, posting will need to be completed from
	// within the app
	IsDraft param.Opt[bool] `json:"is_draft,omitzero"`
	// Sets the privacy status for TikTok (private, public)
	PrivacyStatus param.Opt[string] `json:"privacy_status,omitzero"`
	// Overrides the `title` from the post
	Title param.Opt[string] `json:"title,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []TiktokConfigurationMediaParam `json:"media,omitzero"`
	paramObj
}

func (r TiktokConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow TiktokConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TiktokConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type TiktokConfigurationMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r TiktokConfigurationMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow TiktokConfigurationMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TiktokConfigurationMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TwitterConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []TwitterConfigurationDtoMedia `json:"media,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TwitterConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *TwitterConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TwitterConfigurationDto to a TwitterConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TwitterConfigurationDtoParam.Overrides()
func (r TwitterConfigurationDto) ToParam() TwitterConfigurationDtoParam {
	return param.Override[TwitterConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type TwitterConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TwitterConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *TwitterConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TwitterConfigurationDtoParam struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []TwitterConfigurationDtoMediaParam `json:"media,omitzero"`
	paramObj
}

func (r TwitterConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow TwitterConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TwitterConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type TwitterConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r TwitterConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow TwitterConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TwitterConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type YoutubeConfigurationDto struct {
	// Overrides the `caption` from the post
	Caption any `json:"caption,nullable"`
	// Overrides the `media` from the post
	Media []YoutubeConfigurationDtoMedia `json:"media,nullable"`
	// Overrides the `title` from the post
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption     respjson.Field
		Media       respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YoutubeConfigurationDto) RawJSON() string { return r.JSON.raw }
func (r *YoutubeConfigurationDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this YoutubeConfigurationDto to a YoutubeConfigurationDtoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// YoutubeConfigurationDtoParam.Overrides()
func (r YoutubeConfigurationDto) ToParam() YoutubeConfigurationDtoParam {
	return param.Override[YoutubeConfigurationDtoParam](json.RawMessage(r.RawJSON()))
}

type YoutubeConfigurationDtoMedia struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,nullable"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                  respjson.Field
		ThumbnailTimestampMs respjson.Field
		ThumbnailURL         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r YoutubeConfigurationDtoMedia) RawJSON() string { return r.JSON.raw }
func (r *YoutubeConfigurationDtoMedia) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type YoutubeConfigurationDtoParam struct {
	// Overrides the `title` from the post
	Title param.Opt[string] `json:"title,omitzero"`
	// Overrides the `caption` from the post
	Caption any `json:"caption,omitzero"`
	// Overrides the `media` from the post
	Media []YoutubeConfigurationDtoMediaParam `json:"media,omitzero"`
	paramObj
}

func (r YoutubeConfigurationDtoParam) MarshalJSON() (data []byte, err error) {
	type shadow YoutubeConfigurationDtoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *YoutubeConfigurationDtoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type YoutubeConfigurationDtoMediaParam struct {
	// Public URL of the media
	URL string `json:"url,required"`
	// Timestamp in milliseconds of frame to use as thumbnail for the media
	ThumbnailTimestampMs any `json:"thumbnail_timestamp_ms,omitzero"`
	// Public URL of the thumbnail for the media
	ThumbnailURL any `json:"thumbnail_url,omitzero"`
	paramObj
}

func (r YoutubeConfigurationDtoMediaParam) MarshalJSON() (data []byte, err error) {
	type shadow YoutubeConfigurationDtoMediaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *YoutubeConfigurationDtoMediaParam) UnmarshalJSON(data []byte) error {
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
	// "tiktok", "x", "youtube".
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
