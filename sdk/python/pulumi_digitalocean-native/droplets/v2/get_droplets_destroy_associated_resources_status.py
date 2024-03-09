# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetDropletsDestroyAssociatedResourcesStatusResult',
    'AwaitableGetDropletsDestroyAssociatedResourcesStatusResult',
    'get_droplets_destroy_associated_resources_status',
    'get_droplets_destroy_associated_resources_status_output',
]

@pulumi.output_type
class GetDropletsDestroyAssociatedResourcesStatusResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.AssociatedResourceStatus':
        return pulumi.get(self, "items")


class AwaitableGetDropletsDestroyAssociatedResourcesStatusResult(GetDropletsDestroyAssociatedResourcesStatusResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDropletsDestroyAssociatedResourcesStatusResult(
            items=self.items)


def get_droplets_destroy_associated_resources_status(droplet_id: Optional[str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDropletsDestroyAssociatedResourcesStatusResult:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:getDropletsDestroyAssociatedResourcesStatus', __args__, opts=opts, typ=GetDropletsDestroyAssociatedResourcesStatusResult).value

    return AwaitableGetDropletsDestroyAssociatedResourcesStatusResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(get_droplets_destroy_associated_resources_status)
def get_droplets_destroy_associated_resources_status_output(droplet_id: Optional[pulumi.Input[str]] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDropletsDestroyAssociatedResourcesStatusResult]:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    ...