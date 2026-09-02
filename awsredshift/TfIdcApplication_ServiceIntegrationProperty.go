package awsredshift


// Experimental.
type TfIdcApplication_ServiceIntegrationProperty struct {
	// lake_formation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#lake_formation TfIdcApplication#lake_formation}
	// Experimental.
	LakeFormation interface{} `field:"optional" json:"lakeFormation" yaml:"lakeFormation"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#redshift TfIdcApplication#redshift}
	// Experimental.
	Redshift interface{} `field:"optional" json:"redshift" yaml:"redshift"`
	// s3_access_grants block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#s3_access_grants TfIdcApplication#s3_access_grants}
	// Experimental.
	S3AccessGrants interface{} `field:"optional" json:"s3AccessGrants" yaml:"s3AccessGrants"`
}

