package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerAlgorithm_RangePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CategoricalParameterRangeSpecification() AwsSagemakerAlgorithm_CategoricalParameterRangeSpecificationPropertyList
	// Experimental.
	CategoricalParameterRangeSpecificationInput() interface{}
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
	// Experimental.
	ContinuousParameterRangeSpecification() AwsSagemakerAlgorithm_ContinuousParameterRangeSpecificationPropertyList
	// Experimental.
	ContinuousParameterRangeSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IntegerParameterRangeSpecification() AwsSagemakerAlgorithm_IntegerParameterRangeSpecificationPropertyList
	// Experimental.
	IntegerParameterRangeSpecificationInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
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
	PutCategoricalParameterRangeSpecification(value interface{})
	// Experimental.
	PutContinuousParameterRangeSpecification(value interface{})
	// Experimental.
	PutIntegerParameterRangeSpecification(value interface{})
	// Experimental.
	ResetCategoricalParameterRangeSpecification()
	// Experimental.
	ResetContinuousParameterRangeSpecification()
	// Experimental.
	ResetIntegerParameterRangeSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerAlgorithm_RangePropertyOutputReference
type jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) CategoricalParameterRangeSpecification() AwsSagemakerAlgorithm_CategoricalParameterRangeSpecificationPropertyList {
	var returns AwsSagemakerAlgorithm_CategoricalParameterRangeSpecificationPropertyList
	_jsii_.Get(
		j,
		"categoricalParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) CategoricalParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoricalParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ContinuousParameterRangeSpecification() AwsSagemakerAlgorithm_ContinuousParameterRangeSpecificationPropertyList {
	var returns AwsSagemakerAlgorithm_ContinuousParameterRangeSpecificationPropertyList
	_jsii_.Get(
		j,
		"continuousParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ContinuousParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) IntegerParameterRangeSpecification() AwsSagemakerAlgorithm_IntegerParameterRangeSpecificationPropertyList {
	var returns AwsSagemakerAlgorithm_IntegerParameterRangeSpecificationPropertyList
	_jsii_.Get(
		j,
		"integerParameterRangeSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) IntegerParameterRangeSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParameterRangeSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerAlgorithm_RangePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerAlgorithm_RangePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerAlgorithm_RangePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.RangePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerAlgorithm_RangePropertyOutputReference_Override(a AwsSagemakerAlgorithm_RangePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.RangePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) PutCategoricalParameterRangeSpecification(value interface{}) {
	if err := a.validatePutCategoricalParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCategoricalParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) PutContinuousParameterRangeSpecification(value interface{}) {
	if err := a.validatePutContinuousParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContinuousParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) PutIntegerParameterRangeSpecification(value interface{}) {
	if err := a.validatePutIntegerParameterRangeSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegerParameterRangeSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ResetCategoricalParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCategoricalParameterRangeSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ResetContinuousParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuousParameterRangeSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ResetIntegerParameterRangeSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegerParameterRangeSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_RangePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

