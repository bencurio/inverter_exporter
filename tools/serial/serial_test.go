package serial


//func TestNewCommand(t *testing.T) {
//
//	// region CRC check (with correct CRC)
//
//	var rawCommand = []byte("QPIGS")
//	var exceptedCrc = uint64(47017)
//
//	command, err := NewCommand(rawCommand)
//	if err != nil {
//		t.Errorf("NewCommand: %v", err)
//	}
//
//	if string(rawCommand) == string(command.Raw) {
//		t.Errorf("Command is equal (rawCommand == command.Raw). CRC calc missing?")
//	}
//
//	// +3 for <CRC><CRC><CR>
//	if (len(rawCommand) + 3) != command.Length {
//		t.Errorf("Command length is not equal. CRC calc missing?")
//	}
//
//	if exceptedCrc != command.CRC  {
//		t.Errorf("CRC mismatch")
//	}
//
//	// endregion
//
//	// region CRC check (with incorrect CRC)
//
//	rawCommand = []byte("QPIGS")
//	exceptedCrc = uint64(0)
//
//	command, err = NewCommand(rawCommand)
//	if err != nil {
//		t.Errorf("NewCommand: %v", err)
//	}
//
//	if string(rawCommand) == string(command.Raw) {
//		t.Errorf("Command is equal (rawCommand == command.Raw). CRC calc missing?")
//	}
//
//	if len(rawCommand) == command.Length {
//		t.Errorf("Command length is equal. CRC calc missing?")
//	}
//
//	if command.CRC == exceptedCrc {
//		t.Errorf("CRC check ok, but excepted is fail")
//	}
//
//	// endregion
//}
//
//func TestValidateRead(t *testing.T) {
//
//	// region CRC check (with correct CRC)
//
//	var rawCommand = []byte{0x51, 0x50, 0x49, 0x47, 0x53, 0xB7, 0xA9, 0x0D}
//	var exceptedCrc = uint64(47017)
//
//	validate, err := ValidateRead(rawCommand)
//	if err != nil {
//		t.Errorf("ValidateRead: %v", err)
//	}
//
//	if validate.CRC != exceptedCrc {
//		t.Errorf("CRC check failed")
//	}
//
//	if string(rawCommand) == string(validate.Raw) {
//		t.Errorf("Command is not equal (rawCommand == validate.Raw). CRC calc missing?")
//	}
//
//	// region
//
//	// region CRC check (with incorrect CRC)
//
//	rawCommand = []byte{0x51, 0x50, 0x49, 0x47, 0x53, 0xA9, 0xB7, 0x0D}
//	exceptedCrc = uint64(47017)
//
//	validate, err = ValidateRead(rawCommand)
//	if err == nil {
//		if validate.CRC == exceptedCrc {
//			t.Errorf("CRC check ok, but excepted is fail")
//		}
//	}
//
//	// endregion
//
//	// region Check with missing SERIAL_EOL
//
//	rawCommand = []byte{0x51, 0x50, 0x49, 0x47, 0x53, 0xA9, 0xB7}
//	exceptedCrc = uint64(47017)
//
//	validate, err = ValidateRead(rawCommand)
//	if err == nil {
//		t.Errorf("Passed, but SERIAL_EOL is not in rawCommand")
//	}
//
//	// endregion
//
//}