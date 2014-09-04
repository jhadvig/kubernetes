/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package buildlog

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"io/ioutil"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/apiserver"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/build"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/registry/pod"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/buildlog/buildlogapi"
)

type BuildLogRegistryStorage struct {
	BuildRegistry build.BuildRegistry
	PodRegistry   pod.Registry
}

func NewBuildLogRegistryStorage(b build.BuildRegistry, p pod.Registry) apiserver.RESTStorage {
	return &BuildLogRegistryStorage{
		BuildRegistry: b,
		PodRegistry: p,
	}
}

var logRegexp = regexp.MustCompile(`.*\[([A-Z][a-z]{2}\s+[0-3]{0,1}[0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9].[0-9][0-9][0-9])\]\s+(.*)`)

func (storage *BuildLogRegistryStorage) Get(id string) (interface{}, error) {

	build, err := storage.BuildRegistry.GetBuild(id)
	if err != nil {
		return nil,	fmt.Errorf("No such build")
	}

	pod, err := storage.PodRegistry.GetPod(build.PodID)
	if err != nil {
		return nil,	fmt.Errorf("No such pod")
	}

	// Build will take place only in one container 
	buildType  := pod.DesiredState.Manifest.Containers[0].Name
	buildPodID := build.PodID
	buildHost  := pod.DesiredState.Host

	podInfoGetter := client.HTTPPodInfoGetter{
		Client: http.DefaultClient,
		Port:   10250,
	}

	podInfo, err := podInfoGetter.GetPodInfo(buildHost, buildPodID)
	buildContainerID := podInfo[buildType].ID
	client := &http.DefaultClient

	// TODO Retrive from master IP and Port on which master binds and use it
	// instead of hardcoding the IP and Port.
	req, err := http.NewRequest(
		"GET", 
		fmt.Sprintf("http://127.0.0.1:8080/proxy/minion/%s/containerLogs?containerID=%s", buildHost, buildContainerID), 
		nil,
	)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	log := buildlogapi.BuildLog{}

	logLines, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil,	err
	}

	for _, line := range strings.Split(string(logLines), "\n") {
		if len(line) > 0 {
			matches := logRegexp.FindStringSubmatch(line)
			log.LogItems = append(log.LogItems, buildlogapi.LogItem{Timestamp:matches[1], Message:matches[2]})
		}
	}
	return log, nil
}

func (storage *BuildLogRegistryStorage) New() interface{} {
	return nil
}

func (storage *BuildLogRegistryStorage) List(selector labels.Selector) (interface{}, error) {
	return nil,	fmt.Errorf("BuildLog can only be retrieved")
}

func (storage *BuildLogRegistryStorage) Delete(id string) (<-chan interface{}, error) {
	return apiserver.MakeAsync(func() (interface{}, error) {
		return nil,	fmt.Errorf("BuildLog can only be retrieved")
	}), nil
}

func (storage *BuildLogRegistryStorage) Extract(body []byte) (interface{}, error) {
	return nil,	fmt.Errorf("BuildLog can only be retrieved")
}

func (storage *BuildLogRegistryStorage) Create(obj interface{}) (<-chan interface{}, error) {
	return apiserver.MakeAsync(func() (interface{}, error) {
		return nil,	fmt.Errorf("BuildLog can only be retrieved")
	}), nil
}

func (storage *BuildLogRegistryStorage) Update(obj interface{}) (<-chan interface{}, error) {
	return apiserver.MakeAsync(func() (interface{}, error) {
		return nil,	fmt.Errorf("BuildLog can only be retrieved")
	}), nil
}