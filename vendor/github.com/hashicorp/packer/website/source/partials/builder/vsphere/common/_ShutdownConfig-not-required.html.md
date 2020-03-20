<!-- Code generated from the comments of the ShutdownConfig struct in builder/vsphere/common/step_shutdown.go; DO NOT EDIT MANUALLY -->

-   `shutdown_command` (string) - Specify a VM guest shutdown command. VMware guest tools are used by
    default.
    
-   `shutdown_timeout` (duration string | ex: "1h5m2s") - Amount of time to wait for graceful VM shutdown. Examples 45s and 10m.
    Defaults to 5m(5 minutes).
    