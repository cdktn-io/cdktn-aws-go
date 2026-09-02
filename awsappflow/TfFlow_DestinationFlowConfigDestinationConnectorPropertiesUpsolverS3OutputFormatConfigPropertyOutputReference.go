package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AggregationConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference
	// Experimental.
	AggregationConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty
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
	FileType() *string
	// Experimental.
	SetFileType(val *string)
	// Experimental.
	FileTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty
	// Experimental.
	SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty)
	// Experimental.
	PrefixConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference
	// Experimental.
	PrefixConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty
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
	PutAggregationConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty)
	// Experimental.
	PutPrefixConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty)
	// Experimental.
	ResetAggregationConfig()
	// Experimental.
	ResetFileType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference
type jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) AggregationConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"aggregationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) AggregationConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty
	_jsii_.Get(
		j,
		"aggregationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) FileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) FileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InternalValue() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PrefixConfig() TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference {
	var returns TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PrefixConfigInput() *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty {
	var returns *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty
	_jsii_.Get(
		j,
		"prefixConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference_Override(t TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfFlow.DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetFileType(val *string) {
	if err := j.validateSetFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileType",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetInternalValue(val *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PutAggregationConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty) {
	if err := t.validatePutAggregationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAggregationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PutPrefixConfig(value *TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty) {
	if err := t.validatePutPrefixConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrefixConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ResetAggregationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAggregationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ResetFileType() {
	_jsii_.InvokeVoid(
		t,
		"resetFileType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

