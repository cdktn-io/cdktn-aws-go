package awsdocumentdbelastic

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#admin_user_name TfCluster#admin_user_name}.
	// Experimental.
	AdminUserName *string `field:"required" json:"adminUserName" yaml:"adminUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#admin_user_password TfCluster#admin_user_password}.
	// Experimental.
	AdminUserPassword *string `field:"required" json:"adminUserPassword" yaml:"adminUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#auth_type TfCluster#auth_type}.
	// Experimental.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#name TfCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#shard_capacity TfCluster#shard_capacity}.
	// Experimental.
	ShardCapacity *float64 `field:"required" json:"shardCapacity" yaml:"shardCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#shard_count TfCluster#shard_count}.
	// Experimental.
	ShardCount *float64 `field:"required" json:"shardCount" yaml:"shardCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#backup_retention_period TfCluster#backup_retention_period}.
	// Experimental.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#kms_key_id TfCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#preferred_backup_window TfCluster#preferred_backup_window}.
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#preferred_maintenance_window TfCluster#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#shard_instance_count TfCluster#shard_instance_count}.
	// Experimental.
	ShardInstanceCount *float64 `field:"optional" json:"shardInstanceCount" yaml:"shardInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#subnet_ids TfCluster#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#timeouts TfCluster#timeouts}
	// Experimental.
	Timeouts *TfCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/docdbelastic_cluster#vpc_security_group_ids TfCluster#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

