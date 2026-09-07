package secretsmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSecretRotationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#secret_id AwsSecretRotation#secret_id}.
	// Experimental.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// external_secret_rotation_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#external_secret_rotation_metadata AwsSecretRotation#external_secret_rotation_metadata}
	// Experimental.
	ExternalSecretRotationMetadata interface{} `field:"optional" json:"externalSecretRotationMetadata" yaml:"externalSecretRotationMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#external_secret_rotation_role_arn AwsSecretRotation#external_secret_rotation_role_arn}.
	// Experimental.
	ExternalSecretRotationRoleArn *string `field:"optional" json:"externalSecretRotationRoleArn" yaml:"externalSecretRotationRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#id AwsSecretRotation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#region AwsSecretRotation#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#rotate_immediately AwsSecretRotation#rotate_immediately}.
	// Experimental.
	RotateImmediately interface{} `field:"optional" json:"rotateImmediately" yaml:"rotateImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#rotation_enabled AwsSecretRotation#rotation_enabled}.
	// Experimental.
	RotationEnabled interface{} `field:"optional" json:"rotationEnabled" yaml:"rotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#rotation_lambda_arn AwsSecretRotation#rotation_lambda_arn}.
	// Experimental.
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// rotation_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_rotation#rotation_rules AwsSecretRotation#rotation_rules}
	// Experimental.
	RotationRules *AwsSecretRotation_RotationRulesProperty `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

