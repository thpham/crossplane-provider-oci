//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationObservation) DeepCopyInto(out *ApplicationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationObservation.
func (in *ApplicationObservation) DeepCopy() *ApplicationObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParameters) DeepCopyInto(out *ApplicationParameters) {
	*out = *in
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CompartmentIDRef != nil {
		in, out := &in.CompartmentIDRef, &out.CompartmentIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentIDSelector != nil {
		in, out := &in.CompartmentIDSelector, &out.CompartmentIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.ImagePolicyConfig != nil {
		in, out := &in.ImagePolicyConfig, &out.ImagePolicyConfig
		*out = make([]ImagePolicyConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkSecurityGroupIds != nil {
		in, out := &in.NetworkSecurityGroupIds, &out.NetworkSecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NetworkSecurityGroupIdsRefs != nil {
		in, out := &in.NetworkSecurityGroupIdsRefs, &out.NetworkSecurityGroupIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkSecurityGroupIdsSelector != nil {
		in, out := &in.NetworkSecurityGroupIdsSelector, &out.NetworkSecurityGroupIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIdsRefs != nil {
		in, out := &in.SubnetIdsRefs, &out.SubnetIdsRefs
		*out = make([]v1.Reference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetIdsSelector != nil {
		in, out := &in.SubnetIdsSelector, &out.SubnetIdsSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SyslogURL != nil {
		in, out := &in.SyslogURL, &out.SyslogURL
		*out = new(string)
		**out = **in
	}
	if in.TraceConfig != nil {
		in, out := &in.TraceConfig, &out.TraceConfig
		*out = make([]TraceConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParameters.
func (in *ApplicationParameters) DeepCopy() *ApplicationParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InvokeEndpoint != nil {
		in, out := &in.InvokeEndpoint, &out.InvokeEndpoint
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.ApplicationIDRef != nil {
		in, out := &in.ApplicationIDRef, &out.ApplicationIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplicationIDSelector != nil {
		in, out := &in.ApplicationIDSelector, &out.ApplicationIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImageDigest != nil {
		in, out := &in.ImageDigest, &out.ImageDigest
		*out = new(string)
		**out = **in
	}
	if in.MemoryInMbs != nil {
		in, out := &in.MemoryInMbs, &out.MemoryInMbs
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedConcurrencyConfig != nil {
		in, out := &in.ProvisionedConcurrencyConfig, &out.ProvisionedConcurrencyConfig
		*out = make([]ProvisionedConcurrencyConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SourceDetails != nil {
		in, out := &in.SourceDetails, &out.SourceDetails
		*out = make([]SourceDetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(float64)
		**out = **in
	}
	if in.TraceConfig != nil {
		in, out := &in.TraceConfig, &out.TraceConfig
		*out = make([]FunctionTraceConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionTraceConfigObservation) DeepCopyInto(out *FunctionTraceConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionTraceConfigObservation.
func (in *FunctionTraceConfigObservation) DeepCopy() *FunctionTraceConfigObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionTraceConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionTraceConfigParameters) DeepCopyInto(out *FunctionTraceConfigParameters) {
	*out = *in
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionTraceConfigParameters.
func (in *FunctionTraceConfigParameters) DeepCopy() *FunctionTraceConfigParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionTraceConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyConfigObservation) DeepCopyInto(out *ImagePolicyConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyConfigObservation.
func (in *ImagePolicyConfigObservation) DeepCopy() *ImagePolicyConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePolicyConfigParameters) DeepCopyInto(out *ImagePolicyConfigParameters) {
	*out = *in
	if in.IsPolicyEnabled != nil {
		in, out := &in.IsPolicyEnabled, &out.IsPolicyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyDetails != nil {
		in, out := &in.KeyDetails, &out.KeyDetails
		*out = make([]KeyDetailsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePolicyConfigParameters.
func (in *ImagePolicyConfigParameters) DeepCopy() *ImagePolicyConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ImagePolicyConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunction) DeepCopyInto(out *InvokeFunction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunction.
func (in *InvokeFunction) DeepCopy() *InvokeFunction {
	if in == nil {
		return nil
	}
	out := new(InvokeFunction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvokeFunction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunctionList) DeepCopyInto(out *InvokeFunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InvokeFunction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunctionList.
func (in *InvokeFunctionList) DeepCopy() *InvokeFunctionList {
	if in == nil {
		return nil
	}
	out := new(InvokeFunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvokeFunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunctionObservation) DeepCopyInto(out *InvokeFunctionObservation) {
	*out = *in
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InvokeEndpoint != nil {
		in, out := &in.InvokeEndpoint, &out.InvokeEndpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunctionObservation.
func (in *InvokeFunctionObservation) DeepCopy() *InvokeFunctionObservation {
	if in == nil {
		return nil
	}
	out := new(InvokeFunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunctionParameters) DeepCopyInto(out *InvokeFunctionParameters) {
	*out = *in
	if in.Base64EncodeContent != nil {
		in, out := &in.Base64EncodeContent, &out.Base64EncodeContent
		*out = new(bool)
		**out = **in
	}
	if in.FnIntent != nil {
		in, out := &in.FnIntent, &out.FnIntent
		*out = new(string)
		**out = **in
	}
	if in.FnInvokeType != nil {
		in, out := &in.FnInvokeType, &out.FnInvokeType
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.FunctionIDRef != nil {
		in, out := &in.FunctionIDRef, &out.FunctionIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.FunctionIDSelector != nil {
		in, out := &in.FunctionIDSelector, &out.FunctionIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.InputBodySourcePath != nil {
		in, out := &in.InputBodySourcePath, &out.InputBodySourcePath
		*out = new(string)
		**out = **in
	}
	if in.InvokeFunctionBody != nil {
		in, out := &in.InvokeFunctionBody, &out.InvokeFunctionBody
		*out = new(string)
		**out = **in
	}
	if in.InvokeFunctionBodyBase64Encoded != nil {
		in, out := &in.InvokeFunctionBodyBase64Encoded, &out.InvokeFunctionBodyBase64Encoded
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunctionParameters.
func (in *InvokeFunctionParameters) DeepCopy() *InvokeFunctionParameters {
	if in == nil {
		return nil
	}
	out := new(InvokeFunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunctionSpec) DeepCopyInto(out *InvokeFunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunctionSpec.
func (in *InvokeFunctionSpec) DeepCopy() *InvokeFunctionSpec {
	if in == nil {
		return nil
	}
	out := new(InvokeFunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeFunctionStatus) DeepCopyInto(out *InvokeFunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeFunctionStatus.
func (in *InvokeFunctionStatus) DeepCopy() *InvokeFunctionStatus {
	if in == nil {
		return nil
	}
	out := new(InvokeFunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyDetailsObservation) DeepCopyInto(out *KeyDetailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyDetailsObservation.
func (in *KeyDetailsObservation) DeepCopy() *KeyDetailsObservation {
	if in == nil {
		return nil
	}
	out := new(KeyDetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyDetailsParameters) DeepCopyInto(out *KeyDetailsParameters) {
	*out = *in
	if in.KMSKeyID != nil {
		in, out := &in.KMSKeyID, &out.KMSKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyDetailsParameters.
func (in *KeyDetailsParameters) DeepCopy() *KeyDetailsParameters {
	if in == nil {
		return nil
	}
	out := new(KeyDetailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedConcurrencyConfigObservation) DeepCopyInto(out *ProvisionedConcurrencyConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedConcurrencyConfigObservation.
func (in *ProvisionedConcurrencyConfigObservation) DeepCopy() *ProvisionedConcurrencyConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ProvisionedConcurrencyConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionedConcurrencyConfigParameters) DeepCopyInto(out *ProvisionedConcurrencyConfigParameters) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(float64)
		**out = **in
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionedConcurrencyConfigParameters.
func (in *ProvisionedConcurrencyConfigParameters) DeepCopy() *ProvisionedConcurrencyConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ProvisionedConcurrencyConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceDetailsObservation) DeepCopyInto(out *SourceDetailsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceDetailsObservation.
func (in *SourceDetailsObservation) DeepCopy() *SourceDetailsObservation {
	if in == nil {
		return nil
	}
	out := new(SourceDetailsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceDetailsParameters) DeepCopyInto(out *SourceDetailsParameters) {
	*out = *in
	if in.PbfListingID != nil {
		in, out := &in.PbfListingID, &out.PbfListingID
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceDetailsParameters.
func (in *SourceDetailsParameters) DeepCopy() *SourceDetailsParameters {
	if in == nil {
		return nil
	}
	out := new(SourceDetailsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceConfigObservation) DeepCopyInto(out *TraceConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceConfigObservation.
func (in *TraceConfigObservation) DeepCopy() *TraceConfigObservation {
	if in == nil {
		return nil
	}
	out := new(TraceConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceConfigParameters) DeepCopyInto(out *TraceConfigParameters) {
	*out = *in
	if in.DomainID != nil {
		in, out := &in.DomainID, &out.DomainID
		*out = new(string)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceConfigParameters.
func (in *TraceConfigParameters) DeepCopy() *TraceConfigParameters {
	if in == nil {
		return nil
	}
	out := new(TraceConfigParameters)
	in.DeepCopyInto(out)
	return out
}
