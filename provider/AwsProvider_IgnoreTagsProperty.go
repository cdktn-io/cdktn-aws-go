package provider


// Experimental.
type AwsProvider_IgnoreTagsProperty struct {
	// Resource tag key prefixes to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEY_PREFIXES environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs#key_prefixes AwsProvider#key_prefixes}
	// Experimental.
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Resource tag keys to ignore across all resources. Can also be configured with the TF_AWS_IGNORE_TAGS_KEYS environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs#keys AwsProvider#keys}
	// Experimental.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
}

