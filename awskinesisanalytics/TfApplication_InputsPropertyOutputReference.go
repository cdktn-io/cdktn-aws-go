package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApplication_InputsPropertyOutputReference interface {
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
	Id() *string
	// Experimental.
	InternalValue() *TfApplication_InputsProperty
	// Experimental.
	SetInternalValue(val *TfApplication_InputsProperty)
	// Experimental.
	KinesisFirehose() TfApplication_InputsKinesisFirehosePropertyOutputReference
	// Experimental.
	KinesisFirehoseInput() *TfApplication_InputsKinesisFirehoseProperty
	// Experimental.
	KinesisStream() TfApplication_InputsKinesisStreamPropertyOutputReference
	// Experimental.
	KinesisStreamInput() *TfApplication_InputsKinesisStreamProperty
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// Experimental.
	Parallelism() TfApplication_ParallelismPropertyOutputReference
	// Experimental.
	ParallelismInput() *TfApplication_ParallelismProperty
	// Experimental.
	ProcessingConfiguration() TfApplication_ProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *TfApplication_ProcessingConfigurationProperty
	// Experimental.
	Schema() TfApplication_InputsSchemaPropertyOutputReference
	// Experimental.
	SchemaInput() *TfApplication_InputsSchemaProperty
	// Experimental.
	StartingPositionConfiguration() TfApplication_StartingPositionConfigurationPropertyList
	// Experimental.
	StartingPositionConfigurationInput() interface{}
	// Experimental.
	StreamNames() *[]*string
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
	PutKinesisFirehose(value *TfApplication_InputsKinesisFirehoseProperty)
	// Experimental.
	PutKinesisStream(value *TfApplication_InputsKinesisStreamProperty)
	// Experimental.
	PutParallelism(value *TfApplication_ParallelismProperty)
	// Experimental.
	PutProcessingConfiguration(value *TfApplication_ProcessingConfigurationProperty)
	// Experimental.
	PutSchema(value *TfApplication_InputsSchemaProperty)
	// Experimental.
	PutStartingPositionConfiguration(value interface{})
	// Experimental.
	ResetKinesisFirehose()
	// Experimental.
	ResetKinesisStream()
	// Experimental.
	ResetParallelism()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetStartingPositionConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApplication_InputsPropertyOutputReference
type jsiiProxy_TfApplication_InputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) InternalValue() *TfApplication_InputsProperty {
	var returns *TfApplication_InputsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) KinesisFirehose() TfApplication_InputsKinesisFirehosePropertyOutputReference {
	var returns TfApplication_InputsKinesisFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) KinesisFirehoseInput() *TfApplication_InputsKinesisFirehoseProperty {
	var returns *TfApplication_InputsKinesisFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) KinesisStream() TfApplication_InputsKinesisStreamPropertyOutputReference {
	var returns TfApplication_InputsKinesisStreamPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) KinesisStreamInput() *TfApplication_InputsKinesisStreamProperty {
	var returns *TfApplication_InputsKinesisStreamProperty
	_jsii_.Get(
		j,
		"kinesisStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) Parallelism() TfApplication_ParallelismPropertyOutputReference {
	var returns TfApplication_ParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) ParallelismInput() *TfApplication_ParallelismProperty {
	var returns *TfApplication_ParallelismProperty
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) ProcessingConfiguration() TfApplication_ProcessingConfigurationPropertyOutputReference {
	var returns TfApplication_ProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) ProcessingConfigurationInput() *TfApplication_ProcessingConfigurationProperty {
	var returns *TfApplication_ProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) Schema() TfApplication_InputsSchemaPropertyOutputReference {
	var returns TfApplication_InputsSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) SchemaInput() *TfApplication_InputsSchemaProperty {
	var returns *TfApplication_InputsSchemaProperty
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) StartingPositionConfiguration() TfApplication_StartingPositionConfigurationPropertyList {
	var returns TfApplication_StartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"startingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) StartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) StreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApplication_InputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApplication_InputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApplication_InputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApplication_InputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApplication_InputsPropertyOutputReference_Override(t TfApplication_InputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.TfApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetInternalValue(val *TfApplication_InputsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApplication_InputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutKinesisFirehose(value *TfApplication_InputsKinesisFirehoseProperty) {
	if err := t.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutKinesisStream(value *TfApplication_InputsKinesisStreamProperty) {
	if err := t.validatePutKinesisStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisStream",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutParallelism(value *TfApplication_ParallelismProperty) {
	if err := t.validatePutParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParallelism",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutProcessingConfiguration(value *TfApplication_ProcessingConfigurationProperty) {
	if err := t.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutSchema(value *TfApplication_InputsSchemaProperty) {
	if err := t.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) PutStartingPositionConfiguration(value interface{}) {
	if err := t.validatePutStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ResetKinesisStream() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisStream",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelism",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ResetStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApplication_InputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

