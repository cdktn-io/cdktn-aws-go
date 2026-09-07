package rds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsClusterSnapshotCopyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#source_db_cluster_snapshot_identifier AwsClusterSnapshotCopy#source_db_cluster_snapshot_identifier}.
	// Experimental.
	SourceDbClusterSnapshotIdentifier *string `field:"required" json:"sourceDbClusterSnapshotIdentifier" yaml:"sourceDbClusterSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#target_db_cluster_snapshot_identifier AwsClusterSnapshotCopy#target_db_cluster_snapshot_identifier}.
	// Experimental.
	TargetDbClusterSnapshotIdentifier *string `field:"required" json:"targetDbClusterSnapshotIdentifier" yaml:"targetDbClusterSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#copy_tags AwsClusterSnapshotCopy#copy_tags}.
	// Experimental.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#destination_region AwsClusterSnapshotCopy#destination_region}.
	// Experimental.
	DestinationRegion *string `field:"optional" json:"destinationRegion" yaml:"destinationRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#kms_key_id AwsClusterSnapshotCopy#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#presigned_url AwsClusterSnapshotCopy#presigned_url}.
	// Experimental.
	PresignedUrl *string `field:"optional" json:"presignedUrl" yaml:"presignedUrl"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#region AwsClusterSnapshotCopy#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#shared_accounts AwsClusterSnapshotCopy#shared_accounts}.
	// Experimental.
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#tags AwsClusterSnapshotCopy#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster_snapshot_copy#timeouts AwsClusterSnapshotCopy#timeouts}
	// Experimental.
	Timeouts *AwsClusterSnapshotCopy_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

