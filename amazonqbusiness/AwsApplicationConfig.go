package amazonqbusiness

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplicationConfig struct {
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
	// The display name of the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#display_name AwsApplication#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Amazon Resource Name (ARN) of the IAM service role that provides permissions for the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#iam_service_role_arn AwsApplication#iam_service_role_arn}
	// Experimental.
	IamServiceRoleArn *string `field:"required" json:"iamServiceRoleArn" yaml:"iamServiceRoleArn"`
	// ARN of the IAM Identity Center instance you are either creating for—or connecting to—your Amazon Q Business application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#identity_center_instance_arn AwsApplication#identity_center_instance_arn}
	// Experimental.
	IdentityCenterInstanceArn *string `field:"required" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// attachments_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#attachments_configuration AwsApplication#attachments_configuration}
	// Experimental.
	AttachmentsConfiguration interface{} `field:"optional" json:"attachmentsConfiguration" yaml:"attachmentsConfiguration"`
	// A description of the Amazon Q application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#description AwsApplication#description}
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#encryption_configuration AwsApplication#encryption_configuration}
	// Experimental.
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#region AwsApplication#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#tags AwsApplication#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/qbusiness_application#timeouts AwsApplication#timeouts}
	// Experimental.
	Timeouts *AwsApplication_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

