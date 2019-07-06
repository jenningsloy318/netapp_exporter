1. we can get all performce objects via following command 
  ```
  MainCluster01::> set advanced 

  Warning: These advanced commands are potentially dangerous; use them only when directed to do so by NetApp personnel.
  Do you want to continue? {y|n}: y
  MainCluster01::*> statistics catalog object show 
    ads                         CM object for exporting the Automatic dedupe
                                scheduling(ADS). The different counters tells
                                us the the value of different counters which
                                can be used to debug bythe dev if some
                                exceptions happen or to check if it's
                                preforming normally.
    aggr_efficiency             CM object for exporting the aggregate storage
                                efficiency ratio. The different counters
                                tells us the total physical storage used to
                                the ratio of total logical data stored for
                                volumes, aggregate and in snapshots.
    aggregate                   The Aggregate object reports activity on the
                                disk aggregates. The RAID subsystem exports
                                the physical disks as logical disk units
                                (LUNs). LUNs are grouped into disk aggregates
                                to be carved into logical volumes. The
                                volumes can be exported as disks for SAN
                                protocols or formatted and exported as a file
                                system for the NAS protocols.
    audit_ng                    CM object for exporting audit_ng performance
                                counters
    audit_ng:vserver            CM object for exporting audit_ng performance
                                counters
    cifs                        The CIFS object reports activity of the
                                Common Internet File System protocol
                                subsystem. This is the Microsoft file-sharing
                                protocol that evolved from the Server Message
                                Block (SMB) application layer network
                                protocol to connect PCs to Network Attached
                                Storage devices (NAS). This object reports
                                vserver specific activty for SMB, SMB2 and
                                SMB3 revisions of the CIFS protocol. For
                                information related to network context
                                specific activity see 'cifs_ctx' object. For
                                information related only to SMB, see the
                                'smb1' object for vserver specific activity
                                and 'smb1_ctx' for network context specific
                                activity. For information related only to
                                SMB2/SMB3, see the 'smb2' object for vserver
                                specific activity and 'smb2_ctx' for network
                                context specific activity.
    cifs:node                   The CIFS object reports activity of the
                                Common Internet File System protocol
                                subsystem. This is the Microsoft file-sharing
                                protocol that evolved from the Server Message
                                Block (SMB) application layer network
                                protocol to connect PCs to Network Attached
                                Storage devices (NAS). This object reports
                                vserver specific activty for SMB, SMB2 and
                                SMB3 revisions of the CIFS protocol. For
                                information related to network context
                                specific activity see 'cifs_ctx' object. For
                                information related only to SMB, see the
                                'smb1' object for vserver specific activity
                                and 'smb1_ctx' for network context specific
                                activity. For information related only to
                                SMB2/SMB3, see the 'smb2' object for vserver
                                specific activity and 'smb2_ctx' for network
                                context specific activity.
    cifs:vserver                The CIFS object reports activity of the
                                Common Internet File System protocol           
                                subsystem. This is the Microsoft file-sharing
                                protocol that evolved from the Server Message
                                Block (SMB) application layer network
                                protocol to connect PCs to Network Attached
                                Storage devices (NAS). This object reports
                                vserver specific activty for SMB, SMB2 and
                                SMB3 revisions of the CIFS protocol. For
                                information related to network context
                                specific activity see 'cifs_ctx' object. For
                                information related only to SMB, see the
                                'smb1' object for vserver specific activity
                                and 'smb1_ctx' for network context specific
                                activity. For information related only to
                                SMB2/SMB3, see the 'smb2' object for vserver
                                specific activity and 'smb2_ctx' for network
                                context specific activity.
    cifs_ctx                    The CIFS object reports activity of the
                                Common Internet File System protocol
                                subsystem. This is the Microsoft file-sharing
                                protocol that evolved from the Server Message
                                Block (SMB) application layer network
                                protocol to connect PCs to Network Attached
                                Storage devices (NAS). This object reports
                                network context specific activty for SMB,
                                SMB2 and SMB3 revisions of the CIFS protocol.
                                For information related to vserver specific
                                activity see 'cifs' object. For information
                                related only to SMB, see the 'smb1' object
                                for vserver specific activity and 'smb1_ctx'
                                for network context specific activity. For
                                information related only to SMB2/SMB3, see
                                the 'smb2' object for vserver specific
                                activity and 'smb2_ctx' for network context
                                specific activity.
    cifs_ctx:node               The CIFS object reports activity of the
                                Common Internet File System protocol
                                subsystem. This is the Microsoft file-sharing
                                protocol that evolved from the Server Message
                                Block (SMB) application layer network
                                protocol to connect PCs to Network Attached
                                Storage devices (NAS). This object reports
                                network context specific activty for SMB,
                                SMB2 and SMB3 revisions of the CIFS protocol.
                                For information related to vserver specific
                                activity see 'cifs' object. For information
                                related only to SMB, see the 'smb1' object
                                for vserver specific activity and 'smb1_ctx'
                                for network context specific activity. For
                                information related only to SMB2/SMB3, see
                                the 'smb2' object for vserver specific
                                activity and 'smb2_ctx' for network context
                                specific activity.
    client                      These counters report statistics for CIFS and
                                NFS network clients such as number of
                                operations/packets as well as low level per
                                protocol statistics. The counters can be used
                                to determine which network clients are
                                sending/receiving data to the filer.
    client:vserver              These counters report statistics for CIFS and
                                NFS network clients such as number of          
                                operations/packets as well as low level per
                                protocol statistics. The counters can be used
                                to determine which network clients are
                                sending/receiving data to the filer.
    cluster_peer                The cluster peer object contains peer
                                counters.
    cpx                         CM object for exporting CR Posix counters in
                                N-blade for Infinite Volume. The CR Posix
                                layer translates the protocol request to the
                                required operations on the namespace
                                redirector and the data file.
    cpx_op                      CM object for exporting CR Posix (namespace
                                redirection) per op counters
    ctlbe                       These counters report errors in the BSD CAM
                                target layer.
    disk                        CM object for exporting disk performance
                                counters
    disk:constituent            CM object for exporting disk performance
                                counters
    disk:raid_group             CM object for exporting disk performance
                                counters
    ext_cache                   This object lists global WAFL External Cache
                                state information applicable to the entire
                                software subsystem, not specific to a given
                                type of cache.
    ext_cache_obj               This object provides performance metrics and
                                configuration characteristics for a given
                                WAFL External Cache type, such as Flash Cache
                                or the Predictive Cache Statistics simulator.
                                High-level cache behavior can be monitored
                                using these statistics.
    external_service_op         In their typical operation, SecD and Name
                                Services connect with many servers on the
                                network to access services that they provide.
                                The external_service_op object represents a
                                specific operation provided by a service on
                                an external server.
    external_service_op_error   In their typical operation, SecD and Name
                                Services connect with many servers on the
                                network to access services that they provide.
                                The external_service_op_error object
                                represents a specific error that occurred
                                when calling an operation provided by a
                                service of an external server.
    external_service_server     In their typical operation, SecD and Name
                                Services connect with many servers on the
                                network to access services that they provide.
                                The external_server_service object represents
                                a server providing a given service.
    fcache                      FlexCache-origin Remote Volume operations
    fci_itn                     The fci_itn object reports counters
                                associated with each initiator port and
                                target port nexus for Fibre Channel connected
                                disks and tapes.
    fci_port                    The fci_port object reports counters
                                associated with each adapter port for Fibre
                                Channel connected disks and tapes.
    fcp_lif                     An FCP LIF is a logical interface that
                                connects a Vserver to a physical FCP port.
                                This object collects diagnostics and           
                                performance information for an FCP LIF. The
                                object counters can be used to debug and
                                diagnose connectivity issues with initiators
                                or fabric on a LIF. The object counters can
                                be used to debug and diagnose connectivity
                                issues with initiators or fabric at a LIF
                                level granularity.
    fcp_lif:node                An FCP LIF is a logical interface that
                                connects a Vserver to a physical FCP port.
                                This object collects diagnostics and
                                performance information for an FCP LIF. The
                                object counters can be used to debug and
                                diagnose connectivity issues with initiators
                                or fabric on a LIF. The object counters can
                                be used to debug and diagnose connectivity
                                issues with initiators or fabric at a LIF
                                level granularity.
    fcp_lif:port                An FCP LIF is a logical interface that
                                connects a Vserver to a physical FCP port.
                                This object collects diagnostics and
                                performance information for an FCP LIF. The
                                object counters can be used to debug and
                                diagnose connectivity issues with initiators
                                or fabric on a LIF. The object counters can
                                be used to debug and diagnose connectivity
                                issues with initiators or fabric at a LIF
                                level granularity.
    fcp_lif:vserver             An FCP LIF is a logical interface that
                                connects a Vserver to a physical FCP port.
                                This object collects diagnostics and
                                performance information for an FCP LIF. The
                                object counters can be used to debug and
                                diagnose connectivity issues with initiators
                                or fabric on a LIF. The object counters can
                                be used to debug and diagnose connectivity
                                issues with initiators or fabric at a LIF
                                level granularity.
    fcp_port                    An FCP target port is a hardware endpoint
                                that performs data communications over a
                                physical link using the fibre channel
                                protocol (FCP). This object collects
                                diagnostic and performance information for an
                                FCP target port. The object counters can be
                                used to diagnose link and connectivity issues
                                on the port. These object counters also
                                provide a high level overview of the ports IO
                                performance.
    fcp_port:node               An FCP target port is a hardware endpoint
                                that performs data communications over a
                                physical link using the fibre channel
                                protocol (FCP). This object collects
                                diagnostic and performance information for an
                                FCP target port. The object counters can be
                                used to diagnose link and connectivity issues
                                on the port. These object counters also
                                provide a high level overview of the ports IO
                                performance.
    fileservices_audit          This object is used to monitor the
                                performance statistics for audit user space
                                process. This process consolidates records     
                                from different nodes into a single final
                                output record file.
    fileservices_audit:vserver  This object is used to monitor the
                                performance statistics for audit user space
                                process. This process consolidates records
                                from different nodes into a single final
                                output record file.
    ha                          Counters for Takeover and Giveback statistics.
    hashd                       The hashd object provides counters to measure
                                the performance of the BranchCache hash
                                daemon.
    hostadapter                 The hostadapter object reports activity on
                                the Fibre Channel, Serial Attached SCSI, and
                                parallel SCSI host adapters the storage
                                system uses to connect to disks and tape
                                drives.
    ic_error_stats              This object displays HA-Interconnect error
                                statistics.
    ic_viif_stats               This object displays VI_IF interconnect
                                performance statistics related to 16 VI_IF
                                channels. The instance of the object
                                represents a VI_IF channel.
    icmp                        These counters report ICMP networking
                                activities.
    icmp6                       These counters report networking statistics
                                for ICMP6 packets which are used for control
                                traffic in IPv6.
    igmp                        These counters report IGMP networking
                                activities.
    ip                          These counters report IP networking
                                activities.
    ip6                         Counters for 'ip6' object report statistics
                                for IPv6 data and control traffic from
                                networking perspective.
    iscsi_conn                  CM object for exporting iscsi connection
                                performance counters
    iscsi_conn:session          CM object for exporting iscsi connection
                                performance counters
    iscsi_lif                   CM object for exporting iSCSI LIF performance
                                counters
    iscsi_lif:node              CM object for exporting iSCSI LIF performance
                                counters
    iscsi_lif:vserver           CM object for exporting iSCSI LIF performance
                                counters
    iwarp                       These counters report IP based RDMA
                                activities.RDMA is done through the iWARP
                                protocol.
    lif                         These counters report activity of logical
                                interfaces (LIFs).
    lif:vserver                 These counters report activity of logical
                                interfaces (LIFs).
    lmgr_ng                     CM object for exporting lmgr_ng performance
                                counters
    lun                         This object contains LUN-level SAN counters
                                which are shared between 7-mode and C-mode.
                                These counters are available for every mapped
                                logical unit.
    lun:constituent             This object contains LUN-level SAN counters
                                which are shared between 7-mode and C-mode.
                                These counters are available for every mapped  
                                logical unit.
    lun:node                    This object contains LUN-level SAN counters
                                which are shared between 7-mode and C-mode.
                                These counters are available for every mapped
                                logical unit.
    mcc_storage                 These counters report initiator and target
                                side storage activities.
    mirror_cache                CM object for exporting Cache Mirror
                                performance metrics. Cache Mirror is used for
                                enhancing failover performance by mirroring
                                certain blocks between partner nodes.
    mirror_transport            This object provides performance metrics and
                                configuration characteristics for the
                                transport module of cache mirror. The
                                transport module is used to send data packets
                                between nodes using best effort or guaranteed
                                mechanisms. Transport clients create
                                uni-directional channels that are used to
                                send data between any two nodes in the
                                cluster. Accordingly, a node participates in
                                a channel as a sender or a receiver. This
                                object provides an overall view of the
                                transport module on the node. More detailed
                                counters for each transport channel are
                                available through mirror_transport_channel
                                object and its instances.
    mirror_transport_channel    This object provides performance metrics and
                                configuration characteristics for a given
                                channel in the transport module of cache
                                mirror. The transport module is used to send
                                data packets between nodes using best effort
                                or guaranteed mechanisms. Transport clients
                                create uni-directional channels that are used
                                to send data between any two nodes in the
                                cluster. Accordingly, a node participates in
                                a channel as a sender or a receiver. This
                                object provides per-channel metrics for
                                channel in which a node participates. Overall
                                counters for transport module are available
                                through mirror_transport object.
    mount                       CM object for exporting nfs mount statistics
                                counters. The mount protocol is a protocol
                                related to NFS. It provides the services
                                required to get NFS running. Some of the
                                duties of mount include: looking up server
                                path names, validating user identity, and
                                checking access permissions. THe clients use
                                the mount protocol to get the first file
                                handle.
    mount:constituent           CM object for exporting nfs mount statistics
                                counters. The mount protocol is a protocol
                                related to NFS. It provides the services
                                required to get NFS running. Some of the
                                duties of mount include: looking up server
                                path names, validating user identity, and
                                checking access permissions. THe clients use
                                the mount protocol to get the first file
                                handle.
    mount:cpu                   CM object for exporting nfs mount statistics
                                counters. The mount protocol is a protocol     
                                related to NFS. It provides the services
                                required to get NFS running. Some of the
                                duties of mount include: looking up server
                                path names, validating user identity, and
                                checking access permissions. THe clients use
                                the mount protocol to get the first file
                                handle.
    mount:node                  CM object for exporting nfs mount statistics
                                counters. The mount protocol is a protocol
                                related to NFS. It provides the services
                                required to get NFS running. Some of the
                                duties of mount include: looking up server
                                path names, validating user identity, and
                                checking access permissions. THe clients use
                                the mount protocol to get the first file
                                handle.
    msrpc_tcp                   These counters report activity from the Msrpc
                                over TCP protocol.
    msrpc_tcp:node              These counters report activity from the Msrpc
                                over TCP protocol.
    msrpc_tcp:vserver           These counters report activity from the Msrpc
                                over TCP protocol.
    nblade_cifs                 The Common Internet File System (CIFS)
                                protocol is an implementation of the Server
                                Message Block (SMB) protocol. It is a
                                standard application layer file system
                                protocol used to share files with Windows(TM)
                                systems. This object tracks the data transfer
                                performance at the CIFS protocol layer, in
                                Ontap's Nblade network component. These
                                counters are relevant to the entire node,
                                rather than individual virtual servers.
    nfsv3                       The NFSv3 object reports activity for the
                                Network File System protocol, version 3. This
                                is the Sun file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv3:constituent           The NFSv3 object reports activity for the
                                Network File System protocol, version 3. This
                                is the Sun file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv3:cpu                   The NFSv3 object reports activity for the
                                Network File System protocol, version 3. This
                                is the Sun file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv3:node                  The NFSv3 object reports activity for the
                                Network File System protocol, version 3. This
                                is the Sun file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv4                       The NFSv4 object reports activity for the
                                Network File System protocol, version 4. This
                                is the ISOC file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv4:constituent           The NFSv4 object reports activity for the
                                Network File System protocol, version 4. This
                                is the ISOC file-sharing protocol that is      
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv4:cpu                   The NFSv4 object reports activity for the
                                Network File System protocol, version 4. This
                                is the ISOC file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv4:node                  The NFSv4 object reports activity for the
                                Network File System protocol, version 4. This
                                is the ISOC file-sharing protocol that is
                                predominant on UNIX platforms, used to
                                connect to Network Attached Storage (NAS).
    nfsv4_1                     The NFSv4.1 object reports activity for the
                                Network File System protocol, version 4.1.
                                This is the file-sharing protocol that
                                implements Parallel NFS (pFNS), used to
                                connect to Network Attached Storage (NAS).
    nfsv4_1:constituent         The NFSv4.1 object reports activity for the
                                Network File System protocol, version 4.1.
                                This is the file-sharing protocol that
                                implements Parallel NFS (pFNS), used to
                                connect to Network Attached Storage (NAS).
    nfsv4_1:cpu                 The NFSv4.1 object reports activity for the
                                Network File System protocol, version 4.1.
                                This is the file-sharing protocol that
                                implements Parallel NFS (pFNS), used to
                                connect to Network Attached Storage (NAS).
    nfsv4_1:node                The NFSv4.1 object reports activity for the
                                Network File System protocol, version 4.1.
                                This is the file-sharing protocol that
                                implements Parallel NFS (pFNS), used to
                                connect to Network Attached Storage (NAS).
    nic_bge                     Note: This object is deprecated and will be
                                removed in a future release.  The Broadcom
                                BCM57xx family gigbit Network Interface Card
                                (NIC) is a 10/100/1000 MB copper card found
                                on supported platforms This object tracks
                                hardware network traffic performance and
                                errors for each BGE NIC.
    nic_common                  This object tracks hardware network traffic
                                performance and errors for all supported
                                network interface cards (NIC), such as Intel
                                Niantic, Qlogic CNA, etc.
    nic_e1000                   Note: This object is deprecated and will be
                                removed in a future release.  The Intel e1000
                                Gigabit Ethernet Network Interface Card (NIC)
                                is a 1000baseT card found on supported
                                platforms. This object tracks hardware
                                network traffic performance and errors for
                                each Intel e1000 NIC.
    nic_igbe                    Note: This object is deprecated and will be
                                removed in a future release.  The Intel
                                Gigabit Ethernet Network Interface Card (NIC)
                                is a 1GB card found on supported platforms.
                                This object tracks hardware network traffic
                                performance and errors for each IGBE NIC
    nic_ixgbe                   Note: This object is deprecated and will be
                                removed in a future release.  The Intel IXGBE
                                (Niantic) network interface card is a 10-GB
                                NIC found on supported platforms. This object  
                                tracks hardware network traffic performance
                                and errors for each IXGBE NIC.
    nic_ixl                     Note: This object is deprecated and will be
                                removed in a future release.  The Intel
                                IXL(40G) network interface card is a 40-GB
                                NIC found on supported platforms. This object
                                tracks hardware network traffic performance
                                and errors for each IXL NIC.
    nic_qla                     Note: This object is deprecated and will be
                                removed in a future release.  The Qlogic 8324
                                Network Interface Card (NIC) is a 10GB card
                                found on supported platforms. This object
                                tracks hardware network traffic performance
                                and errors for each QLA NIC.
    nic_qlge                    Note: This object is deprecated and will be
                                removed in a future release.  The Qlogic
                                Schultz Network Interface Card (NIC) is a
                                10GB card found on supported platforms. This
                                object tracks hardware network traffic
                                performance and errors for each Qlogic
                                Schultz NIC.
    nic_t3c                     Note: This object is deprecated and will be
                                removed in a future release.  The Chelsio T3C
                                Network Interface Card (NIC) is a 10GB card
                                found on supported platforms. This object
                                tracks hardware network traffic performance
                                and errors for each T3C NIC.
    nlm                         CM object for exporting nfs nlm statistics
                                counters
    path                        CM object for exporting path performance
                                counters
    policy_group                The policy_group CM object provides
                                information about QoS policy groups. It
                                contain workload statistics counters
                                aggregated by policy-group. These statistics
                                are cluster-scoped.
    policy_group:constituent    The policy_group CM object provides
                                information about QoS policy groups. It
                                contain workload statistics counters
                                aggregated by policy-group. These statistics
                                are cluster-scoped.
    processor                   The processor object exports performance
                                counters for the central processing units of
                                the system.
    processor:node              The processor object exports performance
                                counters for the central processing units of
                                the system.
    qtree                       CM object for exporting qtree performance
                                counters
    raid                        CM object for exporting raid performance
                                counters
    resource                    The resource CM object that provides
                                resource-based utilization information.
    resource_detail             The resource_detail CM object provides per
                                workload resource-based utilization
                                information. A workload is a label given to
                                application traffic that is using the
                                cluster. A resource is a component of the
                                cluster which does work, including CPU, disk,
                                and network. Note: this object returns a very  
                                large number of instances. Querying by
                                instance name may improve response times.
    resource_headroom_aggr      Display message service time variance and
                                message inter-arrival time variance for
                                aggregates in a node.
    resource_headroom_cpu       This object displays message service time
                                variance and message inter-arrival time
                                variance for WAFL, as well as headroom
                                optimal point information for the CPU
                                resource.
    rquota                      CM object for Rquota statistics counters
    sas_port                    The sas_port object reports counters
                                associated with each SAS adapter port.
    sdal                        CM object for exporting MPIO (WAFL MP-Safe IO
                                Layer) statistics
    session_application         This object provides per application session
                                counts.
    session_location            This object provides per location session
                                counts.
    session_user                This object provides per user session counts.
    session_vserver             This object provides per Vserver session
                                counts.
    slc                         CM object for exporting SLC (Storage Location
                                Cache) statistics
    smb1                        These counters report vserver specific
                                activity from the SMB revision of the
                                protocol. For information specific to network
                                context see 'smb1_ctx' object. For
                                information specific to SMB2/SMB3, see 'smb2'
                                object for vserver specific activity and
                                'smb2_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb1:node                   These counters report vserver specific
                                activity from the SMB revision of the
                                protocol. For information specific to network
                                context see 'smb1_ctx' object. For
                                information specific to SMB2/SMB3, see 'smb2'
                                object for vserver specific activity and
                                'smb2_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb1:vserver                These counters report vserver specific
                                activity from the SMB revision of the
                                protocol. For information specific to network
                                context see 'smb1_ctx' object. For
                                information specific to SMB2/SMB3, see 'smb2'
                                object for vserver specific activity and
                                'smb2_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb1_ctx                    These counters report network context
                                specific activity from the SMB revision of
                                the protocol. For information specific to      
                                vserver activity see 'smb1' object. For
                                information specific to SMB2/SMB3, see 'smb2'
                                object for vserver specific activity and
                                'smb2_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb1_ctx:node               These counters report network context
                                specific activity from the SMB revision of
                                the protocol. For information specific to
                                vserver activity see 'smb1' object. For
                                information specific to SMB2/SMB3, see 'smb2'
                                object for vserver specific activity and
                                'smb2_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb2                        These counters report vserver specific
                                activity from SMB2/SMB3 revision of the
                                protocol. For information specific to network
                                context see 'smb2_ctx' object. For
                                information specific to SMB, see 'smb1'
                                object for vserver specific activity and
                                'smb1_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb2:node                   These counters report vserver specific
                                activity from SMB2/SMB3 revision of the
                                protocol. For information specific to network
                                context see 'smb2_ctx' object. For
                                information specific to SMB, see 'smb1'
                                object for vserver specific activity and
                                'smb1_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb2:vserver                These counters report vserver specific
                                activity from SMB2/SMB3 revision of the
                                protocol. For information specific to network
                                context see 'smb2_ctx' object. For
                                information specific to SMB, see 'smb1'
                                object for vserver specific activity and
                                'smb1_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb2_ctx                    These counters report network context
                                specific activity from SMB2/SMB3 revision of
                                the protocol. For information specific to
                                vserver activity see 'smb2' object. For
                                information specific to SMB, see 'smb1'
                                object for vserver specific activity and
                                'smb1_ctx' for network context specific
                                activity. To see an overview across all        
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smb2_ctx:node               These counters report network context
                                specific activity from SMB2/SMB3 revision of
                                the protocol. For information specific to
                                vserver activity see 'smb2' object. For
                                information specific to SMB, see 'smb1'
                                object for vserver specific activity and
                                'smb1_ctx' for network context specific
                                activity. To see an overview across all
                                revisions, see 'cifs' object for vserver
                                specific activity and 'cifs_ctx' for network
                                context specific activity.
    smtape                      CM object for exporting smtape performance
                                counters
    sp_mgmt_counter             Service Processor(SP) management counter
                                object contains counters related to following
                                feature: spcs commands, ssl certificates,
                                alpha user, file transfer module(ftm), SP
                                logfile collection and SP firmware updates.
                                Ontap can perform operations remotely on SP
                                using rpc based spcs commands. These
                                access/operations are secured by ssl
                                certificates and alpha user (an internal
                                account used by Ontap applications to access
                                SP over ssh connection). Files can be
                                transfered securely between ontap and SP
                                using file transfer module. SP logfile module
                                is used to collect logs from SP of a down
                                node in a cluster.SP firmware can be updated
                                via auto-update feature or using ONTAP cli
                                update command. This counter object reports
                                activity for following feature: total, passed
                                and failed spcs commands. Total ssl
                                certificate syncs and number of syncs failed.
                                Total alpha user syncs and number of alpha
                                user syncs failed. Total, passed and failed
                                ftm transfers. Cumulative passed ftm transfer
                                time. Total, passed and failed SP logfiles
                                collections. Cumulative passed splogfile
                                collect time. Total triggred, passed and
                                failed SP updates(includes both auto-update
                                and ONTAP cli updates). Number of
                                auto-updates and ONTAP cli updates triggered.
                                Cumulative passed SP update time. Auto-update
                                toggle count and servprocd process restart
                                count.
    spinhi                      The spinhi object keeps track of various
                                stats related to fileop or filecb (callback)
                                traffic working through the spinhi layer of
                                the system. It includes a breakdown of the
                                type of ops, histograms detailing the latency
                                of the ops, and relevant op counts. The
                                object also contains pertinent information
                                regarding the memory pools for both fileop
                                and filecb requests.
    spinhi_credcache            The spinhi credential cache caches the CIFS
                                and UNIX creds passed in spinnp file-ops,
                                when requested by N-Blade. This counter        
                                manager object exports various performance
                                stats related to spinhi credential cache.
    statusmon                   CM object for exporting nfs sm statistics
                                counters
    striped                     CM object for exporting striped performance
                                counters
    stripedattributes           Note: This object is deprecated and will be
                                removed in a future release.  CM object for
                                exporting stripedattributes performance
                                counters
    stripedfileop               Note: This object is deprecated and will be
                                removed in a future release.  CM object for
                                exporting stripedfileop performance counters
    stripedfileoperrors         Note: This object is deprecated and will be
                                removed in a future release.  CM object for
                                exporting stripedfileoperrors performance
                                counters
    stripedlock                 CM object for exporting stripedlock
                                performance counters
    stripedmemory               CM object for exporting striped memory
                                performance counters. These counters, used by
                                theFlexCache subsystem, monitor the number of
                                objects for performing striped RPCs are
                                available.
    stripedopclient             CM object for exporting stripedopclient
                                performance counters
    stripedopclienterrors       CM object for exporting stripedopclienterrors
                                performance counters. These are now used for
                                monitoring and supporting the client side of
                                operations required to support cross-node
                                operations used by FlexCache.
    stripedopserver             Note: This object is deprecated and will be
                                removed in a future release.  CM object for
                                exporting stripedopserver performance counters
    stripedopservererrors       Note: This object is deprecated and will be
                                removed in a future release.  CM object for
                                exporting stripedopservererrors performance
                                counters These are now used for monitoring
                                and supporting the server side of operations
                                required to support cross-node operations
                                used by FlexCache.
    system                      The System object reports general system
                                activity. This includes global throughput for
                                the main services, I/O latency, and CPU
                                activity.
    system:constituent          The System object reports general system
                                activity. This includes global throughput for
                                the main services, I/O latency, and CPU
                                activity.
    system:node                 The System object reports general system
                                activity. This includes global throughput for
                                the main services, I/O latency, and CPU
                                activity.
    tcp                         These counters report success and error cases
                                in TCP network protocol execution.
    top_client                  Statistically tracked object for identifying
                                the most active clients in Data ONTAP. A
                                client is uniquely identified by its IP
                                Address (or FQDN) and SVM.
    top_file                    Statistically tracked object for identifying   
                                the most actively accessed files in Data
                                ONTAP. A file is uniquely identified by its
                                path, volume, and SVM. The following file
                                types are tracked: regular, internal, and
                                directory. Note: directory instances which
                                are listed as part of this object are due to
                                operations on the directory itself (e.g.
                                listing directory contents), not operations
                                on files within the directory.
    udp                         These counters report UDP protocol datagrams
                                received, encountering an error, dropped, or
                                sent by the system.
    volume                      CM object for exporting volume performance
                                counters
    volume:app_component        CM object for exporting volume performance
                                counters
    volume:application          CM object for exporting volume performance
                                counters
    volume:node                 CM object for exporting volume performance
                                counters
    volume:vserver              CM object for exporting volume performance
                                counters
    wafl_comp_aggr              CM object for exporting Composite Aggregate
                                statistics.
    wafl_comp_aggr:node         CM object for exporting Composite Aggregate
                                statistics.
    wafl_comp_aggr_bin          CM object for exporting Composite Aggregate
                                bin statistics.
    wafl_comp_aggr_node         CM object for exporting node-wide WAFL
                                Composite Aggregate statistics.
    wafl_comp_aggr_vol          CM object for exporting Composit Aggregate
                                volume statistics.
    wafl_comp_aggr_vol:compaggr CM object for exporting Composit Aggregate
                                volume statistics.
    wafl_comp_aggr_vol:node     CM object for exporting Composit Aggregate
                                volume statistics.
    wafl_comp_aggr_vol_bin      CM object for exporting Composit Aggregate
                                bin volume statistics.
    wafl_fastpath               CM object for exporting RAID Fast Path
                                statistics.
    wafl_hya_per_aggr           CM object for exporting Flash Pool
                                per-aggregate statistics.
    wafl_hya_per_vvol           CM object for exporting Flash Pool
                                per-FlexVol volume statistics. This objects
                                exists for all volumes on a Flash Pool
                                aggregate. System Administrators or NGS
                                personnels that would like to know more about
                                the behavior of a particular volume on a
                                Flash Pool aggregate can query this object.
    witness                     These counters report activity from the CIFS
                                Witness protocol.
    witness:node                These counters report activity from the CIFS
                                Witness protocol.
    witness:vserver             These counters report activity from the CIFS
                                Witness protocol.
    workload                    A workload represents work being done on
                                behalf of an application or system process.
                                The workload CM object reports information
                                such as operations per second, a breakdown of
                                where read operations are going, the           
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information. These statistics
                                illustrate the filer's performance with given
                                workloads.
    workload:constituent        A workload represents work being done on
                                behalf of an application or system process.
                                The workload CM object reports information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information. These statistics
                                illustrate the filer's performance with given
                                workloads.
    workload:policy_group       A workload represents work being done on
                                behalf of an application or system process.
                                The workload CM object reports information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information. These statistics
                                illustrate the filer's performance with given
                                workloads.
    workload_aggr               The workload_aggr CM object provides system
                                and internal workload information specific to
                                storage aggregates.
    workload_detail             The workload_detail CM object that provides
                                service center-based statistical information.
                                Note: this object returns a very large number
                                of instances. Querying by instance name and
                                using wild cards may improve response times.
    workload_detail_volume      The workload_detail_volume CM object provides
                                service center-based statistical information
                                for all volumes. Service centers are resource
                                elements that contributes to the latency of a
                                request. This object provides information on
                                the break down of a volume's response time
                                such as the number of visits, service time
                                and wait time across service centers. Note:
                                this object returns a very large number of
                                instances.
    workload_file_lun           A workload represents work being done on
                                behalf of an application or system process.
                                The workload_file_lun CM object reports
                                information such as read/write operations per
                                second and operation latency per workload.
                                These statistics illustrate the filer's
                                performance with given workloads. This object
                                represents File workloads and LUN workloads
                                only to represent VM vvols.
    workload_file_lun:constituent
                                A workload represents work being done on
                                behalf of an application or system process.
                                The workload_file_lun CM object reports
                                information such as read/write operations per  
                                second and operation latency per workload.
                                These statistics illustrate the filer's
                                performance with given workloads. This object
                                represents File workloads and LUN workloads
                                only to represent VM vvols.
    workload_volume             The workload_volume object provides workload
                                statistics on a per volume basis. Workload
                                information at a file or LUN level are not
                                shown by this object (See the workload
                                object). This object provides information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information.
    workload_volume:app_component
                                The workload_volume object provides workload
                                statistics on a per volume basis. Workload
                                information at a file or LUN level are not
                                shown by this object (See the workload
                                object). This object provides information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information.
    workload_volume:application The workload_volume object provides workload
                                statistics on a per volume basis. Workload
                                information at a file or LUN level are not
                                shown by this object (See the workload
                                object). This object provides information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information.
    workload_volume:constituent The workload_volume object provides workload
                                statistics on a per volume basis. Workload
                                information at a file or LUN level are not
                                shown by this object (See the workload
                                object). This object provides information
                                such as operations per second, a breakdown of
                                where read operations are going, the
                                interarrival time of operation request
                                messages, working set size information,
                                operation latency per workload, and deferred
                                workload information.
    zapi                        The zapi object contains all the zapi
                                counters.
174 entries were displayed.
```