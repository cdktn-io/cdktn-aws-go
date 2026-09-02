package awscognitoidp


// Experimental.
type TfUserPool_SmsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#external_id TfUserPool#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sns_caller_arn TfUserPool#sns_caller_arn}.
	// Experimental.
	SnsCallerArn *string `field:"required" json:"snsCallerArn" yaml:"snsCallerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sns_region TfUserPool#sns_region}.
	// Experimental.
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
}

