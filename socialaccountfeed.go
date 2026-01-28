// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package postforme

import (
	"context"
	"encoding/json"
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

// SocialAccountFeedService contains methods and other services that help with
// interacting with the post-for-me API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSocialAccountFeedService] method instead.
type SocialAccountFeedService struct {
	Options []option.RequestOption
}

// NewSocialAccountFeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSocialAccountFeedService(opts ...option.RequestOption) (r SocialAccountFeedService) {
	r = SocialAccountFeedService{}
	r.Options = opts
	return
}

// Get a paginated result for the social account based on the applied filters
func (r *SocialAccountFeedService) List(ctx context.Context, socialAccountID string, query SocialAccountFeedListParams, opts ...option.RequestOption) (res *SocialAccountFeedListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if socialAccountID == "" {
		err = errors.New("missing required social_account_id parameter")
		return
	}
	path := fmt.Sprintf("v1/social-account-feeds/%s", socialAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PlatformPost struct {
	// Caption or text content of the post
	Caption string `json:"caption,required"`
	// Array of media items attached to the post
	Media [][]any `json:"media,required"`
	// Social media platform name
	Platform string `json:"platform,required"`
	// Platform-specific account ID
	PlatformAccountID string `json:"platform_account_id,required"`
	// Platform-specific post ID
	PlatformPostID string `json:"platform_post_id,required"`
	// URL to the post on the platform
	PlatformURL string `json:"platform_url,required"`
	// ID of the social account
	SocialAccountID string `json:"social_account_id,required"`
	// External account ID from the platform
	ExternalAccountID string `json:"external_account_id,nullable"`
	// External post ID from the platform
	ExternalPostID string `json:"external_post_id,nullable"`
	// Post metrics and analytics data
	Metrics PlatformPostMetricsUnion `json:"metrics"`
	// Date the post was published
	PostedAt time.Time `json:"posted_at" format:"date-time"`
	// ID of the social post
	SocialPostID string `json:"social_post_id,nullable"`
	// ID of the social post result
	SocialPostResultID string `json:"social_post_result_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Caption            respjson.Field
		Media              respjson.Field
		Platform           respjson.Field
		PlatformAccountID  respjson.Field
		PlatformPostID     respjson.Field
		PlatformURL        respjson.Field
		SocialAccountID    respjson.Field
		ExternalAccountID  respjson.Field
		ExternalPostID     respjson.Field
		Metrics            respjson.Field
		PostedAt           respjson.Field
		SocialPostID       respjson.Field
		SocialPostResultID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPost) RawJSON() string { return r.JSON.raw }
func (r *PlatformPost) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PlatformPostMetricsUnion contains all possible properties and values from
// [PlatformPostMetricsTikTokBusinessMetricsDto],
// [PlatformPostMetricsTikTokPostMetricsDto],
// [PlatformPostMetricsInstagramPostMetricsDto],
// [PlatformPostMetricsYouTubePostMetricsDto],
// [PlatformPostMetricsFacebookPostMetricsDto],
// [PlatformPostMetricsTwitterPostMetricsDto],
// [PlatformPostMetricsThreadsPostMetricsDto].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PlatformPostMetricsUnion struct {
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AddressClicks float64 `json:"address_clicks"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AppDownloadClicks float64 `json:"app_download_clicks"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AudienceCities []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCity `json:"audience_cities"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AudienceCountries []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCountry `json:"audience_countries"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AudienceGenders []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceGender `json:"audience_genders"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AudienceTypes []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceType `json:"audience_types"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	AverageTimeWatched float64 `json:"average_time_watched"`
	Comments           float64 `json:"comments"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	EmailClicks float64 `json:"email_clicks"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	EngagementLikes []PlatformPostMetricsTikTokBusinessMetricsDtoEngagementLike `json:"engagement_likes"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	Favorites float64 `json:"favorites"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	FullVideoWatchedRate float64 `json:"full_video_watched_rate"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	ImpressionSources []PlatformPostMetricsTikTokBusinessMetricsDtoImpressionSource `json:"impression_sources"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	LeadSubmissions float64 `json:"lead_submissions"`
	Likes           float64 `json:"likes"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	NewFollowers float64 `json:"new_followers"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	PhoneNumberClicks float64 `json:"phone_number_clicks"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	ProfileViews float64 `json:"profile_views"`
	Reach        float64 `json:"reach"`
	Shares       float64 `json:"shares"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	TotalTimeWatched float64 `json:"total_time_watched"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	VideoViewRetention []PlatformPostMetricsTikTokBusinessMetricsDtoVideoViewRetention `json:"video_view_retention"`
	VideoViews         float64                                                         `json:"video_views"`
	// This field is from variant [PlatformPostMetricsTikTokBusinessMetricsDto].
	WebsiteClicks float64 `json:"website_clicks"`
	// This field is from variant [PlatformPostMetricsTikTokPostMetricsDto].
	CommentCount float64 `json:"comment_count"`
	// This field is from variant [PlatformPostMetricsTikTokPostMetricsDto].
	LikeCount float64 `json:"like_count"`
	// This field is from variant [PlatformPostMetricsTikTokPostMetricsDto].
	ShareCount float64 `json:"share_count"`
	// This field is from variant [PlatformPostMetricsTikTokPostMetricsDto].
	ViewCount float64 `json:"view_count"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	Follows float64 `json:"follows"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	IgReelsAvgWatchTime float64 `json:"ig_reels_avg_watch_time"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	IgReelsVideoViewTotalTime float64 `json:"ig_reels_video_view_total_time"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	Navigation float64 `json:"navigation"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	ProfileActivity float64 `json:"profile_activity"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	ProfileVisits float64 `json:"profile_visits"`
	Replies       float64 `json:"replies"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	Saved float64 `json:"saved"`
	// This field is from variant [PlatformPostMetricsInstagramPostMetricsDto].
	TotalInteractions float64 `json:"total_interactions"`
	Views             float64 `json:"views"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	Dislikes float64 `json:"dislikes"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationClickableImpressions float64 `json:"annotationClickableImpressions"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationClicks float64 `json:"annotationClicks"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationClickThroughRate float64 `json:"annotationClickThroughRate"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationClosableImpressions float64 `json:"annotationClosableImpressions"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationCloseRate float64 `json:"annotationCloseRate"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationCloses float64 `json:"annotationCloses"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AnnotationImpressions float64 `json:"annotationImpressions"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AverageViewDuration float64 `json:"averageViewDuration"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	AverageViewPercentage float64 `json:"averageViewPercentage"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardClickRate float64 `json:"cardClickRate"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardClicks float64 `json:"cardClicks"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardImpressions float64 `json:"cardImpressions"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardTeaserClickRate float64 `json:"cardTeaserClickRate"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardTeaserClicks float64 `json:"cardTeaserClicks"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	CardTeaserImpressions float64 `json:"cardTeaserImpressions"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	EngagedViews float64 `json:"engagedViews"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	EstimatedMinutesWatched float64 `json:"estimatedMinutesWatched"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	EstimatedRedMinutesWatched float64 `json:"estimatedRedMinutesWatched"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	RedViews float64 `json:"redViews"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	SubscribersGained float64 `json:"subscribersGained"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	SubscribersLost float64 `json:"subscribersLost"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	VideosAddedToPlaylists float64 `json:"videosAddedToPlaylists"`
	// This field is from variant [PlatformPostMetricsYouTubePostMetricsDto].
	VideosRemovedFromPlaylists float64 `json:"videosRemovedFromPlaylists"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ActivityByActionType []PlatformPostMetricsFacebookPostMetricsDtoActivityByActionType `json:"activity_by_action_type"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ActivityByActionTypeUnique []PlatformPostMetricsFacebookPostMetricsDtoActivityByActionTypeUnique `json:"activity_by_action_type_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	FanReach float64 `json:"fan_reach"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	MediaViews float64 `json:"media_views"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	NonviralReach float64 `json:"nonviral_reach"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	OrganicReach float64 `json:"organic_reach"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	PaidReach float64 `json:"paid_reach"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsAnger float64 `json:"reactions_anger"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsByType any `json:"reactions_by_type"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsHaha float64 `json:"reactions_haha"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsLike float64 `json:"reactions_like"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsLove float64 `json:"reactions_love"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsSorry float64 `json:"reactions_sorry"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsTotal float64 `json:"reactions_total"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ReactionsWow float64 `json:"reactions_wow"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoAvgTimeWatched float64 `json:"video_avg_time_watched"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoCompleteViewsOrganic float64 `json:"video_complete_views_organic"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoCompleteViewsOrganicUnique float64 `json:"video_complete_views_organic_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoCompleteViewsPaid float64 `json:"video_complete_views_paid"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoCompleteViewsPaidUnique float64 `json:"video_complete_views_paid_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoLength float64 `json:"video_length"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoRetentionGraphAutoplayed []PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphAutoplayed `json:"video_retention_graph_autoplayed"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoRetentionGraphClickedToPlay []PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphClickedToPlay `json:"video_retention_graph_clicked_to_play"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoSocialActionsUnique float64 `json:"video_social_actions_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTime float64 `json:"video_view_time"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTimeByAgeGender []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByAgeGender `json:"video_view_time_by_age_gender"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTimeByCountry []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByCountry `json:"video_view_time_by_country"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTimeByDistributionType any `json:"video_view_time_by_distribution_type"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTimeByRegion []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByRegion `json:"video_view_time_by_region"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewTimeOrganic float64 `json:"video_view_time_organic"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViews15s float64 `json:"video_views_15s"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViews60s float64 `json:"video_views_60s"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsAutoplayed float64 `json:"video_views_autoplayed"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsByDistributionType any `json:"video_views_by_distribution_type"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsClickedToPlay float64 `json:"video_views_clicked_to_play"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsOrganic float64 `json:"video_views_organic"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsOrganicUnique float64 `json:"video_views_organic_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsPaid float64 `json:"video_views_paid"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsPaidUnique float64 `json:"video_views_paid_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsSoundOn float64 `json:"video_views_sound_on"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	VideoViewsUnique float64 `json:"video_views_unique"`
	// This field is from variant [PlatformPostMetricsFacebookPostMetricsDto].
	ViralReach float64 `json:"viral_reach"`
	// This field is from variant [PlatformPostMetricsTwitterPostMetricsDto].
	NonPublicMetrics PlatformPostMetricsTwitterPostMetricsDtoNonPublicMetrics `json:"non_public_metrics"`
	// This field is from variant [PlatformPostMetricsTwitterPostMetricsDto].
	OrganicMetrics PlatformPostMetricsTwitterPostMetricsDtoOrganicMetrics `json:"organic_metrics"`
	// This field is from variant [PlatformPostMetricsTwitterPostMetricsDto].
	PublicMetrics PlatformPostMetricsTwitterPostMetricsDtoPublicMetrics `json:"public_metrics"`
	// This field is from variant [PlatformPostMetricsThreadsPostMetricsDto].
	Quotes float64 `json:"quotes"`
	// This field is from variant [PlatformPostMetricsThreadsPostMetricsDto].
	Reposts float64 `json:"reposts"`
	JSON    struct {
		AddressClicks                    respjson.Field
		AppDownloadClicks                respjson.Field
		AudienceCities                   respjson.Field
		AudienceCountries                respjson.Field
		AudienceGenders                  respjson.Field
		AudienceTypes                    respjson.Field
		AverageTimeWatched               respjson.Field
		Comments                         respjson.Field
		EmailClicks                      respjson.Field
		EngagementLikes                  respjson.Field
		Favorites                        respjson.Field
		FullVideoWatchedRate             respjson.Field
		ImpressionSources                respjson.Field
		LeadSubmissions                  respjson.Field
		Likes                            respjson.Field
		NewFollowers                     respjson.Field
		PhoneNumberClicks                respjson.Field
		ProfileViews                     respjson.Field
		Reach                            respjson.Field
		Shares                           respjson.Field
		TotalTimeWatched                 respjson.Field
		VideoViewRetention               respjson.Field
		VideoViews                       respjson.Field
		WebsiteClicks                    respjson.Field
		CommentCount                     respjson.Field
		LikeCount                        respjson.Field
		ShareCount                       respjson.Field
		ViewCount                        respjson.Field
		Follows                          respjson.Field
		IgReelsAvgWatchTime              respjson.Field
		IgReelsVideoViewTotalTime        respjson.Field
		Navigation                       respjson.Field
		ProfileActivity                  respjson.Field
		ProfileVisits                    respjson.Field
		Replies                          respjson.Field
		Saved                            respjson.Field
		TotalInteractions                respjson.Field
		Views                            respjson.Field
		Dislikes                         respjson.Field
		AnnotationClickableImpressions   respjson.Field
		AnnotationClicks                 respjson.Field
		AnnotationClickThroughRate       respjson.Field
		AnnotationClosableImpressions    respjson.Field
		AnnotationCloseRate              respjson.Field
		AnnotationCloses                 respjson.Field
		AnnotationImpressions            respjson.Field
		AverageViewDuration              respjson.Field
		AverageViewPercentage            respjson.Field
		CardClickRate                    respjson.Field
		CardClicks                       respjson.Field
		CardImpressions                  respjson.Field
		CardTeaserClickRate              respjson.Field
		CardTeaserClicks                 respjson.Field
		CardTeaserImpressions            respjson.Field
		EngagedViews                     respjson.Field
		EstimatedMinutesWatched          respjson.Field
		EstimatedRedMinutesWatched       respjson.Field
		RedViews                         respjson.Field
		SubscribersGained                respjson.Field
		SubscribersLost                  respjson.Field
		VideosAddedToPlaylists           respjson.Field
		VideosRemovedFromPlaylists       respjson.Field
		ActivityByActionType             respjson.Field
		ActivityByActionTypeUnique       respjson.Field
		FanReach                         respjson.Field
		MediaViews                       respjson.Field
		NonviralReach                    respjson.Field
		OrganicReach                     respjson.Field
		PaidReach                        respjson.Field
		ReactionsAnger                   respjson.Field
		ReactionsByType                  respjson.Field
		ReactionsHaha                    respjson.Field
		ReactionsLike                    respjson.Field
		ReactionsLove                    respjson.Field
		ReactionsSorry                   respjson.Field
		ReactionsTotal                   respjson.Field
		ReactionsWow                     respjson.Field
		VideoAvgTimeWatched              respjson.Field
		VideoCompleteViewsOrganic        respjson.Field
		VideoCompleteViewsOrganicUnique  respjson.Field
		VideoCompleteViewsPaid           respjson.Field
		VideoCompleteViewsPaidUnique     respjson.Field
		VideoLength                      respjson.Field
		VideoRetentionGraphAutoplayed    respjson.Field
		VideoRetentionGraphClickedToPlay respjson.Field
		VideoSocialActionsUnique         respjson.Field
		VideoViewTime                    respjson.Field
		VideoViewTimeByAgeGender         respjson.Field
		VideoViewTimeByCountry           respjson.Field
		VideoViewTimeByDistributionType  respjson.Field
		VideoViewTimeByRegion            respjson.Field
		VideoViewTimeOrganic             respjson.Field
		VideoViews15s                    respjson.Field
		VideoViews60s                    respjson.Field
		VideoViewsAutoplayed             respjson.Field
		VideoViewsByDistributionType     respjson.Field
		VideoViewsClickedToPlay          respjson.Field
		VideoViewsOrganic                respjson.Field
		VideoViewsOrganicUnique          respjson.Field
		VideoViewsPaid                   respjson.Field
		VideoViewsPaidUnique             respjson.Field
		VideoViewsSoundOn                respjson.Field
		VideoViewsUnique                 respjson.Field
		ViralReach                       respjson.Field
		NonPublicMetrics                 respjson.Field
		OrganicMetrics                   respjson.Field
		PublicMetrics                    respjson.Field
		Quotes                           respjson.Field
		Reposts                          respjson.Field
		raw                              string
	} `json:"-"`
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsTikTokBusinessMetricsDto() (v PlatformPostMetricsTikTokBusinessMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsTikTokPostMetricsDto() (v PlatformPostMetricsTikTokPostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsInstagramPostMetricsDto() (v PlatformPostMetricsInstagramPostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsYouTubePostMetricsDto() (v PlatformPostMetricsYouTubePostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsFacebookPostMetricsDto() (v PlatformPostMetricsFacebookPostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsTwitterPostMetricsDto() (v PlatformPostMetricsTwitterPostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PlatformPostMetricsUnion) AsPlatformPostMetricsThreadsPostMetricsDto() (v PlatformPostMetricsThreadsPostMetricsDto) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PlatformPostMetricsUnion) RawJSON() string { return u.JSON.raw }

func (r *PlatformPostMetricsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDto struct {
	// Number of address clicks
	AddressClicks float64 `json:"address_clicks,required"`
	// Number of app download clicks
	AppDownloadClicks float64 `json:"app_download_clicks,required"`
	// Audience cities breakdown
	AudienceCities []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCity `json:"audience_cities,required"`
	// Audience countries breakdown
	AudienceCountries []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCountry `json:"audience_countries,required"`
	// Audience genders breakdown
	AudienceGenders []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceGender `json:"audience_genders,required"`
	// Audience types breakdown
	AudienceTypes []PlatformPostMetricsTikTokBusinessMetricsDtoAudienceType `json:"audience_types,required"`
	// Average time watched in seconds
	AverageTimeWatched float64 `json:"average_time_watched,required"`
	// Number of comments on the post
	Comments float64 `json:"comments,required"`
	// Number of email clicks
	EmailClicks float64 `json:"email_clicks,required"`
	// Engagement likes data by percentage and time
	EngagementLikes []PlatformPostMetricsTikTokBusinessMetricsDtoEngagementLike `json:"engagement_likes,required"`
	// Number of favorites on the post
	Favorites float64 `json:"favorites,required"`
	// Rate of full video watches as a percentage
	FullVideoWatchedRate float64 `json:"full_video_watched_rate,required"`
	// Impression sources breakdown
	ImpressionSources []PlatformPostMetricsTikTokBusinessMetricsDtoImpressionSource `json:"impression_sources,required"`
	// Number of lead submissions
	LeadSubmissions float64 `json:"lead_submissions,required"`
	// Number of likes on the post
	Likes float64 `json:"likes,required"`
	// Number of new followers gained from the post
	NewFollowers float64 `json:"new_followers,required"`
	// Number of phone number clicks
	PhoneNumberClicks float64 `json:"phone_number_clicks,required"`
	// Number of profile views generated
	ProfileViews float64 `json:"profile_views,required"`
	// Total reach of the post
	Reach float64 `json:"reach,required"`
	// Number of shares on the post
	Shares float64 `json:"shares,required"`
	// Total time watched in seconds
	TotalTimeWatched float64 `json:"total_time_watched,required"`
	// Video view retention data by percentage and time
	VideoViewRetention []PlatformPostMetricsTikTokBusinessMetricsDtoVideoViewRetention `json:"video_view_retention,required"`
	// Total number of video views
	VideoViews float64 `json:"video_views,required"`
	// Number of website clicks
	WebsiteClicks float64 `json:"website_clicks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressClicks        respjson.Field
		AppDownloadClicks    respjson.Field
		AudienceCities       respjson.Field
		AudienceCountries    respjson.Field
		AudienceGenders      respjson.Field
		AudienceTypes        respjson.Field
		AverageTimeWatched   respjson.Field
		Comments             respjson.Field
		EmailClicks          respjson.Field
		EngagementLikes      respjson.Field
		Favorites            respjson.Field
		FullVideoWatchedRate respjson.Field
		ImpressionSources    respjson.Field
		LeadSubmissions      respjson.Field
		Likes                respjson.Field
		NewFollowers         respjson.Field
		PhoneNumberClicks    respjson.Field
		ProfileViews         respjson.Field
		Reach                respjson.Field
		Shares               respjson.Field
		TotalTimeWatched     respjson.Field
		VideoViewRetention   respjson.Field
		VideoViews           respjson.Field
		WebsiteClicks        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTikTokBusinessMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCity struct {
	// City name
	CityName string `json:"city_name,required"`
	// Percentage of audience from this city
	Percentage float64 `json:"percentage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CityName    respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCity) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCountry struct {
	// Country name
	Country string `json:"country,required"`
	// Percentage of audience from this country
	Percentage float64 `json:"percentage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCountry) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoAudienceCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoAudienceGender struct {
	// Gender category
	Gender string `json:"gender,required"`
	// Percentage of audience of this gender
	Percentage float64 `json:"percentage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gender      respjson.Field
		Percentage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoAudienceGender) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoAudienceGender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoAudienceType struct {
	// Percentage of audience of this type
	Percentage float64 `json:"percentage,required"`
	// Type of audience
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoAudienceType) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoAudienceType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoEngagementLike struct {
	// Percentage value for the metric
	Percentage float64 `json:"percentage,required"`
	// Time in seconds for the metric
	Second string `json:"second,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Second      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoEngagementLike) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoEngagementLike) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoImpressionSource struct {
	// Name of the impression source
	ImpressionSource string `json:"impression_source,required"`
	// Percentage of impressions from this source
	Percentage float64 `json:"percentage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImpressionSource respjson.Field
		Percentage       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoImpressionSource) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoImpressionSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokBusinessMetricsDtoVideoViewRetention struct {
	// Percentage value for the metric
	Percentage float64 `json:"percentage,required"`
	// Time in seconds for the metric
	Second string `json:"second,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Percentage  respjson.Field
		Second      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokBusinessMetricsDtoVideoViewRetention) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsTikTokBusinessMetricsDtoVideoViewRetention) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTikTokPostMetricsDto struct {
	// Number of comments on the video
	CommentCount float64 `json:"comment_count,required"`
	// Number of likes on the video
	LikeCount float64 `json:"like_count,required"`
	// Number of shares of the video
	ShareCount float64 `json:"share_count,required"`
	// Number of views on the video
	ViewCount float64 `json:"view_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CommentCount respjson.Field
		LikeCount    respjson.Field
		ShareCount   respjson.Field
		ViewCount    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTikTokPostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTikTokPostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsInstagramPostMetricsDto struct {
	// Number of comments on the post
	Comments float64 `json:"comments"`
	// Number of new follows from this post
	Follows float64 `json:"follows"`
	// Average watch time for Reels (in milliseconds)
	IgReelsAvgWatchTime float64 `json:"ig_reels_avg_watch_time"`
	// Total watch time for Reels (in milliseconds)
	IgReelsVideoViewTotalTime float64 `json:"ig_reels_video_view_total_time"`
	// Number of likes on the post
	Likes float64 `json:"likes"`
	// Navigation actions taken on the media
	Navigation float64 `json:"navigation"`
	// Profile activity generated from this post
	ProfileActivity float64 `json:"profile_activity"`
	// Number of profile visits from this post
	ProfileVisits float64 `json:"profile_visits"`
	// Total number of unique accounts that have seen the media
	Reach float64 `json:"reach"`
	// Number of replies to the story (story media only)
	Replies float64 `json:"replies"`
	// Total number of unique accounts that have saved the media
	Saved float64 `json:"saved"`
	// Total number of shares of the media
	Shares float64 `json:"shares"`
	// Total interactions on the post
	TotalInteractions float64 `json:"total_interactions"`
	// Number of views on the post
	Views float64 `json:"views"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments                  respjson.Field
		Follows                   respjson.Field
		IgReelsAvgWatchTime       respjson.Field
		IgReelsVideoViewTotalTime respjson.Field
		Likes                     respjson.Field
		Navigation                respjson.Field
		ProfileActivity           respjson.Field
		ProfileVisits             respjson.Field
		Reach                     respjson.Field
		Replies                   respjson.Field
		Saved                     respjson.Field
		Shares                    respjson.Field
		TotalInteractions         respjson.Field
		Views                     respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsInstagramPostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsInstagramPostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsYouTubePostMetricsDto struct {
	// Number of comments on the video
	Comments float64 `json:"comments,required"`
	// Number of dislikes on the video
	Dislikes float64 `json:"dislikes,required"`
	// Number of likes on the video
	Likes float64 `json:"likes,required"`
	// Number of views on the video
	Views float64 `json:"views,required"`
	// Number of clickable annotation impressions
	AnnotationClickableImpressions float64 `json:"annotationClickableImpressions"`
	// Number of annotation clicks
	AnnotationClicks float64 `json:"annotationClicks"`
	// Annotation click-through rate
	AnnotationClickThroughRate float64 `json:"annotationClickThroughRate"`
	// Number of closable annotation impressions
	AnnotationClosableImpressions float64 `json:"annotationClosableImpressions"`
	// Annotation close rate
	AnnotationCloseRate float64 `json:"annotationCloseRate"`
	// Number of annotation closes
	AnnotationCloses float64 `json:"annotationCloses"`
	// Number of annotation impressions
	AnnotationImpressions float64 `json:"annotationImpressions"`
	// Average view duration in seconds
	AverageViewDuration float64 `json:"averageViewDuration"`
	// Average percentage of the video watched
	AverageViewPercentage float64 `json:"averageViewPercentage"`
	// Card click-through rate
	CardClickRate float64 `json:"cardClickRate"`
	// Number of card clicks
	CardClicks float64 `json:"cardClicks"`
	// Number of card impressions
	CardImpressions float64 `json:"cardImpressions"`
	// Card teaser click-through rate
	CardTeaserClickRate float64 `json:"cardTeaserClickRate"`
	// Number of card teaser clicks
	CardTeaserClicks float64 `json:"cardTeaserClicks"`
	// Number of card teaser impressions
	CardTeaserImpressions float64 `json:"cardTeaserImpressions"`
	// Number of engaged views
	EngagedViews float64 `json:"engagedViews"`
	// Estimated minutes watched
	EstimatedMinutesWatched float64 `json:"estimatedMinutesWatched"`
	// Estimated minutes watched by YouTube Premium (Red) members
	EstimatedRedMinutesWatched float64 `json:"estimatedRedMinutesWatched"`
	// Number of views from YouTube Premium (Red) members
	RedViews float64 `json:"redViews"`
	// Number of shares
	Shares float64 `json:"shares"`
	// Subscribers gained
	SubscribersGained float64 `json:"subscribersGained"`
	// Subscribers lost
	SubscribersLost float64 `json:"subscribersLost"`
	// Number of times the video was added to playlists
	VideosAddedToPlaylists float64 `json:"videosAddedToPlaylists"`
	// Number of times the video was removed from playlists
	VideosRemovedFromPlaylists float64 `json:"videosRemovedFromPlaylists"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments                       respjson.Field
		Dislikes                       respjson.Field
		Likes                          respjson.Field
		Views                          respjson.Field
		AnnotationClickableImpressions respjson.Field
		AnnotationClicks               respjson.Field
		AnnotationClickThroughRate     respjson.Field
		AnnotationClosableImpressions  respjson.Field
		AnnotationCloseRate            respjson.Field
		AnnotationCloses               respjson.Field
		AnnotationImpressions          respjson.Field
		AverageViewDuration            respjson.Field
		AverageViewPercentage          respjson.Field
		CardClickRate                  respjson.Field
		CardClicks                     respjson.Field
		CardImpressions                respjson.Field
		CardTeaserClickRate            respjson.Field
		CardTeaserClicks               respjson.Field
		CardTeaserImpressions          respjson.Field
		EngagedViews                   respjson.Field
		EstimatedMinutesWatched        respjson.Field
		EstimatedRedMinutesWatched     respjson.Field
		RedViews                       respjson.Field
		Shares                         respjson.Field
		SubscribersGained              respjson.Field
		SubscribersLost                respjson.Field
		VideosAddedToPlaylists         respjson.Field
		VideosRemovedFromPlaylists     respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsYouTubePostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsYouTubePostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDto struct {
	// Total activity breakdown by action type
	ActivityByActionType []PlatformPostMetricsFacebookPostMetricsDtoActivityByActionType `json:"activity_by_action_type"`
	// Unique users activity breakdown by action type
	ActivityByActionTypeUnique []PlatformPostMetricsFacebookPostMetricsDtoActivityByActionTypeUnique `json:"activity_by_action_type_unique"`
	// Number of comments (from post object)
	Comments float64 `json:"comments"`
	// Number of fans who saw the post
	FanReach float64 `json:"fan_reach"`
	// Number of times the photo or video was viewed
	MediaViews float64 `json:"media_views"`
	// Number of people who saw the post via non-viral distribution
	NonviralReach float64 `json:"nonviral_reach"`
	// Number of people who saw the post via organic distribution
	OrganicReach float64 `json:"organic_reach"`
	// Number of people who saw the post via paid distribution
	PaidReach float64 `json:"paid_reach"`
	// Total number of unique people who saw the post
	Reach float64 `json:"reach"`
	// Number of anger reactions
	ReactionsAnger float64 `json:"reactions_anger"`
	// Breakdown of all reaction types
	ReactionsByType any `json:"reactions_by_type"`
	// Number of haha reactions
	ReactionsHaha float64 `json:"reactions_haha"`
	// Number of like reactions
	ReactionsLike float64 `json:"reactions_like"`
	// Number of love reactions
	ReactionsLove float64 `json:"reactions_love"`
	// Number of sad reactions
	ReactionsSorry float64 `json:"reactions_sorry"`
	// Total number of reactions (all types)
	ReactionsTotal float64 `json:"reactions_total"`
	// Number of wow reactions
	ReactionsWow float64 `json:"reactions_wow"`
	// Number of shares (from post object)
	Shares float64 `json:"shares"`
	// Average time video was viewed in milliseconds
	VideoAvgTimeWatched float64 `json:"video_avg_time_watched"`
	// Number of times video was viewed to 95% organically
	VideoCompleteViewsOrganic float64 `json:"video_complete_views_organic"`
	// Number of unique people who viewed video to 95% organically
	VideoCompleteViewsOrganicUnique float64 `json:"video_complete_views_organic_unique"`
	// Number of times video was viewed to 95% via paid distribution
	VideoCompleteViewsPaid float64 `json:"video_complete_views_paid"`
	// Number of unique people who viewed video to 95% via paid distribution
	VideoCompleteViewsPaidUnique float64 `json:"video_complete_views_paid_unique"`
	// Length of the video in milliseconds
	VideoLength float64 `json:"video_length"`
	// Video retention graph for autoplayed views
	VideoRetentionGraphAutoplayed []PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphAutoplayed `json:"video_retention_graph_autoplayed"`
	// Video retention graph for clicked-to-play views
	VideoRetentionGraphClickedToPlay []PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphClickedToPlay `json:"video_retention_graph_clicked_to_play"`
	// Number of unique people who performed social actions on the video
	VideoSocialActionsUnique float64 `json:"video_social_actions_unique"`
	// Total time video was viewed in milliseconds
	VideoViewTime float64 `json:"video_view_time"`
	// Video view time breakdown by age and gender
	VideoViewTimeByAgeGender []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByAgeGender `json:"video_view_time_by_age_gender"`
	// Video view time breakdown by country
	VideoViewTimeByCountry []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByCountry `json:"video_view_time_by_country"`
	// Video view time breakdown by distribution type
	VideoViewTimeByDistributionType any `json:"video_view_time_by_distribution_type"`
	// Video view time breakdown by region
	VideoViewTimeByRegion []PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByRegion `json:"video_view_time_by_region"`
	// Total time video was viewed in milliseconds via organic distribution
	VideoViewTimeOrganic float64 `json:"video_view_time_organic"`
	// Number of times video was viewed for 3+ seconds
	VideoViews float64 `json:"video_views"`
	// Number of times video was viewed for 15+ seconds
	VideoViews15s float64 `json:"video_views_15s"`
	// Number of times video was viewed for 60+ seconds (excludes videos shorter than
	// 60s)
	VideoViews60s float64 `json:"video_views_60s"`
	// Number of times video was autoplayed for 3+ seconds
	VideoViewsAutoplayed float64 `json:"video_views_autoplayed"`
	// Video views breakdown by distribution type
	VideoViewsByDistributionType any `json:"video_views_by_distribution_type"`
	// Number of times video was clicked to play for 3+ seconds
	VideoViewsClickedToPlay float64 `json:"video_views_clicked_to_play"`
	// Number of times video was viewed for 3+ seconds organically
	VideoViewsOrganic float64 `json:"video_views_organic"`
	// Number of unique people who viewed the video for 3+ seconds organically
	VideoViewsOrganicUnique float64 `json:"video_views_organic_unique"`
	// Number of times video was viewed for 3+ seconds via paid distribution
	VideoViewsPaid float64 `json:"video_views_paid"`
	// Number of unique people who viewed the video for 3+ seconds via paid
	// distribution
	VideoViewsPaidUnique float64 `json:"video_views_paid_unique"`
	// Number of times video was viewed with sound on
	VideoViewsSoundOn float64 `json:"video_views_sound_on"`
	// Number of unique people who viewed the video for 3+ seconds
	VideoViewsUnique float64 `json:"video_views_unique"`
	// Number of people who saw the post in News Feed via viral reach
	ViralReach float64 `json:"viral_reach"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActivityByActionType             respjson.Field
		ActivityByActionTypeUnique       respjson.Field
		Comments                         respjson.Field
		FanReach                         respjson.Field
		MediaViews                       respjson.Field
		NonviralReach                    respjson.Field
		OrganicReach                     respjson.Field
		PaidReach                        respjson.Field
		Reach                            respjson.Field
		ReactionsAnger                   respjson.Field
		ReactionsByType                  respjson.Field
		ReactionsHaha                    respjson.Field
		ReactionsLike                    respjson.Field
		ReactionsLove                    respjson.Field
		ReactionsSorry                   respjson.Field
		ReactionsTotal                   respjson.Field
		ReactionsWow                     respjson.Field
		Shares                           respjson.Field
		VideoAvgTimeWatched              respjson.Field
		VideoCompleteViewsOrganic        respjson.Field
		VideoCompleteViewsOrganicUnique  respjson.Field
		VideoCompleteViewsPaid           respjson.Field
		VideoCompleteViewsPaidUnique     respjson.Field
		VideoLength                      respjson.Field
		VideoRetentionGraphAutoplayed    respjson.Field
		VideoRetentionGraphClickedToPlay respjson.Field
		VideoSocialActionsUnique         respjson.Field
		VideoViewTime                    respjson.Field
		VideoViewTimeByAgeGender         respjson.Field
		VideoViewTimeByCountry           respjson.Field
		VideoViewTimeByDistributionType  respjson.Field
		VideoViewTimeByRegion            respjson.Field
		VideoViewTimeOrganic             respjson.Field
		VideoViews                       respjson.Field
		VideoViews15s                    respjson.Field
		VideoViews60s                    respjson.Field
		VideoViewsAutoplayed             respjson.Field
		VideoViewsByDistributionType     respjson.Field
		VideoViewsClickedToPlay          respjson.Field
		VideoViewsOrganic                respjson.Field
		VideoViewsOrganicUnique          respjson.Field
		VideoViewsPaid                   respjson.Field
		VideoViewsPaidUnique             respjson.Field
		VideoViewsSoundOn                respjson.Field
		VideoViewsUnique                 respjson.Field
		ViralReach                       respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsFacebookPostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoActivityByActionType struct {
	// Action type (e.g., like, comment, share)
	ActionType string `json:"action_type,required"`
	// Number of actions
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoActivityByActionType) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoActivityByActionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoActivityByActionTypeUnique struct {
	// Action type (e.g., like, comment, share)
	ActionType string `json:"action_type,required"`
	// Number of actions
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActionType  respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoActivityByActionTypeUnique) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoActivityByActionTypeUnique) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphAutoplayed struct {
	// Percentage of viewers at this time
	Rate float64 `json:"rate,required"`
	// Time in seconds
	Time float64 `json:"time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rate        respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphAutoplayed) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphAutoplayed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphClickedToPlay struct {
	// Percentage of viewers at this time
	Rate float64 `json:"rate,required"`
	// Time in seconds
	Time float64 `json:"time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rate        respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphClickedToPlay) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoVideoRetentionGraphClickedToPlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByAgeGender struct {
	// Demographic key (e.g., age_gender, region, country)
	Key string `json:"key,required"`
	// Total view time in milliseconds
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByAgeGender) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByAgeGender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByCountry struct {
	// Demographic key (e.g., age_gender, region, country)
	Key string `json:"key,required"`
	// Total view time in milliseconds
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByCountry) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByRegion struct {
	// Demographic key (e.g., age_gender, region, country)
	Key string `json:"key,required"`
	// Total view time in milliseconds
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByRegion) RawJSON() string {
	return r.JSON.raw
}
func (r *PlatformPostMetricsFacebookPostMetricsDtoVideoViewTimeByRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsTwitterPostMetricsDto struct {
	// Non-public metrics for the Tweet (available to the Tweet owner or advertisers)
	NonPublicMetrics PlatformPostMetricsTwitterPostMetricsDtoNonPublicMetrics `json:"non_public_metrics"`
	// Organic metrics for the Tweet (available to the Tweet owner)
	OrganicMetrics PlatformPostMetricsTwitterPostMetricsDtoOrganicMetrics `json:"organic_metrics"`
	// Publicly available metrics for the Tweet
	PublicMetrics PlatformPostMetricsTwitterPostMetricsDtoPublicMetrics `json:"public_metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NonPublicMetrics respjson.Field
		OrganicMetrics   respjson.Field
		PublicMetrics    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTwitterPostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTwitterPostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Non-public metrics for the Tweet (available to the Tweet owner or advertisers)
type PlatformPostMetricsTwitterPostMetricsDtoNonPublicMetrics struct {
	// Number of times this Tweet has been viewed via promoted distribution
	ImpressionCount float64 `json:"impression_count,required"`
	// Number of clicks on links in this Tweet via promoted distribution
	URLLinkClicks float64 `json:"url_link_clicks,required"`
	// Number of clicks on the author's profile via promoted distribution
	UserProfileClicks float64 `json:"user_profile_clicks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImpressionCount   respjson.Field
		URLLinkClicks     respjson.Field
		UserProfileClicks respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTwitterPostMetricsDtoNonPublicMetrics) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTwitterPostMetricsDtoNonPublicMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Organic metrics for the Tweet (available to the Tweet owner)
type PlatformPostMetricsTwitterPostMetricsDtoOrganicMetrics struct {
	// Number of times this Tweet has been viewed organically
	ImpressionCount float64 `json:"impression_count,required"`
	// Number of Likes of this Tweet from organic distribution
	LikeCount float64 `json:"like_count,required"`
	// Number of Replies of this Tweet from organic distribution
	ReplyCount float64 `json:"reply_count,required"`
	// Number of Retweets of this Tweet from organic distribution
	RetweetCount float64 `json:"retweet_count,required"`
	// Number of clicks on links in this Tweet from organic distribution
	URLLinkClicks float64 `json:"url_link_clicks,required"`
	// Number of clicks on the author's profile from organic distribution
	UserProfileClicks float64 `json:"user_profile_clicks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImpressionCount   respjson.Field
		LikeCount         respjson.Field
		ReplyCount        respjson.Field
		RetweetCount      respjson.Field
		URLLinkClicks     respjson.Field
		UserProfileClicks respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTwitterPostMetricsDtoOrganicMetrics) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTwitterPostMetricsDtoOrganicMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Publicly available metrics for the Tweet
type PlatformPostMetricsTwitterPostMetricsDtoPublicMetrics struct {
	// Number of times this Tweet has been bookmarked
	BookmarkCount float64 `json:"bookmark_count,required"`
	// Number of times this Tweet has been viewed
	ImpressionCount float64 `json:"impression_count,required"`
	// Number of Likes of this Tweet
	LikeCount float64 `json:"like_count,required"`
	// Number of Quotes of this Tweet
	QuoteCount float64 `json:"quote_count,required"`
	// Number of Replies of this Tweet
	ReplyCount float64 `json:"reply_count,required"`
	// Number of Retweets of this Tweet
	RetweetCount float64 `json:"retweet_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BookmarkCount   respjson.Field
		ImpressionCount respjson.Field
		LikeCount       respjson.Field
		QuoteCount      respjson.Field
		ReplyCount      respjson.Field
		RetweetCount    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsTwitterPostMetricsDtoPublicMetrics) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsTwitterPostMetricsDtoPublicMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PlatformPostMetricsThreadsPostMetricsDto struct {
	// Number of likes on the post
	Likes float64 `json:"likes,required"`
	// Number of quotes of the post
	Quotes float64 `json:"quotes,required"`
	// Number of replies on the post
	Replies float64 `json:"replies,required"`
	// Number of reposts of the post
	Reposts float64 `json:"reposts,required"`
	// Number of shares of the post
	Shares float64 `json:"shares,required"`
	// Number of views on the post
	Views float64 `json:"views,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Likes       respjson.Field
		Quotes      respjson.Field
		Replies     respjson.Field
		Reposts     respjson.Field
		Shares      respjson.Field
		Views       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlatformPostMetricsThreadsPostMetricsDto) RawJSON() string { return r.JSON.raw }
func (r *PlatformPostMetricsThreadsPostMetricsDto) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountFeedListResponse struct {
	Data []PlatformPost                    `json:"data,required"`
	Meta SocialAccountFeedListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccountFeedListResponse) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountFeedListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountFeedListResponseMeta struct {
	// Id representing the next page of items
	Cursor string `json:"cursor,required"`
	// Maximum number of items returned.
	Limit float64 `json:"limit,required"`
	// URL to the next page of results, or null if none.
	Next string `json:"next,required"`
	// Indicates if there are more results or not
	HasMore bool `json:"has_more"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cursor      respjson.Field
		Limit       respjson.Field
		Next        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SocialAccountFeedListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *SocialAccountFeedListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SocialAccountFeedListParams struct {
	// Cursor identifying next page of results
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items to return; Note: some platforms will have different max limits,
	// in the case the provided limit is over the platform's limit we will return the
	// max allowed by the platform.
	Limit param.Opt[float64] `query:"limit,omitzero" json:"-"`
	// Expand additional data in the response. Currently supports: "metrics" to include
	// post analytics data.
	//
	// Any of "metrics".
	Expand []string `query:"expand,omitzero" json:"-"`
	// Filter by Post for Me Social Postexternal ID. Multiple values imply OR logic
	// (e.g., ?external_post_id=xxxxxx&external_post_id=yyyyyy).
	ExternalPostID []string `query:"external_post_id,omitzero" json:"-"`
	// Filter by the platform's id(s). Multiple values imply OR logic (e.g.,
	// ?social_post_id=spr_xxxxxx&social_post_id=spr_yyyyyy).
	PlatformPostID []string `query:"platform_post_id,omitzero" json:"-"`
	// Filter by Post for Me Social Post id(s). Multiple values imply OR logic (e.g.,
	// ?social_post_id=sp_xxxxxx&social_post_id=sp_yyyyyy).
	SocialPostID []string `query:"social_post_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SocialAccountFeedListParams]'s query parameters as
// `url.Values`.
func (r SocialAccountFeedListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
