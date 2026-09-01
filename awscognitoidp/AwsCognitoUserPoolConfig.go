package awscognitoidp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoUserPoolConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#name AwsCognitoUserPool#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// account_recovery_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#account_recovery_setting AwsCognitoUserPool#account_recovery_setting}
	// Experimental.
	AccountRecoverySetting *AwsCognitoUserPool_AccountRecoverySettingProperty `field:"optional" json:"accountRecoverySetting" yaml:"accountRecoverySetting"`
	// admin_create_user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#admin_create_user_config AwsCognitoUserPool#admin_create_user_config}
	// Experimental.
	AdminCreateUserConfig *AwsCognitoUserPool_AdminCreateUserConfigProperty `field:"optional" json:"adminCreateUserConfig" yaml:"adminCreateUserConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#alias_attributes AwsCognitoUserPool#alias_attributes}.
	// Experimental.
	AliasAttributes *[]*string `field:"optional" json:"aliasAttributes" yaml:"aliasAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#auto_verified_attributes AwsCognitoUserPool#auto_verified_attributes}.
	// Experimental.
	AutoVerifiedAttributes *[]*string `field:"optional" json:"autoVerifiedAttributes" yaml:"autoVerifiedAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#deletion_protection AwsCognitoUserPool#deletion_protection}.
	// Experimental.
	DeletionProtection *string `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// device_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#device_configuration AwsCognitoUserPool#device_configuration}
	// Experimental.
	DeviceConfiguration *AwsCognitoUserPool_DeviceConfigurationProperty `field:"optional" json:"deviceConfiguration" yaml:"deviceConfiguration"`
	// email_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_configuration AwsCognitoUserPool#email_configuration}
	// Experimental.
	EmailConfiguration *AwsCognitoUserPool_EmailConfigurationProperty `field:"optional" json:"emailConfiguration" yaml:"emailConfiguration"`
	// email_mfa_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_mfa_configuration AwsCognitoUserPool#email_mfa_configuration}
	// Experimental.
	EmailMfaConfiguration *AwsCognitoUserPool_EmailMfaConfigurationProperty `field:"optional" json:"emailMfaConfiguration" yaml:"emailMfaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_verification_message AwsCognitoUserPool#email_verification_message}.
	// Experimental.
	EmailVerificationMessage *string `field:"optional" json:"emailVerificationMessage" yaml:"emailVerificationMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#email_verification_subject AwsCognitoUserPool#email_verification_subject}.
	// Experimental.
	EmailVerificationSubject *string `field:"optional" json:"emailVerificationSubject" yaml:"emailVerificationSubject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#id AwsCognitoUserPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#lambda_config AwsCognitoUserPool#lambda_config}
	// Experimental.
	LambdaConfig *AwsCognitoUserPool_LambdaConfigProperty `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#mfa_configuration AwsCognitoUserPool#mfa_configuration}.
	// Experimental.
	MfaConfiguration *string `field:"optional" json:"mfaConfiguration" yaml:"mfaConfiguration"`
	// password_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#password_policy AwsCognitoUserPool#password_policy}
	// Experimental.
	PasswordPolicy *AwsCognitoUserPool_PasswordPolicyProperty `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#region AwsCognitoUserPool#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#schema AwsCognitoUserPool#schema}
	// Experimental.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
	// sign_in_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sign_in_policy AwsCognitoUserPool#sign_in_policy}
	// Experimental.
	SignInPolicy *AwsCognitoUserPool_SignInPolicyProperty `field:"optional" json:"signInPolicy" yaml:"signInPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sms_authentication_message AwsCognitoUserPool#sms_authentication_message}.
	// Experimental.
	SmsAuthenticationMessage *string `field:"optional" json:"smsAuthenticationMessage" yaml:"smsAuthenticationMessage"`
	// sms_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sms_configuration AwsCognitoUserPool#sms_configuration}
	// Experimental.
	SmsConfiguration *AwsCognitoUserPool_SmsConfigurationProperty `field:"optional" json:"smsConfiguration" yaml:"smsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#sms_verification_message AwsCognitoUserPool#sms_verification_message}.
	// Experimental.
	SmsVerificationMessage *string `field:"optional" json:"smsVerificationMessage" yaml:"smsVerificationMessage"`
	// software_token_mfa_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#software_token_mfa_configuration AwsCognitoUserPool#software_token_mfa_configuration}
	// Experimental.
	SoftwareTokenMfaConfiguration *AwsCognitoUserPool_SoftwareTokenMfaConfigurationProperty `field:"optional" json:"softwareTokenMfaConfiguration" yaml:"softwareTokenMfaConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#tags AwsCognitoUserPool#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#tags_all AwsCognitoUserPool#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// user_attribute_update_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#user_attribute_update_settings AwsCognitoUserPool#user_attribute_update_settings}
	// Experimental.
	UserAttributeUpdateSettings *AwsCognitoUserPool_UserAttributeUpdateSettingsProperty `field:"optional" json:"userAttributeUpdateSettings" yaml:"userAttributeUpdateSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#username_attributes AwsCognitoUserPool#username_attributes}.
	// Experimental.
	UsernameAttributes *[]*string `field:"optional" json:"usernameAttributes" yaml:"usernameAttributes"`
	// username_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#username_configuration AwsCognitoUserPool#username_configuration}
	// Experimental.
	UsernameConfiguration *AwsCognitoUserPool_UsernameConfigurationProperty `field:"optional" json:"usernameConfiguration" yaml:"usernameConfiguration"`
	// user_pool_add_ons block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#user_pool_add_ons AwsCognitoUserPool#user_pool_add_ons}
	// Experimental.
	UserPoolAddOns *AwsCognitoUserPool_UserPoolAddOnsProperty `field:"optional" json:"userPoolAddOns" yaml:"userPoolAddOns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#user_pool_tier AwsCognitoUserPool#user_pool_tier}.
	// Experimental.
	UserPoolTier *string `field:"optional" json:"userPoolTier" yaml:"userPoolTier"`
	// verification_message_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#verification_message_template AwsCognitoUserPool#verification_message_template}
	// Experimental.
	VerificationMessageTemplate *AwsCognitoUserPool_VerificationMessageTemplateProperty `field:"optional" json:"verificationMessageTemplate" yaml:"verificationMessageTemplate"`
	// web_authn_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cognito_user_pool#web_authn_configuration AwsCognitoUserPool#web_authn_configuration}
	// Experimental.
	WebAuthnConfiguration *AwsCognitoUserPool_WebAuthnConfigurationProperty `field:"optional" json:"webAuthnConfiguration" yaml:"webAuthnConfiguration"`
}

