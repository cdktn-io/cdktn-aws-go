package kinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_InputPropertyOutputReference interface {
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
	InputParallelism() AwsApplication_InputParallelismPropertyOutputReference
	// Experimental.
	InputParallelismInput() *AwsApplication_InputParallelismProperty
	// Experimental.
	InputProcessingConfiguration() AwsApplication_InputProcessingConfigurationPropertyOutputReference
	// Experimental.
	InputProcessingConfigurationInput() *AwsApplication_InputProcessingConfigurationProperty
	// Experimental.
	InputSchema() AwsApplication_InputSchemaPropertyOutputReference
	// Experimental.
	InputSchemaInput() *AwsApplication_InputSchemaProperty
	// Experimental.
	InputStartingPositionConfiguration() AwsApplication_InputStartingPositionConfigurationPropertyList
	// Experimental.
	InputStartingPositionConfigurationInput() interface{}
	// Experimental.
	InternalValue() *AwsApplication_InputProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_InputProperty)
	// Experimental.
	KinesisFirehoseInput() AwsApplication_KinesisFirehoseInputPropertyOutputReference
	// Experimental.
	KinesisFirehoseInputInput() *AwsApplication_KinesisFirehoseInputProperty
	// Experimental.
	KinesisStreamsInput() AwsApplication_KinesisStreamsInputPropertyOutputReference
	// Experimental.
	KinesisStreamsInputInput() *AwsApplication_KinesisStreamsInputProperty
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
	PutInputParallelism(value *AwsApplication_InputParallelismProperty)
	// Experimental.
	PutInputProcessingConfiguration(value *AwsApplication_InputProcessingConfigurationProperty)
	// Experimental.
	PutInputSchema(value *AwsApplication_InputSchemaProperty)
	// Experimental.
	PutInputStartingPositionConfiguration(value interface{})
	// Experimental.
	PutKinesisFirehoseInput(value *AwsApplication_KinesisFirehoseInputProperty)
	// Experimental.
	PutKinesisStreamsInput(value *AwsApplication_KinesisStreamsInputProperty)
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

// The jsii proxy struct for AwsApplication_InputPropertyOutputReference
type jsiiProxy_AwsApplication_InputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InAppStreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inAppStreamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputParallelism() AwsApplication_InputParallelismPropertyOutputReference {
	var returns AwsApplication_InputParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"inputParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputParallelismInput() *AwsApplication_InputParallelismProperty {
	var returns *AwsApplication_InputParallelismProperty
	_jsii_.Get(
		j,
		"inputParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputProcessingConfiguration() AwsApplication_InputProcessingConfigurationPropertyOutputReference {
	var returns AwsApplication_InputProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"inputProcessingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputProcessingConfigurationInput() *AwsApplication_InputProcessingConfigurationProperty {
	var returns *AwsApplication_InputProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"inputProcessingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputSchema() AwsApplication_InputSchemaPropertyOutputReference {
	var returns AwsApplication_InputSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"inputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputSchemaInput() *AwsApplication_InputSchemaProperty {
	var returns *AwsApplication_InputSchemaProperty
	_jsii_.Get(
		j,
		"inputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputStartingPositionConfiguration() AwsApplication_InputStartingPositionConfigurationPropertyList {
	var returns AwsApplication_InputStartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"inputStartingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InputStartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputStartingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) InternalValue() *AwsApplication_InputProperty {
	var returns *AwsApplication_InputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) KinesisFirehoseInput() AwsApplication_KinesisFirehoseInputPropertyOutputReference {
	var returns AwsApplication_KinesisFirehoseInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) KinesisFirehoseInputInput() *AwsApplication_KinesisFirehoseInputProperty {
	var returns *AwsApplication_KinesisFirehoseInputProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) KinesisStreamsInput() AwsApplication_KinesisStreamsInputPropertyOutputReference {
	var returns AwsApplication_KinesisStreamsInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) KinesisStreamsInputInput() *AwsApplication_KinesisStreamsInputProperty {
	var returns *AwsApplication_KinesisStreamsInputProperty
	_jsii_.Get(
		j,
		"kinesisStreamsInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_InputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_InputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_InputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_InputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_InputPropertyOutputReference_Override(a AwsApplication_InputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsApplication.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetInternalValue(val *AwsApplication_InputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutInputParallelism(value *AwsApplication_InputParallelismProperty) {
	if err := a.validatePutInputParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputParallelism",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutInputProcessingConfiguration(value *AwsApplication_InputProcessingConfigurationProperty) {
	if err := a.validatePutInputProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutInputSchema(value *AwsApplication_InputSchemaProperty) {
	if err := a.validatePutInputSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutInputStartingPositionConfiguration(value interface{}) {
	if err := a.validatePutInputStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutKinesisFirehoseInput(value *AwsApplication_KinesisFirehoseInputProperty) {
	if err := a.validatePutKinesisFirehoseInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehoseInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) PutKinesisStreamsInput(value *AwsApplication_KinesisStreamsInputProperty) {
	if err := a.validatePutKinesisStreamsInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamsInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ResetInputParallelism() {
	_jsii_.InvokeVoid(
		a,
		"resetInputParallelism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ResetInputProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInputProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ResetInputStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInputStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ResetKinesisFirehoseInput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehoseInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ResetKinesisStreamsInput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamsInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_InputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

