package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList
type jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList {
	_init_.Initialize()

	if err := validateNewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsAttachedToPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList_Override(a AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEc2NetworkInsightsAnalysis.ReturnPathComponentsAttachedToPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) Get(index *float64) AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2NetworkInsightsAnalysis_ReturnPathComponentsAttachedToPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

