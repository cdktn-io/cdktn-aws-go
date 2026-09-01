package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference interface {
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
	InternalValue() *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty
	// Experimental.
	SetInternalValue(val *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty)
	// Experimental.
	OnDemandSpecification() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList
	// Experimental.
	OnDemandSpecificationInput() interface{}
	// Experimental.
	SpotSpecification() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList
	// Experimental.
	SpotSpecificationInput() interface{}
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
	PutOnDemandSpecification(value interface{})
	// Experimental.
	PutSpotSpecification(value interface{})
	// Experimental.
	ResetOnDemandSpecification()
	// Experimental.
	ResetSpotSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
type jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InternalValue() *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty {
	var returns *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) OnDemandSpecification() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList {
	var returns AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) OnDemandSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) SpotSpecification() AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList {
	var returns AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) SpotSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.CoreInstanceFleetLaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference_Override(a AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.CoreInstanceFleetLaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetInternalValue(val *AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) PutOnDemandSpecification(value interface{}) {
	if err := a.validatePutOnDemandSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOnDemandSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) PutSpotSpecification(value interface{}) {
	if err := a.validatePutSpotSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpotSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

