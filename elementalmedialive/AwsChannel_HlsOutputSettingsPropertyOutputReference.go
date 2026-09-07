package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_HlsOutputSettingsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	H265PackagingType() *string
	// Experimental.
	SetH265PackagingType(val *string)
	// Experimental.
	H265PackagingTypeInput() *string
	// Experimental.
	HlsSettings() AwsChannel_HlsSettingsPropertyOutputReference
	// Experimental.
	HlsSettingsInput() *AwsChannel_HlsSettingsProperty
	// Experimental.
	InternalValue() *AwsChannel_HlsOutputSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_HlsOutputSettingsProperty)
	// Experimental.
	NameModifier() *string
	// Experimental.
	SetNameModifier(val *string)
	// Experimental.
	NameModifierInput() *string
	// Experimental.
	SegmentModifier() *string
	// Experimental.
	SetSegmentModifier(val *string)
	// Experimental.
	SegmentModifierInput() *string
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
	PutHlsSettings(value *AwsChannel_HlsSettingsProperty)
	// Experimental.
	ResetH265PackagingType()
	// Experimental.
	ResetNameModifier()
	// Experimental.
	ResetSegmentModifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsChannel_HlsOutputSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) H265PackagingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"h265PackagingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) H265PackagingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"h265PackagingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) HlsSettings() AwsChannel_HlsSettingsPropertyOutputReference {
	var returns AwsChannel_HlsSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hlsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) HlsSettingsInput() *AwsChannel_HlsSettingsProperty {
	var returns *AwsChannel_HlsSettingsProperty
	_jsii_.Get(
		j,
		"hlsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) InternalValue() *AwsChannel_HlsOutputSettingsProperty {
	var returns *AwsChannel_HlsOutputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) NameModifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) NameModifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) SegmentModifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) SegmentModifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_HlsOutputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_HlsOutputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_HlsOutputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsOutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_HlsOutputSettingsPropertyOutputReference_Override(a AwsChannel_HlsOutputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsOutputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetH265PackagingType(val *string) {
	if err := j.validateSetH265PackagingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"h265PackagingType",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_HlsOutputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetNameModifier(val *string) {
	if err := j.validateSetNameModifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameModifier",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetSegmentModifier(val *string) {
	if err := j.validateSetSegmentModifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentModifier",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) PutHlsSettings(value *AwsChannel_HlsSettingsProperty) {
	if err := a.validatePutHlsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHlsSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ResetH265PackagingType() {
	_jsii_.InvokeVoid(
		a,
		"resetH265PackagingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ResetNameModifier() {
	_jsii_.InvokeVoid(
		a,
		"resetNameModifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ResetSegmentModifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSegmentModifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsOutputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

