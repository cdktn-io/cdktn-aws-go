package datapipeline


// Experimental.
type AwsPipelineDefinition_AttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#key AwsPipelineDefinition#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datapipeline_pipeline_definition#string_value AwsPipelineDefinition#string_value}.
	// Experimental.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

