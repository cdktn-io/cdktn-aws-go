package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AggregationType() *string
	// Experimental.
	SetAggregationType(val *string)
	// Experimental.
	AggregationTypeInput() *string
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
	InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty
	// Experimental.
	SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty)
	// Experimental.
	TargetFileSize() *float64
	// Experimental.
	SetTargetFileSize(val *float64)
	// Experimental.
	TargetFileSizeInput() *float64
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
	ResetAggregationType()
	// Experimental.
	ResetTargetFileSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference
type jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) AggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) AggregationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) InternalValue() *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty {
	var returns *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) TargetFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) TargetFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference_Override(a AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsFlow.DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetAggregationType(val *string) {
	if err := j.validateSetAggregationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aggregationType",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetInternalValue(val *AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetTargetFileSize(val *float64) {
	if err := j.validateSetTargetFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFileSize",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ResetAggregationType() {
	_jsii_.InvokeVoid(
		a,
		"resetAggregationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ResetTargetFileSize() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetFileSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFlow_DestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigAggregationConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

