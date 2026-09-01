package awsecs


// Experimental.
type AwsEcsService_AccessLogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#format AwsEcsService#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#include_query_parameters AwsEcsService#include_query_parameters}.
	// Experimental.
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

