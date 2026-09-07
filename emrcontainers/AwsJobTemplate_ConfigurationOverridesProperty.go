package emrcontainers


// Experimental.
type AwsJobTemplate_ConfigurationOverridesProperty struct {
	// application_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#application_configuration AwsJobTemplate#application_configuration}
	// Experimental.
	ApplicationConfiguration interface{} `field:"optional" json:"applicationConfiguration" yaml:"applicationConfiguration"`
	// monitoring_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrcontainers_job_template#monitoring_configuration AwsJobTemplate#monitoring_configuration}
	// Experimental.
	MonitoringConfiguration *AwsJobTemplate_MonitoringConfigurationProperty `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

