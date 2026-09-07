package bedrockagentcore


// Experimental.
type AwsGatewayRule_StaticOverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_arn AwsGatewayRule#bundle_arn}.
	// Experimental.
	BundleArn *string `field:"required" json:"bundleArn" yaml:"bundleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_version AwsGatewayRule#bundle_version}.
	// Experimental.
	BundleVersion *string `field:"required" json:"bundleVersion" yaml:"bundleVersion"`
}

