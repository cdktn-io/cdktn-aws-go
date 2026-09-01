package awstranscribe


// Experimental.
type AwsTranscribeLanguageModel_InputDataConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transcribe_language_model#data_access_role_arn AwsTranscribeLanguageModel#data_access_role_arn}.
	// Experimental.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transcribe_language_model#s3_uri AwsTranscribeLanguageModel#s3_uri}.
	// Experimental.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transcribe_language_model#tuning_data_s3_uri AwsTranscribeLanguageModel#tuning_data_s3_uri}.
	// Experimental.
	TuningDataS3Uri *string `field:"optional" json:"tuningDataS3Uri" yaml:"tuningDataS3Uri"`
}

