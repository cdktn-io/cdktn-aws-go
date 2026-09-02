package awsroute53


// Experimental.
type DataTfTrafficPolicyDocument_LocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#continent DataTfTrafficPolicyDocument#continent}.
	// Experimental.
	Continent *string `field:"optional" json:"continent" yaml:"continent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#country DataTfTrafficPolicyDocument#country}.
	// Experimental.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataTfTrafficPolicyDocument#endpoint_reference}.
	// Experimental.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#evaluate_target_health DataTfTrafficPolicyDocument#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#health_check DataTfTrafficPolicyDocument#health_check}.
	// Experimental.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#is_default DataTfTrafficPolicyDocument#is_default}.
	// Experimental.
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#rule_reference DataTfTrafficPolicyDocument#rule_reference}.
	// Experimental.
	RuleReference *string `field:"optional" json:"ruleReference" yaml:"ruleReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#subdivision DataTfTrafficPolicyDocument#subdivision}.
	// Experimental.
	Subdivision *string `field:"optional" json:"subdivision" yaml:"subdivision"`
}

