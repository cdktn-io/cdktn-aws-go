package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference interface {
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
	DefaultRetention() AwsFsxOntapVolume_DefaultRetentionPropertyOutputReference
	// Experimental.
	DefaultRetentionInput() *AwsFsxOntapVolume_DefaultRetentionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsFsxOntapVolume_RetentionPeriodProperty
	// Experimental.
	SetInternalValue(val *AwsFsxOntapVolume_RetentionPeriodProperty)
	// Experimental.
	MaximumRetention() AwsFsxOntapVolume_MaximumRetentionPropertyOutputReference
	// Experimental.
	MaximumRetentionInput() *AwsFsxOntapVolume_MaximumRetentionProperty
	// Experimental.
	MinimumRetention() AwsFsxOntapVolume_MinimumRetentionPropertyOutputReference
	// Experimental.
	MinimumRetentionInput() *AwsFsxOntapVolume_MinimumRetentionProperty
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
	PutDefaultRetention(value *AwsFsxOntapVolume_DefaultRetentionProperty)
	// Experimental.
	PutMaximumRetention(value *AwsFsxOntapVolume_MaximumRetentionProperty)
	// Experimental.
	PutMinimumRetention(value *AwsFsxOntapVolume_MinimumRetentionProperty)
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

// The jsii proxy struct for AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference
type jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetention() AwsFsxOntapVolume_DefaultRetentionPropertyOutputReference {
	var returns AwsFsxOntapVolume_DefaultRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetentionInput() *AwsFsxOntapVolume_DefaultRetentionProperty {
	var returns *AwsFsxOntapVolume_DefaultRetentionProperty
	_jsii_.Get(
		j,
		"defaultRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) InternalValue() *AwsFsxOntapVolume_RetentionPeriodProperty {
	var returns *AwsFsxOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetention() AwsFsxOntapVolume_MaximumRetentionPropertyOutputReference {
	var returns AwsFsxOntapVolume_MaximumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"maximumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetentionInput() *AwsFsxOntapVolume_MaximumRetentionProperty {
	var returns *AwsFsxOntapVolume_MaximumRetentionProperty
	_jsii_.Get(
		j,
		"maximumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetention() AwsFsxOntapVolume_MinimumRetentionPropertyOutputReference {
	var returns AwsFsxOntapVolume_MinimumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"minimumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetentionInput() *AwsFsxOntapVolume_MinimumRetentionProperty {
	var returns *AwsFsxOntapVolume_MinimumRetentionProperty
	_jsii_.Get(
		j,
		"minimumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFsxOntapVolume_RetentionPeriodPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFsxOntapVolume_RetentionPeriodPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFsxOntapVolume_RetentionPeriodPropertyOutputReference_Override(a AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference)SetInternalValue(val *AwsFsxOntapVolume_RetentionPeriodProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) PutDefaultRetention(value *AwsFsxOntapVolume_DefaultRetentionProperty) {
	if err := a.validatePutDefaultRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) PutMaximumRetention(value *AwsFsxOntapVolume_MaximumRetentionProperty) {
	if err := a.validatePutMaximumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaximumRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) PutMinimumRetention(value *AwsFsxOntapVolume_MinimumRetentionProperty) {
	if err := a.validatePutMinimumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMinimumRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ResetDefaultRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ResetMaximumRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ResetMinimumRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

