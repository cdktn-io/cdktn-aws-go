package glue


// Experimental.
type AwsTrigger_ActionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#arguments AwsTrigger#arguments}.
	// Experimental.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#crawler_name AwsTrigger#crawler_name}.
	// Experimental.
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#job_name AwsTrigger#job_name}.
	// Experimental.
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// notification_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#notification_property AwsTrigger#notification_property}
	// Experimental.
	NotificationProperty *AwsTrigger_NotificationPropertyProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#security_configuration AwsTrigger#security_configuration}.
	// Experimental.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#timeout AwsTrigger#timeout}.
	// Experimental.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

