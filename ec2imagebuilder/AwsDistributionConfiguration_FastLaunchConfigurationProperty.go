package ec2imagebuilder


// Experimental.
type AwsDistributionConfiguration_FastLaunchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#account_id AwsDistributionConfiguration#account_id}.
	// Experimental.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#enabled AwsDistributionConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_template AwsDistributionConfiguration#launch_template}
	// Experimental.
	LaunchTemplate *AwsDistributionConfiguration_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#max_parallel_launches AwsDistributionConfiguration#max_parallel_launches}.
	// Experimental.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#snapshot_configuration AwsDistributionConfiguration#snapshot_configuration}
	// Experimental.
	SnapshotConfiguration *AwsDistributionConfiguration_SnapshotConfigurationProperty `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

