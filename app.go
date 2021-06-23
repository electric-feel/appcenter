package appcenter

import (
	"fmt"

	"github.com/bitrise-io/appcenter/client"
	"github.com/bitrise-io/appcenter/model"
)

// AppAPI ...
type AppAPI struct {
	API            client.API
	ReleaseOptions model.ReleaseOptions
}

// CreateApplicationAPI ...
func CreateApplicationAPI(api client.API, releaseOptions model.ReleaseOptions) AppAPI {
	return AppAPI{
		API:            api,
		ReleaseOptions: releaseOptions,
	}
}

// NewRelease ...
func (a AppAPI) NewRelease() (model.Release, error) {
	releaseID, err := a.API.CreateRelease(a.ReleaseOptions)
	if err != nil {
		return model.Release{},
			fmt.Errorf("failed to create new release on app: %s, owner: %s, %v",
				a.ReleaseOptions.App.AppName,
				a.ReleaseOptions.App.Owner,
				err)
	}

	return a.API.GetAppReleaseDetails(a.ReleaseOptions.App, releaseID)
}

// Groups ...
func (a AppAPI) Groups(name string) (model.Group, error) {
	return a.API.GetGroupByName(name, a.ReleaseOptions.App)
}

// All Groups...
func (a App) AllGroups() ([]Group, error) {
	var (
		getURL = fmt.Sprintf("%s/v0.1/apps/%s/%s/distribution_groups", baseURL, a.owner, a.name)
		getResponse []Group
	)

	statusCode, err := a.client.jsonRequest(http.MethodGet, getURL, nil, &getResponse)
	if err != nil {
		return []Group{}, err
	}

	if statusCode != http.StatusOK {
		return []Group{}, fmt.Errorf("invalid status code: %d, url: %s, body: %v", statusCode, getURL, getResponse)
	}

	return getResponse, nil
}

// Stores ...
func (a AppAPI) Stores(name string) (model.Store, error) {
	return a.API.GetStore(name, a.ReleaseOptions.App)
}
