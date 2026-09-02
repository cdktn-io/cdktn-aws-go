package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_InputPropertyOutputReference interface {
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
	InAppStreamNames() *[]*string
	// Experimental.
	InputId() *string
	// Experimental.
	InputParallelism() TfApplication_InputParallelismPropertyOutputReference
	// Experimental.
	InputParallelismInput() *TfApplication_InputParallelismProperty
	// Experimental.
	InputProcessingConfiguration() TfApplication_InputProcessingConfigurationPropertyOutputReference
	// Experimental.
	InputProcessingConfigurationInput() *TfApplication_InputProcessingConfigurationProperty
	// Experimental.
	InputSchema() TfApplication_InputSchemaPropertyOutputReference
	// Experimental.
	InputSchemaInput() *TfApplication_InputSchemaProperty
	// Experimental.
	InputStartingPositionConfiguration() TfApplication_InputStartingPositionConfigurationPropertyList
	// Experimental.
	InputStartingPositionConfigurationInput() interface{}
	// Experimental.
	InternalValue() *TfApplication_InputProperty
	// Experimental.
	SetInternalValue(val *TfApplication_InputProperty)
	// Experimental.
	KinesisFirehoseInput() TfApplication_KinesisFirehoseInputPropertyOutputReference
	// Experimental.
	KinesisFirehoseInputInput() *TfApplication_KinesisFirehoseInputProperty
	// Experimental.
	KinesisStreamsInput() TfApplication_KinesisStreamsInputPropertyOutputReference
	// Experimental.
	KinesisStreamsInputInput() *TfApplication_KinesisStreamsInputProperty
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
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
	PutInputParallelism(value *TfApplication_InputParallelismProperty)
	// Experimental.
	PutInputProcessingConfiguration(value *TfApplication_InputProcessingConfigurationProperty)
	// Experimental.
	PutInputSchema(value *TfApplication_InputSchemaProperty)
	// Experimental.
	PutInputStartingPositionConfiguration(value interface{})
	// Experimental.
	PutKinesisFirehoseInput(value *TfApplication_KinesisFirehoseInputProperty)
	// Experimental.
	PutKinesisStreamsInput(value *TfApplication_KinesisStreamsInputProperty)
	// Experimental.
	ResetInputParallelism()
	// Experimental.
	ResetInputProcessingConfiguration()
	// Experimental.
	ResetInputStartingPositionConfiguration()
	// Experimental.
	ResetKinesisFirehoseInput()
	// Experimental.
	ResetKinesisStreamsInput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_InputPropertyOutputReference
type jsiiProxy_TfApplication_InputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InAppStreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inAppStreamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputParallelism() TfApplication_InputParallelismPropertyOutputReference {
	var returns TfApplication_InputParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"inputParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputParallelismInput() *TfApplication_InputParallelismProperty {
	var returns *TfApplication_InputParallelismProperty
	_jsii_.Get(
		j,
		"inputParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputProcessingConfiguration() TfApplication_InputProcessingConfigurationPropertyOutputReference {
	var returns TfApplication_InputProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"inputProcessingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputProcessingConfigurationInput() *TfApplication_InputProcessingConfigurationProperty {
	var returns *TfApplication_InputProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"inputProcessingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputSchema() TfApplication_InputSchemaPropertyOutputReference {
	var returns TfApplication_InputSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"inputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputSchemaInput() *TfApplication_InputSchemaProperty {
	var returns *TfApplication_InputSchemaProperty
	_jsii_.Get(
		j,
		"inputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputStartingPositionConfiguration() TfApplication_InputStartingPositionConfigurationPropertyList {
	var returns TfApplication_InputStartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"inputStartingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InputStartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputStartingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) InternalValue() *TfApplication_InputProperty {
	var returns *TfApplication_InputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) KinesisFirehoseInput() TfApplication_KinesisFirehoseInputPropertyOutputReference {
	var returns TfApplication_KinesisFirehoseInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) KinesisFirehoseInputInput() *TfApplication_KinesisFirehoseInputProperty {
	var returns *TfApplication_KinesisFirehoseInputProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) KinesisStreamsInput() TfApplication_KinesisStreamsInputPropertyOutputReference {
	var returns TfApplication_KinesisStreamsInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) KinesisStreamsInputInput() *TfApplication_KinesisStreamsInputProperty {
	var returns *TfApplication_KinesisStreamsInputProperty
	_jsii_.Get(
		j,
		"kinesisStreamsInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_InputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_InputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_InputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_InputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_InputPropertyOutputReference_Override(t TfApplication_InputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.TfApplication.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetInternalValue(val *TfApplication_InputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutInputParallelism(value *TfApplication_InputParallelismProperty) {
	if err := t.validatePutInputParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputParallelism",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutInputProcessingConfiguration(value *TfApplication_InputProcessingConfigurationProperty) {
	if err := t.validatePutInputProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutInputSchema(value *TfApplication_InputSchemaProperty) {
	if err := t.validatePutInputSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutInputStartingPositionConfiguration(value interface{}) {
	if err := t.validatePutInputStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutKinesisFirehoseInput(value *TfApplication_KinesisFirehoseInputProperty) {
	if err := t.validatePutKinesisFirehoseInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisFirehoseInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) PutKinesisStreamsInput(value *TfApplication_KinesisStreamsInputProperty) {
	if err := t.validatePutKinesisStreamsInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisStreamsInput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ResetInputParallelism() {
	_jsii_.InvokeVoid(
		t,
		"resetInputParallelism",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ResetInputProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetInputProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ResetInputStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetInputStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ResetKinesisFirehoseInput() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisFirehoseInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ResetKinesisStreamsInput() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisStreamsInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_InputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

