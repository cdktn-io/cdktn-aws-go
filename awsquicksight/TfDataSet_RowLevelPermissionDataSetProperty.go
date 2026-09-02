package awsquicksight


// Experimental.
type TfDataSet_RowLevelPermissionDataSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#arn TfDataSet#arn}.
	// Experimental.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#permission_policy TfDataSet#permission_policy}.
	// Experimental.
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#format_version TfDataSet#format_version}.
	// Experimental.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#namespace TfDataSet#namespace}.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#status TfDataSet#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

