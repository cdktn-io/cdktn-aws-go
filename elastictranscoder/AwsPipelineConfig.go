package elastictranscoder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#input_bucket AwsPipeline#input_bucket}.
	// Experimental.
	InputBucket *string `field:"required" json:"inputBucket" yaml:"inputBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#role AwsPipeline#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#aws_kms_key_arn AwsPipeline#aws_kms_key_arn}.
	// Experimental.
	AwsKmsKeyArn *string `field:"optional" json:"awsKmsKeyArn" yaml:"awsKmsKeyArn"`
	// content_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#content_config AwsPipeline#content_config}
	// Experimental.
	ContentConfig *AwsPipeline_ContentConfigProperty `field:"optional" json:"contentConfig" yaml:"contentConfig"`
	// content_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#content_config_permissions AwsPipeline#content_config_permissions}
	// Experimental.
	ContentConfigPermissions interface{} `field:"optional" json:"contentConfigPermissions" yaml:"contentConfigPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#id AwsPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#name AwsPipeline#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#notifications AwsPipeline#notifications}
	// Experimental.
	Notifications *AwsPipeline_NotificationsProperty `field:"optional" json:"notifications" yaml:"notifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#output_bucket AwsPipeline#output_bucket}.
	// Experimental.
	OutputBucket *string `field:"optional" json:"outputBucket" yaml:"outputBucket"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#region AwsPipeline#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// thumbnail_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#thumbnail_config AwsPipeline#thumbnail_config}
	// Experimental.
	ThumbnailConfig *AwsPipeline_ThumbnailConfigProperty `field:"optional" json:"thumbnailConfig" yaml:"thumbnailConfig"`
	// thumbnail_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elastictranscoder_pipeline#thumbnail_config_permissions AwsPipeline#thumbnail_config_permissions}
	// Experimental.
	ThumbnailConfigPermissions interface{} `field:"optional" json:"thumbnailConfigPermissions" yaml:"thumbnailConfigPermissions"`
}

