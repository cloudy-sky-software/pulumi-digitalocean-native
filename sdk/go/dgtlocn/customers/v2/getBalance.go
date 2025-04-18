// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/cloudy-sky-software/pulumi-digitalocean-native/sdk/go/dgtlocn/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBalance(ctx *pulumi.Context, args *GetBalanceArgs, opts ...pulumi.InvokeOption) (*GetBalanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBalanceResult
	err := ctx.Invoke("digitalocean-native:customers/v2:getBalance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetBalanceArgs struct {
}

type GetBalanceResult struct {
	// Current balance of the customer's most recent billing activity.  Does not reflect `month_to_date_usage`.
	AccountBalance *string `pulumi:"accountBalance"`
	// The time at which balances were most recently generated.
	GeneratedAt *string `pulumi:"generatedAt"`
	// Balance as of the `generated_at` time.  This value includes the `account_balance` and `month_to_date_usage`.
	MonthToDateBalance *string `pulumi:"monthToDateBalance"`
	// Amount used in the current billing period as of the `generated_at` time.
	MonthToDateUsage *string `pulumi:"monthToDateUsage"`
}

func GetBalanceOutput(ctx *pulumi.Context, args GetBalanceOutputArgs, opts ...pulumi.InvokeOption) GetBalanceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBalanceResultOutput, error) {
			args := v.(GetBalanceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("digitalocean-native:customers/v2:getBalance", args, GetBalanceResultOutput{}, options).(GetBalanceResultOutput), nil
		}).(GetBalanceResultOutput)
}

type GetBalanceOutputArgs struct {
}

func (GetBalanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBalanceArgs)(nil)).Elem()
}

type GetBalanceResultOutput struct{ *pulumi.OutputState }

func (GetBalanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBalanceResult)(nil)).Elem()
}

func (o GetBalanceResultOutput) ToGetBalanceResultOutput() GetBalanceResultOutput {
	return o
}

func (o GetBalanceResultOutput) ToGetBalanceResultOutputWithContext(ctx context.Context) GetBalanceResultOutput {
	return o
}

// Current balance of the customer's most recent billing activity.  Does not reflect `month_to_date_usage`.
func (o GetBalanceResultOutput) AccountBalance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBalanceResult) *string { return v.AccountBalance }).(pulumi.StringPtrOutput)
}

// The time at which balances were most recently generated.
func (o GetBalanceResultOutput) GeneratedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBalanceResult) *string { return v.GeneratedAt }).(pulumi.StringPtrOutput)
}

// Balance as of the `generated_at` time.  This value includes the `account_balance` and `month_to_date_usage`.
func (o GetBalanceResultOutput) MonthToDateBalance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBalanceResult) *string { return v.MonthToDateBalance }).(pulumi.StringPtrOutput)
}

// Amount used in the current billing period as of the `generated_at` time.
func (o GetBalanceResultOutput) MonthToDateUsage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBalanceResult) *string { return v.MonthToDateUsage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBalanceResultOutput{})
}
