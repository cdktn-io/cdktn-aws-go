package ecs


// Experimental.
type AwsService_AccessLogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#format AwsService#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#include_query_parameters AwsService#include_query_parameters}.
	// Experimental.
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

