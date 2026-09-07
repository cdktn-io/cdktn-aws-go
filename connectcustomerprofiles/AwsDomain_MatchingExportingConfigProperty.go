package connectcustomerprofiles


// Experimental.
type AwsDomain_MatchingExportingConfigProperty struct {
	// s3_exporting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#s3_exporting AwsDomain#s3_exporting}
	// Experimental.
	S3Exporting *AwsDomain_MatchingExportingConfigS3ExportingProperty `field:"optional" json:"s3Exporting" yaml:"s3Exporting"`
}

