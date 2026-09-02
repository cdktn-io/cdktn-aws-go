package awssagemakerai

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMlflowAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#artifact_store_uri TfMlflowApp#artifact_store_uri}.
	// Experimental.
	ArtifactStoreUri *string `field:"required" json:"artifactStoreUri" yaml:"artifactStoreUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#name TfMlflowApp#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#role_arn TfMlflowApp#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#account_default_status TfMlflowApp#account_default_status}.
	// Experimental.
	AccountDefaultStatus *string `field:"optional" json:"accountDefaultStatus" yaml:"accountDefaultStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#default_domain_id_list TfMlflowApp#default_domain_id_list}.
	// Experimental.
	DefaultDomainIdList *[]*string `field:"optional" json:"defaultDomainIdList" yaml:"defaultDomainIdList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#model_registration_mode TfMlflowApp#model_registration_mode}.
	// Experimental.
	ModelRegistrationMode *string `field:"optional" json:"modelRegistrationMode" yaml:"modelRegistrationMode"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#region TfMlflowApp#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#tags TfMlflowApp#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#timeouts TfMlflowApp#timeouts}
	// Experimental.
	Timeouts *TfMlflowApp_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_mlflow_app#weekly_maintenance_window_start TfMlflowApp#weekly_maintenance_window_start}.
	// Experimental.
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

