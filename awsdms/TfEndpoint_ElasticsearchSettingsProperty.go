package awsdms


// Experimental.
type TfEndpoint_ElasticsearchSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#endpoint_uri TfEndpoint#endpoint_uri}.
	// Experimental.
	EndpointUri *string `field:"required" json:"endpointUri" yaml:"endpointUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#service_access_role_arn TfEndpoint#service_access_role_arn}.
	// Experimental.
	ServiceAccessRoleArn *string `field:"required" json:"serviceAccessRoleArn" yaml:"serviceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#error_retry_duration TfEndpoint#error_retry_duration}.
	// Experimental.
	ErrorRetryDuration *float64 `field:"optional" json:"errorRetryDuration" yaml:"errorRetryDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#full_load_error_percentage TfEndpoint#full_load_error_percentage}.
	// Experimental.
	FullLoadErrorPercentage *float64 `field:"optional" json:"fullLoadErrorPercentage" yaml:"fullLoadErrorPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint#use_new_mapping_type TfEndpoint#use_new_mapping_type}.
	// Experimental.
	UseNewMappingType interface{} `field:"optional" json:"useNewMappingType" yaml:"useNewMappingType"`
}

