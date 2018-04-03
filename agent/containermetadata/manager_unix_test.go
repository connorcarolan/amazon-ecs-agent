// +build !integration,!windows
// Copyright 2017-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package containermetadata

import (
	"testing"

	"github.com/aws/amazon-ecs-agent/agent/api"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestCreate is the happypath case for metadata create
func TestCreate(t *testing.T) {
	_, mockIOUtil, mockOS, mockFile, done := managerSetup(t)
	defer done()

	mockTaskARN := validTaskARN
	mockTask := &api.Task{Arn: mockTaskARN}
	mockContainerName := containerName
	mockConfig := &docker.Config{Env: make([]string, 0)}
	mockHostConfig := &docker.HostConfig{Binds: make([]string, 0)}

	gomock.InOrder(
		mockOS.EXPECT().MkdirAll(gomock.Any(), gomock.Any()).Return(nil),
		mockIOUtil.EXPECT().TempFile(gomock.Any(), gomock.Any()).Return(mockFile, nil),
		mockFile.EXPECT().Write(gomock.Any()).Return(0, nil),
		mockFile.EXPECT().Chmod(gomock.Any()).Return(nil),
		mockFile.EXPECT().Name().Return(""),
		mockOS.EXPECT().Rename(gomock.Any(), gomock.Any()).Return(nil),
		mockFile.EXPECT().Close().Return(nil),
	)

	newManager := &metadataManager{
		osWrap:     mockOS,
		ioutilWrap: mockIOUtil,
	}
	err := newManager.Create(mockConfig, mockHostConfig, mockTask, mockContainerName)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(mockConfig.Env), "Unexpected number of environment variables in config")
	assert.Equal(t, 1, len(mockHostConfig.Binds), "Unexpected number of binds in host config")
}

// TestUpdate is happypath case for metadata update
func TestUpdate(t *testing.T) {
	mockClient, mockIOUtil, mockOS, mockFile, done := managerSetup(t)
	defer done()

	mockDockerID := dockerID
	mockTaskARN := validTaskARN
	mockTask := &api.Task{Arn: mockTaskARN}
	mockContainerName := containerName
	mockState := docker.State{
		Running: true,
	}

	mockConfig := &docker.Config{Image: "image"}

	mockNetworks := make(map[string]docker.ContainerNetwork)
	mockNetworkSettings := &docker.NetworkSettings{Networks: mockNetworks}

	mockContainer := &docker.Container{
		State:           mockState,
		Config:          mockConfig,
		NetworkSettings: mockNetworkSettings,
	}

	newManager := &metadataManager{
		client:     mockClient,
		osWrap:     mockOS,
		ioutilWrap: mockIOUtil,
	}

	gomock.InOrder(
		mockClient.EXPECT().InspectContainer(mockDockerID, inspectContainerTimeout).Return(mockContainer, nil),
		mockIOUtil.EXPECT().TempFile(gomock.Any(), gomock.Any()).Return(mockFile, nil),
		mockFile.EXPECT().Write(gomock.Any()).Return(0, nil),
		mockFile.EXPECT().Chmod(gomock.Any()).Return(nil),
		mockFile.EXPECT().Name().Return(""),
		mockOS.EXPECT().Rename(gomock.Any(), gomock.Any()).Return(nil),
		mockFile.EXPECT().Close().Return(nil),
	)
	err := newManager.Update(mockDockerID, mockTask, mockContainerName)

	assert.NoError(t, err)
}
