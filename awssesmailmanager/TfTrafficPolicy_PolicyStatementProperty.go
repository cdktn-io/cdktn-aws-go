package awssesmailmanager


// Experimental.
type TfTrafficPolicy_PolicyStatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#action TfTrafficPolicy#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#condition TfTrafficPolicy#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
}

