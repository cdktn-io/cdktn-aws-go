package fsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/fsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOntapVolume_RetentionPeriodPropertyOutputReference interface {
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
	DefaultRetention() AwsOntapVolume_DefaultRetentionPropertyOutputReference
	// Experimental.
	DefaultRetentionInput() *AwsOntapVolume_DefaultRetentionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOntapVolume_RetentionPeriodProperty
	// Experimental.
	SetInternalValue(val *AwsOntapVolume_RetentionPeriodProperty)
	// Experimental.
	MaximumRetention() AwsOntapVolume_MaximumRetentionPropertyOutputReference
	// Experimental.
	MaximumRetentionInput() *AwsOntapVolume_MaximumRetentionProperty
	// Experimental.
	MinimumRetention() AwsOntapVolume_MinimumRetentionPropertyOutputReference
	// Experimental.
	MinimumRetentionInput() *AwsOntapVolume_MinimumRetentionProperty
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
	PutDefaultRetention(value *AwsOntapVolume_DefaultRetentionProperty)
	// Experimental.
	PutMaximumRetention(value *AwsOntapVolume_MaximumRetentionProperty)
	// Experimental.
	PutMinimumRetention(value *AwsOntapVolume_MinimumRetentionProperty)
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

// The jsii proxy struct for AwsOntapVolume_RetentionPeriodPropertyOutputReference
type jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetention() AwsOntapVolume_DefaultRetentionPropertyOutputReference {
	var returns AwsOntapVolume_DefaultRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"defaultRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) DefaultRetentionInput() *AwsOntapVolume_DefaultRetentionProperty {
	var returns *AwsOntapVolume_DefaultRetentionProperty
	_jsii_.Get(
		j,
		"defaultRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) InternalValue() *AwsOntapVolume_RetentionPeriodProperty {
	var returns *AwsOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetention() AwsOntapVolume_MaximumRetentionPropertyOutputReference {
	var returns AwsOntapVolume_MaximumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"maximumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) MaximumRetentionInput() *AwsOntapVolume_MaximumRetentionProperty {
	var returns *AwsOntapVolume_MaximumRetentionProperty
	_jsii_.Get(
		j,
		"maximumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetention() AwsOntapVolume_MinimumRetentionPropertyOutputReference {
	var returns AwsOntapVolume_MinimumRetentionPropertyOutputReference
	_jsii_.Get(
		j,
		"minimumRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) MinimumRetentionInput() *AwsOntapVolume_MinimumRetentionProperty {
	var returns *AwsOntapVolume_MinimumRetentionProperty
	_jsii_.Get(
		j,
		"minimumRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOntapVolume_RetentionPeriodPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOntapVolume_RetentionPeriodPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOntapVolume_RetentionPeriodPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOntapVolume_RetentionPeriodPropertyOutputReference_Override(a AwsOntapVolume_RetentionPeriodPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsOntapVolume.RetentionPeriodPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference)SetInternalValue(val *AwsOntapVolume_RetentionPeriodProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) PutDefaultRetention(value *AwsOntapVolume_DefaultRetentionProperty) {
	if err := a.validatePutDefaultRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) PutMaximumRetention(value *AwsOntapVolume_MaximumRetentionProperty) {
	if err := a.validatePutMaximumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaximumRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) PutMinimumRetention(value *AwsOntapVolume_MinimumRetentionProperty) {
	if err := a.validatePutMinimumRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMinimumRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ResetDefaultRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ResetMaximumRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ResetMinimumRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOntapVolume_RetentionPeriodPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

