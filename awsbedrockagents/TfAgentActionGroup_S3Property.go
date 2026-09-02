package awsbedrockagents


// Experimental.
type TfAgentActionGroup_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#s3_bucket_name TfAgentActionGroup#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#s3_object_key TfAgentActionGroup#s3_object_key}.
	// Experimental.
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

