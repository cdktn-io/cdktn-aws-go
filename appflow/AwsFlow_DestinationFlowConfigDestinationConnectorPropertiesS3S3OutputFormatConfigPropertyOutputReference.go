package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AggregationConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference
	// Experimental.
	AggregationConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty
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
	InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty)
	// Experimental.
	PrefixConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigPropertyOutputReference
	// Experimental.
	PrefixConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty
	// Experimental.
	PreserveSourceDataTyping() interface{}
	// Experimental.
	SetPreserveSourceDataTyping(val interface{})
	// Experimental.
	PreserveSourceDataTypingInput() interface{}
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
	PutAggregationConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty)
	// Experimental.
	PutPrefixConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty)
	// Experimental.
	ResetAggregationConfig()
	// Experimental.
	ResetFileType()
	// Experimental.
	ResetPrefixConfig()
	// Experimental.
	ResetPreserveSourceDataTyping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference
type jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) AggregationConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"aggregationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) AggregationConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty
	_jsii_.Get(
		j,
		"aggregationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) FileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) FileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PrefixConfig() AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigPropertyOutputReference {
	var returns AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PrefixConfigInput() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty
	_jsii_.Get(
		j,
		"prefixConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PreserveSourceDataTyping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSourceDataTyping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PreserveSourceDataTypingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveSourceDataTypingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference_Override(a AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetFileType(val *string) {
	if err := j.validateSetFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileType",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetPreserveSourceDataTyping(val interface{}) {
	if err := j.validateSetPreserveSourceDataTypingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveSourceDataTyping",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PutAggregationConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty) {
	if err := a.validatePutAggregationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAggregationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) PutPrefixConfig(value *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigProperty) {
	if err := a.validatePutPrefixConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefixConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ResetAggregationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAggregationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ResetFileType() {
	_jsii_.InvokeVoid(
		a,
		"resetFileType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ResetPrefixConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ResetPreserveSourceDataTyping() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveSourceDataTyping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

