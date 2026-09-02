package awscostandusagereport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReportDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#additional_schema_elements TfReportDefinition#additional_schema_elements}.
	// Experimental.
	AdditionalSchemaElements *[]*string `field:"required" json:"additionalSchemaElements" yaml:"additionalSchemaElements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#compression TfReportDefinition#compression}.
	// Experimental.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#format TfReportDefinition#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#report_name TfReportDefinition#report_name}.
	// Experimental.
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#s3_bucket TfReportDefinition#s3_bucket}.
	// Experimental.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#s3_prefix TfReportDefinition#s3_prefix}.
	// Experimental.
	S3Prefix *string `field:"required" json:"s3Prefix" yaml:"s3Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#s3_region TfReportDefinition#s3_region}.
	// Experimental.
	S3Region *string `field:"required" json:"s3Region" yaml:"s3Region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#time_unit TfReportDefinition#time_unit}.
	// Experimental.
	TimeUnit *string `field:"required" json:"timeUnit" yaml:"timeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#additional_artifacts TfReportDefinition#additional_artifacts}.
	// Experimental.
	AdditionalArtifacts *[]*string `field:"optional" json:"additionalArtifacts" yaml:"additionalArtifacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#id TfReportDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#refresh_closed_reports TfReportDefinition#refresh_closed_reports}.
	// Experimental.
	RefreshClosedReports interface{} `field:"optional" json:"refreshClosedReports" yaml:"refreshClosedReports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#report_versioning TfReportDefinition#report_versioning}.
	// Experimental.
	ReportVersioning *string `field:"optional" json:"reportVersioning" yaml:"reportVersioning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#tags TfReportDefinition#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cur_report_definition#tags_all TfReportDefinition#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

