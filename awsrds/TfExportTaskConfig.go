package awsrds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfExportTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#export_task_identifier TfExportTask#export_task_identifier}.
	// Experimental.
	ExportTaskIdentifier *string `field:"required" json:"exportTaskIdentifier" yaml:"exportTaskIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#iam_role_arn TfExportTask#iam_role_arn}.
	// Experimental.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#kms_key_id TfExportTask#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#s3_bucket_name TfExportTask#s3_bucket_name}.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#source_arn TfExportTask#source_arn}.
	// Experimental.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#export_only TfExportTask#export_only}.
	// Experimental.
	ExportOnly *[]*string `field:"optional" json:"exportOnly" yaml:"exportOnly"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#region TfExportTask#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#s3_prefix TfExportTask#s3_prefix}.
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_export_task#timeouts TfExportTask#timeouts}
	// Experimental.
	Timeouts *TfExportTask_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

