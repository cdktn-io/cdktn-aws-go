package securityhub


// Experimental.
type AwsConnectorV2_ServiceNowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_connector_v2#instance_name AwsConnectorV2#instance_name}.
	// Experimental.
	InstanceName *string `field:"required" json:"instanceName" yaml:"instanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_connector_v2#secret_arn AwsConnectorV2#secret_arn}.
	// Experimental.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}

