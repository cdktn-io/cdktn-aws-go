package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_DestinationConnectorPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
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
	CustomConnector() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	// Experimental.
	CustomConnectorInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	// Experimental.
	CustomerProfiles() AwsFlow_CustomerProfilesPropertyOutputReference
	// Experimental.
	CustomerProfilesInput() *AwsFlow_CustomerProfilesProperty
	// Experimental.
	EventBridge() AwsFlow_EventBridgePropertyOutputReference
	// Experimental.
	EventBridgeInput() *AwsFlow_EventBridgeProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Honeycode() AwsFlow_HoneycodePropertyOutputReference
	// Experimental.
	HoneycodeInput() *AwsFlow_HoneycodeProperty
	// Experimental.
	InternalValue() *AwsFlow_DestinationConnectorPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_DestinationConnectorPropertiesProperty)
	// Experimental.
	LookoutMetrics() AwsFlow_LookoutMetricsPropertyOutputReference
	// Experimental.
	LookoutMetricsInput() *AwsFlow_LookoutMetricsProperty
	// Experimental.
	Marketo() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	// Experimental.
	MarketoInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	// Experimental.
	Redshift() AwsFlow_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsFlow_RedshiftProperty
	// Experimental.
	S3() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	// Experimental.
	Salesforce() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	// Experimental.
	SalesforceInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	// Experimental.
	SapoData() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	// Experimental.
	SapoDataInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	// Experimental.
	Snowflake() AwsFlow_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsFlow_SnowflakeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Upsolver() AwsFlow_UpsolverPropertyOutputReference
	// Experimental.
	UpsolverInput() *AwsFlow_UpsolverProperty
	// Experimental.
	Zendesk() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	// Experimental.
	ZendeskInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
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
	PutCustomConnector(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty)
	// Experimental.
	PutCustomerProfiles(value *AwsFlow_CustomerProfilesProperty)
	// Experimental.
	PutEventBridge(value *AwsFlow_EventBridgeProperty)
	// Experimental.
	PutHoneycode(value *AwsFlow_HoneycodeProperty)
	// Experimental.
	PutLookoutMetrics(value *AwsFlow_LookoutMetricsProperty)
	// Experimental.
	PutMarketo(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty)
	// Experimental.
	PutRedshift(value *AwsFlow_RedshiftProperty)
	// Experimental.
	PutS3(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property)
	// Experimental.
	PutSalesforce(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty)
	// Experimental.
	PutSapoData(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty)
	// Experimental.
	PutSnowflake(value *AwsFlow_SnowflakeProperty)
	// Experimental.
	PutUpsolver(value *AwsFlow_UpsolverProperty)
	// Experimental.
	PutZendesk(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty)
	// Experimental.
	ResetCustomConnector()
	// Experimental.
	ResetCustomerProfiles()
	// Experimental.
	ResetEventBridge()
	// Experimental.
	ResetHoneycode()
	// Experimental.
	ResetLookoutMetrics()
	// Experimental.
	ResetMarketo()
	// Experimental.
	ResetRedshift()
	// Experimental.
	ResetS3()
	// Experimental.
	ResetSalesforce()
	// Experimental.
	ResetSapoData()
	// Experimental.
	ResetSnowflake()
	// Experimental.
	ResetUpsolver()
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

// The jsii proxy struct for AwsFlow_DestinationConnectorPropertiesPropertyOutputReference
type jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnector() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorPropertyOutputReference
	_jsii_.Get(
		j,
		"customConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomConnectorInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty
	_jsii_.Get(
		j,
		"customConnectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfiles() AwsFlow_CustomerProfilesPropertyOutputReference {
	var returns AwsFlow_CustomerProfilesPropertyOutputReference
	_jsii_.Get(
		j,
		"customerProfiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) CustomerProfilesInput() *AwsFlow_CustomerProfilesProperty {
	var returns *AwsFlow_CustomerProfilesProperty
	_jsii_.Get(
		j,
		"customerProfilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridge() AwsFlow_EventBridgePropertyOutputReference {
	var returns AwsFlow_EventBridgePropertyOutputReference
	_jsii_.Get(
		j,
		"eventBridge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) EventBridgeInput() *AwsFlow_EventBridgeProperty {
	var returns *AwsFlow_EventBridgeProperty
	_jsii_.Get(
		j,
		"eventBridgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Honeycode() AwsFlow_HoneycodePropertyOutputReference {
	var returns AwsFlow_HoneycodePropertyOutputReference
	_jsii_.Get(
		j,
		"honeycode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) HoneycodeInput() *AwsFlow_HoneycodeProperty {
	var returns *AwsFlow_HoneycodeProperty
	_jsii_.Get(
		j,
		"honeycodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) InternalValue() *AwsFlow_DestinationConnectorPropertiesProperty {
	var returns *AwsFlow_DestinationConnectorPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetrics() AwsFlow_LookoutMetricsPropertyOutputReference {
	var returns AwsFlow_LookoutMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"lookoutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) LookoutMetricsInput() *AwsFlow_LookoutMetricsProperty {
	var returns *AwsFlow_LookoutMetricsProperty
	_jsii_.Get(
		j,
		"lookoutMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Marketo() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoPropertyOutputReference
	_jsii_.Get(
		j,
		"marketo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) MarketoInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty
	_jsii_.Get(
		j,
		"marketoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Redshift() AwsFlow_RedshiftPropertyOutputReference {
	var returns AwsFlow_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) RedshiftInput() *AwsFlow_RedshiftProperty {
	var returns *AwsFlow_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) S3() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) S3Input() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Salesforce() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforcePropertyOutputReference
	_jsii_.Get(
		j,
		"salesforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) SalesforceInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty
	_jsii_.Get(
		j,
		"salesforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoData() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataPropertyOutputReference
	_jsii_.Get(
		j,
		"sapoData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) SapoDataInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty
	_jsii_.Get(
		j,
		"sapoDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Snowflake() AwsFlow_SnowflakePropertyOutputReference {
	var returns AwsFlow_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) SnowflakeInput() *AwsFlow_SnowflakeProperty {
	var returns *AwsFlow_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Upsolver() AwsFlow_UpsolverPropertyOutputReference {
	var returns AwsFlow_UpsolverPropertyOutputReference
	_jsii_.Get(
		j,
		"upsolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) UpsolverInput() *AwsFlow_UpsolverProperty {
	var returns *AwsFlow_UpsolverProperty
	_jsii_.Get(
		j,
		"upsolverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Zendesk() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskPropertyOutputReference
	_jsii_.Get(
		j,
		"zendesk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ZendeskInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty
	_jsii_.Get(
		j,
		"zendeskInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_DestinationConnectorPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_DestinationConnectorPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_DestinationConnectorPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_DestinationConnectorPropertiesPropertyOutputReference_Override(a AwsFlow_DestinationConnectorPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationConnectorPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference)SetInternalValue(val *AwsFlow_DestinationConnectorPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomConnector(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesCustomConnectorProperty) {
	if err := a.validatePutCustomConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutCustomerProfiles(value *AwsFlow_CustomerProfilesProperty) {
	if err := a.validatePutCustomerProfilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerProfiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutEventBridge(value *AwsFlow_EventBridgeProperty) {
	if err := a.validatePutEventBridgeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridge",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutHoneycode(value *AwsFlow_HoneycodeProperty) {
	if err := a.validatePutHoneycodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHoneycode",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutLookoutMetrics(value *AwsFlow_LookoutMetricsProperty) {
	if err := a.validatePutLookoutMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLookoutMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutMarketo(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesMarketoProperty) {
	if err := a.validatePutMarketoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMarketo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutRedshift(value *AwsFlow_RedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutS3(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSalesforce(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSalesforceProperty) {
	if err := a.validatePutSalesforceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSalesforce",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSapoData(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesSapoDataProperty) {
	if err := a.validatePutSapoDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSapoData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutSnowflake(value *AwsFlow_SnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutUpsolver(value *AwsFlow_UpsolverProperty) {
	if err := a.validatePutUpsolverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUpsolver",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) PutZendesk(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskProperty) {
	if err := a.validatePutZendeskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZendesk",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetCustomerProfiles() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerProfiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetEventBridge() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetHoneycode() {
	_jsii_.InvokeVoid(
		a,
		"resetHoneycode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetLookoutMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetLookoutMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetMarketo() {
	_jsii_.InvokeVoid(
		a,
		"resetMarketo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSalesforce() {
	_jsii_.InvokeVoid(
		a,
		"resetSalesforce",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSapoData() {
	_jsii_.InvokeVoid(
		a,
		"resetSapoData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetUpsolver() {
	_jsii_.InvokeVoid(
		a,
		"resetUpsolver",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ResetZendesk() {
	_jsii_.InvokeVoid(
		a,
		"resetZendesk",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationConnectorPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

