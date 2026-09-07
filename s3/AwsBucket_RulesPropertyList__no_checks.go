//go:build no_runtime_type_checking

package s3

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsBucket_RulesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_RulesPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsBucket_RulesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_RulesPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_RulesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_RulesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsBucket_RulesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsBucket_RulesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

