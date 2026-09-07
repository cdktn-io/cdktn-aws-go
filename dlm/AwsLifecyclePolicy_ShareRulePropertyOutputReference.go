package dlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLifecyclePolicy_ShareRulePropertyOutputReference interface {
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
	InternalValue() *AwsLifecyclePolicy_ShareRuleProperty
	// Experimental.
	SetInternalValue(val *AwsLifecyclePolicy_ShareRuleProperty)
	// Experimental.
	TargetAccounts() *[]*string
	// Experimental.
	SetTargetAccounts(val *[]*string)
	// Experimental.
	TargetAccountsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnshareInterval() *float64
	// Experimental.
	SetUnshareInterval(val *float64)
	// Experimental.
	UnshareIntervalInput() *float64
	// Experimental.
	UnshareIntervalUnit() *string
	// Experimental.
	SetUnshareIntervalUnit(val *string)
	// Experimental.
	UnshareIntervalUnitInput() *string
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
	ResetUnshareInterval()
	// Experimental.
	ResetUnshareIntervalUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLifecyclePolicy_ShareRulePropertyOutputReference
type jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) InternalValue() *AwsLifecyclePolicy_ShareRuleProperty {
	var returns *AwsLifecyclePolicy_ShareRuleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) TargetAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) TargetAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) UnshareInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unshareInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) UnshareIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unshareIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) UnshareIntervalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unshareIntervalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) UnshareIntervalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unshareIntervalUnitInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLifecyclePolicy_ShareRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLifecyclePolicy_ShareRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLifecyclePolicy_ShareRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.ShareRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLifecyclePolicy_ShareRulePropertyOutputReference_Override(a AwsLifecyclePolicy_ShareRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.AwsLifecyclePolicy.ShareRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetInternalValue(val *AwsLifecyclePolicy_ShareRuleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetTargetAccounts(val *[]*string) {
	if err := j.validateSetTargetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccounts",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetUnshareInterval(val *float64) {
	if err := j.validateSetUnshareIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unshareInterval",
		val,
	)
}

func (j *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference)SetUnshareIntervalUnit(val *string) {
	if err := j.validateSetUnshareIntervalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unshareIntervalUnit",
		val,
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ResetUnshareInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetUnshareInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ResetUnshareIntervalUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetUnshareIntervalUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLifecyclePolicy_ShareRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

