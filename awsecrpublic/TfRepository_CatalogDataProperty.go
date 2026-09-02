package awsecrpublic


// Experimental.
type TfRepository_CatalogDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#about_text TfRepository#about_text}.
	// Experimental.
	AboutText *string `field:"optional" json:"aboutText" yaml:"aboutText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#architectures TfRepository#architectures}.
	// Experimental.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#description TfRepository#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#logo_image_blob TfRepository#logo_image_blob}.
	// Experimental.
	LogoImageBlob *string `field:"optional" json:"logoImageBlob" yaml:"logoImageBlob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#operating_systems TfRepository#operating_systems}.
	// Experimental.
	OperatingSystems *[]*string `field:"optional" json:"operatingSystems" yaml:"operatingSystems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecrpublic_repository#usage_text TfRepository#usage_text}.
	// Experimental.
	UsageText *string `field:"optional" json:"usageText" yaml:"usageText"`
}

