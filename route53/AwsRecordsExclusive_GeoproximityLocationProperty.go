package route53


// Experimental.
type AwsRecordsExclusive_GeoproximityLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#aws_region AwsRecordsExclusive#aws_region}.
	// Experimental.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#bias AwsRecordsExclusive#bias}.
	// Experimental.
	Bias *float64 `field:"optional" json:"bias" yaml:"bias"`
	// coordinates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#coordinates AwsRecordsExclusive#coordinates}
	// Experimental.
	Coordinates interface{} `field:"optional" json:"coordinates" yaml:"coordinates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#local_zone_group AwsRecordsExclusive#local_zone_group}.
	// Experimental.
	LocalZoneGroup *string `field:"optional" json:"localZoneGroup" yaml:"localZoneGroup"`
}

