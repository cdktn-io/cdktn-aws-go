package awsroute53


// Experimental.
type TfRecord_AliasProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#evaluate_target_health TfRecord#evaluate_target_health}.
	// Experimental.
	EvaluateTargetHealth interface{} `field:"required" json:"evaluateTargetHealth" yaml:"evaluateTargetHealth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#name TfRecord#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#zone_id TfRecord#zone_id}.
	// Experimental.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
}

