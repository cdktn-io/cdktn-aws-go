package awsdms


// Experimental.
type TfReplicationInstance_KerberosAuthenticationSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#key_cache_secret_iam_arn TfReplicationInstance#key_cache_secret_iam_arn}.
	// Experimental.
	KeyCacheSecretIamArn *string `field:"required" json:"keyCacheSecretIamArn" yaml:"keyCacheSecretIamArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#key_cache_secret_id TfReplicationInstance#key_cache_secret_id}.
	// Experimental.
	KeyCacheSecretId *string `field:"required" json:"keyCacheSecretId" yaml:"keyCacheSecretId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_instance#krb5_file_contents TfReplicationInstance#krb5_file_contents}.
	// Experimental.
	Krb5FileContents *string `field:"required" json:"krb5FileContents" yaml:"krb5FileContents"`
}

