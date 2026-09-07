package cognitoidp


// Experimental.
type AwsUserPool_LambdaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#create_auth_challenge AwsUserPool#create_auth_challenge}.
	// Experimental.
	CreateAuthChallenge *string `field:"optional" json:"createAuthChallenge" yaml:"createAuthChallenge"`
	// custom_email_sender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#custom_email_sender AwsUserPool#custom_email_sender}
	// Experimental.
	CustomEmailSender *AwsUserPool_CustomEmailSenderProperty `field:"optional" json:"customEmailSender" yaml:"customEmailSender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#custom_message AwsUserPool#custom_message}.
	// Experimental.
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// custom_sms_sender block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#custom_sms_sender AwsUserPool#custom_sms_sender}
	// Experimental.
	CustomSmsSender *AwsUserPool_CustomSmsSenderProperty `field:"optional" json:"customSmsSender" yaml:"customSmsSender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#define_auth_challenge AwsUserPool#define_auth_challenge}.
	// Experimental.
	DefineAuthChallenge *string `field:"optional" json:"defineAuthChallenge" yaml:"defineAuthChallenge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#kms_key_id AwsUserPool#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#post_authentication AwsUserPool#post_authentication}.
	// Experimental.
	PostAuthentication *string `field:"optional" json:"postAuthentication" yaml:"postAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#post_confirmation AwsUserPool#post_confirmation}.
	// Experimental.
	PostConfirmation *string `field:"optional" json:"postConfirmation" yaml:"postConfirmation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#pre_authentication AwsUserPool#pre_authentication}.
	// Experimental.
	PreAuthentication *string `field:"optional" json:"preAuthentication" yaml:"preAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#pre_sign_up AwsUserPool#pre_sign_up}.
	// Experimental.
	PreSignUp *string `field:"optional" json:"preSignUp" yaml:"preSignUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#pre_token_generation AwsUserPool#pre_token_generation}.
	// Experimental.
	PreTokenGeneration *string `field:"optional" json:"preTokenGeneration" yaml:"preTokenGeneration"`
	// pre_token_generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#pre_token_generation_config AwsUserPool#pre_token_generation_config}
	// Experimental.
	PreTokenGenerationConfig *AwsUserPool_PreTokenGenerationConfigProperty `field:"optional" json:"preTokenGenerationConfig" yaml:"preTokenGenerationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#user_migration AwsUserPool#user_migration}.
	// Experimental.
	UserMigration *string `field:"optional" json:"userMigration" yaml:"userMigration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#verify_auth_challenge_response AwsUserPool#verify_auth_challenge_response}.
	// Experimental.
	VerifyAuthChallengeResponse *string `field:"optional" json:"verifyAuthChallengeResponse" yaml:"verifyAuthChallengeResponse"`
}

