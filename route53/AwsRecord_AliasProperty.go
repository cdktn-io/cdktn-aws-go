package route53


// Experimental.
type AwsRecord_AliasProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#evaluate_target_health AwsRecord#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"required" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#name AwsRecord#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#zone_id AwsRecord#zone_id}.
	// Experimental.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

