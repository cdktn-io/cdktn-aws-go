package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreMemory_KinesisProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#data_stream_arn AwsBedrockagentcoreMemory#data_stream_arn}.
	// Experimental.
	DataStreamArn *string `field:"required" json:"dataStreamArn" yaml:"dataStreamArn"`
	// content_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory#content_configuration AwsBedrockagentcoreMemory#content_configuration}
	// Experimental.
	ContentConfiguration interface{} `field:"optional" json:"contentConfiguration" yaml:"contentConfiguration"`
}

