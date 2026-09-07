package chimesdkmediapipelines


// Experimental.
type AwsMediaInsightsPipelineConfiguration_PostCallAnalyticsSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#data_access_role_arn AwsMediaInsightsPipelineConfiguration#data_access_role_arn}.
	// Experimental.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#output_location AwsMediaInsightsPipelineConfiguration#output_location}.
	// Experimental.
	OutputLocation *string `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_output AwsMediaInsightsPipelineConfiguration#content_redaction_output}.
	// Experimental.
	ContentRedactionOutput *string `field:"optional" json:"contentRedactionOutput" yaml:"contentRedactionOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#output_encryption_kms_key_id AwsMediaInsightsPipelineConfiguration#output_encryption_kms_key_id}.
	// Experimental.
	OutputEncryptionKmsKeyId *string `field:"optional" json:"outputEncryptionKmsKeyId" yaml:"outputEncryptionKmsKeyId"`
}

