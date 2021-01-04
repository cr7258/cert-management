/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gardener/cert-management/pkg/apis/cert/v1alpha1"
	"github.com/gardener/cert-management/pkg/client/cert/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CertV1alpha1Interface interface {
	RESTClient() rest.Interface
	CertificatesGetter
	CertificateRevocationsGetter
	IssuersGetter
}

// CertV1alpha1Client is used to interact with features provided by the cert.gardener.cloud group.
type CertV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CertV1alpha1Client) Certificates(namespace string) CertificateInterface {
	return newCertificates(c, namespace)
}

func (c *CertV1alpha1Client) CertificateRevocations(namespace string) CertificateRevocationInterface {
	return newCertificateRevocations(c, namespace)
}

func (c *CertV1alpha1Client) Issuers(namespace string) IssuerInterface {
	return newIssuers(c, namespace)
}

// NewForConfig creates a new CertV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CertV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CertV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CertV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CertV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CertV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CertV1alpha1Client {
	return &CertV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CertV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
