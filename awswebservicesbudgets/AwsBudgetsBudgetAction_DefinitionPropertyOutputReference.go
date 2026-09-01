package awswebservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudgetsBudgetAction_DefinitionPropertyOutputReference interface {
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
	IamActionDefinition() AwsBudgetsBudgetAction_IamActionDefinitionPropertyOutputReference
	// Experimental.
	IamActionDefinitionInput() *AwsBudgetsBudgetAction_IamActionDefinitionProperty
	// Experimental.
	InternalValue() *AwsBudgetsBudgetAction_DefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsBudgetsBudgetAction_DefinitionProperty)
	// Experimental.
	ScpActionDefinition() AwsBudgetsBudgetAction_ScpActionDefinitionPropertyOutputReference
	// Experimental.
	ScpActionDefinitionInput() *AwsBudgetsBudgetAction_ScpActionDefinitionProperty
	// Experimental.
	SsmActionDefinition() AwsBudgetsBudgetAction_SsmActionDefinitionPropertyOutputReference
	// Experimental.
	SsmActionDefinitionInput() *AwsBudgetsBudgetAction_SsmActionDefinitionProperty
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
	PutIamActionDefinition(value *AwsBudgetsBudgetAction_IamActionDefinitionProperty)
	// Experimental.
	PutScpActionDefinition(value *AwsBudgetsBudgetAction_ScpActionDefinitionProperty)
	// Experimental.
	PutSsmActionDefinition(value *AwsBudgetsBudgetAction_SsmActionDefinitionProperty)
	// Experimental.
	ResetIamActionDefinition()
	// Experimental.
	ResetScpActionDefinition()
	// Experimental.
	ResetSsmActionDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBudgetsBudgetAction_DefinitionPropertyOutputReference
type jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) IamActionDefinition() AwsBudgetsBudgetAction_IamActionDefinitionPropertyOutputReference {
	var returns AwsBudgetsBudgetAction_IamActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"iamActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) IamActionDefinitionInput() *AwsBudgetsBudgetAction_IamActionDefinitionProperty {
	var returns *AwsBudgetsBudgetAction_IamActionDefinitionProperty
	_jsii_.Get(
		j,
		"iamActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) InternalValue() *AwsBudgetsBudgetAction_DefinitionProperty {
	var returns *AwsBudgetsBudgetAction_DefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ScpActionDefinition() AwsBudgetsBudgetAction_ScpActionDefinitionPropertyOutputReference {
	var returns AwsBudgetsBudgetAction_ScpActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"scpActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ScpActionDefinitionInput() *AwsBudgetsBudgetAction_ScpActionDefinitionProperty {
	var returns *AwsBudgetsBudgetAction_ScpActionDefinitionProperty
	_jsii_.Get(
		j,
		"scpActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) SsmActionDefinition() AwsBudgetsBudgetAction_SsmActionDefinitionPropertyOutputReference {
	var returns AwsBudgetsBudgetAction_SsmActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"ssmActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) SsmActionDefinitionInput() *AwsBudgetsBudgetAction_SsmActionDefinitionProperty {
	var returns *AwsBudgetsBudgetAction_SsmActionDefinitionProperty
	_jsii_.Get(
		j,
		"ssmActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBudgetsBudgetAction_DefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBudgetsBudgetAction_DefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBudgetsBudgetAction_DefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudgetAction.DefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBudgetsBudgetAction_DefinitionPropertyOutputReference_Override(a AwsBudgetsBudgetAction_DefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudgetsBudgetAction.DefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference)SetInternalValue(val *AwsBudgetsBudgetAction_DefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) PutIamActionDefinition(value *AwsBudgetsBudgetAction_IamActionDefinitionProperty) {
	if err := a.validatePutIamActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIamActionDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) PutScpActionDefinition(value *AwsBudgetsBudgetAction_ScpActionDefinitionProperty) {
	if err := a.validatePutScpActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScpActionDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) PutSsmActionDefinition(value *AwsBudgetsBudgetAction_SsmActionDefinitionProperty) {
	if err := a.validatePutSsmActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSsmActionDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ResetIamActionDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetIamActionDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ResetScpActionDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetScpActionDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ResetSsmActionDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetSsmActionDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBudgetsBudgetAction_DefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

