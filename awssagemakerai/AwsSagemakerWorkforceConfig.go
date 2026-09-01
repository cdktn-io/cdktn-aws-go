package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerWorkforceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#workforce_name AwsSagemakerWorkforce#workforce_name}.
	// Experimental.
	WorkforceName *string `field:"required" json:"workforceName" yaml:"workforceName"`
	// cognito_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#cognito_config AwsSagemakerWorkforce#cognito_config}
	// Experimental.
	CognitoConfig *AwsSagemakerWorkforce_CognitoConfigProperty `field:"optional" json:"cognitoConfig" yaml:"cognitoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#id AwsSagemakerWorkforce#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// oidc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#oidc_config AwsSagemakerWorkforce#oidc_config}
	// Experimental.
	OidcConfig *AwsSagemakerWorkforce_OidcConfigProperty `field:"optional" json:"oidcConfig" yaml:"oidcConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#region AwsSagemakerWorkforce#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// source_ip_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#source_ip_config AwsSagemakerWorkforce#source_ip_config}
	// Experimental.
	SourceIpConfig *AwsSagemakerWorkforce_SourceIpConfigProperty `field:"optional" json:"sourceIpConfig" yaml:"sourceIpConfig"`
	// workforce_vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_workforce#workforce_vpc_config AwsSagemakerWorkforce#workforce_vpc_config}
	// Experimental.
	WorkforceVpcConfig *AwsSagemakerWorkforce_WorkforceVpcConfigProperty `field:"optional" json:"workforceVpcConfig" yaml:"workforceVpcConfig"`
}

