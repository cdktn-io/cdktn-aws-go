package awssesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationSet_VdmOptionsPropertyOutputReference interface {
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
	DashboardOptions() TfConfigurationSet_DashboardOptionsPropertyOutputReference
	// Experimental.
	DashboardOptionsInput() *TfConfigurationSet_DashboardOptionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GuardianOptions() TfConfigurationSet_GuardianOptionsPropertyOutputReference
	// Experimental.
	GuardianOptionsInput() *TfConfigurationSet_GuardianOptionsProperty
	// Experimental.
	InternalValue() *TfConfigurationSet_VdmOptionsProperty
	// Experimental.
	SetInternalValue(val *TfConfigurationSet_VdmOptionsProperty)
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
	PutDashboardOptions(value *TfConfigurationSet_DashboardOptionsProperty)
	// Experimental.
	PutGuardianOptions(value *TfConfigurationSet_GuardianOptionsProperty)
	// Experimental.
	ResetDashboardOptions()
	// Experimental.
	ResetGuardianOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationSet_VdmOptionsPropertyOutputReference
type jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) DashboardOptions() TfConfigurationSet_DashboardOptionsPropertyOutputReference {
	var returns TfConfigurationSet_DashboardOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dashboardOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) DashboardOptionsInput() *TfConfigurationSet_DashboardOptionsProperty {
	var returns *TfConfigurationSet_DashboardOptionsProperty
	_jsii_.Get(
		j,
		"dashboardOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GuardianOptions() TfConfigurationSet_GuardianOptionsPropertyOutputReference {
	var returns TfConfigurationSet_GuardianOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"guardianOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GuardianOptionsInput() *TfConfigurationSet_GuardianOptionsProperty {
	var returns *TfConfigurationSet_GuardianOptionsProperty
	_jsii_.Get(
		j,
		"guardianOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) InternalValue() *TfConfigurationSet_VdmOptionsProperty {
	var returns *TfConfigurationSet_VdmOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationSet_VdmOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigurationSet_VdmOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationSet_VdmOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfConfigurationSet.VdmOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationSet_VdmOptionsPropertyOutputReference_Override(t TfConfigurationSet_VdmOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfConfigurationSet.VdmOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference)SetInternalValue(val *TfConfigurationSet_VdmOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) PutDashboardOptions(value *TfConfigurationSet_DashboardOptionsProperty) {
	if err := t.validatePutDashboardOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDashboardOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) PutGuardianOptions(value *TfConfigurationSet_GuardianOptionsProperty) {
	if err := t.validatePutGuardianOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGuardianOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ResetDashboardOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetDashboardOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ResetGuardianOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetGuardianOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationSet_VdmOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

