package awsdevopsguru


// Experimental.
type TfServiceIntegration_KmsServerSideEncryptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#kms_key_id TfServiceIntegration#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#opt_in_status TfServiceIntegration#opt_in_status}.
	// Experimental.
	OptInStatus *string `field:"optional" json:"optInStatus" yaml:"optInStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_service_integration#type TfServiceIntegration#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

