package awsglue


// Experimental.
type TfCrawler_LakeFormationConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#account_id TfCrawler#account_id}.
	// Experimental.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#use_lake_formation_credentials TfCrawler#use_lake_formation_credentials}.
	// Experimental.
	UseLakeFormationCredentials interface{} `field:"optional" json:"useLakeFormationCredentials" yaml:"useLakeFormationCredentials"`
}

