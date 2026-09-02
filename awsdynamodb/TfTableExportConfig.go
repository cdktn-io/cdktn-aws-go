package awsdynamodb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTableExportConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#s3_bucket TfTableExport#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#table_arn TfTableExport#table_arn}.
	// Experimental.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_format TfTableExport#export_format}.
	// Experimental.
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_time TfTableExport#export_time}.
	// Experimental.
	ExportTime *string `field:"optional" json:"exportTime" yaml:"exportTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#export_type TfTableExport#export_type}.
	// Experimental.
	ExportType *string `field:"optional" json:"exportType" yaml:"exportType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#id TfTableExport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// incremental_export_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#incremental_export_specification TfTableExport#incremental_export_specification}
	// Experimental.
	IncrementalExportSpecification *TfTableExport_IncrementalExportSpecificationProperty `field:"optional" json:"incrementalExportSpecification" yaml:"incrementalExportSpecification"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#region TfTableExport#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#s3_bucket_owner TfTableExport#s3_bucket_owner}.
	// Experimental.
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#s3_prefix TfTableExport#s3_prefix}.
	// Experimental.
	S3Prefix *string `field:"optional" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#s3_sse_algorithm TfTableExport#s3_sse_algorithm}.
	// Experimental.
	S3SseAlgorithm *string `field:"optional" json:"s3SseAlgorithm" yaml:"s3SseAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#s3_sse_kms_key_id TfTableExport#s3_sse_kms_key_id}.
	// Experimental.
	S3SseKmsKeyId *string `field:"optional" json:"s3SseKmsKeyId" yaml:"s3SseKmsKeyId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export#timeouts TfTableExport#timeouts}
	// Experimental.
	Timeouts *TfTableExport_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

