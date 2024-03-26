//go:build !ignore_autogenerated_capd
// +build !ignore_autogenerated_capd

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	corev1alpha3 "sigs.k8s.io/cluster-api/internal/apis/core/v1alpha3"
	v1beta1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1beta1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*APIEndpoint)(nil), (*v1beta1.APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(a.(*APIEndpoint), b.(*v1beta1.APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.APIEndpoint)(nil), (*APIEndpoint)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(a.(*v1beta1.APIEndpoint), b.(*APIEndpoint), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerCluster)(nil), (*v1beta1.DockerCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster(a.(*DockerCluster), b.(*v1beta1.DockerCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerCluster)(nil), (*DockerCluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster(a.(*v1beta1.DockerCluster), b.(*DockerCluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerClusterList)(nil), (*v1beta1.DockerClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerClusterList_To_v1beta1_DockerClusterList(a.(*DockerClusterList), b.(*v1beta1.DockerClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerClusterList)(nil), (*DockerClusterList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerClusterList_To_v1alpha3_DockerClusterList(a.(*v1beta1.DockerClusterList), b.(*DockerClusterList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerClusterSpec)(nil), (*v1beta1.DockerClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec(a.(*DockerClusterSpec), b.(*v1beta1.DockerClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerClusterStatus)(nil), (*v1beta1.DockerClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus(a.(*DockerClusterStatus), b.(*v1beta1.DockerClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerClusterStatus)(nil), (*DockerClusterStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus(a.(*v1beta1.DockerClusterStatus), b.(*DockerClusterStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachine)(nil), (*v1beta1.DockerMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine(a.(*DockerMachine), b.(*v1beta1.DockerMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachine)(nil), (*DockerMachine)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine(a.(*v1beta1.DockerMachine), b.(*DockerMachine), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineList)(nil), (*v1beta1.DockerMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineList_To_v1beta1_DockerMachineList(a.(*DockerMachineList), b.(*v1beta1.DockerMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineList)(nil), (*DockerMachineList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineList_To_v1alpha3_DockerMachineList(a.(*v1beta1.DockerMachineList), b.(*DockerMachineList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineSpec)(nil), (*v1beta1.DockerMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(a.(*DockerMachineSpec), b.(*v1beta1.DockerMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineSpec)(nil), (*DockerMachineSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(a.(*v1beta1.DockerMachineSpec), b.(*DockerMachineSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineStatus)(nil), (*v1beta1.DockerMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus(a.(*DockerMachineStatus), b.(*v1beta1.DockerMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineStatus)(nil), (*DockerMachineStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus(a.(*v1beta1.DockerMachineStatus), b.(*DockerMachineStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineTemplate)(nil), (*v1beta1.DockerMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate(a.(*DockerMachineTemplate), b.(*v1beta1.DockerMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineTemplate)(nil), (*DockerMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate(a.(*v1beta1.DockerMachineTemplate), b.(*DockerMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineTemplateList)(nil), (*v1beta1.DockerMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineTemplateList_To_v1beta1_DockerMachineTemplateList(a.(*DockerMachineTemplateList), b.(*v1beta1.DockerMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineTemplateList)(nil), (*DockerMachineTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineTemplateList_To_v1alpha3_DockerMachineTemplateList(a.(*v1beta1.DockerMachineTemplateList), b.(*DockerMachineTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineTemplateResource)(nil), (*v1beta1.DockerMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource(a.(*DockerMachineTemplateResource), b.(*v1beta1.DockerMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DockerMachineTemplateSpec)(nil), (*v1beta1.DockerMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec(a.(*DockerMachineTemplateSpec), b.(*v1beta1.DockerMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.DockerMachineTemplateSpec)(nil), (*DockerMachineTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec(a.(*v1beta1.DockerMachineTemplateSpec), b.(*DockerMachineTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Mount)(nil), (*v1beta1.Mount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Mount_To_v1beta1_Mount(a.(*Mount), b.(*v1beta1.Mount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.Mount)(nil), (*Mount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_Mount_To_v1alpha3_Mount(a.(*v1beta1.Mount), b.(*Mount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.DockerClusterSpec)(nil), (*DockerClusterSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerClusterSpec_To_v1alpha3_DockerClusterSpec(a.(*v1beta1.DockerClusterSpec), b.(*DockerClusterSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.DockerMachineTemplateResource)(nil), (*DockerMachineTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_DockerMachineTemplateResource_To_v1alpha3_DockerMachineTemplateResource(a.(*v1beta1.DockerMachineTemplateResource), b.(*DockerMachineTemplateResource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(in *APIEndpoint, out *v1beta1.APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(in *APIEndpoint, out *v1beta1.APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(in, out, s)
}

func autoConvert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(in *v1beta1.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	out.Host = in.Host
	out.Port = in.Port
	return nil
}

// Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint is an autogenerated conversion function.
func Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(in *v1beta1.APIEndpoint, out *APIEndpoint, s conversion.Scope) error {
	return autoConvert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(in, out, s)
}

func autoConvert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster(in *DockerCluster, out *v1beta1.DockerCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster is an autogenerated conversion function.
func Convert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster(in *DockerCluster, out *v1beta1.DockerCluster, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster(in, out, s)
}

func autoConvert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster(in *v1beta1.DockerCluster, out *DockerCluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_DockerClusterSpec_To_v1alpha3_DockerClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster is an autogenerated conversion function.
func Convert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster(in *v1beta1.DockerCluster, out *DockerCluster, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster(in, out, s)
}

func autoConvert_v1alpha3_DockerClusterList_To_v1beta1_DockerClusterList(in *DockerClusterList, out *v1beta1.DockerClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.DockerCluster, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_DockerCluster_To_v1beta1_DockerCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_DockerClusterList_To_v1beta1_DockerClusterList is an autogenerated conversion function.
func Convert_v1alpha3_DockerClusterList_To_v1beta1_DockerClusterList(in *DockerClusterList, out *v1beta1.DockerClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerClusterList_To_v1beta1_DockerClusterList(in, out, s)
}

func autoConvert_v1beta1_DockerClusterList_To_v1alpha3_DockerClusterList(in *v1beta1.DockerClusterList, out *DockerClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerCluster, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_DockerCluster_To_v1alpha3_DockerCluster(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_DockerClusterList_To_v1alpha3_DockerClusterList is an autogenerated conversion function.
func Convert_v1beta1_DockerClusterList_To_v1alpha3_DockerClusterList(in *v1beta1.DockerClusterList, out *DockerClusterList, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerClusterList_To_v1alpha3_DockerClusterList(in, out, s)
}

func autoConvert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec(in *DockerClusterSpec, out *v1beta1.DockerClusterSpec, s conversion.Scope) error {
	if err := Convert_v1alpha3_APIEndpoint_To_v1beta1_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(apiv1beta1.FailureDomainSpec)
			if err := corev1alpha3.Convert_v1alpha3_FailureDomainSpec_To_v1beta1_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	return nil
}

// Convert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec is an autogenerated conversion function.
func Convert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec(in *DockerClusterSpec, out *v1beta1.DockerClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerClusterSpec_To_v1beta1_DockerClusterSpec(in, out, s)
}

func autoConvert_v1beta1_DockerClusterSpec_To_v1alpha3_DockerClusterSpec(in *v1beta1.DockerClusterSpec, out *DockerClusterSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_APIEndpoint_To_v1alpha3_APIEndpoint(&in.ControlPlaneEndpoint, &out.ControlPlaneEndpoint, s); err != nil {
		return err
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(corev1alpha3.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(corev1alpha3.FailureDomainSpec)
			if err := corev1alpha3.Convert_v1beta1_FailureDomainSpec_To_v1alpha3_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	// WARNING: in.LoadBalancer requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus(in *DockerClusterStatus, out *v1beta1.DockerClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(apiv1beta1.FailureDomainSpec)
			if err := corev1alpha3.Convert_v1alpha3_FailureDomainSpec_To_v1beta1_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1alpha3_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus is an autogenerated conversion function.
func Convert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus(in *DockerClusterStatus, out *v1beta1.DockerClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerClusterStatus_To_v1beta1_DockerClusterStatus(in, out, s)
}

func autoConvert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus(in *v1beta1.DockerClusterStatus, out *DockerClusterStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(corev1alpha3.FailureDomains, len(*in))
		for key, val := range *in {
			newVal := new(corev1alpha3.FailureDomainSpec)
			if err := corev1alpha3.Convert_v1beta1_FailureDomainSpec_To_v1alpha3_FailureDomainSpec(&val, newVal, s); err != nil {
				return err
			}
			(*out)[key] = *newVal
		}
	} else {
		out.FailureDomains = nil
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1alpha3.Conditions, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1beta1_Condition_To_v1alpha3_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus is an autogenerated conversion function.
func Convert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus(in *v1beta1.DockerClusterStatus, out *DockerClusterStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerClusterStatus_To_v1alpha3_DockerClusterStatus(in, out, s)
}

func autoConvert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine(in *DockerMachine, out *v1beta1.DockerMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine(in *DockerMachine, out *v1beta1.DockerMachine, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine(in, out, s)
}

func autoConvert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine(in *v1beta1.DockerMachine, out *DockerMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine is an autogenerated conversion function.
func Convert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine(in *v1beta1.DockerMachine, out *DockerMachine, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineList_To_v1beta1_DockerMachineList(in *DockerMachineList, out *v1beta1.DockerMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.DockerMachine, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_DockerMachine_To_v1beta1_DockerMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_DockerMachineList_To_v1beta1_DockerMachineList is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineList_To_v1beta1_DockerMachineList(in *DockerMachineList, out *v1beta1.DockerMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineList_To_v1beta1_DockerMachineList(in, out, s)
}

func autoConvert_v1beta1_DockerMachineList_To_v1alpha3_DockerMachineList(in *v1beta1.DockerMachineList, out *DockerMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerMachine, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_DockerMachine_To_v1alpha3_DockerMachine(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_DockerMachineList_To_v1alpha3_DockerMachineList is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineList_To_v1alpha3_DockerMachineList(in *v1beta1.DockerMachineList, out *DockerMachineList, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineList_To_v1alpha3_DockerMachineList(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(in *DockerMachineSpec, out *v1beta1.DockerMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.CustomImage = in.CustomImage
	out.PreLoadImages = *(*[]string)(unsafe.Pointer(&in.PreLoadImages))
	out.ExtraMounts = *(*[]v1beta1.Mount)(unsafe.Pointer(&in.ExtraMounts))
	out.Bootstrapped = in.Bootstrapped
	return nil
}

// Convert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(in *DockerMachineSpec, out *v1beta1.DockerMachineSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(in, out, s)
}

func autoConvert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(in *v1beta1.DockerMachineSpec, out *DockerMachineSpec, s conversion.Scope) error {
	out.ProviderID = (*string)(unsafe.Pointer(in.ProviderID))
	out.CustomImage = in.CustomImage
	out.PreLoadImages = *(*[]string)(unsafe.Pointer(&in.PreLoadImages))
	out.ExtraMounts = *(*[]Mount)(unsafe.Pointer(&in.ExtraMounts))
	out.Bootstrapped = in.Bootstrapped
	return nil
}

// Convert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(in *v1beta1.DockerMachineSpec, out *DockerMachineSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus(in *DockerMachineStatus, out *v1beta1.DockerMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.LoadBalancerConfigured = in.LoadBalancerConfigured
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]apiv1beta1.MachineAddress, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1alpha3_MachineAddress_To_v1beta1_MachineAddress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Addresses = nil
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1alpha3_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus(in *DockerMachineStatus, out *v1beta1.DockerMachineStatus, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineStatus_To_v1beta1_DockerMachineStatus(in, out, s)
}

func autoConvert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus(in *v1beta1.DockerMachineStatus, out *DockerMachineStatus, s conversion.Scope) error {
	out.Ready = in.Ready
	out.LoadBalancerConfigured = in.LoadBalancerConfigured
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]corev1alpha3.MachineAddress, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1beta1_MachineAddress_To_v1alpha3_MachineAddress(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Addresses = nil
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(corev1alpha3.Conditions, len(*in))
		for i := range *in {
			if err := corev1alpha3.Convert_v1beta1_Condition_To_v1alpha3_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus(in *v1beta1.DockerMachineStatus, out *DockerMachineStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineStatus_To_v1alpha3_DockerMachineStatus(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate(in *DockerMachineTemplate, out *v1beta1.DockerMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate(in *DockerMachineTemplate, out *v1beta1.DockerMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate(in, out, s)
}

func autoConvert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate(in *v1beta1.DockerMachineTemplate, out *DockerMachineTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate(in *v1beta1.DockerMachineTemplate, out *DockerMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineTemplateList_To_v1beta1_DockerMachineTemplateList(in *DockerMachineTemplateList, out *v1beta1.DockerMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.DockerMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha3_DockerMachineTemplate_To_v1beta1_DockerMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha3_DockerMachineTemplateList_To_v1beta1_DockerMachineTemplateList is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineTemplateList_To_v1beta1_DockerMachineTemplateList(in *DockerMachineTemplateList, out *v1beta1.DockerMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineTemplateList_To_v1beta1_DockerMachineTemplateList(in, out, s)
}

func autoConvert_v1beta1_DockerMachineTemplateList_To_v1alpha3_DockerMachineTemplateList(in *v1beta1.DockerMachineTemplateList, out *DockerMachineTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerMachineTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_DockerMachineTemplate_To_v1alpha3_DockerMachineTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_DockerMachineTemplateList_To_v1alpha3_DockerMachineTemplateList is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineTemplateList_To_v1alpha3_DockerMachineTemplateList(in *v1beta1.DockerMachineTemplateList, out *DockerMachineTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineTemplateList_To_v1alpha3_DockerMachineTemplateList(in, out, s)
}

func autoConvert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource(in *DockerMachineTemplateResource, out *v1beta1.DockerMachineTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha3_DockerMachineSpec_To_v1beta1_DockerMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource(in *DockerMachineTemplateResource, out *v1beta1.DockerMachineTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource(in, out, s)
}

func autoConvert_v1beta1_DockerMachineTemplateResource_To_v1alpha3_DockerMachineTemplateResource(in *v1beta1.DockerMachineTemplateResource, out *DockerMachineTemplateResource, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	if err := Convert_v1beta1_DockerMachineSpec_To_v1alpha3_DockerMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec(in *DockerMachineTemplateSpec, out *v1beta1.DockerMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha3_DockerMachineTemplateResource_To_v1beta1_DockerMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec(in *DockerMachineTemplateSpec, out *v1beta1.DockerMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha3_DockerMachineTemplateSpec_To_v1beta1_DockerMachineTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec(in *v1beta1.DockerMachineTemplateSpec, out *DockerMachineTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_DockerMachineTemplateResource_To_v1alpha3_DockerMachineTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec(in *v1beta1.DockerMachineTemplateSpec, out *DockerMachineTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_DockerMachineTemplateSpec_To_v1alpha3_DockerMachineTemplateSpec(in, out, s)
}

func autoConvert_v1alpha3_Mount_To_v1beta1_Mount(in *Mount, out *v1beta1.Mount, s conversion.Scope) error {
	out.ContainerPath = in.ContainerPath
	out.HostPath = in.HostPath
	out.Readonly = in.Readonly
	return nil
}

// Convert_v1alpha3_Mount_To_v1beta1_Mount is an autogenerated conversion function.
func Convert_v1alpha3_Mount_To_v1beta1_Mount(in *Mount, out *v1beta1.Mount, s conversion.Scope) error {
	return autoConvert_v1alpha3_Mount_To_v1beta1_Mount(in, out, s)
}

func autoConvert_v1beta1_Mount_To_v1alpha3_Mount(in *v1beta1.Mount, out *Mount, s conversion.Scope) error {
	out.ContainerPath = in.ContainerPath
	out.HostPath = in.HostPath
	out.Readonly = in.Readonly
	return nil
}

// Convert_v1beta1_Mount_To_v1alpha3_Mount is an autogenerated conversion function.
func Convert_v1beta1_Mount_To_v1alpha3_Mount(in *v1beta1.Mount, out *Mount, s conversion.Scope) error {
	return autoConvert_v1beta1_Mount_To_v1alpha3_Mount(in, out, s)
}
