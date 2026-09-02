package awswebservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBudgetAction_DefinitionPropertyOutputReference interface {
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
	IamActionDefinition() TfBudgetAction_IamActionDefinitionPropertyOutputReference
	// Experimental.
	IamActionDefinitionInput() *TfBudgetAction_IamActionDefinitionProperty
	// Experimental.
	InternalValue() *TfBudgetAction_DefinitionProperty
	// Experimental.
	SetInternalValue(val *TfBudgetAction_DefinitionProperty)
	// Experimental.
	ScpActionDefinition() TfBudgetAction_ScpActionDefinitionPropertyOutputReference
	// Experimental.
	ScpActionDefinitionInput() *TfBudgetAction_ScpActionDefinitionProperty
	// Experimental.
	SsmActionDefinition() TfBudgetAction_SsmActionDefinitionPropertyOutputReference
	// Experimental.
	SsmActionDefinitionInput() *TfBudgetAction_SsmActionDefinitionProperty
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
	PutIamActionDefinition(value *TfBudgetAction_IamActionDefinitionProperty)
	// Experimental.
	PutScpActionDefinition(value *TfBudgetAction_ScpActionDefinitionProperty)
	// Experimental.
	PutSsmActionDefinition(value *TfBudgetAction_SsmActionDefinitionProperty)
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

// The jsii proxy struct for TfBudgetAction_DefinitionPropertyOutputReference
type jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) IamActionDefinition() TfBudgetAction_IamActionDefinitionPropertyOutputReference {
	var returns TfBudgetAction_IamActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"iamActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) IamActionDefinitionInput() *TfBudgetAction_IamActionDefinitionProperty {
	var returns *TfBudgetAction_IamActionDefinitionProperty
	_jsii_.Get(
		j,
		"iamActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) InternalValue() *TfBudgetAction_DefinitionProperty {
	var returns *TfBudgetAction_DefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ScpActionDefinition() TfBudgetAction_ScpActionDefinitionPropertyOutputReference {
	var returns TfBudgetAction_ScpActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"scpActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ScpActionDefinitionInput() *TfBudgetAction_ScpActionDefinitionProperty {
	var returns *TfBudgetAction_ScpActionDefinitionProperty
	_jsii_.Get(
		j,
		"scpActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) SsmActionDefinition() TfBudgetAction_SsmActionDefinitionPropertyOutputReference {
	var returns TfBudgetAction_SsmActionDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"ssmActionDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) SsmActionDefinitionInput() *TfBudgetAction_SsmActionDefinitionProperty {
	var returns *TfBudgetAction_SsmActionDefinitionProperty
	_jsii_.Get(
		j,
		"ssmActionDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBudgetAction_DefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBudgetAction_DefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBudgetAction_DefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudgetAction.DefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBudgetAction_DefinitionPropertyOutputReference_Override(t TfBudgetAction_DefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudgetAction.DefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference)SetInternalValue(val *TfBudgetAction_DefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) PutIamActionDefinition(value *TfBudgetAction_IamActionDefinitionProperty) {
	if err := t.validatePutIamActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIamActionDefinition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) PutScpActionDefinition(value *TfBudgetAction_ScpActionDefinitionProperty) {
	if err := t.validatePutScpActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScpActionDefinition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) PutSsmActionDefinition(value *TfBudgetAction_SsmActionDefinitionProperty) {
	if err := t.validatePutSsmActionDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSsmActionDefinition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ResetIamActionDefinition() {
	_jsii_.InvokeVoid(
		t,
		"resetIamActionDefinition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ResetScpActionDefinition() {
	_jsii_.InvokeVoid(
		t,
		"resetScpActionDefinition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ResetSsmActionDefinition() {
	_jsii_.InvokeVoid(
		t,
		"resetSsmActionDefinition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBudgetAction_DefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

