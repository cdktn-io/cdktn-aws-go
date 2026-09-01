package awsrds


// Experimental.
type AwsRdsCluster_RestoreToPointInTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#restore_to_time AwsRdsCluster#restore_to_time}.
	// Experimental.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#restore_type AwsRdsCluster#restore_type}.
	// Experimental.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#source_cluster_identifier AwsRdsCluster#source_cluster_identifier}.
	// Experimental.
	SourceClusterIdentifier *string `field:"optional" json:"sourceClusterIdentifier" yaml:"sourceClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#source_cluster_resource_id AwsRdsCluster#source_cluster_resource_id}.
	// Experimental.
	SourceClusterResourceId *string `field:"optional" json:"sourceClusterResourceId" yaml:"sourceClusterResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#use_latest_restorable_time AwsRdsCluster#use_latest_restorable_time}.
	// Experimental.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

