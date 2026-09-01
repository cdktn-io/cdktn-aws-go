package awsdatapipeline


// Experimental.
type AwsDatapipelinePipelineDefinition_FieldProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#key AwsDatapipelinePipelineDefinition#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#ref_value AwsDatapipelinePipelineDefinition#ref_value}.
	// Experimental.
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#string_value AwsDatapipelinePipelineDefinition#string_value}.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

