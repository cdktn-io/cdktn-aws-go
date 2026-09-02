package awssecretsmanager

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type EphemeralTfRandomPasswordConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_characters EphemeralTfRandomPassword#exclude_characters}.
	// Experimental.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_lowercase EphemeralTfRandomPassword#exclude_lowercase}.
	// Experimental.
	ExcludeLowercase interface{} `field:"optional" json:"excludeLowercase" yaml:"excludeLowercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_numbers EphemeralTfRandomPassword#exclude_numbers}.
	// Experimental.
	ExcludeNumbers interface{} `field:"optional" json:"excludeNumbers" yaml:"excludeNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_punctuation EphemeralTfRandomPassword#exclude_punctuation}.
	// Experimental.
	ExcludePunctuation interface{} `field:"optional" json:"excludePunctuation" yaml:"excludePunctuation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#exclude_uppercase EphemeralTfRandomPassword#exclude_uppercase}.
	// Experimental.
	ExcludeUppercase interface{} `field:"optional" json:"excludeUppercase" yaml:"excludeUppercase"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#include_space EphemeralTfRandomPassword#include_space}.
	// Experimental.
	IncludeSpace interface{} `field:"optional" json:"includeSpace" yaml:"includeSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#password_length EphemeralTfRandomPassword#password_length}.
	// Experimental.
	PasswordLength *float64 `field:"optional" json:"passwordLength" yaml:"passwordLength"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#region EphemeralTfRandomPassword#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/ephemeral-resources/secretsmanager_random_password#require_each_included_type EphemeralTfRandomPassword#require_each_included_type}.
	// Experimental.
	RequireEachIncludedType interface{} `field:"optional" json:"requireEachIncludedType" yaml:"requireEachIncludedType"`
}

