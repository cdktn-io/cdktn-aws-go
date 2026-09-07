package route53


// Experimental.
type DataAwsTrafficPolicyDocument_SecondaryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataAwsTrafficPolicyDocument#endpoint_reference}.
	// Experimental.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#evaluate_target_health DataAwsTrafficPolicyDocument#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#health_check DataAwsTrafficPolicyDocument#health_check}.
	// Experimental.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#rule_reference DataAwsTrafficPolicyDocument#rule_reference}.
	// Experimental.
	RuleReference *string `field:"optional" json:"ruleReference" yaml:"ruleReference"`
}

