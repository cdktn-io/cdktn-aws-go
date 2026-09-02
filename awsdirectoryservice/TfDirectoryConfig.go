package awsdirectoryservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDirectoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#name TfDirectory#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#password TfDirectory#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#alias TfDirectory#alias}.
	// Experimental.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// connect_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#connect_settings TfDirectory#connect_settings}
	// Experimental.
	ConnectSettings *TfDirectory_ConnectSettingsProperty `field:"optional" json:"connectSettings" yaml:"connectSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#description TfDirectory#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#desired_number_of_domain_controllers TfDirectory#desired_number_of_domain_controllers}.
	// Experimental.
	DesiredNumberOfDomainControllers *float64 `field:"optional" json:"desiredNumberOfDomainControllers" yaml:"desiredNumberOfDomainControllers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#edition TfDirectory#edition}.
	// Experimental.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#enable_directory_data_access TfDirectory#enable_directory_data_access}.
	// Experimental.
	EnableDirectoryDataAccess interface{} `field:"optional" json:"enableDirectoryDataAccess" yaml:"enableDirectoryDataAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#enable_sso TfDirectory#enable_sso}.
	// Experimental.
	EnableSso interface{} `field:"optional" json:"enableSso" yaml:"enableSso"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#id TfDirectory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#region TfDirectory#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#short_name TfDirectory#short_name}.
	// Experimental.
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#size TfDirectory#size}.
	// Experimental.
	Size *string `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#tags TfDirectory#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#tags_all TfDirectory#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#timeouts TfDirectory#timeouts}
	// Experimental.
	Timeouts *TfDirectory_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#type TfDirectory#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// vpc_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_directory#vpc_settings TfDirectory#vpc_settings}
	// Experimental.
	VpcSettings *TfDirectory_VpcSettingsProperty `field:"optional" json:"vpcSettings" yaml:"vpcSettings"`
}

