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

// Webhooks enable you to subscribe to certain events. This involves Post for Me
// making a POST request to the URL of any webhooks you create. Only the events you
// subscribe to will be sent to your webhook URL.
//
// ## Payload
//
// When an event happens that your webhook is subscribed to, we will make a POST
// request with the following JSON body
//
// ```
//
//	{
//	    "event_type": "",
//	    "data": {}
//	}
//
// ```
//
// The event_type will be the event that triggered the webhook POST, data will be
// the resulting entity from the event
//
// ## Security
//
// To verify the POST to your webhook URL is from us we will include a secret in
// the header "Post-For-Me-Webhook-Secret". When you create a webhook you will
// receive the secret in the response.
//
// ## Retries
//
// If your server fails to respond with a 2XX code, requests to it will be retried
// with exponential backoff around 8 times over the course of just over a day.
//
// WebhookService contains methods and other services that help with interacting
// with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

// Create Webhook
func (r *WebhookService) New(ctx context.Context, body WebhookNewParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get webhook by ID
func (r *WebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update Webhook
func (r *WebhookService) Update(ctx context.Context, id string, body WebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Get a paginated result for webhooks based on the applied filters
func (r *WebhookService) List(ctx context.Context, query WebhookListParams, opts ...option.RequestOption) (res *WebhookListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Webhook
func (r *WebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *DeleteEntityResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type Webhook struct {
	// The unique identifier of the webhook
	ID string `json:"id" api:"required"`
	// Events that will be sent to the webhook
	EventTypes []string `json:"event_types" api:"required"`
	// Secret key used to verify webhook post
	Secret string `json:"secret" api:"required"`
	// The public webhook url
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		EventTypes  respjson.Field
		Secret      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponse struct {
	Data []Webhook               `json:"data" api:"required"`
	Meta WebhookListResponseMeta `json:"meta" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListResponseMeta struct {
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
func (r WebhookListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *WebhookListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookNewParams struct {
	// List of events the webhook will recieve
	//
	// Any of "social.post.created", "social.post.updated", "social.post.deleted",
	// "social.post.result.created", "social.account.created",
	// "social.account.updated".
	EventTypes []string `json:"event_types,omitzero" api:"required"`
	// Public url to recieve event data
	URL string `json:"url" api:"required"`
	paramObj
}

func (r WebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookUpdateParams struct {
	// Public url to recieve event data
	URL param.Opt[string] `json:"url,omitzero"`
	// List of events the webhook will recieve
	//
	// Any of "social.post.created", "social.post.updated", "social.post.deleted",
	// "social.post.result.created", "social.account.created",
	// "social.account.updated".
	EventTypes []string `json:"event_types,omitzero"`
	paramObj
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookListParams struct {
	// Number of items to return
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip
	Offset param.Opt[float64] `query:"offset,omitzero" json:"-"`
	// Filter by id(s). Multiple values imply OR logic (e.g.,
	// ?id=wbh_xxxxxx&id=wbh_yyyyyy).
	ID []string `query:"id,omitzero" json:"-"`
	// Filter by event type(s). Multiple values imply OR logic (e.g.,
	// ?event_type=social.post.created&event_type=social.post.updated).
	EventType []string `query:"event_type,omitzero" json:"-"`
	// Filter by url(s). Multiple values imply OR logic (e.g.,
	// ?url=https://example.com&url=https://postforme.dev).
	URL []string `query:"url,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WebhookListParams]'s query parameters as `url.Values`.
func (r WebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
