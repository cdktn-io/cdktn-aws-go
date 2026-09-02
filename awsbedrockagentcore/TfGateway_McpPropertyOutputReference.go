package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGateway_McpPropertyOutputReference interface {
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
	Instructions() *string
	// Experimental.
	SetInstructions(val *string)
	// Experimental.
	InstructionsInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SearchType() *string
	// Experimental.
	SetSearchType(val *string)
	// Experimental.
	SearchTypeInput() *string
	// Experimental.
	SessionConfiguration() TfGateway_SessionConfigurationPropertyList
	// Experimental.
	SessionConfigurationInput() interface{}
	// Experimental.
	StreamingConfiguration() TfGateway_StreamingConfigurationPropertyList
	// Experimental.
	StreamingConfigurationInput() interface{}
	// Experimental.
	SupportedVersions() *[]*string
	// Experimental.
	SetSupportedVersions(val *[]*string)
	// Experimental.
	SupportedVersionsInput() *[]*string
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
	PutSessionConfiguration(value interface{})
	// Experimental.
	PutStreamingConfiguration(value interface{})
	// Experimental.
	ResetInstructions()
	// Experimental.
	ResetSearchType()
	// Experimental.
	ResetSessionConfiguration()
	// Experimental.
	ResetStreamingConfiguration()
	// Experimental.
	ResetSupportedVersions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGateway_McpPropertyOutputReference
type jsiiProxy_TfGateway_McpPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) Instructions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) InstructionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SearchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SearchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SessionConfiguration() TfGateway_SessionConfigurationPropertyList {
	var returns TfGateway_SessionConfigurationPropertyList
	_jsii_.Get(
		j,
		"sessionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SessionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) StreamingConfiguration() TfGateway_StreamingConfigurationPropertyList {
	var returns TfGateway_StreamingConfigurationPropertyList
	_jsii_.Get(
		j,
		"streamingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) StreamingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SupportedVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) SupportedVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGateway_McpPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGateway_McpPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGateway_McpPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGateway_McpPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGateway_McpPropertyOutputReference_Override(t TfGateway_McpPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGateway.McpPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetInstructions(val *string) {
	if err := j.validateSetInstructionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instructions",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetSearchType(val *string) {
	if err := j.validateSetSearchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchType",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetSupportedVersions(val *[]*string) {
	if err := j.validateSetSupportedVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportedVersions",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGateway_McpPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) PutSessionConfiguration(value interface{}) {
	if err := t.validatePutSessionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSessionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) PutStreamingConfiguration(value interface{}) {
	if err := t.validatePutStreamingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStreamingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ResetInstructions() {
	_jsii_.InvokeVoid(
		t,
		"resetInstructions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ResetSearchType() {
	_jsii_.InvokeVoid(
		t,
		"resetSearchType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ResetSessionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ResetStreamingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ResetSupportedVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetSupportedVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGateway_McpPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

