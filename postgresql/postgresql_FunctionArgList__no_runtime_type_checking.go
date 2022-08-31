//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt postgresql Provider for Terraform CDK (cdktf)
package postgresql

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionArgList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FunctionArgList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FunctionArgList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FunctionArgList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FunctionArgList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FunctionArgList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFunctionArgListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

