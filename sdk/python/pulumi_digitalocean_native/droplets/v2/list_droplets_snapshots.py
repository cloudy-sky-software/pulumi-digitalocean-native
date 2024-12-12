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
    'ListDropletsSnapshotsItems',
    'AwaitableListDropletsSnapshotsItems',
    'list_droplets_snapshots',
    'list_droplets_snapshots_output',
]

@pulumi.output_type
class ListDropletsSnapshotsItems:
    def __init__(__self__, links=None, meta=None, snapshots=None):
        if links and not isinstance(links, dict):
            raise TypeError("Expected argument 'links' to be a dict")
        pulumi.set(__self__, "links", links)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if snapshots and not isinstance(snapshots, list):
            raise TypeError("Expected argument 'snapshots' to be a list")
        pulumi.set(__self__, "snapshots", snapshots)

    @property
    @pulumi.getter
    def links(self) -> Optional['outputs.PageLinks']:
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def meta(self) -> 'outputs.MetaMeta':
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter
    def snapshots(self) -> Optional[Sequence['outputs.DropletSnapshot']]:
        return pulumi.get(self, "snapshots")


class AwaitableListDropletsSnapshotsItems(ListDropletsSnapshotsItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDropletsSnapshotsItems(
            links=self.links,
            meta=self.meta,
            snapshots=self.snapshots)


def list_droplets_snapshots(droplet_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListDropletsSnapshotsItems:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:droplets/v2:listDropletsSnapshots', __args__, opts=opts, typ=ListDropletsSnapshotsItems).value

    return AwaitableListDropletsSnapshotsItems(
        links=pulumi.get(__ret__, 'links'),
        meta=pulumi.get(__ret__, 'meta'),
        snapshots=pulumi.get(__ret__, 'snapshots'))
def list_droplets_snapshots_output(droplet_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListDropletsSnapshotsItems]:
    """
    Use this data source to access information about an existing resource.

    :param str droplet_id: A unique identifier for a Droplet instance.
    """
    __args__ = dict()
    __args__['dropletId'] = droplet_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:droplets/v2:listDropletsSnapshots', __args__, opts=opts, typ=ListDropletsSnapshotsItems)
    return __ret__.apply(lambda __response__: ListDropletsSnapshotsItems(
        links=pulumi.get(__response__, 'links'),
        meta=pulumi.get(__response__, 'meta'),
        snapshots=pulumi.get(__response__, 'snapshots')))
