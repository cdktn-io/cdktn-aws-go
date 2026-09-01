package awsopensearchserverless

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchserverlessSecurityConfigConfig struct {
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
	// Name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#name AwsOpensearchserverlessSecurityConfig#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of configuration. Valid values: `saml`, `iamidentitycenter` or `iamfederation`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#type AwsOpensearchserverlessSecurityConfig#type}
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Description of the security configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#description AwsOpensearchserverlessSecurityConfig#description}
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// iam_federation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#iam_federation_options AwsOpensearchserverlessSecurityConfig#iam_federation_options}
	// Experimental.
	IamFederationOptions interface{} `field:"optional" json:"iamFederationOptions" yaml:"iamFederationOptions"`
	// iam_identity_center_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#iam_identity_center_options AwsOpensearchserverlessSecurityConfig#iam_identity_center_options}
	// Experimental.
	IamIdentityCenterOptions interface{} `field:"optional" json:"iamIdentityCenterOptions" yaml:"iamIdentityCenterOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#region AwsOpensearchserverlessSecurityConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// saml_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearchserverless_security_config#saml_options AwsOpensearchserverlessSecurityConfig#saml_options}
	// Experimental.
	SamlOptions interface{} `field:"optional" json:"samlOptions" yaml:"samlOptions"`
}

