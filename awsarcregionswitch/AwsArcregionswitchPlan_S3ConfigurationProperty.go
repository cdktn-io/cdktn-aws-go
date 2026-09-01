package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_S3ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#bucket_owner AwsArcregionswitchPlan#bucket_owner}.
	// Experimental.
	BucketOwner *string `field:"required" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#bucket_path AwsArcregionswitchPlan#bucket_path}.
	// Experimental.
	BucketPath *string `field:"required" json:"bucketPath" yaml:"bucketPath"`
}

