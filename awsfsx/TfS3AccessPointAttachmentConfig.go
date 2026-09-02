package awsfsx

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfS3AccessPointAttachmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#name TfS3AccessPointAttachment#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#type TfS3AccessPointAttachment#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// openzfs_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#openzfs_configuration TfS3AccessPointAttachment#openzfs_configuration}
	// Experimental.
	OpenzfsConfiguration interface{} `field:"optional" json:"openzfsConfiguration" yaml:"openzfsConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#region TfS3AccessPointAttachment#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_access_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#s3_access_point TfS3AccessPointAttachment#s3_access_point}
	// Experimental.
	S3AccessPoint interface{} `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_s3_access_point_attachment#timeouts TfS3AccessPointAttachment#timeouts}
	// Experimental.
	Timeouts *TfS3AccessPointAttachment_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

