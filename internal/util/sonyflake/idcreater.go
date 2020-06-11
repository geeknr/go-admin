package sonyflake

import (
	"fmt"
	"time"
)

type IdCreater struct {
	sf        *Sonyflake
	machineID uint16
}

var IdCter IdCreater
var machineID uint16 = 0

func (creater *IdCreater) Init(machID int) {
	t, _ := time.Parse("2006-01-02", "2020-01-01")
	creater.machineID = uint16(machID)

	settings := Settings{
		StartTime:      t,
		MachineID:      GetMachineID,
		CheckMachineID: nil,
	}

	creater.sf = NewSonyflake(settings)
	/*
		for i:=0; i<100; i++ {
			id, err := mgr.sf.NextID()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("id:", id)
		}*/
}

func (creater *IdCreater) Init2(machID int) {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	creater.machineID = uint16(machID)

	settings := Settings{
		StartTime:      t,
		MachineID:      GetMachineID2,
		CheckMachineID: nil,
	}

	creater.sf = NewSonyflake(settings)
	/*
		for i:=0; i<100; i++ {
			id, err := mgr.sf.NextID()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("id:", id)
		}*/
}

func GetMachineID() (uint16, error) {
	return machineID, nil
}

func GetMachineID2() (uint16, error) {
	return IdCter.machineID, nil
}

func (creater *IdCreater) GetNextID() (int64, error) {
	id, err := creater.sf.NextID()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int64(id), nil
}
