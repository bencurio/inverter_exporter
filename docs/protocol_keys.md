
# GET

The key parameter can be used to pass the value to Home Assistant sensors or Prometheus Exporter.

| Command | Key | Type | UoM | Description | Example |
|:-------:|---|----|:---:|-----------|----|
| `QID` | `SerialNumber` | `int` | `-` | The device serial number | - |
| **Command** | **Key** | **Type** | **UoM** | **Description** | **Example** |
| `QPI` | `ProtocolID` | `int` | `-` | Device Protocol ID | - |
| **Command** | **Key** | **Type** | **UoM** | **Description** | **Example** |
| `QMOD` | `PowerOnMode` | `bool` | `-` | `P` Power on mode | - |
| `QMOD` | `StandbyMode` | `bool` | `-` | `S` Standby mode | - |
| `QMOD` | `LineMode` | `bool` | `-` | `L` Line mode | - |
| `QMOD` | `BatteryMode` | `bool` | `-` | `B` Battery mode | - |
| `QMOD` | `FaultMode` | `bool` | `-` | `F` Fault mode | - |
| `QMOD` | `PowerSavingMode` | `bool` | `-` | `H` Power saving mode | - |
| **Command** | **Key** | **Type** | **UoM** | **Description** | **Example** |
| `QPIGS` | `GridVoltage` | `float` | `V` | Grid voltage | 230.0 |
| `QPIGS` | `GridFrequency` | `float` | `Hz` | Grid frequency | 50.0 |
| `QPIGS` | `ACOutputVoltage` | `float` | `V` | AC output voltage | 230.0 |
| `QPIGS` | `ACOutputFrequency` | `float` | `Hz` | AC output frequency | 50.0 |
| `QPIGS` | `ACOutputApparentPower` | `int` | `VA` | AC output apparent power | 1500 |
| `QPIGS` | `ACOutputActivePower` | `int` | `W` | AC output active power | 1450 |
| `QPIGS` | `BUSVoltage` | `int` | `V` | BUS voltage | 400 |
| `QPIGS` | `BatteryVoltage` | `float` | `V` | Battery voltage | 48.23 |
| `QPIGS` | `BatteryChargingCurrent` | `int` | `A` | Battery charging current | 12 |
| `QPIGS` | `BatteryCapacity` | `int` | `%` | Battery capacity | 083 |
| `QPIGS` | `InverterHeatSinkTemperature` | `int` | `°C` | Inverter heat sink temperature (NTC A/D value for 1~3K) | 28 |
| `QPIGS` | `PVInputCurrentForBattery` | `int` | `A` | PV Input current for battery | 2 |
| `QPIGS` | `PVInputVoltage` | `float` | `V` | PV Input voltage 1 | 120.2 |
| `QPIGS` | `BatteryVoltageFromSCC` | `float` | `V` | Battery voltage from SCC | 48.00 |
| `QPIGS` | `BatteryDischargeCurrent` | `int` | `A` | Battery discharge current | 16 |
| `QPIGS` | `StatusSBUPriority` | `bool` | `-` | Battery voltage from SCC | - |
| `QPIGS` | `StatusConfiguration` | `bool` | `-` | Configuration status (True = Change, False = Unchanged) | - |
| `QPIGS` | `StatusSCCFirmware` | `bool` | `-` | SCC firmware version (True = Updated, False = Unchanged) | - |
| `QPIGS` | `StatusLoad` | `bool` | `-` | Load status (True = Load on, False = Load off) | - |
| `QPIGS` | `StatusBatteryVoltageStdy` | `bool` | `-` | Battery voltage to steady while charging | - |
| `QPIGS` | `StatusCharging` | `bool` | `-` | Charging on/off | - |
| `QPIGS` | `StatusSCCCharging` | `bool` | `-` | SCC charging on/off | - |
| `QPIGS` | `StatusACCharging` | `bool` | `-` | AC charging on/off | - |
| `QPIGS` | `BatteryVoltageOffsetForFansOn` | `int` | `mV` | Battery voltage offset for fans on | - |
| `QPIGS` | `EEPROMVersion` | `int` | `-` | EEPROM version | - |
| `QPIGS` | `PVChargingPower` | `int` | `W` | PV Charging power | - |
| `QPIGS` | `StatusBatteryFloatingCharge` | `bool` | `-` | Flag for charging to floating mode | - |
| `QPIGS` | `StatusSwitchOn` | `bool` | `-` | Switch on | - |
| `QPIGS` | `StatusDustproofInstalled` | `bool` | `-` | Flag for dustproof installed | - |
| **Command** | **Key** | **Type** | **UoM** | **Description** | **Example** |
| `QPIRI` | `GridRatingVoltage` | `float` | `V` | Grid rating voltage | - |
| `QPIRI` | `GridRatingCurrent` | `float` | `A` | Grid rating current | - |
| `QPIRI` | `ACOutputRatingVoltage` | `float` | `V` | AC output rating voltage | - |
| `QPIRI` | `ACOutputRatingFrequency` | `float` | `Hz` | AC output rating frequency | - |
| `QPIRI` | `ACOutputRatingCurrent` | `float` | `A` | AC output rating current | - |
| `QPIRI` | `ACOutputRatingApparentPower` | `int` | `VA` | AC output rating apparent power | 5000 |
| `QPIRI` | `ACOutputRatingActivePower` | `int` | `W` | AC output rating active power | 5000 |
| `QPIRI` | `BatteryRatingVoltage` | `float` | `V` | Battery rating voltage | 00.0 |
| `QPIRI` | `BatteryRechargeVoltage` | `float` | `V` | Battery re-charge voltage | 00.0 |
| `QPIRI` | `BatteryUnderVoltage` | `float` | `V` | Battery under voltage | 00.0 |
| `QPIRI` | `BatteryBulkVoltage` | `float` | `V` | Battery bulk voltage | 00.0 |
| `QPIRI` | `BatteryFloatVoltage` | `float` | `V` | Battery float voltage | 00.0 |
| `QPIRI` | `BatteryType` | `int` | `-` |  Battery type (0: AGM, 1: Flooded, 2: User) | - |
| `QPIRI` | `BatteryTypeAGM` | `bool` | `-` |  Battery type is AGM? | - |
| `QPIRI` | `BatteryTypeFlooded` | `bool` | `-` |  Battery type is Flooded? | - |
| `QPIRI` | `BatteryTypeUser` | `bool` | `-` |  Battery type is User? | - |
| `QPIRI` | `CurrentMaxACChargingCurrent` | `int` | `A` | Current max AC charging current | 00 |
| `QPIRI` | `CurrentMaxChargingCurrent` | `int` | `A` | Current max charging current | - |
| `QPIRI` | `InputVoltageRange` | `int` | `-` | Input voltage range (0: Appliance, 1: UPS) | - |
| `QPIRI` | `InputVoltageRangeAppliance` | `bool` | `-` | Input voltage is Appliance? | - |
| `QPIRI` | `InputVoltageRangeUPS` | `bool` | `-` | Input voltage range is UPS? | - |
| `QPIRI` | `OutputSourcePriority` | `int` | `-` | Output source priority (0: Utility first, 1: Solar first, 2: SBU first) | - |
| `QPIRI` | `OutputSourcePriorityUtilityFirst` | `bool` | `-` | Output source priority is Utility first? | - |
| `QPIRI` | `OutputSourcePrioritySolarFirst` | `bool` | `-` | Output source priority is Solar first? | - |
| `QPIRI` | `OutputSourcePrioritySBUPriority` | `bool` | `-` | Output source priority is SBU first? | - |
| `QPIRI` | `ChargerSourcePirority` | `int` | `-` | Charger source priority (0: Utility first, 1: Solar first, 2: Solar + Utility, 3: Only solar charging permitted if battery voltage not too low) | - |
| `QPIRI` | `ChargerSourcePiroritySolarFirst` | `bool` | `-` | Charger source priority is Utility first? | - |
| `QPIRI` | `ChargerSourcePirorityUtilityAndSolar` | `bool` | `-` | Charger source priority is Solar first? | - |
| `QPIRI` | `ChargerSourcePirorityOnlySolar` | `bool` | `-` | Charger source priority is Only solar? | - |
| `QPIRI` | `ParallelMaxNumber` | `int` | `-` | Parallel max number | - |
| `QPIRI` | `MachineType` | `int` | `-` | Machine type (00: Grid tie, 01: Off grid, 10: Hybrid) | 00 |
| `QPIRI` | `Topology` | `int` | `-` | Topology (0: Transformerless, 1: Transformer) | 0 |
| `QPIRI` | `TopologyTransformerless` | `bool` | `-` | Topology is Transformerless? | 0 |
| `QPIRI` | `TopologyTransformer` | `bool` | `-` | Topology is Transformer? | 0 |
| `QPIRI` | `OutputMode` | `int` | `-` | Output mode (00: single machine output, 01: parallel output, 02: phase 1 of 3 phase output, 03: phase 2 of 3 phase output, 04: phase 3 of 3 phase output) | 00 |
| `QPIRI` | `BatteryRedischargeVoltage` | `float` | `V` | Battery re-discharge voltage | 00.0 |
| `QPIRI` | `PVOKConditionForParallel` | `int` | `-` | PV OK condition for parallel (0: As long as one unit of inverters has connect PV, parallel system will consider PV OK; 1: Only All of inverters have connect PV, parallel system will consider PV OK) | - |
| `QPIRI` | `PVPowerBalance` | `int` | `-` | PV power balance (0: PV input max current will be the max charged current; 1: PV input max power will be the sum of the max charged power and loads power) | - |
| **Command** | **Key** | **Type** | **UoM** | **Description** | **Example** |
| `QPIWS` | `InverterFaultA1` | `bool` | `-` | `Fault` Inverter fault | - |
| `QPIWS` | `BusOverA2` | `bool` | `-` | `Fault` Bus Over | - |
| `QPIWS` | `BusUnderA3` | `bool` | `-` | `Fault` Bus Under | - |
| `QPIWS` | `BusSoftFailA4` | `bool` | `-` | `Fault` Bus Soft Fail | - |
| `QPIWS` | `LineFailA5` | `bool` | `-` | `Warning` Line fail | - |
| `QPIWS` | `OPVShortA6` | `bool` | `-` | `Warning` OPVShort | - |
| `QPIWS` | `InverterVoltageTooLowA7` | `bool` | `-` | `Fault` Inverter voltage too low | - |
| `QPIWS` | `InverterVoltageTooHighA8` | `bool` | `-` | `Fault` Inverter voltage too high | - |
| `QPIWS` | `OverTemperatureA9` | `bool` | `-` | Over temperature | - |
| `QPIWS` | `FanLockedA10` | `bool` | `-` | Fan locked | - |
| `QPIWS` | `BatteryVoltageHighA11` | `bool` | `-` | Battery voltage igh | - |
| `QPIWS` | `BatteryLowAlarmA12` | `bool` | `-` | `Warning` Battery low alarm | - |
| `QPIWS` | `OverchargeA13` | `bool` | `-` | `Fault` Overcharge | - |
| `QPIWS` | `BatteryUnderShutdownA14` | `bool` | `-` | `Warning` Battery under shutdown | - |
| `QPIWS` | `BatteryDeratingA15` | `bool` | `-` | `Warning` Battery derating | - |
| `QPIWS` | `OverLoadA16` | `bool` | `-` | Over load | - |
| `QPIWS` | `EepromFaultA17` | `bool` | `-` | `Warning` Eeprom fault | - |
| `QPIWS` | `InverterOverCurrentA18` | `bool` | `-` | `Fault` Inverter Over Current | - |
| `QPIWS` | `InverterSoftFailA19` | `bool` | `-` | `Fault` Inverter Soft Fail | - |
| `QPIWS` | `SelfTestFailA20` | `bool` | `-` | `Fault` Self Test Fail | - |
| `QPIWS` | `OPDCVoltageOverA21` | `bool` | `-` | `Fault` OP DC Voltage Over | - |
| `QPIWS` | `BatteryOpenA22` | `bool` | `-` | `Fault` Battery connection is open | - |
| `QPIWS` | `CurrentSensorFailA23` | `bool` | `-` | `Fault` Current Sensor Fail | - |
| `QPIWS` | `BatteryShortA24` | `bool` | `-` | `Fault` Battery Short | - |
| `QPIWS` | `PowerLimitA25` | `bool` | `-` | `Warning` Power limit | - |
| `QPIWS` | `PVVoltageHighA26` | `bool` | `-` | `Warning` PV voltage high | - |
| `QPIWS` | `MPPTOverloadFaultA27` | `bool` | `-` | `Warning` MPPT overload fault | - |
| `QPIWS` | `MPPTOverloadWarningA28` | `bool` | `-` | `Warning` MPPT overload warning | - |
| `QPIWS` | `BatteryTooLowToChargeA29` | `bool` | `-` | `Warning` Battery too low to charge | - |

# SET

| Command | Key | Description | Status key |
|---|---|---|---|
| `POP00` | `SetOutputSourcePriorityToUtilityFirst`  | Output source: Utility first | `OutputSourcePriorityUtilityFirst` |
| `POP01` | `SetOutputSourcePriorityToSolarFirst`  | Output source: Solar first | `OutputSourcePrioritySolarFirst` |
| `POP02` | `SetOutputSourcePriorityToSBUPriority`  | Output source: SBU Priority | `OutputSourcePrioritySBUPriority` |
| **Command** | **Key** | **Description** | **Status key** |
| `PCP00` | `SetDeviceChargerPriorityToUtilityFirst`  | Charger priority: Utility first | `ChargerSourcePirorityUtilityFirst` |
| `PCP01` | `SetDeviceChargerPriorityToSolarFirst`  | Charger priority: Solar first | `ChargerSourcePiroritySolarFirst` |
| `PCP02` | `SetDeviceChargerPriorityToSolarAndUtility`  | Charger priority: Solar and Utility | `ChargerSourcePirorityUtilityAndSolar` |
| `PCP03` | `SetDeviceChargerPriorityToOnlySolarCharging`  | Charger priority: Only solar charging | `ChargerSourcePirorityOnlySolar` |
| **Command** | **Key** | **Description** | **Status key** |
| `PGR00` | `SetDeviceGridWorkingRangeToAppliance`  | Grid working range: Appliance | `InputVoltageRangeAppliance` |
| `PGR01` | `SetDeviceGridWorkingRangeToUPS` | Grid working range: UPS | `InputVoltageRangeUPS` |

## SET Responses

| Code | Description |
|---|---|
| `ACK` | Inverter accepts this command, otherwise... |
| `NAK` | responds NAK |

## Raw responses

| Command | Response |
|---|---|
| `QID` `<CRC>` | `(92932007104584` `<CRC>` `<CR>` |
| `QVFW` `<CRC>` | `(VERFW:00041.15` `<CRC>` `<CR>` |
| `QPIRI` `<CRC>` | `(230.0 21.7 230.0 50.0 21.7 5000 5000 48.0 44.0 42.0 58.4 54.0 1 02 040 0 2 3 1 01 0 0 00.0 0 1` `<CRC>` `<CR>` |
| `QFLAG` `<CRC>` | `(EbuxzDajkvy` `<CRC>` `<CR>` |
| `QPIGS` `<CRC>` | `(232.1 50.0 232.1 50.0 2088 2081 041 363 51.70 006 071 0032 01.1 337.5 00.00 00000 00010110 00 00 00395 010` `<CRC>` `<CR>` |
| `QMOD` `<CRC>` | `(L` `<CRC>` `<CR>` |
| `QPIWS` `<CRC>` | `(000000000000000000000000000000000000` `<CRC>` `<CR>` |
| `QDI` `<CRC>` | `(230.0 50.0 0030 42.0 54.0 56.4 46.0 60 0 0 2 0 0 0 0 0 1 1 1 0 1 0 54.0 0 1` `<CRC>` `<CR>` |
| `QMCHGCR` `<CRC>` | `(010 020 030 040 050 060 070 080` `<CRC>` `<CR>` |
| `QMUCHGCR` `<CRC>` | `(002 010 020 030 040 050 060` `<CRC>` `<CR>` |
| `QBEQI` `<CRC>` | `(1 060 030 040 003 58.40 000 120 0 0650` `<CRC>` `<CR>` |

# Resources

- OPTI-Solar: SP Efecto & Brilliant RS232 communication Protocol
- MPP Solar: PIP-HS/MS/MSX & HYBRID V Communication Protocol