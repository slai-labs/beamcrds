/*
MIT License

Copyright (c) 2022 Slai

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	beamrunnerv1 "github.com/slai-labs/beamcrds/pkg/beamrunner/v1"
	versioned "github.com/slai-labs/beamcrds/pkg/beamrunner/v1/clientset/versioned"
	internalinterfaces "github.com/slai-labs/beamcrds/pkg/beamrunner/v1/informers/externalversions/internalinterfaces"
	v1 "github.com/slai-labs/beamcrds/pkg/beamrunner/v1/listers/beamrunner/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BeamRunnerInformer provides access to a shared informer and lister for
// BeamRunners.
type BeamRunnerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BeamRunnerLister
}

type beamRunnerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBeamRunnerInformer constructs a new informer for BeamRunner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBeamRunnerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBeamRunnerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBeamRunnerInformer constructs a new informer for BeamRunner type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBeamRunnerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1().BeamRunners(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorkV1().BeamRunners(namespace).Watch(context.TODO(), options)
			},
		},
		&beamrunnerv1.BeamRunner{},
		resyncPeriod,
		indexers,
	)
}

func (f *beamRunnerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBeamRunnerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *beamRunnerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&beamrunnerv1.BeamRunner{}, f.defaultInformer)
}

func (f *beamRunnerInformer) Lister() v1.BeamRunnerLister {
	return v1.NewBeamRunnerLister(f.Informer().GetIndexer())
}
