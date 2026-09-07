package rds


// Experimental.
type AwsCustomDbEngineVersion_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#create AwsCustomDbEngineVersion#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#delete AwsCustomDbEngineVersion#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_custom_db_engine_version#update AwsCustomDbEngineVersion#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

