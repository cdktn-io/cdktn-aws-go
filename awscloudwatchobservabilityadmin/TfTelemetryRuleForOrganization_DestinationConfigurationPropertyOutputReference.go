package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudtrailParameters() TfTelemetryRuleForOrganization_CloudtrailParametersPropertyList
	// Experimental.
	CloudtrailParametersInput() interface{}
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
	DestinationPattern() *string
	// Experimental.
	SetDestinationPattern(val *string)
	// Experimental.
	DestinationPatternInput() *string
	// Experimental.
	DestinationType() *string
	// Experimental.
	SetDestinationType(val *string)
	// Experimental.
	DestinationTypeInput() *string
	// Experimental.
	ElbLoadBalancerLoggingParameters() TfTelemetryRuleForOrganization_ElbLoadBalancerLoggingParametersPropertyList
	// Experimental.
	ElbLoadBalancerLoggingParametersInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogDeliveryParameters() TfTelemetryRuleForOrganization_LogDeliveryParametersPropertyList
	// Experimental.
	LogDeliveryParametersInput() interface{}
	// Experimental.
	MskMonitoringParameters() TfTelemetryRuleForOrganization_MskMonitoringParametersPropertyList
	// Experimental.
	MskMonitoringParametersInput() interface{}
	// Experimental.
	RetentionInDays() *float64
	// Experimental.
	SetRetentionInDays(val *float64)
	// Experimental.
	RetentionInDaysInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcFlowLogParameters() TfTelemetryRuleForOrganization_VpcFlowLogParametersPropertyList
	// Experimental.
	VpcFlowLogParametersInput() interface{}
	// Experimental.
	WafLoggingParameters() TfTelemetryRuleForOrganization_WafLoggingParametersPropertyList
	// Experimental.
	WafLoggingParametersInput() interface{}
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
	PutCloudtrailParameters(value interface{})
	// Experimental.
	PutElbLoadBalancerLoggingParameters(value interface{})
	// Experimental.
	PutLogDeliveryParameters(value interface{})
	// Experimental.
	PutMskMonitoringParameters(value interface{})
	// Experimental.
	PutVpcFlowLogParameters(value interface{})
	// Experimental.
	PutWafLoggingParameters(value interface{})
	// Experimental.
	ResetCloudtrailParameters()
	// Experimental.
	ResetDestinationPattern()
	// Experimental.
	ResetDestinationType()
	// Experimental.
	ResetElbLoadBalancerLoggingParameters()
	// Experimental.
	ResetLogDeliveryParameters()
	// Experimental.
	ResetMskMonitoringParameters()
	// Experimental.
	ResetRetentionInDays()
	// Experimental.
	ResetVpcFlowLogParameters()
	// Experimental.
	ResetWafLoggingParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference
type jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) CloudtrailParameters() TfTelemetryRuleForOrganization_CloudtrailParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_CloudtrailParametersPropertyList
	_jsii_.Get(
		j,
		"cloudtrailParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) CloudtrailParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudtrailParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) DestinationPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) DestinationPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ElbLoadBalancerLoggingParameters() TfTelemetryRuleForOrganization_ElbLoadBalancerLoggingParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_ElbLoadBalancerLoggingParametersPropertyList
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ElbLoadBalancerLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) LogDeliveryParameters() TfTelemetryRuleForOrganization_LogDeliveryParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_LogDeliveryParametersPropertyList
	_jsii_.Get(
		j,
		"logDeliveryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) LogDeliveryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) MskMonitoringParameters() TfTelemetryRuleForOrganization_MskMonitoringParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_MskMonitoringParametersPropertyList
	_jsii_.Get(
		j,
		"mskMonitoringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) MskMonitoringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mskMonitoringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) VpcFlowLogParameters() TfTelemetryRuleForOrganization_VpcFlowLogParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_VpcFlowLogParametersPropertyList
	_jsii_.Get(
		j,
		"vpcFlowLogParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) VpcFlowLogParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcFlowLogParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) WafLoggingParameters() TfTelemetryRuleForOrganization_WafLoggingParametersPropertyList {
	var returns TfTelemetryRuleForOrganization_WafLoggingParametersPropertyList
	_jsii_.Get(
		j,
		"wafLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) WafLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafLoggingParametersInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfTelemetryRuleForOrganization.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference_Override(t TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.TfTelemetryRuleForOrganization.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetDestinationPattern(val *string) {
	if err := j.validateSetDestinationPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPattern",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutCloudtrailParameters(value interface{}) {
	if err := t.validatePutCloudtrailParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudtrailParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutElbLoadBalancerLoggingParameters(value interface{}) {
	if err := t.validatePutElbLoadBalancerLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElbLoadBalancerLoggingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutLogDeliveryParameters(value interface{}) {
	if err := t.validatePutLogDeliveryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogDeliveryParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutMskMonitoringParameters(value interface{}) {
	if err := t.validatePutMskMonitoringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMskMonitoringParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutVpcFlowLogParameters(value interface{}) {
	if err := t.validatePutVpcFlowLogParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcFlowLogParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) PutWafLoggingParameters(value interface{}) {
	if err := t.validatePutWafLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWafLoggingParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetCloudtrailParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudtrailParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetDestinationPattern() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationPattern",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetDestinationType() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetElbLoadBalancerLoggingParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetElbLoadBalancerLoggingParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetLogDeliveryParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetLogDeliveryParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetMskMonitoringParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetMskMonitoringParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		t,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetVpcFlowLogParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcFlowLogParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ResetWafLoggingParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetWafLoggingParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTelemetryRuleForOrganization_DestinationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

