package awssesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference interface {
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
	DashboardOptions() AwsSesv2ConfigurationSet_DashboardOptionsPropertyOutputReference
	// Experimental.
	DashboardOptionsInput() *AwsSesv2ConfigurationSet_DashboardOptionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	GuardianOptions() AwsSesv2ConfigurationSet_GuardianOptionsPropertyOutputReference
	// Experimental.
	GuardianOptionsInput() *AwsSesv2ConfigurationSet_GuardianOptionsProperty
	// Experimental.
	InternalValue() *AwsSesv2ConfigurationSet_VdmOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsSesv2ConfigurationSet_VdmOptionsProperty)
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
	PutDashboardOptions(value *AwsSesv2ConfigurationSet_DashboardOptionsProperty)
	// Experimental.
	PutGuardianOptions(value *AwsSesv2ConfigurationSet_GuardianOptionsProperty)
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

// The jsii proxy struct for AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference
type jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) DashboardOptions() AwsSesv2ConfigurationSet_DashboardOptionsPropertyOutputReference {
	var returns AwsSesv2ConfigurationSet_DashboardOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dashboardOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) DashboardOptionsInput() *AwsSesv2ConfigurationSet_DashboardOptionsProperty {
	var returns *AwsSesv2ConfigurationSet_DashboardOptionsProperty
	_jsii_.Get(
		j,
		"dashboardOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GuardianOptions() AwsSesv2ConfigurationSet_GuardianOptionsPropertyOutputReference {
	var returns AwsSesv2ConfigurationSet_GuardianOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"guardianOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GuardianOptionsInput() *AwsSesv2ConfigurationSet_GuardianOptionsProperty {
	var returns *AwsSesv2ConfigurationSet_GuardianOptionsProperty
	_jsii_.Get(
		j,
		"guardianOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) InternalValue() *AwsSesv2ConfigurationSet_VdmOptionsProperty {
	var returns *AwsSesv2ConfigurationSet_VdmOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsSesv2ConfigurationSet.VdmOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference_Override(a AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.AwsSesv2ConfigurationSet.VdmOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference)SetInternalValue(val *AwsSesv2ConfigurationSet_VdmOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) PutDashboardOptions(value *AwsSesv2ConfigurationSet_DashboardOptionsProperty) {
	if err := a.validatePutDashboardOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDashboardOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) PutGuardianOptions(value *AwsSesv2ConfigurationSet_GuardianOptionsProperty) {
	if err := a.validatePutGuardianOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGuardianOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ResetDashboardOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetDashboardOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ResetGuardianOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetGuardianOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSesv2ConfigurationSet_VdmOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

