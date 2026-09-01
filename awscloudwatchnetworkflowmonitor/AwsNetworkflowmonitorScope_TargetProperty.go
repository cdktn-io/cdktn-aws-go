package awscloudwatchnetworkflowmonitor


// Experimental.
type AwsNetworkflowmonitorScope_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_scope#region AwsNetworkflowmonitorScope#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// target_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_scope#target_identifier AwsNetworkflowmonitorScope#target_identifier}
	// Experimental.
	TargetIdentifier interface{} `field:"optional" json:"targetIdentifier" yaml:"targetIdentifier"`
}

