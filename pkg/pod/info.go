/*
Copyright © 2020 FairOS Authors

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

package pod

import (
	"github.com/fairdatasociety/fairOS-dfs/pkg/collection"

	"github.com/fairdatasociety/fairOS-dfs/pkg/account"
	di "github.com/fairdatasociety/fairOS-dfs/pkg/dir"
	"github.com/fairdatasociety/fairOS-dfs/pkg/feed"
	f "github.com/fairdatasociety/fairOS-dfs/pkg/file"
	"github.com/fairdatasociety/fairOS-dfs/pkg/utils"
)

// Info is the struct which holds the pod information
type Info struct {
	podName     string
	podPassword string
	userAddress utils.Address
	dir         *di.Directory
	file        *f.File
	accountInfo *account.Info
	feed        *feed.API
	kvStore     *collection.KeyValue
	docStore    *collection.Document
}

// GetPodName returns the pod name
func (i *Info) GetPodName() string {
	return i.podName
}

// GetPodAddress returns the pod address
func (i *Info) GetPodAddress() utils.Address {
	return i.userAddress
}

// GetPodPassword returns the pod password
func (i *Info) GetPodPassword() string {
	return i.podPassword
}

// GetDirectory returns the directory object
func (i *Info) GetDirectory() *di.Directory {
	return i.dir
}

// GetFile returns the file object
func (i *Info) GetFile() *f.File {
	return i.file
}

// GetAccountInfo returns the pod account info
func (i *Info) GetAccountInfo() *account.Info {
	return i.accountInfo
}

// GetFeed returns the feed object
func (i *Info) GetFeed() *feed.API {
	return i.feed
}

// GetKVStore returns kvStore
// skipcq: TCV-001
func (i *Info) GetKVStore() *collection.KeyValue {
	return i.kvStore
}

// GetDocStore returns docStore
// skipcq: TCV-001
func (i *Info) GetDocStore() *collection.Document {
	return i.docStore
}
