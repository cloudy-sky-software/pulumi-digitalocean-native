# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from ... import _utilities
from ._enums import *

__all__ = [
    'ResourcesItemPropertiesArgs',
    'ResourcesItemPropertiesArgsDict',
]

MYPY = False

if not MYPY:
    class ResourcesItemPropertiesArgsDict(TypedDict):
        resource_id: NotRequired[pulumi.Input[builtins.str]]
        """
        The identifier of a resource.
        """
        resource_type: NotRequired[pulumi.Input['ResourcesItemPropertiesResourceType']]
        """
        The type of the resource.
        """
elif False:
    ResourcesItemPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ResourcesItemPropertiesArgs:
    def __init__(__self__, *,
                 resource_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input['ResourcesItemPropertiesResourceType']] = None):
        """
        :param pulumi.Input[builtins.str] resource_id: The identifier of a resource.
        :param pulumi.Input['ResourcesItemPropertiesResourceType'] resource_type: The type of the resource.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The identifier of a resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input['ResourcesItemPropertiesResourceType']]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input['ResourcesItemPropertiesResourceType']]):
        pulumi.set(self, "resource_type", value)


