package awsdatapipeline


// Experimental.
type AwsDatapipelinePipelineDefinition_PipelineObjectProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#id AwsDatapipelinePipelineDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#name AwsDatapipelinePipelineDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#field AwsDatapipelinePipelineDefinition#field}
	// Experimental.
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}

