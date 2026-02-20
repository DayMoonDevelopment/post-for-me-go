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

func TestSocialAccountFeedListWithOptionalParams(t *testing.T) {
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
	_, err := client.SocialAccountFeeds.List(
		context.TODO(),
		"social_account_id",
		postforme.SocialAccountFeedListParams{
			Cursor:         postforme.String("cursor"),
			Expand:         []string{"metrics"},
			ExternalPostID: []string{"string"},
			Limit:          postforme.Float(0),
			PlatformPostID: []string{"string"},
			SocialPostID:   []string{"string"},
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
