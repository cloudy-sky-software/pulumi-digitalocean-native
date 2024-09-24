# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from ._enums import *

__all__ = [
    'ListDropletsNeighborsItems',
    'AwaitableListDropletsNeighborsItems',
    'list_droplets_neighbors',
    'list_droplets_neighbors_output',
]

@pulumi.output_type
class ListDropletsNeighborsItems:
    def __init__(__self__, droplets=None):
        if droplets and not isinstance(droplets, list):
            raise TypeError("Expected argument 'droplets' to be a list")
        pulumi.set(__self__, "droplets", droplets)

    @property
    @pulumi.getter
    def droplets(self) -> Optional[Sequence['outputs.Droplet']]:
        return pulumi.get(self, "droplets")


class AwaitableListDropletsNeighborsItems(ListDropletsNeighborsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDropletsNeighborsItems(
            droplets=self.droplets)


def list_droplets_neighbors(droplet_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDropletsNeighborsItems:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:listDropletsNeighbors', __args__, opts=opts, typ=ListDropletsNeighborsItems).value

    return AwaitableListDropletsNeighborsItems(
        droplets=pulumi.get(__ret__, 'droplets'))
def list_droplets_neighbors_output(droplet_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDropletsNeighborsItems]:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:droplets/v2:listDropletsNeighbors', __args__, opts=opts, typ=ListDropletsNeighborsItems)
    return __ret__.apply(lambda __response__: ListDropletsNeighborsItems(
        droplets=pulumi.get(__response__, 'droplets')))
