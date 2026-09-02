package awsemr


// Experimental.
type TfCluster_KerberosAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#kdc_admin_password TfCluster#kdc_admin_password}.
	// Experimental.
	KdcAdminPassword *string `field:"required" json:"kdcAdminPassword" yaml:"kdcAdminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#realm TfCluster#realm}.
	// Experimental.
	Realm *string `field:"required" json:"realm" yaml:"realm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ad_domain_join_password TfCluster#ad_domain_join_password}.
	// Experimental.
	AdDomainJoinPassword *string `field:"optional" json:"adDomainJoinPassword" yaml:"adDomainJoinPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ad_domain_join_user TfCluster#ad_domain_join_user}.
	// Experimental.
	AdDomainJoinUser *string `field:"optional" json:"adDomainJoinUser" yaml:"adDomainJoinUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#cross_realm_trust_principal_password TfCluster#cross_realm_trust_principal_password}.
	// Experimental.
	CrossRealmTrustPrincipalPassword *string `field:"optional" json:"crossRealmTrustPrincipalPassword" yaml:"crossRealmTrustPrincipalPassword"`
}

