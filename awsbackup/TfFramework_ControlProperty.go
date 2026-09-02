package awsbackup


// Experimental.
type TfFramework_ControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#name TfFramework#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// input_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#input_parameter TfFramework#input_parameter}
	// Experimental.
	InputParameter interface{} `field:"optional" json:"inputParameter" yaml:"inputParameter"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/backup_framework#scope TfFramework#scope}
	// Experimental.
	Scope *TfFramework_ScopeProperty `field:"optional" json:"scope" yaml:"scope"`
}

