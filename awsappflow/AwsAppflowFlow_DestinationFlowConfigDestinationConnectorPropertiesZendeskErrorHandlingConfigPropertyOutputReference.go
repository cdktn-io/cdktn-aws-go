package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BucketName() *string
	// Experimental.
	SetBucketName(val *string)
	// Experimental.
	BucketNameInput() *string
	// Experimental.
	BucketPrefix() *string
	// Experimental.
	SetBucketPrefix(val *string)
	// Experimental.
	BucketPrefixInput() *string
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
	FailOnFirstDestinationError() interface{}
	// Experimental.
	SetFailOnFirstDestinationError(val interface{})
	// Experimental.
	FailOnFirstDestinationErrorInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty
	// Experimental.
	SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty)
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
	ResetBucketName()
	// Experimental.
	ResetBucketPrefix()
	// Experimental.
	ResetFailOnFirstDestinationError()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference
type jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) BucketPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) BucketPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) FailOnFirstDestinationError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstDestinationError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) FailOnFirstDestinationErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failOnFirstDestinationErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) InternalValue() *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty {
	var returns *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference_Override(a AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsAppflowFlow.DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetBucketPrefix(val *string) {
	if err := j.validateSetBucketPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetFailOnFirstDestinationError(val interface{}) {
	if err := j.validateSetFailOnFirstDestinationErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failOnFirstDestinationError",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetInternalValue(val *AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ResetBucketName() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ResetBucketPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ResetFailOnFirstDestinationError() {
	_jsii_.InvokeVoid(
		a,
		"resetFailOnFirstDestinationError",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppflowFlow_DestinationFlowConfigDestinationConnectorPropertiesZendeskErrorHandlingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

