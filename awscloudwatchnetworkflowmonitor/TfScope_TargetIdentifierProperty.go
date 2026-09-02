package awscloudwatchnetworkflowmonitor


// Experimental.
type TfScope_TargetIdentifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_scope#target_type TfScope#target_type}.
	// Experimental.
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// target_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/networkflowmonitor_scope#target_id TfScope#target_id}
	// Experimental.
	TargetId interface{} `field:"optional" json:"targetId" yaml:"targetId"`
}

