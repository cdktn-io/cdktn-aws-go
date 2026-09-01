package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
	// Experimental.
	AccessConfigInput() interface{}
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
	EventRecordFormat() *string
	// Experimental.
	SetEventRecordFormat(val *string)
	// Experimental.
	EventRecordFormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	// Experimental.
	SetInternalValue(val *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty)
	// Experimental.
	SchemaRegistryUri() *string
	// Experimental.
	SetSchemaRegistryUri(val *string)
	// Experimental.
	SchemaRegistryUriInput() *string
	// Experimental.
	SchemaValidationConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
	// Experimental.
	SchemaValidationConfigInput() interface{}
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
	PutAccessConfig(value interface{})
	// Experimental.
	PutSchemaValidationConfig(value interface{})
	// Experimental.
	ResetAccessConfig()
	// Experimental.
	ResetEventRecordFormat()
	// Experimental.
	ResetSchemaRegistryUri()
	// Experimental.
	ResetSchemaValidationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference
type jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList {
	var returns AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigAccessConfigPropertyList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) EventRecordFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventRecordFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InternalValue() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty {
	var returns *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaRegistryUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaRegistryUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList {
	var returns AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList
	_jsii_.Get(
		j,
		"schemaValidationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) SchemaValidationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaValidationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping.AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference_Override(a AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping.AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetEventRecordFormat(val *string) {
	if err := j.validateSetEventRecordFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventRecordFormat",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetInternalValue(val *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetSchemaRegistryUri(val *string) {
	if err := j.validateSetSchemaRegistryUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaRegistryUri",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutAccessConfig(value interface{}) {
	if err := a.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) PutSchemaValidationConfig(value interface{}) {
	if err := a.validatePutSchemaValidationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaValidationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetEventRecordFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetEventRecordFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaRegistryUri() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaRegistryUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ResetSchemaValidationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaValidationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigSchemaRegistryConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

