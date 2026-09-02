package awsbedrockagentcore


// Experimental.
type TfGatewayRule_ActionConfigurationBundleWeightedOverrideTrafficSplitConfigurationBundleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_arn TfGatewayRule#bundle_arn}.
	// Experimental.
	BundleArn *string `field:"required" json:"bundleArn" yaml:"bundleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_gateway_rule#bundle_version TfGatewayRule#bundle_version}.
	// Experimental.
	BundleVersion *string `field:"required" json:"bundleVersion" yaml:"bundleVersion"`
}

