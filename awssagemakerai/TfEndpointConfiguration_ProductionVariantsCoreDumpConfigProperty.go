package awssagemakerai


// Experimental.
type TfEndpointConfiguration_ProductionVariantsCoreDumpConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#destination_s3_uri TfEndpointConfiguration#destination_s3_uri}.
	// Experimental.
	DestinationS3Uri *string `field:"required" json:"destinationS3Uri" yaml:"destinationS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#kms_key_id TfEndpointConfiguration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

