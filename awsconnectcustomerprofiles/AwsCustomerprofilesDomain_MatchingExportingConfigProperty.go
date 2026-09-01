package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_MatchingExportingConfigProperty struct {
	// s3_exporting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#s3_exporting AwsCustomerprofilesDomain#s3_exporting}
	// Experimental.
	S3Exporting *AwsCustomerprofilesDomain_MatchingExportingConfigS3ExportingProperty `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

