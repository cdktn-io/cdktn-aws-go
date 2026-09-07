package sesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchDestination() AwsConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference
	// Experimental.
	CloudWatchDestinationInput() *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	EventBridgeDestination() AwsConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference
	// Experimental.
	EventBridgeDestinationInput() *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConfigurationSetEventDestination_EventDestinationProperty
	// Experimental.
	SetInternalValue(val *AwsConfigurationSetEventDestination_EventDestinationProperty)
	// Experimental.
	KinesisFirehoseDestination() AwsConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference
	// Experimental.
	KinesisFirehoseDestinationInput() *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty
	// Experimental.
	MatchingEventTypes() *[]*string
	// Experimental.
	SetMatchingEventTypes(val *[]*string)
	// Experimental.
	MatchingEventTypesInput() *[]*string
	// Experimental.
	PinpointDestination() AwsConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference
	// Experimental.
	PinpointDestinationInput() *AwsConfigurationSetEventDestination_PinpointDestinationProperty
	// Experimental.
	SnsDestination() AwsConfigurationSetEventDestination_SnsDestinationPropertyOutputReference
	// Experimental.
	SnsDestinationInput() *AwsConfigurationSetEventDestination_SnsDestinationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCloudWatchDestination(value *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty)
	// Experimental.
	PutEventBridgeDestination(value *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty)
	// Experimental.
	PutKinesisFirehoseDestination(value *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty)
	// Experimental.
	PutPinpointDestination(value *AwsConfigurationSetEventDestination_PinpointDestinationProperty)
	// Experimental.
	PutSnsDestination(value *AwsConfigurationSetEventDestination_SnsDestinationProperty)
	// Experimental.
	ResetCloudWatchDestination()
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetEventBridgeDestination()
	// Experimental.
	ResetKinesisFirehoseDestination()
	// Experimental.
	ResetPinpointDestination()
	// Experimental.
	ResetSnsDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference
type jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CloudWatchDestination() AwsConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference {
	var returns AwsConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CloudWatchDestinationInput() *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty
	_jsii_.Get(
		j,
		"cloudWatchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EventBridgeDestination() AwsConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference {
	var returns AwsConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"eventBridgeDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EventBridgeDestinationInput() *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty
	_jsii_.Get(
		j,
		"eventBridgeDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InternalValue() *AwsConfigurationSetEventDestination_EventDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_EventDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) KinesisFirehoseDestination() AwsConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference {
	var returns AwsConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) KinesisFirehoseDestinationInput() *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) MatchingEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) MatchingEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PinpointDestination() AwsConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference {
	var returns AwsConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"pinpointDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PinpointDestinationInput() *AwsConfigurationSetEventDestination_PinpointDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_PinpointDestinationProperty
	_jsii_.Get(
		j,
		"pinpointDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) SnsDestination() AwsConfigurationSetEventDestination_SnsDestinationPropertyOutputReference {
	var returns AwsConfigurationSetEventDestination_SnsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"snsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) SnsDestinationInput() *AwsConfigurationSetEventDestination_SnsDestinationProperty {
	var returns *AwsConfigurationSetEventDestination_SnsDestinationProperty
	_jsii_.Get(
		j,
		"snsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConfigurationSetEventDestination_EventDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsConfigurationSetEventDestination.EventDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference_Override(a AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsConfigurationSetEventDestination.EventDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetInternalValue(val *AwsConfigurationSetEventDestination_EventDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetMatchingEventTypes(val *[]*string) {
	if err := j.validateSetMatchingEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingEventTypes",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutCloudWatchDestination(value *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty) {
	if err := a.validatePutCloudWatchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudWatchDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutEventBridgeDestination(value *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty) {
	if err := a.validatePutEventBridgeDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventBridgeDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutKinesisFirehoseDestination(value *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty) {
	if err := a.validatePutKinesisFirehoseDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehoseDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutPinpointDestination(value *AwsConfigurationSetEventDestination_PinpointDestinationProperty) {
	if err := a.validatePutPinpointDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPinpointDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutSnsDestination(value *AwsConfigurationSetEventDestination_SnsDestinationProperty) {
	if err := a.validatePutSnsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnsDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetCloudWatchDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudWatchDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetEventBridgeDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetEventBridgeDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetKinesisFirehoseDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehoseDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetPinpointDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetPinpointDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetSnsDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetSnsDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

