package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference interface {
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
	InternalValue() *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty
	// Experimental.
	SetInternalValue(val *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty)
	// Experimental.
	OnDemandSpecification() TfCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList
	// Experimental.
	OnDemandSpecificationInput() interface{}
	// Experimental.
	SpotSpecification() TfCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList
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

// The jsii proxy struct for TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference
type jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InternalValue() *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty {
	var returns *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) OnDemandSpecification() TfCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList {
	var returns TfCluster_CoreInstanceFleetLaunchSpecificationsOnDemandSpecificationPropertyList
	_jsii_.Get(
		j,
		"onDemandSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) OnDemandSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) SpotSpecification() TfCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList {
	var returns TfCluster_CoreInstanceFleetLaunchSpecificationsSpotSpecificationPropertyList
	_jsii_.Get(
		j,
		"spotSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) SpotSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.CoreInstanceFleetLaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference_Override(t TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.TfCluster.CoreInstanceFleetLaunchSpecificationsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetInternalValue(val *TfCluster_CoreInstanceFleetLaunchSpecificationsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) PutOnDemandSpecification(value interface{}) {
	if err := t.validatePutOnDemandSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOnDemandSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) PutSpotSpecification(value interface{}) {
	if err := t.validatePutSpotSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpotSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ResetOnDemandSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ResetSpotSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_CoreInstanceFleetLaunchSpecificationsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

