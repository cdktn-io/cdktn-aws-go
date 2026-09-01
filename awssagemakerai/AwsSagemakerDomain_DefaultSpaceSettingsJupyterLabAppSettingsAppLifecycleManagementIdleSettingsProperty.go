package awssagemakerai


// Experimental.
type AwsSagemakerDomain_DefaultSpaceSettingsJupyterLabAppSettingsAppLifecycleManagementIdleSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#idle_timeout_in_minutes AwsSagemakerDomain#idle_timeout_in_minutes}.
	// Experimental.
	IdleTimeoutInMinutes *float64 `field:"optional" json:"idleTimeoutInMinutes" yaml:"idleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#lifecycle_management AwsSagemakerDomain#lifecycle_management}.
	// Experimental.
	LifecycleManagement *string `field:"optional" json:"lifecycleManagement" yaml:"lifecycleManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#max_idle_timeout_in_minutes AwsSagemakerDomain#max_idle_timeout_in_minutes}.
	// Experimental.
	MaxIdleTimeoutInMinutes *float64 `field:"optional" json:"maxIdleTimeoutInMinutes" yaml:"maxIdleTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_domain#min_idle_timeout_in_minutes AwsSagemakerDomain#min_idle_timeout_in_minutes}.
	// Experimental.
	MinIdleTimeoutInMinutes *float64 `field:"optional" json:"minIdleTimeoutInMinutes" yaml:"minIdleTimeoutInMinutes"`
}

