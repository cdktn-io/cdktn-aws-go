package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() AwsAppflowFlow_AmplitudePropertyOutputReference
	// Experimental.
	AmplitudeInput() *AwsAppflowFlow_AmplitudeProperty
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
	CustomConnector() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	// Experimental.
	Datadog() AwsAppflowFlow_DatadogPropertyOutputReference
	// Experimental.
	DatadogInput() *AwsAppflowFlow_DatadogProperty
	// Experimental.
	Dynatrace() AwsAppflowFlow_DynatracePropertyOutputReference
	// Experimental.
	DynatraceInput() *AwsAppflowFlow_DynatraceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() AwsAppflowFlow_GoogleAnalyticsPropertyOutputReference
	// Experimental.
	GoogleAnalyticsInput() *AwsAppflowFlow_GoogleAnalyticsProperty
	// Experimental.
	InforNexus() AwsAppflowFlow_InforNexusPropertyOutputReference
	// Experimental.
	InforNexusInput() *AwsAppflowFlow_InforNexusProperty
	// Experimental.
	InternalValue() *AwsAppflowFlow_SourceConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowFlow_SourceConnectorPropertiesProperty)
	// Experimental.
	Marketo() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	// Experimental.
	S3() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	// Experimental.
	Salesforce() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	// Experimental.
	ServiceNow() AwsAppflowFlow_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsAppflowFlow_ServiceNowProperty
	// Experimental.
	Singular() AwsAppflowFlow_SingularPropertyOutputReference
	// Experimental.
	SingularInput() *AwsAppflowFlow_SingularProperty
	// Experimental.
	Slack() AwsAppflowFlow_SlackPropertyOutputReference
	// Experimental.
	SlackInput() *AwsAppflowFlow_SlackProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() AwsAppflowFlow_TrendmicroPropertyOutputReference
	// Experimental.
	TrendmicroInput() *AwsAppflowFlow_TrendmicroProperty
	// Experimental.
	Veeva() AwsAppflowFlow_VeevaPropertyOutputReference
	// Experimental.
	VeevaInput() *AwsAppflowFlow_VeevaProperty
	// Experimental.
	Zendesk() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
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
	PutAmplitude(value *AwsAppflowFlow_AmplitudeProperty)
	// Experimental.
	PutCustomConnector(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutDatadog(value *AwsAppflowFlow_DatadogProperty)
	// Experimental.
	PutDynatrace(value *AwsAppflowFlow_DynatraceProperty)
	// Experimental.
	PutGoogleAnalytics(value *AwsAppflowFlow_GoogleAnalyticsProperty)
	// Experimental.
	PutInforNexus(value *AwsAppflowFlow_InforNexusProperty)
	// Experimental.
	PutMarketo(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty)
	// Experimental.
	PutS3(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutServiceNow(value *AwsAppflowFlow_ServiceNowProperty)
	// Experimental.
	PutSingular(value *AwsAppflowFlow_SingularProperty)
	// Experimental.
	PutSlack(value *AwsAppflowFlow_SlackProperty)
	// Experimental.
	PutTrendmicro(value *AwsAppflowFlow_TrendmicroProperty)
	// Experimental.
	PutVeeva(value *AwsAppflowFlow_VeevaProperty)
	// Experimental.
	PutZendesk(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty)
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
	ResetInforNexus()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetS3()
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

// The jsii proxy struct for AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Amplitude() AwsAppflowFlow_AmplitudePropertyOutputReference {
	var returns AwsAppflowFlow_AmplitudePropertyOutputReference
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) AmplitudeInput() *AwsAppflowFlow_AmplitudeProperty {
	var returns *AwsAppflowFlow_AmplitudeProperty
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnector() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Datadog() AwsAppflowFlow_DatadogPropertyOutputReference {
	var returns AwsAppflowFlow_DatadogPropertyOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) DatadogInput() *AwsAppflowFlow_DatadogProperty {
	var returns *AwsAppflowFlow_DatadogProperty
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Dynatrace() AwsAppflowFlow_DynatracePropertyOutputReference {
	var returns AwsAppflowFlow_DynatracePropertyOutputReference
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) DynatraceInput() *AwsAppflowFlow_DynatraceProperty {
	var returns *AwsAppflowFlow_DynatraceProperty
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalytics() AwsAppflowFlow_GoogleAnalyticsPropertyOutputReference {
	var returns AwsAppflowFlow_GoogleAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GoogleAnalyticsInput() *AwsAppflowFlow_GoogleAnalyticsProperty {
	var returns *AwsAppflowFlow_GoogleAnalyticsProperty
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexus() AwsAppflowFlow_InforNexusPropertyOutputReference {
	var returns AwsAppflowFlow_InforNexusPropertyOutputReference
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) InforNexusInput() *AwsAppflowFlow_InforNexusProperty {
	var returns *AwsAppflowFlow_InforNexusProperty
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) InternalValue() *AwsAppflowFlow_SourceConnectorPropertiesProperty {
	var returns *AwsAppflowFlow_SourceConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Marketo() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) MarketoInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) S3() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) S3Input() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Salesforce() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) SalesforceInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) SapoData() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) SapoDataInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNow() AwsAppflowFlow_ServiceNowPropertyOutputReference {
	var returns AwsAppflowFlow_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ServiceNowInput() *AwsAppflowFlow_ServiceNowProperty {
	var returns *AwsAppflowFlow_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Singular() AwsAppflowFlow_SingularPropertyOutputReference {
	var returns AwsAppflowFlow_SingularPropertyOutputReference
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) SingularInput() *AwsAppflowFlow_SingularProperty {
	var returns *AwsAppflowFlow_SingularProperty
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Slack() AwsAppflowFlow_SlackPropertyOutputReference {
	var returns AwsAppflowFlow_SlackPropertyOutputReference
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) SlackInput() *AwsAppflowFlow_SlackProperty {
	var returns *AwsAppflowFlow_SlackProperty
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Trendmicro() AwsAppflowFlow_TrendmicroPropertyOutputReference {
	var returns AwsAppflowFlow_TrendmicroPropertyOutputReference
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) TrendmicroInput() *AwsAppflowFlow_TrendmicroProperty {
	var returns *AwsAppflowFlow_TrendmicroProperty
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Veeva() AwsAppflowFlow_VeevaPropertyOutputReference {
	var returns AwsAppflowFlow_VeevaPropertyOutputReference
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) VeevaInput() *AwsAppflowFlow_VeevaProperty {
	var returns *AwsAppflowFlow_VeevaProperty
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Zendesk() AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference {
	var returns AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ZendeskInput() *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty {
	var returns *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference_Override(a AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.SourceConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference)SetInternalValue(val *AwsAppflowFlow_SourceConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutAmplitude(value *AwsAppflowFlow_AmplitudeProperty) {
	if err := a.validatePutAmplitudeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmplitude",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutDatadog(value *AwsAppflowFlow_DatadogProperty) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutDynatrace(value *AwsAppflowFlow_DynatraceProperty) {
	if err := a.validatePutDynatraceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynatrace",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutGoogleAnalytics(value *AwsAppflowFlow_GoogleAnalyticsProperty) {
	if err := a.validatePutGoogleAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutInforNexus(value *AwsAppflowFlow_InforNexusProperty) {
	if err := a.validatePutInforNexusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInforNexus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutMarketo(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutS3(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutSalesforce(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutSapoData(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutServiceNow(value *AwsAppflowFlow_ServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutSingular(value *AwsAppflowFlow_SingularProperty) {
	if err := a.validatePutSingularParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSingular",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutSlack(value *AwsAppflowFlow_SlackProperty) {
	if err := a.validatePutSlackParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlack",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutTrendmicro(value *AwsAppflowFlow_TrendmicroProperty) {
	if err := a.validatePutTrendmicroParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrendmicro",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutVeeva(value *AwsAppflowFlow_VeevaProperty) {
	if err := a.validatePutVeevaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVeeva",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) PutZendesk(value *AwsAppflowFlow_SourceFlowConfigSourceConnectorPropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_SourceConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

