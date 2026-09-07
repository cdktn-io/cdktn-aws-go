package cloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTelemetryRule_DestinationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudtrailParameters() AwsTelemetryRule_CloudtrailParametersPropertyList
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
	ElbLoadBalancerLoggingParameters() AwsTelemetryRule_ElbLoadBalancerLoggingParametersPropertyList
	// Experimental.
	ElbLoadBalancerLoggingParametersInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LogDeliveryParameters() AwsTelemetryRule_LogDeliveryParametersPropertyList
	// Experimental.
	LogDeliveryParametersInput() interface{}
	// Experimental.
	MskMonitoringParameters() AwsTelemetryRule_MskMonitoringParametersPropertyList
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
	VpcFlowLogParameters() AwsTelemetryRule_VpcFlowLogParametersPropertyList
	// Experimental.
	VpcFlowLogParametersInput() interface{}
	// Experimental.
	WafLoggingParameters() AwsTelemetryRule_WafLoggingParametersPropertyList
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

// The jsii proxy struct for AwsTelemetryRule_DestinationConfigurationPropertyOutputReference
type jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) CloudtrailParameters() AwsTelemetryRule_CloudtrailParametersPropertyList {
	var returns AwsTelemetryRule_CloudtrailParametersPropertyList
	_jsii_.Get(
		j,
		"cloudtrailParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) CloudtrailParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudtrailParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) DestinationPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) DestinationPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) DestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) DestinationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ElbLoadBalancerLoggingParameters() AwsTelemetryRule_ElbLoadBalancerLoggingParametersPropertyList {
	var returns AwsTelemetryRule_ElbLoadBalancerLoggingParametersPropertyList
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ElbLoadBalancerLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbLoadBalancerLoggingParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) LogDeliveryParameters() AwsTelemetryRule_LogDeliveryParametersPropertyList {
	var returns AwsTelemetryRule_LogDeliveryParametersPropertyList
	_jsii_.Get(
		j,
		"logDeliveryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) LogDeliveryParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) MskMonitoringParameters() AwsTelemetryRule_MskMonitoringParametersPropertyList {
	var returns AwsTelemetryRule_MskMonitoringParametersPropertyList
	_jsii_.Get(
		j,
		"mskMonitoringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) MskMonitoringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mskMonitoringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) VpcFlowLogParameters() AwsTelemetryRule_VpcFlowLogParametersPropertyList {
	var returns AwsTelemetryRule_VpcFlowLogParametersPropertyList
	_jsii_.Get(
		j,
		"vpcFlowLogParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) VpcFlowLogParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcFlowLogParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) WafLoggingParameters() AwsTelemetryRule_WafLoggingParametersPropertyList {
	var returns AwsTelemetryRule_WafLoggingParametersPropertyList
	_jsii_.Get(
		j,
		"wafLoggingParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) WafLoggingParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wafLoggingParametersInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTelemetryRule_DestinationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTelemetryRule_DestinationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTelemetryRule_DestinationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsTelemetryRule.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTelemetryRule_DestinationConfigurationPropertyOutputReference_Override(a AwsTelemetryRule_DestinationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsTelemetryRule.DestinationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetDestinationPattern(val *string) {
	if err := j.validateSetDestinationPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPattern",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetDestinationType(val *string) {
	if err := j.validateSetDestinationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationType",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutCloudtrailParameters(value interface{}) {
	if err := a.validatePutCloudtrailParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudtrailParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutElbLoadBalancerLoggingParameters(value interface{}) {
	if err := a.validatePutElbLoadBalancerLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElbLoadBalancerLoggingParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutLogDeliveryParameters(value interface{}) {
	if err := a.validatePutLogDeliveryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDeliveryParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutMskMonitoringParameters(value interface{}) {
	if err := a.validatePutMskMonitoringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMskMonitoringParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutVpcFlowLogParameters(value interface{}) {
	if err := a.validatePutVpcFlowLogParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcFlowLogParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) PutWafLoggingParameters(value interface{}) {
	if err := a.validatePutWafLoggingParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWafLoggingParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetCloudtrailParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudtrailParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetDestinationPattern() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationPattern",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetDestinationType() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetElbLoadBalancerLoggingParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetElbLoadBalancerLoggingParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetLogDeliveryParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDeliveryParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetMskMonitoringParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetMskMonitoringParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetRetentionInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetVpcFlowLogParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcFlowLogParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ResetWafLoggingParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetWafLoggingParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTelemetryRule_DestinationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

