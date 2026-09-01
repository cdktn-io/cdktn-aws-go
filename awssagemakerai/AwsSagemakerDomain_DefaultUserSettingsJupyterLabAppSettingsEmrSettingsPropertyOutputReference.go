package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AssumableRoleArns() *[]*string
	// Experimental.
	SetAssumableRoleArns(val *[]*string)
	// Experimental.
	AssumableRoleArnsInput() *[]*string
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
	ExecutionRoleArns() *[]*string
	// Experimental.
	SetExecutionRoleArns(val *[]*string)
	// Experimental.
	ExecutionRoleArnsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty)
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
	ResetAssumableRoleArns()
	// Experimental.
	ResetExecutionRoleArns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference
type jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) AssumableRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assumableRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) AssumableRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assumableRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ExecutionRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"executionRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ExecutionRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"executionRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) InternalValue() *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty {
	var returns *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference_Override(a AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerDomain.DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetAssumableRoleArns(val *[]*string) {
	if err := j.validateSetAssumableRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumableRoleArns",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetExecutionRoleArns(val *[]*string) {
	if err := j.validateSetExecutionRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArns",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetInternalValue(val *AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ResetAssumableRoleArns() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumableRoleArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ResetExecutionRoleArns() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRoleArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerDomain_DefaultUserSettingsJupyterLabAppSettingsEmrSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

