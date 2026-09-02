package awsverifiedaccess

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTrustProviderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#policy_reference_name TfTrustProvider#policy_reference_name}.
	// Experimental.
	PolicyReferenceName *string `field:"required" json:"policyReferenceName" yaml:"policyReferenceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#trust_provider_type TfTrustProvider#trust_provider_type}.
	// Experimental.
	TrustProviderType *string `field:"required" json:"trustProviderType" yaml:"trustProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#description TfTrustProvider#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// device_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#device_options TfTrustProvider#device_options}
	// Experimental.
	DeviceOptions *TfTrustProvider_DeviceOptionsProperty `field:"optional" json:"deviceOptions" yaml:"deviceOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#device_trust_provider_type TfTrustProvider#device_trust_provider_type}.
	// Experimental.
	DeviceTrustProviderType *string `field:"optional" json:"deviceTrustProviderType" yaml:"deviceTrustProviderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#id TfTrustProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// native_application_oidc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#native_application_oidc_options TfTrustProvider#native_application_oidc_options}
	// Experimental.
	NativeApplicationOidcOptions *TfTrustProvider_NativeApplicationOidcOptionsProperty `field:"optional" json:"nativeApplicationOidcOptions" yaml:"nativeApplicationOidcOptions"`
	// oidc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#oidc_options TfTrustProvider#oidc_options}
	// Experimental.
	OidcOptions *TfTrustProvider_OidcOptionsProperty `field:"optional" json:"oidcOptions" yaml:"oidcOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#region TfTrustProvider#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// sse_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#sse_specification TfTrustProvider#sse_specification}
	// Experimental.
	SseSpecification *TfTrustProvider_SseSpecificationProperty `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#tags TfTrustProvider#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#tags_all TfTrustProvider#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#timeouts TfTrustProvider#timeouts}
	// Experimental.
	Timeouts *TfTrustProvider_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_trust_provider#user_trust_provider_type TfTrustProvider#user_trust_provider_type}.
	// Experimental.
	UserTrustProviderType *string `field:"optional" json:"userTrustProviderType" yaml:"userTrustProviderType"`
}

