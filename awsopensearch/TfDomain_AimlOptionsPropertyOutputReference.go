package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_AimlOptionsPropertyOutputReference interface {
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
	InternalValue() *TfDomain_AimlOptionsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_AimlOptionsProperty)
	// Experimental.
	NaturalLanguageQueryGenerationOptions() TfDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	// Experimental.
	NaturalLanguageQueryGenerationOptionsInput() *TfDomain_NaturalLanguageQueryGenerationOptionsProperty
	// Experimental.
	S3VectorsEngine() TfDomain_S3VectorsEnginePropertyOutputReference
	// Experimental.
	S3VectorsEngineInput() *TfDomain_S3VectorsEngineProperty
	// Experimental.
	ServerlessVectorAcceleration() TfDomain_ServerlessVectorAccelerationPropertyOutputReference
	// Experimental.
	ServerlessVectorAccelerationInput() *TfDomain_ServerlessVectorAccelerationProperty
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
	PutNaturalLanguageQueryGenerationOptions(value *TfDomain_NaturalLanguageQueryGenerationOptionsProperty)
	// Experimental.
	PutS3VectorsEngine(value *TfDomain_S3VectorsEngineProperty)
	// Experimental.
	PutServerlessVectorAcceleration(value *TfDomain_ServerlessVectorAccelerationProperty)
	// Experimental.
	ResetNaturalLanguageQueryGenerationOptions()
	// Experimental.
	ResetS3VectorsEngine()
	// Experimental.
	ResetServerlessVectorAcceleration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_AimlOptionsPropertyOutputReference
type jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) InternalValue() *TfDomain_AimlOptionsProperty {
	var returns *TfDomain_AimlOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptions() TfDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference {
	var returns TfDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptionsInput() *TfDomain_NaturalLanguageQueryGenerationOptionsProperty {
	var returns *TfDomain_NaturalLanguageQueryGenerationOptionsProperty
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) S3VectorsEngine() TfDomain_S3VectorsEnginePropertyOutputReference {
	var returns TfDomain_S3VectorsEnginePropertyOutputReference
	_jsii_.Get(
		j,
		"s3VectorsEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) S3VectorsEngineInput() *TfDomain_S3VectorsEngineProperty {
	var returns *TfDomain_S3VectorsEngineProperty
	_jsii_.Get(
		j,
		"s3VectorsEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAcceleration() TfDomain_ServerlessVectorAccelerationPropertyOutputReference {
	var returns TfDomain_ServerlessVectorAccelerationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessVectorAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAccelerationInput() *TfDomain_ServerlessVectorAccelerationProperty {
	var returns *TfDomain_ServerlessVectorAccelerationProperty
	_jsii_.Get(
		j,
		"serverlessVectorAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_AimlOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_AimlOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_AimlOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_AimlOptionsPropertyOutputReference_Override(t TfDomain_AimlOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.TfDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference)SetInternalValue(val *TfDomain_AimlOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) PutNaturalLanguageQueryGenerationOptions(value *TfDomain_NaturalLanguageQueryGenerationOptionsProperty) {
	if err := t.validatePutNaturalLanguageQueryGenerationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNaturalLanguageQueryGenerationOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) PutS3VectorsEngine(value *TfDomain_S3VectorsEngineProperty) {
	if err := t.validatePutS3VectorsEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3VectorsEngine",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) PutServerlessVectorAcceleration(value *TfDomain_ServerlessVectorAccelerationProperty) {
	if err := t.validatePutServerlessVectorAccelerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerlessVectorAcceleration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ResetNaturalLanguageQueryGenerationOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetNaturalLanguageQueryGenerationOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ResetS3VectorsEngine() {
	_jsii_.InvokeVoid(
		t,
		"resetS3VectorsEngine",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ResetServerlessVectorAcceleration() {
	_jsii_.InvokeVoid(
		t,
		"resetServerlessVectorAcceleration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_AimlOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

