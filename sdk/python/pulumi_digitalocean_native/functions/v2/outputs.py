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
from . import outputs

__all__ = [
    'NamespaceInfo',
    'ScheduledDetails',
    'ScheduledDetailsBodyProperties',
    'TriggerInfo',
    'TriggerInfoScheduledRunsProperties',
]

@pulumi.output_type
class NamespaceInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiHost":
            suggest = "api_host"
        elif key == "createdAt":
            suggest = "created_at"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in NamespaceInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        NamespaceInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        NamespaceInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_host: Optional[builtins.str] = None,
                 created_at: Optional[builtins.str] = None,
                 key: Optional[builtins.str] = None,
                 label: Optional[builtins.str] = None,
                 namespace: Optional[builtins.str] = None,
                 region: Optional[builtins.str] = None,
                 updated_at: Optional[builtins.str] = None,
                 uuid: Optional[builtins.str] = None):
        """
        :param builtins.str api_host: The namespace's API hostname. Each function in a namespace is provided an endpoint at the namespace's hostname.
        :param builtins.str created_at: UTC time string.
        :param builtins.str key: A random alpha numeric string. This key is used in conjunction with the namespace's UUID to authenticate 
               a user to use the namespace via `doctl`, DigitalOcean's official CLI.
        :param builtins.str label: The namespace's unique name.
        :param builtins.str namespace: A unique string format of UUID with a prefix fn-.
        :param builtins.str region: The namespace's datacenter region.
        :param builtins.str updated_at: UTC time string.
        :param builtins.str uuid: The namespace's Universally Unique Identifier.
        """
        if api_host is not None:
            pulumi.set(__self__, "api_host", api_host)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="apiHost")
    def api_host(self) -> Optional[builtins.str]:
        """
        The namespace's API hostname. Each function in a namespace is provided an endpoint at the namespace's hostname.
        """
        return pulumi.get(self, "api_host")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        UTC time string.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def key(self) -> Optional[builtins.str]:
        """
        A random alpha numeric string. This key is used in conjunction with the namespace's UUID to authenticate 
        a user to use the namespace via `doctl`, DigitalOcean's official CLI.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def label(self) -> Optional[builtins.str]:
        """
        The namespace's unique name.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        """
        A unique string format of UUID with a prefix fn-.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        """
        The namespace's datacenter region.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        UTC time string.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter
    def uuid(self) -> Optional[builtins.str]:
        """
        The namespace's Universally Unique Identifier.
        """
        return pulumi.get(self, "uuid")


@pulumi.output_type
class ScheduledDetails(dict):
    """
    Trigger details for SCHEDULED type, where body is optional.
    """
    def __init__(__self__, *,
                 cron: builtins.str,
                 body: Optional['outputs.ScheduledDetailsBodyProperties'] = None):
        """
        Trigger details for SCHEDULED type, where body is optional.

        :param builtins.str cron: valid cron expression string which is required for SCHEDULED type triggers.
        :param 'ScheduledDetailsBodyProperties' body: Optional data to be sent to function while triggering the function.
        """
        pulumi.set(__self__, "cron", cron)
        if body is not None:
            pulumi.set(__self__, "body", body)

    @property
    @pulumi.getter
    def cron(self) -> builtins.str:
        """
        valid cron expression string which is required for SCHEDULED type triggers.
        """
        return pulumi.get(self, "cron")

    @property
    @pulumi.getter
    def body(self) -> Optional['outputs.ScheduledDetailsBodyProperties']:
        """
        Optional data to be sent to function while triggering the function.
        """
        return pulumi.get(self, "body")


@pulumi.output_type
class ScheduledDetailsBodyProperties(dict):
    """
    Optional data to be sent to function while triggering the function.
    """
    def __init__(__self__, *,
                 name: Optional[builtins.str] = None):
        """
        Optional data to be sent to function while triggering the function.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        return pulumi.get(self, "name")


@pulumi.output_type
class TriggerInfo(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "createdAt":
            suggest = "created_at"
        elif key == "isEnabled":
            suggest = "is_enabled"
        elif key == "scheduledDetails":
            suggest = "scheduled_details"
        elif key == "scheduledRuns":
            suggest = "scheduled_runs"
        elif key == "updatedAt":
            suggest = "updated_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerInfo. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerInfo.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerInfo.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 created_at: Optional[builtins.str] = None,
                 function: Optional[builtins.str] = None,
                 is_enabled: Optional[builtins.bool] = None,
                 name: Optional[builtins.str] = None,
                 namespace: Optional[builtins.str] = None,
                 scheduled_details: Optional['outputs.ScheduledDetails'] = None,
                 scheduled_runs: Optional['outputs.TriggerInfoScheduledRunsProperties'] = None,
                 type: Optional[builtins.str] = None,
                 updated_at: Optional[builtins.str] = None):
        """
        :param builtins.str created_at: UTC time string.
        :param builtins.str function: Name of function(action) that exists in the given namespace.
        :param builtins.bool is_enabled: Indicates weather the trigger is paused or unpaused.
        :param builtins.str name: The trigger's unique name within the namespace.
        :param builtins.str namespace: A unique string format of UUID with a prefix fn-.
        :param 'ScheduledDetails' scheduled_details: Trigger details for SCHEDULED type, where body is optional.
        :param builtins.str type: String which indicates the type of trigger source like SCHEDULED.
        :param builtins.str updated_at: UTC time string.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if function is not None:
            pulumi.set(__self__, "function", function)
        if is_enabled is not None:
            pulumi.set(__self__, "is_enabled", is_enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if scheduled_details is not None:
            pulumi.set(__self__, "scheduled_details", scheduled_details)
        if scheduled_runs is not None:
            pulumi.set(__self__, "scheduled_runs", scheduled_runs)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[builtins.str]:
        """
        UTC time string.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def function(self) -> Optional[builtins.str]:
        """
        Name of function(action) that exists in the given namespace.
        """
        return pulumi.get(self, "function")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[builtins.bool]:
        """
        Indicates weather the trigger is paused or unpaused.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The trigger's unique name within the namespace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        """
        A unique string format of UUID with a prefix fn-.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="scheduledDetails")
    def scheduled_details(self) -> Optional['outputs.ScheduledDetails']:
        """
        Trigger details for SCHEDULED type, where body is optional.
        """
        return pulumi.get(self, "scheduled_details")

    @property
    @pulumi.getter(name="scheduledRuns")
    def scheduled_runs(self) -> Optional['outputs.TriggerInfoScheduledRunsProperties']:
        return pulumi.get(self, "scheduled_runs")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        """
        String which indicates the type of trigger source like SCHEDULED.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[builtins.str]:
        """
        UTC time string.
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class TriggerInfoScheduledRunsProperties(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "lastRunAt":
            suggest = "last_run_at"
        elif key == "nextRunAt":
            suggest = "next_run_at"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TriggerInfoScheduledRunsProperties. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TriggerInfoScheduledRunsProperties.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TriggerInfoScheduledRunsProperties.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 last_run_at: Optional[builtins.str] = None,
                 next_run_at: Optional[builtins.str] = None):
        """
        :param builtins.str last_run_at: Indicates last run time. null value indicates trigger not run yet.
        :param builtins.str next_run_at: Indicates next run time. null value indicates trigger will not run.
        """
        if last_run_at is not None:
            pulumi.set(__self__, "last_run_at", last_run_at)
        if next_run_at is not None:
            pulumi.set(__self__, "next_run_at", next_run_at)

    @property
    @pulumi.getter(name="lastRunAt")
    def last_run_at(self) -> Optional[builtins.str]:
        """
        Indicates last run time. null value indicates trigger not run yet.
        """
        return pulumi.get(self, "last_run_at")

    @property
    @pulumi.getter(name="nextRunAt")
    def next_run_at(self) -> Optional[builtins.str]:
        """
        Indicates next run time. null value indicates trigger will not run.
        """
        return pulumi.get(self, "next_run_at")


