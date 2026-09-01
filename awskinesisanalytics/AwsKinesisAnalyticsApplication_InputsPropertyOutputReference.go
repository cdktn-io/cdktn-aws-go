package awskinesisanalytics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalytics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisAnalyticsApplication_InputsPropertyOutputReference interface {
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
	InternalValue() *AwsKinesisAnalyticsApplication_InputsProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisAnalyticsApplication_InputsProperty)
	// Experimental.
	KinesisFirehose() AwsKinesisAnalyticsApplication_InputsKinesisFirehosePropertyOutputReference
	// Experimental.
	KinesisFirehoseInput() *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty
	// Experimental.
	KinesisStream() AwsKinesisAnalyticsApplication_InputsKinesisStreamPropertyOutputReference
	// Experimental.
	KinesisStreamInput() *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// Experimental.
	Parallelism() AwsKinesisAnalyticsApplication_ParallelismPropertyOutputReference
	// Experimental.
	ParallelismInput() *AwsKinesisAnalyticsApplication_ParallelismProperty
	// Experimental.
	ProcessingConfiguration() AwsKinesisAnalyticsApplication_ProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty
	// Experimental.
	Schema() AwsKinesisAnalyticsApplication_InputsSchemaPropertyOutputReference
	// Experimental.
	SchemaInput() *AwsKinesisAnalyticsApplication_InputsSchemaProperty
	// Experimental.
	StartingPositionConfiguration() AwsKinesisAnalyticsApplication_StartingPositionConfigurationPropertyList
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
	PutKinesisFirehose(value *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty)
	// Experimental.
	PutKinesisStream(value *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty)
	// Experimental.
	PutParallelism(value *AwsKinesisAnalyticsApplication_ParallelismProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty)
	// Experimental.
	PutSchema(value *AwsKinesisAnalyticsApplication_InputsSchemaProperty)
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

// The jsii proxy struct for AwsKinesisAnalyticsApplication_InputsPropertyOutputReference
type jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) InternalValue() *AwsKinesisAnalyticsApplication_InputsProperty {
	var returns *AwsKinesisAnalyticsApplication_InputsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) KinesisFirehose() AwsKinesisAnalyticsApplication_InputsKinesisFirehosePropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_InputsKinesisFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) KinesisFirehoseInput() *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty {
	var returns *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) KinesisStream() AwsKinesisAnalyticsApplication_InputsKinesisStreamPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_InputsKinesisStreamPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) KinesisStreamInput() *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty {
	var returns *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty
	_jsii_.Get(
		j,
		"kinesisStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) Parallelism() AwsKinesisAnalyticsApplication_ParallelismPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_ParallelismPropertyOutputReference
	_jsii_.Get(
		j,
		"parallelism",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ParallelismInput() *AwsKinesisAnalyticsApplication_ParallelismProperty {
	var returns *AwsKinesisAnalyticsApplication_ParallelismProperty
	_jsii_.Get(
		j,
		"parallelismInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ProcessingConfiguration() AwsKinesisAnalyticsApplication_ProcessingConfigurationPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_ProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ProcessingConfigurationInput() *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty {
	var returns *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) Schema() AwsKinesisAnalyticsApplication_InputsSchemaPropertyOutputReference {
	var returns AwsKinesisAnalyticsApplication_InputsSchemaPropertyOutputReference
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) SchemaInput() *AwsKinesisAnalyticsApplication_InputsSchemaProperty {
	var returns *AwsKinesisAnalyticsApplication_InputsSchemaProperty
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) StartingPositionConfiguration() AwsKinesisAnalyticsApplication_StartingPositionConfigurationPropertyList {
	var returns AwsKinesisAnalyticsApplication_StartingPositionConfigurationPropertyList
	_jsii_.Get(
		j,
		"startingPositionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) StartingPositionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startingPositionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) StreamNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisAnalyticsApplication_InputsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisAnalyticsApplication_InputsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisAnalyticsApplication_InputsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsKinesisAnalyticsApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisAnalyticsApplication_InputsPropertyOutputReference_Override(a AwsKinesisAnalyticsApplication_InputsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics.AwsKinesisAnalyticsApplication.InputsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetInternalValue(val *AwsKinesisAnalyticsApplication_InputsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutKinesisFirehose(value *AwsKinesisAnalyticsApplication_InputsKinesisFirehoseProperty) {
	if err := a.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutKinesisStream(value *AwsKinesisAnalyticsApplication_InputsKinesisStreamProperty) {
	if err := a.validatePutKinesisStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutParallelism(value *AwsKinesisAnalyticsApplication_ParallelismProperty) {
	if err := a.validatePutParallelismParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParallelism",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutProcessingConfiguration(value *AwsKinesisAnalyticsApplication_ProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutSchema(value *AwsKinesisAnalyticsApplication_InputsSchemaProperty) {
	if err := a.validatePutSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchema",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) PutStartingPositionConfiguration(value interface{}) {
	if err := a.validatePutStartingPositionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStartingPositionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ResetKinesisStream() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ResetParallelism() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelism",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ResetStartingPositionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPositionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisAnalyticsApplication_InputsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

