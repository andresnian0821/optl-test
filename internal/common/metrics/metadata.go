// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

// MetadataExporter provides an interface to implement
// ConsumeMetadata in Exporters that support metadata.
type MetadataExporter interface {
	// ConsumeMetadata will be invoked every time there's an
	// update to a resource that results in one or more MetadataUpdate.
	ConsumeMetadata(metadata []*MetadataUpdate) error
}

type ResourceID string

// MetadataDelta keeps track of changes to metadata on resources.
// The fields on this struct should help determine if there have
// been changes to resource metadata such as Kubernetes labels.
// An example of how this is used. Let's say we are dealing with a
// Pod that has the following labels -
// {"env": "test", "team": "otell", "usser": "bob"}. Now, let's say
// there's an update to one or more labels on the same Pod and the
// labels now look like the following -
// {"env": "test", "team": "otel", "user": "bob"}. The k8sclusterreceiver
// upon receiving the event corresponding to the labels updates will
// generate a MetadataDelta with the following values -
// 					MetadataToAdd: {"user": "bob"}
// 					MetadataToRemove: {"usser": "bob"}
// 					MetadataToUpdate: {"team": "otel"}
// Apart from Kubernetes labels, the other metadata collected by this
// receiver are also handled in the same manner.
type MetadataDelta struct {
	// MetadataToAdd contains key-value pairs that are newly added to
	// the resource description in the current revision.
	MetadataToAdd map[string]string
	// MetadataToRemove contains key-value pairs that no longer describe
	// a resource and needs to be removed.
	MetadataToRemove map[string]string
	// MetadataToUpdate contains key-value pairs that have been updated
	// in the current revision compared to the previous revisions(s).
	MetadataToUpdate map[string]string
}

// MetadataUpdate provides a delta view of metadata on a resource between
// two revisions of a resource.
type MetadataUpdate struct {
	// ResourceIDKey is the label key of UID label for the resource.
	ResourceIDKey string
	// ResourceID is the Kubernetes UID of the resource. In case of
	// containers, this value is the container id.
	ResourceID ResourceID
	MetadataDelta
}
