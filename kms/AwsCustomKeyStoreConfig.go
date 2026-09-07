package kms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCustomKeyStoreConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#custom_key_store_name AwsCustomKeyStore#custom_key_store_name}.
	// Experimental.
	CustomKeyStoreName *string `field:"required" json:"customKeyStoreName" yaml:"customKeyStoreName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#cloud_hsm_cluster_id AwsCustomKeyStore#cloud_hsm_cluster_id}.
	// Experimental.
	CloudHsmClusterId *string `field:"optional" json:"cloudHsmClusterId" yaml:"cloudHsmClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#custom_key_store_type AwsCustomKeyStore#custom_key_store_type}.
	// Experimental.
	CustomKeyStoreType *string `field:"optional" json:"customKeyStoreType" yaml:"customKeyStoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#id AwsCustomKeyStore#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#key_store_password AwsCustomKeyStore#key_store_password}.
	// Experimental.
	KeyStorePassword *string `field:"optional" json:"keyStorePassword" yaml:"keyStorePassword"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#region AwsCustomKeyStore#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#timeouts AwsCustomKeyStore#timeouts}
	// Experimental.
	Timeouts *AwsCustomKeyStore_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#trust_anchor_certificate AwsCustomKeyStore#trust_anchor_certificate}.
	// Experimental.
	TrustAnchorCertificate *string `field:"optional" json:"trustAnchorCertificate" yaml:"trustAnchorCertificate"`
	// xks_proxy_authentication_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#xks_proxy_authentication_credential AwsCustomKeyStore#xks_proxy_authentication_credential}
	// Experimental.
	XksProxyAuthenticationCredential *AwsCustomKeyStore_XksProxyAuthenticationCredentialProperty `field:"optional" json:"xksProxyAuthenticationCredential" yaml:"xksProxyAuthenticationCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#xks_proxy_connectivity AwsCustomKeyStore#xks_proxy_connectivity}.
	// Experimental.
	XksProxyConnectivity *string `field:"optional" json:"xksProxyConnectivity" yaml:"xksProxyConnectivity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#xks_proxy_uri_endpoint AwsCustomKeyStore#xks_proxy_uri_endpoint}.
	// Experimental.
	XksProxyUriEndpoint *string `field:"optional" json:"xksProxyUriEndpoint" yaml:"xksProxyUriEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#xks_proxy_uri_path AwsCustomKeyStore#xks_proxy_uri_path}.
	// Experimental.
	XksProxyUriPath *string `field:"optional" json:"xksProxyUriPath" yaml:"xksProxyUriPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#xks_proxy_vpc_endpoint_service_name AwsCustomKeyStore#xks_proxy_vpc_endpoint_service_name}.
	// Experimental.
	XksProxyVpcEndpointServiceName *string `field:"optional" json:"xksProxyVpcEndpointServiceName" yaml:"xksProxyVpcEndpointServiceName"`
}

