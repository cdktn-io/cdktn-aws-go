package cloudsearch


// Experimental.
type AwsDomain_IndexFieldProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#name AwsDomain#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#type AwsDomain#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#analysis_scheme AwsDomain#analysis_scheme}.
	// Experimental.
	AnalysisScheme *string `field:"optional" json:"analysisScheme" yaml:"analysisScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#default_value AwsDomain#default_value}.
	// Experimental.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#facet AwsDomain#facet}.
	// Experimental.
	Facet interface{} `field:"optional" json:"facet" yaml:"facet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#highlight AwsDomain#highlight}.
	// Experimental.
	Highlight interface{} `field:"optional" json:"highlight" yaml:"highlight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#return AwsDomain#return}.
	// Experimental.
	Return interface{} `field:"optional" json:"return" yaml:"return"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#search AwsDomain#search}.
	// Experimental.
	Search interface{} `field:"optional" json:"search" yaml:"search"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#sort AwsDomain#sort}.
	// Experimental.
	Sort interface{} `field:"optional" json:"sort" yaml:"sort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudsearch_domain#source_fields AwsDomain#source_fields}.
	// Experimental.
	SourceFields *string `field:"optional" json:"sourceFields" yaml:"sourceFields"`
}

