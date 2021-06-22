// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSet) DeepCopyInto(out *ApplicationSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSet.
func (in *ApplicationSet) DeepCopy() *ApplicationSet {
	if in == nil {
		return nil
	}
	out := new(ApplicationSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetBaseGenerator) DeepCopyInto(out *ApplicationSetBaseGenerator) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = new(ListGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = new(ClusterGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.SCMProvider != nil {
		in, out := &in.SCMProvider, &out.SCMProvider
		*out = new(SCMProviderGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterDecisionResource != nil {
		in, out := &in.ClusterDecisionResource, &out.ClusterDecisionResource
		*out = new(DuckTypeGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetBaseGenerator.
func (in *ApplicationSetBaseGenerator) DeepCopy() *ApplicationSetBaseGenerator {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetBaseGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetGenerator) DeepCopyInto(out *ApplicationSetGenerator) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = new(ListGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = new(ClusterGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(GitGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.Matrix != nil {
		in, out := &in.Matrix, &out.Matrix
		*out = new(MatrixGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.SCMProvider != nil {
		in, out := &in.SCMProvider, &out.SCMProvider
		*out = new(SCMProviderGenerator)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterDecisionResource != nil {
		in, out := &in.ClusterDecisionResource, &out.ClusterDecisionResource
		*out = new(DuckTypeGenerator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetGenerator.
func (in *ApplicationSetGenerator) DeepCopy() *ApplicationSetGenerator {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetList) DeepCopyInto(out *ApplicationSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetList.
func (in *ApplicationSetList) DeepCopy() *ApplicationSetList {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetSpec) DeepCopyInto(out *ApplicationSetSpec) {
	*out = *in
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]ApplicationSetGenerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	if in.SyncPolicy != nil {
		in, out := &in.SyncPolicy, &out.SyncPolicy
		*out = new(ApplicationSetSyncPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetSpec.
func (in *ApplicationSetSpec) DeepCopy() *ApplicationSetSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetStatus) DeepCopyInto(out *ApplicationSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetStatus.
func (in *ApplicationSetStatus) DeepCopy() *ApplicationSetStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetSyncPolicy) DeepCopyInto(out *ApplicationSetSyncPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetSyncPolicy.
func (in *ApplicationSetSyncPolicy) DeepCopy() *ApplicationSetSyncPolicy {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetSyncPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetTemplate) DeepCopyInto(out *ApplicationSetTemplate) {
	*out = *in
	in.ApplicationSetTemplateMeta.DeepCopyInto(&out.ApplicationSetTemplateMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetTemplate.
func (in *ApplicationSetTemplate) DeepCopy() *ApplicationSetTemplate {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSetTemplateMeta) DeepCopyInto(out *ApplicationSetTemplateMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSetTemplateMeta.
func (in *ApplicationSetTemplateMeta) DeepCopy() *ApplicationSetTemplateMeta {
	if in == nil {
		return nil
	}
	out := new(ApplicationSetTemplateMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterGenerator) DeepCopyInto(out *ClusterGenerator) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Template.DeepCopyInto(&out.Template)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterGenerator.
func (in *ClusterGenerator) DeepCopy() *ClusterGenerator {
	if in == nil {
		return nil
	}
	out := new(ClusterGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DuckTypeGenerator) DeepCopyInto(out *DuckTypeGenerator) {
	*out = *in
	if in.RequeueAfterSeconds != nil {
		in, out := &in.RequeueAfterSeconds, &out.RequeueAfterSeconds
		*out = new(int64)
		**out = **in
	}
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	in.Template.DeepCopyInto(&out.Template)
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DuckTypeGenerator.
func (in *DuckTypeGenerator) DeepCopy() *DuckTypeGenerator {
	if in == nil {
		return nil
	}
	out := new(DuckTypeGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitDirectoryGeneratorItem) DeepCopyInto(out *GitDirectoryGeneratorItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitDirectoryGeneratorItem.
func (in *GitDirectoryGeneratorItem) DeepCopy() *GitDirectoryGeneratorItem {
	if in == nil {
		return nil
	}
	out := new(GitDirectoryGeneratorItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitFileGeneratorItem) DeepCopyInto(out *GitFileGeneratorItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitFileGeneratorItem.
func (in *GitFileGeneratorItem) DeepCopy() *GitFileGeneratorItem {
	if in == nil {
		return nil
	}
	out := new(GitFileGeneratorItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitGenerator) DeepCopyInto(out *GitGenerator) {
	*out = *in
	if in.Directories != nil {
		in, out := &in.Directories, &out.Directories
		*out = make([]GitDirectoryGeneratorItem, len(*in))
		copy(*out, *in)
	}
	if in.Files != nil {
		in, out := &in.Files, &out.Files
		*out = make([]GitFileGeneratorItem, len(*in))
		copy(*out, *in)
	}
	if in.RequeueAfterSeconds != nil {
		in, out := &in.RequeueAfterSeconds, &out.RequeueAfterSeconds
		*out = new(int64)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitGenerator.
func (in *GitGenerator) DeepCopy() *GitGenerator {
	if in == nil {
		return nil
	}
	out := new(GitGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGenerator) DeepCopyInto(out *ListGenerator) {
	*out = *in
	if in.Elements != nil {
		in, out := &in.Elements, &out.Elements
		*out = make([]ListGeneratorElement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGenerator.
func (in *ListGenerator) DeepCopy() *ListGenerator {
	if in == nil {
		return nil
	}
	out := new(ListGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListGeneratorElement) DeepCopyInto(out *ListGeneratorElement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListGeneratorElement.
func (in *ListGeneratorElement) DeepCopy() *ListGeneratorElement {
	if in == nil {
		return nil
	}
	out := new(ListGeneratorElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatrixGenerator) DeepCopyInto(out *MatrixGenerator) {
	*out = *in
	if in.Generators != nil {
		in, out := &in.Generators, &out.Generators
		*out = make([]ApplicationSetBaseGenerator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatrixGenerator.
func (in *MatrixGenerator) DeepCopy() *MatrixGenerator {
	if in == nil {
		return nil
	}
	out := new(MatrixGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMProviderGenerator) DeepCopyInto(out *SCMProviderGenerator) {
	*out = *in
	if in.Github != nil {
		in, out := &in.Github, &out.Github
		*out = new(SCMProviderGeneratorGithub)
		(*in).DeepCopyInto(*out)
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]SCMProviderGeneratorFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RequeueAfterSeconds != nil {
		in, out := &in.RequeueAfterSeconds, &out.RequeueAfterSeconds
		*out = new(int64)
		**out = **in
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMProviderGenerator.
func (in *SCMProviderGenerator) DeepCopy() *SCMProviderGenerator {
	if in == nil {
		return nil
	}
	out := new(SCMProviderGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMProviderGeneratorFilter) DeepCopyInto(out *SCMProviderGeneratorFilter) {
	*out = *in
	if in.RepositoryMatch != nil {
		in, out := &in.RepositoryMatch, &out.RepositoryMatch
		*out = new(string)
		**out = **in
	}
	if in.PathsExist != nil {
		in, out := &in.PathsExist, &out.PathsExist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LabelMatch != nil {
		in, out := &in.LabelMatch, &out.LabelMatch
		*out = new(string)
		**out = **in
	}
	if in.BranchMatch != nil {
		in, out := &in.BranchMatch, &out.BranchMatch
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMProviderGeneratorFilter.
func (in *SCMProviderGeneratorFilter) DeepCopy() *SCMProviderGeneratorFilter {
	if in == nil {
		return nil
	}
	out := new(SCMProviderGeneratorFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SCMProviderGeneratorGithub) DeepCopyInto(out *SCMProviderGeneratorGithub) {
	*out = *in
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(SecretRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SCMProviderGeneratorGithub.
func (in *SCMProviderGeneratorGithub) DeepCopy() *SCMProviderGeneratorGithub {
	if in == nil {
		return nil
	}
	out := new(SCMProviderGeneratorGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}
