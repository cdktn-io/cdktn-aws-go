package securitylake


// Experimental.
type AwsSubscriber_CustomLogSourceResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#source_name AwsSubscriber#source_name}.
	// Experimental.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#source_version AwsSubscriber#source_version}.
	// Experimental.
	SourceVersion *string `field:"optional" json:"sourceVersion" yaml:"sourceVersion"`
}

