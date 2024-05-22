package pnumberint

import "strconv"

func GetPnumber(pnumberParam string) (int, error){
	phonenumber, err:=strconv.Atoi(pnumberParam)
	if err != nil{
		return 0, err
	}
	return phonenumber, nil
}