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

__all__ = [
    'GetTagProperties',
    'AwaitableGetTagProperties',
    'get_tag',
    'get_tag_output',
]

@pulumi.output_type
class GetTagProperties:
    def __init__(__self__, tag=None):
        if tag and not isinstance(tag, dict):
            raise TypeError("Expected argument 'tag' to be a dict")
        pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def tag(self) -> Optional['outputs.Tags']:
        """
        A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.
        Tags have two attributes: a user defined `name` attribute and an embedded `resources` attribute with information about resources that have been tagged.
        """
        return pulumi.get(self, "tag")


class AwaitableGetTagProperties(GetTagProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagProperties(
            tag=self.tag)


def get_tag(tag_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagProperties:
    """
    Use this data source to access information about an existing resource.

    :param str tag_id: The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag.
    """
    __args__ = dict()
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:tags/v2:getTag', __args__, opts=opts, typ=GetTagProperties).value

    return AwaitableGetTagProperties(
        tag=pulumi.get(__ret__, 'tag'))
def get_tag_output(tag_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTagProperties]:
    """
    Use this data source to access information about an existing resource.

    :param str tag_id: The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag.
    """
    __args__ = dict()
    __args__['tagId'] = tag_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:tags/v2:getTag', __args__, opts=opts, typ=GetTagProperties)
    return __ret__.apply(lambda __response__: GetTagProperties(
        tag=pulumi.get(__response__, 'tag')))
