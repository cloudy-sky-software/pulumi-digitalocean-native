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
    'GetImageProperties',
    'AwaitableGetImageProperties',
    'get_image',
    'get_image_output',
]

@pulumi.output_type
class GetImageProperties:
    def __init__(__self__, image=None):
        if image and not isinstance(image, dict):
            raise TypeError("Expected argument 'image' to be a dict")
        pulumi.set(__self__, "image", image)

    @property
    @pulumi.getter
    def image(self) -> 'outputs.Image':
        return pulumi.get(self, "image")


class AwaitableGetImageProperties(GetImageProperties):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetImageProperties(
            image=self.image)


def get_image(image_id: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetImageProperties:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str image_id: A unique number (id) or string (slug) used to identify and reference a
           specific image.
           
           **Public** images can be identified by image `id` or `slug`.
           
           **Private** images *must* be identified by image `id`.
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:images/v2:getImage', __args__, opts=opts, typ=GetImageProperties).value

    return AwaitableGetImageProperties(
        image=pulumi.get(__ret__, 'image'))
def get_image_output(image_id: Optional[pulumi.Input[builtins.str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetImageProperties]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str image_id: A unique number (id) or string (slug) used to identify and reference a
           specific image.
           
           **Public** images can be identified by image `id` or `slug`.
           
           **Private** images *must* be identified by image `id`.
    """
    __args__ = dict()
    __args__['imageId'] = image_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:images/v2:getImage', __args__, opts=opts, typ=GetImageProperties)
    return __ret__.apply(lambda __response__: GetImageProperties(
        image=pulumi.get(__response__, 'image')))
