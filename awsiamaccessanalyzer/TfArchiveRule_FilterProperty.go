package awsiamaccessanalyzer


// Experimental.
type TfArchiveRule_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_archive_rule#criteria TfArchiveRule#criteria}.
	// Experimental.
	Criteria *string `field:"required" json:"criteria" yaml:"criteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_archive_rule#contains TfArchiveRule#contains}.
	// Experimental.
	Contains *[]*string `field:"optional" json:"contains" yaml:"contains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_archive_rule#eq TfArchiveRule#eq}.
	// Experimental.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_archive_rule#exists TfArchiveRule#exists}.
	// Experimental.
	Exists *string `field:"optional" json:"exists" yaml:"exists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/accessanalyzer_archive_rule#neq TfArchiveRule#neq}.
	// Experimental.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

