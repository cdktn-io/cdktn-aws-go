package awsbackup


// Experimental.
type AwsBackupFramework_ControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#name AwsBackupFramework#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#input_parameter AwsBackupFramework#input_parameter}
	// Experimental.
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#scope AwsBackupFramework#scope}
	// Experimental.
	Scope *AwsBackupFramework_ScopeProperty `field:"optional" json:"scope" yaml:"scope"`
}

