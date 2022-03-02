//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMEExternalAccountBinding) DeepCopyInto(out *ACMEExternalAccountBinding) {
	*out = *in
	if in.KeySecretRef != nil {
		in, out := &in.KeySecretRef, &out.KeySecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMEExternalAccountBinding.
func (in *ACMEExternalAccountBinding) DeepCopy() *ACMEExternalAccountBinding {
	if in == nil {
		return nil
	}
	out := new(ACMEExternalAccountBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACMESpec) DeepCopyInto(out *ACMESpec) {
	*out = *in
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.ExternalAccountBinding != nil {
		in, out := &in.ExternalAccountBinding, &out.ExternalAccountBinding
		*out = new(ACMEExternalAccountBinding)
		(*in).DeepCopyInto(*out)
	}
	if in.SkipDNSChallengeValidation != nil {
		in, out := &in.SkipDNSChallengeValidation, &out.SkipDNSChallengeValidation
		*out = new(bool)
		**out = **in
	}
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = new(DNSSelection)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACMESpec.
func (in *ACMESpec) DeepCopy() *ACMESpec {
	if in == nil {
		return nil
	}
	out := new(ACMESpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackOffState) DeepCopyInto(out *BackOffState) {
	*out = *in
	in.RetryAfter.DeepCopyInto(&out.RetryAfter)
	out.RetryInterval = in.RetryInterval
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackOffState.
func (in *BackOffState) DeepCopy() *BackOffState {
	if in == nil {
		return nil
	}
	out := new(BackOffState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CASpec) DeepCopyInto(out *CASpec) {
	*out = *in
	if in.PrivateKeySecretRef != nil {
		in, out := &in.PrivateKeySecretRef, &out.PrivateKeySecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CASpec.
func (in *CASpec) DeepCopy() *CASpec {
	if in == nil {
		return nil
	}
	out := new(CASpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Certificate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateList) DeepCopyInto(out *CertificateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Certificate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateList.
func (in *CertificateList) DeepCopy() *CertificateList {
	if in == nil {
		return nil
	}
	out := new(CertificateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRef) DeepCopyInto(out *CertificateRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRef.
func (in *CertificateRef) DeepCopy() *CertificateRef {
	if in == nil {
		return nil
	}
	out := new(CertificateRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRevocation) DeepCopyInto(out *CertificateRevocation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRevocation.
func (in *CertificateRevocation) DeepCopy() *CertificateRevocation {
	if in == nil {
		return nil
	}
	out := new(CertificateRevocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRevocation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRevocationList) DeepCopyInto(out *CertificateRevocationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertificateRevocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRevocationList.
func (in *CertificateRevocationList) DeepCopy() *CertificateRevocationList {
	if in == nil {
		return nil
	}
	out := new(CertificateRevocationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertificateRevocationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRevocationSpec) DeepCopyInto(out *CertificateRevocationSpec) {
	*out = *in
	out.CertificateRef = in.CertificateRef
	if in.Renew != nil {
		in, out := &in.Renew, &out.Renew
		*out = new(bool)
		**out = **in
	}
	if in.QualifyingDate != nil {
		in, out := &in.QualifyingDate, &out.QualifyingDate
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRevocationSpec.
func (in *CertificateRevocationSpec) DeepCopy() *CertificateRevocationSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateRevocationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateRevocationStatus) DeepCopyInto(out *CertificateRevocationStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = new(ObjectStatuses)
		(*in).DeepCopyInto(*out)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SecretStatuses)
		(*in).DeepCopyInto(*out)
	}
	if in.RevocationApplied != nil {
		in, out := &in.RevocationApplied, &out.RevocationApplied
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateRevocationStatus.
func (in *CertificateRevocationStatus) DeepCopy() *CertificateRevocationStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateRevocationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSecretRef) DeepCopyInto(out *CertificateSecretRef) {
	*out = *in
	out.SecretReference = in.SecretReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSecretRef.
func (in *CertificateSecretRef) DeepCopy() *CertificateSecretRef {
	if in == nil {
		return nil
	}
	out := new(CertificateSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSpec) DeepCopyInto(out *CertificateSpec) {
	*out = *in
	if in.CommonName != nil {
		in, out := &in.CommonName, &out.CommonName
		*out = new(string)
		**out = **in
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CSR != nil {
		in, out := &in.CSR, &out.CSR
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(IssuerRef)
		**out = **in
	}
	if in.SecretName != nil {
		in, out := &in.SecretName, &out.SecretName
		*out = new(string)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Renew != nil {
		in, out := &in.Renew, &out.Renew
		*out = new(bool)
		**out = **in
	}
	if in.EnsureRenewedAfter != nil {
		in, out := &in.EnsureRenewedAfter, &out.EnsureRenewedAfter
		*out = (*in).DeepCopy()
	}
	if in.FollowCNAME != nil {
		in, out := &in.FollowCNAME, &out.FollowCNAME
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSpec.
func (in *CertificateSpec) DeepCopy() *CertificateSpec {
	if in == nil {
		return nil
	}
	out := new(CertificateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateStatus) DeepCopyInto(out *CertificateStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.LastPendingTimestamp != nil {
		in, out := &in.LastPendingTimestamp, &out.LastPendingTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CommonName != nil {
		in, out := &in.CommonName, &out.CommonName
		*out = new(string)
		**out = **in
	}
	if in.DNSNames != nil {
		in, out := &in.DNSNames, &out.DNSNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IssuerRef != nil {
		in, out := &in.IssuerRef, &out.IssuerRef
		*out = new(QualifiedIssuerRef)
		**out = **in
	}
	if in.ExpirationDate != nil {
		in, out := &in.ExpirationDate, &out.ExpirationDate
		*out = new(string)
		**out = **in
	}
	if in.BackOff != nil {
		in, out := &in.BackOff, &out.BackOff
		*out = new(BackOffState)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateStatus.
func (in *CertificateStatus) DeepCopy() *CertificateStatus {
	if in == nil {
		return nil
	}
	out := new(CertificateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSSelection) DeepCopyInto(out *DNSSelection) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSSelection.
func (in *DNSSelection) DeepCopy() *DNSSelection {
	if in == nil {
		return nil
	}
	out := new(DNSSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Issuer) DeepCopyInto(out *Issuer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Issuer.
func (in *Issuer) DeepCopy() *Issuer {
	if in == nil {
		return nil
	}
	out := new(Issuer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Issuer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerList) DeepCopyInto(out *IssuerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Issuer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerList.
func (in *IssuerList) DeepCopy() *IssuerList {
	if in == nil {
		return nil
	}
	out := new(IssuerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IssuerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerRef) DeepCopyInto(out *IssuerRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerRef.
func (in *IssuerRef) DeepCopy() *IssuerRef {
	if in == nil {
		return nil
	}
	out := new(IssuerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerSpec) DeepCopyInto(out *IssuerSpec) {
	*out = *in
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		*out = new(ACMESpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(CASpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RequestsPerDayQuota != nil {
		in, out := &in.RequestsPerDayQuota, &out.RequestsPerDayQuota
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerSpec.
func (in *IssuerSpec) DeepCopy() *IssuerSpec {
	if in == nil {
		return nil
	}
	out := new(IssuerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssuerStatus) DeepCopyInto(out *IssuerStatus) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ACME != nil {
		in, out := &in.ACME, &out.ACME
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.CA != nil {
		in, out := &in.CA, &out.CA
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssuerStatus.
func (in *IssuerStatus) DeepCopy() *IssuerStatus {
	if in == nil {
		return nil
	}
	out := new(IssuerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStatuses) DeepCopyInto(out *ObjectStatuses) {
	*out = *in
	if in.Processing != nil {
		in, out := &in.Processing, &out.Processing
		*out = make([]CertificateRef, len(*in))
		copy(*out, *in)
	}
	if in.Renewed != nil {
		in, out := &in.Renewed, &out.Renewed
		*out = make([]CertificateRef, len(*in))
		copy(*out, *in)
	}
	if in.Revoked != nil {
		in, out := &in.Revoked, &out.Revoked
		*out = make([]CertificateRef, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]CertificateRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStatuses.
func (in *ObjectStatuses) DeepCopy() *ObjectStatuses {
	if in == nil {
		return nil
	}
	out := new(ObjectStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QualifiedIssuerRef) DeepCopyInto(out *QualifiedIssuerRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QualifiedIssuerRef.
func (in *QualifiedIssuerRef) DeepCopy() *QualifiedIssuerRef {
	if in == nil {
		return nil
	}
	out := new(QualifiedIssuerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatuses) DeepCopyInto(out *SecretStatuses) {
	*out = *in
	if in.Processing != nil {
		in, out := &in.Processing, &out.Processing
		*out = make([]CertificateSecretRef, len(*in))
		copy(*out, *in)
	}
	if in.Revoked != nil {
		in, out := &in.Revoked, &out.Revoked
		*out = make([]CertificateSecretRef, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]CertificateSecretRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatuses.
func (in *SecretStatuses) DeepCopy() *SecretStatuses {
	if in == nil {
		return nil
	}
	out := new(SecretStatuses)
	in.DeepCopyInto(out)
	return out
}
