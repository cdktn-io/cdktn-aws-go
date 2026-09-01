package awsapigatewayv2


// Experimental.
type AwsApigatewayv2Stage_AccessLogSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_stage#destination_arn AwsApigatewayv2Stage#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_stage#format AwsApigatewayv2Stage#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
}

