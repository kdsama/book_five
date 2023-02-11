package service

import "testing"

func TestUserActivity(t *testing.T) {
	TestSaveUserActivity(t)
}

func TestSaveUserActivity(t *testing.T) {
	TestSaveUserActivityErrors(t)
}
func TestSaveUserActivityErrors(t *testing.T) {

	//Users DoNot Exist
	//List is passed but it donot exist
	//Comment is passed but it donot exist
	//List and Comment passed but both dono exist

}
