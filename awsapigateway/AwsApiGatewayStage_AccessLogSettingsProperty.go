package awsapigateway


// Experimental.
type AwsApiGatewayStage_AccessLogSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#destination_arn AwsApiGatewayStage#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#format AwsApiGatewayStage#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
}

