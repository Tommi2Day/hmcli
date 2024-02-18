package cmd

// MockURL is a sample URL to be used for testing
const MockURL = "http://localhost:8080"

// MockToken is a sample token to be used for testing
const MockToken = "1234567890"

// DeviceListTest is a sample response from a devicelist.cgi request
// devicelist.cgi?device_id=4740
const DeviceListTest = `
<deviceList>
<device name="Bewegungsmelder Garage" address="000955699D3D84" ise_id="4740" 
	interface="HmIP-RF" device_type="HmIP-SMO" ready_config="true">
<channel name="Bewegungsmelder Garage:0" type="30" address="000955699D3D84:0" ise_id="4741" 
	direction="UNKNOWN" parent_device="4740" index="0" group_partner="" aes_available="false" transmission_mode="AES" 
	visible="true" ready_config="true" operate="true"/>
<channel name="Bewegungsmelder Garage" type="17" address="000955699D3D84:1" ise_id="4763" 
	direction="SENDER" parent_device="4740" index="1" group_partner="" aes_available="false" transmission_mode="AES" 
	visible="true" ready_config="true" operate="true"/>
<channel name="HmIP-SMO 000955699D3D84:2" type="17" address="000955699D3D84:2" ise_id="7802" 
	direction="RECEIVER" parent_device="4740" index="2" group_partner="" aes_available="false" transmission_mode="AES"
	visible="true" ready_config="true" operate="true"/>
<channel name="HmIP-SMO 000955699D3D84:3" type="26" address="000955699D3D84:3" ise_id="7803" 
	direction="UNKNOWN" parent_device="4740" index="3" group_partner="" aes_available="false" transmission_mode="AES"
	visible="false" ready_config="true" operate="true"/>
</device>
</deviceList>
`

/*
// DeviceListEmptyTest is a sample response from a devicelist.cgi request with empty result
const DeviceListEmptyTest = `
<deviceList/>
`

const DeviceTypeListNotAuthTest = `
<deviceTypeList>
<not_authenticated/>
</deviceTypeList>
`

// DeviceTypeListTest is a sample response from a devicetypelist.cgi request
const DeviceTypeListTest = `
<deviceTypeList>
<deviceType name="HM-RC-Sec4-3" description="HM-RC-4" thumbnailPath="/config/img/deviceList/50/84_hm-rc-4-x_thumb.png"
	imagePath="/config/img/deviceList/250/85_hm-rc-sec4-3.png">
<form type="line" name="arrow_part1" x1="0.312" y1="0.288" x2="0.416" y2="0.288" stroke="0.012"/>
<form type="line" name="arrow_part2" x1="0.312" y1="0.288" x2="0.352" y2="0.248" stroke="0.012"/>
<form type="line" name="arrow_part3" x1="0.312" y1="0.288" x2="0.352" y2="0.328" stroke="0.012"/>
<form type="formset" name="Arrow" formList="arrow_part1,arrow_part2,arrow_part3"/>
<form type="offset" name="1_Arrow" formName="Arrow" x="0.25" y="0.0"/>
<form type="offset" name="2_Arrow" formName="Arrow" x="0.238" y="0.156"/>
<form type="offset" name="3_Arrow" formName="Arrow" x="0.228" y="0.312"/>
<form type="offset" name="4_Arrow" formName="Arrow" x="0.212" y="0.468"/>
<form type="formset" name="1" formList="2_Arrow"/>
<form type="formset" name="2" formList="1_Arrow"/>
<form type="formset" name="3" formList="4_Arrow"/>
<form type="formset" name="4" formList="3_Arrow"/>
<form type="formset" name="1+2" formList="1_Arrow,2_Arrow"/>
<form type="formset" name="3+4" formList="3_Arrow,4_Arrow"/>
</deviceType>
<deviceType name="HmIP-eTRV-C" description="TRV-C" thumbnailPath="/config/img/deviceList/50/188_hmip-etrv-c_thumb.png"
	imagePath="/config/img/deviceList/250/188_hmip-etrv-c.png"/>
<deviceType name="HMW-LC-Bl1-DR" description="HMW-LC-Bl1-DR" thumbnailPath="/config/img/deviceList/50/27_hmw-lc-bl1-dr_thumb.png"
	imagePath="/config/img/deviceList/250/27_hmw-lc-bl1-dr.png">
<form type="rectangle" name="1" x="0.452" y="0.772" width="0.044" height="0.06"/>
<form type="rectangle" name="2" x="0.5" y="0.772" width="0.048" height="0.06"/>
<form type="rectangle" name="3" x="0.452" y="0.388" width="0.096" height="0.06"/>
</deviceType>
<deviceType name="HmIP-WTH-B-2" description="HmIP-WTH-B" thumbnailPath="/config/img/deviceList/50/200_hmip-wth-b_thumb.png"
	imagePath="/config/img/deviceList/250/200_hmip-wth-b.png"/>
<deviceType name="HmIPW-STH" description="HmIPW-STH" thumbnailPath="/config/img/deviceList/50/146_hmip-sth_thumb.png"
	imagePath="/config/img/deviceList/250/146_hmip-sth.png"/>
</deviceTypeList>
`
*/
// NotificationsTest is a sample response from a notification.cgi request
const NotificationsTest = `
<systemNotification>
<notification ise_id="2850" name="BidCos-RF.NEQ0117117:0.STICKY_UNREACH" type="STICKY_UNREACH" timestamp="1704112392"/>
<notification ise_id="4748" name="HmIP-RF.000955699D3D84:0.LOW_BAT" type="LOWBAT" timestamp="1704142392"/>
</systemNotification>
`

// MasterValueTest is a sample response from a mastervalue.cgi request
// mastervalue.cgi?device_id=4740,4763
const MasterValueTest = `
<mastervalue>
<device name="Bewegungsmelder Garage" ise_id="4740" device_type="HmIP-SMO">
<mastervalue name="ARR_TIMEOUT" value="10"/>
<mastervalue name="CYCLIC_BIDI_INFO_MSG_DISCARD_FACTOR" value="1"/>
<mastervalue name="CYCLIC_BIDI_INFO_MSG_DISCARD_VALUE" value="57"/>
<mastervalue name="CYCLIC_INFO_MSG" value="1"/>
<mastervalue name="CYCLIC_INFO_MSG_DIS" value="1"/>
<mastervalue name="CYCLIC_INFO_MSG_DIS_UNCHANGED" value="20"/>
<mastervalue name="CYCLIC_INFO_MSG_OVERDUE_THRESHOLD" value="2"/>
<mastervalue name="DISABLE_MSG_TO_AC" value="0"/>
<mastervalue name="DUTYCYCLE_LIMIT" value="180"/>
<mastervalue name="ENABLE_ROUTING" value="1"/>
<mastervalue name="LOCAL_RESET_DISABLED" value="0"/>
<mastervalue name="LOW_BAT_LIMIT" value="2.200000"/>
</device>
<device name="Bewegungsmelder Garage" ise_id="4763" device_type="MOTIONDETECTOR_TRANSCEIVER">
<mastervalue name="ALARM_MODE_TYPE" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_1" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_2" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_3" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_4" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_5" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_6" value="0"/>
<mastervalue name="ALARM_MODE_ZONE_7" value="0"/>
<mastervalue name="BRIGHTNESS_FILTER" value="7"/>
<mastervalue name="CAPTURE_WITHIN_INTERVAL" value="0"/>
<mastervalue name="COND_TX_THRESHOLD_LO" value="1000"/>
<mastervalue name="EVENT_FILTER_NUMBER" value="2"/>
<mastervalue name="EVENT_FILTER_PERIOD" value="1.500000"/>
<mastervalue name="LED_DISABLE_CHANNELSTATE" value="1"/>
<mastervalue name="MIN_INTERVAL" value="2"/>
<mastervalue name="MOTION_ACTIVE_TIME" value="1"/>
<mastervalue name="PIR_OPERATION_MODE" value="0"/>
<mastervalue name="PIR_SENSITIVITY" value="24"/>
</device>
</mastervalue>
`

/*
// MasterValueErrorTest is a sample response from a mastervalue.cgi request with error
const MasterValueErrorTest = `
<mastervalue>
<device ise_id="2850" error="true">DEVICE NOT FOUND</device>
</mastervalue>
`

// MasterValueChangeTest is a sample response from a mastervaluechange.cgi request
// mastervaluechange.cgi?device_id=4740&name=ARR_TIMEOUT&value=11
const MasterValueChangeTest = `
<mastervalue>
<device name="Bewegungsmelder Garage" ise_id="4740" device_type="HmIP-SMO">
<mastervalue name="ARR_TIMEOUT" value="11"/>
</device>
</mastervalue>
`

// StateTest is a sample response from a state.cgi request
// state.cgi?device_id=4740
const StateTest = `<state>
<device name="Bewegungsmelder Garage" ise_id="4740" unreach="false" config_pending="false">
<channel name="Bewegungsmelder Garage:0" ise_id="4741" lastdpactiontime="1704496651">
<datapoint name="HmIP-RF.000955699D3D84:0.CONFIG_PENDING" type="CONFIG_PENDING" ise_id="4742" value="false" valuetype="2"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.DUTY_CYCLE" type="DUTY_CYCLE" ise_id="4746" value="false" valuetype="2"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.ERROR_CODE" type="ERROR_CODE" ise_id="4747" value="0" valuetype="16"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.LOW_BAT" type="LOW_BAT" ise_id="4748" value="false" valuetype="2"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.OPERATING_VOLTAGE" type="OPERATING_VOLTAGE" ise_id="4752"
	value="2.500000" valuetype="4" valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.RSSI_DEVICE" type="RSSI_DEVICE" ise_id="4753"
	value="-72" valuetype="16" valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.RSSI_PEER" type="RSSI_PEER" ise_id="4754" value="-71" valuetype="16"
	valueunit="" timestamp="1704112433" lasttimestamp="0"/>
<datapoint name="HmIP-RF.000955699D3D84:0.UNREACH" type="UNREACH" ise_id="4755" value="false" valuetype="2"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:0.UPDATE_PENDING" type="UPDATE_PENDING" ise_id="4759" value="false" valuetype="2"
	valueunit="" timestamp="1704112431" lasttimestamp="0"/>
<datapoint name="HmIP-RF.000955699D3D84:0.OPERATING_VOLTAGE_STATUS" type="OPERATING_VOLTAGE_STATUS" ise_id="7798"
	value="0" valuetype="16" valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
</channel>
<channel name="Bewegungsmelder Garage" ise_id="4763" lastdpactiontime="1704484301">
<datapoint name="HmIP-RF.000955699D3D84:1.ILLUMINATION" type="ILLUMINATION" ise_id="4764" value="0.060000" valuetype="4"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:1.MOTION" type="MOTION" ise_id="4765" value="false" valuetype="2"
	valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:1.MOTION_DETECTION_ACTIVE" type="MOTION_DETECTION_ACTIVE" ise_id="4766"
	value="true" valuetype="2" valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
<datapoint name="HmIP-RF.000955699D3D84:1.RESET_MOTION" type="RESET_MOTION" ise_id="4767" value="" valuetype="2"
	valueunit="" timestamp="0" lasttimestamp="0"/>
<datapoint name="HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION" type="CURRENT_ILLUMINATION" ise_id="7799"
	value="633.800000" valuetype="4" valueunit="" timestamp="1704112434" lasttimestamp="1704112433"/>
<datapoint name="HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION_STATUS" type="CURRENT_ILLUMINATION_STATUS" ise_id="7800"
	value="0" valuetype="16" valueunit="" timestamp="1704112434" lasttimestamp="1704112434"/>
<datapoint name="HmIP-RF.000955699D3D84:1.ILLUMINATION_STATUS" type="ILLUMINATION_STATUS" ise_id="7801"
	value="0" valuetype="16" valueunit="" timestamp="1704518761" lasttimestamp="1704512000"/>
</channel>
<channel name="HmIP-SMO 000955699D3D84:2" ise_id="7802" lastdpactiontime="0"/>
<channel name="HmIP-SMO 000955699D3D84:3" ise_id="7803" lastdpactiontime="0"/>
</device>
</state>
`
*/
// StateListTest is a sample response from a statelist.cgi request
const StateListTest = `
<stateList>
<device name="Bewegungsmelder Garage" ise_id="4740" unreach="false" config_pending="false">
<channel name="Bewegungsmelder Garage:0" ise_id="4741" index="0" visible="true" operate="true">
<datapoint name="HmIP-RF.000955699D3D84:0.CONFIG_PENDING" type="CONFIG_PENDING" ise_id="4742" value="false" valuetype="2"
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.DUTY_CYCLE" type="DUTY_CYCLE" ise_id="4746" value="false" valuetype="2" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.ERROR_CODE" type="ERROR_CODE" ise_id="4747" value="0" valuetype="16" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.LOW_BAT" type="LOW_BAT" ise_id="4748" value="false" valuetype="2" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.OPERATING_VOLTAGE" type="OPERATING_VOLTAGE" ise_id="4752" 
	value="2.500000" valuetype="4" valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.RSSI_DEVICE" type="RSSI_DEVICE" ise_id="4753" value="-70" valuetype="16" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.RSSI_PEER" type="RSSI_PEER" ise_id="4754" value="-71" valuetype="16" 
	valueunit="" timestamp="1704112433" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.UNREACH" type="UNREACH" ise_id="4755" value="false" valuetype="2" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.UPDATE_PENDING" type="UPDATE_PENDING" ise_id="4759" 
	value="false" valuetype="2" valueunit="" timestamp="1704112431" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:0.OPERATING_VOLTAGE_STATUS" type="OPERATING_VOLTAGE_STATUS" ise_id="7798" 
	value="0" valuetype="16" valueunit="" timestamp="1704543866" operations="5"/>
</channel>
<channel name="Bewegungsmelder Garage" ise_id="4763" index="1" visible="true" operate="true">
<datapoint name="HmIP-RF.000955699D3D84:1.ILLUMINATION" type="ILLUMINATION" ise_id="4764" 
	value="139.850000" valuetype="4" valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:1.MOTION" type="MOTION" ise_id="4765" value="false" valuetype="2" 
	valueunit="" timestamp="1704543866" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:1.MOTION_DETECTION_ACTIVE" type="MOTION_DETECTION_ACTIVE" ise_id="4766" 
	value="true" valuetype="2" valueunit="" timestamp="1704543866" operations="7"/>
<datapoint name="HmIP-RF.000955699D3D84:1.RESET_MOTION" type="RESET_MOTION" ise_id="4767" value="" valuetype="2" 
	valueunit="" timestamp="0" operations="2"/>
<datapoint name="HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION" type="CURRENT_ILLUMINATION" ise_id="7799" 
	value="633.800000" valuetype="4" valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:1.CURRENT_ILLUMINATION_STATUS" type="CURRENT_ILLUMINATION_STATUS" ise_id="7800" 
	value="0" valuetype="16" valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="HmIP-RF.000955699D3D84:1.ILLUMINATION_STATUS" type="ILLUMINATION_STATUS" ise_id="7801" 
	value="0" valuetype="16" valueunit="" timestamp="1704543866" operations="5"/>
</channel>
<channel name="HmIP-SMO 000955699D3D84:2" ise_id="7802" index="2" visible="true" operate="true"/>
<channel name="HmIP-SMO 000955699D3D84:3" ise_id="7803" index="3" visible="false" operate="true"/>
</device>
<device name="FB Licht studio" ise_id="3816" unreach="false" sticky_unreach="false" config_pending="false">
<channel name="FB Licht studio:0" ise_id="3817" index="0" visible="" operate="">
<datapoint name="BidCos-RF.NEQ0902973:0.UNREACH" type="UNREACH" ise_id="3837" value="false" valuetype="2" 
	valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.STICKY_UNREACH" type="STICKY_UNREACH" ise_id="3833" value="false" valuetype="2" 
	valueunit="" timestamp="1704112434" operations="7"/>
<datapoint name="BidCos-RF.NEQ0902973:0.CONFIG_PENDING" type="CONFIG_PENDING" ise_id="3819" value="false" valuetype="2" 
	valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.LOWBAT" type="LOWBAT" ise_id="3827" value="false" valuetype="2" 
	valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.RSSI_DEVICE" type="RSSI_DEVICE" ise_id="3831" value="-65535" valuetype="16" 
	valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.RSSI_PEER" type="RSSI_PEER" ise_id="3832" value="-65535" valuetype="16" 
	valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.DEVICE_IN_BOOTLOADER" type="DEVICE_IN_BOOTLOADER" ise_id="3823" value="false" 
	valuetype="2" valueunit="" timestamp="1704112434" operations="5"/>
<datapoint name="BidCos-RF.NEQ0902973:0.UPDATE_PENDING" type="UPDATE_PENDING" ise_id="3841" value="false" valuetype="2" 
	valueunit="" timestamp="1704112434" operations="5"/>
</channel>
<channel name="FB Studio Taste 2" ise_id="3845" index="1" visible="true" operate="true">
<datapoint name="BidCos-RF.NEQ0902973:1.PRESS_SHORT" type="PRESS_SHORT" ise_id="3850" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
<datapoint name="BidCos-RF.NEQ0902973:1.PRESS_LONG" type="PRESS_LONG" ise_id="3848" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
</channel>
<channel name="FB Studio Taste 1" ise_id="3851" index="2" visible="true" operate="true">
<datapoint name="BidCos-RF.NEQ0902973:2.PRESS_SHORT" type="PRESS_SHORT" ise_id="3856" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
<datapoint name="BidCos-RF.NEQ0902973:2.PRESS_LONG" type="PRESS_LONG" ise_id="3854" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
</channel>
<channel name="FB Studio Taste 4" ise_id="3857" index="3" visible="true" operate="true">
<datapoint name="BidCos-RF.NEQ0902973:3.PRESS_SHORT" type="PRESS_SHORT" ise_id="3862" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
<datapoint name="BidCos-RF.NEQ0902973:3.PRESS_LONG" type="PRESS_LONG" ise_id="3860" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
</channel>
<channel name="FB Studio Taste 3" ise_id="3863" index="4" visible="true" operate="true">
<datapoint name="BidCos-RF.NEQ0902973:4.PRESS_SHORT" type="PRESS_SHORT" ise_id="3868" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
<datapoint name="BidCos-RF.NEQ0902973:4.PRESS_LONG" type="PRESS_LONG" ise_id="3866" value="" valuetype="2" valueunit="" timestamp="0" operations="6"/>
</channel>
</device>
</stateList>
`

/*
// StateEmptyTest is a sample response from a statelist.cgi request with empty result
const StateEmptyTest = `
<state/>
`

// StateDP4748 is a sample response from a state.cgi request for a single datapoint
const StateDP4748 = `
<state>
<datapoint ise_id="4748" value="false"/>
</state>
`

// StateChangeTest is a sample response from a state.cgi request for a single datapoint
const StateChangeTest = `
<result>
<changed id="4740" new_value="" success="true"/>
</result>
`

// StateChangeEmptyTest is a sample response from a failed state.cgi request
const StateChangeEmptyTest = `
<result/>
`

// StateChangeNotFoundTest is a sample response from a incomplete state.cgi request
const StateChangeNotFoundTest = `
<result>
<not_found/>
<changed id="4740" new_value="" success="true"/>
</result>
`
*/
// RssiTest is a sample response from a rssi.cgi request
const RssiTest = `
<rssiList>
<rssi device="BidCoS-RF" rx="65536" tx="-58"/>
<rssi device="MEQ0481419" rx="-66" tx="-69"/>
<rssi device="MEQ0482843" rx="-66" tx="-60"/>
<rssi device="MEQ1887951" rx="-58" tx="65536"/>
</rssiList>
`

// SysVarListTest is a sample response from a sysvar.cgi request
const SysVarListTest = `
<systemVariables>
<systemVariable name="Alarmzone 1" variable="4" value="" value_list="" ise_id="1233" min="" max="" unit="" type="2" 
	subtype="6" logged="false" visible="true" timestamp="0" value_name_0="nicht ausgelÃ¶st" value_name_1="ausgelÃ¶st" 
	info="Alarmmeldung Alarmzone 1"/>
<systemVariable name="Anwesenheit" variable="1" value="true" value_list="" ise_id="950" min="" max="" unit="" type="2" 
	subtype="2" logged="false" visible="true" timestamp="1704112382" value_name_0="nicht anwesend" value_name_1="anwesend" 
	info="Anwesenheit"/>
<systemVariable name="DutyCycle" variable="32.000000" value="32.000000" value_list="" ise_id="6548" min="-1" max="100" 
	unit="%" type="4" subtype="0" logged="false" visible="true" timestamp="1705231500" value_name_0="" value_name_1="" 
	info="DutyCycle CCU"/>
<systemVariable name="DutyCycle-Alarm" variable="4" value="" value_list="" ise_id="7931" min="" max="" unit="" type="2" 
	subtype="6" logged="false" visible="true" timestamp="0" value_name_0="nicht ausgelöst" value_name_1="ausgelöst" 
	info="DutyCycle 98% (PEQ0626928)"/>
<systemVariable name="DutyCycle-LGW" variable="6.000000" value="6.000000" value_list="" ise_id="8254" min="-1" max="100"
	unit="%" type="4" subtype="0" logged="false" visible="true" timestamp="1705231500" value_name_0="" value_name_1="" 
	info="DutyCycle LGW (MEQ1887951)"/>
<systemVariable name="RF-Gateway-Alarm" variable="4" value="" value_list="" ise_id="7987" min="" max="" unit="" type="2"
	subtype="6" logged="false" visible="true" timestamp="0" value_name_0="nicht ausgelöst" value_name_1="ausgelöst" 
	info="RF-Gateway rasbpi8 (MEQ1887951) not connected"/>
<systemVariable name="WatchDog-Alarm" variable="4" value="" value_list="" ise_id="7794" min="" max="" unit="" type="2" 
	subtype="6" logged="false" visible="true" timestamp="0" value_name_0="nicht ausgelöst" value_name_1="ausgelöst" 
	info="Unclean shutdown or system crash identified"/>
</systemVariables>
`

// SysVarTest is a sample response from a sysvar.cgi request
// sysvar.cgi?ise_id=8254
const SysVarTest = `
<systemVariables>
<systemVariable name="DutyCycle-LGW" variable="6.000000" value="6.000000" value_list="" value_text="" ise_id="8254" 
	min="-1" max="100" unit="%" type="4" subtype="0" timestamp="1705231800" value_name_0="" value_name_1=""/>
</systemVariables>
`

/*
// SysVarTextTest is a sample response from a sysvar.cgi request with text option
// sysvar.cgi?&ise_id=8254&text=true
const SysVarTextTest = `
<systemVariables>
<systemVariable name="DutyCycle-LGW" variable="6.000000" value="6.000000" value_list="" value_text="" ise_id="8254"
	min="-1" max="100" unit="%" type="4" subtype="0" timestamp="1705231980" value_name_0="" value_name_1=""/>
</systemVariables>
`

// SysVarEmptyTest is a sample response from a sysvar.cgi request with empty result
const SysVarEmptyTest = `
<systemVariables/>
`

// RoomListTest is a sample response from a roomlist.cgi request
const RoomListTest = `
<roomList>
<room name="Aussen" ise_id="4255">
<channel ise_id="4763"/>
<channel ise_id="4068"/>
<channel ise_id="4071"/>
<channel ise_id="8995"/>
<channel ise_id="8691"/>
<channel ise_id="8663"/>
</room>
<room name="Bad" ise_id="1229">
<channel ise_id="3076"/>
</room>
</roomList>`
*/
