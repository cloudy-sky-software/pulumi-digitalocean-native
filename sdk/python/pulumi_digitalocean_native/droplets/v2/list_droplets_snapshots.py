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
from ._enums import *

__all__ = [
    'ListDropletsSnapshotsResult',
    'AwaitableListDropletsSnapshotsResult',
    'list_droplets_snapshots',
    'list_droplets_snapshots_output',
]

@pulumi.output_type
class ListDropletsSnapshotsResult:
    def __init__(__self__, items=None):
        if items and not isinstance(items, dict):
            raise TypeError("Expected argument 'items' to be a dict")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def items(self) -> 'outputs.ListDropletsSnapshotsItems':
        return pulumi.get(self, "items")


class AwaitableListDropletsSnapshotsResult(ListDropletsSnapshotsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDropletsSnapshotsResult(
            items=self.items)


def list_droplets_snapshots(droplet_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDropletsSnapshotsResult:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:listDropletsSnapshots', __args__, opts=opts, typ=ListDropletsSnapshotsResult).value

    return AwaitableListDropletsSnapshotsResult(
        items=pulumi.get(__ret__, 'items'))


@_utilities.lift_output_func(list_droplets_snapshots)
def list_droplets_snapshots_output(droplet_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ListDropletsSnapshotsResult]:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    ...
