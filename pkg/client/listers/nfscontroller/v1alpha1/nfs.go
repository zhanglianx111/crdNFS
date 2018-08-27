/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/zhanglianx111/nfscontroller/pkg/apis/nfscontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NFSLister helps list NFSs.
type NFSLister interface {
	// List lists all NFSs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NFS, err error)
	// Get retrieves the NFS from the index for a given name.
	Get(name string) (*v1alpha1.NFS, error)
	NFSListerExpansion
}

// nFSLister implements the NFSLister interface.
type nFSLister struct {
	indexer cache.Indexer
}

// NewNFSLister returns a new NFSLister.
func NewNFSLister(indexer cache.Indexer) NFSLister {
	return &nFSLister{indexer: indexer}
}

// List lists all NFSs in the indexer.
func (s *nFSLister) List(selector labels.Selector) (ret []*v1alpha1.NFS, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NFS))
	})
	return ret, err
}

// Get retrieves the NFS from the index for a given name.
func (s *nFSLister) Get(name string) (*v1alpha1.NFS, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nfs"), name)
	}
	return obj.(*v1alpha1.NFS), nil
}