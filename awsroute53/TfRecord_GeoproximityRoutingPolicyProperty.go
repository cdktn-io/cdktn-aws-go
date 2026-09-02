package awsroute53


// Experimental.
type TfRecord_GeoproximityRoutingPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#aws_region TfRecord#aws_region}.
	// Experimental.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#bias TfRecord#bias}.
	// Experimental.
	Bias *float64 `field:"optional" json:"bias" yaml:"bias"`
	// coordinates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#coordinates TfRecord#coordinates}
	// Experimental.
	Coordinates interface{} `field:"optional" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#local_zone_group TfRecord#local_zone_group}.
	// Experimental.
	LocalZoneGroup *string `field:"optional" json:"localZoneGroup" yaml:"localZoneGroup"`
}

