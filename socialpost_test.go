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
					AutoAddMusic:           postforme.Bool(true),
					BoardIDs:               []string{"string"},
					Caption:                map[string]interface{}{},
					Collaborators:          [][]any{{map[string]interface{}{}}},
					CommunityID:            postforme.String("community_id"),
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					Link:                   postforme.String("link"),
					Location:               postforme.String("location"),
					Media:                  []string{"string"},
					Placement:              "reels",
					Poll: postforme.CreateSocialPostAccountConfigurationConfigurationPollParam{
						DurationMinutes: 0,
						Options:         []string{"string"},
						ReplySettings:   "following",
					},
					PrivacyStatus: postforme.String("privacy_status"),
					QuoteTweetID:  postforme.String("quote_tweet_id"),
					ReplySettings: "following",
					ShareToFeed:   postforme.Bool(true),
					Title:         postforme.String("title"),
				},
				SocialAccountID: "social_account_id",
			}},
			ExternalID: postforme.String("external_id"),
			IsDraft:    postforme.Bool(true),
			Media: []postforme.CreateSocialPostMediaParam{{
				URL: "url",
				Tags: []postforme.CreateSocialPostMediaTagParam{{
					ID:       "id",
					Platform: "facebook",
					Type:     "user",
					X:        postforme.Float(0),
					Y:        postforme.Float(0),
				}},
				ThumbnailTimestampMs: map[string]interface{}{},
				ThumbnailURL:         map[string]interface{}{},
			}},
			PlatformConfigurations: postforme.PlatformConfigurationsDtoParam{
				Bluesky: postforme.BlueskyConfigurationDtoParam{
					Caption: map[string]interface{}{},
					Media: []postforme.BlueskyConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.BlueskyConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
				},
				Facebook: postforme.FacebookConfigurationDtoParam{
					Caption:       map[string]interface{}{},
					Collaborators: [][]any{{map[string]interface{}{}}},
					Location:      postforme.String("location"),
					Media: []postforme.FacebookConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.FacebookConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					Placement: postforme.FacebookConfigurationDtoPlacementReels,
				},
				Instagram: postforme.InstagramConfigurationDtoParam{
					Caption:       map[string]interface{}{},
					Collaborators: []string{"string"},
					Location:      postforme.String("location"),
					Media: []postforme.InstagramConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.InstagramConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					Placement:   postforme.InstagramConfigurationDtoPlacementReels,
					ShareToFeed: postforme.Bool(true),
				},
				Linkedin: postforme.LinkedinConfigurationDtoParam{
					Caption: map[string]interface{}{},
					Media: []postforme.LinkedinConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.LinkedinConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
				},
				Pinterest: postforme.PinterestConfigurationDtoParam{
					BoardIDs: []string{"string"},
					Caption:  map[string]interface{}{},
					Link:     postforme.String("link"),
					Media: []postforme.PinterestConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.PinterestConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
				},
				Threads: postforme.ThreadsConfigurationDtoParam{
					Caption: map[string]interface{}{},
					Media: []postforme.ThreadsConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.ThreadsConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					Placement: postforme.ThreadsConfigurationDtoPlacementReels,
				},
				Tiktok: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					AutoAddMusic:           postforme.Bool(true),
					Caption:                map[string]interface{}{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					Media: []postforme.TiktokConfigurationMediaParam{{
						URL: "url",
						Tags: []postforme.TiktokConfigurationMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					PrivacyStatus: postforme.String("privacy_status"),
					Title:         postforme.String("title"),
				},
				TiktokBusiness: postforme.TiktokConfigurationParam{
					AllowComment:           postforme.Bool(true),
					AllowDuet:              postforme.Bool(true),
					AllowStitch:            postforme.Bool(true),
					AutoAddMusic:           postforme.Bool(true),
					Caption:                map[string]interface{}{},
					DiscloseBrandedContent: postforme.Bool(true),
					DiscloseYourBrand:      postforme.Bool(true),
					IsAIGenerated:          postforme.Bool(true),
					IsDraft:                postforme.Bool(true),
					Media: []postforme.TiktokConfigurationMediaParam{{
						URL: "url",
						Tags: []postforme.TiktokConfigurationMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					PrivacyStatus: postforme.String("privacy_status"),
					Title:         postforme.String("title"),
				},
				X: postforme.TwitterConfigurationDtoParam{
					Caption:     map[string]interface{}{},
					CommunityID: postforme.String("community_id"),
					Media: []postforme.TwitterConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.TwitterConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					Poll: postforme.TwitterConfigurationDtoPollParam{
						DurationMinutes: 0,
						Options:         []string{"string"},
						ReplySettings:   "following",
					},
					QuoteTweetID:  postforme.String("quote_tweet_id"),
					ReplySettings: postforme.TwitterConfigurationDtoReplySettingsFollowing,
				},
				Youtube: postforme.YoutubeConfigurationDtoParam{
					Caption: map[string]interface{}{},
					Media: []postforme.YoutubeConfigurationDtoMediaParam{{
						URL: "url",
						Tags: []postforme.YoutubeConfigurationDtoMediaTagParam{{
							ID:       "id",
							Platform: "facebook",
							Type:     "user",
							X:        postforme.Float(0),
							Y:        postforme.Float(0),
						}},
						ThumbnailTimestampMs: map[string]interface{}{},
						ThumbnailURL:         map[string]interface{}{},
					}},
					Title: postforme.String("title"),
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
						AutoAddMusic:           postforme.Bool(true),
						BoardIDs:               []string{"string"},
						Caption:                map[string]interface{}{},
						Collaborators:          [][]any{{map[string]interface{}{}}},
						CommunityID:            postforme.String("community_id"),
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						IsAIGenerated:          postforme.Bool(true),
						IsDraft:                postforme.Bool(true),
						Link:                   postforme.String("link"),
						Location:               postforme.String("location"),
						Media:                  []string{"string"},
						Placement:              "reels",
						Poll: postforme.CreateSocialPostAccountConfigurationConfigurationPollParam{
							DurationMinutes: 0,
							Options:         []string{"string"},
							ReplySettings:   "following",
						},
						PrivacyStatus: postforme.String("privacy_status"),
						QuoteTweetID:  postforme.String("quote_tweet_id"),
						ReplySettings: "following",
						ShareToFeed:   postforme.Bool(true),
						Title:         postforme.String("title"),
					},
					SocialAccountID: "social_account_id",
				}},
				ExternalID: postforme.String("external_id"),
				IsDraft:    postforme.Bool(true),
				Media: []postforme.CreateSocialPostMediaParam{{
					URL: "url",
					Tags: []postforme.CreateSocialPostMediaTagParam{{
						ID:       "id",
						Platform: "facebook",
						Type:     "user",
						X:        postforme.Float(0),
						Y:        postforme.Float(0),
					}},
					ThumbnailTimestampMs: map[string]interface{}{},
					ThumbnailURL:         map[string]interface{}{},
				}},
				PlatformConfigurations: postforme.PlatformConfigurationsDtoParam{
					Bluesky: postforme.BlueskyConfigurationDtoParam{
						Caption: map[string]interface{}{},
						Media: []postforme.BlueskyConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.BlueskyConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
					},
					Facebook: postforme.FacebookConfigurationDtoParam{
						Caption:       map[string]interface{}{},
						Collaborators: [][]any{{map[string]interface{}{}}},
						Location:      postforme.String("location"),
						Media: []postforme.FacebookConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.FacebookConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						Placement: postforme.FacebookConfigurationDtoPlacementReels,
					},
					Instagram: postforme.InstagramConfigurationDtoParam{
						Caption:       map[string]interface{}{},
						Collaborators: []string{"string"},
						Location:      postforme.String("location"),
						Media: []postforme.InstagramConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.InstagramConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						Placement:   postforme.InstagramConfigurationDtoPlacementReels,
						ShareToFeed: postforme.Bool(true),
					},
					Linkedin: postforme.LinkedinConfigurationDtoParam{
						Caption: map[string]interface{}{},
						Media: []postforme.LinkedinConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.LinkedinConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
					},
					Pinterest: postforme.PinterestConfigurationDtoParam{
						BoardIDs: []string{"string"},
						Caption:  map[string]interface{}{},
						Link:     postforme.String("link"),
						Media: []postforme.PinterestConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.PinterestConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
					},
					Threads: postforme.ThreadsConfigurationDtoParam{
						Caption: map[string]interface{}{},
						Media: []postforme.ThreadsConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.ThreadsConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						Placement: postforme.ThreadsConfigurationDtoPlacementReels,
					},
					Tiktok: postforme.TiktokConfigurationParam{
						AllowComment:           postforme.Bool(true),
						AllowDuet:              postforme.Bool(true),
						AllowStitch:            postforme.Bool(true),
						AutoAddMusic:           postforme.Bool(true),
						Caption:                map[string]interface{}{},
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						IsAIGenerated:          postforme.Bool(true),
						IsDraft:                postforme.Bool(true),
						Media: []postforme.TiktokConfigurationMediaParam{{
							URL: "url",
							Tags: []postforme.TiktokConfigurationMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						PrivacyStatus: postforme.String("privacy_status"),
						Title:         postforme.String("title"),
					},
					TiktokBusiness: postforme.TiktokConfigurationParam{
						AllowComment:           postforme.Bool(true),
						AllowDuet:              postforme.Bool(true),
						AllowStitch:            postforme.Bool(true),
						AutoAddMusic:           postforme.Bool(true),
						Caption:                map[string]interface{}{},
						DiscloseBrandedContent: postforme.Bool(true),
						DiscloseYourBrand:      postforme.Bool(true),
						IsAIGenerated:          postforme.Bool(true),
						IsDraft:                postforme.Bool(true),
						Media: []postforme.TiktokConfigurationMediaParam{{
							URL: "url",
							Tags: []postforme.TiktokConfigurationMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						PrivacyStatus: postforme.String("privacy_status"),
						Title:         postforme.String("title"),
					},
					X: postforme.TwitterConfigurationDtoParam{
						Caption:     map[string]interface{}{},
						CommunityID: postforme.String("community_id"),
						Media: []postforme.TwitterConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.TwitterConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						Poll: postforme.TwitterConfigurationDtoPollParam{
							DurationMinutes: 0,
							Options:         []string{"string"},
							ReplySettings:   "following",
						},
						QuoteTweetID:  postforme.String("quote_tweet_id"),
						ReplySettings: postforme.TwitterConfigurationDtoReplySettingsFollowing,
					},
					Youtube: postforme.YoutubeConfigurationDtoParam{
						Caption: map[string]interface{}{},
						Media: []postforme.YoutubeConfigurationDtoMediaParam{{
							URL: "url",
							Tags: []postforme.YoutubeConfigurationDtoMediaTagParam{{
								ID:       "id",
								Platform: "facebook",
								Type:     "user",
								X:        postforme.Float(0),
								Y:        postforme.Float(0),
							}},
							ThumbnailTimestampMs: map[string]interface{}{},
							ThumbnailURL:         map[string]interface{}{},
						}},
						Title: postforme.String("title"),
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
