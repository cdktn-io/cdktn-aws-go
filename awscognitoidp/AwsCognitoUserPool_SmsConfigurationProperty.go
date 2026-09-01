package awscognitoidp


// Experimental.
type AwsCognitoUserPool_SmsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#external_id AwsCognitoUserPool#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sns_caller_arn AwsCognitoUserPool#sns_caller_arn}.
	// Experimental.
	SnsCallerArn *string `field:"required" json:"snsCallerArn" yaml:"snsCallerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sns_region AwsCognitoUserPool#sns_region}.
	// Experimental.
	SnsRegion *string `field:"optional" json:"snsRegion" yaml:"snsRegion"`
}

