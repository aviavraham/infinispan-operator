//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authorization) DeepCopyInto(out *Authorization) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]AuthorizationRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authorization.
func (in *Authorization) DeepCopy() *Authorization {
	if in == nil {
		return nil
	}
	out := new(Authorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthorizationRole) DeepCopyInto(out *AuthorizationRole) {
	*out = *in
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthorizationRole.
func (in *AuthorizationRole) DeepCopy() *AuthorizationRole {
	if in == nil {
		return nil
	}
	out := new(AuthorizationRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Autoscale) DeepCopyInto(out *Autoscale) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Autoscale.
func (in *Autoscale) DeepCopy() *Autoscale {
	if in == nil {
		return nil
	}
	out := new(Autoscale)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigListenerLoggingSpec) DeepCopyInto(out *ConfigListenerLoggingSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigListenerLoggingSpec.
func (in *ConfigListenerLoggingSpec) DeepCopy() *ConfigListenerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigListenerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigListenerSpec) DeepCopyInto(out *ConfigListenerSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(ConfigListenerLoggingSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigListenerSpec.
func (in *ConfigListenerSpec) DeepCopy() *ConfigListenerSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerProbeSpec) DeepCopyInto(out *ContainerProbeSpec) {
	*out = *in
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SuccessThreshold != nil {
		in, out := &in.SuccessThreshold, &out.SuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerProbeSpec.
func (in *ContainerProbeSpec) DeepCopy() *ContainerProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossSiteExposeSpec) DeepCopyInto(out *CrossSiteExposeSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossSiteExposeSpec.
func (in *CrossSiteExposeSpec) DeepCopy() *CrossSiteExposeSpec {
	if in == nil {
		return nil
	}
	out := new(CrossSiteExposeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossSiteKeyStore) DeepCopyInto(out *CrossSiteKeyStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossSiteKeyStore.
func (in *CrossSiteKeyStore) DeepCopy() *CrossSiteKeyStore {
	if in == nil {
		return nil
	}
	out := new(CrossSiteKeyStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossSiteTrustStore) DeepCopyInto(out *CrossSiteTrustStore) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossSiteTrustStore.
func (in *CrossSiteTrustStore) DeepCopy() *CrossSiteTrustStore {
	if in == nil {
		return nil
	}
	out := new(CrossSiteTrustStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Starting != nil {
		in, out := &in.Starting, &out.Starting
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Stopped != nil {
		in, out := &in.Stopped, &out.Stopped
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoverySiteSpec) DeepCopyInto(out *DiscoverySiteSpec) {
	*out = *in
	if in.LaunchGossipRouter != nil {
		in, out := &in.LaunchGossipRouter, &out.LaunchGossipRouter
		*out = new(bool)
		**out = **in
	}
	if in.Heartbeats != nil {
		in, out := &in.Heartbeats, &out.Heartbeats
		*out = new(GossipRouterHeartbeatSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoverySiteSpec.
func (in *DiscoverySiteSpec) DeepCopy() *DiscoverySiteSpec {
	if in == nil {
		return nil
	}
	out := new(DiscoverySiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSiteSpec) DeepCopyInto(out *EncryptionSiteSpec) {
	*out = *in
	out.TransportKeyStore = in.TransportKeyStore
	out.RouterKeyStore = in.RouterKeyStore
	if in.TrustStore != nil {
		in, out := &in.TrustStore, &out.TrustStore
		*out = new(CrossSiteTrustStore)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSiteSpec.
func (in *EncryptionSiteSpec) DeepCopy() *EncryptionSiteSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionSiteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointEncryption) DeepCopyInto(out *EndpointEncryption) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointEncryption.
func (in *EndpointEncryption) DeepCopy() *EndpointEncryption {
	if in == nil {
		return nil
	}
	out := new(EndpointEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposeSpec) DeepCopyInto(out *ExposeSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposeSpec.
func (in *ExposeSpec) DeepCopy() *ExposeSpec {
	if in == nil {
		return nil
	}
	out := new(ExposeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GossipRouterHeartbeatSpec) DeepCopyInto(out *GossipRouterHeartbeatSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GossipRouterHeartbeatSpec.
func (in *GossipRouterHeartbeatSpec) DeepCopy() *GossipRouterHeartbeatSpec {
	if in == nil {
		return nil
	}
	out := new(GossipRouterHeartbeatSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HotRodRollingUpgradeStatus) DeepCopyInto(out *HotRodRollingUpgradeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HotRodRollingUpgradeStatus.
func (in *HotRodRollingUpgradeStatus) DeepCopy() *HotRodRollingUpgradeStatus {
	if in == nil {
		return nil
	}
	out := new(HotRodRollingUpgradeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Infinispan) DeepCopyInto(out *Infinispan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Infinispan.
func (in *Infinispan) DeepCopy() *Infinispan {
	if in == nil {
		return nil
	}
	out := new(Infinispan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Infinispan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanCloudEvents) DeepCopyInto(out *InfinispanCloudEvents) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanCloudEvents.
func (in *InfinispanCloudEvents) DeepCopy() *InfinispanCloudEvents {
	if in == nil {
		return nil
	}
	out := new(InfinispanCloudEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanCondition) DeepCopyInto(out *InfinispanCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanCondition.
func (in *InfinispanCondition) DeepCopy() *InfinispanCondition {
	if in == nil {
		return nil
	}
	out := new(InfinispanCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanContainerSpec) DeepCopyInto(out *InfinispanContainerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanContainerSpec.
func (in *InfinispanContainerSpec) DeepCopy() *InfinispanContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanExternalArtifacts) DeepCopyInto(out *InfinispanExternalArtifacts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanExternalArtifacts.
func (in *InfinispanExternalArtifacts) DeepCopy() *InfinispanExternalArtifacts {
	if in == nil {
		return nil
	}
	out := new(InfinispanExternalArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanExternalDependencies) DeepCopyInto(out *InfinispanExternalDependencies) {
	*out = *in
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make([]InfinispanExternalArtifacts, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanExternalDependencies.
func (in *InfinispanExternalDependencies) DeepCopy() *InfinispanExternalDependencies {
	if in == nil {
		return nil
	}
	out := new(InfinispanExternalDependencies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanList) DeepCopyInto(out *InfinispanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Infinispan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanList.
func (in *InfinispanList) DeepCopy() *InfinispanList {
	if in == nil {
		return nil
	}
	out := new(InfinispanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfinispanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanLoggingSpec) DeepCopyInto(out *InfinispanLoggingSpec) {
	*out = *in
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make(map[string]LoggingLevelType, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanLoggingSpec.
func (in *InfinispanLoggingSpec) DeepCopy() *InfinispanLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSecurity) DeepCopyInto(out *InfinispanSecurity) {
	*out = *in
	if in.Authorization != nil {
		in, out := &in.Authorization, &out.Authorization
		*out = new(Authorization)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointAuthentication != nil {
		in, out := &in.EndpointAuthentication, &out.EndpointAuthentication
		*out = new(bool)
		**out = **in
	}
	if in.EndpointEncryption != nil {
		in, out := &in.EndpointEncryption, &out.EndpointEncryption
		*out = new(EndpointEncryption)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSecurity.
func (in *InfinispanSecurity) DeepCopy() *InfinispanSecurity {
	if in == nil {
		return nil
	}
	out := new(InfinispanSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceContainerSpec) DeepCopyInto(out *InfinispanServiceContainerSpec) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	in.LivenessProbe.DeepCopyInto(&out.LivenessProbe)
	in.ReadinessProbe.DeepCopyInto(&out.ReadinessProbe)
	in.StartupProbe.DeepCopyInto(&out.StartupProbe)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceContainerSpec.
func (in *InfinispanServiceContainerSpec) DeepCopy() *InfinispanServiceContainerSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanServiceSpec) DeepCopyInto(out *InfinispanServiceSpec) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(InfinispanServiceContainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Sites != nil {
		in, out := &in.Sites, &out.Sites
		*out = new(InfinispanSitesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanServiceSpec.
func (in *InfinispanServiceSpec) DeepCopy() *InfinispanServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSiteLocationSpec) DeepCopyInto(out *InfinispanSiteLocationSpec) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSiteLocationSpec.
func (in *InfinispanSiteLocationSpec) DeepCopy() *InfinispanSiteLocationSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSiteLocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesLocalSpec) DeepCopyInto(out *InfinispanSitesLocalSpec) {
	*out = *in
	in.Expose.DeepCopyInto(&out.Expose)
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionSiteSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Discovery != nil {
		in, out := &in.Discovery, &out.Discovery
		*out = new(DiscoverySiteSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesLocalSpec.
func (in *InfinispanSitesLocalSpec) DeepCopy() *InfinispanSitesLocalSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesLocalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSitesSpec) DeepCopyInto(out *InfinispanSitesSpec) {
	*out = *in
	in.Local.DeepCopyInto(&out.Local)
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]InfinispanSiteLocationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSitesSpec.
func (in *InfinispanSitesSpec) DeepCopy() *InfinispanSitesSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSitesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanSpec) DeepCopyInto(out *InfinispanSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	in.Security.DeepCopyInto(&out.Security)
	out.Container = in.Container
	in.Service.DeepCopyInto(&out.Service)
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(InfinispanLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Expose != nil {
		in, out := &in.Expose, &out.Expose
		*out = new(ExposeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscale != nil {
		in, out := &in.Autoscale, &out.Autoscale
		*out = new(Autoscale)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudEvents != nil {
		in, out := &in.CloudEvents, &out.CloudEvents
		*out = new(InfinispanCloudEvents)
		**out = **in
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = new(InfinispanExternalDependencies)
		(*in).DeepCopyInto(*out)
	}
	if in.Upgrades != nil {
		in, out := &in.Upgrades, &out.Upgrades
		*out = new(InfinispanUpgradesSpec)
		**out = **in
	}
	if in.ConfigListener != nil {
		in, out := &in.ConfigListener, &out.ConfigListener
		*out = new(ConfigListenerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Jmx != nil {
		in, out := &in.Jmx, &out.Jmx
		*out = new(JmxSpec)
		**out = **in
	}
	if in.Scheduling != nil {
		in, out := &in.Scheduling, &out.Scheduling
		*out = new(SchedulingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanSpec.
func (in *InfinispanSpec) DeepCopy() *InfinispanSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanStatus) DeepCopyInto(out *InfinispanStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]InfinispanCondition, len(*in))
		copy(*out, *in)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(InfinispanSecurity)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.PodStatus.DeepCopyInto(&out.PodStatus)
	if in.ConsoleUrl != nil {
		in, out := &in.ConsoleUrl, &out.ConsoleUrl
		*out = new(string)
		**out = **in
	}
	if in.HotRodRollingUpgradeStatus != nil {
		in, out := &in.HotRodRollingUpgradeStatus, &out.HotRodRollingUpgradeStatus
		*out = new(HotRodRollingUpgradeStatus)
		**out = **in
	}
	out.Operand = in.Operand
	out.Operator = in.Operator
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanStatus.
func (in *InfinispanStatus) DeepCopy() *InfinispanStatus {
	if in == nil {
		return nil
	}
	out := new(InfinispanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfinispanUpgradesSpec) DeepCopyInto(out *InfinispanUpgradesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfinispanUpgradesSpec.
func (in *InfinispanUpgradesSpec) DeepCopy() *InfinispanUpgradesSpec {
	if in == nil {
		return nil
	}
	out := new(InfinispanUpgradesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JmxSpec) DeepCopyInto(out *JmxSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JmxSpec.
func (in *JmxSpec) DeepCopy() *JmxSpec {
	if in == nil {
		return nil
	}
	out := new(JmxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperandStatus) DeepCopyInto(out *OperandStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperandStatus.
func (in *OperandStatus) DeepCopy() *OperandStatus {
	if in == nil {
		return nil
	}
	out := new(OperandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operator) DeepCopyInto(out *Operator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operator.
func (in *Operator) DeepCopy() *Operator {
	if in == nil {
		return nil
	}
	out := new(Operator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingSpec) DeepCopyInto(out *SchedulingSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]corev1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingSpec.
func (in *SchedulingSpec) DeepCopy() *SchedulingSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulingSpec)
	in.DeepCopyInto(out)
	return out
}
