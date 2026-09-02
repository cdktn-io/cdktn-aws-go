package awsroute53


// Experimental.
type TfRecordsExclusive_GeoproximityLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#aws_region TfRecordsExclusive#aws_region}.
	// Experimental.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#bias TfRecordsExclusive#bias}.
	// Experimental.
	Bias *float64 `field:"optional" json:"bias" yaml:"bias"`
	// coordinates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#coordinates TfRecordsExclusive#coordinates}
	// Experimental.
	Coordinates interface{} `field:"optional" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#local_zone_group TfRecordsExclusive#local_zone_group}.
	// Experimental.
	LocalZoneGroup *string `field:"optional" json:"localZoneGroup" yaml:"localZoneGroup"`
}

