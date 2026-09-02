package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
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
	CustomConnector() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	// Experimental.
	Datadog() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	// Experimental.
	Dynatrace() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	// Experimental.
	Honeycode() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	// Experimental.
	InforNexus() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	// Experimental.
	InternalValue() *TfConnectorProfile_ConnectorProfileCredentialsProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfileCredentialsProperty)
	// Experimental.
	Marketo() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	// Experimental.
	Redshift() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	// Experimental.
	Salesforce() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	// Experimental.
	SapoData() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	// Experimental.
	ServiceNow() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	// Experimental.
	Singular() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	// Experimental.
	SingularInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	// Experimental.
	Slack() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	// Experimental.
	SlackInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	// Experimental.
	Snowflake() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	// Experimental.
	Veeva() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	// Experimental.
	Zendesk() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
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
	PutAmplitude(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty)
	// Experimental.
	PutDynatrace(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty)
	// Experimental.
	PutHoneycode(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty)
	// Experimental.
	PutInforNexus(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty)
	// Experimental.
	PutMarketo(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty)
	// Experimental.
	PutRedshift(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty)
	// Experimental.
	PutSalesforce(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty)
	// Experimental.
	PutSapoData(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty)
	// Experimental.
	PutServiceNow(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty)
	// Experimental.
	PutSingular(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty)
	// Experimental.
	PutSlack(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty)
	// Experimental.
	PutSnowflake(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty)
	// Experimental.
	PutTrendmicro(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty)
	// Experimental.
	PutVeeva(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty)
	// Experimental.
	PutZendesk(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty)
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

// The jsii proxy struct for TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Amplitude() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) AmplitudeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnector() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) CustomConnectorInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Datadog() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DatadogInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Dynatrace() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) DynatraceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalytics() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GoogleAnalyticsInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Honeycode() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) HoneycodeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexus() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InforNexusInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfileCredentialsProperty {
	var returns *TfConnectorProfile_ConnectorProfileCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Marketo() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) MarketoInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Redshift() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) RedshiftInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Salesforce() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SalesforceInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoData() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SapoDataInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNow() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ServiceNowInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Singular() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SingularInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Slack() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SlackInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Snowflake() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) SnowflakeInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Trendmicro() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) TrendmicroInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Veeva() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) VeevaInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Zendesk() TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference {
	var returns TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ZendeskInput() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfileCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutAmplitude(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsAmplitudeProperty) {
	if err := t.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutCustomConnector(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty) {
	if err := t.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDatadog(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDatadogProperty) {
	if err := t.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatadog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutDynatrace(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsDynatraceProperty) {
	if err := t.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutGoogleAnalytics(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsProperty) {
	if err := t.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutHoneycode(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsHoneycodeProperty) {
	if err := t.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutInforNexus(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsInforNexusProperty) {
	if err := t.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutMarketo(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsMarketoProperty) {
	if err := t.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMarketo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutRedshift(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsRedshiftProperty) {
	if err := t.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshift",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSalesforce(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSalesforceProperty) {
	if err := t.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSapoData(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSapoDataProperty) {
	if err := t.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSapoData",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutServiceNow(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsServiceNowProperty) {
	if err := t.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSingular(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSingularProperty) {
	if err := t.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSingular",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSlack(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSlackProperty) {
	if err := t.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlack",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutSnowflake(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsSnowflakeProperty) {
	if err := t.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutTrendmicro(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsTrendmicroProperty) {
	if err := t.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutVeeva(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsVeevaProperty) {
	if err := t.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVeeva",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) PutZendesk(value *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsZendeskProperty) {
	if err := t.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putZendesk",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		t,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		t,
		"resetDatadog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		t,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		t,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		t,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		t,
		"resetMarketo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshift",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		t,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		t,
		"resetSapoData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		t,
		"resetSingular",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		t,
		"resetSlack",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		t,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		t,
		"resetVeeva",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		t,
		"resetZendesk",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

