package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisanalyticsv2Application_InputPropertyOutputReference interface {
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
	InputParallelism() AwsKinesisanalyticsv2Application_InputParallelismPropertyOutputReference
	// Experimental.
	InputParallelismInput() *AwsKinesisanalyticsv2Application_InputParallelismProperty
	// Experimental.
	InputProcessingConfiguration() AwsKinesisanalyticsv2Application_InputProcessingConfigurationPropertyOutputReference
	// Experimental.
	InputProcessingConfigurationInput() *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty
	// Experimental.
	InputSchema() AwsKinesisanalyticsv2Application_InputSchemaPropertyOutputReference
	// Experimental.
	InputSchemaInput() *AwsKinesisanalyticsv2Application_InputSchemaProperty
	// Experimental.
	InputStartingPositionConfiguration() AwsKinesisanalyticsv2Application_InputStartingPositionConfigurationPropertyList
	// Experimental.
	InputStartingPositionConfigurationInput() interface{}
	// Experimental.
	InternalValue() *AwsKinesisanalyticsv2Application_InputProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisanalyticsv2Application_InputProperty)
	// Experimental.
	KinesisFirehoseInput() AwsKinesisanalyticsv2Application_KinesisFirehoseInputPropertyOutputReference
	// Experimental.
	KinesisFirehoseInputInput() *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty
	// Experimental.
	KinesisStreamsInput() AwsKinesisanalyticsv2Application_KinesisStreamsInputPropertyOutputReference
	// Experimental.
	KinesisStreamsInputInput() *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty
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
	PutInputParallelism(value *AwsKinesisanalyticsv2Application_InputParallelismProperty)
	// Experimental.
	PutInputProcessingConfiguration(value *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty)
	// Experimental.
	PutInputSchema(value *AwsKinesisanalyticsv2Application_InputSchemaProperty)
	// Experimental.
	PutInputStartingPositionConfiguration(value interface{})
	// Experimental.
	PutKinesisFirehoseInput(value *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty)
	// Experimental.
	PutKinesisStreamsInput(value *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty)
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

// The jsii proxy struct for AwsKinesisanalyticsv2Application_InputPropertyOutputReference
type jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InAppStreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inAppStreamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputParallelism() AwsKinesisanalyticsv2Application_InputParallelismPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_InputParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"inputParallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputParallelismInput() *AwsKinesisanalyticsv2Application_InputParallelismProperty {
	var returns *AwsKinesisanalyticsv2Application_InputParallelismProperty
	_jsii_.Get(
		j,
		"inputParallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputProcessingConfiguration() AwsKinesisanalyticsv2Application_InputProcessingConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_InputProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"inputProcessingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputProcessingConfigurationInput() *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"inputProcessingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputSchema() AwsKinesisanalyticsv2Application_InputSchemaPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_InputSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"inputSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputSchemaInput() *AwsKinesisanalyticsv2Application_InputSchemaProperty {
	var returns *AwsKinesisanalyticsv2Application_InputSchemaProperty
	_jsii_.Get(
		j,
		"inputSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputStartingPositionConfiguration() AwsKinesisanalyticsv2Application_InputStartingPositionConfigurationPropertyList {
	var returns AwsKinesisanalyticsv2Application_InputStartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"inputStartingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InputStartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputStartingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InternalValue() *AwsKinesisanalyticsv2Application_InputProperty {
	var returns *AwsKinesisanalyticsv2Application_InputProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) KinesisFirehoseInput() AwsKinesisanalyticsv2Application_KinesisFirehoseInputPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_KinesisFirehoseInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) KinesisFirehoseInputInput() *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty {
	var returns *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) KinesisStreamsInput() AwsKinesisanalyticsv2Application_KinesisStreamsInputPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_KinesisStreamsInputPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) KinesisStreamsInputInput() *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty {
	var returns *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty
	_jsii_.Get(
		j,
		"kinesisStreamsInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisanalyticsv2Application_InputPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisanalyticsv2Application_InputPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisanalyticsv2Application_InputPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisanalyticsv2Application_InputPropertyOutputReference_Override(a AwsKinesisanalyticsv2Application_InputPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.InputPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetInternalValue(val *AwsKinesisanalyticsv2Application_InputProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutInputParallelism(value *AwsKinesisanalyticsv2Application_InputParallelismProperty) {
	if err := a.validatePutInputParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputParallelism",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutInputProcessingConfiguration(value *AwsKinesisanalyticsv2Application_InputProcessingConfigurationProperty) {
	if err := a.validatePutInputProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutInputSchema(value *AwsKinesisanalyticsv2Application_InputSchemaProperty) {
	if err := a.validatePutInputSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutInputStartingPositionConfiguration(value interface{}) {
	if err := a.validatePutInputStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutKinesisFirehoseInput(value *AwsKinesisanalyticsv2Application_KinesisFirehoseInputProperty) {
	if err := a.validatePutKinesisFirehoseInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehoseInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) PutKinesisStreamsInput(value *AwsKinesisanalyticsv2Application_KinesisStreamsInputProperty) {
	if err := a.validatePutKinesisStreamsInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamsInput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ResetInputParallelism() {
	_jsii_.InvokeVoid(
		a,
		"resetInputParallelism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ResetInputProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInputProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ResetInputStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInputStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ResetKinesisFirehoseInput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehoseInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ResetKinesisStreamsInput() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamsInput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_InputPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

