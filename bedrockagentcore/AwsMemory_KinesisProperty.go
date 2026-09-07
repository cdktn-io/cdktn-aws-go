package bedrockagentcore


// Experimental.
type AwsMemory_KinesisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#data_stream_arn AwsMemory#data_stream_arn}.
	// Experimental.
	DataStreamArn *string `field:"required" json:"dataStreamArn" yaml:"dataStreamArn"`
	// content_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#content_configuration AwsMemory#content_configuration}
	// Experimental.
	ContentConfiguration interface{} `field:"optional" json:"contentConfiguration" yaml:"contentConfiguration"`
}

