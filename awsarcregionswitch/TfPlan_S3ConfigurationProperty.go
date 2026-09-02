package awsarcregionswitch


// Experimental.
type TfPlan_S3ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#bucket_owner TfPlan#bucket_owner}.
	// Experimental.
	BucketOwner *string `field:"required" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#bucket_path TfPlan#bucket_path}.
	// Experimental.
	BucketPath *string `field:"required" json:"bucketPath" yaml:"bucketPath"`
}

