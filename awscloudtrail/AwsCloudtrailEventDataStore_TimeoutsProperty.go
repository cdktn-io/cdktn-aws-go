package awscloudtrail


// Experimental.
type AwsCloudtrailEventDataStore_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#create AwsCloudtrailEventDataStore#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#delete AwsCloudtrailEventDataStore#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#update AwsCloudtrailEventDataStore#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

