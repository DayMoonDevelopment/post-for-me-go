// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DayMoonDevelopment/post-for-me-go"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/testutil"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
)

func TestSocialPostPreviewNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.SocialPostPreviews.New(context.TODO(), postforme.SocialPostPreviewNewParams{
		CreateSocialPostPreview: postforme.CreateSocialPostPreviewParam{
			Caption: "caption",
			PreviewSocialAccounts: []postforme.CreateSocialPostPreviewPreviewSocialAccountParam{{
				ID:       "id",
				Platform: "platform",
				Username: postforme.String("username"),
			}},
			AccountConfigurations: []postforme.AccountConfigurationParam{{
				Configuration: postforme.AccountConfigurationConfigurationParam{
					Localizations: map[string]postforme.AccountConfigurationConfigurationLocalizationParam{
						"foo": {
							Description: postforme.String("description"),
							Title:       postforme.String("title"),
						},
					},
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					AudioName:              postforme.String("audio_name"),
					AutoAddMusic:           postforme.Bool(true),
					BoardIDs:               []string{"string"},
					Caption:                map[string]any{},
					CategoryID:             postforme.String("category_id"),
					Collaborators:          [][]any{{map[string]any{}}},
					CommunityID:            postforme.String("community_id"),
					ContainsSyntheticMedia: postforme.Bool(true),
					DefaultLanguage:        postforme.String("default_language"),
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					Embeddable:             postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					License:                "youtube",
					Link:                   postforme.String("link"),
					Location:               postforme.String("location"),
					MadeForKids:            postforme.Bool(true),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Placement: "reels",
					Poll: postforme.TwitterPollParam{
						DurationMinutes: 0,
						Options:         []string{"string"},
						ReplySettings:   postforme.TwitterPollReplySettingsFollowing,
					},
					PrivacyStatus:          "public",
					PublicStatsViewable:    postforme.Bool(true),
					PublishAt:              postforme.String("publish_at"),
					QuoteTweetID:           postforme.String("quote_tweet_id"),
					RecordingDate:          postforme.String("recording_date"),
					ReplySettings:          "following",
					ResharePostID:          postforme.String("reshare_post_id"),
					SetCaptionForEachImage: postforme.Bool(true),
					ShareToFeed:            postforme.Bool(true),
					Tags:                   []string{"string"},
					Title:                  postforme.String("title"),
					TrialReelType:          "manual",
				},
				SocialAccountID: "social_account_id",
			}},
			Media: []postforme.SocialPostMediaParam{{
				URL:            "url",
				SkipProcessing: postforme.Bool(true),
				Tags: []postforme.SocialPostMediaTagParam{{
					ID:       "id",
					Platform: "facebook",
					Type:     "user",
					X:        postforme.Float(0),
					Y:        postforme.Float(0),
				}},
				ThumbnailTimestampMs: map[string]any{},
				ThumbnailURL:         map[string]any{},
			}},
			PlatformConfigurations: postforme.PlatformConfigurationsDtoParam{
				Bluesky: postforme.BlueskyConfigurationDtoParam{
					Caption: map[string]any{},
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
				},
				Facebook: postforme.FacebookConfigurationDtoParam{
					Caption:       map[string]any{},
					Collaborators: [][]any{{map[string]any{}}},
					Location:      postforme.String("location"),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Placement:              postforme.FacebookConfigurationDtoPlacementReels,
					SetCaptionForEachImage: postforme.Bool(true),
				},
				Instagram: postforme.InstagramConfigurationDtoParam{
					AudioName:     postforme.String("audio_name"),
					Caption:       map[string]any{},
					Collaborators: []string{"string"},
					Location:      postforme.String("location"),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Placement:     postforme.InstagramConfigurationDtoPlacementReels,
					ShareToFeed:   postforme.Bool(true),
					TrialReelType: postforme.InstagramConfigurationDtoTrialReelTypeManual,
				},
				Linkedin: postforme.LinkedinConfigurationDtoParam{
					Caption: map[string]any{},
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					ResharePostID: postforme.String("reshare_post_id"),
				},
				Pinterest: postforme.PinterestConfigurationDtoParam{
					BoardIDs: []string{"string"},
					Caption:  map[string]any{},
					Link:     postforme.String("link"),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Title: postforme.String("title"),
				},
				Threads: postforme.ThreadsConfigurationDtoParam{
					Caption: map[string]any{},
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Placement: postforme.ThreadsConfigurationDtoPlacementReels,
				},
				Tiktok: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					AutoAddMusic:           postforme.Bool(true),
					Caption:                map[string]any{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					PrivacyStatus: postforme.String("privacy_status"),
					Title:         postforme.String("title"),
				},
				TiktokBusiness: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					AutoAddMusic:           postforme.Bool(true),
					Caption:                map[string]any{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					PrivacyStatus: postforme.String("privacy_status"),
					Title:         postforme.String("title"),
				},
				X: postforme.TwitterConfigurationDtoParam{
					Caption:     map[string]any{},
					CommunityID: postforme.String("community_id"),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					Poll: postforme.TwitterPollParam{
						DurationMinutes: 0,
						Options:         []string{"string"},
						ReplySettings:   postforme.TwitterPollReplySettingsFollowing,
					},
					QuoteTweetID:  postforme.String("quote_tweet_id"),
					ReplySettings: postforme.TwitterConfigurationDtoReplySettingsFollowing,
				},
				Youtube: postforme.YoutubeConfigurationDtoParam{
					Localizations: map[string]postforme.YoutubeConfigurationDtoLocalizationParam{
						"foo": {
							Description: postforme.String("description"),
							Title:       postforme.String("title"),
						},
					},
					Caption:                map[string]any{},
					CategoryID:             postforme.String("category_id"),
					ContainsSyntheticMedia: postforme.Bool(true),
					DefaultLanguage:        postforme.String("default_language"),
					Description:            postforme.String("description"),
					Embeddable:             postforme.Bool(true),
					License:                postforme.YoutubeConfigurationDtoLicenseYoutube,
					MadeForKids:            postforme.Bool(true),
					Media: []postforme.SocialPostMediaParam{{
						URL:            "url",
						SkipProcessing: postforme.Bool(true),
						Tags: []postforme.SocialPostMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]any{},
						ThumbnailURL:         map[string]any{},
					}},
					PrivacyStatus:       postforme.YoutubeConfigurationDtoPrivacyStatusPublic,
					PublicStatsViewable: postforme.Bool(true),
					PublishAt:           postforme.String("publish_at"),
					RecordingDate:       postforme.String("recording_date"),
					Tags:                []string{"string"},
					Title:               postforme.String("title"),
				},
			},
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
