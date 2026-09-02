package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference interface {
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
	RetrieveAndGenerateSourceConfig() TfEvaluationJob_RetrieveAndGenerateSourceConfigPropertyList
	// Experimental.
	RetrieveAndGenerateSourceConfigInput() interface{}
	// Experimental.
	RetrieveSourceConfig() TfEvaluationJob_RetrieveSourceConfigPropertyList
	// Experimental.
	RetrieveSourceConfigInput() interface{}
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
	PutRetrieveAndGenerateSourceConfig(value interface{})
	// Experimental.
	PutRetrieveSourceConfig(value interface{})
	// Experimental.
	ResetRetrieveAndGenerateSourceConfig()
	// Experimental.
	ResetRetrieveSourceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference
type jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) RetrieveAndGenerateSourceConfig() TfEvaluationJob_RetrieveAndGenerateSourceConfigPropertyList {
	var returns TfEvaluationJob_RetrieveAndGenerateSourceConfigPropertyList
	_jsii_.Get(
		j,
		"retrieveAndGenerateSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) RetrieveAndGenerateSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveAndGenerateSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) RetrieveSourceConfig() TfEvaluationJob_RetrieveSourceConfigPropertyList {
	var returns TfEvaluationJob_RetrieveSourceConfigPropertyList
	_jsii_.Get(
		j,
		"retrieveSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) RetrieveSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retrieveSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfEvaluationJob.PrecomputedRagSourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference_Override(t TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfEvaluationJob.PrecomputedRagSourceConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) PutRetrieveAndGenerateSourceConfig(value interface{}) {
	if err := t.validatePutRetrieveAndGenerateSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetrieveAndGenerateSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) PutRetrieveSourceConfig(value interface{}) {
	if err := t.validatePutRetrieveSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetrieveSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ResetRetrieveAndGenerateSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRetrieveAndGenerateSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ResetRetrieveSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRetrieveSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEvaluationJob_PrecomputedRagSourceConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

