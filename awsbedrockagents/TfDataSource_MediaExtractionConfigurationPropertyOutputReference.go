package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_MediaExtractionConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioExtractionConfiguration() TfDataSource_AudioExtractionConfigurationPropertyList
	// Experimental.
	AudioExtractionConfigurationInput() interface{}
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
	ImageExtractionConfiguration() TfDataSource_ImageExtractionConfigurationPropertyList
	// Experimental.
	ImageExtractionConfigurationInput() interface{}
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
	VideoExtractionConfiguration() TfDataSource_VideoExtractionConfigurationPropertyList
	// Experimental.
	VideoExtractionConfigurationInput() interface{}
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
	PutAudioExtractionConfiguration(value interface{})
	// Experimental.
	PutImageExtractionConfiguration(value interface{})
	// Experimental.
	PutVideoExtractionConfiguration(value interface{})
	// Experimental.
	ResetAudioExtractionConfiguration()
	// Experimental.
	ResetImageExtractionConfiguration()
	// Experimental.
	ResetVideoExtractionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSource_MediaExtractionConfigurationPropertyOutputReference
type jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) AudioExtractionConfiguration() TfDataSource_AudioExtractionConfigurationPropertyList {
	var returns TfDataSource_AudioExtractionConfigurationPropertyList
	_jsii_.Get(
		j,
		"audioExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) AudioExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ImageExtractionConfiguration() TfDataSource_ImageExtractionConfigurationPropertyList {
	var returns TfDataSource_ImageExtractionConfigurationPropertyList
	_jsii_.Get(
		j,
		"imageExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ImageExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) VideoExtractionConfiguration() TfDataSource_VideoExtractionConfigurationPropertyList {
	var returns TfDataSource_VideoExtractionConfigurationPropertyList
	_jsii_.Get(
		j,
		"videoExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) VideoExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoExtractionConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_MediaExtractionConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSource_MediaExtractionConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_MediaExtractionConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.MediaExtractionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_MediaExtractionConfigurationPropertyOutputReference_Override(t TfDataSource_MediaExtractionConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfDataSource.MediaExtractionConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) PutAudioExtractionConfiguration(value interface{}) {
	if err := t.validatePutAudioExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioExtractionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) PutImageExtractionConfiguration(value interface{}) {
	if err := t.validatePutImageExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageExtractionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) PutVideoExtractionConfiguration(value interface{}) {
	if err := t.validatePutVideoExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVideoExtractionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ResetAudioExtractionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioExtractionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ResetImageExtractionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetImageExtractionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ResetVideoExtractionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoExtractionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSource_MediaExtractionConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

