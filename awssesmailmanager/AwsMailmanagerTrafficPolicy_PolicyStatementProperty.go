package awssesmailmanager


// Experimental.
type AwsMailmanagerTrafficPolicy_PolicyStatementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#action AwsMailmanagerTrafficPolicy#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_traffic_policy#condition AwsMailmanagerTrafficPolicy#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
}

