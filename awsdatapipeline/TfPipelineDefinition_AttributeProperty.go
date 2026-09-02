package awsdatapipeline


// Experimental.
type TfPipelineDefinition_AttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#key TfPipelineDefinition#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#string_value TfPipelineDefinition#string_value}.
	// Experimental.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

