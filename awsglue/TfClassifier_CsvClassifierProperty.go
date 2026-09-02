package awsglue


// Experimental.
type TfClassifier_CsvClassifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#allow_single_column TfClassifier#allow_single_column}.
	// Experimental.
	AllowSingleColumn interface{} `field:"optional" json:"allowSingleColumn" yaml:"allowSingleColumn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#contains_header TfClassifier#contains_header}.
	// Experimental.
	ContainsHeader *string `field:"optional" json:"containsHeader" yaml:"containsHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#custom_datatype_configured TfClassifier#custom_datatype_configured}.
	// Experimental.
	CustomDatatypeConfigured interface{} `field:"optional" json:"customDatatypeConfigured" yaml:"customDatatypeConfigured"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#custom_datatypes TfClassifier#custom_datatypes}.
	// Experimental.
	CustomDatatypes *[]*string `field:"optional" json:"customDatatypes" yaml:"customDatatypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#delimiter TfClassifier#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#disable_value_trimming TfClassifier#disable_value_trimming}.
	// Experimental.
	DisableValueTrimming interface{} `field:"optional" json:"disableValueTrimming" yaml:"disableValueTrimming"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#header TfClassifier#header}.
	// Experimental.
	Header *[]*string `field:"optional" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#quote_symbol TfClassifier#quote_symbol}.
	// Experimental.
	QuoteSymbol *string `field:"optional" json:"quoteSymbol" yaml:"quoteSymbol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#serde TfClassifier#serde}.
	// Experimental.
	Serde *string `field:"optional" json:"serde" yaml:"serde"`
}

