package awsquicksight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRefreshScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#data_set_id TfRefreshSchedule#data_set_id}.
	// Experimental.
	DataSetId *string `field:"required" json:"dataSetId" yaml:"dataSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#schedule_id TfRefreshSchedule#schedule_id}.
	// Experimental.
	ScheduleId *string `field:"required" json:"scheduleId" yaml:"scheduleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#aws_account_id TfRefreshSchedule#aws_account_id}.
	// Experimental.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#region TfRefreshSchedule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_refresh_schedule#schedule TfRefreshSchedule#schedule}
	// Experimental.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

