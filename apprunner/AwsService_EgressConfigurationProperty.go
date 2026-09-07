package apprunner


// Experimental.
type AwsService_EgressConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#egress_type AwsService#egress_type}.
	// Experimental.
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#vpc_connector_arn AwsService#vpc_connector_arn}.
	// Experimental.
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

