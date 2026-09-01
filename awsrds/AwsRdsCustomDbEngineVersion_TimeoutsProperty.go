package awsrds


// Experimental.
type AwsRdsCustomDbEngineVersion_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#create AwsRdsCustomDbEngineVersion#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#delete AwsRdsCustomDbEngineVersion#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#update AwsRdsCustomDbEngineVersion#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

