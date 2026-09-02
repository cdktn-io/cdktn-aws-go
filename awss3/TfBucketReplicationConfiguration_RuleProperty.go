package awss3


// Experimental.
type TfBucketReplicationConfiguration_RuleProperty struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#destination TfBucketReplicationConfiguration#destination}
	// Experimental.
	Destination *TfBucketReplicationConfiguration_DestinationProperty `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#status TfBucketReplicationConfiguration#status}.
	// Experimental.
	Status *string `field:"required" json:"status" yaml:"status"`
	// delete_marker_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#delete_marker_replication TfBucketReplicationConfiguration#delete_marker_replication}
	// Experimental.
	DeleteMarkerReplication *TfBucketReplicationConfiguration_DeleteMarkerReplicationProperty `field:"optional" json:"deleteMarkerReplication" yaml:"deleteMarkerReplication"`
	// existing_object_replication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#existing_object_replication TfBucketReplicationConfiguration#existing_object_replication}
	// Experimental.
	ExistingObjectReplication *TfBucketReplicationConfiguration_ExistingObjectReplicationProperty `field:"optional" json:"existingObjectReplication" yaml:"existingObjectReplication"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#filter TfBucketReplicationConfiguration#filter}
	// Experimental.
	Filter *TfBucketReplicationConfiguration_FilterProperty `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#id TfBucketReplicationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#prefix TfBucketReplicationConfiguration#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#priority TfBucketReplicationConfiguration#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// source_selection_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_replication_configuration#source_selection_criteria TfBucketReplicationConfiguration#source_selection_criteria}
	// Experimental.
	SourceSelectionCriteria *TfBucketReplicationConfiguration_SourceSelectionCriteriaProperty `field:"optional" json:"sourceSelectionCriteria" yaml:"sourceSelectionCriteria"`
}

