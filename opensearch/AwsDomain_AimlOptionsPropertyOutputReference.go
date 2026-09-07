package opensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/opensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/opensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_AimlOptionsPropertyOutputReference interface {
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
	InternalValue() *AwsDomain_AimlOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_AimlOptionsProperty)
	// Experimental.
	NaturalLanguageQueryGenerationOptions() AwsDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	// Experimental.
	NaturalLanguageQueryGenerationOptionsInput() *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty
	// Experimental.
	S3VectorsEngine() AwsDomain_S3VectorsEnginePropertyOutputReference
	// Experimental.
	S3VectorsEngineInput() *AwsDomain_S3VectorsEngineProperty
	// Experimental.
	ServerlessVectorAcceleration() AwsDomain_ServerlessVectorAccelerationPropertyOutputReference
	// Experimental.
	ServerlessVectorAccelerationInput() *AwsDomain_ServerlessVectorAccelerationProperty
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
	PutNaturalLanguageQueryGenerationOptions(value *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty)
	// Experimental.
	PutS3VectorsEngine(value *AwsDomain_S3VectorsEngineProperty)
	// Experimental.
	PutServerlessVectorAcceleration(value *AwsDomain_ServerlessVectorAccelerationProperty)
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

// The jsii proxy struct for AwsDomain_AimlOptionsPropertyOutputReference
type jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) InternalValue() *AwsDomain_AimlOptionsProperty {
	var returns *AwsDomain_AimlOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptions() AwsDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference {
	var returns AwsDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptionsInput() *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty {
	var returns *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) S3VectorsEngine() AwsDomain_S3VectorsEnginePropertyOutputReference {
	var returns AwsDomain_S3VectorsEnginePropertyOutputReference
	_jsii_.Get(
		j,
		"s3VectorsEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) S3VectorsEngineInput() *AwsDomain_S3VectorsEngineProperty {
	var returns *AwsDomain_S3VectorsEngineProperty
	_jsii_.Get(
		j,
		"s3VectorsEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAcceleration() AwsDomain_ServerlessVectorAccelerationPropertyOutputReference {
	var returns AwsDomain_ServerlessVectorAccelerationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessVectorAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAccelerationInput() *AwsDomain_ServerlessVectorAccelerationProperty {
	var returns *AwsDomain_ServerlessVectorAccelerationProperty
	_jsii_.Get(
		j,
		"serverlessVectorAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_AimlOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_AimlOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_AimlOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_AimlOptionsPropertyOutputReference_Override(a AwsDomain_AimlOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference)SetInternalValue(val *AwsDomain_AimlOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) PutNaturalLanguageQueryGenerationOptions(value *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty) {
	if err := a.validatePutNaturalLanguageQueryGenerationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNaturalLanguageQueryGenerationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) PutS3VectorsEngine(value *AwsDomain_S3VectorsEngineProperty) {
	if err := a.validatePutS3VectorsEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3VectorsEngine",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) PutServerlessVectorAcceleration(value *AwsDomain_ServerlessVectorAccelerationProperty) {
	if err := a.validatePutServerlessVectorAccelerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerlessVectorAcceleration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ResetNaturalLanguageQueryGenerationOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetNaturalLanguageQueryGenerationOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ResetS3VectorsEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetS3VectorsEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ResetServerlessVectorAcceleration() {
	_jsii_.InvokeVoid(
		a,
		"resetServerlessVectorAcceleration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_AimlOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

