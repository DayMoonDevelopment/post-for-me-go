// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"net/http"
	"slices"

	"github.com/DayMoonDevelopment/post-for-me-go/internal/apijson"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/requestconfig"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
	"github.com/DayMoonDevelopment/post-for-me-go/packages/respjson"
)

// MediaService contains methods and other services that help with interacting with
// the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaService] method instead.
type MediaService struct {
	Options []option.RequestOption
}

// NewMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaService(opts ...option.RequestOption) (r MediaService) {
	r = MediaService{}
	r.Options = opts
	return
}

// To upload media to attach to your post, make a `POST` request to the
// `/media/create-upload-url` endpoint.
//
// You'll receive the public url of your media item (which can be used when making
// a post) and will include an `upload_url` which is a signed URL of the storage
// location for uploading your file to.
//
// This URL is unique and publicly signed for a short time, so make sure to upload
// your files in a timely manner.
//
// **Example flow using JavaScript and the Fetch API:**
//
// **Request an upload URL**
//
// ```js
// // Step 1: Request an upload URL from your API
// const response = await fetch(
//
//	"https://api.postforme.dev/v1/media/create-upload-url",
//	{
//	  method: "POST",
//	  headers: {
//	    "Content-Type": "application/json",
//	  },
//	}
//
// );
//
// const { media_url, upload_url } = await response.json();
// ```
//
// **Upload your file to the signed URL**
//
// ```js
// // Step 2: Upload your file to the signed URL
// const file = /* your File or Blob object, e.g., from an <input type="file"> */;
//
//	await fetch(upload_url, {
//	  method: 'PUT',
//	  headers: {
//	    'Content-Type': 'image/jpeg'
//	  },
//	  body: file
//	});
//
// ```
//
// **Use the `media_url` when creating your post**
//
//	```js
//	// Step 3: Use the `media_url` when creating your post
//	const response = await fetch('https://api.postforme.dev/v1/social-posts', {
//	  method: 'POST',
//	  headers: {
//	    'Content-Type': 'application/json'
//	  },
//	  body: JSON.stringify({
//	    social_accounts: ['spc_...', ...],
//	    caption: 'My caption',
//	    media: [
//	      {
//	        url: media_url
//	      }
//	    ]
//	  })
//	});
//	```
func (r *MediaService) NewUploadURL(ctx context.Context, opts ...option.RequestOption) (res *MediaNewUploadURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/media/create-upload-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type MediaNewUploadURLResponse struct {
	// The public URL for the media, to use once file has been uploaded
	MediaURL string `json:"media_url,required"`
	// The signed upload URL for the client to upload the file
	UploadURL string `json:"upload_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MediaURL    respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaNewUploadURLResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaNewUploadURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
