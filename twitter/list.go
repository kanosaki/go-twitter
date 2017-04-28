package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ListService provides methods for accessing Twitter list API endpoints.
type ListService struct {
	sling *sling.Sling
}

// newListService returns a new ListService.
func newListService(sling *sling.Sling) *ListService {
	return &ListService{
		sling: sling.Path("lists/"),
	}
}

// ListStatusesParams are the parameters for ListService.Statuses
type ListStatusesParams struct {
	Slug            string `url:"slug"`
	ListID          string `url:"list_id,omittmpty"`
	OwnerScreenName string `url:"owner_screen_name,omittmpty"`
	SinceID         int64 `url:"since_id,omitempty"`
	MaxID           int64 `url:"max_id,omitempty"`
	Count           int `url:"count,omitempty"`
	IncludeEntities bool `url:"include_entities,omitempty"`
	IncludeRTs      bool `url:"include_rts,omitempty"`
}

// Tweets returns a collection of Tweets matching a search query.
// https://dev.twitter.com/rest/reference/get/search/tweets
func (s *ListService) Statuses(params *ListStatusesParams) ([]Tweet, *http.Response, error) {
	var tweets []Tweet
	apiError := new(APIError)
	resp, err := s.sling.New().Get("statuses.json").QueryStruct(params).Receive(&tweets, apiError)
	return tweets, resp, relevantError(err, *apiError)
}
