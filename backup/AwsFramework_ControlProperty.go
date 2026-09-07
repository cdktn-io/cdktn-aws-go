package backup


// Experimental.
type AwsFramework_ControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#name AwsFramework#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#input_parameter AwsFramework#input_parameter}
	// Experimental.
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#scope AwsFramework#scope}
	// Experimental.
	Scope *AwsFramework_ScopeProperty `field:"optional" json:"scope" yaml:"scope"`
}

