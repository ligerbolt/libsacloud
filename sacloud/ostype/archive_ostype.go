// Package ostype is define OS type of SakuraCloud public archive
package ostype

//go:generate stringer -type=ArchiveOSTypes

// ArchiveOSTypes パブリックアーカイブOS種別
type ArchiveOSTypes int

const (
	// CentOS OS種別:CentOS
	CentOS ArchiveOSTypes = iota
	// Ubuntu OS種別:Ubuntu
	Ubuntu
	// Debian OS種別:Debian
	Debian
	// VyOS OS種別:VyOS
	VyOS
	// CoreOS OS種別:CoreOS
	CoreOS
	// Kusanagi OS種別:Kusanagi(CentOS)
	Kusanagi
	// SiteGuard OS種別:SiteGuard(CentOS)
	SiteGuard
	// FreeBSD OS種別:FreeBSD
	FreeBSD
	// Windows2008 OS種別:Windows Server 2008 R2 Datacenter Edition
	Windows2008
	// Windows2008RDS OS種別:Windows Server 2008 R2 for RDS
	Windows2008RDS
	// Windows2008RDSOffice OS種別:Windows Server 2008 R2 for RDS(Office)
	Windows2008RDSOffice
	// Windows2012 OS種別:Windows Server 2012 R2 Datacenter Edition
	Windows2012
	// Windows2012RDS OS種別:Windows Server 2012 R2 for RDS
	Windows2012RDS
	// Windows2012RDSOffice OS種別:Windows Server 2012 R2 for RDS(Office)
	Windows2012RDSOffice
	// Windows2016 OS種別:Windows Server 2016 Datacenter Edition
	Windows2016
	// Custom OS種別:カスタム
	Custom
)
