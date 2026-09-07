package sesv2


// Experimental.
type AwsConfigurationSet_SuppressionOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set#suppressed_reasons AwsConfigurationSet#suppressed_reasons}.
	// Experimental.
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
}

