// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/DayMoonDevelopment/post-for-me-go"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/testutil"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
)

func TestSocialPostNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := postforme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SocialPosts.New(context.TODO(), postforme.SocialPostNewParams{
		CreateSocialPost: postforme.CreateSocialPostParam{
			Caption:        "caption",
			SocialAccounts: []string{"string"},
			AccountConfigurations: []postforme.CreateSocialPostAccountConfigurationParam{{
				Configuration: postforme.CreateSocialPostAccountConfigurationConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					BoardIDs:               []string{"string"},
					Caption:                map[string]interface{}{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					Link:                   postforme.String("link"),
					Location:               "reels",
					Media:                  []string{"string"},
					PrivacyStatus:          postforme.String("privacy_status"),
					Title:                  postforme.String("title"),
				},
				SocialAccountID: "social_account_id",
			}},
			ExternalID: postforme.String("external_id"),
			IsDraft:    postforme.Bool(true),
			Media: []postforme.CreateSocialPostMediaParam{{
				URL:                  "url",
				ThumbnailTimestampMs: map[string]interface{}{},
				ThumbnailURL:         map[string]interface{}{},
			}},
			PlatformConfigurations: postforme.CreateSocialPostPlatformConfigurationsParam{
				Bluesky: postforme.CreateSocialPostPlatformConfigurationsBlueskyParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
				},
				Facebook: postforme.CreateSocialPostPlatformConfigurationsFacebookParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
				},
				Instagram: postforme.CreateSocialPostPlatformConfigurationsInstagramParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
				},
				Linkedin: postforme.CreateSocialPostPlatformConfigurationsLinkedinParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
				},
				Pinterest: postforme.CreateSocialPostPlatformConfigurationsPinterestParam{
					BoardIDs: []string{"string"},
					Caption:  map[string]interface{}{},
					Link:     postforme.String("link"),
					Media:    []string{"string"},
				},
				Threads: postforme.CreateSocialPostPlatformConfigurationsThreadsParam{
					Caption:  map[string]interface{}{},
					Location: "reels",
					Media:    []string{"string"},
				},
				Tiktok: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					Caption:                map[string]interface{}{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					Media:                  []string{"string"},
					PrivacyStatus:          postforme.String("privacy_status"),
					Title:                  postforme.String("title"),
				},
				TiktokBusiness: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					Caption:                map[string]interface{}{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					Media:                  []string{"string"},
					PrivacyStatus:          postforme.String("privacy_status"),
					Title:                  postforme.String("title"),
				},
				X: postforme.CreateSocialPostPlatformConfigurationsXParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
				},
				Youtube: postforme.CreateSocialPostPlatformConfigurationsYoutubeParam{
					Caption: map[string]interface{}{},
					Media:   []string{"string"},
					Title:   postforme.String("title"),
				},
			},
			ScheduledAt: postforme.Time(time.Now()),
		},
	})
	if err != nil {
		var apierr *postforme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSocialPostGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := postforme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SocialPosts.Get(context.TODO(), "id")
	if err != nil {
		var apierr *postforme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSocialPostUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := postforme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SocialPosts.Update(
		context.TODO(),
		"id",
		postforme.SocialPostUpdateParams{
			CreateSocialPost: postforme.CreateSocialPostParam{
				Caption:        "caption",
				SocialAccounts: []string{"string"},
				AccountConfigurations: []postforme.CreateSocialPostAccountConfigurationParam{{
					Configuration: postforme.CreateSocialPostAccountConfigurationConfigurationParam{
						AllowComment:           postforme.Bool(true),
						AllowDuet:              postforme.Bool(true),
						AllowStitch:            postforme.Bool(true),
						BoardIDs:               []string{"string"},
						Caption:                map[string]interface{}{},
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						Link:                   postforme.String("link"),
						Location:               "reels",
						Media:                  []string{"string"},
						PrivacyStatus:          postforme.String("privacy_status"),
						Title:                  postforme.String("title"),
					},
					SocialAccountID: "social_account_id",
				}},
				ExternalID: postforme.String("external_id"),
				IsDraft:    postforme.Bool(true),
				Media: []postforme.CreateSocialPostMediaParam{{
					URL:                  "url",
					ThumbnailTimestampMs: map[string]interface{}{},
					ThumbnailURL:         map[string]interface{}{},
				}},
				PlatformConfigurations: postforme.CreateSocialPostPlatformConfigurationsParam{
					Bluesky: postforme.CreateSocialPostPlatformConfigurationsBlueskyParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
					},
					Facebook: postforme.CreateSocialPostPlatformConfigurationsFacebookParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
					},
					Instagram: postforme.CreateSocialPostPlatformConfigurationsInstagramParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
					},
					Linkedin: postforme.CreateSocialPostPlatformConfigurationsLinkedinParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
					},
					Pinterest: postforme.CreateSocialPostPlatformConfigurationsPinterestParam{
						BoardIDs: []string{"string"},
						Caption:  map[string]interface{}{},
						Link:     postforme.String("link"),
						Media:    []string{"string"},
					},
					Threads: postforme.CreateSocialPostPlatformConfigurationsThreadsParam{
						Caption:  map[string]interface{}{},
						Location: "reels",
						Media:    []string{"string"},
					},
					Tiktok: postforme.TiktokConfigurationParam{
						AllowComment:           postforme.Bool(true),
						AllowDuet:              postforme.Bool(true),
						AllowStitch:            postforme.Bool(true),
						Caption:                map[string]interface{}{},
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						Media:                  []string{"string"},
						PrivacyStatus:          postforme.String("privacy_status"),
						Title:                  postforme.String("title"),
					},
					TiktokBusiness: postforme.TiktokConfigurationParam{
						AllowComment:           postforme.Bool(true),
						AllowDuet:              postforme.Bool(true),
						AllowStitch:            postforme.Bool(true),
						Caption:                map[string]interface{}{},
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						Media:                  []string{"string"},
						PrivacyStatus:          postforme.String("privacy_status"),
						Title:                  postforme.String("title"),
					},
					X: postforme.CreateSocialPostPlatformConfigurationsXParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
					},
					Youtube: postforme.CreateSocialPostPlatformConfigurationsYoutubeParam{
						Caption: map[string]interface{}{},
						Media:   []string{"string"},
						Title:   postforme.String("title"),
					},
				},
				ScheduledAt: postforme.Time(time.Now()),
			},
		},
	)
	if err != nil {
		var apierr *postforme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSocialPostListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := postforme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SocialPosts.List(context.TODO(), postforme.SocialPostListParams{
		ExternalID: []string{"string"},
		Limit:      postforme.Float(0),
		Offset:     postforme.Float(0),
		Platform:   []string{"bluesky"},
		Status:     []string{"draft"},
	})
	if err != nil {
		var apierr *postforme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSocialPostDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := postforme.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SocialPosts.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *postforme.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
