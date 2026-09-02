package awssecretsmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSecretVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#secret_id TfSecretVersion#secret_id}.
	// Experimental.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#id TfSecretVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#region TfSecretVersion#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#secret_binary TfSecretVersion#secret_binary}.
	// Experimental.
	SecretBinary *string `field:"optional" json:"secretBinary" yaml:"secretBinary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#secret_string TfSecretVersion#secret_string}.
	// Experimental.
	SecretString *string `field:"optional" json:"secretString" yaml:"secretString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#secret_string_wo TfSecretVersion#secret_string_wo}.
	// Experimental.
	SecretStringWo *string `field:"optional" json:"secretStringWo" yaml:"secretStringWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#secret_string_wo_version TfSecretVersion#secret_string_wo_version}.
	// Experimental.
	SecretStringWoVersion *float64 `field:"optional" json:"secretStringWoVersion" yaml:"secretStringWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/secretsmanager_secret_version#version_stages TfSecretVersion#version_stages}.
	// Experimental.
	VersionStages *[]*string `field:"optional" json:"versionStages" yaml:"versionStages"`
}

