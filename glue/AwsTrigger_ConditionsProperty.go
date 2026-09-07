package glue


// Experimental.
type AwsTrigger_ConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#crawler_name AwsTrigger#crawler_name}.
	// Experimental.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#crawl_state AwsTrigger#crawl_state}.
	// Experimental.
	CrawlState *string `field:"optional" json:"crawlState" yaml:"crawlState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#job_name AwsTrigger#job_name}.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#logical_operator AwsTrigger#logical_operator}.
	// Experimental.
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#state AwsTrigger#state}.
	// Experimental.
	State *string `field:"optional" json:"state" yaml:"state"`
}

