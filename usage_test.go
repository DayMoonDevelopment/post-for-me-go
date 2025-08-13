// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme_test

import (
	"context"
	"os"
	"testing"

	"github.com/DayMoonDevelopment/post-for-me-go"
	"github.com/DayMoonDevelopment/post-for-me-go/internal/testutil"
	"github.com/DayMoonDevelopment/post-for-me-go/option"
)

func TestUsage(t *testing.T) {
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
	socialPost, err := client.SocialPosts.New(context.TODO(), postforme.SocialPostNewParams{
		CreateSocialPost: postforme.CreateSocialPostParam{
			Caption:        "caption",
			SocialAccounts: []string{"string"},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", socialPost.ID)
}
