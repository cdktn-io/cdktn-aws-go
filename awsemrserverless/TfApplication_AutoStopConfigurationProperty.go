package awsemrserverless


// Experimental.
type TfApplication_AutoStopConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#enabled TfApplication#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#idle_timeout_minutes TfApplication#idle_timeout_minutes}.
	// Experimental.
	IdleTimeoutMinutes *float64 `field:"optional" json:"idleTimeoutMinutes" yaml:"idleTimeoutMinutes"`
}

