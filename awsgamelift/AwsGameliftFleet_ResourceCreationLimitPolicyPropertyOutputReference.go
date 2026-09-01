package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsgamelift/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsgamelift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference interface {
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
	InternalValue() *AwsGameliftFleet_ResourceCreationLimitPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsGameliftFleet_ResourceCreationLimitPolicyProperty)
	// Experimental.
	NewGameSessionsPerCreator() *float64
	// Experimental.
	SetNewGameSessionsPerCreator(val *float64)
	// Experimental.
	NewGameSessionsPerCreatorInput() *float64
	// Experimental.
	PolicyPeriodInMinutes() *float64
	// Experimental.
	SetPolicyPeriodInMinutes(val *float64)
	// Experimental.
	PolicyPeriodInMinutesInput() *float64
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
	ResetNewGameSessionsPerCreator()
	// Experimental.
	ResetPolicyPeriodInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference
type jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) InternalValue() *AwsGameliftFleet_ResourceCreationLimitPolicyProperty {
	var returns *AwsGameliftFleet_ResourceCreationLimitPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) NewGameSessionsPerCreator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newGameSessionsPerCreator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) NewGameSessionsPerCreatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"newGameSessionsPerCreatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) PolicyPeriodInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyPeriodInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) PolicyPeriodInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyPeriodInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-gamelift.AwsGameliftFleet.ResourceCreationLimitPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference_Override(a AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-gamelift.AwsGameliftFleet.ResourceCreationLimitPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetInternalValue(val *AwsGameliftFleet_ResourceCreationLimitPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetNewGameSessionsPerCreator(val *float64) {
	if err := j.validateSetNewGameSessionsPerCreatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newGameSessionsPerCreator",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetPolicyPeriodInMinutes(val *float64) {
	if err := j.validateSetPolicyPeriodInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyPeriodInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ResetNewGameSessionsPerCreator() {
	_jsii_.InvokeVoid(
		a,
		"resetNewGameSessionsPerCreator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ResetPolicyPeriodInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyPeriodInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGameliftFleet_ResourceCreationLimitPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

