package apigatewayv2


// Experimental.
type AwsStage_AccessLogSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_stage#destination_arn AwsStage#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apigatewayv2_stage#format AwsStage#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
}

