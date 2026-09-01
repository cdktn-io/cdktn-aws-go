package awsdocumentdb


// Experimental.
type AwsDocdbCluster_RestoreToPointInTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#source_cluster_identifier AwsDocdbCluster#source_cluster_identifier}.
	// Experimental.
	SourceClusterIdentifier *string `field:"required" json:"sourceClusterIdentifier" yaml:"sourceClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#restore_to_time AwsDocdbCluster#restore_to_time}.
	// Experimental.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#restore_type AwsDocdbCluster#restore_type}.
	// Experimental.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdb_cluster#use_latest_restorable_time AwsDocdbCluster#use_latest_restorable_time}.
	// Experimental.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

