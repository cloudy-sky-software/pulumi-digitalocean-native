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
    'Options',
    'AwaitableOptions',
    'list_databases_options',
    'list_databases_options_output',
]

@pulumi.output_type
class Options:
    def __init__(__self__, options=None, version_availability=None):
        if options and not isinstance(options, dict):
            raise TypeError("Expected argument 'options' to be a dict")
        pulumi.set(__self__, "options", options)
        if version_availability and not isinstance(version_availability, dict):
            raise TypeError("Expected argument 'version_availability' to be a dict")
        pulumi.set(__self__, "version_availability", version_availability)

    @property
    @pulumi.getter
    def options(self) -> Optional['outputs.OptionsOptionsProperties']:
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="versionAvailability")
    def version_availability(self) -> Optional['outputs.OptionsVersionAvailabilityProperties']:
        return pulumi.get(self, "version_availability")


class AwaitableOptions(Options):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Options(
            options=self.options,
            version_availability=self.version_availability)


def list_databases_options(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableOptions:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('digitalocean-native:databases/v2:listDatabasesOptions', __args__, opts=opts, typ=Options).value

    return AwaitableOptions(
        options=pulumi.get(__ret__, 'options'),
        version_availability=pulumi.get(__ret__, 'version_availability'))
def list_databases_options_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[Options]:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('digitalocean-native:databases/v2:listDatabasesOptions', __args__, opts=opts, typ=Options)
    return __ret__.apply(lambda __response__: Options(
        options=pulumi.get(__response__, 'options'),
        version_availability=pulumi.get(__response__, 'version_availability')))
