package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SimpleScopeTerm() TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermPropertyOutputReference
	// Experimental.
	SimpleScopeTermInput() *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty
	// Experimental.
	TagScopeTerm() TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermPropertyOutputReference
	// Experimental.
	TagScopeTermInput() *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty
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
	PutSimpleScopeTerm(value *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty)
	// Experimental.
	PutTagScopeTerm(value *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty)
	// Experimental.
	ResetSimpleScopeTerm()
	// Experimental.
	ResetTagScopeTerm()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference
type jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) SimpleScopeTerm() TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermPropertyOutputReference {
	var returns TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermPropertyOutputReference
	_jsii_.Get(
		j,
		"simpleScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) SimpleScopeTermInput() *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty {
	var returns *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty
	_jsii_.Get(
		j,
		"simpleScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) TagScopeTerm() TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermPropertyOutputReference {
	var returns TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermPropertyOutputReference
	_jsii_.Get(
		j,
		"tagScopeTerm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) TagScopeTermInput() *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty {
	var returns *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty
	_jsii_.Get(
		j,
		"tagScopeTermInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionScopingIncludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference_Override(t TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionScopingIncludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) PutSimpleScopeTerm(value *TfClassificationJob_S3JobDefinitionScopingIncludesAndSimpleScopeTermProperty) {
	if err := t.validatePutSimpleScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSimpleScopeTerm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) PutTagScopeTerm(value *TfClassificationJob_S3JobDefinitionScopingIncludesAndTagScopeTermProperty) {
	if err := t.validatePutTagScopeTermParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagScopeTerm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ResetSimpleScopeTerm() {
	_jsii_.InvokeVoid(
		t,
		"resetSimpleScopeTerm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ResetTagScopeTerm() {
	_jsii_.InvokeVoid(
		t,
		"resetTagScopeTerm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionScopingIncludesAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

