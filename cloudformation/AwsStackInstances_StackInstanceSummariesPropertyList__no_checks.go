//go:build no_runtime_type_checking

package cloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsStackInstances_StackInstanceSummariesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsStackInstances_StackInstanceSummariesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

