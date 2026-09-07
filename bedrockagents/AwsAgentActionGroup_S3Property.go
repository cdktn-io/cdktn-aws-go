package bedrockagents


// Experimental.
type AwsAgentActionGroup_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#s3_bucket_name AwsAgentActionGroup#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#s3_object_key AwsAgentActionGroup#s3_object_key}.
	// Experimental.
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

