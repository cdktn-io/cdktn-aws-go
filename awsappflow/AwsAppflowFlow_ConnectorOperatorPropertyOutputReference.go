package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_ConnectorOperatorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Amplitude() *string
	// Experimental.
	SetAmplitude(val *string)
	// Experimental.
	AmplitudeInput() *string
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
	CustomConnector() *string
	// Experimental.
	SetCustomConnector(val *string)
	// Experimental.
	CustomConnectorInput() *string
	// Experimental.
	Datadog() *string
	// Experimental.
	SetDatadog(val *string)
	// Experimental.
	DatadogInput() *string
	// Experimental.
	Dynatrace() *string
	// Experimental.
	SetDynatrace(val *string)
	// Experimental.
	DynatraceInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GoogleAnalytics() *string
	// Experimental.
	SetGoogleAnalytics(val *string)
	// Experimental.
	GoogleAnalyticsInput() *string
	// Experimental.
	InforNexus() *string
	// Experimental.
	SetInforNexus(val *string)
	// Experimental.
	InforNexusInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Marketo() *string
	// Experimental.
	SetMarketo(val *string)
	// Experimental.
	MarketoInput() *string
	// Experimental.
	S3() *string
	// Experimental.
	SetS3(val *string)
	// Experimental.
	S3Input() *string
	// Experimental.
	Salesforce() *string
	// Experimental.
	SetSalesforce(val *string)
	// Experimental.
	SalesforceInput() *string
	// Experimental.
	SapoData() *string
	// Experimental.
	SetSapoData(val *string)
	// Experimental.
	SapoDataInput() *string
	// Experimental.
	ServiceNow() *string
	// Experimental.
	SetServiceNow(val *string)
	// Experimental.
	ServiceNowInput() *string
	// Experimental.
	Singular() *string
	// Experimental.
	SetSingular(val *string)
	// Experimental.
	SingularInput() *string
	// Experimental.
	Slack() *string
	// Experimental.
	SetSlack(val *string)
	// Experimental.
	SlackInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trendmicro() *string
	// Experimental.
	SetTrendmicro(val *string)
	// Experimental.
	TrendmicroInput() *string
	// Experimental.
	Veeva() *string
	// Experimental.
	SetVeeva(val *string)
	// Experimental.
	VeevaInput() *string
	// Experimental.
	Zendesk() *string
	// Experimental.
	SetZendesk(val *string)
	// Experimental.
	ZendeskInput() *string
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

// The jsii proxy struct for AwsAppflowFlow_ConnectorOperatorPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Amplitude() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) AmplitudeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amplitudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) CustomConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) CustomConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Datadog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) DatadogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Dynatrace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatrace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) DynatraceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynatraceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GoogleAnalytics() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GoogleAnalyticsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) InforNexus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) InforNexusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inforNexusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Marketo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) MarketoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) S3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) S3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Salesforce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) SalesforceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) SapoData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) SapoDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ServiceNow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ServiceNowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Singular() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singular",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) SingularInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"singularInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Slack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) SlackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Trendmicro() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicro",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) TrendmicroInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trendmicroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Veeva() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veeva",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) VeevaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"veevaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Zendesk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ZendeskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_ConnectorOperatorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAppflowFlow_ConnectorOperatorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_ConnectorOperatorPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.ConnectorOperatorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_ConnectorOperatorPropertyOutputReference_Override(a AwsAppflowFlow_ConnectorOperatorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.ConnectorOperatorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetAmplitude(val *string) {
	if err := j.validateSetAmplitudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amplitude",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetCustomConnector(val *string) {
	if err := j.validateSetCustomConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customConnector",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetDatadog(val *string) {
	if err := j.validateSetDatadogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datadog",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetDynatrace(val *string) {
	if err := j.validateSetDynatraceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynatrace",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetGoogleAnalytics(val *string) {
	if err := j.validateSetGoogleAnalyticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAnalytics",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetInforNexus(val *string) {
	if err := j.validateSetInforNexusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inforNexus",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetMarketo(val *string) {
	if err := j.validateSetMarketoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketo",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetS3(val *string) {
	if err := j.validateSetS3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetSalesforce(val *string) {
	if err := j.validateSetSalesforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforce",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetSapoData(val *string) {
	if err := j.validateSetSapoDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapoData",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetServiceNow(val *string) {
	if err := j.validateSetServiceNowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNow",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetSingular(val *string) {
	if err := j.validateSetSingularParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singular",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetSlack(val *string) {
	if err := j.validateSetSlackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slack",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetTrendmicro(val *string) {
	if err := j.validateSetTrendmicroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trendmicro",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetVeeva(val *string) {
	if err := j.validateSetVeevaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"veeva",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference)SetZendesk(val *string) {
	if err := j.validateSetZendeskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zendesk",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetAmplitude() {
	_jsii_.InvokeVoid(
		a,
		"resetAmplitude",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetDynatrace() {
	_jsii_.InvokeVoid(
		a,
		"resetDynatrace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetGoogleAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetInforNexus() {
	_jsii_.InvokeVoid(
		a,
		"resetInforNexus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetSingular() {
	_jsii_.InvokeVoid(
		a,
		"resetSingular",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetSlack() {
	_jsii_.InvokeVoid(
		a,
		"resetSlack",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetTrendmicro() {
	_jsii_.InvokeVoid(
		a,
		"resetTrendmicro",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetVeeva() {
	_jsii_.InvokeVoid(
		a,
		"resetVeeva",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_ConnectorOperatorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

