package kinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApplication_InputsPropertyOutputReference interface {
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
	InternalValue() *AwsApplication_InputsProperty
	// Experimental.
	SetInternalValue(val *AwsApplication_InputsProperty)
	// Experimental.
	KinesisFirehose() AwsApplication_InputsKinesisFirehosePropertyOutputReference
	// Experimental.
	KinesisFirehoseInput() *AwsApplication_InputsKinesisFirehoseProperty
	// Experimental.
	KinesisStream() AwsApplication_InputsKinesisStreamPropertyOutputReference
	// Experimental.
	KinesisStreamInput() *AwsApplication_InputsKinesisStreamProperty
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// Experimental.
	Parallelism() AwsApplication_ParallelismPropertyOutputReference
	// Experimental.
	ParallelismInput() *AwsApplication_ParallelismProperty
	// Experimental.
	ProcessingConfiguration() AwsApplication_ProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsApplication_ProcessingConfigurationProperty
	// Experimental.
	Schema() AwsApplication_InputsSchemaPropertyOutputReference
	// Experimental.
	SchemaInput() *AwsApplication_InputsSchemaProperty
	// Experimental.
	StartingPositionConfiguration() AwsApplication_StartingPositionConfigurationPropertyList
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
	PutKinesisFirehose(value *AwsApplication_InputsKinesisFirehoseProperty)
	// Experimental.
	PutKinesisStream(value *AwsApplication_InputsKinesisStreamProperty)
	// Experimental.
	PutParallelism(value *AwsApplication_ParallelismProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsApplication_ProcessingConfigurationProperty)
	// Experimental.
	PutSchema(value *AwsApplication_InputsSchemaProperty)
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

// The jsii proxy struct for AwsApplication_InputsPropertyOutputReference
type jsiiProxy_AwsApplication_InputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) InternalValue() *AwsApplication_InputsProperty {
	var returns *AwsApplication_InputsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) KinesisFirehose() AwsApplication_InputsKinesisFirehosePropertyOutputReference {
	var returns AwsApplication_InputsKinesisFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) KinesisFirehoseInput() *AwsApplication_InputsKinesisFirehoseProperty {
	var returns *AwsApplication_InputsKinesisFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) KinesisStream() AwsApplication_InputsKinesisStreamPropertyOutputReference {
	var returns AwsApplication_InputsKinesisStreamPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) KinesisStreamInput() *AwsApplication_InputsKinesisStreamProperty {
	var returns *AwsApplication_InputsKinesisStreamProperty
	_jsii_.Get(
		j,
		"kinesisStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) Parallelism() AwsApplication_ParallelismPropertyOutputReference {
	var returns AwsApplication_ParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ParallelismInput() *AwsApplication_ParallelismProperty {
	var returns *AwsApplication_ParallelismProperty
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ProcessingConfiguration() AwsApplication_ProcessingConfigurationPropertyOutputReference {
	var returns AwsApplication_ProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ProcessingConfigurationInput() *AwsApplication_ProcessingConfigurationProperty {
	var returns *AwsApplication_ProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) Schema() AwsApplication_InputsSchemaPropertyOutputReference {
	var returns AwsApplication_InputsSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) SchemaInput() *AwsApplication_InputsSchemaProperty {
	var returns *AwsApplication_InputsSchemaProperty
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) StartingPositionConfiguration() AwsApplication_StartingPositionConfigurationPropertyList {
	var returns AwsApplication_StartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"startingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) StartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) StreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApplication_InputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApplication_InputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApplication_InputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApplication_InputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApplication_InputsPropertyOutputReference_Override(a AwsApplication_InputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetInternalValue(val *AwsApplication_InputsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApplication_InputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutKinesisFirehose(value *AwsApplication_InputsKinesisFirehoseProperty) {
	if err := a.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutKinesisStream(value *AwsApplication_InputsKinesisStreamProperty) {
	if err := a.validatePutKinesisStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutParallelism(value *AwsApplication_ParallelismProperty) {
	if err := a.validatePutParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelism",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutProcessingConfiguration(value *AwsApplication_ProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutSchema(value *AwsApplication_InputsSchemaProperty) {
	if err := a.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) PutStartingPositionConfiguration(value interface{}) {
	if err := a.validatePutStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ResetKinesisStream() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ResetStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApplication_InputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

