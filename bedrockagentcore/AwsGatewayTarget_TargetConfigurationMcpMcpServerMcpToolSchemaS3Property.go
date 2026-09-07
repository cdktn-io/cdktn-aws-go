package bedrockagentcore


// Experimental.
type AwsGatewayTarget_TargetConfigurationMcpMcpServerMcpToolSchemaS3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#uri AwsGatewayTarget#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_target#bucket_owner_account_id AwsGatewayTarget#bucket_owner_account_id}.
	// Experimental.
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
}

