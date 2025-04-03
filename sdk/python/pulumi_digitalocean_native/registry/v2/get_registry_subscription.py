# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload, Awaitable
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from . import outputs

__all__ = [
    'GetRegistrySubscriptionProperties',
    'AwaitableGetRegistrySubscriptionProperties',
    'get_registry_subscription',
    'get_registry_subscription_output',
]

@pulumi.output_type
class GetRegistrySubscriptionProperties:
    def __init__(__self__, subscription=None):
        if subscription and not isinstance(subscription, dict):
            raise TypeError("Expected argument 'subscription' to be a dict")
        pulumi.set(__self__, "subscription", subscription)

    @property
    @pulumi.getter
    def subscription(self) -> Optional['outputs.Subscription']:
        return pulumi.get(self, "subscription")


class AwaitableGetRegistrySubscriptionProperties(GetRegistrySubscriptionProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegistrySubscriptionProperties(
            subscription=self.subscription)


def get_registry_subscription(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegistrySubscriptionProperties:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:registry/v2:getRegistrySubscription', __args__, opts=opts, typ=GetRegistrySubscriptionProperties).value

    return AwaitableGetRegistrySubscriptionProperties(
        subscription=pulumi.get(__ret__, 'subscription'))
def get_registry_subscription_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRegistrySubscriptionProperties]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:registry/v2:getRegistrySubscription', __args__, opts=opts, typ=GetRegistrySubscriptionProperties)
    return __ret__.apply(lambda __response__: GetRegistrySubscriptionProperties(
        subscription=pulumi.get(__response__, 'subscription')))
