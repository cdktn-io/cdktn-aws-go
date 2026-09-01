package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreGatewayRule_ActionConfigurationBundleWeightedOverrideTrafficSplitConfigurationBundleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_arn AwsBedrockagentcoreGatewayRule#bundle_arn}.
	// Experimental.
	BundleArn *string `field:"required" json:"bundleArn" yaml:"bundleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_version AwsBedrockagentcoreGatewayRule#bundle_version}.
	// Experimental.
	BundleVersion *string `field:"required" json:"bundleVersion" yaml:"bundleVersion"`
}

