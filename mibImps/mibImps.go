package mibImps

import "gitlab.hexcloud.cn/third-party/GoSNMPServer/mibImps/dismanEventMib"
import "gitlab.hexcloud.cn/third-party/GoSNMPServer/mibImps/ifMib"
import "gitlab.hexcloud.cn/third-party/GoSNMPServer/mibImps/ucdMib"

func init() {
	g_Logger = GoSNMPServer.NewDiscardLogger()
}

var g_Logger GoSNMPServer.ILogger

//SetupLogger Setups Logger for All sub mibs.
func SetupLogger(i GoSNMPServer.ILogger) {
	g_Logger = i
	dismanEventMib.SetupLogger(i)
	ifMib.SetupLogger(i)
	ucdMib.SetupLogger(i)
}

// All function provides a list of common used OID
//    includes part of ucdMib, ifMib, and dismanEventMib
func All() []*GoSNMPServer.PDUValueControlItem {
	toRet := []*GoSNMPServer.PDUValueControlItem{}
	toRet = append(toRet, dismanEventMib.All()...)
	toRet = append(toRet, ifMib.All()...)
	toRet = append(toRet, ucdMib.All()...)
	return toRet
}
