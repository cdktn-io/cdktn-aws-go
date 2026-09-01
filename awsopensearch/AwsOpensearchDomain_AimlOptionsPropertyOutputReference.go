package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchDomain_AimlOptionsPropertyOutputReference interface {
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
	InternalValue() *AwsOpensearchDomain_AimlOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsOpensearchDomain_AimlOptionsProperty)
	// Experimental.
	NaturalLanguageQueryGenerationOptions() AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	// Experimental.
	NaturalLanguageQueryGenerationOptionsInput() *AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsProperty
	// Experimental.
	S3VectorsEngine() AwsOpensearchDomain_S3VectorsEnginePropertyOutputReference
	// Experimental.
	S3VectorsEngineInput() *AwsOpensearchDomain_S3VectorsEngineProperty
	// Experimental.
	ServerlessVectorAcceleration() AwsOpensearchDomain_ServerlessVectorAccelerationPropertyOutputReference
	// Experimental.
	ServerlessVectorAccelerationInput() *AwsOpensearchDomain_ServerlessVectorAccelerationProperty
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
	PutNaturalLanguageQueryGenerationOptions(value *AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsProperty)
	// Experimental.
	PutS3VectorsEngine(value *AwsOpensearchDomain_S3VectorsEngineProperty)
	// Experimental.
	PutServerlessVectorAcceleration(value *AwsOpensearchDomain_ServerlessVectorAccelerationProperty)
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

// The jsii proxy struct for AwsOpensearchDomain_AimlOptionsPropertyOutputReference
type jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) InternalValue() *AwsOpensearchDomain_AimlOptionsProperty {
	var returns *AwsOpensearchDomain_AimlOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptions() AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) NaturalLanguageQueryGenerationOptionsInput() *AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsProperty {
	var returns *AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsProperty
	_jsii_.Get(
		j,
		"naturalLanguageQueryGenerationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) S3VectorsEngine() AwsOpensearchDomain_S3VectorsEnginePropertyOutputReference {
	var returns AwsOpensearchDomain_S3VectorsEnginePropertyOutputReference
	_jsii_.Get(
		j,
		"s3VectorsEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) S3VectorsEngineInput() *AwsOpensearchDomain_S3VectorsEngineProperty {
	var returns *AwsOpensearchDomain_S3VectorsEngineProperty
	_jsii_.Get(
		j,
		"s3VectorsEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAcceleration() AwsOpensearchDomain_ServerlessVectorAccelerationPropertyOutputReference {
	var returns AwsOpensearchDomain_ServerlessVectorAccelerationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessVectorAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ServerlessVectorAccelerationInput() *AwsOpensearchDomain_ServerlessVectorAccelerationProperty {
	var returns *AwsOpensearchDomain_ServerlessVectorAccelerationProperty
	_jsii_.Get(
		j,
		"serverlessVectorAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOpensearchDomain_AimlOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOpensearchDomain_AimlOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOpensearchDomain_AimlOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOpensearchDomain_AimlOptionsPropertyOutputReference_Override(a AwsOpensearchDomain_AimlOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.AimlOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference)SetInternalValue(val *AwsOpensearchDomain_AimlOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) PutNaturalLanguageQueryGenerationOptions(value *AwsOpensearchDomain_NaturalLanguageQueryGenerationOptionsProperty) {
	if err := a.validatePutNaturalLanguageQueryGenerationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNaturalLanguageQueryGenerationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) PutS3VectorsEngine(value *AwsOpensearchDomain_S3VectorsEngineProperty) {
	if err := a.validatePutS3VectorsEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3VectorsEngine",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) PutServerlessVectorAcceleration(value *AwsOpensearchDomain_ServerlessVectorAccelerationProperty) {
	if err := a.validatePutServerlessVectorAccelerationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerlessVectorAcceleration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ResetNaturalLanguageQueryGenerationOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetNaturalLanguageQueryGenerationOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ResetS3VectorsEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetS3VectorsEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ResetServerlessVectorAcceleration() {
	_jsii_.InvokeVoid(
		a,
		"resetServerlessVectorAcceleration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_AimlOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

