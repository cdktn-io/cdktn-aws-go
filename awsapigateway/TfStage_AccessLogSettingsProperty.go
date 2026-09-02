package awsapigateway


// Experimental.
type TfStage_AccessLogSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#destination_arn TfStage#destination_arn}.
	// Experimental.
	DestinationArn *string `field:"required" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_stage#format TfStage#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
}

