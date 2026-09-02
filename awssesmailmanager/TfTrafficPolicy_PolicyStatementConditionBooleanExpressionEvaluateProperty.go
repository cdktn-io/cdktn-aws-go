package awssesmailmanager


// Experimental.
type TfTrafficPolicy_PolicyStatementConditionBooleanExpressionEvaluateProperty struct {
	// analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#analysis TfTrafficPolicy#analysis}
	// Experimental.
	Analysis interface{} `field:"optional" json:"analysis" yaml:"analysis"`
	// is_in_address_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#is_in_address_list TfTrafficPolicy#is_in_address_list}
	// Experimental.
	IsInAddressList interface{} `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

