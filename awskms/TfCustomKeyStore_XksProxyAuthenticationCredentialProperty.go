package awskms


// Experimental.
type TfCustomKeyStore_XksProxyAuthenticationCredentialProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#access_key_id TfCustomKeyStore#access_key_id}.
	// Experimental.
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kms_custom_key_store#raw_secret_access_key TfCustomKeyStore#raw_secret_access_key}.
	// Experimental.
	RawSecretAccessKey *string `field:"required" json:"rawSecretAccessKey" yaml:"rawSecretAccessKey"`
}

