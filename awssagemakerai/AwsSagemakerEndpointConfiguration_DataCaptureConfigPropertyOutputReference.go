package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CaptureContentTypeHeader() AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderPropertyOutputReference
	// Experimental.
	CaptureContentTypeHeaderInput() *AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderProperty
	// Experimental.
	CaptureOptions() AwsSagemakerEndpointConfiguration_CaptureOptionsPropertyList
	// Experimental.
	CaptureOptionsInput() interface{}
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
	DestinationS3Uri() *string
	// Experimental.
	SetDestinationS3Uri(val *string)
	// Experimental.
	DestinationS3UriInput() *string
	// Experimental.
	EnableCapture() interface{}
	// Experimental.
	SetEnableCapture(val interface{})
	// Experimental.
	EnableCaptureInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InitialSamplingPercentage() *float64
	// Experimental.
	SetInitialSamplingPercentage(val *float64)
	// Experimental.
	InitialSamplingPercentageInput() *float64
	// Experimental.
	InternalValue() *AwsSagemakerEndpointConfiguration_DataCaptureConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerEndpointConfiguration_DataCaptureConfigProperty)
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
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
	PutCaptureContentTypeHeader(value *AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderProperty)
	// Experimental.
	PutCaptureOptions(value interface{})
	// Experimental.
	ResetCaptureContentTypeHeader()
	// Experimental.
	ResetEnableCapture()
	// Experimental.
	ResetKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) CaptureContentTypeHeader() AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderPropertyOutputReference {
	var returns AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderPropertyOutputReference
	_jsii_.Get(
		j,
		"captureContentTypeHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) CaptureContentTypeHeaderInput() *AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderProperty {
	var returns *AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderProperty
	_jsii_.Get(
		j,
		"captureContentTypeHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) CaptureOptions() AwsSagemakerEndpointConfiguration_CaptureOptionsPropertyList {
	var returns AwsSagemakerEndpointConfiguration_CaptureOptionsPropertyList
	_jsii_.Get(
		j,
		"captureOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) CaptureOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captureOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) DestinationS3Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationS3Uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) DestinationS3UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationS3UriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) EnableCapture() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCapture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) EnableCaptureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCaptureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) InitialSamplingPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSamplingPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) InitialSamplingPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialSamplingPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) InternalValue() *AwsSagemakerEndpointConfiguration_DataCaptureConfigProperty {
	var returns *AwsSagemakerEndpointConfiguration_DataCaptureConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.DataCaptureConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference_Override(a AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpointConfiguration.DataCaptureConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetDestinationS3Uri(val *string) {
	if err := j.validateSetDestinationS3UriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationS3Uri",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetEnableCapture(val interface{}) {
	if err := j.validateSetEnableCaptureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCapture",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetInitialSamplingPercentage(val *float64) {
	if err := j.validateSetInitialSamplingPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialSamplingPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerEndpointConfiguration_DataCaptureConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) PutCaptureContentTypeHeader(value *AwsSagemakerEndpointConfiguration_CaptureContentTypeHeaderProperty) {
	if err := a.validatePutCaptureContentTypeHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptureContentTypeHeader",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) PutCaptureOptions(value interface{}) {
	if err := a.validatePutCaptureOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptureOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ResetCaptureContentTypeHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptureContentTypeHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ResetEnableCapture() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableCapture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpointConfiguration_DataCaptureConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

