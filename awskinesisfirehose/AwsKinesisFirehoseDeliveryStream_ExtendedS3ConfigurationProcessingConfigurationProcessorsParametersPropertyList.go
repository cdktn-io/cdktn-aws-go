package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList interface {
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
	Get(index *float64) AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList
type jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList {
	_init_.Initialize()

	if err := validateNewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList_Override(a AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) Get(index *float64) AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProcessorsParametersPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

