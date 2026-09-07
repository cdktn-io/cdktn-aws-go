package route53


// Experimental.
type DataAwsTrafficPolicyDocument_GeoProximityLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#bias DataAwsTrafficPolicyDocument#bias}.
	// Experimental.
	Bias *string `field:"optional" json:"bias" yaml:"bias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#endpoint_reference DataAwsTrafficPolicyDocument#endpoint_reference}.
	// Experimental.
	EndpointReference *string `field:"optional" json:"endpointReference" yaml:"endpointReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#evaluate_target_health DataAwsTrafficPolicyDocument#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"optional" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#health_check DataAwsTrafficPolicyDocument#health_check}.
	// Experimental.
	HealthCheck *string `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#latitude DataAwsTrafficPolicyDocument#latitude}.
	// Experimental.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#longitude DataAwsTrafficPolicyDocument#longitude}.
	// Experimental.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#region DataAwsTrafficPolicyDocument#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#rule_reference DataAwsTrafficPolicyDocument#rule_reference}.
	// Experimental.
	RuleReference *string `field:"optional" json:"ruleReference" yaml:"ruleReference"`
}

