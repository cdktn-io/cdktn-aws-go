package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsMedialiveChannel_NielsenWatermarksSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_NielsenWatermarksSettingsProperty)
	// Experimental.
	NielsenCbetSettings() AwsMedialiveChannel_NielsenCbetSettingsPropertyOutputReference
	// Experimental.
	NielsenCbetSettingsInput() *AwsMedialiveChannel_NielsenCbetSettingsProperty
	// Experimental.
	NielsenDistributionType() *string
	// Experimental.
	SetNielsenDistributionType(val *string)
	// Experimental.
	NielsenDistributionTypeInput() *string
	// Experimental.
	NielsenNaesIiNwSettings() AwsMedialiveChannel_NielsenNaesIiNwSettingsPropertyList
	// Experimental.
	NielsenNaesIiNwSettingsInput() interface{}
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
	PutNielsenCbetSettings(value *AwsMedialiveChannel_NielsenCbetSettingsProperty)
	// Experimental.
	PutNielsenNaesIiNwSettings(value interface{})
	// Experimental.
	ResetNielsenCbetSettings()
	// Experimental.
	ResetNielsenDistributionType()
	// Experimental.
	ResetNielsenNaesIiNwSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_NielsenWatermarksSettingsProperty {
	var returns *AwsMedialiveChannel_NielsenWatermarksSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenCbetSettings() AwsMedialiveChannel_NielsenCbetSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_NielsenCbetSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"nielsenCbetSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenCbetSettingsInput() *AwsMedialiveChannel_NielsenCbetSettingsProperty {
	var returns *AwsMedialiveChannel_NielsenCbetSettingsProperty
	_jsii_.Get(
		j,
		"nielsenCbetSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenDistributionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenDistributionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenDistributionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenDistributionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenNaesIiNwSettings() AwsMedialiveChannel_NielsenNaesIiNwSettingsPropertyList {
	var returns AwsMedialiveChannel_NielsenNaesIiNwSettingsPropertyList
	_jsii_.Get(
		j,
		"nielsenNaesIiNwSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) NielsenNaesIiNwSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nielsenNaesIiNwSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.NielsenWatermarksSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.NielsenWatermarksSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_NielsenWatermarksSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetNielsenDistributionType(val *string) {
	if err := j.validateSetNielsenDistributionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nielsenDistributionType",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) PutNielsenCbetSettings(value *AwsMedialiveChannel_NielsenCbetSettingsProperty) {
	if err := a.validatePutNielsenCbetSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNielsenCbetSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) PutNielsenNaesIiNwSettings(value interface{}) {
	if err := a.validatePutNielsenNaesIiNwSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNielsenNaesIiNwSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenCbetSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenCbetSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenDistributionType() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenDistributionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ResetNielsenNaesIiNwSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenNaesIiNwSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_NielsenWatermarksSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

