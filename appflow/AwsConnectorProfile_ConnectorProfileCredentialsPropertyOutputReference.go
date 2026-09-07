package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomConnector() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	// Experimental.
	Datadog() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	// Experimental.
	Dynatrace() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	// Experimental.
	Honeycode() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	// Experimental.
	InforNexus() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	// Experimental.
	InternalValue() *AwsConnectorProfile_ConnectorProfileCredentialsProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfileCredentialsProperty)
	// Experimental.
	Marketo() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	// Experimental.
	Redshift() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	// Experimental.
	Salesforce() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	// Experimental.
	SapoData() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	// Experimental.
	ServiceNow() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	// Experimental.
	Singular() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	// Experimental.
	Slack() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	// Experimental.
	Snowflake() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	// Experimental.
	Veeva() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	// Experimental.
	Zendesk() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAmplitude(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty)
	// Experimental.
	PutSalesforce(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty)
	// Experimental.
	PutSlack(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty)
	// Experimental.
	PutSnowflake(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty)
	// Experimental.
	PutZendesk(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty)
	// Experimental.
	ResetAmplitude()
	// Experimental.
	ResetCustomConnector()
	// Experimental.
	ResetDatadog()
	// Experimental.
	ResetDynatrace()
	// Experimental.
	ResetGoogleAnalytics()
	// Experimental.
	ResetHoneycode()
	// Experimental.
	ResetInforNexus()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetRedshift()
	// Experimental.
	ResetSalesforce()
	// Experimental.
	ResetSapoData()
	// Experimental.
	ResetServiceNow()
	// Experimental.
	ResetSingular()
	// Experimental.
	ResetSlack()
	// Experimental.
	ResetSnowflake()
	// Experimental.
	ResetTrendmicro()
	// Experimental.
	ResetVeeva()
	// Experimental.
	ResetZendesk()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Amplitude() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) AmplitudeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnector() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnectorInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Datadog() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DatadogInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Dynatrace() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DynatraceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalytics() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalyticsInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Honeycode() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) HoneycodeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexus() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexusInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfileCredentialsProperty {
	var returns *AwsConnectorProfile_ConnectorProfileCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Marketo() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) MarketoInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Redshift() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) RedshiftInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Salesforce() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SalesforceInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoData() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoDataInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNow() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNowInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Singular() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SingularInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Slack() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SlackInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Snowflake() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SnowflakeInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Trendmicro() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TrendmicroInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Veeva() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) VeevaInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Zendesk() AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference {
	var returns AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ZendeskInput() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfileCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutAmplitude(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutCustomConnector(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDatadog(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDynatrace(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutGoogleAnalytics(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutHoneycode(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutInforNexus(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutMarketo(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutRedshift(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSalesforce(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSapoData(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutServiceNow(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSingular(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSlack(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSnowflake(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutTrendmicro(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutVeeva(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutZendesk(value *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

