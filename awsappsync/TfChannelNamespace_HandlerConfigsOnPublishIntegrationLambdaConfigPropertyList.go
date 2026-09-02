package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList
type jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList {
	_init_.Initialize()

	if err := validateNewTfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfChannelNamespace.HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList_Override(t TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfChannelNamespace.HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) Get(index *float64) TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannelNamespace_HandlerConfigsOnPublishIntegrationLambdaConfigPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

