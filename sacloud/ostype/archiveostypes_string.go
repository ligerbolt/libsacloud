// Code generated by "stringer -type=ArchiveOSTypes"; DO NOT EDIT

package ostype

import "fmt"

const _ArchiveOSTypes_name = "CentOSUbuntuDebianVyOSCoreOSKusanagiSiteGuardFreeBSDWindows2008Windows2008RDSWindows2008RDSOfficeWindows2012Windows2012RDSWindows2012RDSOfficeWindows2016Custom"

var _ArchiveOSTypes_index = [...]uint8{0, 6, 12, 18, 22, 28, 36, 45, 52, 63, 77, 97, 108, 122, 142, 153, 159}

func (i ArchiveOSTypes) String() string {
	if i < 0 || i >= ArchiveOSTypes(len(_ArchiveOSTypes_index)-1) {
		return fmt.Sprintf("ArchiveOSTypes(%d)", i)
	}
	return _ArchiveOSTypes_name[_ArchiveOSTypes_index[i]:_ArchiveOSTypes_index[i+1]]
}
