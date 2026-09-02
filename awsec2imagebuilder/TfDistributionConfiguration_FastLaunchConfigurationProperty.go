package awsec2imagebuilder


// Experimental.
type TfDistributionConfiguration_FastLaunchConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#account_id TfDistributionConfiguration#account_id}.
	// Experimental.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#enabled TfDistributionConfiguration#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#launch_template TfDistributionConfiguration#launch_template}
	// Experimental.
	LaunchTemplate *TfDistributionConfiguration_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#max_parallel_launches TfDistributionConfiguration#max_parallel_launches}.
	// Experimental.
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// snapshot_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_distribution_configuration#snapshot_configuration TfDistributionConfiguration#snapshot_configuration}
	// Experimental.
	SnapshotConfiguration *TfDistributionConfiguration_SnapshotConfigurationProperty `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

