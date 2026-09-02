package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOntapVolume_RetentionPeriodPropertyOutputReference interface {
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
	DefaultRetention() TfOntapVolume_DefaultRetentionPropertyOutputReference
	// Experimental.
	DefaultRetentionInput() *TfOntapVolume_DefaultRetentionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfOntapVolume_RetentionPeriodProperty
	// Experimental.
	SetInternalValue(val *TfOntapVolume_RetentionPeriodProperty)
	// Experimental.
	MaximumRetention() TfOntapVolume_MaximumRetentionPropertyOutputReference
	// Experimental.
	MaximumRetentionInput() *TfOntapVolume_MaximumRetentionProperty
	// Experimental.
	MinimumRetention() TfOntapVolume_MinimumRetentionPropertyOutputReference
	// Experimental.
	MinimumRetentionInput() *TfOntapVolume_MinimumRetentionProperty
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
	PutDefaultRetention(value *TfOntapVolume_DefaultRetentionProperty)
	// Experimental.
	PutMaximumRetention(value *TfOntapVolume_MaximumRetentionProperty)
	// Experimental.
	PutMinimumRetention(value *TfOntapVolume_MinimumRetentionProperty)
	// Experimental.
	ResetDefaultRetention()
	// Experimental.
	ResetMaximumRetention()
	// Experimental.
	ResetMinimumRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOntapVolume_RetentionPeriodPropertyOutputReference
type jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetention() TfOntapVolume_DefaultRetentionPropertyOutputReference {
	var returns TfOntapVolume_DefaultRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetentionInput() *TfOntapVolume_DefaultRetentionProperty {
	var returns *TfOntapVolume_DefaultRetentionProperty
	_jsii_.Get(
		j,
		"defaultRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) InternalValue() *TfOntapVolume_RetentionPeriodProperty {
	var returns *TfOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetention() TfOntapVolume_MaximumRetentionPropertyOutputReference {
	var returns TfOntapVolume_MaximumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"maximumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetentionInput() *TfOntapVolume_MaximumRetentionProperty {
	var returns *TfOntapVolume_MaximumRetentionProperty
	_jsii_.Get(
		j,
		"maximumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetention() TfOntapVolume_MinimumRetentionPropertyOutputReference {
	var returns TfOntapVolume_MinimumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"minimumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetentionInput() *TfOntapVolume_MinimumRetentionProperty {
	var returns *TfOntapVolume_MinimumRetentionProperty
	_jsii_.Get(
		j,
		"minimumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOntapVolume_RetentionPeriodPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOntapVolume_RetentionPeriodPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOntapVolume_RetentionPeriodPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOntapVolume_RetentionPeriodPropertyOutputReference_Override(t TfOntapVolume_RetentionPeriodPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference)SetInternalValue(val *TfOntapVolume_RetentionPeriodProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) PutDefaultRetention(value *TfOntapVolume_DefaultRetentionProperty) {
	if err := t.validatePutDefaultRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultRetention",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) PutMaximumRetention(value *TfOntapVolume_MaximumRetentionProperty) {
	if err := t.validatePutMaximumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaximumRetention",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) PutMinimumRetention(value *TfOntapVolume_MinimumRetentionProperty) {
	if err := t.validatePutMinimumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMinimumRetention",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ResetDefaultRetention() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultRetention",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ResetMaximumRetention() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRetention",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ResetMinimumRetention() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumRetention",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOntapVolume_RetentionPeriodPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

