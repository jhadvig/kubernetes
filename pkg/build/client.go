package build

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/build/buildapi"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
)

// BuildInterface is the client interface for builds
type BuildInterface interface {
	ListBuilds() (buildapi.BuildList, error)
	UpdateBuild(buildapi.Build) (buildapi.Build, error)
}

// BuildClient is the builds client implementation
type BuildClient struct {
	*client.RESTClient
}

// NewBuildClient creates a build client
func NewBuildClient(host string, auth *client.AuthInfo) BuildInterface {
	return &BuildClient{client.NewRESTClient(host, auth, "/api/v1beta1/")}
}

// ListBuilds returns a list of builds.
func (c *BuildClient) ListBuilds() (result buildapi.BuildList, err error) {
	err = c.Get().Path("builds").Do().Into(&result)
	return
}

// UpdateBuild updates an existing build.
func (c *BuildClient) UpdateBuild(build buildapi.Build) (result buildapi.Build, err error) {
	err = c.Put().Path("builds").Path(build.ID).Body(build).Do().Into(&result)
	return
}
