package awskinesisanalyticsv2


// Experimental.
type TfApplication_ParallelismConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#configuration_type TfApplication#configuration_type}.
	// Experimental.
	ConfigurationType *string `field:"required" json:"configurationType" yaml:"configurationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#auto_scaling_enabled TfApplication#auto_scaling_enabled}.
	// Experimental.
	AutoScalingEnabled interface{} `field:"optional" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#parallelism TfApplication#parallelism}.
	// Experimental.
	Parallelism *float64 `field:"optional" json:"parallelism" yaml:"parallelism"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#parallelism_per_kpu TfApplication#parallelism_per_kpu}.
	// Experimental.
	ParallelismPerKpu *float64 `field:"optional" json:"parallelismPerKpu" yaml:"parallelismPerKpu"`
}

