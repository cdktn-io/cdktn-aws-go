package awsosis

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#max_units TfPipeline#max_units}.
	// Experimental.
	MaxUnits *float64 `field:"required" json:"maxUnits" yaml:"maxUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#min_units TfPipeline#min_units}.
	// Experimental.
	MinUnits *float64 `field:"required" json:"minUnits" yaml:"minUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#pipeline_configuration_body TfPipeline#pipeline_configuration_body}.
	// Experimental.
	PipelineConfigurationBody *string `field:"required" json:"pipelineConfigurationBody" yaml:"pipelineConfigurationBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#pipeline_name TfPipeline#pipeline_name}.
	// Experimental.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// buffer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#buffer_options TfPipeline#buffer_options}
	// Experimental.
	BufferOptions interface{} `field:"optional" json:"bufferOptions" yaml:"bufferOptions"`
	// encryption_at_rest_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#encryption_at_rest_options TfPipeline#encryption_at_rest_options}
	// Experimental.
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#log_publishing_options TfPipeline#log_publishing_options}
	// Experimental.
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#pipeline_role_arn TfPipeline#pipeline_role_arn}.
	// Experimental.
	PipelineRoleArn *string `field:"optional" json:"pipelineRoleArn" yaml:"pipelineRoleArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#region TfPipeline#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#tags TfPipeline#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#timeouts TfPipeline#timeouts}
	// Experimental.
	Timeouts *TfPipeline_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/osis_pipeline#vpc_options TfPipeline#vpc_options}
	// Experimental.
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

