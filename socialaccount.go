// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
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
	opts = append(r.Options[:], opts...)
	path := "v1/social-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get social account by ID
func (r *SocialAccountService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialAccount, err error) {
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
	path := "v1/social-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Generates a URL that initiates the authentication flow for a user's social media
// account. When visited, the user is redirected to the selected social platform's
// login/authorization page. Upon successful authentication, they are redirected
// back to your application
func (r *SocialAccountService) NewAuthURL(ctx context.Context, body SocialAccountNewAuthURLParams, opts ...option.RequestOption) (res *SocialAccountNewAuthURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/social-accounts/auth-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disconnecting an account with remove all auth tokens and mark the account as
// disconnected. The record of the account will be kept and can be retrieved and
// reconnected by the owner of the account.
func (r *SocialAccountService) Disconnect(ctx context.Context, id string, opts ...option.RequestOption) (res *SocialAccountDisconnectResponse, err error) {
	opts = append(r.Options[:], opts...)
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
	// Additional data for connecting linkedin accounts
	Linkedin SocialAccountNewAuthURLParamsPlatformDataLinkedin `json:"linkedin,omitzero"`
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

// Additional data for connecting linkedin accounts
//
// The property ConnectionType is required.
type SocialAccountNewAuthURLParamsPlatformDataLinkedin struct {
	// The type of connection; personal for posting on behalf of the user only,
	// organization for posting on behalf of both an organization and the user
	//
	// Any of "personal", "organization".
	ConnectionType string `json:"connection_type,omitzero,required"`
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
