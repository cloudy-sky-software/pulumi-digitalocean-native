// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AppAlertPhase = {
    Unknown: "UNKNOWN",
    Pending: "PENDING",
    Configuring: "CONFIGURING",
    Active: "ACTIVE",
    Error: "ERROR",
} as const;

export type AppAlertPhase = (typeof AppAlertPhase)[keyof typeof AppAlertPhase];

export const AppAlertProgressStepStatus = {
    Unknown: "UNKNOWN",
    Pending: "PENDING",
    Running: "RUNNING",
    Error: "ERROR",
    Success: "SUCCESS",
} as const;

export type AppAlertProgressStepStatus = (typeof AppAlertProgressStepStatus)[keyof typeof AppAlertProgressStepStatus];

export const AppAlertSpecOperator = {
    UnspecifiedOperator: "UNSPECIFIED_OPERATOR",
    GreaterThan: "GREATER_THAN",
    LessThan: "LESS_THAN",
} as const;

export type AppAlertSpecOperator = (typeof AppAlertSpecOperator)[keyof typeof AppAlertSpecOperator];

export const AppAlertSpecRule = {
    UnspecifiedRule: "UNSPECIFIED_RULE",
    CpuUtilization: "CPU_UTILIZATION",
    MemUtilization: "MEM_UTILIZATION",
    RestartCount: "RESTART_COUNT",
    DeploymentFailed: "DEPLOYMENT_FAILED",
    DeploymentLive: "DEPLOYMENT_LIVE",
    DomainFailed: "DOMAIN_FAILED",
    DomainLive: "DOMAIN_LIVE",
    FunctionsActivationCount: "FUNCTIONS_ACTIVATION_COUNT",
    FunctionsAverageDurationMs: "FUNCTIONS_AVERAGE_DURATION_MS",
    FunctionsErrorRatePerMinute: "FUNCTIONS_ERROR_RATE_PER_MINUTE",
    FunctionsAverageWaitTimeMs: "FUNCTIONS_AVERAGE_WAIT_TIME_MS",
    FunctionsErrorCount: "FUNCTIONS_ERROR_COUNT",
    FunctionsGbRatePerSecond: "FUNCTIONS_GB_RATE_PER_SECOND",
} as const;

export type AppAlertSpecRule = (typeof AppAlertSpecRule)[keyof typeof AppAlertSpecRule];

export const AppAlertSpecWindow = {
    UnspecifiedWindow: "UNSPECIFIED_WINDOW",
    FiveMinutes: "FIVE_MINUTES",
    TenMinutes: "TEN_MINUTES",
    ThirtyMinutes: "THIRTY_MINUTES",
    OneHour: "ONE_HOUR",
} as const;

export type AppAlertSpecWindow = (typeof AppAlertSpecWindow)[keyof typeof AppAlertSpecWindow];

export const AppComponentInstanceBaseInstanceSizeSlug = {
    BasicXxs: "basic-xxs",
    BasicXs: "basic-xs",
    BasicS: "basic-s",
    BasicM: "basic-m",
    ProfessionalXs: "professional-xs",
    ProfessionalS: "professional-s",
    ProfessionalM: "professional-m",
    Professional1l: "professional-1l",
    ProfessionalL: "professional-l",
    ProfessionalXl: "professional-xl",
} as const;

/**
 * The instance size to use for this component. Default: `basic-xxs`
 */
export type AppComponentInstanceBaseInstanceSizeSlug = (typeof AppComponentInstanceBaseInstanceSizeSlug)[keyof typeof AppComponentInstanceBaseInstanceSizeSlug];

export const AppDatabaseSpecEngine = {
    Unset: "UNSET",
    Mysql: "MYSQL",
    Pg: "PG",
    Redis: "REDIS",
} as const;

/**
 * - MYSQL: MySQL
 * - PG: PostgreSQL
 * - REDIS: Redis
 */
export type AppDatabaseSpecEngine = (typeof AppDatabaseSpecEngine)[keyof typeof AppDatabaseSpecEngine];

export const AppDomainSpecMinimumTlsVersion = {
    AppDomainSpecMinimumTlsVersion_12: "1.2",
    AppDomainSpecMinimumTlsVersion_13: "1.3",
} as const;

/**
 * The minimum version of TLS a client application can use to access resources for the domain.  Must be one of the following values wrapped within quotations: `"1.2"` or `"1.3"`.
 */
export type AppDomainSpecMinimumTlsVersion = (typeof AppDomainSpecMinimumTlsVersion)[keyof typeof AppDomainSpecMinimumTlsVersion];

export const AppDomainSpecType = {
    Unspecified: "UNSPECIFIED",
    Default: "DEFAULT",
    Primary: "PRIMARY",
    Alias: "ALIAS",
} as const;

/**
 * - DEFAULT: The default `.ondigitalocean.app` domain assigned to this app
 * - PRIMARY: The primary domain for this app that is displayed as the default in the control panel, used in bindable environment variables, and any other places that reference an app's live URL. Only one domain may be set as primary.
 * - ALIAS: A non-primary domain
 */
export type AppDomainSpecType = (typeof AppDomainSpecType)[keyof typeof AppDomainSpecType];

export const AppJobSpecPropertiesKind = {
    Unspecified: "UNSPECIFIED",
    PreDeploy: "PRE_DEPLOY",
    PostDeploy: "POST_DEPLOY",
    FailedDeploy: "FAILED_DEPLOY",
} as const;

/**
 * - UNSPECIFIED: Default job type, will auto-complete to POST_DEPLOY kind.
 * - PRE_DEPLOY: Indicates a job that runs before an app deployment.
 * - POST_DEPLOY: Indicates a job that runs after an app deployment.
 * - FAILED_DEPLOY: Indicates a job that runs after a component fails to deploy.
 */
export type AppJobSpecPropertiesKind = (typeof AppJobSpecPropertiesKind)[keyof typeof AppJobSpecPropertiesKind];

export const AppRollbackValidationConditionCode = {
    IncompatiblePhase: "incompatible_phase",
    IncompatibleResult: "incompatible_result",
    ExceededRevisionLimit: "exceeded_revision_limit",
    AppPinned: "app_pinned",
    DatabaseConfigConflict: "database_config_conflict",
    RegionConflict: "region_conflict",
    StaticSiteRequiresRebuild: "static_site_requires_rebuild",
    ImageSourceMissingDigest: "image_source_missing_digest",
} as const;

/**
 * A code identifier that represents the failing condition.
 *
 * Failing conditions:
 *   - `incompatible_phase` - indicates that the deployment's phase is not suitable for rollback.
 *   - `incompatible_result` - indicates that the deployment's result is not suitable for rollback.
 *   - `exceeded_revision_limit` - indicates that the app has exceeded the rollback revision limits for its tier.
 *   - `app_pinned` - indicates that there is already a rollback in progress and the app is pinned.
 *   - `database_config_conflict` - indicates that the deployment's database config is different than the current config.
 *   - `region_conflict` - indicates that the deployment's region differs from the current app region.
 *   
 * Warning conditions:
 *   - `static_site_requires_rebuild` - indicates that the deployment contains at least one static site that will require a rebuild.
 *   - `image_source_missing_digest` - indicates that the deployment contains at least one component with an image source that is missing a digest.
 */
export type AppRollbackValidationConditionCode = (typeof AppRollbackValidationConditionCode)[keyof typeof AppRollbackValidationConditionCode];

export const AppSpecRegion = {
    Ams: "ams",
    Nyc: "nyc",
    Fra: "fra",
    Sfo: "sfo",
    Sgp: "sgp",
    Blr: "blr",
    Tor: "tor",
    Lon: "lon",
    Syd: "syd",
} as const;

/**
 * The slug form of the geographical origin of the app. Default: `nearest available`
 */
export type AppSpecRegion = (typeof AppSpecRegion)[keyof typeof AppSpecRegion];

export const AppVariableDefinitionScope = {
    Unset: "UNSET",
    RunTime: "RUN_TIME",
    BuildTime: "BUILD_TIME",
    RunAndBuildTime: "RUN_AND_BUILD_TIME",
} as const;

/**
 * - RUN_TIME: Made available only at run-time
 * - BUILD_TIME: Made available only at build-time
 * - RUN_AND_BUILD_TIME: Made available at both build and run-time
 */
export type AppVariableDefinitionScope = (typeof AppVariableDefinitionScope)[keyof typeof AppVariableDefinitionScope];

export const AppVariableDefinitionType = {
    General: "GENERAL",
    Secret: "SECRET",
} as const;

/**
 * - GENERAL: A plain-text environment variable
 * - SECRET: A secret encrypted environment variable
 */
export type AppVariableDefinitionType = (typeof AppVariableDefinitionType)[keyof typeof AppVariableDefinitionType];

export const AppsDeploymentPhase = {
    Unknown: "UNKNOWN",
    PendingBuild: "PENDING_BUILD",
    Building: "BUILDING",
    PendingDeploy: "PENDING_DEPLOY",
    Deploying: "DEPLOYING",
    Active: "ACTIVE",
    Superseded: "SUPERSEDED",
    Error: "ERROR",
    Canceled: "CANCELED",
} as const;

export type AppsDeploymentPhase = (typeof AppsDeploymentPhase)[keyof typeof AppsDeploymentPhase];

export const AppsDeploymentProgressStepStatus = {
    Unknown: "UNKNOWN",
    Pending: "PENDING",
    Running: "RUNNING",
    Error: "ERROR",
    Success: "SUCCESS",
} as const;

export type AppsDeploymentProgressStepStatus = (typeof AppsDeploymentProgressStepStatus)[keyof typeof AppsDeploymentProgressStepStatus];

export const AppsDomainPhase = {
    Unknown: "UNKNOWN",
    Pending: "PENDING",
    Configuring: "CONFIGURING",
    Active: "ACTIVE",
    Error: "ERROR",
} as const;

export type AppsDomainPhase = (typeof AppsDomainPhase)[keyof typeof AppsDomainPhase];

export const AppsImageSourceSpecRegistryType = {
    DockerHub: "DOCKER_HUB",
    Docr: "DOCR",
} as const;

/**
 * - DOCKER_HUB: The DockerHub container registry type.
 * - DOCR: The DigitalOcean container registry type.
 */
export type AppsImageSourceSpecRegistryType = (typeof AppsImageSourceSpecRegistryType)[keyof typeof AppsImageSourceSpecRegistryType];

export const AppsInstanceSizeCpuType = {
    Unspecified: "UNSPECIFIED",
    Shared: "SHARED",
    Dedicated: "DEDICATED",
} as const;

export type AppsInstanceSizeCpuType = (typeof AppsInstanceSizeCpuType)[keyof typeof AppsInstanceSizeCpuType];
