<!-- Code generated from the comments of the HardwareConfig struct in builder/vsphere/common/step_hardware.go; DO NOT EDIT MANUALLY -->

-   `CPUs` (int32) - Number of CPU sockets.
    
-   `cpu_cores` (int32) - Number of CPU cores per socket.
    
-   `CPU_reservation` (int64) - Amount of reserved CPU resources in MHz.
    
-   `CPU_limit` (int64) - Upper limit of available CPU resources in MHz.
    
-   `CPU_hot_plug` (bool) - Enable CPU hot plug setting for virtual machine. Defaults to `false`.
    
-   `RAM` (int64) - Amount of RAM in MB.
    
-   `RAM_reservation` (int64) - Amount of reserved RAM in MB.
    
-   `RAM_reserve_all` (bool) - Reserve all available RAM. Defaults to `false`. Cannot be used together
    with `RAM_reservation`.
    
-   `RAM_hot_plug` (bool) - Enable RAM hot plug setting for virtual machine. Defaults to `false`.
    
-   `video_ram` (int64) - Amount of video memory in MB.
    
-   `NestedHV` (bool) - Enable nested hardware virtualization for VM. Defaults to `false`.
    