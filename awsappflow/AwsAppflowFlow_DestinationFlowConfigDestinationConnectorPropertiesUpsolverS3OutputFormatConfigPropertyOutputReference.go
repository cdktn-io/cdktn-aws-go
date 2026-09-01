package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AggregationConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference
	// Experimental.
	AggregationConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty
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
	InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty)
	// Experimental.
	PrefixConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference
	// Experimental.
	PrefixConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty
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
	PutAggregationConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty)
	// Experimental.
	PutPrefixConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty)
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

// The jsii proxy struct for AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) AggregationConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"aggregationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) AggregationConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty
	_jsii_.Get(
		j,
		"aggregationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) FileType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) FileTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PrefixConfig() AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference {
	var returns AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PrefixConfigInput() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty
	_jsii_.Get(
		j,
		"prefixConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference_Override(a AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetFileType(val *string) {
	if err := j.validateSetFileTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileType",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PutAggregationConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigAggregationConfigProperty) {
	if err := a.validatePutAggregationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAggregationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) PutPrefixConfig(value *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPrefixConfigProperty) {
	if err := a.validatePutPrefixConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefixConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ResetAggregationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAggregationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ResetFileType() {
	_jsii_.InvokeVoid(
		a,
		"resetFileType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesUpsolverS3OutputFormatConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

