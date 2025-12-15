// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/DayMoonDevelopment/post-for-me-go/internal/apijson"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/apiquery"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/requestconfig"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/param"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/respjson"
)

// SocialAccountService contains methods and other services that help with
// interacting with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSocialAccountService] method instead.
type SocialAccountService struct {
	Options []option.RequestOption
}

// NewSocialAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSocialAccountService(opts ...option.RequestOption) (r SocialAccountService) {
	r = SocialAccountService{}
	r.Options = opts
	return
}

// If a social account with the same platform and user_id already exists, it will
// be updated. If not, a new social account will be created.
func (r *SocialAccountService) New(ctx context.Context, body SocialAccountNewParams, opts ...option.RequestOption) (res *SocialAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/social-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get social account by ID
func (r *SocialAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update social account
func (r *SocialAccountService) Update(ctx context.Context, id string, body SocialAccountUpdateParams, opts ...option.RequestOption) (res *SocialAccount, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-accounts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated result for social accounts based on the applied filters
func (r *SocialAccountService) List(ctx context.Context, query SocialAccountListParams, opts ...option.RequestOption) (res *SocialAccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/social-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Generates a URL that initiates the authentication flow for a user's social media
// account. When visited, the user is redirected to the selected social platform's
// login/authorization page. Upon successful authentication, they are redirected
// back to your application
func (r *SocialAccountService) NewAuthURL(ctx context.Context, body SocialAccountNewAuthURLParams, opts ...option.RequestOption) (res *SocialAccountNewAuthURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/social-accounts/auth-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disconnecting an account with remove all auth tokens and mark the account as
// disconnected. The record of the account will be kept and can be retrieved and
// reconnected by the owner of the account.
func (r *SocialAccountService) Disconnect(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialAccountDisconnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-accounts/%s/disconnect", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SocialAccount struct {
	// The unique identifier of the social account
	ID string `json:"id,required"`
	// The access token of the social account
	AccessToken string `json:"access_token,required"`
	// The access token expiration date of the social account
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at,required" format:"date-time"`
	// The external id of the social account
	ExternalID string `json:"external_id,required"`
	// The metadata of the social account
	Metadata any `json:"metadata,required"`
	// The platform of the social account
	Platform string `json:"platform,required"`
	// The platform's profile photo of the social account
	ProfilePhotoURL string `json:"profile_photo_url,required"`
	// The refresh token of the social account
	RefreshToken string `json:"refresh_token,required"`
	// The refresh token expiration date of the social account
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at,required" format:"date-time"`
	// Status of the account
	//
	// Any of "connected", "disconnected".
	Status SocialAccountStatus `json:"status,required"`
	// The platform's id of the social account
	UserID string `json:"user_id,required"`
	// The platform's username of the social account
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessToken           respjson.Field
		AccessTokenExpiresAt  respjson.Field
		ExternalID            respjson.Field
		Metadata              respjson.Field
		Platform              respjson.Field
		ProfilePhotoURL       respjson.Field
		RefreshToken          respjson.Field
		RefreshTokenExpiresAt respjson.Field
		Status                respjson.Field
		UserID                respjson.Field
		Username              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccount) RawJSON() string { return r.JSON.raw }
func (r *SocialAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the account
type SocialAccountStatus string

const (
	SocialAccountStatusConnected    SocialAccountStatus = "connected"
	SocialAccountStatusDisconnected SocialAccountStatus = "disconnected"
)

type SocialAccountListResponse struct {
	Data []SocialAccount               `json:"data,required"`
	Meta SocialAccountListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountListResponseMeta struct {
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
func (r SocialAccountListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountNewAuthURLResponse struct {
	// The social account provider
	Platform string `json:"platform,required"`
	// The url to redirect the user to, in order to connect their account
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Platform    respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccountNewAuthURLResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountNewAuthURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountDisconnectResponse struct {
	// The unique identifier of the social account
	ID string `json:"id,required"`
	// The access token of the social account
	AccessToken string `json:"access_token,required"`
	// The access token expiration date of the social account
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at,required" format:"date-time"`
	// The external id of the social account
	ExternalID string `json:"external_id,required"`
	// The metadata of the social account
	Metadata any `json:"metadata,required"`
	// The platform of the social account
	Platform string `json:"platform,required"`
	// The platform's profile photo of the social account
	ProfilePhotoURL string `json:"profile_photo_url,required"`
	// The refresh token of the social account
	RefreshToken string `json:"refresh_token,required"`
	// The refresh token expiration date of the social account
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at,required" format:"date-time"`
	// Status of the account
	//
	// Any of "disconnected".
	Status SocialAccountDisconnectResponseStatus `json:"status,required"`
	// The platform's id of the social account
	UserID string `json:"user_id,required"`
	// The platform's username of the social account
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessToken           respjson.Field
		AccessTokenExpiresAt  respjson.Field
		ExternalID            respjson.Field
		Metadata              respjson.Field
		Platform              respjson.Field
		ProfilePhotoURL       respjson.Field
		RefreshToken          respjson.Field
		RefreshTokenExpiresAt respjson.Field
		Status                respjson.Field
		UserID                respjson.Field
		Username              respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccountDisconnectResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountDisconnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the account
type SocialAccountDisconnectResponseStatus string

const (
	SocialAccountDisconnectResponseStatusDisconnected SocialAccountDisconnectResponseStatus = "disconnected"
)

type SocialAccountNewParams struct {
	// The access token of the social account
	AccessToken string `json:"access_token,required"`
	// The access token expiration date of the social account
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at,required" format:"date-time"`
	// The platform of the social account
	//
	// Any of "facebook", "instagram", "x", "tiktok", "youtube", "pinterest",
	// "linkedin", "bluesky", "threads", "tiktok_business".
	Platform SocialAccountNewParamsPlatform `json:"platform,omitzero,required"`
	// The user id of the social account
	UserID string `json:"user_id,required"`
	// The external id of the social account
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// The refresh token of the social account
	RefreshToken param.Opt[string] `json:"refresh_token,omitzero"`
	// The refresh token expiration date of the social account
	RefreshTokenExpiresAt param.Opt[time.Time] `json:"refresh_token_expires_at,omitzero" format:"date-time"`
	// The platform's username of the social account
	Username param.Opt[string] `json:"username,omitzero"`
	// The metadata of the social account
	Metadata any `json:"metadata,omitzero"`
	paramObj
}

func (r SocialAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The platform of the social account
type SocialAccountNewParamsPlatform string

const (
	SocialAccountNewParamsPlatformFacebook       SocialAccountNewParamsPlatform = "facebook"
	SocialAccountNewParamsPlatformInstagram      SocialAccountNewParamsPlatform = "instagram"
	SocialAccountNewParamsPlatformX              SocialAccountNewParamsPlatform = "x"
	SocialAccountNewParamsPlatformTiktok         SocialAccountNewParamsPlatform = "tiktok"
	SocialAccountNewParamsPlatformYoutube        SocialAccountNewParamsPlatform = "youtube"
	SocialAccountNewParamsPlatformPinterest      SocialAccountNewParamsPlatform = "pinterest"
	SocialAccountNewParamsPlatformLinkedin       SocialAccountNewParamsPlatform = "linkedin"
	SocialAccountNewParamsPlatformBluesky        SocialAccountNewParamsPlatform = "bluesky"
	SocialAccountNewParamsPlatformThreads        SocialAccountNewParamsPlatform = "threads"
	SocialAccountNewParamsPlatformTiktokBusiness SocialAccountNewParamsPlatform = "tiktok_business"
)

type SocialAccountUpdateParams struct {
	// The platform's external id of the social account
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// The platform's username of the social account
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r SocialAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountListParams struct {
	// Number of items to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Filter by id(s). Multiple values imply OR logic (e.g.,
	// ?id=spc_xxxxxx&id=spc_yyyyyy).
	ID []string `query:"id,omitzero" json:"-"`
	// Filter by externalId(s). Multiple values imply OR logic (e.g.,
	// ?externalId=test&externalId=test2).
	ExternalID []string `query:"external_id,omitzero" json:"-"`
	// Filter by platform(s). Multiple values imply OR logic (e.g.,
	// ?platform=x&platform=facebook).
	Platform []string `query:"platform,omitzero" json:"-"`
	// Filter by username(s). Multiple values imply OR logic (e.g.,
	// ?username=test&username=test2).
	Username []string `query:"username,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SocialAccountListParams]'s query parameters as
// `url.Values`.
func (r SocialAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SocialAccountNewAuthURLParams struct {
	// The social account provider
	Platform string `json:"platform,required"`
	// Your unique identifier for the social account
	ExternalID param.Opt[string] `json:"external_id,omitzero"`
	// Override the default redirect URL for the OAuth flow. If provided, this URL will
	// be used instead of our redirect URL. Make sure this URL is included in your
	// app's authorized redirect urls. This override will not work when using our
	// system credientals.
	RedirectURLOverride param.Opt[string] `json:"redirect_url_override,omitzero"`
	// List of permissions you want to allow. Will default to only post permissions.
	// You must include the "feeds" permission to request an account feed and metrics
	//
	// Any of "posts", "feeds".
	Permissions []string `json:"permissions,omitzero"`
	// Additional data needed for the provider
	PlatformData SocialAccountNewAuthURLParamsPlatformData `json:"platform_data,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParams) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data needed for the provider
type SocialAccountNewAuthURLParamsPlatformData struct {
	// Additional data needed for connecting bluesky accounts
	Bluesky SocialAccountNewAuthURLParamsPlatformDataBluesky `json:"bluesky,omitzero"`
	// Additional data for connecting facebook accounts
	Facebook SocialAccountNewAuthURLParamsPlatformDataFacebook `json:"facebook,omitzero"`
	// Additional data for connecting instagram accounts
	Instagram SocialAccountNewAuthURLParamsPlatformDataInstagram `json:"instagram,omitzero"`
	// Additional data for connecting linkedin accounts
	Linkedin SocialAccountNewAuthURLParamsPlatformDataLinkedin `json:"linkedin,omitzero"`
	// Additional data for connecting Pinterest accounts
	Pinterest SocialAccountNewAuthURLParamsPlatformDataPinterest `json:"pinterest,omitzero"`
	// Additional data for connecting Threads accounts
	Threads SocialAccountNewAuthURLParamsPlatformDataThreads `json:"threads,omitzero"`
	// Additional data for connecting TikTok accounts
	Tiktok SocialAccountNewAuthURLParamsPlatformDataTiktok `json:"tiktok,omitzero"`
	// Additional data for connecting TikTok Business accounts
	TiktokBusiness SocialAccountNewAuthURLParamsPlatformDataTiktokBusiness `json:"tiktok_business,omitzero"`
	// Additional data for connecting YouTube accounts
	Youtube SocialAccountNewAuthURLParamsPlatformDataYoutube `json:"youtube,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformData) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data needed for connecting bluesky accounts
//
// The properties AppPassword, Handle are required.
type SocialAccountNewAuthURLParamsPlatformDataBluesky struct {
	// The app password of the account
	AppPassword string `json:"app_password,required"`
	// The handle of the account
	Handle string `json:"handle,required"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataBluesky) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataBluesky
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataBluesky) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting facebook accounts
type SocialAccountNewAuthURLParamsPlatformDataFacebook struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// public_profile, pages_show_list, pages_read_engagement, pages_manage_posts,
	// business_management
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataFacebook) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataFacebook
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataFacebook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting instagram accounts
//
// The property ConnectionType is required.
type SocialAccountNewAuthURLParamsPlatformDataInstagram struct {
	// The type of connection; instagram for using login with instagram, facebook for
	// using login with facebook.
	//
	// Any of "instagram", "facebook".
	ConnectionType string `json:"connection_type,omitzero,required"`
	// Override the default permissions/scopes requested during OAuth. Default
	// instagram scopes: instagram_business_basic, instagram_business_content_publish.
	// Default facebook scopes: instagram_basic, instagram_content_publish,
	// pages_show_list, public_profile, business_management
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataInstagram) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataInstagram
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataInstagram) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SocialAccountNewAuthURLParamsPlatformDataInstagram](
		"connection_type", "instagram", "facebook",
	)
}

// Additional data for connecting linkedin accounts
//
// The property ConnectionType is required.
type SocialAccountNewAuthURLParamsPlatformDataLinkedin struct {
	// The type of connection; If using our provided credentials always use
	// "organization". If using your own crednetials then only use "organization" if
	// you are using the Community API
	//
	// Any of "personal", "organization".
	ConnectionType string `json:"connection_type,omitzero,required"`
	// Override the default permissions/scopes requested during OAuth. Default personal
	// scopes: openid, w_member_social, profile, email. Default organization scopes:
	// r_basicprofile, w_member_social, r_organization_social, w_organization_social,
	// rw_organization_admin
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataLinkedin) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataLinkedin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataLinkedin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SocialAccountNewAuthURLParamsPlatformDataLinkedin](
		"connection_type", "personal", "organization",
	)
}

// Additional data for connecting Pinterest accounts
type SocialAccountNewAuthURLParamsPlatformDataPinterest struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// boards:read, boards:write, pins:read, pins:write, user_accounts:read
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataPinterest) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataPinterest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataPinterest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting Threads accounts
type SocialAccountNewAuthURLParamsPlatformDataThreads struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// threads_basic, threads_content_publish
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataThreads) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataThreads
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataThreads) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting TikTok accounts
type SocialAccountNewAuthURLParamsPlatformDataTiktok struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// user.info.basic, video.list, video.upload, video.publish
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataTiktok) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataTiktok
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataTiktok) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting TikTok Business accounts
type SocialAccountNewAuthURLParamsPlatformDataTiktokBusiness struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// user.info.basic, user.info.username, user.info.stats, user.info.profile,
	// user.account.type, user.insights, video.list, video.insights, comment.list,
	// comment.list.manage, video.publish, video.upload, biz.spark.auth,
	// discovery.search.words
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataTiktokBusiness) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataTiktokBusiness
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataTiktokBusiness) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional data for connecting YouTube accounts
type SocialAccountNewAuthURLParamsPlatformDataYoutube struct {
	// Override the default permissions/scopes requested during OAuth. Default scopes:
	// https://www.googleapis.com/auth/youtube.force-ssl,
	// https://www.googleapis.com/auth/youtube.upload,
	// https://www.googleapis.com/auth/youtube.readonly,
	// https://www.googleapis.com/auth/userinfo.profile
	PermissionOverrides [][]any `json:"permission_overrides,omitzero"`
	paramObj
}

func (r SocialAccountNewAuthURLParamsPlatformDataYoutube) MarshalJSON() (data []byte, err error) {
	type shadow SocialAccountNewAuthURLParamsPlatformDataYoutube
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SocialAccountNewAuthURLParamsPlatformDataYoutube) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
