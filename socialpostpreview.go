// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"net/http"
	"slices"

	"github.com/DayMoonDevelopment/post-for-me-go/internal/apijson"
	shimjson "github.com/DayMoonDevelopment/post-for-me-go/internal/encoding/json"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/requestconfig"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/param"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/respjson"
)

// Social Post Previews allow you to see what a Social Post will create for each
// account in the post.
//
// SocialPostPreviewService contains methods and other services that help with
// interacting with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSocialPostPreviewService] method instead.
type SocialPostPreviewService struct {
	Options []option.RequestOption
}

// NewSocialPostPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSocialPostPreviewService(opts ...option.RequestOption) (r SocialPostPreviewService) {
	r = SocialPostPreviewService{}
	r.Options = opts
	return
}

// Create Post Previews
func (r *SocialPostPreviewService) New(ctx context.Context, body SocialPostPreviewNewParams, opts ...option.RequestOption) (res *[]SocialPostPreview, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/social-post-previews"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// The properties Caption, PreviewSocialAccounts are required.
type CreateSocialPostPreviewParam struct {
	// Caption text for the post
	Caption string `json:"caption" api:"required"`
	// Array of social accounts. Can preview non connected accounts, just specify a
	// random ID
	PreviewSocialAccounts []CreateSocialPostPreviewPreviewSocialAccountParam `json:"preview_social_accounts,omitzero" api:"required"`
	// Account-specific configurations for the post
	AccountConfigurations []AccountConfigurationParam `json:"account_configurations,omitzero"`
	// Array of media URLs associated with the post
	Media []SocialPostMediaParam `json:"media,omitzero"`
	// Platform-specific configurations for the post
	PlatformConfigurations PlatformConfigurationsDtoParam `json:"platform_configurations,omitzero"`
	paramObj
}

func (r CreateSocialPostPreviewParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPreviewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPreviewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Platform are required.
type CreateSocialPostPreviewPreviewSocialAccountParam struct {
	// ID of the social account, ex: spc_12312
	ID string `json:"id" api:"required"`
	// Platform of the social account
	Platform string `json:"platform" api:"required"`
	// username of the social account
	Username param.Opt[string] `json:"username,omitzero"`
	paramObj
}

func (r CreateSocialPostPreviewPreviewSocialAccountParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateSocialPostPreviewPreviewSocialAccountParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateSocialPostPreviewPreviewSocialAccountParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostPreview struct {
	// Caption text for the post
	Caption string `json:"caption" api:"required"`
	// Platform of the post
	Platform string `json:"platform" api:"required"`
	// Id of the social account
	SocialAccountID string `json:"social_account_id" api:"required"`
	// Additional configuration for this platform
	Configuration any `json:"configuration"`
	// Array of media URLs associated with the post
	Media []SocialPostMedia `json:"media" api:"nullable"`
	// Url of the social account profile picture
	SocialAccountProfilePictureURL any `json:"social_account_profile_picture_url"`
	// Username of the social account
	SocialAccountUsername any `json:"social_account_username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption                        respjson.Field
		Platform                       respjson.Field
		SocialAccountID                respjson.Field
		Configuration                  respjson.Field
		Media                          respjson.Field
		SocialAccountProfilePictureURL respjson.Field
		SocialAccountUsername          respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialPostPreview) RawJSON() string { return r.JSON.raw }
func (r *SocialPostPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialPostPreviewNewParams struct {
	CreateSocialPostPreview CreateSocialPostPreviewParam
	paramObj
}

func (r SocialPostPreviewNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateSocialPostPreview)
}
func (r *SocialPostPreviewNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
