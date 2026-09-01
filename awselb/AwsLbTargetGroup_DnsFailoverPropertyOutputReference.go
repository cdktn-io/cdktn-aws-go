package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLbTargetGroup_DnsFailoverPropertyOutputReference interface {
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
	InternalValue() *AwsLbTargetGroup_DnsFailoverProperty
	// Experimental.
	SetInternalValue(val *AwsLbTargetGroup_DnsFailoverProperty)
	// Experimental.
	MinimumHealthyTargetsCount() *string
	// Experimental.
	SetMinimumHealthyTargetsCount(val *string)
	// Experimental.
	MinimumHealthyTargetsCountInput() *string
	// Experimental.
	MinimumHealthyTargetsPercentage() *string
	// Experimental.
	SetMinimumHealthyTargetsPercentage(val *string)
	// Experimental.
	MinimumHealthyTargetsPercentageInput() *string
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
	ResetMinimumHealthyTargetsCount()
	// Experimental.
	ResetMinimumHealthyTargetsPercentage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLbTargetGroup_DnsFailoverPropertyOutputReference
type jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) InternalValue() *AwsLbTargetGroup_DnsFailoverProperty {
	var returns *AwsLbTargetGroup_DnsFailoverProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) MinimumHealthyTargetsCount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) MinimumHealthyTargetsCountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) MinimumHealthyTargetsPercentage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) MinimumHealthyTargetsPercentageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumHealthyTargetsPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLbTargetGroup_DnsFailoverPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLbTargetGroup_DnsFailoverPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLbTargetGroup_DnsFailoverPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbTargetGroup.DnsFailoverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLbTargetGroup_DnsFailoverPropertyOutputReference_Override(a AwsLbTargetGroup_DnsFailoverPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsLbTargetGroup.DnsFailoverPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetInternalValue(val *AwsLbTargetGroup_DnsFailoverProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetMinimumHealthyTargetsCount(val *string) {
	if err := j.validateSetMinimumHealthyTargetsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyTargetsCount",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetMinimumHealthyTargetsPercentage(val *string) {
	if err := j.validateSetMinimumHealthyTargetsPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumHealthyTargetsPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ResetMinimumHealthyTargetsCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumHealthyTargetsCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ResetMinimumHealthyTargetsPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumHealthyTargetsPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLbTargetGroup_DnsFailoverPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

