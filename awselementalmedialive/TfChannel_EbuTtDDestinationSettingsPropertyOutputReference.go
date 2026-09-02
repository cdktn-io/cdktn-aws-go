package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EbuTtDDestinationSettingsPropertyOutputReference interface {
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
	// Experimental.
	CopyrightHolder() *string
	// Experimental.
	SetCopyrightHolder(val *string)
	// Experimental.
	CopyrightHolderInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FillLineGap() *string
	// Experimental.
	SetFillLineGap(val *string)
	// Experimental.
	FillLineGapInput() *string
	// Experimental.
	FontFamily() *string
	// Experimental.
	SetFontFamily(val *string)
	// Experimental.
	FontFamilyInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_EbuTtDDestinationSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EbuTtDDestinationSettingsProperty)
	// Experimental.
	StyleControl() *string
	// Experimental.
	SetStyleControl(val *string)
	// Experimental.
	StyleControlInput() *string
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
	ResetCopyrightHolder()
	// Experimental.
	ResetFillLineGap()
	// Experimental.
	ResetFontFamily()
	// Experimental.
	ResetStyleControl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_EbuTtDDestinationSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) CopyrightHolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyrightHolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) CopyrightHolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyrightHolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) FillLineGap() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fillLineGap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) FillLineGapInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fillLineGapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) FontFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) FontFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fontFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) InternalValue() *TfChannel_EbuTtDDestinationSettingsProperty {
	var returns *TfChannel_EbuTtDDestinationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) StyleControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"styleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) StyleControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"styleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EbuTtDDestinationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EbuTtDDestinationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EbuTtDDestinationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EbuTtDDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EbuTtDDestinationSettingsPropertyOutputReference_Override(t TfChannel_EbuTtDDestinationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EbuTtDDestinationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetCopyrightHolder(val *string) {
	if err := j.validateSetCopyrightHolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyrightHolder",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetFillLineGap(val *string) {
	if err := j.validateSetFillLineGapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fillLineGap",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetFontFamily(val *string) {
	if err := j.validateSetFontFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fontFamily",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EbuTtDDestinationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetStyleControl(val *string) {
	if err := j.validateSetStyleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"styleControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ResetCopyrightHolder() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyrightHolder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ResetFillLineGap() {
	_jsii_.InvokeVoid(
		t,
		"resetFillLineGap",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ResetFontFamily() {
	_jsii_.InvokeVoid(
		t,
		"resetFontFamily",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ResetStyleControl() {
	_jsii_.InvokeVoid(
		t,
		"resetStyleControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EbuTtDDestinationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

