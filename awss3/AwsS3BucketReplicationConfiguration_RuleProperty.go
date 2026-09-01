package awss3


// Experimental.
type AwsS3BucketReplicationConfiguration_RuleProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#destination AwsS3BucketReplicationConfiguration#destination}
	// Experimental.
	Destination *AwsS3BucketReplicationConfiguration_DestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status AwsS3BucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// delete_marker_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#delete_marker_replication AwsS3BucketReplicationConfiguration#delete_marker_replication}
	// Experimental.
	DeleteMarkerReplication *AwsS3BucketReplicationConfiguration_DeleteMarkerReplicationProperty `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// existing_object_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#existing_object_replication AwsS3BucketReplicationConfiguration#existing_object_replication}
	// Experimental.
	ExistingObjectReplication *AwsS3BucketReplicationConfiguration_ExistingObjectReplicationProperty `field:"optional" json:"existingObjectReplication" yaml:"existingObjectReplication"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#filter AwsS3BucketReplicationConfiguration#filter}
	// Experimental.
	Filter *AwsS3BucketReplicationConfiguration_FilterProperty `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#id AwsS3BucketReplicationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#prefix AwsS3BucketReplicationConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#priority AwsS3BucketReplicationConfiguration#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// source_selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#source_selection_criteria AwsS3BucketReplicationConfiguration#source_selection_criteria}
	// Experimental.
	SourceSelectionCriteria *AwsS3BucketReplicationConfiguration_SourceSelectionCriteriaProperty `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

