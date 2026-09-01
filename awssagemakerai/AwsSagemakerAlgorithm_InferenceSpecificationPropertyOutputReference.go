package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference interface {
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
	// Experimental.
	Containers() AwsSagemakerAlgorithm_ContainersPropertyList
	// Experimental.
	ContainersInput() interface{}
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
	SupportedContentTypes() *[]*string
	// Experimental.
	SetSupportedContentTypes(val *[]*string)
	// Experimental.
	SupportedContentTypesInput() *[]*string
	// Experimental.
	SupportedRealtimeInferenceInstanceTypes() *[]*string
	// Experimental.
	SetSupportedRealtimeInferenceInstanceTypes(val *[]*string)
	// Experimental.
	SupportedRealtimeInferenceInstanceTypesInput() *[]*string
	// Experimental.
	SupportedResponseMimeTypes() *[]*string
	// Experimental.
	SetSupportedResponseMimeTypes(val *[]*string)
	// Experimental.
	SupportedResponseMimeTypesInput() *[]*string
	// Experimental.
	SupportedTransformInstanceTypes() *[]*string
	// Experimental.
	SetSupportedTransformInstanceTypes(val *[]*string)
	// Experimental.
	SupportedTransformInstanceTypesInput() *[]*string
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
	PutContainers(value interface{})
	// Experimental.
	ResetContainers()
	// Experimental.
	ResetSupportedContentTypes()
	// Experimental.
	ResetSupportedRealtimeInferenceInstanceTypes()
	// Experimental.
	ResetSupportedResponseMimeTypes()
	// Experimental.
	ResetSupportedTransformInstanceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference
type jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) Containers() AwsSagemakerAlgorithm_ContainersPropertyList {
	var returns AwsSagemakerAlgorithm_ContainersPropertyList
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ContainersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedContentTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedContentTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedContentTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedRealtimeInferenceInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedRealtimeInferenceInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedRealtimeInferenceInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedResponseMimeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedResponseMimeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedResponseMimeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedTransformInstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) SupportedTransformInstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTransformInstanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.InferenceSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference_Override(a AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerAlgorithm.InferenceSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetSupportedContentTypes(val *[]*string) {
	if err := j.validateSetSupportedContentTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedContentTypes",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetSupportedRealtimeInferenceInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedRealtimeInferenceInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedRealtimeInferenceInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetSupportedResponseMimeTypes(val *[]*string) {
	if err := j.validateSetSupportedResponseMimeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedResponseMimeTypes",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetSupportedTransformInstanceTypes(val *[]*string) {
	if err := j.validateSetSupportedTransformInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedTransformInstanceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) PutContainers(value interface{}) {
	if err := a.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ResetContainers() {
	_jsii_.InvokeVoid(
		a,
		"resetContainers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ResetSupportedContentTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportedContentTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ResetSupportedRealtimeInferenceInstanceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportedRealtimeInferenceInstanceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ResetSupportedResponseMimeTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportedResponseMimeTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ResetSupportedTransformInstanceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetSupportedTransformInstanceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerAlgorithm_InferenceSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

