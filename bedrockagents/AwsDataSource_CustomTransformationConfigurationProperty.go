package bedrockagents


// Experimental.
type AwsDataSource_CustomTransformationConfigurationProperty struct {
	// intermediate_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#intermediate_storage AwsDataSource#intermediate_storage}
	// Experimental.
	IntermediateStorage interface{} `field:"optional" json:"intermediateStorage" yaml:"intermediateStorage"`
	// transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#transformation AwsDataSource#transformation}
	// Experimental.
	Transformation interface{} `field:"optional" json:"transformation" yaml:"transformation"`
}

