package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference interface {
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
	DestinationRefId() *string
	// Experimental.
	SetDestinationRefId(val *string)
	// Experimental.
	DestinationRefIdInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference
type jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) DestinationRefId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRefId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) DestinationRefIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRefIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) InternalValue() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty {
	var returns *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference_Override(a AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetDestinationRefId(val *string) {
	if err := j.validateSetDestinationRefIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRefId",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetInternalValue(val *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

