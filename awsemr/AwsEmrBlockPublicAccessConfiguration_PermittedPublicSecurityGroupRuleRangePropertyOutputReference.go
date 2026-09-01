package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxRange() *float64
	// Experimental.
	SetMaxRange(val *float64)
	// Experimental.
	MaxRangeInput() *float64
	// Experimental.
	MinRange() *float64
	// Experimental.
	SetMinRange(val *float64)
	// Experimental.
	MinRangeInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference
type jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) MaxRange() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) MaxRangeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) MinRange() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) MinRangeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrBlockPublicAccessConfiguration.PermittedPublicSecurityGroupRuleRangePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference_Override(a AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrBlockPublicAccessConfiguration.PermittedPublicSecurityGroupRuleRangePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetMaxRange(val *float64) {
	if err := j.validateSetMaxRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRange",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetMinRange(val *float64) {
	if err := j.validateSetMinRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRange",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrBlockPublicAccessConfiguration_PermittedPublicSecurityGroupRuleRangePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

