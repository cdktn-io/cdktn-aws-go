package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
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
	CustomConnector() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	// Experimental.
	Datadog() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	// Experimental.
	Dynatrace() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	// Experimental.
	Honeycode() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	// Experimental.
	InforNexus() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	// Experimental.
	InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty)
	// Experimental.
	Marketo() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	// Experimental.
	Redshift() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	// Experimental.
	Salesforce() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	// Experimental.
	SapoData() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	// Experimental.
	ServiceNow() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	// Experimental.
	Singular() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	// Experimental.
	Slack() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	// Experimental.
	Snowflake() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	// Experimental.
	Veeva() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	// Experimental.
	Zendesk() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
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
	PutAmplitude(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty)
	// Experimental.
	PutSalesforce(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty)
	// Experimental.
	PutSlack(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty)
	// Experimental.
	PutSnowflake(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty)
	// Experimental.
	PutZendesk(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty)
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

// The jsii proxy struct for AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference
type jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Amplitude() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) AmplitudeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnector() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnectorInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Datadog() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DatadogInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Dynatrace() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DynatraceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalytics() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalyticsInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Honeycode() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) HoneycodeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexus() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexusInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InternalValue() *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Marketo() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) MarketoInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Redshift() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) RedshiftInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Salesforce() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SalesforceInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoData() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoDataInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNow() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNowInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Singular() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SingularInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Slack() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SlackInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Snowflake() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SnowflakeInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Trendmicro() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TrendmicroInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Veeva() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) VeevaInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Zendesk() AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference {
	var returns AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ZendeskInput() *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty {
	var returns *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference_Override(a AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetInternalValue(val *AwsAppflowConnectorProfile_ConnectorProfileCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutAmplitude(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutCustomConnector(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDatadog(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDynatrace(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutGoogleAnalytics(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutHoneycode(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutInforNexus(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutMarketo(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutRedshift(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSalesforce(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSapoData(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutServiceNow(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSingular(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSlack(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSnowflake(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutTrendmicro(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutVeeva(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutZendesk(value *AwsAppflowConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

