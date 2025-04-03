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
    'ListFunctionsTriggersItems',
    'AwaitableListFunctionsTriggersItems',
    'list_functions_triggers',
    'list_functions_triggers_output',
]

@pulumi.output_type
class ListFunctionsTriggersItems:
    def __init__(__self__, triggers=None):
        if triggers and not isinstance(triggers, list):
            raise TypeError("Expected argument 'triggers' to be a list")
        pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[Sequence['outputs.TriggerInfo']]:
        return pulumi.get(self, "triggers")


class AwaitableListFunctionsTriggersItems(ListFunctionsTriggersItems):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListFunctionsTriggersItems(
            triggers=self.triggers)


def list_functions_triggers(namespace_id: Optional[builtins.str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListFunctionsTriggersItems:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str namespace_id: The ID of the namespace to be managed.
    """
    __args__ = dict()
    __args__['namespaceId'] = namespace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:functions/v2:listFunctionsTriggers', __args__, opts=opts, typ=ListFunctionsTriggersItems).value

    return AwaitableListFunctionsTriggersItems(
        triggers=pulumi.get(__ret__, 'triggers'))
def list_functions_triggers_output(namespace_id: Optional[pulumi.Input[builtins.str]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ListFunctionsTriggersItems]:
    """
    Use this data source to access information about an existing resource.

    :param builtins.str namespace_id: The ID of the namespace to be managed.
    """
    __args__ = dict()
    __args__['namespaceId'] = namespace_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:functions/v2:listFunctionsTriggers', __args__, opts=opts, typ=ListFunctionsTriggersItems)
    return __ret__.apply(lambda __response__: ListFunctionsTriggersItems(
        triggers=pulumi.get(__response__, 'triggers')))
