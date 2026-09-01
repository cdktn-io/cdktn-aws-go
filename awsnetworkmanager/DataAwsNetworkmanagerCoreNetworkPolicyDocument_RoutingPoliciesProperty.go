package awsnetworkmanager


// Experimental.
type DataAwsNetworkmanagerCoreNetworkPolicyDocument_RoutingPoliciesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_direction DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_direction}.
	// Experimental.
	RoutingPolicyDirection *string `field:"required" json:"routingPolicyDirection" yaml:"routingPolicyDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_name DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_name}.
	// Experimental.
	RoutingPolicyName *string `field:"required" json:"routingPolicyName" yaml:"routingPolicyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_number DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_number}.
	// Experimental.
	RoutingPolicyNumber *float64 `field:"required" json:"routingPolicyNumber" yaml:"routingPolicyNumber"`
	// routing_policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_rules DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_rules}
	// Experimental.
	RoutingPolicyRules interface{} `field:"required" json:"routingPolicyRules" yaml:"routingPolicyRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document#routing_policy_description DataAwsNetworkmanagerCoreNetworkPolicyDocument#routing_policy_description}.
	// Experimental.
	RoutingPolicyDescription *string `field:"optional" json:"routingPolicyDescription" yaml:"routingPolicyDescription"`
}

