package main

import "fmt"

func main() {

	if result, err := Devide(100, 0); err == "" {
		fmt.Println("100 / 12 = ", result)
	} else {
		fmt.Println("err message :", err)
	}

}

// 定义一个 DivideError 结构
type DevideError struct {
	devidee int
	devider int
}

func (de *DevideError) Error() string { // 实现error这个行为
	strError := `conot process, the devider is zero, 
	devidee is %d;
	devider is 0`
	return fmt.Sprintf(strError, de.devider)
}

func Devide(devider int, devidee int) (result int, errorMsg string) {
	if devidee == 0 {
		dDate := DevideError{
			devider: devider,
			devidee: devidee,
		}
		errorMsg = dDate.Error()
		return
	} else {
		return devider / devidee, ""
	}

}
