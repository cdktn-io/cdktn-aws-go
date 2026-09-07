package emr


// Experimental.
type AwsCluster_KerberosAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#kdc_admin_password AwsCluster#kdc_admin_password}.
	// Experimental.
	KdcAdminPassword *string `field:"required" json:"kdcAdminPassword" yaml:"kdcAdminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#realm AwsCluster#realm}.
	// Experimental.
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ad_domain_join_password AwsCluster#ad_domain_join_password}.
	// Experimental.
	AdDomainJoinPassword *string `field:"optional" json:"adDomainJoinPassword" yaml:"adDomainJoinPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ad_domain_join_user AwsCluster#ad_domain_join_user}.
	// Experimental.
	AdDomainJoinUser *string `field:"optional" json:"adDomainJoinUser" yaml:"adDomainJoinUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#cross_realm_trust_principal_password AwsCluster#cross_realm_trust_principal_password}.
	// Experimental.
	CrossRealmTrustPrincipalPassword *string `field:"optional" json:"crossRealmTrustPrincipalPassword" yaml:"crossRealmTrustPrincipalPassword"`
}

