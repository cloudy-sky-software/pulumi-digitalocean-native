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

__all__ = [
    'ScheduledDetailsBodyPropertiesArgs',
    'ScheduledDetailsBodyPropertiesArgsDict',
    'ScheduledDetailsArgs',
    'ScheduledDetailsArgsDict',
]

MYPY = False

if not MYPY:
    class ScheduledDetailsBodyPropertiesArgsDict(TypedDict):
        """
        Optional data to be sent to function while triggering the function.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
elif False:
    ScheduledDetailsBodyPropertiesArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ScheduledDetailsBodyPropertiesArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Optional data to be sent to function while triggering the function.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


if not MYPY:
    class ScheduledDetailsArgsDict(TypedDict):
        """
        Trigger details for SCHEDULED type, where body is optional.
        """
        cron: pulumi.Input[builtins.str]
        """
        valid cron expression string which is required for SCHEDULED type triggers.
        """
        body: NotRequired[pulumi.Input['ScheduledDetailsBodyPropertiesArgsDict']]
        """
        Optional data to be sent to function while triggering the function.
        """
elif False:
    ScheduledDetailsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ScheduledDetailsArgs:
    def __init__(__self__, *,
                 cron: pulumi.Input[builtins.str],
                 body: Optional[pulumi.Input['ScheduledDetailsBodyPropertiesArgs']] = None):
        """
        Trigger details for SCHEDULED type, where body is optional.

        :param pulumi.Input[builtins.str] cron: valid cron expression string which is required for SCHEDULED type triggers.
        :param pulumi.Input['ScheduledDetailsBodyPropertiesArgs'] body: Optional data to be sent to function while triggering the function.
        """
        pulumi.set(__self__, "cron", cron)
        if body is not None:
            pulumi.set(__self__, "body", body)

    @property
    @pulumi.getter
    def cron(self) -> pulumi.Input[builtins.str]:
        """
        valid cron expression string which is required for SCHEDULED type triggers.
        """
        return pulumi.get(self, "cron")

    @cron.setter
    def cron(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cron", value)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input['ScheduledDetailsBodyPropertiesArgs']]:
        """
        Optional data to be sent to function while triggering the function.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input['ScheduledDetailsBodyPropertiesArgs']]):
        pulumi.set(self, "body", value)


