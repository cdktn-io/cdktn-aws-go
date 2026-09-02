package awsamp


// Experimental.
type TfScraper_RoleConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#source_role_arn TfScraper#source_role_arn}.
	// Experimental.
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_scraper#target_role_arn TfScraper#target_role_arn}.
	// Experimental.
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

