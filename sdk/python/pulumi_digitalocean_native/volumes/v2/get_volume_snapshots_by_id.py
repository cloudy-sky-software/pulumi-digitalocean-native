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
from ._enums import *

__all__ = [
    'GetVolumeSnapshotsByIdProperties',
    'AwaitableGetVolumeSnapshotsByIdProperties',
    'get_volume_snapshots_by_id',
    'get_volume_snapshots_by_id_output',
]

@pulumi.output_type
class GetVolumeSnapshotsByIdProperties:
    def __init__(__self__, snapshot=None):
        if snapshot and not isinstance(snapshot, dict):
            raise TypeError("Expected argument 'snapshot' to be a dict")
        pulumi.set(__self__, "snapshot", snapshot)

    @property
    @pulumi.getter
    def snapshot(self) -> Optional['outputs.Snapshots']:
        return pulumi.get(self, "snapshot")


class AwaitableGetVolumeSnapshotsByIdProperties(GetVolumeSnapshotsByIdProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeSnapshotsByIdProperties(
            snapshot=self.snapshot)


def get_volume_snapshots_by_id(snapshot_id: Optional[builtins.str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeSnapshotsByIdProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str snapshot_id: Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
    """
    __args__ = dict()
    __args__['snapshotId'] = snapshot_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:volumes/v2:getVolumeSnapshotsById', __args__, opts=opts, typ=GetVolumeSnapshotsByIdProperties).value

    return AwaitableGetVolumeSnapshotsByIdProperties(
        snapshot=pulumi.get(__ret__, 'snapshot'))
def get_volume_snapshots_by_id_output(snapshot_id: Optional[pulumi.Input[builtins.str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVolumeSnapshotsByIdProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str snapshot_id: Either the ID of an existing snapshot. This will be an integer for a Droplet snapshot or a string for a volume snapshot.
    """
    __args__ = dict()
    __args__['snapshotId'] = snapshot_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:volumes/v2:getVolumeSnapshotsById', __args__, opts=opts, typ=GetVolumeSnapshotsByIdProperties)
    return __ret__.apply(lambda __response__: GetVolumeSnapshotsByIdProperties(
        snapshot=pulumi.get(__response__, 'snapshot')))
