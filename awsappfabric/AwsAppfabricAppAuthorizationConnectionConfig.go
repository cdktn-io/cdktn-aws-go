package awsappfabric

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppfabricAppAuthorizationConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#app_authorization_arn AwsAppfabricAppAuthorizationConnection#app_authorization_arn}.
	// Experimental.
	AppAuthorizationArn *string `field:"required" json:"appAuthorizationArn" yaml:"appAuthorizationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#app_bundle_arn AwsAppfabricAppAuthorizationConnection#app_bundle_arn}.
	// Experimental.
	AppBundleArn *string `field:"required" json:"appBundleArn" yaml:"appBundleArn"`
	// auth_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#auth_request AwsAppfabricAppAuthorizationConnection#auth_request}
	// Experimental.
	AuthRequest interface{} `field:"optional" json:"authRequest" yaml:"authRequest"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#region AwsAppfabricAppAuthorizationConnection#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_app_authorization_connection#timeouts AwsAppfabricAppAuthorizationConnection#timeouts}
	// Experimental.
	Timeouts *AwsAppfabricAppAuthorizationConnection_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

