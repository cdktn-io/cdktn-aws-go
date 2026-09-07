package athena


// Experimental.
type AwsWorkgroup_QueryResultsS3AccessGrantsConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#authentication_type AwsWorkgroup#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enable_s3_access_grants AwsWorkgroup#enable_s3_access_grants}.
	// Experimental.
	EnableS3AccessGrants interface{} `field:"required" json:"enableS3AccessGrants" yaml:"enableS3AccessGrants"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#create_user_level_prefix AwsWorkgroup#create_user_level_prefix}.
	// Experimental.
	CreateUserLevelPrefix interface{} `field:"optional" json:"createUserLevelPrefix" yaml:"createUserLevelPrefix"`
}

