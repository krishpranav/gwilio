/**
 * @file constant.go
 * @author Krisna Pranav
 * @version 1.0
 * @date 2025-05-09
 *
 * @copyright Copyright (c) 2025 Krisna Pranav
 *
 */

package constant

var gwilioConst = map[string]interface{}{
	"GatewayDialString":"sofia/gateway/pstn_trunk",
	"SipGatwayDialString": "sofia/gateway/pstn_trunk",
	"DefaultVendorAuthId": "TINIYO1SECRET1AUTHID",
	"ExportFsVars": "'tiniyo_accid,parent_call_sid,parent_call_uuid'",
	"NumberExportFsVars": "'tiniyo_accid,parent_call_sid,parent_call_uuid,tiniyo_did_number'",
	"ApiVersion": "2010-04-01",
}

func GetConstant(varName string) interface{} {
	return gwilioConst[varName]
}