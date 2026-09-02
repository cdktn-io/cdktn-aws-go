package awsdatapipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipelineDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#pipeline_id TfPipelineDefinition#pipeline_id}.
	// Experimental.
	PipelineId *string `field:"required" json:"pipelineId" yaml:"pipelineId"`
	// pipeline_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#pipeline_object TfPipelineDefinition#pipeline_object}
	// Experimental.
	PipelineObject interface{} `field:"required" json:"pipelineObject" yaml:"pipelineObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#id TfPipelineDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// parameter_object block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#parameter_object TfPipelineDefinition#parameter_object}
	// Experimental.
	ParameterObject interface{} `field:"optional" json:"parameterObject" yaml:"parameterObject"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#parameter_value TfPipelineDefinition#parameter_value}
	// Experimental.
	ParameterValue interface{} `field:"optional" json:"parameterValue" yaml:"parameterValue"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#region TfPipelineDefinition#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

